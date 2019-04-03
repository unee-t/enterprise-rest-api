// model.go

package main

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type person struct {
	ID                   int        `db:"id_person" json:"id"`
	ExternalID           *string    `db:"external_id" json:"external_id,omitempty"`
	ExternalSystem       *string    `db:"external_system" json:"external_system,omitempty"`
	ExternalTable        *string    `db:"external_table" json:"external_table,omitempty"`
	SystCreatedDatetime  *time.Time `db:"syst_created_datetime" json:"syst_created_datetime,omitempty"`
	CreationSystemID     *int       `db:"creation_system_id" json:"creation_system_id,omitempty"`
	CreatedByID          *string    `db:"created_by_id" json:"created_by_id,omitempty"`
	SystUpdatedDatetime  *time.Time `db:"syst_updated_datetime" json:"syst_updated_datetime,omitempty"`
	UpdateSystemID       *int       `db:"update_system_id" json:"update_system_id,omitempty"`
	UpdatedByID          *string    `db:"updated_by_id" json:"updated_by_id,omitempty"`
	PersonStatusID       *int       `db:"person_status_id" json:"person_status_id,omitempty"`
	DupeID               *int       `db:"dupe_id" json:"dupe_id,omitempty"`
	HandlerID            *int       `db:"handler_id" json:"handler_id,omitempty"`
	IsUneeTAccountNeeded *int       `db:"is_unee_t_account_needed" json:"is_unee_t_account_needed,omitempty"`
	UneeTUserTypeID      *int       `db:"unee_t_user_type_id" json:"unee_t_user_type_id,omitempty"`
	CountryCode          *string    `db:"country_code" json:"country_code,omitempty"`
	Gender               *int       `db:"gender" json:"gender,omitempty"`
	SalutationID         *int       `db:"salutation_id" json:"salutation_id,omitempty"`
	GivenName            string     `db:"given_name" json:"given_name,omitempty"`
	MiddleName           *string    `db:"middle_name" json:"middle_name,omitempty"`
	FamilyName           *string    `db:"family_name" json:"family_name,omitempty"`
	DateOfBirth          *time.Time `db:"date_of_birth" json:"date_of_birth,omitempty"`
	Alias                *string    `db:"alias" json:"alias,omitempty"`
	JobTitle             *string    `db:"job_title" json:"job_title,omitempty"`
	Organization         *string    `db:"organization" json:"organization,omitempty"`
	Email                *string    `db:"email" json:"email,omitempty"`
	Tel1                 *string    `db:"tel_1" json:"tel_1,omitempty"`
	Tel2                 *string    `db:"tel_2" json:"tel_2,omitempty"`
	Whatsapp             *string    `db:"whatsapp" json:"whatsapp,omitempty"`
	Linkedin             *string    `db:"linkedin" json:"linkedin,omitempty"`
	Facebook             *string    `db:"facebook" json:"facebook,omitempty"`
	Adr1                 *string    `db:"adr1" json:"adr_1,omitempty"`
	Adr2                 *string    `db:"adr2" json:"adr_2,omitempty"`
	Adr3                 *string    `db:"adr3" json:"adr_3,omitempty"`
	City                 *string    `db:"City" json:"city,omitempty"`
	ZipPostcode          *string    `db:"zip_postcode" json:"zip_postcode,omitempty"`
	RegionOrState        *string    `db:"region_or_state" json:"region_or_state,omitempty"`
	Country              *string    `db:"country" json:"country,omitempty"`
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
