package main

import "github.com/84codes/go-api/api"

func NewApi(baseUrl, apiKey string) *api.API {
	return api.New(baseUrl, apiKey)
}

func NewCustomerApi(baseUrl, apiKey string) *api.CustomerAPI {
	return api.NewCustomerApi(baseUrl, apiKey)
}

func NewAlarmApi(baseUrl, apiKey string) *api.AlarmAPI {
	return api.NewAlarmApi(baseUrl, apiKey)
}

