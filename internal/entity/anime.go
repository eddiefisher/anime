package entity

import "github.com/gdamore/tcell/v2"

type Anime struct {
	Source         string
	Title          string
	Description    string
	Type           Type
	CurrentEpisode string
	Episodes       int64
	Status         Status
	StatusColor    tcell.Color
	AnimeSeason    AnimeSeason
	Picture        string
	Thumbnail      string
	Synonyms       []string
	Relations      []string
	Tags           []string
	Date           string
}

type Status struct {
	Name  StatusName
	Color StatusColor
}

type Type int64

const (
	Movie Type = iota + 1
	ONA
	OVA
	Special
	TV
	Undefined
)

// String - Creating common behavior - give the type a String function
func (t Type) String() string {
	return [...]string{"tv", "movie", "ova", "ona", "special", "undefined"}[t-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (t Type) EnumIndex() int {
	return int(t)
}

type StatusColor string

const (
	FinishedColor StatusColor = "red"
	AnnonceColor  StatusColor = "blue"
	OngoingColor  StatusColor = "green"
)

type StatusName int64

const (
	Finished StatusName = iota + 1
	Annonce
	Ongoing
)

// String - Creating common behavior - give the type a String function
func (t StatusName) String() string {
	return [...]string{"finished", "annonce", "ongoing"}[t-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (t StatusName) EnumIndex() int {
	return int(t)
}
