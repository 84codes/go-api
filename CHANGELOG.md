## 1.1.0 (Dec 13, 2019) // Unreleased

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
Additional error handling
* api/alarms.go
* api/instance.go
* api/notifications.go
* api/plugins.go

BUG FIXES:
Correct corresponding status code received by API calls
* api/alarms.go
* api/instance.go
* api/notifications.go
* api/plugins.go

## 0.3.0 (Sep 10, 2019) // Unreleased

NOTES:
* Refactor api.go for single API calls and use proxy

REMOVED:
* api/api.go - Deleted older resource

FEATURES:
CRUD services for both customer and api-proxy calls
* api/alarms.go
* api/instance.go
* api/plugins.go
* go.mod
* go.sums
* api/api.go - New resource for single API endpoints

## 0.2.0 (Jul 5, 2019) // Unreleased

NOTES:
* Refactor to handle multiple API endpoints

FEATURES:
Extend API wrapper
* api/api_alarms.go
* api/customer_instance.go
* main.go

IMPROVEMENTS:
Multiple API calls
* api/api_alarms.go
* main.go

## 0.1.0 (Jun 18, 2019) // Unreleased

NOTES:
* Initial commit

FEATURES:
* api.go
* main.go
* readme.md
