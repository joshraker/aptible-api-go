# \VpcPeersAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVpcPeer**](VpcPeersAPI.md#GetVpcPeer) | **Get** /vpc_peers/{id} | show vpc_peer
[**ListVpcPeersForStack**](VpcPeersAPI.md#ListVpcPeersForStack) | **Get** /stacks/{stack_id}/vpc_peers | list vpc_peers



## GetVpcPeer

> VpcPeer GetVpcPeer(ctx, id).Execute()

show vpc_peer

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
	resp, r, err := apiClient.VpcPeersAPI.GetVpcPeer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcPeersAPI.GetVpcPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcPeer`: VpcPeer
	fmt.Fprintf(os.Stdout, "Response from `VpcPeersAPI.GetVpcPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpcPeer**](VpcPeer.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcPeersForStack

> ListVpcPeersForStack200Response ListVpcPeersForStack(ctx, stackId).Page(page).Execute()

list vpc_peers

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
	stackId := int32(56) // int32 | stack_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcPeersAPI.ListVpcPeersForStack(context.Background(), stackId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcPeersAPI.ListVpcPeersForStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcPeersForStack`: ListVpcPeersForStack200Response
	fmt.Fprintf(os.Stdout, "Response from `VpcPeersAPI.ListVpcPeersForStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **int32** | stack_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVpcPeersForStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListVpcPeersForStack200Response**](ListVpcPeersForStack200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

