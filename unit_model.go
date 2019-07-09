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
	result, err := db.NamedExec(`INSERT INTO unte_api_add_unit (external_id,
	designation,
	unee_t_unit_type,
	country_code,
	mgt_cny_default_assignee,
	request_id,
	syst_created_datetime,
	organization_key)
	VALUES (:external_id,
	:designation,
	:unee_t_unit_type,
	:country_code,
	:mgt_cny_default_assignee,
	uuid(),
	now(),
	:organization_key)`, u)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	err = db.Get(u, "SELECT * FROM unte_api_add_unit WHERE id_unte_api_add_unit=?", int(id))
	return err
}
