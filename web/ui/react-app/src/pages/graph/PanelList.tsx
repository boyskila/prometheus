import React, { Component } from 'react';
import { RouteComponentProps } from '@reach/router';

import { Alert, Button, Col, Row } from 'reactstrap';

import Panel, { PanelOptions, PanelDefaultOptions } from './Panel';
import PathPrefixProps from '../../types/PathPrefixProps';
import { generateID, decodePanelOptionsFromQueryString, encodePanelOptionsToQueryString } from '../../utils';
import LocalStorageListener, { setLocalStorageItem, getLocalStorageItem } from '../../components/LocalStorageHelpers';

export type PanelMeta = { key: string; options: PanelOptions; id: string };

interface PanelListState {
  panels: PanelMeta[];
  metricNames: string[];
  fetchMetricsError: string | null;
  timeDriftError: string | null;
}

class PanelList extends Component<RouteComponentProps & PathPrefixProps, PanelListState> {
  constructor(props: RouteComponentProps & PathPrefixProps) {
    super(props);

    this.state = {
      panels: decodePanelOptionsFromQueryString(window.location.search),
      metricNames: [],
      fetchMetricsError: null,
      timeDriftError: null,
    };
  }

  componentDidMount() {
    !this.state.panels.length && this.addPanel();
    fetch(`${this.props.pathPrefix}/api/v1/label/__name__/values`, { cache: 'no-store', credentials: 'same-origin' })
      .then(resp => {
        if (resp.ok) {
          return resp.json();
        } else {
          throw new Error('Unexpected response status when fetching metric names: ' + resp.statusText); // TODO extract error
        }
      })
      .then(json => {
        this.setState({ metricNames: json.data });
      })
      .catch(error => this.setState({ fetchMetricsError: error.message }));

    const browserTime = new Date().getTime() / 1000;
    fetch(`${this.props.pathPrefix}/api/v1/query?query=time()`, { cache: 'no-store', credentials: 'same-origin' })
      .then(resp => {
        if (resp.ok) {
          return resp.json();
        } else {
          throw new Error('Unexpected response status when fetching metric names: ' + resp.statusText); // TODO extract error
        }
      })
      .then(json => {
        const serverTime = json.data.result[0];
        const delta = Math.abs(browserTime - serverTime);

        if (delta >= 30) {
          throw new Error(
            'Detected ' +
              delta +
              ' seconds time difference between your browser and the server. Prometheus relies on accurate time and time drift might cause unexpected query results.'
          );
        }
      })
      .catch(error => this.setState({ timeDriftError: error.message }));

    window.onpopstate = () => {
      const panels = decodePanelOptionsFromQueryString(window.location.search);
      if (panels.length > 0) {
        this.setState({ panels });
      }
    };
  }

  handleExecuteQuery = (query: string) => {
    const isSimpleMetric = this.state.metricNames.indexOf(query) !== -1;
    if (isSimpleMetric || !query.length) {
      return;
    }
    const historyItems = getLocalStorageItem<string[]>('history', '[]');
    
    const extendedItems = historyItems.reduce(
      (acc, metric) => {
        return metric === query ? acc : [...acc, metric]; // Prevent adding query twice.
      },
      [query]
    )
    setLocalStorageItem('history', extendedItems.slice(0, 50), true)
  };

  updateURL() {
    const query = encodePanelOptionsToQueryString(this.state.panels);
    window.history.pushState({}, '', query);
  }

  handleOptionsChanged = (id: string, options: PanelOptions) => {
    const updatedPanels = this.state.panels.map(p => (id === p.id ? { ...p, options } : p));
    this.setState({ panels: updatedPanels }, this.updateURL);
  };

  addPanel = () => {
    const { panels } = this.state;
    const nextPanels = [
      ...panels,
      {
        id: generateID(),
        key: `${panels.length}`,
        options: PanelDefaultOptions,
      },
    ];
    this.setState({ panels: nextPanels }, this.updateURL);
  };

  removePanel = (id: string) => {
    this.setState(
      {
        panels: this.state.panels.reduce<PanelMeta[]>((acc, panel) => {
          return panel.id !== id ? [...acc, { ...panel, key: `${acc.length}` }] : acc;
        }, []),
      },
      this.updateURL
    );
  };

  render() {
    const { metricNames, timeDriftError, fetchMetricsError, panels } = this.state;
    const { pathPrefix } = this.props;
    return (
      <>
        <Row>
          <Col>
            {timeDriftError && (
              <Alert color="danger">
                <strong>Warning:</strong> Error fetching server time: {timeDriftError}
              </Alert>
            )}
          </Col>
        </Row>
        <Row>
          <Col>
            {fetchMetricsError && (
              <Alert color="danger">
                <strong>Warning:</strong> Error fetching metrics list: {fetchMetricsError}
              </Alert>
            )}
          </Col>
        </Row>
        <LocalStorageListener
          state={{
            'enable-query-history': 'false',
            'use-local-time': 'false',
            history: '[]'
          }}
        >
          {(storage: any) => {
            return (
              <>
                {panels.map(({ id, options }) => (
                  <Panel
                    onExecuteQuery={this.handleExecuteQuery}
                    key={id}
                    options={options}
                    onOptionsChanged={opts => this.handleOptionsChanged(id, opts)}
                    useLocalTime={storage['use-local-time']}
                    removePanel={() => this.removePanel(id)}
                    metricNames={metricNames}
                    pastQueries={storage['enable-query-history'] ? storage['history'] : []}
                    pathPrefix={pathPrefix}
                  />
                ))}
              </>
            )
          }}
        </LocalStorageListener>
        <Button color="primary" className="add-panel-btn" onClick={this.addPanel}>
          Add Panel
        </Button>
      </>
    );
  }
}

export default PanelList;
