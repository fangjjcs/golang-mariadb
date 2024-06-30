package main

import (
	"app/internal/driver"
	"app/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type Config struct {
	Port int    `json:"port"`
	ENV  string `json:"env"`
	DSN  string `json:"dsn"`
}

type application struct {
	config   Config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	DB       models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.Port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Println(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.ENV, app.config.Port))

	return srv.ListenAndServe()
}

func main() {
	var cfg Config

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open Json Config
	pwd, _ := os.Getwd()
	configFile, err := os.Open(pwd + "/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &cfg)

	conn, err := driver.OpenDB(cfg.DSN)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		DB:       models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
