import React from 'react';
import PropTypes from 'prop-types';
import getEntityName from 'modules/getEntityName';
import { entityNameQueryMap } from 'modules/queryMap';
import entityLabels from 'messages/entity';

import Query from 'Components/ThrowingQuery';
import PageHeader from 'Components/PageHeader';
import capitalize from 'lodash/capitalize';
import ExportButton from 'Components/ExportButton';

const getQueryAndVariables = (entityType, entityId) => {
    const query = entityNameQueryMap[entityType] || null;
    return {
        query,
        variables: {
            id: entityId
        }
    };
};

const EntityPageHeader = ({ entityType, entityId, urlParams }) => {
    const { query, variables } = getQueryAndVariables(entityType, entityId);
    if (!query) return null;

    return (
        <Query query={query} variables={variables}>
            {({ data }) => {
                const header = getEntityName(entityType, data) || '-';
                const subHeader = entityLabels[entityType];
                const exportFilename = `${capitalize(
                    entityType.toLowerCase()
                )} Report: "${header}"`;

                let pdfId = 'capture-dashboard-stretch';
                if (urlParams && urlParams.entityListType1) {
                    pdfId = 'capture-list';
                }
                return (
                    <PageHeader classes="bg-primary-100 z-1" header={header} subHeader={subHeader}>
                        <div className="flex flex-1 justify-end">
                            <div className="flex">
                                <div className="flex items-center">
                                    <ExportButton
                                        fileName={exportFilename}
                                        type={entityType}
                                        page="configManagement"
                                        pdfId={pdfId}
                                    />
                                </div>
                            </div>
                        </div>
                    </PageHeader>
                );
            }}
        </Query>
    );
};

EntityPageHeader.propTypes = {
    entityType: PropTypes.string.isRequired,
    entityId: PropTypes.string.isRequired,
    urlParams: PropTypes.shape({})
};

EntityPageHeader.defaultProps = {
    urlParams: null
};

export default EntityPageHeader;
