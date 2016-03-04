package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "testing"
    "authtokenws/client"
    "gopkg.in/yaml.v2"
    "net/http"
)

type TestConfig struct {
    Endpoint  string
    Token     string
}

var cfg = loadConfig( )

var goodWhom = "*"
var goodWhat = "*"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

func TestHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status := tester( goodWhom, goodWhat, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestEmptyWhom( t *testing.T ) {
    expected := http.StatusBadRequest
    status := tester( empty, goodWhat, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestEmptyWhat( t *testing.T ) {
    expected := http.StatusBadRequest
    status := tester( goodWhom, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestEmptyToken( t *testing.T ) {
    expected := http.StatusBadRequest
    status := tester( goodWhom, goodWhat, empty )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    err := tester( goodWhom, goodWhat, badToken )
    if err != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, err )
    }
}

func tester( whom string, what string, token string ) int {
    return client.Auth( cfg.Endpoint, whom, what, token )
}

func loadConfig( ) TestConfig {

    data, err := ioutil.ReadFile( "service_test.yml" )
    if err != nil {
        log.Fatal( err )
    }

    var c TestConfig
    if err := yaml.Unmarshal( data, &c ); err != nil {
        log.Fatal( err )
    }

    fmt.Printf( "endpoint [%s]\n", c.Endpoint )
    fmt.Printf( "token    [%s]\n", c.Token )

    return c
}