package main

import (
	"github.com/84codes/go-api/api"
)

func main() {

}

func New(baseUrl, apiKey string) *api.API {
	return api.New(baseUrl, apiKey, "test")
}
