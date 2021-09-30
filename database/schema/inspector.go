package schema

import "github.com/topcoder-club/club/crud/database/dialects"

type Connection struct {
	dialect dialects.Dialect
	host string
	user string
	password string
	database string
	charset string
}

type Database struct {
	databaseName string
	Tables []Table
}
