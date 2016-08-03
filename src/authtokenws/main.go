package main

import (
    "fmt"
    "log"
    "net/http"
    "authtokenws/config"
    "authtokenws/logger"
)

func main( ) {

    logger.Log( fmt.Sprintf( "===> version: '%s' <===", Version( ) ) )

    // load the token cache
    err := LoadTokenCache( )
    if err != nil {
        log.Fatal( err )
    }

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.ServicePort ), router ) )
}

