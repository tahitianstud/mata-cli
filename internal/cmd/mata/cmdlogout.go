package mata

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/io/file"
	"github.com/tahitianstud/mata-cli/internal"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
)

// logoutCommand will define the logout command
func logoutCommand() cli.Command {
	return cli.Command{
		Name:   "logout",
		Usage:  "Logs out of the current session",
		Action: doLogout,
	}
}

func doLogout(_ *cli.Context) error {

	ok := file.DeleteIfExistsIn(internal.ConfigLocation, ".session")

	if ok == false {
		return errors.New("could not logout")
	}

	log.InfoWith("Logout successful")

	return nil
}
