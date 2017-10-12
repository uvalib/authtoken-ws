package handlers

import (
   "authtokenws/cache"
   "github.com/gorilla/mux"
   "net/http"
   "strings"
)

//
// TokenLookup -- handler for the token lookup method
//
func TokenLookup(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   whom := vars["whom"]
   what := vars["what"]
   token := vars["token"]

   // parameters OK ?
   if parametersOk(whom, what, token) == false {
      encodeLookupResponse(w, http.StatusBadRequest)
      return
   }

   // is this a good token ?
   if cache.ActivityIsOk(whom, what, token) == false {
      encodeLookupResponse(w, http.StatusForbidden)
      return
   }

   encodeLookupResponse(w, http.StatusOK)
}


func parametersOk(whom string, what string, token string) bool {

   // validate inbound parameters
   return len(strings.TrimSpace(whom)) != 0 &&
      len(strings.TrimSpace(what)) != 0 &&
      len(strings.TrimSpace(token)) != 0
}

//
// end of file
//