package main

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/ReferenceShop/auth"
	"github.com/ReferenceShop/common"
	"github.com/ReferenceShop/model"
	"github.com/ReferenceShop/user"
	gorm "github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Logger logs all the requests
func Logger(inner func(http.ResponseWriter, *http.Request) (int, string, error), name string, hasAuth bool, dbConnection *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ctx := r.Context()

		if hasAuth {
			err := auth.TokenValid(r)
			if err != nil {
				common.KOJSON(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
		}

		// log the URL (proper formatting) and the method
		url, _ := url.QueryUnescape(r.URL.String())
		requestLogger := logrus.WithFields(logrus.Fields{"url": url, "handler": name, "method": r.Method})

		//add db connexion in the context
		ctx = context.WithValue(ctx, common.DBConnectionKey{}, dbConnection)
		dbConnection.Debug().AutoMigrate(&model.Category{})
		dbConnection.Debug().AutoMigrate(&model.Product{})
		dbConnection.Debug().AutoMigrate(&model.Variante{})
		dbConnection.Debug().AutoMigrate(&model.VarianteValue{})
		dbConnection.Debug().AutoMigrate(&model.ProductStock{})
		dbConnection.Debug().AutoMigrate(&model.ProductPrice{})
		dbConnection.Debug().AutoMigrate(&model.ProductImage{})
		dbConnection.Debug().AutoMigrate(&user.User{})

		r = r.WithContext(ctx)

		//Log
		status, _, err := inner(w, r)
		duration := time.Since(start).Seconds()

		requestLogger = requestLogger.WithFields(logrus.Fields{"code": status, "took": duration})
		if err != nil {
			// log the error
			requestLogger = requestLogger.WithField("error", err.Error())
			requestLogger.Error()
			return
		}
		requestLogger.Info()
	})
}
