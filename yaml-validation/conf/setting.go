package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server               Server
	Database             Database
	KubernetesJsonschema KubernetesJsonschema
}

type Server struct {
	HttpPort string
}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type KubernetesJsonschema struct {
	KubernetesVersion string
	Url               string
}

var Configuration Configurations

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
