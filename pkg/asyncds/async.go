package asyncds

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// handleQuery will call query, and attempt to reconnect if the query failed
func (adq *AsyncQueryDataHandler) handleAsyncQuery(ctx context.Context, query backend.DataQuery) (data.Frames, error) {
	asyncQuery, err := getAsyncQuery(query)
	if err != nil {
		return getErrorFrameFromQuery(asyncQuery), err
	}

	if asyncQuery.QueryID == "" {
		queryID, err := adq.provider.StartQuery(ctx, query)
		if err != nil {
			return getErrorFrameFromQuery(asyncQuery), err
		}
		return data.Frames{
			{Meta: &data.FrameMeta{
				// ExecutedQueryString: q.RawSQL,
				Custom: queryMeta{QueryID: queryID, Status: "started"}},
			},
		}, nil
	}

	status, err := adq.provider.GetQueryStatus(ctx, asyncQuery.QueryID)
	if err != nil {
		return getErrorFrameFromQuery(asyncQuery), err
	}
	customMeta := queryMeta{QueryID: asyncQuery.QueryID, Status: status.String()}
	if !status.Finished() {
		return data.Frames{
			{Meta: &data.FrameMeta{
				// ExecutedQueryString: q.RawSQL,
				Custom: customMeta},
			},
		}, nil
	}

	return adq.provider.GetResult(ctx, asyncQuery.QueryID)
}
