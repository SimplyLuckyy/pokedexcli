package pokeapi


import (
	"encoding/json"
	"io"
	"net/http"
)

const (baseURL = "https://pokeapi.co/api/v2")

type RespShallowLoca struct {
	Count int 	 	 `json:"count"`
	Next  *string 	 `json:"next"`
	Prev  *string 	 `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLoca(pageURL *string) (RespShallowLoca, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {return RespShallowLoca{}, err}

	resp, err := c.httpClient.Do(req)
	if err != nil {return RespShallowLoca{}, err}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {return RespShallowLoca{}, err}

	locaResp := RespShallowLoca{}
	err = json.Unmarshal(data, &locaResp)
	if err != nil {return RespShallowLoca{}, err}

	return locaResp, nil
}