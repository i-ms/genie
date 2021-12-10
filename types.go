package genie

import "database/sql"

// rootPath : working directory for the project
// folderNames : list of folders to be created
type initPaths struct {
	rootPath    string
	folderNames []string
}

// name: name of cookie
// lifetime: lifetime of cookie
// persistent: if true, cookie will be persistent between browser closes
// secure : if true , cookie will be encrypted
// domain: domain cookie is associated with
type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}

type databaseConfig struct {
	dsn      string
	database string
}

// Database - type of database (redis/mysql/postgres)
type Database struct {
	DataType string
	Pool     *sql.DB
}

// redisConfig holds config values for redis
type redisConfig struct {
	host     string
	password string
	prefix   string
}
