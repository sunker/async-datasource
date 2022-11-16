package asyncquerydata

import (
	"context"
	"errors"
	"time"

	// "github.com/graph-gophers/graphql-go/log"
	"github.com/jpillora/backoff"
)

var (
	backoffMin = 200 * time.Millisecond
	backoffMax = 10 * time.Minute
)

func (adq *AsyncQueryDataHandler) waitOnQuery(ctx context.Context, queryId string) error {
	backoffInstance := backoff.Backoff{
		Min:    backoffMin,
		Max:    backoffMax,
		Factor: 1.1,
	}
	for {
		status, err := adq.provider.GetQueryStatus(ctx, queryId)
		if err != nil {
			return err
		}
		if status.Finished() {
			return nil
		}
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.Canceled) {
				err := adq.provider.CancelQuery(ctx, queryId)
				if err != nil {
					return err
				}
			}
			// log.DefaultLogger.Debug("request failed", "query ID", query.RefID, "error", err)
			return err
		case <-time.After(backoffInstance.Duration()):
			continue
		}
	}
}
