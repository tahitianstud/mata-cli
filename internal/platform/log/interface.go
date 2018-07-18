package log

// Logger defines the interface to use to interact with errors
type Logger interface {
	DieIf(err error)
	SetLevelTo(level string)
	WriteAt(level string, message string, additionalInfo ...string)
	WriteIf(err error, level string)
}
