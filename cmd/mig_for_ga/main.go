package main

import (
	"fmt"
	"github.com/syanhaiD/ga_prac/pkg/rdb"
	gomig "github.com/syanhaiD/gomig/pkg/proc"
	"os"
)

func main() {
	env := "test"

	// github_actionsの場合はPJルートからの相対パスで指定
	dbTomlPath := fmt.Sprintf("cmd/mig_for_ga/database.toml")
	rdb.DatabaseTomlPath = dbTomlPath
	rdb.Connect()
	sqlFilesPath := fmt.Sprintf("pkg/rdb/sqlfiles/")
	rdb.SqlFilesPath = sqlFilesPath
	schemaTomlPath := fmt.Sprintf("pkg/rdb/schema.toml")
	connTomlPath := fmt.Sprintf("cmd/mig_for_ga/mig.toml")
	err := gomig.Exec(schemaTomlPath, env, connTomlPath, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rdb.Seed()

	os.Exit(0)
}

func target() bool {
	return true
}
