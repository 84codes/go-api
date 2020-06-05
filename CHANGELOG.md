## 1.3.2 (Jun, 05, 2020)

IMPREVEMENTS:

* lint: name convention
* api/plugin.go - Added delete function. Missmatch when deleting instance.

## 1.3.1 (Apr, 29, 2020)

IMPROVEMENTS:

* api/nodes.go - Added API call to retrieve node information

## 1.3.0 (Apr, 27, 2020)

IMPROVEMENTS:

* api/integration.go - Added combined log and metric integration API calls
* api/alarms.go - Updated wait function for alarm deletion

## 1.2.0 (Feb 27, 2020)

NOTES:

* Changes to underlying API

IMPROVEMENTS:

* api/notifications.go - Updated to handle changes to API
* api/alarm.go - Updated to handle changes to API

## 1.1.2 (Feb 18, 2020)

BUG FIXES:

* api/credentials.go - Url regex expression to handle both amqp and amqps

## 1.1.1 (Feb 15, 2020)

IMPROVEMENTS:

* api/notification.go - Function to retrieve all recipients

BUG FIXES:

* api/instance.go - Url regex expression to handle both amqp and amqps

## 1.1.0 (Dec 13, 2019)

NOTES:

* First release

FEATURES:

* api/instance.go - Extract information from AMQP url

IMPROVEMENTS:

* api/*.go - Added debug logging
* api/*.go - Additional validation

BUG FIXES:

* api/plugins.go - Changed body format and content-type
* api/plugins_community.go - Changed body format and content-type

## 1.0.0 (Nov 22, 2019) // Unreleased

NOTES:

* Major change, user agent to api calls and new sling objects.

FEATURES:

* api/vpc_peering.go - CRUD service for VPC and peering
* api/api.go - Add user agent to all API calls // breakande
* api/*.go - New sling object with base properties
* api/security_firewall.go - CRUD service for firewall rules

BUG FIXES:

* api/vpc_peering.go - Removed wrongly data[id]

## 0.6.0 (Oct 28, 2019) // Unreleased

FEATURES:

* api/credentials.go - CloudAMQP specific: Extract credentials from url
* api/plugins_community.go - Additional CRUD service for community plugins

BUG FIXES:

* api/notifications.go - Error handling on resource identifiers

## 0.5.0 (Oct 7, 2019) // Unreleased

IMPROVEMENTS:

* api/alarms.go - Add func call to list alarms
* api/api.go - CloudAMQP specific: Use current default Rabbit MQ version

## 0.4.0 (Sep 24, 2019) // Unreleased

FEATURES:

* LICENSE.md - Added MIT license

IMPROVEMENTS:

* api/alarms.go - Additional error handling
* api/instance.go - Additional error handling
* api/notifications.go - Additional error handling
* api/plugins.go - Additional error handling

BUG FIXES:

* api/alarms.go - Correct corresponding status code received by API calls
* api/instance.go - Correct corresponding status code received by API calls
* api/notifications.go - Correct corresponding status code received by API calls
* api/plugins.go - Correct corresponding status code received by API calls

## 0.3.0 (Sep 10, 2019) // Unreleased

NOTES:

* Refactor api.go for single API calls and use proxy

REMOVED:

* api/api.go - Deleted older resource

FEATURES:

* api/alarms.go - CRUD services for both customer and api-proxy calls
* api/instance.go - CRUD services for both customer and api-proxy calls
* api/plugins.go - CRUD services for both customer and api-proxy calls
* api/api.go - New resource for single API endpoints
* go.mod - Module dependencies
* go.sum - Module dependencies

## 0.2.0 (Jul 5, 2019) // Unreleased

NOTES:

* Refactor to handle multiple API endpoints

FEATURES:

* api/api_alarms.go - Extend API wrapper
* api/customer_instance.go - Extend API wrapper
* main.go - Extend API wrapper

IMPROVEMENTS:

* api/api_alarms.go - Multiple API calls
* main.go - Multiple API calls

## 0.1.0 (Jun 18, 2019) // Unreleased

NOTES:

* Initial commit

FEATURES:

* api.go
* main.go
* readme.md
