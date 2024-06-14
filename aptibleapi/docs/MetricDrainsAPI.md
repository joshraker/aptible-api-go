# \MetricDrainsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricDrain**](MetricDrainsAPI.md#CreateMetricDrain) | **Post** /accounts/{account_id}/metric_drains | create metric_drain
[**GetMetricDrain**](MetricDrainsAPI.md#GetMetricDrain) | **Get** /metric_drains/{id} | show metric_drain
[**ListMetricDrainsForAccount**](MetricDrainsAPI.md#ListMetricDrainsForAccount) | **Get** /accounts/{account_id}/metric_drains | list metric_drains



## CreateMetricDrain

> MetricDrain CreateMetricDrain(ctx, accountId).CreateMetricDrainRequest(createMetricDrainRequest).Execute()

create metric_drain

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	accountId := int32(56) // int32 | account_id
	createMetricDrainRequest := *openapiclient.NewCreateMetricDrainRequest("Handle_example", "DrainType_example") // CreateMetricDrainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricDrainsAPI.CreateMetricDrain(context.Background(), accountId).CreateMetricDrainRequest(createMetricDrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricDrainsAPI.CreateMetricDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetricDrain`: MetricDrain
	fmt.Fprintf(os.Stdout, "Response from `MetricDrainsAPI.CreateMetricDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMetricDrainRequest** | [**CreateMetricDrainRequest**](CreateMetricDrainRequest.md) |  | 

### Return type

[**MetricDrain**](MetricDrain.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricDrain

> MetricDrain GetMetricDrain(ctx, id).Execute()

show metric_drain

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	id := int32(56) // int32 | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricDrainsAPI.GetMetricDrain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricDrainsAPI.GetMetricDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricDrain`: MetricDrain
	fmt.Fprintf(os.Stdout, "Response from `MetricDrainsAPI.GetMetricDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricDrain**](MetricDrain.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricDrainsForAccount

> ListMetricDrainsForAccount200Response ListMetricDrainsForAccount(ctx, accountId).Page(page).Execute()

list metric_drains

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	accountId := int32(56) // int32 | account_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricDrainsAPI.ListMetricDrainsForAccount(context.Background(), accountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricDrainsAPI.ListMetricDrainsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricDrainsForAccount`: ListMetricDrainsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricDrainsAPI.ListMetricDrainsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetricDrainsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListMetricDrainsForAccount200Response**](ListMetricDrainsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

