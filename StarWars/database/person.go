package database

// A Person is an individual person or character within the Star Wars universe.
type Person struct {
	Name         string        `json:"name"`
	Height       string        `json:"height"`
	Mass         string        `json:"mass"`
	HairColor    string        `json:"hair_color"`
	SkinColor    string        `json:"skin_color"`
	EyeColor     string        `json:"eye_color"`
	BirthYear    string        `json:"birth_year"`
	Gender       string        `json:"gender"`
	Homeworld    string        `json:"homeworld"`
	FilmURLs     []filmURL     `json:"films"`
	SpeciesURLs  []speciesURL  `json:"species"`
	VehicleURLs  []vehicleURL  `json:"vehicles"`
	StarshipURLs []starshipURL `json:"starships"`
	Created      string        `json:"created"`
	Edited       string        `json:"edited"`
	URL          string        `json:"url"`
}

type personURL string

func (p *Person) getName() (string){
	return p.Name
}