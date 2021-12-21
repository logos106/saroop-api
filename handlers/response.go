package handler

// import model "example.com/logos106/saroop-api/models"

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type jsonError struct {
	// in: body
	Message string `json:"message"`
}
