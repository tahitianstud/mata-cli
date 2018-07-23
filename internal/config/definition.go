package config

// Definition describe a configuration
type Definition struct {
	ShowDebugOutput bool `description:"Enable debug-level output" key:"output.debug"`
	VerboseOutput   bool `description:"Enable more output" key:"output.verbose"`
}

// NewConfig will initialize a `Config` data structure with default values
func NewConfig() Definition {
	return Definition{
		ShowDebugOutput: false,
		VerboseOutput:   false,
	}
}
