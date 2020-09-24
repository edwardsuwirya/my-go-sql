package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type HttpConf struct {
	Host string
	Port string
}
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
	HttpConf       *HttpConf
	env            string
}

func NewConfig(env string) *Config {
	c := &Config{env: env}
	c.dbConf = &dbConf{
		dbUser:     c.GetEnv("DBUSER", "root"),
		dbPassword: c.GetEnv("DBPASSWORD", ""),
		dbHost:     c.GetEnv("DBHOST", "0.0.0.0"),
		dbPort:     c.GetEnv("DBPORT", "3306"),
		schema:     c.GetEnv("DBSCHEMA", "test"),
		dbEngine:   c.GetEnv("DBENGINE", "mysql"),
	}
	c.HttpConf = &HttpConf{
		Host: c.GetEnv("HTTPHOST", "localhost"),
		Port: c.GetEnv("HTTPPORT", "8080"),
	}
	return c
}

func (c *Config) InitDb() error {
	sf, err := NewSessionFactory(c.dbConf.dbEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbConf.dbUser, c.dbConf.dbPassword, c.dbConf.dbHost, c.dbConf.dbPort, c.dbConf.schema))
	if err != nil {
		return err
	}
	c.SessionFactory = sf
	return nil
}

func (c *Config) GetEnv(key, defaultValue string) string {
	viper.AddConfigPath(".")
	e := c.env
	if e == "" {
		viper.SetConfigFile(".env")
	} else {
		f := fmt.Sprintf("%s", e)
		viper.AddConfigPath(".")
		viper.SetConfigType("env")
		viper.SetConfigName(f)
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
