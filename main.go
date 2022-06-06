package main

import (
	"log"

	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/database"
	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/router"
)

func main() {
	database.InitDB()
	routes := router.OrderRouters()
	log.Fatal(routes.Run(":3000"))
}
