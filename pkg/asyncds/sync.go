package asyncds

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func (adq *AsyncQueryDataHandler) handleSyncQuery(ctx context.Context, query backend.DataQuery) (data.Frames, error) {
	asyncQuery, err := getAsyncQuery(query)
	if err != nil {
		return getErrorFrameFromQuery(asyncQuery), err
	}

	queryId, err := adq.provider.StartQuery(ctx, query)
	if err != nil {
		return getErrorFrameFromQuery(asyncQuery), err
	}

	if err := adq.waitOnQuery(ctx, queryId); err != nil {
		return nil, err
	}

	return adq.provider.GetResult(ctx, queryId)
}
