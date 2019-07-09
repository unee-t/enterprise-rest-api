package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) getunit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	u := UnteApiAddUnit{ExternalID: vars["id"]}
	if err := u.getunit(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "unit not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}
