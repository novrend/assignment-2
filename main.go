package main

import (
	"assignment-2/database"
	"assignment-2/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8081"
	routers.StartServer().Run(PORT)
}
