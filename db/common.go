package db

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
	DB_HOST			= "localhost"
	DB_PORT			= 5534
	DB_USER     = "postgres"
	DB_PASSWORD = "glowglow"
	DB_NAME     = "wang"
)

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}
