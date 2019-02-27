package core

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")          // for main
	viper.AddConfigPath("../config")       // for main tests
	viper.AddConfigPath("../../config")    // for sub package testing purpose
	viper.AddConfigPath("../../../config") // for sub package testing purpose
	err := viper.ReadInConfig()            // Find and read the config file
	if err != nil {                        // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// ConfString Get config as string
func ConfString(name string) string {
	return viper.GetString(name)
}

// ConfInt Get config as string
func ConfInt(name string) int {
	return viper.GetInt(name)
}

// ConfBool Get config as bool
func ConfBool(name string) bool {
	return viper.GetBool(name)
}
