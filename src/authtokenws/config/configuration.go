package config

import (
	"authtokenws/logger"
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	ServicePort  string
	DbHost       string
	DbName       string
	DbUser       string
	DbPassphrase string
	Debug        bool
}

var Configuration = LoadConfig()

func LoadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.DbHost, "dbhost", "mysqldev.lib.virginia.edu:3306", "The database server hostname:port")
	flag.StringVar(&c.DbName, "dbname", "authtoken_development", "The database name")
	flag.StringVar(&c.DbUser, "dbuser", "authtoken", "The database username")
	flag.StringVar(&c.DbPassphrase, "dbpassword", "dbpassword", "The database passphrase")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("ServicePort:  %s", c.ServicePort))
	logger.Log(fmt.Sprintf("DbHost:       %s", c.DbHost))
	logger.Log(fmt.Sprintf("DbName:       %s", c.DbName))
	logger.Log(fmt.Sprintf("DbUser:       %s", c.DbUser))
	logger.Log(fmt.Sprintf("DbPassphrase: %s", strings.Repeat("*", len(c.DbPassphrase))))
	logger.Log(fmt.Sprintf("Debug         %t", c.Debug))

	return c
}
