package config

import (
	"database/sql"
	"fmt"
)

type dbConf struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	schema     string
	dbEngine   string
}
type Config struct {
	Db     *sql.DB
	dbConf *dbConf
}

func NewConfig() *Config {
	return &Config{
		dbConf: &dbConf{
			dbUser:     "root",
			dbPassword: "P@ssw0rd",
			dbHost:     "localhost",
			dbPort:     "3306",
			schema:     "enigma",
			dbEngine:   "mysql",
		},
	}
}

func (c *Config) InitDb() error {
	fmt.Println("======= Create DB Connection =======")
	db, err := sql.Open(c.dbConf.dbEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbConf.dbUser, c.dbConf.dbPassword, c.dbConf.dbHost, c.dbConf.dbPort, c.dbConf.schema))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	c.Db = db
	return nil
}
