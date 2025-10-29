package pokeapi

import (
	"net/http"
	"time"

	"github.com/simplyluckyy/pokedexcli/internal"
)

type Client struct {
	httpClient http.Client
	cache	   pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{},
		cache: pokecache.NewCache(cacheInterval),
	}
}