import ClusterContext from 'common/cluster-context';
import { Script } from 'utils/script-bundle';
import { ScriptsContext } from 'containers/App/scripts-context';
import {
  LIVE_VIEW_SCRIPT_ARGS_KEY, LIVE_VIEW_SCRIPT_ID_KEY, LIVE_VIEW_PIXIE_SCRIPT_KEY,
  LIVE_VIEW_VIS_SPEC_KEY, useSessionStorage,
} from 'common/storage';
import * as React from 'react';
import { withRouter } from 'react-router';

import { parseVis, toJSON, Vis } from 'containers/live/vis';
import { argsEquals, argsForVis, Arguments } from 'utils/args-utils';
import { debounce } from 'utils/debounce';
import urlParams from 'utils/url-params';
import { SetStateFunc } from './common';
import {
  EntityURLParams, getLiveViewTitle, LiveViewEntityParams, LiveViewPage,
  LiveViewPageScriptIds, matchLiveViewEntity, toEntityPathname,
} from '../components/live-widgets/utils/live-view-params';

interface ScriptContextProps {
  liveViewPage: LiveViewPage;

  args: Arguments;
  setArgs: SetStateFunc<Arguments>;

  visJSON: string;
  vis: Vis;
  setVis: SetStateFunc<Vis>;
  setVisJSON: SetStateFunc<string>;

  pxlEditorText: string;
  visEditorText: string;
  setVisEditorText: SetStateFunc<string>;
  setPxlEditorText: SetStateFunc<string>;

  pxl: string;
  setPxl: SetStateFunc<string>;
  title: string;
  id: string;

  setScript: (vis: Vis, pxl: string, args: Arguments, id: string,
    liveViewPage?: LiveViewPage) => void;
  commitURL: () => void;
}

export const ScriptContext = React.createContext<ScriptContextProps>(null);

function emptyVis(): Vis {
  return { variables: [], widgets: [], globalFuncs: [] };
}

function getTitleOfScript(scriptId: string, scripts: Map<string, Script>): string {
  if (scripts.has(scriptId)) {
    return scripts.get(scriptId).title;
  }
  return scriptId;
}

function getEntityParams(liveViewPage: LiveViewPage, args: Arguments): EntityURLParams {
  const entityParamNames = LiveViewEntityParams[liveViewPage] || new Set();
  const entityParams = {};
  entityParamNames.forEach((paramName: string) => {
    if (args[paramName] != null) {
      entityParams[paramName] = args[paramName];
    }
  });
  return entityParams;
}

function getNonEntityParams(liveViewPage: LiveViewPage, args: Arguments): Arguments {
  const entityParamNames = LiveViewEntityParams[liveViewPage] || new Set();
  const nonEntityParams = {};
  Object.keys(args).forEach((argName: string) => {
    if (!entityParamNames.has(argName)) {
      nonEntityParams[argName] = args[argName];
    }
  });
  return nonEntityParams;
}

