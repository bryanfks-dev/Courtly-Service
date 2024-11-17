package enums

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

// Label is a function that returns the label of the court type.
//
// Returns the label of the court type.
func (c CourtType) Label() string {
	return map[CourtType]string{
		Football:   "Football",
		Basketball: "Basketball",
		Tennis:     "Tennis",
		Volleyball: "Volleyball",
		Badminton:  "Badminton",
	}[c]
}
