import React, { FC } from 'react';
import { FilterData } from './Filter';
import { useFetch } from '../../utils/useFetch';
import { groupTargets, Target } from './target';
import ScrapePoolPanel from './ScrapePoolPanel';
import PathPrefixProps from '../../PathPrefixProps';
import { withStatusIndicator } from '../../withStatusIndicator';

interface ScrapePoolListProps {
  filter: FilterData;
  activeTargets: Target[];
}

const ScrapePoolListWithStatusIndicator = withStatusIndicator<ScrapePoolListProps>(({ filter, activeTargets }) => {
  const targetGroups = groupTargets(activeTargets);
  const { showHealthy, showUnhealthy } = filter;
  return (
    <>
      {Object.keys(targetGroups).reduce<JSX.Element[]>((panels, scrapePool) => {
        const targetGroup = targetGroups[scrapePool];
        const isHealthy = targetGroup.upCount === targetGroup.targets.length;
        return (isHealthy && showHealthy) || (!isHealthy && showUnhealthy)
          ? [...panels, <ScrapePoolPanel key={scrapePool} scrapePool={scrapePool} targetGroup={targetGroup} />]
          : panels;
      }, [])}
    </>
  );
});

const ScrapePoolList: FC<{ filter: FilterData } & PathPrefixProps> = ({ pathPrefix, filter }) => {
  const { response, error, isLoading } = useFetch<Partial<ScrapePoolListProps>>(`${pathPrefix}/api/v1/targets?state=active`);
  const { status: responseStatus } = response;
  const badResponseStatus = responseStatus !== 'success' && responseStatus !== 'start fetching';
  return (
    <ScrapePoolListWithStatusIndicator
      {...response.data}
      filter={filter}
      error={badResponseStatus ? new Error() : error}
      isLoading={isLoading}
      customErrorMsg={
        <>
          <strong>Error fetching targets:</strong> {badResponseStatus ? responseStatus : error && error.message}
        </>
      }
    />
  );
};

export default ScrapePoolList;
