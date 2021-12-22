package db

import (
	model "example.com/logos106/saroop-api/models"
)

// Insert allows populating database
func InsertAdmin(admin model.Admin) string {
	db := setupDB()

	var lastInsertID int
	err := db.QueryRow("INSERT INTO admin(name, domain_id, password, role, status) VALUES($1, $2, $3, $4, $5) returning id;",
										admin.Name, admin.Domain, admin.Password, admin.Role, admin.Status).Scan(&lastInsertID)
	checkErr(err)

	return "The admin has been created successfully!"
}

// Get returns the whole database
func SelectAdmin() []model.Admin {
	db := setupDB()

	rows, err := db.Query("SELECT * FROM admin")
	checkErr(err)

  var admins []model.Admin

	for rows.Next() {
		var id int
		var name string
		var domain int
		var pass string
		var role string
		var status string

		err = rows.Scan(&id, &name, &domain, &pass, &role, &status)
		checkErr(err)
		admins = append(admins, model.Admin{ID: id, Name: name, Domain: domain, Password: pass, Role: role, Status: status})
	}

// close database
	defer db.Close()

	return admins
}
