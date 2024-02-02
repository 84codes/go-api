package api

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func (api *API) CreateNotification(instanceID int, params map[string]interface{}) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/recipients", instanceID)
	)

	log.Printf("[DEBUG] go-api::notification::create path: %s, params: %v", path, params)
	response, err := api.sling.New().Post(path).BodyJSON(params).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::notification::create data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, fmt.Errorf("create notification failed, status: %d, message: %s",
			response.StatusCode, failed)
	}

	if v, ok := data["id"]; ok {
		data["id"] = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		log.Printf("[DEBUG] go-api::notification::create id set: %v", data["id"])
	} else {
		msg := fmt.Sprintf("go-api::notification::create Invalid notification identifier: %v", data["id"])
		log.Printf("[ERROR] %s", msg)
		return nil, errors.New(msg)
	}

	return data, err
}

func (api *API) ReadNotification(instanceID int, recipientID string) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/recipients/%s", instanceID, recipientID)
	)

	log.Printf("[DEBUG] go-api::notification::read path: %s", path)
	response, err := api.sling.New().Path(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::notification::read data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("read notification failed, status: %d, message: %s",
			response.StatusCode, failed)
	}

	return data, err
}

func (api *API) ReadNotifications(instanceID int) ([]map[string]interface{}, error) {
	var (
		data   []map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/recipients", instanceID)
	)

	log.Printf("[DEBUG] go-api::ReadNotifications::read path: %s", path)
	response, err := api.sling.New().Path(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::ReadNotifications::read data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("read notifications failed, status: %d, message: %s",
			response.StatusCode, failed)
	}

	return data, err
}

func (api *API) UpdateNotification(instanceID int, params map[string]interface{}) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/recipients/%s", instanceID, params["id"])
	)

	log.Printf("[DEBUG] go-api::notification::update path: %s, params: %v", path, params)

	response, err := api.sling.New().Put(path).BodyJSON(params).Receive(nil, &failed)

	if response.StatusCode != 200 {
		return fmt.Errorf("update notification failed, status: %d, message: %s",
			response.StatusCode, failed)
	}

	return err
}

func (api *API) DeleteNotification(instanceID int, params map[string]interface{}) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/recipients/%s", instanceID, params["id"])
	)

	log.Printf("[DEBUG] go-api::notification::delete path: %s, params: %v", path, params)
	response, err := api.sling.New().Delete(path).BodyJSON(params).Receive(nil, &failed)

	if response.StatusCode != 204 {
		return fmt.Errorf("delete notification failed, status: %d, message: %s",
			response.StatusCode, failed)
	}

	return err
}
