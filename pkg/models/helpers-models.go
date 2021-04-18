package models

// ReturnDBHelpers will return the file to create the db struct
// that is utilized by the services for interface chaining
func ReturnDBHelpers() ([]byte, error) {
	return []byte(`package helpers

import (
	"github.com/jmoiron/sqlx"
)
	
/**********************************************************************
/
/	Helpers for DB
/
/**********************************************************************/
	
// DBContext is for passing gorm to db layer
type DBContext struct {
	DB *sqlx.DB
}`), nil
}

// ReturnResponseHelpers will return the generic response
// files that are utilized by the controllers
func ReturnResponseHelpers() ([]byte, error) {
	return []byte(`package helpers

import (
	"encoding/json"
	"net/http"
)
	
// SendSuccessHeader is a generic method for setting and sending headers to
// our front end
func SendSuccessHeader(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
	
// SendErrorHeader is a generic method for setting and sending errors to
// the front end
func SendErrorHeader(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}`), nil
}
