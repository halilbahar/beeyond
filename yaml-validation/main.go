package main

import (
	"./routers"
	"./services"
)

func main() {
	services.Init()
	routers.Init()
}
