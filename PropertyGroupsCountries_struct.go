package main

type PropertyGroupsCountries struct {
	CountryCode string `db:"country_code" json:"country_code"`
	CountryName string `db:"country_name" json:"country_name"`
}
