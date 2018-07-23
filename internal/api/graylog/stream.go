package graylog

const (
	enabledStreamsAPI = "/streams/enabled"
)

// StreamsList describes the result of the listing of streams
type StreamsList struct {
	NumberOfStreams int      `json:"total"`
	Streams         []Stream `json:"streams"`
}

func (sl StreamsList) GetNumberOfStreams() int {
	return sl.NumberOfStreams
}

func (sl StreamsList) GetStreams() []Stream {
	return sl.Streams
}

// Stream describes a Graylog stream
type Stream struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s Stream) GetID() string {
	return s.ID
}

func (s Stream) GetTitle() string {
	return s.Title
}

func (s Stream) GetDescription() string {
	return s.Description
}