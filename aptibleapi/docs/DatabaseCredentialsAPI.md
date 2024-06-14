# \DatabaseCredentialsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseCredential**](DatabaseCredentialsAPI.md#GetDatabaseCredential) | **Get** /database_credentials/{id} | show database_credential
[**ListDatabaseCredentialsForDatabase**](DatabaseCredentialsAPI.md#ListDatabaseCredentialsForDatabase) | **Get** /databases/{database_id}/database_credentials | list database_credentials



## GetDatabaseCredential

> DatabaseCredential GetDatabaseCredential(ctx, id).Execute()

show database_credential

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
	resp, r, err := apiClient.DatabaseCredentialsAPI.GetDatabaseCredential(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseCredentialsAPI.GetDatabaseCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseCredential`: DatabaseCredential
	fmt.Fprintf(os.Stdout, "Response from `DatabaseCredentialsAPI.GetDatabaseCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseCredential**](DatabaseCredential.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseCredentialsForDatabase

> ListDatabaseCredentialsForDatabase200Response ListDatabaseCredentialsForDatabase(ctx, databaseId).Page(page).Execute()

list database_credentials

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
	databaseId := int32(56) // int32 | database_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseCredentialsAPI.ListDatabaseCredentialsForDatabase(context.Background(), databaseId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseCredentialsAPI.ListDatabaseCredentialsForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseCredentialsForDatabase`: ListDatabaseCredentialsForDatabase200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabaseCredentialsAPI.ListDatabaseCredentialsForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseCredentialsForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDatabaseCredentialsForDatabase200Response**](ListDatabaseCredentialsForDatabase200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

