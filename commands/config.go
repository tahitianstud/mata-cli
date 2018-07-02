package commands

import (
	"github.com/urfave/cli"
	"golang.org/x/crypto/openpgp/errors"
)

// usage examples:
//   mata config                         # global configuration

/**
 * Config will allow the save global configuration settings
 */
func Config(c *cli.Context) (err error) {

	// defensive conditions
	if c.NArg() > 0 {
		return errors.InvalidArgumentError("config does not take any argument")
	}

	// TODO: do global app configuration

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