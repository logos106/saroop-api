package handler

import (
	"encoding/json"
	"net/http"
	"time"
  "strconv"
	"fmt"

	"github.com/gorilla/mux"
	db "example.com/logos106/saroop-api/db"
	model "example.com/logos106/saroop-api/models"
)

// IDParam is used to identify a person
//
// swagger:parameters listPerson
type IDParam struct {
	// The ID of a person
	//
	// in: path
	// required: true
	ID int64 `json:"id"`
}

// SetAdmin is an httpHandler for route POST /admin
func SetAdmin(w http.ResponseWriter, r *http.Request) {
	// Get parameters
	// id := r.FormValue("id")
  // name := r.FormValue("name")
  // domain := r.FormValue("domain")
  // pass := r.FormValue("password")
  // role := r.FormValue("role")
  // status := r.FormValue("status")
	//
	// admin = Admin(ID: id, Name: name, Domain: domain, Password: pass, Role: role, Status: status)

	var admin model.Admin
	err := decodeJSONBody(w, r, &admin)
	fmt.Printf("%+v\n", admin)
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.InsertAdmin(admin))
}

// GetAdmins is an httpHandler for route GET /admins
func GetAdmins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.SelectAdmin())
}

// GetAdmin is an httpHandler for route GET /admin/{id}
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.SelectAdmin() {
    intVar, _ := strconv.Atoi(params["id"])
		if item.ID == intVar {
			w.WriteHeader(http.StatusOK)
			// add a arbitraty pause of 1 second
			time.Sleep(1000 * time.Millisecond)
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}
