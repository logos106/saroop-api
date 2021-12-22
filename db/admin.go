package db

import (
	"database/sql"
	model "example.com/logos106/saroop-api/models"
)

var admin []model.Admin

// Insert allows populating database
func Insert(person model.Admin) {
	admin = append(admin, person)
}

// Get returns the whole database
func Get() []model.Admin {
	db := setupDB
	checkErr(err)

	return admin
}
