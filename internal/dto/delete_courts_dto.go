package dto

// DeleteCourtsDTO is a data transfer object that represents the data to delete courts.
type DeleteCourtsDTO struct {
	// CourtIDs is the list of court IDs to be deleted.
	CourtIDs []uint `json:"court_ids"`
}
