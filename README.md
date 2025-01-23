# API implementation for 84codes products in Go

> [!NOTE]
> This library ([and its history](https://github.com/cloudamqp/terraform-provider-cloudamqp/pull/282)) moved June 2024 to the [Terraform Provider for CloudAMQP](https://github.com/cloudamqp/terraform-provider-cloudamqp) repository.

Go implementation for interacting with the API for [CloudAMQP](https://www.cloudamqp.com),
[CloudKarafka](https://www.cloudkarafka.com), [CloudMQTT](https://www.cloudmqtt.com) and [ElephantSQL](https://www.elephantsql.com)

```
useragent := "" // Default set to '84codes go-api'
api := api.New("https://customer.cloudamqp.com", "<YOUR_API_KEY>", useragent)
params := map[string]interface{}{"name": "test", "plan": "bunny", "region": "amazon-web-services::us-east-1"}
instance_info := api.Create(params)
```
