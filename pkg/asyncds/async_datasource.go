package asyncds

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type QueryStatus uint32

const (
	QueryUnknown QueryStatus = iota
	QuerySubmitted
	QueryRunning
	QueryFinished
	QueryCanceled
	QueryFailed
)

// QueryIDs is a map of RefIDs (Unique Query ID) to query ids.
type QueryIDs map[string]string

type AsyncDatasource interface {
	StartQueries(ctx context.Context, req *backend.QueryDataRequest) (*QueryIDs, error)
	GetQueryIDs(ctx context.Context, req *backend.QueryDataRequest) (*QueryIDs, error)
	QueryStatus(ctx context.Context, queryIDs *QueryIDs) (QueryStatus, error)
	CancelQuery(ctx context.Context, queryIDs *QueryIDs) error
	GetResults(ctx context.Context, queryIDs *QueryIDs) (*backend.QueryDataResponse, error)
}

type AsyncDatasourceRunner struct {
	backend.QueryDataHandler
	AsyncDatasource
	// ///ctx context.Context, req *QueryDataRequest) (*QueryDataResponse, error)
	// StartQuery(ctx context.Context, query string, args ...interface{}) (string, error)
	// GetQueryID(ctx context.Context, query string, args ...interface{}) (bool, string, error)
	// QueryStatus(ctx context.Context, queryID string) (QueryStatus, error)
	// CancelQuery(ctx context.Context, queryID string) error
	// GetRows(ctx context.Context, queryID string) (driver.Rows, error)
}
