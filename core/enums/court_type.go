package enums

import (
	"main/pkg/utils"
	"slices"
)

// CourtType is an enum that defines the court types.
type CourtType int

// CourtType is an enum that defines the court types.
const (
	Football CourtType = iota
	Basketball
	Tennis
	Volleyball
	Badminton
)

// courtTypes is a list of court types.
var courtTypes = []string{
	"Football",
	"Basketball",
	"Tennis",
	"Volleyball",
	"Badminton",
}

// courtTypesApiValue is a list of court types API value.
var courtTypesApiValue = map[string]int{
	"Football":   1,
	"Basketball": 2,
	"Tennis":     3,
	"Volleyball": 4,
	"Badminton":  5,
}

// Label is a function that returns the label of the court type.
//
// Returns the label of the court type.
func (c CourtType) Label() string {
	return courtTypes[c]
}

// GetCourtTypeID is a function that returns the ID of the court type.
//
// val: The value of the court type.
//
// Returns the ID of the court type.
func GetCourtTypeID(val string) int {
	return courtTypesApiValue[val]
}

// InCourtType is a function that checks if the given string is a court type.
//
// s: The string to check.
//
// Returns true if the string is a court type.
func InCourtType(s string) bool {
	return slices.Contains(courtTypes, utils.UpperFirstLetter(s))
}

// CourtTypes is a function that returns the court types.
//
// Returns the court types.
func CourtTypes() []string {
	return courtTypes
}
