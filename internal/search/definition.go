package search

// Definition is the data structure used to describe a search search
type Definition struct {
	Terms string `description:"search for"`
	Range string `description:"the range from which to start the search on (e.g. 300 means 5 minutes ago)"`
	Stream string
	From string
	To string
}

// DefaultQuery creates a Query definition with default values
func Default() Definition {
	return Definition{
		Terms: "*",
		Range: "300",
		Stream: "",
	}
}
