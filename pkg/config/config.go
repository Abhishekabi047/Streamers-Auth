package config

import "github.com/spf13/viper"

type Config struct {
	DB_port     string `mapstructure:"DB_PORT"`
	DB_host     string `mapstructure:"DB_HOST"`
	DB_username string `mapstructure:"DB_USER"`
	DB_password string `mapstructure:"DB_PASSWORD"`
	DB_name     string `mapstructure:"DB_NAME"`
	Port        string `mapstructure:"PORT"`
	Key1        string `mapstructure:"KEY1"`
	Key2        string `mapstructure:"KEY2"`
	Key3        string `mapstructure:"KEY3"`
}

var envs = []string{"DB_PORT", "DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "PORT","KEY1","KEY2","KEY3"}

var config *Config

func LoadConfig() (config *Config,err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return nil, err
		}
	}
	err=viper.Unmarshal(&config)
	return config, nil
}

func GetConfig() *Config {
	return config
}
