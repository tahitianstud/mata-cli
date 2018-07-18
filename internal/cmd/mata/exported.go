package mata

import "github.com/tahitianstud/mata-cli/internal/cmd/mata/urfave"

var (
	cli = urfave.CreateCLI()
)

// Run will execute the application that was created
// using the chosen implementation
func Run() {
	cli.Run()
}