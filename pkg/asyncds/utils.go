package asyncds

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// getErrorFrameFromQuery returns a error frames with empty data and meta fields
func getErrorFrameFromQuery(query *AsyncQuery, err error) backend.DataResponse {
	frames := data.Frames{}
	frame := data.NewFrame(query.RefID)
	frame.Meta = &data.FrameMeta{}
	frames = append(frames, frame)
	return backend.DataResponse{
		Frames: frames,
		Error:  err,
	}
}
