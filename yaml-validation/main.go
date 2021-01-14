package main

import (
	"./routers"
)
func main() {
	r := routers.InitRouter()
	r.Run(":8180")
}
