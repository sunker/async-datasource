package asyncquerydata

import (
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type QueryMeta struct {
	QueryFlow string `json:"queryFlow,omitempty"`
}

type queryMeta struct {
	QueryID string `json:"queryID,omitempty"`
	Status  string `json:"status"`
}

type AsyncQuery struct {
	RefID   string    `json:"-"`
	QueryID string    `json:"queryID,omitempty"`
	Meta    QueryMeta `json:"meta,omitempty"`
}

func (q *AsyncQuery) IsAsync() bool {
	return q.Meta.QueryFlow == "async"
}

// GetQuery returns a Query object given a backend.DataQuery using json.Unmarshal
func getAsyncQuery(query backend.DataQuery) (*AsyncQuery, error) {
	model := &AsyncQuery{}

	if err := json.Unmarshal(query.JSON, &model); err != nil {
		return nil, err
	}

	return model, nil
}
