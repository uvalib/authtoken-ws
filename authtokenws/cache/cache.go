package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/uvalib/authtoken-ws/authtokenws/dao"
	"github.com/uvalib/authtoken-ws/authtokenws/logger"
)

// create the cache
var theCache = cache.New(cache.NoExpiration, cache.NoExpiration)

//
// LoadTokenCache -- one time load of the token cache
//
func LoadTokenCache() error {

	// create the storage singleton
	err := dao.NewDatastore()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return err
	}

	// load the active tokens
	tokens, err := dao.Store.GetActiveTokens()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return err
	}

	// destroy our storage instance
	_ = dao.Store.Destroy()

	// add the tokens to the local cache
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
