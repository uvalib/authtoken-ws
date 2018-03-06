package dao

import (
	"database/sql"
	// needed by the linter
	_ "github.com/go-sql-driver/mysql"
)

type dbStruct struct {
	*sql.DB
}

// Permissions -- a permission associated with a token
type Permissions struct {
	Whom string
	What string
}

// TokenPermissions -- a map of tokens and their associated permissions
type TokenPermissions map[string]Permissions

//
// DB -- the database instance
//
var DB *dbStruct

//
// NewDB -- create the database singleton
//
func NewDB(dataSourceName string) error {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	DB = &dbStruct{db}
	return nil
}

//
// DestroyDB -- destroy database singleton
//
func (db *dbStruct) DestroyDB() error {
	return db.Close()
}

//
// GetAuthTokens -- get all authentication tokens
//
func (db *dbStruct) GetAuthTokens() (TokenPermissions, error) {

	rows, err := db.Query("SELECT token, whom, what from authtokens")
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
