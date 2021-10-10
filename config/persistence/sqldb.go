package persistence

import "os"

type SqlDBConfig struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func NewSqlDBConfig() *SqlDBConfig {
	return &SqlDBConfig{
		host:     os.Getenv("mysql_host"),
		port:     os.Getenv("mysql_port"),
		user:     os.Getenv("mysql_user"),
		password: os.Getenv("mysql_password"),
		database: os.Getenv("mysql_database"),
	}
}

func (config *SqlDBConfig) Addr() string {
	return config.host + ":" + config.port
}

func (config *SqlDBConfig) User() string {
	return config.user
}

func (config *SqlDBConfig) Password() string {
	return config.password
}

func (config *SqlDBConfig) Database() string {
	return config.database
}
