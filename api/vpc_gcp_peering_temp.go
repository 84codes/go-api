package api

// Temporary functions between versions in order to change managed VPC peering.
// Instead of using instanceID as identifier, use managed vpcID as identifier.

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func (api *API) waitForGcpPeeringStatusTemp(vpcID, peerID string) error {
	for {
		data, err := api.ReadVpcGcpPeeringTemp(vpcID, peerID)
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
			time.Sleep(10 * time.Second)
		}
	}
}

func (api *API) RequestVpcGcpPeeringTemp(vpcID string, params map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::request params: %v", params)
	path := fmt.Sprintf("api/vpcs/%s/vpc-peering", vpcID)
	response, err := api.sling.New().Post(path).BodyJSON(params).Receive(&data, &failed)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Request VPC peering failed, status: %v, message: %s", response.StatusCode, failed)
	}

	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::request waiting for active state")
	api.waitForGcpPeeringStatusTemp(vpcID, data["peering"].(string))
	return data, nil
}

func (api *API) ReadVpcGcpPeeringTemp(vpcID, peerID string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::read instance_id: %s, peer_id: %s", vpcID, peerID)
	path := fmt.Sprintf("/api/vpcs/%s/vpc-peering", vpcID)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::read data: %v", data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("ReadRequest failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return data, nil
}

func (api *API) UpdateVpcGcpPeeringTemp(vpcID, peerID string) (map[string]interface{}, error) {
	return api.ReadVpcGcpPeeringTemp(vpcID, peerID)
}

func (api *API) RemoveVpcGcpPeeringTemp(vpcID, peerID string) error {
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::remove vpc id: %s, peering id: %s", vpcID, peerID)
	path := fmt.Sprintf("/api/vpcs/%s/vpc-peering/%s", vpcID, peerID)
	response, err := api.sling.New().Delete(path).Receive(nil, &failed)

	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return fmt.Errorf("RemoveVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	}
	return nil
}

func (api *API) ReadVpcGcpInfoTemp(vpcID string) (map[string]interface{}, error) {
	// Initiale values, 5 attempts and 20 second sleep
	return api.readVpcGcpInfoWithRetryTemp(vpcID, 5, 20)
}

func (api *API) readVpcGcpInfoWithRetryTemp(vpcID string, attempts, sleep int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::info vpc id: %s", vpcID)
	path := fmt.Sprintf("/api/vpcs/%s/vpc-peering/info", vpcID)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::info data: %v", data)

	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_temp::info statusCode: %d", statusCode)
	switch {
	case statusCode == 400:
		// Todo: Add error code to avoid using string comparison
		if strings.Compare(failed["error"].(string), "Timeout talking to backend") == 0 {
			if attempts--; attempts > 0 {
				log.Printf("[INFO] go-api::vpc_gcp_peering_temp::info Timeout talking to backend "+
					"attempts left %d and retry in %d seconds", attempts, sleep)
				time.Sleep(time.Duration(sleep) * time.Second)
				return api.readVpcGcpInfoWithRetryTemp(vpcID, attempts, 2*sleep)
			} else {
				return nil, fmt.Errorf("ReadInfo failed, status: %v, message: %s", response.StatusCode, failed)
			}
		}
	}
	return data, nil
}
