package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	DatabaseURI string `mapstructure:"DATABASE_URI"`
	Port        string `mapstructure:"PORT"`
}

type ConfigLoader struct {
	cfg Configuration
}

func (c *ConfigLoader) Config() Configuration {
	return c.cfg
}

func (c *ConfigLoader) Refresh() error {
	cfg := Configuration{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}
	fmt.Printf("config changed: %+v\n", cfg)
	c.cfg = cfg
	return nil
}

func init() {
	if os.Getenv("ENV") == "test" {
		viper.SetConfigFile(".test.env")
	} else {
		viper.SetConfigFile(".env")
	}

	log.Println("init config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		viper.AutomaticEnv()
	}
}

func GetLoader() ConfigLoader {
	log.Println("read config")

	var cfg Configuration

	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return ConfigLoader{cfg: cfg}
}
