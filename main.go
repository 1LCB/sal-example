package main

import (
	"examples/controllers"

	"github.com/1LCB/sal"
)

func main() {
	api := sal.NewAPI("Products API")
	
	controllers.RegisterRoutesFromController()

	api.Run(":8000")
}
