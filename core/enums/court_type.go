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

// Label is a function that returns the label of the court type.
//
// Returns the label of the court type.
func (c CourtType) Label() string {
	return courtTypes[c]
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
