package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	// load .env file
	v := viper.New()
	v.SetConfigFile(".env")
	v.AddConfigPath("./../../")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return v
}
