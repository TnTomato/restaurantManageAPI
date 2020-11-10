package main

import (
	"fmt"
	"log"
	"os"

	_ "restaurantManageAPI/init/runtime"
	"restaurantManageAPI/pkg/router"
)

func main() {
	err := router.Router.Run(fmt.Sprintf(":%s", os.Getenv("GIN_SERVER_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
