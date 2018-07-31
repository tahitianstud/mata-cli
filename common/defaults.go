package common

import (
	"fmt"
	"os"
)

var (
	// where the config files (global + targets) will be stored
	ConfigLocation = fmt.Sprintf("%s/.mata", os.Getenv("HOME"))
)