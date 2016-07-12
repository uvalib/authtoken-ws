package main

import (
    "log"
    "net/http"
    "time"
    "authtokenws/config"
)

func Logger(inner http.Handler, name string) http.Handler {

   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

      start := time.Now( )

      inner.ServeHTTP( w, r )

      log.Printf(
         "%s: %s (%s) -> method %s, time %s",
          config.Configuration.ServiceName,
          r.Method,
          r.RequestURI,
          name,
          time.Since( start ),
      )
   } )
}
