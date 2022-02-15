package api

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func (api *API) CreateVpcInstance(params map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc::create params: %v", params)
	response, err := api.sling.New().Post("/api/vpcs").BodyJSON(params).Receive(&data, &failed)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("CreateVpcInstance failed, status: %v, message: %s", response.StatusCode, failed)
	}

	if id, ok := data["id"]; ok {
		data["id"] = strconv.FormatFloat(id.(float64), 'f', 0, 64)
		log.Printf("[DEBUG] go-api::vpc::create id set: %v", data["id"])
	} else {
		msg := fmt.Sprintf("go-api::vpc::create Invalid instance identifier: %v", data["id"])
		log.Printf("[ERROR] %s", msg)
		return nil, errors.New(msg)
	}

	return data, nil
}

func (api *API) ReadVpcInstance(instanceID string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc::read instance ID: %s", instanceID)

	path := fmt.Sprintf("/api/vpcs/%s", instanceID)
	response, err := api.sling.New().Path(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc::read data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("ReadVpcInstance failed, status: %v, message: %v", response.StatusCode, failed)
	}

	return data, nil
}

func (api *API) UpdateVpcInstance(instanceID string, params map[string]interface{}) error {
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::instance::update instance ID: %s, params: %v", instanceID, params)
	path := fmt.Sprintf("api/vpcs/%s", instanceID)
	response, err := api.sling.New().Put(path).BodyJSON(params).Receive(nil, &failed)

	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("UpdateInstance failed, status: %v, message: %v", response.StatusCode, failed)
	}

	return api.waitUntilAllNodesReady(instanceID)
}

func (api *API) DeleteVpcInstance(instanceID string) error {
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc::delete instance ID: %s", instanceID)
	response, err := api.sling.New().Path("/api/vpcs/").Delete(instanceID).Receive(nil, &failed)

	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return fmt.Errorf("DeleteVpcInstance failed, status: %v, message: %v", response.StatusCode, failed)
	}

	return nil
}
