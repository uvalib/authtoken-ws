package tests

import (
   "authtokenws/client"
   "net/http"
   "testing"
)

func TestAuthTokenHappyDay(t *testing.T) {
   expected := http.StatusOK
   status := client.Auth(cfg.Endpoint, goodWhom, goodWhat, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestAuthTokenEmptyWhom(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.Auth(cfg.Endpoint, empty, goodWhat, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestAuthTokenEmptyWhat(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.Auth(cfg.Endpoint, goodWhom, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestAuthTokenEmptyToken(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.Auth(cfg.Endpoint, goodWhom, goodWhat, empty)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestAuthTokenBadToken(t *testing.T) {
   expected := http.StatusForbidden
   err := client.Auth(cfg.Endpoint, goodWhom, goodWhat, badToken)
   if err != expected {
      t.Fatalf("Expected %v, got %v\n", expected, err)
   }
}

//
// end of file
//