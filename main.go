package main

import (
	"MyGram/config"
	"MyGram/router"
)

func main() {
	configuration := config.New()
	config.StartDB(configuration)

	var PORT = ":" + configuration.Get("APP_PORT")
	router.StartApp().Run(PORT)
}
