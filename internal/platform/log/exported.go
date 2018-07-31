package log

import (
	"github.com/tahitianstud/mata-cli/internal/platform/log/logrus"
	"fmt"
)


// Level describes the different log levels to deal with
type Level int
const (
	DEBUG Level=0 + iota
	INFO
	WARNING
	ERROR
	FATAL
)



var (
	Log Logger = logrus.GetLogrusHandler()
)




// SetLevelTo determines at what level we are logging
func SetLevelTo(level Level) {

	switch level {

	case DEBUG:
		Log.SetLevelTo("DEBUG")

	case INFO:
		Log.SetLevelTo("INFO")

	case WARNING:
		Log.SetLevelTo("WARN")

	case ERROR:
		Log.SetLevelTo("ERROR")

	case FATAL:
		Log.SetLevelTo("FATAL")
	}
}


// DieIf will exit application if error is not nil
func DieIf(err error) {
	Log.DieIf(err)
}



// DebugIf will raise a debug-level message if err is not nil
func DebugIf(err error) {
	Log.WriteIf(err, "DEBUG")
}


// InfoIf will raise an info-level message if err is not nil
func InfoIf(err error) {
	Log.WriteIf(err, "INFO")
}


// WarningIf will raise a warning-level message if err is not nil
func WarningIf(err error) {
	Log.WriteIf(err, "WARN")
}


// ErrorIf will raise an error-level message if err is not nil
func ErrorIf(err error) {
	Log.WriteIf(err, "ERROR")
}


// DebugWith will output a debug-level message
func DebugWith(message string, additionalInfo ...string) {
	Log.WriteAt("DEBUG", message, additionalInfo...)
}


// InfoWith will output an info-level message
func InfoWith(message string, additionalInfo ...string) {
	Log.WriteAt("INFO", message, additionalInfo...)
}


// WarnWith will output a warning-level message
func WarnWith(message string, additionalInfo ...string) {
	Log.WriteAt("WARN", message, additionalInfo...)
}


// ErrorWith will output an error-level message
func ErrorWith(message string, additionalInfo ...string) {
	Log.WriteAt("ERROR", message, additionalInfo...)
}


// Data will create a compatible way to represent additional log information
func Data(key string, value interface{}) string {
	return fmt.Sprintf("%s|%s", key, value)
}