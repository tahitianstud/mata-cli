package mata

import (
	"github.com/spf13/viper"
)

var (
//configFileName = "mata.yml"

// ConfigFile is the location of the configuration file
//ConfigFile = fmt.Sprintf("%s/%s", configLocation, configFileName)
)

// package initializer
func init() {
	viper.AddConfigPath(ConfigLocation)
}

// Target defines the data structure to use to describe a server target
type Target struct {
	GraylogServerURL string
	GraylogUsername  string
	GraylogPassword  string
}

// Default constructor for `Config` data structure
func newTarget() Target {
	return Target{
		GraylogServerURL: "http://graylog.server.com:12229",
		GraylogUsername:  "user",
		GraylogPassword:  "password",
	}
}

/**
 * FetchConfig fetches a `Config` struct from a Yaml file
 */
//func FetchConfig(targetName string) (isNewConfig bool, target Target) {
////	config = newConfig()
////	isNewConfig = true
////
////	viper.
////	err := viper.ReadInConfig()
////	if err != nil {
////		logrus.
////			WithField("configFile", ConfigFile).
////			Info("No configuration found")
////	} else {
////		logrus.
////			WithField("configFile", ConfigFile).
////			Info("Using configuration from file")
////
////		// read configuration file into `Config` struct
////		config.GraylogServerUrl = viper.GetString("graylog.server.url")
////		config.GraylogUsername  = viper.GetString("graylog.username")
////		config.GraylogPassword  = viper.GetString("graylog.password")
////
////		isNewConfig = false
////	}
////
////	return isNewConfig, config
//
//}

// SaveConfig will write out the `Config` struct to the config file
func SaveConfig(config Target) (err error) {
	//viper.Set("graylog.server.url", config.GraylogServerUrl)
	//viper.Set("graylog.username", config.GraylogUsername)
	//viper.Set("graylog.password", config.GraylogPassword)
	//
	//if _, err := os.Stat(ConfigLocation); os.IsNotExist(err) {
	//	logrus.
	//		WithField("configFile", ConfigFile).
	//		Debug("File does not exist, creating it...")
	//
	//	err = os.MkdirAll(ConfigLocation, os.ModePerm)
	//}
	//
	//err = viper.WriteConfig()
	//
	//if err != nil {
	//	logrus.
	//		WithField("configFile", ConfigFile).
	//		Error("Error writing configuraton file")
	//} else {
	//	logrus.
	//		WithField("configFile", ConfigFile).
	//		Debug("Configuration file written")
	//}

	return err
}

/**
 * ConfigFile will print out the file corresponding to a certain configuration
 */
//func ConfigFile(targetName string) string {
//	//return fmt.Sprintf("%s/%s", ConfigLocation, config)
//}
