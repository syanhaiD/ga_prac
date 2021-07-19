package main

import (
	"fmt"
	"github.com/syanhaiD/ga_prac/pkg/rdb"
	gomig "github.com/syanhaiD/gomig/pkg/proc"
	"os"
	"strings"
)

func main() {
	projectName := "ga_prac"
	env := "test"

	p, _ := os.Getwd()
	splitPath := strings.Split(p, projectName)
	dbTomlPath := fmt.Sprintf("%v%v/cmd/mig_for_ga/database.toml", splitPath[0], projectName)
	dbTomlPath = fmt.Sprintf("cmd/mig_for_ga/database.toml")
	rdb.DatabaseTomlPath = dbTomlPath
	rdb.Connect()
	sqlFilesPath := fmt.Sprintf("%v%v/pkg/rdb/sqlfiles/", splitPath[0], projectName)
	sqlFilesPath = fmt.Sprintf("pkg/rdb/sqlfiles/")
	rdb.SqlFilesPath = sqlFilesPath
	schemaTomlPath := fmt.Sprintf("%v%v/pkg/rdb/schema.toml", splitPath[0], projectName)
	schemaTomlPath = fmt.Sprintf("pkg/rdb/schema.toml")
	connTomlPath := fmt.Sprintf("%v%v/cmd/mig_for_ga/conn.toml", splitPath[0], projectName)
	connTomlPath = fmt.Sprintf("cmd/mig_for_ga/conn.toml")
	err := gomig.Exec(schemaTomlPath, env, connTomlPath, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rdb.Seed()

	os.Exit(0)
}
