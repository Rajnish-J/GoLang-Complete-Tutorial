package utility

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() *sql.DB{

	// dsn := "root:Saranya@2001@tcp(127.0.0.1:3306)/golangcui"
	dsn := "root:Saranya@2001@tcp(127.0.0.1:3306)/golangcui?parseTime=true"
    db, err :=sql.Open("mysql",dsn)

	if err != nil{
		log.Fatal("DB Connection failed :", err)
	} 

	if err:=db.Ping(); err != nil{
		log.Fatal("Failed to ping DB:",err)
	}
	
	// fmt.Println("Connected to MySQL Database!")
	InfoLogger.Println("Database connected");
	return db
}