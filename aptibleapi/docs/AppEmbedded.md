# AppEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]Service**](Service.md) |  | [optional] 
**LastOperation** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**LastDeployOperation** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**CurrentImage** | Pointer to [**Image**](Image.md) |  | [optional] 

## Methods

### NewAppEmbedded

`func NewAppEmbedded() *AppEmbedded`

NewAppEmbedded instantiates a new AppEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEmbeddedWithDefaults

`func NewAppEmbeddedWithDefaults() *AppEmbedded`

NewAppEmbeddedWithDefaults instantiates a new AppEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *AppEmbedded) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AppEmbedded) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AppEmbedded) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *AppEmbedded) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLastOperation

`func (o *AppEmbedded) GetLastOperation() Operation`

GetLastOperation returns the LastOperation field if non-nil, zero value otherwise.

### GetLastOperationOk

`func (o *AppEmbedded) GetLastOperationOk() (*Operation, bool)`

GetLastOperationOk returns a tuple with the LastOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOperation

`func (o *AppEmbedded) SetLastOperation(v Operation)`

SetLastOperation sets LastOperation field to given value.

### HasLastOperation

`func (o *AppEmbedded) HasLastOperation() bool`

HasLastOperation returns a boolean if a field has been set.

### GetLastDeployOperation

`func (o *AppEmbedded) GetLastDeployOperation() Operation`

GetLastDeployOperation returns the LastDeployOperation field if non-nil, zero value otherwise.

### GetLastDeployOperationOk

`func (o *AppEmbedded) GetLastDeployOperationOk() (*Operation, bool)`

GetLastDeployOperationOk returns a tuple with the LastDeployOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployOperation

`func (o *AppEmbedded) SetLastDeployOperation(v Operation)`

SetLastDeployOperation sets LastDeployOperation field to given value.

### HasLastDeployOperation

`func (o *AppEmbedded) HasLastDeployOperation() bool`

HasLastDeployOperation returns a boolean if a field has been set.

### GetCurrentImage

`func (o *AppEmbedded) GetCurrentImage() Image`

GetCurrentImage returns the CurrentImage field if non-nil, zero value otherwise.

### GetCurrentImageOk

`func (o *AppEmbedded) GetCurrentImageOk() (*Image, bool)`

GetCurrentImageOk returns a tuple with the CurrentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentImage

`func (o *AppEmbedded) SetCurrentImage(v Image)`

SetCurrentImage sets CurrentImage field to given value.

### HasCurrentImage

`func (o *AppEmbedded) HasCurrentImage() bool`

HasCurrentImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


