package pokeapi


import (
	"encoding/json"
	"io"
	"net/http"
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

func (c *Client) ListLoca(pageURL *string) (PokeAPILoca, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		loca := PokeAPILoca{}
		err := json.Unmarshal(val, &loca)
		if err != nil {
			return PokeAPILoca{}, err
		}

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
	err = json.Unmarshal(data, &locaResp)
	if err != nil {return PokeAPILoca{}, err}

	c.cache.Add(url, data)
	return loca, nil
}