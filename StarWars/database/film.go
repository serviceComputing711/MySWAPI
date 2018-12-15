package database

// A Film is an single film.
type Film struct {
	Title         string         `json:"title"`
	EpisodeID     int            `json:"episode_id"`
	OpeningCrawl  string         `json:"opening_crawl"`
	Director      string         `json:"director"`
	Producer      string         `json:"producer"`
	CharacterURLs []characterURL `json:"characters"`
	PlanetURLs    []planetURL    `json:"planets"`
	StarshipURLs  []starshipURL  `json:"starships"`
	VehicleURLs   []vehicleURL   `json:"vehicles"`
	SpeciesURLs   []speciesURL   `json:"species"`
	Created       string         `json:"created"`
	Edited        string         `json:"edited"`
	URL           string         `json:"url"`
}


type filmURL string
type characterURL string

func (p *Film) getName() (string){
	return p.Title
}