package graylog

const (
	enabledStreamsAPI = "/streams/enabled"
)

// StreamsList describes the result of the listing of streams
type StreamsList struct {
	NumberOfStreams int      `json:"total"`
	Streams         []Stream `json:"streams"`
}

// Stream describes a Graylog stream
type Stream struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}