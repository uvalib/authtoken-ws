package config

import (
    "flag"
    "log"
)

type Config struct {
    ServicePort   string
    DbHost        string
    DbName        string
    DbUser        string
    DbPassphrase  string
}

var Configuration = LoadConfig( )

func LoadConfig( ) Config {

    c := Config{}

    // process command line flags and setup configuration
    flag.StringVar( &c.ServicePort, "port", "8080", "The service listen port" )
    flag.StringVar( &c.DbHost, "dbhost", "mysqldev.lib.virginia.edu:3306", "The database server hostname:port" )
    flag.StringVar( &c.DbName, "dbname", "authtoken_development", "The database name" )
    flag.StringVar( &c.DbUser, "dbuser", "authtoken", "The database username" )
    flag.StringVar( &c.DbPassphrase, "dbpassword", "dbpassword", "The database passphrase")

    flag.Parse()

    log.Printf( "ServicePort:  %s", c.ServicePort )
    log.Printf( "DbHost:       %s", c.DbHost )
    log.Printf( "DbName:       %s", c.DbName )
    log.Printf( "DbUser:       %s", c.DbUser )
    log.Printf( "DbPassphrase: %s", c.DbPassphrase )

    return c
}

