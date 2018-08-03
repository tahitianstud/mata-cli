// Package stream describes what a stream of logging messages is.
package stream

// StreamsList describes the result of the listing of streams
type StreamsList struct {
	NumberOfStreams int      `json:"total"`
	Streams         []Stream `json:"streams"`
}

// Stream describes a Graylog Stream
type Stream struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}