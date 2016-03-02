package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
)

func Auth( endpoint string, whom string, what string, token string ) int {

    url := fmt.Sprintf( "%s/authorize/%s/%s/%s", endpoint, whom, what, token )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}