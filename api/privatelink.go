package api

import (
	"fmt"
	"log"
	"time"
)

func (api *API) EnablePrivatelink(instanceID int) error {
	// TODO: Check if instance have VPC already?
	// TODO: Otherwise just need to first enable it before moving on.
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/privatelink", instanceID)
	)
	response, err := api.sling.New().Post(path).Receive(nil, &failed)
	if err != nil {
		return err
	} else if response.StatusCode == 200 {
		return api.waitForEnablePrivatelinkWithRetry(instanceID, 5, 20)
	}
	return fmt.Errorf("Enable PrivateLink failed, status: %v, message: %s", response.StatusCode, failed)
}

func (api *API) ReadPrivatelink(instanceID int) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/privatelink", instanceID)
	)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	if err != nil {
		return nil, err
	} else if response.StatusCode == 200 {
		return data, nil
	}
	return nil, fmt.Errorf("Read PrivateLink failed, status: %v, message: %s", response.StatusCode, failed)
}

func (api *API) UpdatePrivatelink(instanceID int, params map[string]interface{}) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/privatelink", instanceID)
	)
	response, err := api.sling.New().Put(path).BodyJSON(params).Receive(nil, &failed)
	if err != nil {
		return err
	} else if response.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("Update Privatelink failed, status: %v, message: %s", response.StatusCode, failed)
}

func (api *API) DisablePrivatelink(instanceID int) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/privatelink", instanceID)
	)
	response, err := api.sling.New().Delete(path).Receive(nil, &failed)
	if err != nil {
		return err
	} else if response.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("Delete Privatelink failed, status: %v, message: %s", response.StatusCode, failed)
}

func (api *API) waitForEnablePrivatelinkWithRetry(instanceID, attempts, sleep int) error {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/privatelink", instanceID)
	)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 200:
		switch data["status"].(string) {
		case "enabled":
			return nil
		case "pending":
			if attempts--; attempts > 0 {
				log.Printf("[INFO] go-api::privatelink::waitForEnablePrivatelink "+
					"attempts left %d and retry in %d seconds", attempts, sleep)
				time.Sleep(time.Duration(sleep) * time.Second)
				return api.waitForEnablePrivatelinkWithRetry(instanceID, attempts, 2*sleep)
			}
		}
	}
	return fmt.Errorf("Wait for enable PrivateLink failed, status: %v, message: %s",
		response.StatusCode, failed)
}
