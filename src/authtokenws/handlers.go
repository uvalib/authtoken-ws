package main

import (
    "log"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func TokenLookup( w http.ResponseWriter, r *http.Request ) {
    vars := mux.Vars( r )
    whom := vars[ "whom" ]
    what := vars[ "what" ]
    token := vars[ "token" ]

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )

    // is this a good token and
    if ActivityIsOk( whom, what, token ) {
        status := http.StatusOK
        w.WriteHeader( status )
        if err := json.NewEncoder( w ).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
            log.Fatal( err )
        }
        return
    }

    // If this token is not OK then 403
    status := http.StatusForbidden
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( Response{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        log.Fatal( err )
    }
}

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    healthy := true
    message := ""

    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( http.StatusOK )

    if err := json.NewEncoder(w).Encode( HealthCheckResponse { CheckType: HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}