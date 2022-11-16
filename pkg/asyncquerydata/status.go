package asyncquerydata

// QueryStatus represents the status of an async query
type QueryStatus uint32

const (
	QueryUnknown QueryStatus = iota
	QuerySubmitted
	QueryRunning
	QueryFinished
	QueryCanceled
	QueryFailed
)

func (qs QueryStatus) Finished() bool {
	return qs == QueryCanceled || qs == QueryFailed || qs == QueryFinished
}

func (qs QueryStatus) String() string {
	switch qs {
	case QuerySubmitted:
		return "submitted"
	case QueryRunning:
		return "running"
	case QueryFinished:
		return "finished"
	case QueryCanceled:
		return "canceled"
	case QueryFailed:
		return "failed"
	default:
		return "unknown"
	}
}
