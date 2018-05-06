package model

import (
    "database/sql"
    "log"

    // mysql driver
    _ "github.com/go-sql-driver/mysql"
)

// DBConnect returns *sql.DB
func DBConnect() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root" // 環境変数から取得
    dbName := "go_crud_sample"
    dbOption := "?parseTime=true"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
    if err != nil {
        log.Fatal(err)
    }
    return db
}
