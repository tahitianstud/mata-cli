package logrus

import (
	"github.com/sirupsen/logrus"
	"strings"
	"io"
	"os"
)


// Handler is the implementation of the Logger interface
type Handler struct {
	Output io.Writer
}


func GetLogrusHandler() *Handler {
	return &Handler{
		Output: os.Stderr,
	}
}


// DieIf will use the logrus library in order to log at FATAL-level and exit application if err is not nil
func (h *Handler) DieIf(err error) {

	// if err is not nil, log fatal and exit

	logrus.SetOutput(h.Output)

	if err != nil {
		// TODO: probably need a stacktrace here and show more
		logrus.Fatal(err.Error())
	}

	// else do nothing

}



// If will use the logrus library in order to log at <level> if err is not nil
func (h *Handler) WriteIf(err error, level string) {
	
	if err != nil {
		parseLevel, e := logrus.ParseLevel(level)

		if e != nil {

			return

		}

		logrus.SetOutput(h.Output)

		errorMessage := err.Error()
		switch parseLevel {

		case logrus.DebugLevel:
			logrus.Debug(errorMessage)

		case logrus.InfoLevel:
			logrus.Info(errorMessage)

		case logrus.WarnLevel:
			logrus.Warn(errorMessage)

		case logrus.ErrorLevel:
			logrus.Error(errorMessage)

		}

	}

	// else do nothing

}



// At will log a message at a particular level
func (h *Handler) WriteAt(level string, message string, additionalInfo ...string) {

	logrus.SetOutput(h.Output)

	// create entries from additional information

	var entry = logrus.NewEntry(logrus.StandardLogger())

	for _, info := range additionalInfo {
		key, val := parseInfo(info)

		if len(key) > 0 {
			entry = entry.WithField(key, val)
		}
	}

	parseLevel, e := logrus.ParseLevel(level)

	if e == nil {

		switch parseLevel {

		case logrus.DebugLevel:
			entry.Debug(message)

		case logrus.InfoLevel:
			entry.Info(message)

		case logrus.WarnLevel:
			entry.Warn(message)

		case logrus.ErrorLevel:
			entry.Error(message)

		}

	}

}



func parseInfo(info string) (string, string) {

	if ! strings.Contains(info, "|") {

		return "", ""

	} else {

		splitStrings := strings.Split(info, "|")

		return splitStrings[0], splitStrings[1]

	}
}



// SetLevelTo will set the right Logrus level
func (h *Handler) SetLevelTo(level string) {

	parseLevel, e := logrus.ParseLevel(level)

	if e == nil {

		logrus.SetLevel(parseLevel)

	}

}