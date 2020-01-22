package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// MySQL holds configs to connet MySQL
type MySQL struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	SchemaPath string `mapstructure:"schema_path"`
}

// AppConfig is the overall config struct
type AppConfig struct {
	MySQL `mapstructure:"mysql"`
}

// LoadConfig will reads JSON config file with explicit path, from ENV, or use default main.json
func LoadConfig(confPath string) (conf AppConfig, err error) {

	viper.SetConfigType("json")
	viper.AutomaticEnv()        // read in environment variables that match
	viper.SetEnvPrefix("AMIGO") // AMIGO_
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		viper.AddConfigPath("./config/")
		viper.SetConfigName("main")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
			return conf, err
		}
		log.Println("Using config file:", viper.ConfigFileUsed())

		if err := viper.Unmarshal(&conf); err != nil {
			log.Fatalf("Error unmarshal config file, %s", err)
			return conf, err
		}
	}
	return conf, nil
}
