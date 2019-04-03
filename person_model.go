// model.go

package main

import (
	"github.com/jmoiron/sqlx"
)

func (u *person) getperson(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM persons WHERE id_person=?", u.ID)
	return err
}

func (u *person) updateperson(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE persons SET name=?, age=? WHERE id=?", u.GivenName, u.FamilyName, u.ID)
	return err
}

func (u *person) deleteperson(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM persons WHERE id=?", u.ID)
	return err
}

func (u *person) createperson(db *sqlx.DB) error {
	result, err := db.Exec("insert into persons(given_name, family_name) values(?,?)",
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
	err = db.Select(&persons, "SELECT * FROM persons WHERE id_person >= ? LIMIT ?", startid, count)
	if err != nil {
		return persons, err
	}
	return persons, nil
}
