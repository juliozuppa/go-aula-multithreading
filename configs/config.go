package configs

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Application struct {
	SearchTimeout int `mapstructure:"search_timeout"`
}

type BrasilApi struct {
	Name string `mapstructure:"name"`
	URL  string `mapstructure:"url"`
}

type ViaCep struct {
	Name string `mapstructure:"name"`
	URL  string `mapstructure:"url"`
}

type CepServices struct {
	BrasilApi `mapstructure:"brasil_api"`
	ViaCep    `mapstructure:"via_cep"`
}

type Config struct {
	Application Application `mapstructure:"application"`
	CepServices `mapstructure:"cep_services"`
}

// LoadConfig loads the configuration from a specified path.
// It uses Viper to read a YAML configuration file named "application.yml".
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.Set("verbose", true)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
