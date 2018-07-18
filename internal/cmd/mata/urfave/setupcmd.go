package urfave

import (
	"fmt"

	"github.com/urfave/cli"
	"golang.org/x/crypto/openpgp/errors"
	"github.com/tahitianstud/mata-cli/internal/mata"
)

// SetupCommand defines the setup command
func SetupCommand() cli.Command {
	return cli.Command{
		Name:      "setup",
		Usage:     "Sets up the application's preferences",
		UsageText: fmt.Sprintf("%s --target <settings> setup", mata.AppName),
		Action:    DoSetup,
	}
}

// usage examples:
//   mata --target dev setup                              # setup dev settings
//   mata --target farepreprod setup                      # setup farepreprod settings

// DoSetup will create the configuration file needed to make the application work
func DoSetup(c *cli.Context) (err error) {

	var target = c.GlobalString("target")

	if len(target) > 0 {
		setupTarget(target)
	} else {
		// TODO: or show usage information for this command

		return errors.InvalidArgumentError("target argument is mandatory for setup")
	}

	//var configNeverWritten, config = model.FetchConfig(isTargetSetup)
	//if configNeverWritten == true {
	//	logrus.
	//		WithField("configFile", model.ConfigFile(isTargetSetup)).
	//		Info("Create new configuration")
	//} else {
	//	logrus.
	//		WithField("configFile", model.ConfigFile(isTargetSetup)).
	//		Info("Overwrite configuration")
	//}
	//
	//var questions = []*survey.Question{
	//	{
	//		Name: "graylogServerUrl",
	//		Prompt: &survey.Input{
	//			Message: "Graylog URL",
	//			Default: config.GraylogServerUrl,
	//			Help:    "What is the URL of the Graylog Server ?",
	//		},
	//	},
	//	{
	//		Name: "graylogUsername",
	//		Prompt: &survey.Input{
	//			Message: "Graylog username",
	//			Default: config.GraylogUsername,
	//			Help:    "What username will mata use to login to Graylog server ?",
	//		},
	//	},
	//	{
	//		Name: "graylogPassword",
	//		Prompt: &survey.Password{
	//			Message: "Graylog password",
	//			Help:    "What password will mata use to login to Graylog server ?",
	//		},
	//	},
	//}
	//var answers model.Config
	//
	//err := survey.Ask(questions, &answers)
	//
	//// encode password to base64 equivalent so that it's not stored as clear
	//encodedPassword := base64.StdEncoding.EncodeToString([]byte(answers.GraylogPassword))
	//answers.GraylogPassword = encodedPassword
	//
	//// write file out config to file
	//err = model.SaveConfig(answers)
	//if err != nil {
	//	logrus.
	//		WithField("configFile", model.ConfigFile(isTargetSetup)).
	//		Error("Could not save configuration")
	//} else {
	//	logrus.
	//		WithField("configFile", model.ConfigFile(isTargetSetup)).
	//		Info("Configuration saved")
	//}

	return err
}

func setupTarget(s string) {

}

func setupGlobal() {

	//var configNeverWritten, config = model.FetchConfig(target)
	//if configNeverWritten == true {
	//	logrus.
	//		WithField("configFile", model.ConfigFile(target)).
	//		Info("Create new configuration")
	//} else {
	//	logrus.
	//		WithField("configFile", model.ConfigFile(target)).
	//		Info("Overwrite configuration")
}
