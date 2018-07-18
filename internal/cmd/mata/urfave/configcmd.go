package urfave

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"golang.org/x/crypto/openpgp/errors"
	modelUtils "gopkg.in/jeevatkm/go-model.v1"
	"github.com/tahitianstud/mata-cli/internal/mata"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
)

// ConfigCommand defines the config command
func ConfigCommand() cli.Command {
	return cli.Command{
		Name:      "config",
		Usage:     "Configures global settings",
		Action:    DoConfig,
		UsageText: fmt.Sprintf("%s config", mata.AppName),
	}
}

// usage examples:
//   mata config                         # global configuration

// DoConfig will save the global configuration settings
func DoConfig(c *cli.Context) (err error) {

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

var (
	// ConfigFileName is the name of the config file to look for
	ConfigFileName = "mata.yml"
)

// ReadConfigFromFile will read a `Config` struct from a Yaml file
func ReadConfigFromFile(file string) (config mata.Config, err error) {
	config = mata.NewConfig()

	viper.SetConfigFile(file)
	err = viper.ReadInConfig()
	if err != nil {

		log.ErrorWith("Unable to read configuration file",
			log.Data("configFile", file))

		return config, err
	}

	// use introspection to fetch all fields declared in the model
	configFields, err := modelUtils.Fields(config)
	if err == nil {
		for _, configField := range configFields {
			configName := configField.Name
			configKey := configField.Tag.Get("key")

			modelUtils.Set(&config, configName, viper.Get(configKey))
		}
	}

	return config, nil

}

// WriteConfigToFile will write out the current config to a file
func WriteConfigToFile(config mata.Config, file string) (err error) {
	// TODO: check if file exists and create if not

	if _, err := os.Stat(file); os.IsNotExist(err) {

		log.DebugWith("File does not exist, creating it...",
			log.Data("configFile", file))

		parentDir := filepath.Dir(file)
		err = os.MkdirAll(parentDir, os.ModePerm)
		if err == nil {
			_, err = os.Create(file)
		}
	}

	if err == nil {

		viper.SetConfigFile(file)

		// use introspection to populate config file using all fields declared in the model
		configFields, err := modelUtils.Fields(config)
		if err == nil {
			for _, configField := range configFields {
				configName := configField.Name
				configKey := configField.Tag.Get("key")

				val, err := modelUtils.Get(&config, configName)
				if err != nil {
					return err
				}
				viper.Set(configKey, val)
			}
		}

		err = viper.WriteConfig()
	}

	return err
}
