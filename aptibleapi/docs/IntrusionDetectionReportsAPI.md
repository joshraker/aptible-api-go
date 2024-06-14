# \IntrusionDetectionReportsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIntrustionDetectionReport**](IntrusionDetectionReportsAPI.md#GetIntrustionDetectionReport) | **Get** /intrusion_detection_reports/{id} | show intrusion_detection_report
[**GetIntrustionDetectionReportCsvDownload**](IntrusionDetectionReportsAPI.md#GetIntrustionDetectionReportCsvDownload) | **Get** /intrusion_detection_reports/{intrusion_report_id}/download_csv | download_csv intrusion_detection_report
[**GetIntrustionDetectionReportPdfDownload**](IntrusionDetectionReportsAPI.md#GetIntrustionDetectionReportPdfDownload) | **Get** /intrusion_detection_reports/{intrusion_report_id}/download_pdf | download_pdf intrusion_detection_report
[**ListIntrustionDetectionReportsForStack**](IntrusionDetectionReportsAPI.md#ListIntrustionDetectionReportsForStack) | **Get** /stacks/{stack_id}/intrusion_detection_reports | list intrusion_detection_reports



## GetIntrustionDetectionReport

> IntrusionDetectionReport GetIntrustionDetectionReport(ctx, id).Execute()

show intrusion_detection_report

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
	resp, r, err := apiClient.IntrusionDetectionReportsAPI.GetIntrustionDetectionReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionDetectionReportsAPI.GetIntrustionDetectionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntrustionDetectionReport`: IntrusionDetectionReport
	fmt.Fprintf(os.Stdout, "Response from `IntrusionDetectionReportsAPI.GetIntrustionDetectionReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntrustionDetectionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntrusionDetectionReport**](IntrusionDetectionReport.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntrustionDetectionReportCsvDownload

> GetIntrustionDetectionReportCsvDownload(ctx, intrusionReportId).Execute()

download_csv intrusion_detection_report

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
	intrusionReportId := int32(56) // int32 | intrusion_report_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntrusionDetectionReportsAPI.GetIntrustionDetectionReportCsvDownload(context.Background(), intrusionReportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionDetectionReportsAPI.GetIntrustionDetectionReportCsvDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intrusionReportId** | **int32** | intrusion_report_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntrustionDetectionReportCsvDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntrustionDetectionReportPdfDownload

> GetIntrustionDetectionReportPdfDownload(ctx, intrusionReportId).Execute()

download_pdf intrusion_detection_report

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
	intrusionReportId := int32(56) // int32 | intrusion_report_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntrusionDetectionReportsAPI.GetIntrustionDetectionReportPdfDownload(context.Background(), intrusionReportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionDetectionReportsAPI.GetIntrustionDetectionReportPdfDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intrusionReportId** | **int32** | intrusion_report_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntrustionDetectionReportPdfDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntrustionDetectionReportsForStack

> ListIntrustionDetectionReportsForStack200Response ListIntrustionDetectionReportsForStack(ctx, stackId).Page(page).Execute()

list intrusion_detection_reports

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
	resp, r, err := apiClient.IntrusionDetectionReportsAPI.ListIntrustionDetectionReportsForStack(context.Background(), stackId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionDetectionReportsAPI.ListIntrustionDetectionReportsForStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntrustionDetectionReportsForStack`: ListIntrustionDetectionReportsForStack200Response
	fmt.Fprintf(os.Stdout, "Response from `IntrusionDetectionReportsAPI.ListIntrustionDetectionReportsForStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **int32** | stack_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntrustionDetectionReportsForStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListIntrustionDetectionReportsForStack200Response**](ListIntrustionDetectionReportsForStack200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

