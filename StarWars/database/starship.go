package database

// A Starship is a single transport craft that has hyperdrive capability.
type Starship struct {
	Name                 string      `json:"name"`
	Model                string      `json:"model"`
	Manufacturer         string      `json:"manufacturer"`
	CostInCredits        string      `json:"cost_in_credits"`
	Length               string      `json:"length"`
	MaxAtmospheringSpeed string      `json:"max_atmosphering_speed"`
	Crew                 string      `json:"crew"`
	Passengers           string      `json:"passengers"`
	CargoCapacity        string      `json:"cargo_capacity"`
	Consumables          string      `json:"consumables"`
	HyperdriveRating     string      `json:"hyperdrive_rating"`
	MGLT                 string      `json:"MGLT"`
	StarshipClass        string      `json:"starship_class"`
	PilotURLs            []personURL `json:"pilots"`
	FilmURLs             []filmURL   `json:"films"`
	Created              string      `json:"created"`
	Edited               string      `json:"edited"`
	URL                  string      `json:"url"`
}

type starshipURL string

func (p *Starship) getName() (string){
	return p.Name
}