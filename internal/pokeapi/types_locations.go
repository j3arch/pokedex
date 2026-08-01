package pokeapi

type RespShallowLocations struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []result `json:"results"`
}
