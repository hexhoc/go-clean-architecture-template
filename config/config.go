package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App        App
		HTTP       HTTP
		Log        Log
		Datasource Datasource
		Prometheus Prometheus
	}

	App struct {
		Name    string
		Version string
	}

	HTTP struct {
		Port    string
		Version string
	}

	Log struct {
		Level string
	}

	Datasource struct {
		Username string
		Password string
		Host     string
		Port     string
		Database string
		Sslmode  string
	}

	Prometheus struct {
		gateway string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName("config_dev") // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")   // optionally look for config in the working directory
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	cfg.App = App{
		viper.GetString("app.name"),
		viper.GetString("app.version"),
	}

	cfg.HTTP = HTTP{
		viper.GetString("http.port"),
		viper.GetString("http.version"),
	}

	cfg.Log = Log{
		viper.GetString("logger.log_level"),
	}

	cfg.Datasource = Datasource{
		viper.GetString("datasource.pg.username"),
		viper.GetString("datasource.pg.password"),
		viper.GetString("datasource.pg.host"),
		viper.GetString("datasource.pg.port"),
		viper.GetString("datasource.pg.database"),
		viper.GetString("datasource.pg.sslmode"),
	}

	cfg.Prometheus = Prometheus{
		viper.GetString("prometheus.gateway"),
	}

	return cfg, nil
}
