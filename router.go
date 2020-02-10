package main

import (
	"github.com/gorilla/mux"
	gorm "github.com/jinzhu/gorm"
)

// NewRouter provides a new router instance
func NewRouter(db *gorm.DB) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		handler := route.Handler
		logger := Logger(handler, route.Name, route.Auth, db)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Queries(route.Queries...).
			Handler(logger)
	}

	return router
}
