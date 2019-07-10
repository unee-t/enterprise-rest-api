// model.go

package main

import (
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

// https://github.com/unee-t/enterprise-rest-api/issues/10
func (u *UnteApiAddUnit) getunit(db *sqlx.DB) (err error) {
	if u.ExternalID != "" {
		log.WithField("external_id", u.ExternalID).Info("select")
		err = db.Get(u, "SELECT * FROM unte_api_add_unit WHERE external_id=?", u.ExternalID)
	} else if u.RequestID != "" {
		log.WithField("request_id", u.RequestID).Info("select")
		err = db.Get(u, "SELECT * FROM unte_api_add_unit WHERE request_id=?", u.RequestID)
	}
	return err
}

// https://github.com/unee-t/enterprise-rest-api/issues/3
func (u *UnteApiAddUnit) createunit(db *sqlx.DB) error {
	result, err := db.NamedExec(`INSERT INTO unte_api_add_unit (
external_id,
organization_key,
unee_t_unit_type,
designation,
country_code,
mgt_cny_default_assignee,
external_system_id,
external_table,
creation_method,
parent_mefe_id,
tower,
unit_id,
address_1,
address_2,
zip_postal_code,
state,
city,
description,
count_rooms,
surface,
surface_measurement_unit,
number_of_beds,
landlord_default_assignee,
tenant_default_assignee,
agent_default_assignee,
request_id,
syst_created_datetime)
	VALUES (
:external_id,
:organization_key,
:unee_t_unit_type,
:designation,
:country_code,
:mgt_cny_default_assignee,
:external_system_id,
:external_table,
:creation_method,
:parent_mefe_id,
:tower,
:unit_id,
:address_1,
:address_2,
:zip_postal_code,
:state,
:city,
:description,
:count_rooms,
:surface,
:surface_measurement_unit,
:number_of_beds,
:landlord_default_assignee,
:tenant_default_assignee,
:agent_default_assignee,
	uuid(),
	now())`, u)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	err = db.Get(&u.RequestID, "SELECT request_id FROM unte_api_add_unit WHERE id_unte_api_add_unit=?", int(id))
	return err
}
