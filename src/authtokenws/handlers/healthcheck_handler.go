package handlers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	encodeHealthCheckResponse(w, http.StatusOK, "")
}
