package pokemon

type Sprites struct {
	ImageUrl string `json:"front_shiny"`
}
type Type struct {
	Name string `json:"name"`
}
type Types struct {
	Type Type `json:"type"`
}
type Stat struct {
	Name string `json:"name"`
}

type Stats struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
}

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Pokemon struct {
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	Abilities      []Abilities `json:"abilities"`
	Stats          []Stats     `json:"stats"`
	Id             int         `json:"id"`
	Types          []Types     `json:"types"`
	Sprites        Sprites     `json:"sprites"`
}
