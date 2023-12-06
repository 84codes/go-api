package api

// VPC peering for GCP, using vpcID as identifier.

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func (api *API) waitForGcpPeeringStatusWithVpcId(vpcID, peerID string) error {
	for {
		time.Sleep(10 * time.Second)
		data, err := api.ReadVpcGcpPeeringWithVpcId(vpcID, peerID)
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

func (api *API) RequestVpcGcpPeeringWithVpcId(vpcID string, params map[string]interface{},
	waitOnStatus bool, sleep, timeout int) (map[string]interface{}, error) {

	path := fmt.Sprintf("api/vpcs/%s/vpc-peering", vpcID)
	data, err := api.requestVpcGcpPeeringWithRetry(path, params, waitOnStatus, 1, sleep, timeout)
	if err != nil {
		return nil, err
	}

	if waitOnStatus {
		log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::request waiting for active state")
		api.waitForGcpPeeringStatusWithVpcId(vpcID, data["peering"].(string))
	}

	return data, nil
}

// func (api *API) RequestVpcGcpPeeringWithVpcId(vpcID string, params map[string]interface{},
// 	waitOnStatus bool) (map[string]interface{}, error) {
// 	var (
// 		data   map[string]interface{}
// 		failed map[string]interface{}
// 		path   = fmt.Sprintf("api/vpcs/%s/vpc-peering", vpcID)
// 	)

// 	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::request params: %v", params)
// 	response, err := api.sling.New().Post(path).BodyJSON(params).Receive(&data, &failed)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if response.StatusCode != 200 {
// 		return nil, fmt.Errorf("request VPC peering failed, status: %v, message: %s", response.StatusCode, failed)
// 	}

// 	if waitOnStatus {
// 		log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::request waiting for active state")
// 		api.waitForGcpPeeringStatusWithVpcId(vpcID, data["peering"].(string))
// 	}
// 	return data, nil
// }

// func (api *API) ReadVpcGcpPeeringWithVpcId(vpcID, peerID string) (map[string]interface{}, error) {
// 	var (
// 		data   map[string]interface{}
// 		failed map[string]interface{}
// 		path   = fmt.Sprintf("/api/vpcs/%s/vpc-peering", vpcID)
// 	)

// 	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::read instance_id: %s, peer_id: %s", vpcID, peerID)
// 	response, err := api.sling.New().Get(path).Receive(&data, &failed)
// 	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::read data: %v", data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if response.StatusCode != 200 {
// 		return nil, fmt.Errorf("ReadRequest failed, status: %v, message: %s", response.StatusCode, failed)
// 	}

// 	return data, nil
// }

func (api *API) UpdateVpcGcpPeeringWithVpcId(vpcID, peerID string) (map[string]interface{}, error) {
	return api.ReadVpcGcpPeeringWithVpcId(vpcID, peerID)
}

func (api *API) RemoveVpcGcpPeeringWithVpcId(vpcID, peerID string) error {
	var (
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/vpcs/%s/vpc-peering/%s", vpcID, peerID)
	)

	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::remove vpc id: %s, peering id: %s", vpcID, peerID)
	response, err := api.sling.New().Delete(path).Receive(nil, &failed)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return fmt.Errorf("RemoveVpcPeering failed, status: %v, message: %s", response.StatusCode, failed)
	}
	return nil
}

func (api *API) ReadVpcGcpInfoWithVpcId(vpcID string, sleep, timeout int) (map[string]interface{}, error) {
	// Initiale values, 5 attempts and 20 second sleep
	return api.readVpcGcpInfoWithRetryWithVpcId(vpcID, 5, 20)
}

func (api *API) readVpcGcpInfoWithRetryWithVpcId(vpcID string, attempts, sleep int) (map[string]interface{}, error) {
	var (
		data   map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/vpcs/%s/vpc-peering/info", vpcID)
	)

	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::info vpc id: %s", vpcID)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::info data: %v", data)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode
	log.Printf("[DEBUG] go-api::vpc_gcp_peering_withvpcid::info statusCode: %d", statusCode)
	switch {
	case statusCode == 400:
		// Todo: Add error code to avoid using string comparison
		if strings.Compare(failed["error"].(string), "Timeout talking to backend") == 0 {
			if attempts--; attempts > 0 {
				log.Printf("[INFO] go-api::vpc_gcp_peering_withvpcid::info Timeout talking to backend "+
					"attempts left %d and retry in %d seconds", attempts, sleep)
				time.Sleep(time.Duration(sleep) * time.Second)
				return api.readVpcGcpInfoWithRetryWithVpcId(vpcID, attempts, 2*sleep)
			} else {
				return nil, fmt.Errorf("ReadInfo failed, status: %v, message: %s", response.StatusCode, failed)
			}
		}
	}
	return data, nil
}
