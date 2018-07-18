package mata

import (
	"fmt"
	"os"
)

var (
	// AppName is the name of the cli application
	AppName = "mata"

	// ConfigLocation is where the config files (global + targets) will be stored
	ConfigLocation = fmt.Sprintf("%s/.mata", os.Getenv("HOME"))
)
