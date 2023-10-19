package settings

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Version   string `mapstructure:"version"`
	Contracts *ContractsConfig `mapstructure:"contracts"`
	MySql *MySqlConfig `mapstructure:"mysql"`
}

type ContractsConfig struct {
	BSCT struct {
		CollectPaidMw string `mapstructure:"CollectPaidMw"`
		ProfileNFT    string `mapstructure:"ProfileNFT"`
	}
	BSC struct {
		CollectPaidMw string `mapstructure:"CollectPaidMw"`
		ProfileNFT    string `mapstructure:"ProfileNFT"`
	}
}

type MySqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
} 

var Config = new(AppConfig)

func Init() (err error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = viper.Unmarshal(Config); err != nil {
		log.Fatalln("viper.Unmarshal failed, err:", err)
	}

	return
}
