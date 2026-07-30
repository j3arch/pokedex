package pokeapi

import (
	"net/http"
)

type Client struct {
	http.Client http.Client
}