package api

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func (api *API) WaitForPeeringStatus(instanceID int, peeringID string, attempt, sleep, timeout int) (int, error) {
	time.Sleep(10 * time.Second)
	path := fmt.Sprintf("/api/instances/%v/vpc-peering/status/%v", instanceID, peeringID)
	return api.waitForPeeringStatusWithRetry(path, attempt, sleep, timeout)
}

func (api *API) waitForPeeringStatusWithRetry(path string, attempt, sleep, timeout int) (int, error) {
	log.Printf("[DEBUG] go-api::vpc_peering::waitForPeeringStatusWithRetry path: %s, "+
		"attempt: %d, sleep: %d, timeout: %d", path, attempt, sleep, timeout)
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	response, err := api.sling.New().Path(path).Receive(&data, &failed)
	if err != nil {
		return attempt, err
	}

	switch {
	case attempt*sleep >= timeout:
		return attempt, fmt.Errorf("Remove VPC peering failed, reached timeout of %d seconds", timeout)
	case response.StatusCode == 200:
		switch data["status"] {
		case "active", "pending-acceptance":
			return attempt, nil
		}
		// Todo: Check if needed?
		log.Printf("[DEBUG] go-api::vpc_peering::waitForPeeringStatusWithRetry No status yet, "+
			"attempt: %d, until timeout: %d", attempt, (timeout - (attempt * sleep)))
		attempt++
		time.Sleep(time.Duration(sleep) * time.Second)
		return api.waitForPeeringStatusWithRetry(path, attempt, sleep, timeout)
	case response.StatusCode == 400:
		switch {
		case failed["error_code"] == nil:
			break
		case failed["error_code"].(float64) == 40003:
			log.Printf("[DEBUG] go-api::vpc_peering::waitForPeeringStatusWithRetry Could not find VPC peering, "+
				"attempt: %d, until timeout: %d", attempt, (timeout - (attempt * sleep)))
			attempt++
			time.Sleep(time.Duration(sleep) * time.Second)
			return api.waitForPeeringStatusWithRetry(path, attempt, sleep, timeout)
		}
	}
	return attempt, fmt.Errorf("Accept VPC peering failed, status: %v, message: %v", response.StatusCode, failed)
}

func (api *API) ReadVpcInfo(instanceID int) (map[string]interface{}, error) {
	// Initiale values, 5 attempts and 20 second sleep
	return api.readVpcInfoWithRetry(instanceID, 5, 20)
}

func (api *API) readVpcInfoWithRetry(instanceID, attempts, sleep int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_peering::info instance id: %v", instanceID)
	path := fmt.Sprintf("/api/instances/%v/vpc-peering/info", instanceID)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc_peering::info data: %v", data)

	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode
	log.Printf("[DEBUG] go-api::vpc_peering::info statusCode: %d", statusCode)
	switch {
	case statusCode == 400:
		if strings.Compare(failed["error"].(string), "Timeout talking to backend") == 0 {
			if attempts--; attempts > 0 {
				log.Printf("[INFO] go-api::vpc_peering::info Timeout talking to backend "+
					"attempts left %d and retry in %d seconds", attempts, sleep)
				time.Sleep(time.Duration(sleep) * time.Second)
				return api.readVpcInfoWithRetry(instanceID, attempts, 2*sleep)
			}
			return nil, fmt.Errorf("ReadInfo failed, status: %v, message: %s", response.StatusCode, failed)
		}
	}
	return data, nil
}

func (api *API) ReadVpcPeeringRequest(instanceID int, peeringID string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_peering::request instance id: %v, peering id: %v", instanceID, peeringID)
	path := fmt.Sprintf("/api/instances/%v/vpc-peering/request/%v", instanceID, peeringID)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc_peering::request data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("ReadRequest failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return data, nil
}

func (api *API) retryAcceptVpcPeering(instanceID int, peeringID string, attempt, sleep, timeout int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	path := fmt.Sprintf("/api/instances/%v/vpc-peering/request/%v", instanceID, peeringID)
	response, err := api.sling.New().Put(path).Receive(&data, &failed)

	if err != nil {
		return nil, err
	}
	if attempt*sleep >= timeout {
		return nil, fmt.Errorf("AcceptVpcPeering failed, reached timeout of %d seconds", timeout)
	} else if response.StatusCode == 400 {
		errorCode := failed["error_code"].(float64)
		if errorCode == 40001 {
			log.Printf("[DEBUG] go-api::vpc_peering::accept firewall not finished configuring will retry "+
				"accept VPC peering, attempt: %d, until timeout: %d", attempt, (timeout - (attempt * sleep)))
			attempt++
			time.Sleep(time.Duration(sleep) * time.Second)
			return api.retryAcceptVpcPeering(instanceID, peeringID, attempt, sleep, timeout)
		}
		return nil, fmt.Errorf("AcceptVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("AcceptVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	}
	return data, nil
}

func (api *API) AcceptVpcPeering(instanceID int, peeringID string, sleep, timeout int) (map[string]interface{}, error) {
	_, err := api.waitForPeeringStatus(instanceID, peeringID)
	if err != nil {
		return nil, err
	}
	return api.retryAcceptVpcPeering(instanceID, peeringID, 1, sleep, timeout)
}

func (api *API) retryRemoveVpcPeering(instanceID int, peeringID string, attempt, sleep, timeout int) error {
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_peering::remove instance id: %v, peering id: %v", instanceID, peeringID)
	path := fmt.Sprintf("/api/instances/%v/vpc-peering/%v", instanceID, peeringID)
	response, err := api.sling.New().Delete(path).Receive(nil, &failed)

	if err != nil {
		return err
	}
	if attempt*sleep >= timeout {
		return fmt.Errorf("RemoveVpcPeering failed, reached timeout of %d seconds", timeout)
	} else if response.StatusCode == 400 {
		errorCode := failed["error_code"].(float64)
		if errorCode == 40001 {
			log.Printf("[DEBUG] go-api::vpc_peering::remove firewall not finished configuring will retry "+
				"accept VPC peering, attempt: %d, until timeout: %d", attempt, (timeout - (attempt * sleep)))
			attempt++
			time.Sleep(time.Duration(sleep) * time.Second)
			return api.retryRemoveVpcPeering(instanceID, peeringID, attempt, sleep, timeout)
		}
		return fmt.Errorf("RemoveVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	} else if response.StatusCode != 204 {
		return fmt.Errorf("RemoveVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	}
	return nil
}

func (api *API) RemoveVpcPeering(instanceID int, peeringID string, sleep, timeout int) error {
	return api.retryRemoveVpcPeering(instanceID, peeringID, 1, sleep, timeout)
}
