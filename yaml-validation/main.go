package main

import (
	"yaml-validation/conf"
	"yaml-validation/routers"
	"yaml-validation/services"
)

func main() {
	conf.Init()
	services.Init()
	routers.Init()
}
