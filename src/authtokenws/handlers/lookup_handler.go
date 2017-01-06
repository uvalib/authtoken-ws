package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "authtokenws/cache"
)

func TokenLookup( w http.ResponseWriter, r *http.Request ) {
    vars := mux.Vars( r )
    whom := vars[ "whom" ]
    what := vars[ "what" ]
    token := vars[ "token" ]

    // parameters OK ?
    if cache.ParametersOk( whom, what, token ) == false {
        encodeLookupResponse(w, http.StatusBadRequest )
        return
    }

    // is this a good token ?
    if cache.ActivityIsOk(whom, what, token) == false {
        encodeLookupResponse(w, http.StatusForbidden)
        return
    }

    encodeLookupResponse(w, http.StatusOK)
}