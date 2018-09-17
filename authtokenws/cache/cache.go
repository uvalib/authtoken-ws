package cache

import (
	"fmt"
	"github.com/uvalib/authtoken-ws/authtokenws/config"
	"github.com/uvalib/authtoken-ws/authtokenws/dao"
	"github.com/uvalib/authtoken-ws/authtokenws/logger"
	// needed
	_ "github.com/go-sql-driver/mysql"
	"github.com/patrickmn/go-cache"
)

// create the cache
var theCache = cache.New(cache.NoExpiration, cache.NoExpiration)

//
// LoadTokenCache -- one time load of the token cache
//
func LoadTokenCache() error {

	// access the database
	err := dao.NewDB(
		config.Configuration.DbHost,
		config.Configuration.DbSecure,
		config.Configuration.DbName,
		config.Configuration.DbUser,
		config.Configuration.DbPassphrase,
		config.Configuration.DbTimeout)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return err
	}

	tokens, err := dao.DB.GetAuthTokens()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return err
	}

	// close our connection
	dao.DB.DestroyDB()

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
