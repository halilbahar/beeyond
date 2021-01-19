package main

import (
	"yaml-validation/routers"
	"yaml-validation/services"
)

func main() {
	services.Init()
	routers.Init()
}
