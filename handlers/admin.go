package handler

import (
	"encoding/json"
	"net/http"
	"time"
  "strconv"

	"github.com/gorilla/mux"
	db "example.com/logos106/saroop-api/db"
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

// GetPeople is an httpHandler for route GET /people
func GetAdmins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Get())
}

// GetPerson is an httpHandler for route GET /people/{id}
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.Get() {
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
