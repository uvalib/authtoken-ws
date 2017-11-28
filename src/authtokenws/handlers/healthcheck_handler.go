package handlers

import (
	"net/http"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	encodeHealthCheckResponse(w, http.StatusOK, "")
}

//
// end of file
//
