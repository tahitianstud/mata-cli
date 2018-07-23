package query

// Definition is the data structure used to describe a search query
type Definition struct {
	Terms string `description:"search for"`
}

// Service determines what you can do with queries
type Service interface {
	Validate(query Definition) bool
}

// DefaultQuery creates a Query definition with default values
func Default() Definition {
	return Definition{
		Terms: "*",
	}
}
