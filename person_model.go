// model.go

package main

import (
	"github.com/jmoiron/sqlx"
)

func (u *person) getperson(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM external_persons WHERE id_person=?", u.ID)
	return err
}

func (u *person) updateperson(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE external_persons SET external_id=?, external_system=?, external_table=?, given_name=?, family_name=? WHERE id_person=?",
		u.ExternalID,
		u.ExternalSystem,
		u.ExternalTable,
		u.GivenName,
		u.FamilyName,
		u.ID)
	return err
}

func (u *person) deleteperson(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM external_persons WHERE id_person=?", u.ID)
	return err
}

func (u *person) createperson(db *sqlx.DB) error {
	result, err := db.Exec("insert into external_persons(external_id, external_system, external_table, created_by_id, given_name, family_name) values(?,?,?,?,?,?)",
		u.ExternalID,
		u.ExternalSystem,
		u.ExternalTable,
		u.CreatedByID,
		u.GivenName,
		u.FamilyName)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = int(id)
	return err
}

func getpersons(db *sqlx.DB, startid, count int) (persons []person, err error) {
	err = db.Select(&persons, "SELECT * FROM external_persons WHERE id_person >= ? ORDER BY id_person ASC LIMIT ?", startid, count)
	if err != nil {
		return persons, err
	}
	if len(persons) == 0 {
		return []person{}, nil
	}
	return persons, nil
}
