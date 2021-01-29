package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type Server struct {
	HttpPort string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

type KubernetesJsonschema struct {
	KubernetesVersion string
	Url               string
}

var KubernetesJsonschemaSetting = &KubernetesJsonschema{}

var cfg *ini.File

// Setup initialize the configuration instance
func Init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("kubernetes-jsonschema", KubernetesJsonschemaSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
