package api

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// waitUntilPluginUninstalled wait until a community plugin been uninstalled.
func (api *API) waitUntilPluginUninstalled(instanceID int, pluginName string, sleep, timeout int) (
	map[string]interface{}, error) {

	log.Printf("[DEBUG] go-api::plugin_community::waitUntilPluginUninstalled instance id: %v, name: %v",
		instanceID, pluginName)
	for {
		response, err := api.ReadPlugin(instanceID, pluginName, sleep, timeout)

		if err != nil {
			return nil, err
		}
		if len(response) == 0 {
			return response, nil
		}
	}
}

// InstallPlugin install a community plugin on an instance.
func (api *API) InstallPluginCommunity(instanceID int, pluginName string, sleep, timeout int) (
	map[string]interface{}, error) {

	var (
		failed map[string]interface{}
		params = &PluginParams{Name: pluginName}
		path   = fmt.Sprintf("/api/instances/%d/plugins/community?async=true", instanceID)
	)

	log.Printf("[DEBUG] go-api::plugin_community::enable path: %s", path)
	response, err := api.sling.New().Post(path).BodyJSON(params).Receive(nil, &failed)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 204 {
		return nil,
			fmt.Errorf("EnablePluginCommunity failed, status: %v, message: %v", response.StatusCode, failed)
	}

	return api.waitUntilPluginChanged(instanceID, pluginName, true, sleep, timeout)
}

// ReadPlugin reads a specific community plugin from an instance.
func (api *API) ReadPluginCommunity(instanceID int, pluginName string, sleep, timeout int) (
	map[string]interface{}, error) {

	log.Printf("[DEBUG] go-api::plugin_community::read instance ID: %v, name: %v", instanceID, pluginName)
	data, err := api.ListPluginsCommunity(instanceID, sleep, timeout)

	if err != nil {
		return nil, err
	}

	for _, plugin := range data {
		if plugin["name"] == pluginName {
			log.Printf("[DEBUG] go-api::plugin_community::read found plugin: %v", pluginName)
			return plugin, nil
		}
	}

	return nil, nil
}

// ListPluginsCommunity list all community plugins for an instance.
func (api *API) ListPluginsCommunity(instanceID, sleep, timeout int) ([]map[string]interface{}, error) {
	return api.listPluginsCommunityWithRetry(instanceID, 1, sleep, timeout)
}

// listPluginsCommunityWithRetry list all community plugins for an instance,
// with retry if the backend is busy.
func (api *API) listPluginsCommunityWithRetry(instanceID, attempt, sleep, timeout int) (
	[]map[string]interface{}, error) {

	var (
		data   []map[string]interface{}
		failed map[string]interface{}
		path   = fmt.Sprintf("/api/instances/%d/plugins/community", instanceID)
	)

	log.Printf("[DEBUG] go-api::plugin_community::listPluginsCommunityWithRetry path: %s", path)
	response, err := api.sling.New().Get(path).Receive(&data, &failed)

	if err != nil {
		return nil, err
	} else if attempt*sleep > timeout {
		return nil, fmt.Errorf("read plugins reached timeout of %d seconds", timeout)
	}

	statusCode := response.StatusCode
	log.Printf("[DEBUG] go-api::plugin_community::listPluginsCommunityWithRetry statusCode: %d", statusCode)
	switch {
	case statusCode == 400:
		if strings.Compare(failed["error"].(string), "Timeout talking to backend") == 0 {
			log.Printf("[INFO] go-api::plugins-community::read Timeout talking to backend "+
				"attempt: %d, until timeout: %d", attempt, (timeout - (attempt * sleep)))
			attempt++
			time.Sleep(time.Duration(sleep) * time.Second)
			return api.listPluginsCommunityWithRetry(instanceID, attempt, sleep, timeout)
		}
	}
	return data, nil
}

// UpdatePluginCommunity updates a community plugin from an instance.
func (api *API) UpdatePluginCommunity(instanceID int, params map[string]interface{}, sleep, timeout int) (
	map[string]interface{}, error) {

	var (
		failed       map[string]interface{}
		pluginName   = params["name"].(string)
		enabled      = params["enabled"].(bool)
		pluginParams = &PluginParams{Name: pluginName, Enabled: enabled}
		path         = fmt.Sprintf("/api/instances/%d/plugins/community?async=true", instanceID)
	)

	log.Printf("[DEBUG] go-api::plugin_community::update path: %s", path)
	response, err := api.sling.New().Put(path).BodyJSON(pluginParams).Receive(nil, &failed)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 204 {
		return nil,
			fmt.Errorf("UpdatePluginCommunity failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return api.waitUntilPluginChanged(instanceID, pluginName, enabled, sleep, timeout)
}

// UninstallPluginCommunity uninstall a community plugin from an instance.
func (api *API) UninstallPluginCommunity(instanceID int, pluginName string, sleep, timeout int) (
	map[string]interface{}, error) {

	var (
		failed = make(map[string]interface{})
		path   = fmt.Sprintf("/api/instances/%d/plugins/community/%s?async=true", instanceID, pluginName)
	)

	log.Printf("[DEBUG] go-api::plugin_community::disable path: %s", path)
	response, err := api.sling.New().Delete(path).Receive(nil, &failed)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 204 {
		return nil,
			fmt.Errorf("DisablePluginCommunity failed, status: %v, message: %s", response.StatusCode, failed)
	}

	return api.waitUntilPluginUninstalled(instanceID, pluginName, sleep, timeout)
}
