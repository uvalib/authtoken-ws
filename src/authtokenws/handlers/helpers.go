package handlers

import (
	"authtokenws/api"
	"encoding/json"
	"log"
	"net/http"
)

func encodeLookupResponse(w http.ResponseWriter, status int) {
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.DefaultResponse{Status: status, Message: http.StatusText(status)}); err != nil {
		log.Fatal(err)
	}
}

func encodeHealthCheckResponse(w http.ResponseWriter, status int, message string) {
	healthy := status == http.StatusOK
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.HealthCheckResponse{CheckType: api.HealthCheckResult{Healthy: healthy, Message: message}}); err != nil {
		log.Fatal(err)
	}
}

func encodeVersionResponse(w http.ResponseWriter, status int, version string) {
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.VersionResponse{Version: version}); err != nil {
		log.Fatal(err)
	}
}

func encodeRuntimeResponse(w http.ResponseWriter, status int, version string, cpus int, goroutines int, heapcount uint64, alloc uint64) {
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.RuntimeResponse{Version: version, CPUCount: cpus, GoRoutineCount: goroutines, ObjectCount: heapcount, AllocatedMemory: alloc}); err != nil {
		log.Fatal(err)
	}
}

func jsonAttributes(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
