package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

func (a *App) getunit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

func (a *App) createunit(w http.ResponseWriter, r *http.Request) {
	var u UnteApiAddUnit
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		log.WithError(err).Error("bad payload")
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := u.createunit(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}
