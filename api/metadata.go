package api

import (
	"fmt"
)

type Plan struct {
	Name    string `json:"name"`
	Backend string `json:"backend"`
	Shared  bool   `json:"shared"`
}

type Region struct {
	Provider string `json:"provider"`
	Region   string `json:"region"`
}

func (api *API) ValidatePlan(name string) error {
	var (
		data   []Plan
		failed map[string]interface{}
		path   = fmt.Sprintf("api/plans")
	)

	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("%s", failed["message"].(string))
	}

	for _, plan := range data {
		if name == plan.Name {
			return nil
		}
	}
	return fmt.Errorf("Subscription plan: %s is not valid", name)
}

func (api *API) ValidateRegion(region string) error {
	var (
		data     []Region
		failed   map[string]interface{}
		path     = fmt.Sprintf("api/regions")
		platform string
	)

	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("%s", failed["message"].(string))
	}

	for _, v := range data {
		platform = fmt.Sprintf("%s::%s", v.Provider, v.Region)
		if region == platform {
			return nil
		}
	}

	return fmt.Errorf("Provider & region: %s is not valid", region)
}
