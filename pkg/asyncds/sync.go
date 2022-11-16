package asyncds

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func (adq *AsyncQueryDataHandler) handleSyncQuery(ctx context.Context, query backend.DataQuery) backend.DataResponse {
	asyncQuery, err := getAsyncQuery(query)
	if err != nil {
		return getErrorFrameFromQuery(asyncQuery, err)
	}

	queryId, err := adq.provider.StartQuery(ctx, query)
	if err != nil {
		return getErrorFrameFromQuery(asyncQuery, err)
	}

	if err := adq.waitOnQuery(ctx, queryId); err != nil {
		return getErrorFrameFromQuery(asyncQuery, err)
	}

	return adq.provider.GetResult(ctx, query.RefID, queryId)
}
