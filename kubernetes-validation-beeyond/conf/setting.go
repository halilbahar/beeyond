package conf

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server               Server
	Database             Database
	KubernetesJsonschema KubernetesJsonschema
	TestDataBase         TestDataBase
	Authentication       Authentication
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

type Authentication struct {
	Url      string
	Port     string
	ClientId string
	Realm    string
	Username string
	Password string
}

var Configuration Configurations
var V *viper.Viper

func Init() {
	V = viper.New()
	V.SetConfigName("config")
	V.SetConfigType("yml")
	V.AddConfigPath("./conf")
	V.AutomaticEnv()

	if err := V.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := V.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}

func ConvertStringToRSA(key string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	return pub.(*rsa.PublicKey)
}
