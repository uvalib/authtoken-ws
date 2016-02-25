package main

import (
    "log"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "authtokenws/api"
)

func TokenLookup( w http.ResponseWriter, r *http.Request ) {
    vars := mux.Vars( r )
    whom := vars[ "whom" ]
    what := vars[ "what" ]
    token := vars[ "token" ]

    // is this a good token ?
    if ActivityIsOk( whom, what, token ) {
        encodeResponse( w, http.StatusOK )
    } else {
        encodeResponse( w, http.StatusForbidden )
    }
}

func HealthCheck( w http.ResponseWriter, r *http.Request ) {
    encodeHealthCheckResponse( w, http.StatusOK, "" )
}

func encodeResponse( w http.ResponseWriter, status int ) {
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.DefaultResponse{ Status: status, Message: http.StatusText( status ) } ); err != nil {
        log.Fatal( err )
    }
}

func encodeHealthCheckResponse( w http.ResponseWriter, status int, message string ) {
    healthy := status == http.StatusOK
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func jsonResponse( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}