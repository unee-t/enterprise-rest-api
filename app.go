// app.go

package main

import (
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/unee-t/env"
)

type App struct {
	Router         *mux.Router
	DB             *sqlx.DB
	APIAccessToken string
}

func New() (a App, err error) {
	connectionString := "root:secret@tcp(localhost:3306)/unee_t_enterprise?multiStatements=true&sql_mode=TRADITIONAL&timeout=5s&parseTime=true"
	a.DB, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		return a, err
	}
	cfg, err := external.LoadDefaultAWSConfig(external.WithSharedConfigProfile("uneet-dev"))
	if err != nil {
		log.WithError(err).Fatal("setting up credentials")
		return
	}
	cfg.Region = endpoints.ApSoutheast1RegionID
	e, err := env.New(cfg)
	if err != nil {
		log.WithError(err).Warn("error getting AWS unee-t env")
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.APIAccessToken = e.GetSecret("API_ACCESS_TOKEN")
	return a, err
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/persons", a.getpersons).Methods("GET")
	a.Router.HandleFunc("/person", a.createperson).Methods("POST")
	a.Router.HandleFunc("/person/{id:[0-9]+}", a.getperson).Methods("GET")
	a.Router.HandleFunc("/person/{id:[0-9]+}", a.updateperson).Methods("PUT")
	a.Router.HandleFunc("/person/{id:[0-9]+}", a.deleteperson).Methods("DELETE")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
