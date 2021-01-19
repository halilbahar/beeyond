package main

import (
	"./pkg/setting"
	"./routers"
	"./services"
)

func main() {
	setting.Init()
	services.Init()
	routers.Init()
}
