package infrastructure

import (
	"github.com/spf13/viper"
	"github.com/vnnyx/article-service/exception"
)

type Config struct {
	MongoPoolMin           int    `mapstructure:"MONGO_POOL_MIN"`
	MongoPoolMax           int    `mapstructure:"MONGO_POOL_MAX"`
	MongoMaxIdleTimeSecond int    `mapstructure:"MONGO_MAX_IDLE_TIME_SECOND"`
	MongoURI               string `mapstructure:"MONGO_URI"`
	MongoDatabae           string `mapstructure:"MONGO_DATABASE"`
	AppPort                string `mapstructure:"APP_PORT"`
	GRPCHost               string `mapstructure:"GRPC_HOST"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)
	err = viper.Unmarshal(&config)
	exception.PanicIfNeeded(err)
	return config
}
