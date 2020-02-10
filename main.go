package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var configFilename string
var config tomlConfig

type tomlConfig struct {
	PostgresDb PostgresConfig `toml:"postgres"`
}

//PostgresConfig postgres config
type PostgresConfig struct {
	Host string `toml:"host"`
	User string `toml:"user"`
	Pwd  string `toml:"pwd"`
	Db   string `toml:"db"`
}

func init() {

	// point to the right configuration file
	env, ok := os.LookupEnv("ENV")
	if ok && env == "production" {
		configFilename = "./config.production.toml"
	} else if ok && env == "development" {
		configFilename = "./config.development.toml"
	} else {
		configFilename = "./config.local.toml"
	}

	// read toml config file
	_, err := toml.DecodeFile(configFilename, &config)
	if err != nil {
		log.Fatal("Could not decode the toml configuration file")
	}
}
func main() {
	//Init Postgres connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.PostgresDb.Host, 5432, config.PostgresDb.User, config.PostgresDb.Pwd, config.PostgresDb.Db)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	router := NewRouter(db)

	// deploy on the right port
	var port string
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9090"
	}
	env, ok := os.LookupEnv("ENV")
	if ok && (env == "production" || env == "development") {
		log.Fatal(http.ListenAndServe(":"+port, router))
	} else {
		// local
		log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
	}

}