const ScriptContextProvider = (props) => {
  const { location } = props;

  const { scripts } = React.useContext(ScriptsContext);
  const { selectedClusterName, setClusterByName } = React.useContext(ClusterContext);

  const entity = matchLiveViewEntity(location.pathname);
  const [liveViewPage, setLiveViewPage] = React.useState<LiveViewPage>(entity.page);
  const [visEditorText, setVisEditorText] = React.useState<string>();
  const [pxlEditorText, setPxlEditorText] = React.useState<string>();

  // Args that are not part of an entity.
  const [args, setArgs] = useSessionStorage<Arguments | null>(LIVE_VIEW_SCRIPT_ARGS_KEY, entity.params);

  const [pxl, setPxl] = useSessionStorage(LIVE_VIEW_PIXIE_SCRIPT_KEY, '');
  const [id, setId] = useSessionStorage(LIVE_VIEW_SCRIPT_ID_KEY,
    entity.page === LiveViewPage.Default ? '' : LiveViewPageScriptIds[entity.page]);

  // We use a separation of visJSON states to prevent infinite update loops from happening.
  // Otherwise, an update in the editor could trigger an update to some other piece of the pie.
  // visJSON should be the incoming state to the vis editor.
  // Vis Raw is the outgoing state from the vis editor.
  const [visJSON, setVisJSONBase] = useSessionStorage<string>(LIVE_VIEW_VIS_SPEC_KEY);
  const [vis, setVisBase] = React.useState(() => {
    const parsed = parseVis(visJSON);
    if (parsed) {
      return parsed;
    }
    return emptyVis();
  });

  React.useEffect(() => {
    const newArgs = argsForVis(vis, args);
    // Need this check in order to avoid infinite rerender on 'args' changing.
    if (!argsEquals(newArgs, args)) {
      setArgs(newArgs);
    }
  }, [vis, args, setArgs]);

  // title is dependent on whether or not we are in an entity page.
  const title = React.useMemo(() => {
    const entityParams = getEntityParams(liveViewPage, args);
    const newTitle = getLiveViewTitle(getTitleOfScript(id, scripts), liveViewPage, entityParams);
    document.querySelector('title').textContent = newTitle;
    return newTitle;
  }, [liveViewPage, scripts, id, args]);

  // Logic to set cluster

  React.useEffect(() => {
    if (entity.clusterName && entity.clusterName !== selectedClusterName) {
      setClusterByName(entity.clusterName);
    }
  }, []);

  // Logic to set url params when location changes

  React.useEffect(() => {
    if (location.pathname !== '/live') {
      urlParams.triggerOnChange();
    }
  }, [location]);

  // Logic to update entity paths when live view page or cluster changes.

  React.useEffect(() => {
    const entityParams = getEntityParams(liveViewPage, args);
    const entityURL = {
      clusterName: selectedClusterName,
      page: liveViewPage,
      params: entityParams,
    };
    urlParams.setPathname(toEntityPathname(entityURL));
  }, [selectedClusterName, liveViewPage, args]);

  React.useEffect(() => {
    const scriptID = liveViewPage === LiveViewPage.Default ? id : '';
    urlParams.setScript(scriptID, /* diff */'');
  }, [liveViewPage, id]);

  // Logic to update arguments (which are a combo of entity params and normal args)

  React.useEffect(() => {
    const nonEntityArgs = getNonEntityParams(liveViewPage, args);
    urlParams.setArgs(nonEntityArgs);
  }, [liveViewPage, args]);

  // Logic to update vis spec.

  // Note that this function does not update args, so it should only be called
  // when variables will not be modified (such as for layouts).
  const setVis = (newVis: Vis) => {
    if (!newVis) {
      newVis = emptyVis();
    }
    setVisJSONBase(toJSON(newVis));
    setVisBase(newVis);
  };

  const setVisDebounce = React.useRef(debounce((newJSON: string) => {
    const parsed = parseVis(newJSON);
    if (parsed) {
      setVisBase(parsed);
    }
  }, 2000));

  const setVisJSON = (newJSON: string) => {
    setVisJSONBase(newJSON);
    setVisDebounce.current(newJSON);
  };

  // Logic to update the full script

  const setScript = (newVis: Vis, newPxl: string, newArgs: Arguments, newID: string,
    newLiveViewPage?: LiveViewPage) => {
    setVis(newVis);
    setPxl(newPxl);
    setArgs(newArgs);
    setId(newID);
    if (newLiveViewPage != null) {
      setLiveViewPage(newLiveViewPage);
    }
  };

  // Logic to commit the URL to the history.
  const commitURL = () => {
    // Only show the script as a query arg when we are not on an entity page.
    const scriptId = liveViewPage === LiveViewPage.Default ? (id || '') : '';
    const nonEntityArgs = getNonEntityParams(liveViewPage, args);
    urlParams.commitAll(scriptId, '', nonEntityArgs);
  };

  return (
    <ScriptContext.Provider
      value={{
        liveViewPage,
        args,
        setArgs,
        pxlEditorText,
        visEditorText,
        setPxlEditorText,
        setVisEditorText,
        vis,
        setVis,
        visJSON,
        setVisJSON,
        pxl,
        setPxl,
        title,
        id,
        setScript,
        commitURL,
      }}
    >
      {props.children}
    </ScriptContext.Provider>
  );
};

export default withRouter(ScriptContextProvider);
