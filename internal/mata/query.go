package mata

// Query is the data structure used to describe a search query
type Query struct {
	Terms string `description:"search for"`
}

// QueryService determines what you can do with queries
type QueryService interface {
	Validate(query Query) bool
}

// DefaultQuery creates a Query definition with default values
func DefaultQuery() Query {
	return Query{
		Terms: "*",
	}
}
