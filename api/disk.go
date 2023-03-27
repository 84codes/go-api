package api

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func (api *API) ResizeDisk(instanceID int, params map[string]interface{}, sleep, timeout int) (map[string]interface{}, error) {
	var (
		id   = strconv.Itoa(instanceID)
		path = fmt.Sprintf("api/instances/%s/disk", id)
	)
	log.Printf("[DEBUG] go-api::resizeDisk::resizeDiskWithRetry path: %s, "+
		"attempt: %d, sleep: %d, timeout: %d", path, 1, sleep, timeout)
	return api.resizeDiskWithRetry(id, params, 1, sleep, timeout)
}

func (api *API) resizeDiskWithRetry(id string, params map[string]interface{}, attempt, sleep, timeout int) (map[string]interface{}, error) {
	var (
		data   = make(map[string]interface{})
		failed = make(map[string]interface{})
		path   = fmt.Sprintf("api/instances/%s/disk", id)
	)

	response, err := api.sling.New().Put(path).BodyJSON(params).Receive(&data, &failed)
	if err != nil {
		return nil, err
	} else if attempt*sleep > timeout {
		return nil, fmt.Errorf("Wait until resize disk failed, reached timeout of %d seconds", timeout)
	}

	switch response.StatusCode {
	case 200:
		if err = api.waitUntilAllNodesReady(id); err != nil {
			return nil, err
		}
		return data, nil
	case 400:
		switch {
		case failed["error_code"] == nil:
			break
		case failed["error_code"].(float64) == 40002:
			return nil, fmt.Errorf("Resize disk failed: %s", failed["error"].(string))
		case failed["error_code"].(float64) == 40003:
			if failed["error"] == "Timeout talking to backend" {
				log.Printf("[DEBUG] go-api::resizeDisk::resizeDiskWithRetry Timeout talking to backend, "+
					" will try again, attempt: %d, until timeout: %d", attempt, (timeout - (attempt * sleep)))
				attempt++
				time.Sleep(time.Duration(sleep) * time.Second)
				return api.resizeDiskWithRetry(id, params, attempt, sleep, timeout)
			} else {
				break
			}
		}
	}
	return nil, fmt.Errorf("Resize disk failed, status: %v, message: %s", response.StatusCode, failed["error"])
}
