// model.go

package main

import (
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

func (u *UnteApiAddUnit) getunit(db *sqlx.DB) error {
	log.WithField("external_id", u.ExternalID).Info("select")
	err := db.Get(u, "SELECT * FROM unte_api_add_unit WHERE external_id=?", u.ExternalID)
	return err
}

func (u *UnteApiAddUnit) createunit(db *sqlx.DB) error {
	_, err := db.NamedExec(`INSERT INTO unte_api_add_unit (external_id,
	designation,
	organization_key)
	VALUES (:external_id,
	:designation,
	:organization_key)`, u)
	if err != nil {
		return err
	}
	return err
}
