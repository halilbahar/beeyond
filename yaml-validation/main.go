package main

import (
	"yaml-validation/pkg/setting"
	"yaml-validation/routers"
	"yaml-validation/services"
)

func main() {
	setting.Init()
	services.Init()
	routers.Init()
}
