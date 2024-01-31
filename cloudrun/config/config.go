package config

import "github.com/spf13/viper"

type Conf struct {
	LogLevel          string `mapstructure:"LOG_LEVEL"`
	WebServerPort     int    `mapstructure:"WEB_SERVER_PORT"`
	HttpClientTimeout int    `mapstructure:"HTTP_CLIENT_TIMEOUT_MS"`
	ViaCepApiBaseUrl  string `mapstructure:"VIACEP_API_BASE_URL"`
	WeatherApiBaseUrl string `mapstructure:"WEATHER_API_BASE_URL"`
	WeatherApiKey     string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*Conf, error) {
	var c *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	return c, nil
}
