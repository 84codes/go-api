package main

import "github.com/84codes/go-api/api"

func NewApi(baseUrl, apiKey string) *api.API {
	return api.New(baseUrl, apiKey)
}
