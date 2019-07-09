package main

import (
	"database/sql"
	"time"
)

type UnteApiAddUnit struct {
	IDUnteApiAddUnit        int            `db:"id_unte_api_add_unit"`
	RequestID               *string        `db:"request_id"`
	ExternalID              string         `db:"external_id"`
	ExternalSystemID        string         `db:"external_system_id"`
	ExternalTable           string         `db:"external_table"`
	SystCreatedDatetime     *time.Time     `db:"syst_created_datetime"`
	CreationSystemID        *string        `db:"creation_system_id"`
	OrganizationKey         string         `db:"organization_key"`
	CreationMethod          *string        `db:"creation_method"`
	SystUpdatedDatetime     *time.Time     `db:"syst_updated_datetime"`
	UpdateSystemID          *string        `db:"update_system_id"`
	UpdatedByID             *int           `db:"updated_by_id"`
	UpdateMethod            *string        `db:"update_method"`
	IsObsolete              *int           `db:"is_obsolete"`
	Order                   *int           `db:"order"`
	ParentMefeID            *string        `db:"parent_mefe_id"`
	UneeTUnitType           *string        `db:"unee_t_unit_type"`
	Designation             string         `db:"designation"`
	Tower                   string         `db:"tower"`
	UnitID                  *string        `db:"unit_id"`
	Address1                *string        `db:"address_1"`
	Address2                *string        `db:"address_2"`
	ZipPostalCode           *string        `db:"zip_postal_code"`
	State                   *string        `db:"state"`
	City                    *string        `db:"city"`
	CountryCode             *string        `db:"country_code"`
	Description             *string        `db:"description"`
	CountRooms              *int           `db:"count_rooms"`
	Surface                 *int           `db:"surface"`
	SurfaceMeasurementUnit  *string        `db:"surface_measurement_unit"`
	NumberOfBeds            *int           `db:"number_of_beds"`
	MgtCnyDefaultAssignee   *string        `db:"mgt_cny_default_assignee"`
	LandlordDefaultAssignee *string        `db:"landlord_default_assignee"`
	TenantDefaultAssignee   *string        `db:"tenant_default_assignee"`
	AgentDefaultAssignee    *string        `db:"agent_default_assignee"`
	IsCreationNeededInUneeT *int           `db:"is_creation_needed_in_unee_t"`
	MefeUnitID              *string        `db:"mefe_unit_id"`
	UneetCreatedDatetime    *time.Time     `db:"uneet_created_datetime"`
	IsApiSuccess            *sql.NullInt64 `db:"is_api_success"`
	ApiErrorMessage         *string        `db:"api_error_message"`
}
