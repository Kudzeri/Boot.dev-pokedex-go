package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClint http.Client
}

func NewClient() Client {
	return Client{
		httpClint: http.Client{
			Timeout: time.Minute,
		},
	}
}
