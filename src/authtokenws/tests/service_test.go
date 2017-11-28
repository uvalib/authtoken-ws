package tests

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type testConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

var goodWhom = "*"
var goodWhat = "*"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("endpoint [%s]\n", c.Endpoint)
	log.Printf("token    [%s]\n", c.Token)

	return c
}

//
// end of file
//
