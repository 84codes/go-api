## 1.10.2 (Dec 14, 2022)

IMPROVEMENTS:

* api/security_firewall.go - Add configurable sleep and timeout when configuring firewall settings [#30](https://github.com/84codes/go-api/pull/30)

## 1.10.1 (Dec 07, 2022)

IMPROVEMENTS:

* api/rabbitmq_configuration.gp - Update RabbitMQ configuration response handling when backend is busy

## 1.10.0 (Oct 24, 2022)

FEATURES:

* api/privatelink.go - Add new endpoints to support PrivateLink [#28](https://github.com/84codes/go-api/pull/28)

## 1.9.2 (Oct 07, 2022)

IMPROVEMENTS:

* api/vpc_peering.go - Retry VPC peering and wait for status [#27](https://github.com/84codes/go-api/pull/27)
* api/vpc_peering_withvpcid.go - Retry VPC peering and wait for status [#27](https://github.com/84codes/go-api/pull/27)

## 1.9.1 (Sep 14, 2022)

NOTES:

* Asynchronous request for plugin/community actions.

IMPROVEMENTS:

* api/plugin.go - Update plugin action endpoints to be called asynchronous [#24](https://github.com/84codes/go-api/pull/24)
* api/plugins_community.go - Update plugin action endpoints to be called asynchronous [#24](https://github.com/84codes/go-api/pull/24)
* api/security_firewall.go - Update error handling [#25](https://github.com/84codes/go-api/pull/25)
* api/disk.go - Update error handling [#26](https://github.com/84codes/go-api/pull/26)

## 1.9.0 (Jul 01, 2022)

FEATURES:

* api/disk.go - Add new endpoint for resize disk [#22](https://github.com/84codes/go-api/pull/22)

## 1.8.1 (May 31, 2022)

IMPROVEMENTS:

* api/nodes.go - Use the existing instance actions API [#21](https://github.com/84codes/go-api/pull/21)
* api/rabbitmq_configuration.go - Rename file and functions

## 1.8.0 (May 24, 2022)

FEATURES:

* api/vpc_peering.go - Add support for retries and timeout to accept/remove VPC peering [#20](https://github.com/84codes/go-api/pull/20)
* api/vpc_peering_withvpcid.go - Add support for retries and timeout to accept/remove VPC peering [#20](https://github.com/84codes/go-api/pull/20)

## 1.7.0 (May 13, 2022)

FEATURES:

* api/upgrade_rabbitmq.go - Added new endpoints to retrieve and upgrade RabbitMQ and Erlang versions [#19](https://github.com/84codes/go-api/pull/19)
* api/nodes.go - Added invoke action for the node (e.g. start/stop/restart RabbitMQ)

## 1.6.0 (May 04, 2022)

FEATURES:

* api/vpc.go - Added new endpoints for standalone VPC request [#18](https://github.com/84codes/go-api/pull/18)
* api/vpc_gcp_peering_withvpcid.go - Added new endpoints for GCP peering using VPC identifier [#18](https://github.com/84codes/go-api/pull/18)
* api/vpc_peering_withvpcid.go - Added new endpoints for AWS peering using VPC identifier [#18](https://github.com/84codes/go-api/pull/18)
* api/rabbit_configuration.go - Added new endpoints for RabbitMQ configuration

IMPROVEMENTS:

* api/instances.go - Added keep_vpc paramater to keep vpc when deleting instance
* all peerings - Changed sleep time to start waiting before requesting peering status

## 1.5.4 (Dec 21, 2021)

* api/security_firewall.go - Added multiple retry functionality for create and update firewall rules

## 1.5.3 (Dec 20, 2021)

* api/vpc_gcp_peering.go - Added VPC peering and info for Google Cloud Platform. [#17](https://github.com/84codes/go-api/pull/17)

## 1.5.2 (Nov 17, 2021)

* api/security_firewall.go - Include `STREAM`, `STREAM_SSL` services in default rules. [#9](https://github.com/84codes/go-api/pull/16)

## 1.5.1 (Oct 05, 2021)

NOTES:

* Use Go 1.17

## 1.5.0 (Oct 05, 2021)

FEATURES:

* api/account.go - Added account support, to be able to list all instances for an account
* api/custom_domain.go - Added custom domain support

IMPROVEMENTS:

* api/plugins_community.go - Added multiple retry functionality when fetching community plugin information
* api/instance.go - Updated iteration loop to determine if all nodes are configured
* api/instance.go - Updated endpoint to calculate number of nodes

## 1.4.0 (Sep 20, 2021)

IMPROVEMENTS:

* api/security_firewall.go - Include `HTTPS` service in default rules. [#9](https://github.com/84codes/go-api/pull/9)
* api/vpc_peering.go - Added multiple retry functionallity when fetching vpc information
* api/plugins.go - Added multiple retry functionallity when fetching plugin information
* Merged pull request [#13](https://github.com/84codes/go-api/pull/13) to fix spelling

## 1.3.4 (Nov 12, 2020)

NOTES:

* Webhook added already Oct 6, but no release until Nov 12.

FEATURES:

* api/webhook.go - Added webhook implementation against CloudAMQP API.

IMPROVEMENTS:

* api/instance.go - Wait on nodes to fully be configured and running when doing up/downgrades and updating instance.

## 1.3.3 (Jun 30, 2020)

IMPROVEMENTS:

* api/security_firewall.go - Added wait func for create, update and delete of rules. In order for cache delay to finish.

## 1.3.2 (Jun, 05, 2020)

IMPROVEMENTS:

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
