package dao

import (
	"database/sql"
	"github.com/uvalib/authtoken-ws/authtokenws/config"

	// needed by the linter
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// this is our DB implementation
type storage struct {
	*sql.DB
}

//
// New -- create a DB version of the storage singleton
//
func newDBStore() (Storage, error) {

	// access the database
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1&tls=%s&timeout=%s&readTimeout=%s&writeTimeout=%s",
		config.Configuration.DbUser,
		config.Configuration.DbPassphrase,
		config.Configuration.DbHost,
		config.Configuration.DbName,
		config.Configuration.DbSecure,
		config.Configuration.DbTimeout,
		config.Configuration.DbTimeout,
		config.Configuration.DbTimeout)

	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &storage{db}, nil
}

//
// Destroy -- destroy the storage instance
//
func (s *storage) Destroy() error {
	return s.Close()
}

//
// GetActiveTokens -- get all active authentication tokens
//
func (s *storage) GetActiveTokens() (TokenPermissions, error) {

	rows, err := s.Query("SELECT token, whom, what from authtokens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authTokens := make(TokenPermissions)

	for rows.Next() {
		var token, whom, what string
		err := rows.Scan(&token, &whom, &what)
		if err != nil {
			return nil, err
		}

		authTokens[token] = Permissions{Whom: whom, What: what}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return authTokens, nil
}

//
// end of file
//
