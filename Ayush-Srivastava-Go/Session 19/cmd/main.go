package main

import (
	"log"
	"net/http"
	"os"
	"userdb/internal/routes"
	"userdb/internal/utils"

	"github.com/joho/godotenv"
)

func main() {

	routes.SetupRoutes()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(utils.ENVFileLoadError)
		return
	}



	log.Println("Server runnning on PORT", os.Getenv("PORT"))
	err = http.ListenAndServe(os.Getenv("ENDPOINT_ADDRESS"), nil)
	if err != nil {
		log.Println("Error in starting the server", err)
		return
	}
	
}