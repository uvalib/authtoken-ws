package handlers

import (
    "log"
    "encoding/json"
    "net/http"
    "authtokenws/api"
)

func encodeLookupResponse( w http.ResponseWriter, status int ) {
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

func encodeVersionResponse( w http.ResponseWriter, status int, version string ) {
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.VersionResponse { Version: version } ); err != nil {
        log.Fatal( err )
    }
}

func jsonResponse( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}