package common

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

//DBConnectionKey points to the DB connection in the context
type DBConnectionKey struct{}

//GetDatabaseConnection returns the database connection if it was found
// connection is nil if ctx has no value for the key
// The type assertion *mongodb.Connection returns false for nil
func GetDatabaseConnection(ctx context.Context) (*gorm.DB, bool) {
	connection, ok := ctx.Value(DBConnectionKey{}).(*gorm.DB)
	return connection, ok
}

//ErrorResponse represents the struct error response.
type ErrorResponse struct {
	Error      string `json:"error"`
	Descripion string `json:"error_descripion"`
}

// OkJSON encodes a JSON with a 200 status code
func OkJSON(w http.ResponseWriter, v interface{}, gzipCompression bool) {
	// sending a JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if gzipCompression == true {
		// add the Content-Encoding header
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(http.StatusOK)

		// gzip compression
		gz := gzip.NewWriter(w)
		defer gz.Close()
		// encode v as JSON
		json.NewEncoder(gz).Encode(v)
	} else {
		w.WriteHeader(http.StatusOK)
		// encode V as JSON
		json.NewEncoder(w).Encode(v)
	}
}

// KOJSON encodes a JSON with a error status code
func KOJSON(w http.ResponseWriter, status int, description string) {

	// sending a JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var error string

	switch status {
	case http.StatusNoContent:
		error = "No content"
	case http.StatusForbidden:
		http.Error(w, http.StatusText(status), status)
	case http.StatusUnauthorized:
		error = "Unauthorized"
		description = "Bad username or password"
	case http.StatusBadRequest:
		error = "Bad request"
	case http.StatusNotFound:
		error = "Not Found"
	case http.StatusMethodNotAllowed:
		error = "Method not allowed"
	case http.StatusNotAcceptable:
		error = "Not acceptable"
	case http.StatusInternalServerError:
		error = "Internal error"
	default:
		error = "Internal error"
		description = "Internal error"
	}

	response := ErrorResponse{
		Error:      error,
		Descripion: description,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)

}
