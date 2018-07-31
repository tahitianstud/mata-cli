package model

// Config is a struct that will describe the global configuration of the application
type Config struct {
	ShowDebugOutput bool `description:"Enable debug-level output" key:"output.debug"`
	VerboseOutput bool `description:"Enable more output" key:"output.verbose"`
}

// NewConfig will initialize a `Config` data structure with default values
func NewConfig() Config {
	return Config{
		ShowDebugOutput: false,
		VerboseOutput: false,
	}
}