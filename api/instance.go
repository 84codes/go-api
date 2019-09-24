package api

import (
	"fmt"
	"strconv"
	"time"
)

func (api *API) waitUntilReady(id string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	for {
		resp, err := api.sling.Path("/api/instances/").Get(id).ReceiveSuccess(&data)
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("Got statuscode %d from api ", resp.StatusCode)
		}
		if err != nil {
			return nil, err
		}
		if data["ready"] == true {
			data["id"] = id
			return data, nil
		}
		time.Sleep(10 * time.Second)
	}
}

func (api *API) CreateInstance(params map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	resp, err := api.sling.Post("/api/instances").BodyJSON(params).ReceiveSuccess(&data)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got statuscode %d from api ", resp.StatusCode)
	}
	if err != nil {
		return nil, err
	}
	string_id := strconv.Itoa(int(data["id"].(float64)))
	return api.waitUntilReady(string_id)
}

func (api *API) ReadInstance(id string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	resp, err := api.sling.Path("/api/instances/").Get(id).ReceiveSuccess(&data)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got statuscode %d from api ", resp.StatusCode)
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (api *API) UpdateInstance(id string, params map[string]interface{}) error {
	resp, err := api.sling.Put("/api/instances/" + id).BodyJSON(params).ReceiveSuccess(nil)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Got statuscode %d from api ", resp.StatusCode)
	}
	return err
}

func (api *API) DeleteInstance(id string) error {
	resp, err := api.sling.Path("/api/instances/").Delete(id).ReceiveSuccess(nil)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Got statuscode %d from api ", resp.StatusCode)
	}
	return err
}
