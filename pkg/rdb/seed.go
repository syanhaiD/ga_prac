package rdb

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var SqlFilesPath = "./sqlfiles/"

func Seed() {
	allFiles, err := ioutil.ReadDir(SqlFilesPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var sqlFiles []string
	for _, file := range allFiles {
		if !file.IsDir() && filepath.Ext(file.Name()) == `.sql` && !strings.HasPrefix(file.Name(), "_") {
			sqlFiles = append(sqlFiles, SqlFilesPath+file.Name())
		}
	}

	for _, sqlFile := range sqlFiles {
		sqlBytes, err := ioutil.ReadFile(sqlFile)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sqlString := string(sqlBytes)
		_, err = DbConn.Exec(sqlString)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}
