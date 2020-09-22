package config

import (
	"fmt"
	"os"
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
			dbUser:     os.Getenv("dbuser"),
			dbPassword: os.Getenv("dbpassword"),
			dbHost:     os.Getenv("dbhost"),
			dbPort:     os.Getenv("dbport"),
			schema:     os.Getenv("dbschema"),
			dbEngine:   os.Getenv("dbengine"),
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
