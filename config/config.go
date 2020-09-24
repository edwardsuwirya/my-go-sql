package config

import (
	"fmt"
	"github.com/spf13/viper"
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
	env            string
}

func NewConfig(env string) *Config {
	c := &Config{env: env}
	c.dbConf = &dbConf{
		dbUser:     c.GetEnv("dbuser", "root"),
		dbPassword: c.GetEnv("dbpassword", ""),
		dbHost:     c.GetEnv("dbhost", "localhost"),
		dbPort:     c.GetEnv("dbport", "3306"),
		schema:     c.GetEnv("dbschema", "test"),
		dbEngine:   c.GetEnv("dbengine", "mysql"),
	}
	return c
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
