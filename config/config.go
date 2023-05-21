package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	spfviper "github.com/spf13/viper"
)

type Configuration struct {
	DatabaseURI string `mapstructure:"DATABASE_URI"`
	Port        string `mapstructure:"PORT"`
}

type Loader struct {
	cfg                     Configuration
	viper                   *spfviper.Viper
	onConfigChangeCallbacks []func(oldCfg Configuration)
}

func (c *Loader) Config() Configuration {
	return c.cfg
}

func (c *Loader) applyCallbacks() {
	c.viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		oldCfg := c.cfg
		fmt.Printf("Old config: %v\n", c.cfg)
		if err := c.Load(); err != nil {
			log.Println("error when refreshing config", err)
		}
		fmt.Printf("New config: %v\n", c.cfg)
		for _, fn := range c.onConfigChangeCallbacks {
			fn(oldCfg)
		}
	})
}

func (c *Loader) OnDatabaseURIChange(fn func(oldURI string, newURI string)) {
	c.onConfigChangeCallbacks = append(c.onConfigChangeCallbacks, func(oldCfg Configuration) {
		if oldCfg.DatabaseURI != c.cfg.DatabaseURI {
			fn(oldCfg.DatabaseURI, c.cfg.DatabaseURI)
		}
	})
	c.applyCallbacks()
}

func (c *Loader) OnPortChange(fn func(oldPort string, newPort string)) {
	c.onConfigChangeCallbacks = append(c.onConfigChangeCallbacks, func(oldCfg Configuration) {
		if oldCfg.Port != c.cfg.Port {
			fn(oldCfg.Port, c.cfg.Port)
		}
	})
	c.applyCallbacks()
}

func (c *Loader) Load() error {
	cfg := Configuration{}
	err := c.viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}
	fmt.Printf("config changed: %+v\n", cfg)
	c.cfg = cfg
	return nil
}

func NewLoader() Loader {
	log.Println("read config")

	var cfg Configuration
	viper := spfviper.New()

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

	loader := Loader{
		cfg:   cfg,
		viper: viper,
	}

	return loader
}

func NewLoaderAndLoad() (Loader, error) {
	loader := NewLoader()
	err := loader.Load()
	return loader, err
}
