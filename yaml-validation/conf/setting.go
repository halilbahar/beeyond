package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server               Server
	Database             Database
	KubernetesJsonschema KubernetesJsonschema
	TestDataBase		 TestDataBase
}

type Server struct {
	HttpPort string
}

type TestDataBase struct {
	Port string
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
var V *viper.Viper
var EnvVarPrefix = "BEEYOND_KUBERNETES_VALIDATION"
var EnvVarBindVar = "DB_PORT"
var EnvVar = EnvVarPrefix+"_"+EnvVarBindVar
var defaultDbPort = "27017"

func Init() {
	V = viper.New()
	V.SetConfigName("config")
	V.SetConfigType("yml")
	V.AddConfigPath("./conf")
	V.AutomaticEnv()

	if err := V.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//if _, ok := os.LookupEnv(EnvVar); ok{
	//	v := os.Getenv(EnvVar)
	//	fmt.Println("-------------------------------------ENV VAR VALUE: "+v)
	//	V.Set("Database.Port", os.Getenv(EnvVar))
	//}

	err := V.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
