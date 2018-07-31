package accessors

import (
	"github.com/sirupsen/logrus"
	"github.com/tahitianstud/mata-cli/model"
	"github.com/spf13/viper"
	modelUtils "gopkg.in/jeevatkm/go-model.v1"
)

var (
	ConfigFileName = "mata.yml"
)

/**
 * FetchConfigConfig fetches a `Config` struct from a Yaml file
 */
func ReadConfigFromFile(file string) (config model.Config, err error) {
	config = model.NewConfig()

	viper.SetConfigFile(file)
	err = viper.ReadInConfig()
	if err != nil {
		logrus.
			WithField("configFile", file).
			Error("Unable to read configuration file")

		return config, err
	} else {
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

}

/**
 * WriteConfigToFile will write out the current config to a file
 */
func WriteConfigToFile(config model.Config, file string) (err error) {
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

	return err
}


