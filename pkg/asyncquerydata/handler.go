package asyncquerydata

import (
	"context"
	"sync"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type AsyncQueryDataProvider interface {
	StartQuery(ctx context.Context, query backend.DataQuery) (string, error)
	GetQueryID(ctx context.Context, query backend.DataQuery) (string, error)
	GetQueryStatus(ctx context.Context, queryId string) (QueryStatus, error)
	CancelQuery(ctx context.Context, queryId string) error
	GetResult(ctx context.Context, refId, queryId string) backend.DataResponse
}

type AsyncQueryDataHandler struct {
	provider AsyncQueryDataProvider
}

func NewAsyncQueryDataHandler(asyncQueryDataProvider AsyncQueryDataProvider) *AsyncQueryDataHandler {
	return &AsyncQueryDataHandler{
		provider: asyncQueryDataProvider,
	}
}

func (ds *AsyncQueryDataHandler) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	firstAsyncQuery, err := getAsyncQuery(req.Queries[0])
	if err != nil {
		return nil, err
	}

	_, isFromAlert := req.Headers["FromAlert"]
	isAsyncMode := firstAsyncQuery.IsAsync() || isFromAlert

	// async flow
	var (
		response = &Response{
			res: backend.NewQueryDataResponse(),
			mtx: &sync.Mutex{},
		}
		wg = sync.WaitGroup{}
	)

	for _, q := range req.Queries {
		wg.Add(1)
		go func(query backend.DataQuery) {
			var dataResponse backend.DataResponse
			if isAsyncMode {
				dataResponse = ds.handleAsyncQuery(ctx, query)
			} else {
				dataResponse = ds.handleSyncQuery(ctx, query)
			}
			response.Set(query.RefID, dataResponse)
			wg.Done()
		}(q)
	}

	wg.Wait()
	return response.Response(), nil
}
