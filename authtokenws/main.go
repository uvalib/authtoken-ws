package main

import (
	"fmt"
	"github.com/uvalib/authtoken-ws/authtokenws/cache"
	"github.com/uvalib/authtoken-ws/authtokenws/config"
	"github.com/uvalib/authtoken-ws/authtokenws/handlers"
	"github.com/uvalib/authtoken-ws/authtokenws/logger"
	"log"
	"net/http"
	"time"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// load the token cache
	err := cache.LoadTokenCache()
	if err != nil {
		log.Fatal(err)
	}

	// setup router and server...
	serviceTimeout := 15 * time.Second
	router := NewRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Configuration.ServicePort),
		Handler:      router,
		ReadTimeout:  serviceTimeout,
		WriteTimeout: serviceTimeout,
	}
	log.Fatal(server.ListenAndServe())
}
