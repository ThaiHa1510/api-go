package main

import "api-go/database"

func main() {
	//fmt.Println("Hello World!")

	database.InitialMigration()
	initializeRouter()
}
