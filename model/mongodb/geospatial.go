package model

type (
	NearBy struct {
		LocationName string     `json:"location_name" validate:"required"`
		Type         string     `json:"type" validate:"required"`
		CoorName     string     `json:"coor_name" validate:"required"`
		Coor         Coordinate `json:"coordinate" validate:"required"`
	}
	Coordinate struct {
		Long float64 `json:"long" validate:"required"`
		Lat  float64 `json:"lat" validate:"required"`
	}
)
