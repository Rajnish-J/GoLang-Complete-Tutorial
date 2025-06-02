package db

import (
    "database/sql"
    "fmt"
    "go-assignment/config"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(cfg config.Config) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Failed to open DB: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping DB: %v", err)
    }

    log.Println("✅ Connected to MySQL")
}
