package api

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (api *API) CreateAlarm(instanceID int, params map[string]interface{}) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms", instanceID)
	)

	log.Printf("[DEBUG] go-api::alarm::create path: %s", path)
	response, err := api.sling.New().Post(path).BodyJSON(params).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::alarm::create data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, fmt.Errorf("CreateAlarm failed, status: %d, message: %s", response.StatusCode, failed)
	}

	if id, ok := data["id"]; ok {
		data["id"] = strconv.FormatFloat(id.(float64), 'f', 0, 64)
	} else {
		msg := fmt.Sprintf("go-api::instance::create Invalid alarm identifier: %v", data["id"])
		return nil, errors.New(msg)
	}

	return data, err
}

func (api *API) ReadAlarm(instanceID int, alarmID string) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/%v", instanceID, alarmID)
	)

	log.Printf("[DEBUG] go-api::alarm::read path: %s", path)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::alarm::read data : %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("read alarm failed, status: %d, message: %s", response.StatusCode, failed)
	}

	return data, err
}

func (api *API) ReadAlarms(instanceID int) ([]map[string]interface{}, error) {
	var (
		data   []map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms", instanceID)
	)

	log.Printf("[DEBUG] go-api::alarm::read path: %s", path)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::alarm::read data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("read alarms failed, status: %d, message: %s", response.StatusCode, failed)
	}

	return data, err
}

func (api *API) UpdateAlarm(instanceID int, params map[string]interface{}) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/%v", instanceID, params["id"])
	)

	log.Printf("[DEBUG] go-api::alarm::update path: %s", path)
	response, err := api.sling.New().Put(path).BodyJSON(params).Receive(nil, &failed)

	if err != nil {
		return err
	}
	if response.StatusCode != 201 {
		return fmt.Errorf("update alarm failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return err
}

func (api *API) DeleteAlarm(instanceID int, params map[string]interface{}) error {
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::alarm::delete instance id: %v, params: %v", instanceID, params)
	path := fmt.Sprintf("/api/instances/%v/alarms/%v", instanceID, params["id"])
	response, _ := api.sling.New().Delete(path).BodyJSON(params).Receive(nil, &failed)

	if response.StatusCode != 204 {
		return fmt.Errorf("Alarm::DeleteAlarm failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return api.waitUntilAlarmDeletion(instanceID, params["id"].(string))
}

func (api *API) waitUntilAlarmDeletion(instanceID int, id string) error {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/alarms/%s", instanceID, id)
	)
	log.Printf("[DEBUG] go-api::alarm::waitUntilAlarmDeletion waiting")

	for {
		response, err := api.sling.New().Path(path).Receive(&data, &failed)

		if err != nil {
			return err
		}
		if response.StatusCode == 404 {
			return nil
		}

		time.Sleep(10 * time.Second)
	}
}
