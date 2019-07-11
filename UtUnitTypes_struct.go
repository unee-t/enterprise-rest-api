package main

import "time"

type UtUnitTypes struct {
	IDPropertyType      int        `db:"id_property_type" json:"id_property_type,omitempty"`
	SystCreatedDatetime *time.Time `db:"syst_created_datetime" json:"syst_created_datetime,omitempty"`
	CreationSystemID    int        `db:"creation_system_id" json:"creation_system_id,omitempty"`
	CreatedByID         *string    `db:"created_by_id" json:"created_by_id,omitempty"`
	SystUpdatedDatetime *time.Time `db:"syst_updated_datetime" json:"syst_updated_datetime,omitempty"`
	UpdateSystemID      *int       `db:"update_system_id" json:"update_system_id,omitempty"`
	UpdatedByID         *string    `db:"updated_by_id" json:"updated_by_id,omitempty"`
	Order               *int       `db:"order" json:"order,omitempty"`
	IsLevel1            *int       `db:"is_level_1" json:"is_level_1,omitempty"`
	IsLevel2            *int       `db:"is_level_2" json:"is_level_2,omitempty"`
	IsLevel3            *int       `db:"is_level_3" json:"is_level_3,omitempty"`
	IsObsolete          *int       `db:"is_obsolete" json:"inactive,omitempty"`
	Designation         string     `db:"designation" json:"designation,omitempty"`
	Description         *string    `db:"description" json:"description,omitempty"`
}
