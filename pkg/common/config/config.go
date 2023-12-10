package config

import (
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/utils"
	"github.com/spf13/viper"
)

type Config struct {
    Port 			string	`mapstructure:"PORT"`
    DBUrl			string 	`mapstructure:"DB_URL"`
	AwsRegion		string 	`mapstructure:"AWS_REGION"`
	AwsAccessKey	string 	`mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey 	string 	`mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("../pkg/common/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	utils.CheckErr(err)

	err = viper.Unmarshal(&c)

	return
}