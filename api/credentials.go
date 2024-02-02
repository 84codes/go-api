package api

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func (api *API) ReadCredentials(id int) (map[string]interface{}, error) {
	var (
		data       map[string]interface{}
		failed     map[string]interface{}
		instanceID = strconv.Itoa(id)
		path       = fmt.Sprintf("/api/instances/%s", instanceID)
	)

	log.Printf("[DEBUG] go-api::credentials::read path: %s", path)
	response, err := api.sling.New().Path(path).Receive(&data, &failed)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("read credentials failed, status: %d, message: %s",
			response.StatusCode, failed)
	}

	return extractInfo(data["url"].(string)), nil
}

func extractInfo(url string) map[string]interface{} {
	paramsMap := make(map[string]interface{})
	r := regexp.MustCompile(`^.*:\/\/(?P<username>(.*)):(?P<password>(.*))@`)
	match := r.FindStringSubmatch(url)

	for i, name := range r.SubexpNames() {
		if name == "username" {
			paramsMap["username"] = match[i]
		}
		if name == "password" {
			paramsMap["password"] = match[i]
		}
	}

	return paramsMap
}
