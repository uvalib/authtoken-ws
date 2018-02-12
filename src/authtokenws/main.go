package main

import (
	"authtokenws/cache"
	"authtokenws/config"
	"authtokenws/handlers"
	"authtokenws/logger"
	"fmt"
	"log"
	"net/http"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// load the token cache
	err := cache.LoadTokenCache()
	if err != nil {
		log.Fatal(err)
	}

	// setup router and serve...
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Configuration.ServicePort), router))
}
