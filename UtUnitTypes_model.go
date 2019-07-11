package main

import (
	"github.com/jmoiron/sqlx"
)

func getUnitTypes(db *sqlx.DB) (types []UtUnitTypes, err error) {
	err = db.Select(&types, "SELECT designation FROM ut_unit_types")
	if err != nil {
		return
	}
	if len(types) == 0 {
		return []UtUnitTypes{}, nil
	}
	return
}
