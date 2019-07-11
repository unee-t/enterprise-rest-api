package main

import (
	"net/http"

	"github.com/tj/go/http/response"
)

func (a *App) getUnitTypes(w http.ResponseWriter, r *http.Request) {
	UnitTypes, err := getUnitTypes(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, UnitTypes)
}
