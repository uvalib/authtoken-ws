package logger

import (
    "log"
)

func Log( msg string ) {
    log.Printf( "AUTHTOKEN: %s", msg )
}