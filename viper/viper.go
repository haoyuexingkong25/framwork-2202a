package viper

import (
	"github.com/spf13/viper"
)

const CONFIG_TYPE = "yaml"

func GetViper(path, name string) error {
	viper.SetConfigName(name)        // name of config file (without extension)
	viper.SetConfigType(CONFIG_TYPE) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)        // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		return err
	}
	return nil
}
