package main

import "time"

type person struct {
	ID                   int        `db:"id_person"`
	ExternalId           string     `db:"external_id"`
	ExternalSystem       string     `db:"external_system"`
	ExternalTable        string     `db:"external_table"`
	SystCreatedDatetime  *time.Time `db:"syst_created_datetime"`
	CreationSystemId     *int       `db:"creation_system_id"`
	CreatedById          *string    `db:"created_by_id"`
	CreationMethod       *string    `db:"creation_method"`
	SystUpdatedDatetime  *time.Time `db:"syst_updated_datetime"`
	UpdateSystemId       *int       `db:"update_system_id"`
	UpdatedById          *string    `db:"updated_by_id"`
	UpdateMethod         *string    `db:"update_method"`
	PersonStatusId       *int       `db:"person_status_id"`
	DupeId               *int       `db:"dupe_id"`
	HandlerId            *int       `db:"handler_id"`
	IsUneeTAccountNeeded *int       `db:"is_unee_t_account_needed"`
	UneeTUserTypeId      *int       `db:"unee_t_user_type_id"`
	CountryCode          *string    `db:"country_code"`
	Gender               *int       `db:"gender"`
	SalutationId         *int       `db:"salutation_id"`
	GivenName            string     `db:"given_name"`
	MiddleName           *string    `db:"middle_name"`
	FamilyName           *string    `db:"family_name"`
	DateOfBirth          *time.Time `db:"date_of_birth"`
	Alias                *string    `db:"alias"`
	JobTitle             *string    `db:"job_title"`
	Organization         *string    `db:"organization"`
	Email                *string    `db:"email"`
	Tel1                 *string    `db:"tel_1"`
	Tel2                 *string    `db:"tel_2"`
	Whatsapp             *string    `db:"whatsapp"`
	Linkedin             *string    `db:"linkedin"`
	Facebook             *string    `db:"facebook"`
	Adr1                 *string    `db:"adr1"`
	Adr2                 *string    `db:"adr2"`
	Adr3                 *string    `db:"adr3"`
	City                 *string    `db:"City"`
	ZipPostcode          *string    `db:"zip_postcode"`
	RegionOrState        *string    `db:"region_or_state"`
	Country              *string    `db:"country"`
}
