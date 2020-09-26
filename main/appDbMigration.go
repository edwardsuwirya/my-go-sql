package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myfirstgosql/config"
	"os"
)

type migration struct {
	Up   []migrationStatement `json:"up"`
	Down []migrationStatement `json:"down"`
}
type migrationStatement struct {
	Description string `json:"description"`
	Statement   string `json:"statement"`
}

func dbMigration(sf *config.SessionFactory, jsonFile string, mode string) {
	fmt.Println("Reading file ", jsonFile)
	j, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}
	defer j.Close()

	jval, err := ioutil.ReadAll(j)
	var mig migration
	err = json.Unmarshal(jval, &mig)
	if err != nil {
		panic(err)
	}
	session := sf.GetSession()
	var migStmt = make([]migrationStatement, 0)
	if mode == "up" {
		migStmt = mig.Up
	} else {
		migStmt = mig.Down
	}
	for _, ms := range migStmt {
		fmt.Printf("%-30s %30s\n", ms.Description, "[OK]")
		_, err := session.Exec(ms.Statement)
		if err != nil {
			fmt.Printf("%-30s %30s\n", ms.Description, "[FAILED]")
			panic(err)
		}
	}
}
