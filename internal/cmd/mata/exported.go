package mata

// cli describes the interface to the mata package
type cliRunner interface {
	run()
}

// run will execute the application that was created
// using the chosen implementation
func Run(version string) {
	createCLI(version).run()
}