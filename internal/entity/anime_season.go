package entity

type AnimeSeason struct {
	Season string
	Year   int64
}

type Season int64

const (
	Summer Season = iota + 1
	Autumn
	Winter
	Spring
)

// String - Creating common behavior - give the type a String function
func (t Season) String() string {
	return [...]string{"summer", "autumn", "winter", "spring"}[t-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (t Season) EnumIndex() int {
	return int(t)
}
