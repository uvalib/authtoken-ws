package main

import (
	"authtokenws/client"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"gopkg.in/yaml.v2"
)

type TestConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

var goodWhom = "*"
var goodWhat = "*"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

func TestHealthCheck(t *testing.T) {
	expected := http.StatusOK
	status := client.HealthCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestVersionCheck(t *testing.T) {
	expected := http.StatusOK
	status, version := client.VersionCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(version) == 0 {
		t.Fatalf("Expected non-zero length version string\n")
	}
}

func TestRuntimeCheck(t *testing.T) {
	expected := http.StatusOK
	status, runtime := client.RuntimeCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if runtime == nil {
		t.Fatalf("Expected non-nil runtime info\n")
	}

	if  len( runtime.Version ) == 0 ||
		runtime.AllocatedMemory == 0 ||
		runtime.CpuCount == 0 ||
		runtime.GoRoutineCount == 0 ||
		runtime.ObjectCount == 0 {
		t.Fatalf("Expected non-zero value in runtime info but one is zero\n")
	}
}

func TestHappyDay(t *testing.T) {
	expected := http.StatusOK
	status := client.Auth(cfg.Endpoint, goodWhom, goodWhat, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestEmptyWhom(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Auth(cfg.Endpoint, empty, goodWhat, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestEmptyWhat(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Auth(cfg.Endpoint, goodWhom, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestEmptyToken(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.Auth(cfg.Endpoint, goodWhom, goodWhat, empty)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestBadToken(t *testing.T) {
	expected := http.StatusForbidden
	err := client.Auth(cfg.Endpoint, goodWhom, goodWhat, badToken)
	if err != expected {
		t.Fatalf("Expected %v, got %v\n", expected, err)
	}
}

func loadConfig() TestConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c TestConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("endpoint [%s]\n", c.Endpoint)
	log.Printf("token    [%s]\n", c.Token)

	return c
}
