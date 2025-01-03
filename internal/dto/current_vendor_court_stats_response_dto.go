package dto

import (
	"main/core/enums"
	"main/core/types"
)

// CurrentVendorCourtStatsResponseDTO is a struct that defines the current vendor court stats response DTO.
type CurrentVendorCourtStatsResponseDTO struct {
	// TotalCourtCount is the total court count.
	TotalCourtCount int64 `json:"total_court_count"`

	// FootballCourtCount is the football court count.
	FootballCourtCount int64 `json:"football_court_count"`

	// BasketballCourtCount is the basketball court count.
	BasketballCourtCount int64 `json:"basketball_court_count"`

	// TennisCourtCount is the tennis court count.
	TennisCourtCount int64 `json:"tennis_court_count"`

	// VolleyballCourtCount is the volleyball court count.
	VolleyballCourtCount int64 `json:"volleyball_court_count"`

	// BadmintonCourtCount is the badminton court count.
	BadmintonCourtCount int64 `json:"badminton_court_count"`
}

// FromMap is a function that converts a court counts map to a current vendor court stats response DTO.
//
// m: The court counts map.
//
// Returns the current vendor court stats response DTO.
func (c CurrentVendorCourtStatsResponseDTO) FromMap(m *types.CourtCountsMap) *CurrentVendorCourtStatsResponseDTO {
	// Calculate the total court count
	total := (*m)[enums.Football.Label()] + (*m)[enums.Basketball.Label()] + (*m)[enums.Tennis.Label()] + (*m)[enums.Volleyball.Label()] + (*m)[enums.Badminton.Label()]

	return &CurrentVendorCourtStatsResponseDTO{
		TotalCourtCount:      total,
		FootballCourtCount:   (*m)[enums.Football.Label()],
		BasketballCourtCount: (*m)[enums.Basketball.Label()],
		TennisCourtCount:     (*m)[enums.Tennis.Label()],
		VolleyballCourtCount: (*m)[enums.Volleyball.Label()],
		BadmintonCourtCount:  (*m)[enums.Badminton.Label()],
	}
}
