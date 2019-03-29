// app.go

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

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
	connectionString := "root:secret@tcp(localhost:3306)/unee_t_enterprise?multiStatements=true&sql_mode=TRADITIONAL&timeout=5s"
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

func (a *App) getpersons(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	startid, _ := strconv.Atoi(r.FormValue("id"))

	if count > 10 || count < 1 {
		count = 10
	}
	if startid < 0 {
		startid = 0
	}

	Persons, err := getpersons(a.DB, startid, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, Persons)
}

func (a *App) createperson(w http.ResponseWriter, r *http.Request) {
	var u person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := u.createperson(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) getperson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid person ID")
		return
	}

	u := person{ID: id}
	if err := u.getperson(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "person not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) updateperson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid person ID")
		return
	}

	var u person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	u.ID = id

	if err := u.updateperson(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) deleteperson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid person ID")
		return
	}

	u := person{ID: id}
	if err := u.deleteperson(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
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
