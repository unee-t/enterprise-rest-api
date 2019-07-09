package main

import (
	"time"
)

type UnteApiAddUnit struct {
	IDUnteApiAddUnit        int        `db:"id_unte_api_add_unit" json:"id_unte_api_add_unit,omitempty"`
	RequestID               *string    `db:"request_id" json:"request_id,omitempty"`
	ExternalID              string     `db:"external_id" json:"external_id,omitempty"`
	ExternalSystemID        string     `db:"external_system_id" json:"external_system_id,omitempty"`
	ExternalTable           string     `db:"external_table" json:"external_table,omitempty"`
	SystCreatedDatetime     *time.Time `db:"syst_created_datetime" json:"syst_created_datetime,omitempty"`
	CreationSystemID        *string    `db:"creation_system_id" json:"creation_system_id,omitempty"`
	OrganizationKey         string     `db:"organization_key" json:"organization_key,omitempty"`
	CreationMethod          *string    `db:"creation_method" json:"creation_method,omitempty"`
	SystUpdatedDatetime     *time.Time `db:"syst_updated_datetime" json:"syst_updated_datetime,omitempty"`
	UpdateSystemID          *string    `db:"update_system_id" json:"update_system_id,omitempty"`
	UpdatedByID             *int       `db:"updated_by_id" json:"updated_by_id,omitempty"`
	UpdateMethod            *string    `db:"update_method" json:"update_method,omitempty"`
	IsObsolete              *int       `db:"is_obsolete" json:"is_obsolete,omitempty"`
	Order                   *int       `db:"order" json:"order,omitempty"`
	ParentMefeID            *string    `db:"parent_mefe_id" json:"parent_mefe_id,omitempty"`
	UneeTUnitType           *string    `db:"unee_t_unit_type" json:"unee_t_unit_type,omitempty"`
	Designation             string     `db:"designation" json:"designation,omitempty"`
	Tower                   string     `db:"tower" json:"tower,omitempty"`
	UnitID                  *string    `db:"unit_id" json:"unit_id,omitempty"`
	Address1                *string    `db:"address_1" json:"address_1,omitempty"`
	Address2                *string    `db:"address_2" json:"address_2,omitempty"`
	ZipPostalCode           *string    `db:"zip_postal_code" json:"zip_postal_code,omitempty"`
	State                   *string    `db:"state" json:"state,omitempty"`
	City                    *string    `db:"city" json:"city,omitempty"`
	CountryCode             *string    `db:"country_code" json:"country_code,omitempty"`
	Description             *string    `db:"description" json:"description,omitempty"`
	CountRooms              *int       `db:"count_rooms" json:"count_rooms,omitempty"`
	Surface                 *int       `db:"surface" json:"surface,omitempty"`
	SurfaceMeasurementUnit  *string    `db:"surface_measurement_unit" json:"surface_measurement_unit,omitempty"`
	NumberOfBeds            *int       `db:"number_of_beds" json:"number_of_beds,omitempty"`
	MgtCnyDefaultAssignee   *string    `db:"mgt_cny_default_assignee" json:"mgt_cny_default_assignee,omitempty"`
	LandlordDefaultAssignee *string    `db:"landlord_default_assignee" json:"landlord_default_assignee,omitempty"`
	TenantDefaultAssignee   *string    `db:"tenant_default_assignee" json:"tenant_default_assignee,omitempty"`
	AgentDefaultAssignee    *string    `db:"agent_default_assignee" json:"agent_default_assignee,omitempty"`
	IsCreationNeededInUneeT *int       `db:"is_creation_needed_in_unee_t" json:"is_creation_needed_in_unee_t,omitempty"`
	MefeUnitID              *string    `db:"mefe_unit_id" json:"mefe_unit_id,omitempty"`
	UneetCreatedDatetime    *time.Time `db:"uneet_created_datetime" json:"uneet_created_datetime,omitempty"`
	IsApiSuccess            *int       `db:"is_api_success" json:"is_api_success,omitempty"`
	ApiErrorMessage         *string    `db:"api_error_message" json:"api_error_message,omitempty"`
}
