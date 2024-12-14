package types

// CourtCountsMap is a type that represents the court counts map.
// CourtCountsMap should be formatted as following.
// {
//     "Football": ...,
//     "Basketball": ...,
//     "Volleyball": ...,
//     "Tennis": ...,
//     "Badminton": ...
// }
type CourtCountsMap map[string]int64
