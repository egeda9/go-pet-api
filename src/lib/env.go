package lib

import (
	"fmt"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	Environment     string `mapstructure:"ENV"`
	StorageFilePath string `mapstructure:"STORAGE_FILE_PATH"`
}

// NewEnv creates a new environment
// Constructs type Env, depends on Logger
func NewEnv() Env {
	localPath := "."

	viper.AddConfigPath(localPath)
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	env := Env{}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Sprintf("%s%s", "☠️ cannot read configuration ", err)
		} else {
			fmt.Sprintf("%s%s", "☠️config file was found but another error was produced", err)
			// Config file was found but another error was produced
		}
	}

	err := viper.Unmarshal(&env)
	if err != nil {
		fmt.Sprintf("%s%s", "☠️ environment can't be loaded: ", err)
	}

	ForceMapping(&env)

	fmt.Sprintf("%v", env)
	return env
}

func ForceMapping(env *Env) {
	if env.Environment == "" {
		env.Environment = viper.GetString("ENV")
	}

	if env.StorageFilePath == "" {
		env.StorageFilePath = viper.GetString("STORAGE_FILE_PATH")
	}
}
