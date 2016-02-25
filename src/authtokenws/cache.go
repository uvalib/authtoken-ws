package main

import (
    "log"
    "fmt"
    "github.com/patrickmn/go-cache"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// create the cache
var c = cache.New( cache.NoExpiration, cache.NoExpiration )

// whet we put in the cache
type Permissions struct {
    Whom    string
    What    string
}

func LoadTokenCache( ) error {

    connectStr := fmt.Sprintf( "%s:%s@tcp(%s)/%s", config.DbUser, config.DbPassphrase, config.DbHost, config.DbName )
    db, err := sql.Open( "mysql", connectStr )
    if err != nil {
        log.Fatal( err )
    }
    defer db.Close( )

    rows, err := db.Query( "SELECT token, whom, what from authtokens" )
    if err != nil {
        log.Fatal( err )
    }
    defer rows.Close( )

    for rows.Next( ) {
        var token, whom, what string
        err := rows.Scan( &token, &whom, &what )
        if err != nil {
            log.Fatal( err )
        }
        log.Println( token, whom, what )
        log.Printf( "Adding: %s/%s -> %s", whom, what, token )
        c.Set( token, Permissions{ Whom: whom, What: what }, cache.NoExpiration )
    }
    err = rows.Err( )
    if err != nil {
        log.Fatal( err )
    }

    return nil
}

func ActivityIsOk( whom string, what string, token string ) bool {

    // validate inbound parameters
    if len( whom ) == 0 || len( what ) == 0 || len( token ) == 0 {
        return false
    }

    log.Printf( "Token lookup: whom [%s], what [%s], token [%s]", whom, what, token )

    // lookup the token in the cache
    hit, found := c.Get( token )
    if found {
        // determine if we have the correct permissions
        permission := hit.(Permissions)
        found = ( permission.Whom == "*" || permission.Whom == whom ) &&
                ( permission.What == "*" || permission.What == what )
    }

    log.Printf( "Token found: %v", found )
    return found
}