package cache

import (
	"authtokenws/config"
	"authtokenws/dao"
	"authtokenws/logger"
	"fmt"
	// needed
	_ "github.com/go-sql-driver/mysql"
	"github.com/patrickmn/go-cache"
	"log"
)

// create the cache
var theCache = cache.New(cache.NoExpiration, cache.NoExpiration)

//
// LoadTokenCache -- one time load of the token cache
//
func LoadTokenCache() error {

	// access the database
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1", config.Configuration.DbUser,
		config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName)

	err := dao.NewDB(connectStr)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		log.Fatal(err)
	}

	tokens, err := dao.DB.GetAuthTokens()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		log.Fatal(err)
	}

	// close our connection
	dao.DB.DestroyDB( )

	for token := range tokens {
		p := tokens[token]
		logger.Log(fmt.Sprintf("Adding: %s/%s -> %s", p.Whom, p.What, token))
		theCache.Set(token, p, cache.NoExpiration)
	}

	return nil
}

//
// ActivityIsOk - looks up token in token cache
//
func ActivityIsOk(whom string, what string, token string) bool {

	logger.Log(fmt.Sprintf("Token lookup: whom [%s], what [%s], token [%s]", whom, what, token))

	// lookup the token in the cache
	hit, found := theCache.Get(token)
	if found {
		// determine if we have the correct permissions
		permission := hit.(dao.Permissions)
		found = (permission.Whom == "*" || permission.Whom == whom) &&
			(permission.What == "*" || permission.What == what)
	}

	logger.Log(fmt.Sprintf("Token found: %v", found))
	return found
}

//
// end of file
//
