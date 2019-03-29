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
	ExternalSystem       string         `db:"external_system"`
	ExternalTable        string         `db:"external_table"`
	SystCreatedDatetime  mysql.NullTime `db:"syst_created_datetime"`
	CreationSystemID     sql.NullInt64  `db:"creation_system_id"`
	CreatedByID          sql.NullString `db:"created_by_id"`
	SystUpdatedDatetime  mysql.NullTime `db:"syst_updated_datetime"`
	UpdateSystemID       sql.NullInt64  `db:"update_system_id"`
	UpdatedByID          sql.NullString `db:"updated_by_id"`
	PersonStatusID       sql.NullInt64  `db:"person_status_id"`
	DupeID               sql.NullInt64  `db:"dupe_id"`
	HandlerID            sql.NullInt64  `db:"handler_id"`
	IsUneeTAccountNeeded sql.NullInt64  `db:"is_unee_t_account_needed"`
	UneeTUserTypeID      sql.NullInt64  `db:"unee_t_user_type_id"`
	CountryCode          sql.NullString `db:"country_code"`
	Gender               sql.NullInt64  `db:"gender"`
	SalutationID         sql.NullInt64  `db:"salutation_id"`
	GivenName            string         `db:"given_name"`
	MiddleName           sql.NullString `db:"middle_name"`
	FamilyName           sql.NullString `db:"family_name"`
	DateOfBirth          mysql.NullTime `db:"date_of_birth"`
	Alias                sql.NullString `db:"alias"`
	JobTitle             sql.NullString `db:"job_title"`
	Organization         sql.NullString `db:"organization"`
	Email                sql.NullString `db:"email"`
	Tel1                 sql.NullString `db:"tel_1"`
	Tel2                 sql.NullString `db:"tel_2"`
	Whatsapp             sql.NullString `db:"whatsapp"`
	Linkedin             sql.NullString `db:"linkedin"`
	Facebook             sql.NullString `db:"facebook"`
	Adr1                 sql.NullString `db:"adr1"`
	Adr2                 sql.NullString `db:"adr2"`
	Adr3                 sql.NullString `db:"adr3"`
	City                 sql.NullString `db:"City"`
	ZipPostcode          sql.NullString `db:"zip_postcode"`
	RegionOrState        sql.NullString `db:"region_or_state"`
	Country              sql.NullString `db:"country"`
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
