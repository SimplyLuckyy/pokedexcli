package pokeapi

import (
	"net/http"

	"github.com/simplyluckyy/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache	   pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
}