package main

import (
	"net/http"

	"github.com/tj/go/http/response"
)

func (a *App) getUnitTypes(w http.ResponseWriter, r *http.Request) {
	_, all := r.URL.Query()["all"]
	UnitTypes, err := getUnitTypes(a.DB, all)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, UnitTypes)
}
