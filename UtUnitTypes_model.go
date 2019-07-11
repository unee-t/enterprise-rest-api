package main

import (
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

func getUnitTypes(db *sqlx.DB, all bool) (types []UtUnitTypes, err error) {
	if all {
		log.Info("all including obsolete records")
		err = db.Select(&types, "SELECT designation FROM ut_unit_types")
	} else { // active is default
		log.Info("active records")
		err = db.Select(&types, "SELECT designation FROM ut_unit_types WHERE is_obsolete = 0")
	}
	if err != nil {
		return
	}
	if len(types) == 0 {
		return []UtUnitTypes{}, nil
	}
	return
}
