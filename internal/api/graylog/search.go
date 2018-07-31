package graylog

const (
	absoluteSearchUniversalAPI = "/search/universal/absolute"
	relativeSearchUniversalAPI = "/search/universal/relative"
)

// Definition is the data structure used to describe a search Query
type Search struct {
	Query string `url:"query,omitempty"`
	Range string `url:"range,omitempty"`
	From string `url:"from,omitempty"`
	To string `url:"to,omitempty"`
	Limit string `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
	Filter string `url:"filter,omitempty"`
	Fields string `url:"fields,omitempty"`
	Sort string `url:"sort,omitempty"`
}

// DefaultSearch creates a search definition with default values
func DefaultSearch() Search {
	return Search{
		Query: "*",
		Range: "300",
		Sort: "timestamp:desc",
	}
}

// SearchResult represents the data returned from a search
type SearchResult struct {
	From string
	To string
	Query string
	Messages []interface{}
	Fields []string
	Time int
	TotalResults int `json:"total_results"`
}