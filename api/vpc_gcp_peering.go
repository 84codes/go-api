package api

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Include in retry logic
func (api *API) waitForGcpPeeringStatus(instanceID int, peerID string) error {
	for {
		time.Sleep(10 * time.Second)
		data, err := api.ReadVpcGcpPeering(instanceID, peerID)
		if err != nil {
			return err
		}
		rows := data["rows"].([]interface{})
		if len(rows) > 0 {
			for _, row := range rows {
				tempRow := row.(map[string]interface{})
				if tempRow["name"] != peerID {
					continue
				}
				if tempRow["state"] == "ACTIVE" {
					return nil
				}
			}
		}
	}
}

func (api *API) RequestVpcGcpPeering(instanceID int, params map[string]interface{},
	waitOnStatus bool, sleep, timeout int) (map[string]interface{}, error) {

	path := fmt.Sprintf("api/instances/%v/vpc-peering", instanceID)
	data, err := api.requestVpcGcpPeeringWithRetry(path, params, waitOnStatus, 1, sleep, timeout)
	if err != nil {
		return nil, err
	}

	if waitOnStatus {
		log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::request waiting for active state")
		api.waitForGcpPeeringStatus(instanceID, data["peering"].(string))
	}

	return data, nil
}

func (api *API) requestVpcGcpPeeringWithRetry(path string, params map[string]interface{},
	waitOnStatus bool, attempt, sleep, timeout int) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
	)

	log.Printf("[DEBUG] go-api::vpc_gcp_peering::request path: %s, params: %v", path, params)
	response, err := api.sling.New().Post(path).BodyJSON(params).Receive(&data, &failed)
	if err != nil {
		return nil, err
	} else if attempt*sleep > timeout {
		return nil, fmt.Errorf("request VPC peering failed, reached timeout of %d seconds", timeout)
	}

	switch response.StatusCode {
	case 200:
		return data, nil
	case 400:
		if strings.Compare(failed["error"].(string), "Timeout talking to backend") == 0 {
			log.Printf("[INFO] go-api::vpc_gcp_peering::request Timeout talking to backend "+
				"attempt %d until timeout: %d", attempt, (timeout - (attempt * sleep)))
			attempt++
			time.Sleep(time.Duration(sleep) * time.Second)
			return api.requestVpcGcpPeeringWithRetry(path, params, waitOnStatus, attempt, sleep, timeout)
		}
	}
	return nil, fmt.Errorf("request VPC peering failed, status: %v, message: %s",
		response.StatusCode, failed)
}

// TODO: Add retry logic
func (api *API) ReadVpcGcpPeering(instanceID int, peerID string) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%v/vpc-peering", instanceID)
	)

	log.Printf("[DEBUG] go-api::vpc_gcp_peering::read instance_id: %v, peer_id: %v", instanceID, peerID)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc_gcp_peering::read data: %v", data)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("ReadRequest failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return data, nil
}

func (api *API) UpdateVpcGcpPeering(instanceID int, peerID string) (map[string]interface{}, error) {
	return api.ReadVpcGcpPeering(instanceID, peerID)
}

// TODO: Add retry logic
func (api *API) RemoveVpcGcpPeering(instanceID int, peerID string) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%v/vpc-peering/%v", instanceID, peerID)
	)

	log.Printf("[DEBUG] go-api::vpc_gcp_peering::remove instance id: %v, peering id: %v", instanceID, peerID)
	response, err := api.sling.New().Delete(path).Receive(nil, &failed)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return fmt.Errorf("RemoveVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	}
	return nil
}

func (api *API) ReadVpcGcpInfo(instanceID, sleep, timeout int) (map[string]interface{}, error) {
	return api.readVpcGcpInfoWithRetry(instanceID, 1, sleep, timeout)
}

func (api *API) readVpcGcpInfoWithRetry(instanceID, attempt, sleep, timeout int) (map[string]interface{},
	error) {

	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%v/vpc-peering/info", instanceID)
	)

	log.Printf("[DEBUG] go-api::vpc_gcp_peering::info path: %s", path)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	if err != nil {
		return nil, err
	} else if attempt*sleep > timeout {
		return nil, fmt.Errorf("read VPC info, reached timeout of %d seconds", timeout)
	}

	switch response.StatusCode {
	case 200:
		return data, nil
	case 400:
		if strings.Compare(failed["error"].(string), "Timeout talking to backend") == 0 {
			log.Printf("[INFO] go-api::vpc_gcp_peering::info Timeout talking to backend "+
				"attempt %d until timeout: %d", attempt, (timeout - (attempt * sleep)))
			attempt++
			time.Sleep(time.Duration(sleep) * time.Second)
			return api.readVpcGcpInfoWithRetry(instanceID, attempt, sleep, timeout)
		}
	}
	return nil, fmt.Errorf("read VPC info failed, status: %v, message: %s",
		response.StatusCode, failed)
}
