package main

import (
	"github.com/jmoiron/sqlx"
)

func getcountries(db *sqlx.DB) (countries []PropertyGroupsCountries, err error) {
	err = db.Select(&countries, "SELECT country_code, country_name FROM property_groups_countries")
	if err != nil {
		return
	}
	if len(countries) == 0 {
		return []PropertyGroupsCountries{}, nil
	}
	return
}
