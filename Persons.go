package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Person struct {
	IDPerson             int            `db:"IdPerson"`
	ExternalID           string         `db:"ExternalId"`
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
