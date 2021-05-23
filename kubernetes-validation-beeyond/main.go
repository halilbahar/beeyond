package main

import (
	"kubernetes-validation-beeyond/conf"
	"kubernetes-validation-beeyond/routers"
	"kubernetes-validation-beeyond/services"
)

func main() {
	conf.Init()
	services.Init()
	routers.Init()
}
