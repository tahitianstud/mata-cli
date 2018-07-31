package search

// Result represents the data returned from a search
type Result struct {
	From string
	To string
	Query string
	Messages []map[string]interface {}
	Fields []string
	Time int
	TotalResults string `json:"total_results"`
}