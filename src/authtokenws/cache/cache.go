package cache

import (
    "log"
    "fmt"
    "github.com/patrickmn/go-cache"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "strings"
    "authtokenws/config"
    "authtokenws/logger"
)

// create the cache
var c = cache.New( cache.NoExpiration, cache.NoExpiration )

// whet we put in the cache
type Permissions struct {
    Whom    string
    What    string
}

func LoadTokenCache( ) error {

    connectStr := fmt.Sprintf( "%s:%s@tcp(%s)/%s?allowOldPasswords=1", config.Configuration.DbUser,
        config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName )
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
        logger.Log( fmt.Sprintf( "Adding: %s/%s -> %s", whom, what, token ) )
        c.Set( token, Permissions{ Whom: whom, What: what }, cache.NoExpiration )
    }
    err = rows.Err( )
    if err != nil {
        log.Fatal( err )
    }

    return nil
}

func ParametersOk( whom string, what string, token string ) bool {

    // validate inbound parameters
    return len( strings.TrimSpace( whom ) ) != 0 &&
           len( strings.TrimSpace( what ) ) != 0 &&
           len( strings.TrimSpace( token ) ) != 0
}

func ActivityIsOk( whom string, what string, token string ) bool {

    logger.Log( fmt.Sprintf( "Token lookup: whom [%s], what [%s], token [%s]", whom, what, token ) )

    // lookup the token in the cache
    hit, found := c.Get( token )
    if found {
        // determine if we have the correct permissions
        permission := hit.(Permissions)
        found = ( permission.Whom == "*" || permission.Whom == whom ) &&
                ( permission.What == "*" || permission.What == what )
    }

    logger.Log( fmt.Sprintf( "Token found: %v", found ) )
    return found
}