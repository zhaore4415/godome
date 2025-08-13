package dbset

import (
	"myproject/internal/model/entity"
)

type User struct {
	entity.User
	Address Address
}

type Address struct {
	CountryID  int64
	ProvinceID int64
	CityID     int64
	CountyID   int64
	StreetID   int64
}
