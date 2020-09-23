package config

import (
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
	SessionFactory *SessionFactory
	dbConf         *dbConf
}

func NewConfig() *Config {
	return &Config{
		dbConf: &dbConf{
			dbUser:     GetEnv("dbuser", "root"),
			dbPassword: GetEnv("dbpassword", ""),
			dbHost:     GetEnv("dbhost", "localhost"),
			dbPort:     GetEnv("dbport", "3306"),
			schema:     GetEnv("dbschema", "test"),
			dbEngine:   GetEnv("dbengine", "mysql"),
		},
	}
}

func (c *Config) InitDb() error {
	fmt.Println("======= Create DB Connection =======")
	sf, err := NewSessionFactory(c.dbConf.dbEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbConf.dbUser, c.dbConf.dbPassword, c.dbConf.dbHost, c.dbConf.dbPort, c.dbConf.schema))
	if err != nil {
		return err
	}
	c.SessionFactory = sf
	return nil
}
