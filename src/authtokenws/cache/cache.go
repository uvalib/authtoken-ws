package cache

import (
   "authtokenws/config"
   "authtokenws/logger"
   "database/sql"
   "fmt"
   // needed
   _ "github.com/go-sql-driver/mysql"
   "github.com/patrickmn/go-cache"
   "log"
)

// create the cache
var theCache = cache.New(cache.NoExpiration, cache.NoExpiration)

// whet we put in the cache
type permissions struct {
   Whom string
   What string
}

//
// LoadTokenCache -- one time load of the token cache
//
func LoadTokenCache() error {

   connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1", config.Configuration.DbUser,
      config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName)
   db, err := sql.Open("mysql", connectStr)
   if err != nil {
      log.Fatal(err)
   }
   defer db.Close()

   rows, err := db.Query("SELECT token, whom, what from authtokens")
   if err != nil {
      log.Fatal(err)
   }
   defer rows.Close()

   for rows.Next() {
      var token, whom, what string
      err := rows.Scan(&token, &whom, &what)
      if err != nil {
         log.Fatal(err)
      }
      logger.Log(fmt.Sprintf("Adding: %s/%s -> %s", whom, what, token))
      theCache.Set(token, permissions{Whom: whom, What: what}, cache.NoExpiration)
   }
   err = rows.Err()
   if err != nil {
      log.Fatal(err)
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
      permission := hit.(permissions)
      found = (permission.Whom == "*" || permission.Whom == whom) &&
         (permission.What == "*" || permission.What == what)
   }

   logger.Log(fmt.Sprintf("Token found: %v", found))
   return found
}

//
// end of file
//
