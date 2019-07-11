package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/unee-t/env"
)

type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func New() (a App, err error) {
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
	var connectionString string
	dbOptions := "?multiStatements=true&sql_mode=TRADITIONAL&timeout=5s&parseTime=true&collation=utf8mb4_unicode_520_ci"
	if os.Getenv("UP_STAGE") == "" {
		connectionString = "root:secret@tcp(localhost:3306)/unee_t_enterprise" + dbOptions
	} else {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/unee_t_enterprise%s",
			"root",
			e.GetSecret("MYSQL_ROOT_PASSWORD"),
			// e.GetSecret("UNEE-T_ENTERPRISE_RDS_MASTER_USER"),
			// e.GetSecret("UNEE-T_ENTERPRISE_RDS_MASTER_USER_PASSWORD"),
			e.Udomain("auroradb"),
			dbOptions)
	}
	log.WithField("connection", connectionString).Info("Connecting to database")
	a.DB, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		return a, err
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	return a, err
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/unit", a.createunit).Methods("POST")
	a.Router.HandleFunc("/unitinfo", a.unitcreatedinfo).Methods("POST")
	a.Router.HandleFunc("/countries", a.getcountries).Methods("GET")
	a.Router.HandleFunc("/unittypes", a.getUnitTypes).Methods("GET")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, " ", "   ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
