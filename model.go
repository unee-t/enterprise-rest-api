// model.go

package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type person struct {
	ID                   int            `db:"id_person"`
	ExternalID           string         `db:"external_id"`
	ExternalSystem       string         `db:"ExternalSystem"`
	ExternalTable        string         `db:"ExternalTable"`
	SystCreatedDatetime  mysql.NullTime `db:"SystCreatedDatetime"`
	CreationSystemID     sql.NullInt64  `db:"CreationSystemId"`
	CreatedByID          sql.NullString `db:"CreatedById"`
	SystUpdatedDatetime  mysql.NullTime `db:"SystUpdatedDatetime"`
	UpdateSystemID       sql.NullInt64  `db:"UpdateSystemId"`
	UpdatedByID          sql.NullString `db:"UpdatedById"`
	PersonStatusID       sql.NullInt64  `db:"PersonStatusId"`
	DupeID               sql.NullInt64  `db:"DupeId"`
	HandlerID            sql.NullInt64  `db:"HandlerId"`
	IsUneeTAccountNeeded sql.NullInt64  `db:"IsUneeTAccountNeeded"`
	UneeTUserTypeID      sql.NullInt64  `db:"UneeTUserTypeId"`
	CountryCode          sql.NullString `db:"CountryCode"`
	Gender               sql.NullInt64  `db:"Gender"`
	SalutationID         sql.NullInt64  `db:"SalutationId"`
	GivenName            string         `db:"GivenName"`
	MiddleName           sql.NullString `db:"MiddleName"`
	FamilyName           sql.NullString `db:"FamilyName"`
	DateOfBirth          mysql.NullTime `db:"DateOfBirth"`
	Alias                sql.NullString `db:"Alias"`
	JobTitle             sql.NullString `db:"JobTitle"`
	Organization         sql.NullString `db:"Organization"`
	Email                sql.NullString `db:"Email"`
	Tel1                 sql.NullString `db:"Tel1"`
	Tel2                 sql.NullString `db:"Tel2"`
	Whatsapp             sql.NullString `db:"Whatsapp"`
	Linkedin             sql.NullString `db:"Linkedin"`
	Facebook             sql.NullString `db:"Facebook"`
	Adr1                 sql.NullString `db:"Adr1"`
	Adr2                 sql.NullString `db:"Adr2"`
	Adr3                 sql.NullString `db:"Adr3"`
	City                 sql.NullString `db:"City"`
	ZipPostcode          sql.NullString `db:"ZipPostcode"`
	RegionOrState        sql.NullString `db:"RegionOrState"`
	Country              sql.NullString `db:"Country"`
}

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
	result, err := db.Exec("insert into persons(name, age) values(?,?)",
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
