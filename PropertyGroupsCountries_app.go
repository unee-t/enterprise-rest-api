package main

import (
	"net/http"

	"github.com/tj/go/http/response"
)

func (a *App) getcountries(w http.ResponseWriter, r *http.Request) {
	Countries, err := getcountries(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, Countries)
}
