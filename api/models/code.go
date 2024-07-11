package models

import (
	"github.com/google/uuid"
)

type PointsOfInterest struct {
	Poi_Id      uuid.UUID `json:"poi_id"`
	Tenant_Id   string    `json:"tenant_id"`
	Address        string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Description string    `json:"description"`
	Archived    bool      `json:"archived"`
}
