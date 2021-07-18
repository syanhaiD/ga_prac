package rdb

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type DbConfig struct {
	Test DbInfo
}
type DbInfo struct {
	Mysql MysqlInfo
}
type MysqlInfo struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

var DbConn *sql.DB

type TranConn struct {
	Tx *sql.Tx
}

func Connect() {
	dbInfo := newDbConfig()
	dsn := fmt.Sprintf(
		`%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8mb4,utf8&collation=utf8mb4_general_ci`,
		dbInfo.Mysql.User,
		dbInfo.Mysql.Pass,
		dbInfo.Mysql.Host,
		dbInfo.Mysql.Port,
		dbInfo.Mysql.Db,
	)

	var err error
	DbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func Close() {
	_ = DbConn.Close()
}

var DatabaseTomlPath = `./database.toml`
var Env = "test"

func newDbConfig() DbInfo {
	var dbConfig DbConfig
	_, err := toml.DecodeFile(DatabaseTomlPath, &dbConfig)
	if err != nil {
		panic(err)
	}
	v := reflect.ValueOf(dbConfig)
	currentDbInfo := v.FieldByName(strings.Title(Env)).Interface().(DbInfo)

	return currentDbInfo
}
