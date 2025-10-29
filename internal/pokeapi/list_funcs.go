package pokeapi


import (
	"encoding/json"
	"io"
	"net/http"
	//"fmt"
)

const (baseURL = "https://pokeapi.co/api/v2")

type PokeAPILoca struct {
	Count int 	 	 `json:"count"`
	Next  *string 	 `json:"next"`
	Prev  *string 	 `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeAPILocaINFO struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLoca(pageURL *string) (PokeAPILoca, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		loca := PokeAPILoca{}
		err := json.Unmarshal(val, &loca)
		if err != nil {return PokeAPILoca{}, err}

		return loca, nil
	} 

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {return PokeAPILoca{}, err}

	resp, err := c.httpClient.Do(req)
	if err != nil {return PokeAPILoca{}, err}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {return PokeAPILoca{}, err}

	loca := PokeAPILoca{}
	err = json.Unmarshal(data, &loca)
	if err != nil {return PokeAPILoca{}, err}
	
	c.cache.Add(url, data)
	return loca, nil		

}

func (c *Client) ListLocaINFO(location string) (PokeAPILocaINFO, error) {
	url := baseURL + "/location-area/" + location

	val, ok := c.cache.Get(url)
	if ok {
		locaINFO := PokeAPILocaINFO{}
		err := json.Unmarshal(val, &locaINFO)
		if err != nil {return PokeAPILocaINFO{}, err}

		return locaINFO, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {return PokeAPILocaINFO{}, err}

	resp, err := c.httpClient.Do(req)
	if err != nil {return PokeAPILocaINFO{}, err}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {return PokeAPILocaINFO{}, err}

	locaINFO := PokeAPILocaINFO{}
	err = json.Unmarshal(data, &locaINFO)
	if err != nil {return PokeAPILocaINFO{}, err}
	
	c.cache.Add(url, data)
	return locaINFO, nil
}