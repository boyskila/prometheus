import React, { FC, useState, Fragment } from 'react';
import { Row, Badge } from 'reactstrap';
import CollapsibleAlertPanel from './CollapsibleAlertPanel';
import Checkbox from '../../Checkbox';
import { isPresent } from '../../utils/func';

export type RuleState = keyof RuleStatus<any>;

export interface Rule {
  alerts: Alert[];
  annotations: Record<string, string>;
  duration: number;
  health: string;
  labels: Record<string, string>;
  name: string;
  query: string;
  state: RuleState;
  type: string;
}

export interface RuleStatus<T> {
  firing: T;
  pending: T;
  inactive: T;
}

export interface AlertsProps {
  groups?: RuleGroup[];
  statsCount: RuleStatus<number>;
}

interface Alert {
  labels: Record<string, string>;
  state: RuleState;
  value: string;
  annotations: Record<string, string>;
  activeAt: string;
}

interface RuleGroup {
  name: string;
  file: string;
  rules: Rule[];
  interval: number;
}

const alertKinds: RuleState[] = ['inactive', 'pending', 'firing'];

const AlertsContent: FC<AlertsProps> = ({ groups = [], statsCount }) => {
  const [state, setState] = useState<RuleStatus<boolean>>({
    firing: true,
    pending: true,
    inactive: true,
  });
  const [showAnnotations, setShowAnnotations] = useState(false);

  const toggle = (ruleState: RuleState) => () => {
    setState({
      ...state,
      [ruleState]: !state[ruleState],
    });
  };

  return (
    <>
      <Row className="mb-2">
        <Checkbox
          id="show-annotations"
          wrapperStyles={{ margin: '0 0 0 15px', alignSelf: 'center' }}
          onClick={() => setShowAnnotations(!showAnnotations)}
        >
          Show annotations
        </Checkbox>
        {alertKinds.map((alertKind, i) => {
          return (
            <Checkbox
              defaultChecked
              id={`${alertKind}-toggler`}
              wrapperStyles={{ margin: '0 0 0 15px' }}
              onClick={toggle(alertKind)}
            >
              <Badge color={['success', 'warning', 'danger'][i]} className="text-capitalize">
                {alertKind} ({statsCount[alertKind]})
              </Badge>
            </Checkbox>
          );
        })}
      </Row>
      {groups.map((group, i) => {
        return (
          <Fragment key={i}>
            <StatusBadges rules={group.rules}>
              {group.file} > {group.name}
            </StatusBadges>
            {group.rules.map((rule, j) => {
              return (
                state[rule.state] && (
                  <CollapsibleAlertPanel key={rule.name + j} showAnnotations={showAnnotations} rule={rule} />
                )
              );
            })}
          </Fragment>
        );
      })}
    </>
  );
};

interface StatusBadgesProps {
  rules: Rule[];
}

const StatusBadges: FC<StatusBadgesProps> = ({ rules, children }) => {
  const statesCounter = rules.reduce<any>(
    (acc, r) => {
      return {
        ...acc,
        [r.state]: acc[r.state] + r.alerts.length,
      };
    },
    {
      firing: 0,
      pending: 0,
    }
  );

  return (
    <div className="status-badges border rounded-sm" style={{ lineHeight: 1.1 }}>
      {children}
      <div className="badges-wrapper">
        {isPresent(statesCounter.inactive) && <Badge color="success">inactive</Badge>}
        {statesCounter.pending > 0 && <Badge color="warning">pending ({statesCounter.pending})</Badge>}
        {statesCounter.firing > 0 && <Badge color="danger">firing ({statesCounter.firing})</Badge>}
      </div>
    </div>
  );
};

AlertsContent.displayName = 'Alerts';

export default AlertsContent;
