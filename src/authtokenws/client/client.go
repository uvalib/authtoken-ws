package client

import (
	"authtokenws/api"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var debugHTTP = false
var serviceTimeout = 5

//
// HealthCheck -- calls the service health check method
//
func HealthCheck(endpoint string) int {

	url := fmt.Sprintf("%s/healthcheck", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer resp.Body.Close()

	return resp.StatusCode
}

//
// VersionCheck -- calls the service version check method
//
func VersionCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/version", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer resp.Body.Close()

	r := api.VersionResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, ""
	}

	return resp.StatusCode, r.Version
}

//
// RuntimeCheck -- calls the service runtime method
//
func RuntimeCheck(endpoint string) (int, *api.RuntimeResponse) {

	url := fmt.Sprintf("%s/runtime", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.RuntimeResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, &r
}

//
// Auth -- authenticates the supplied token
//
func Auth(endpoint string, whom string, what string, token string) int {

	url := fmt.Sprintf("%s/authorize/%s/%s/%s", endpoint, whom, what, token)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer resp.Body.Close()

	return resp.StatusCode
}

//
// end of file
//
