package main

import (
	"kubernetes-validation-beeyond/conf"
	"kubernetes-validation-beeyond/routers"
	"kubernetes-validation-beeyond/services"
)

// @title Swagger Kubernetes Validation Beeyond API
// @version 1.0
// @description This is an API for the validation of kubernetes specifications (yaml) with constraints.
func main() {
	conf.Init()
	services.Init()
	routers.Init()
}
