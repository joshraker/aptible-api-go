# CreateMetricDrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** |  | 
**DrainType** | **string** |  | 
**DrainConfiguration** | Pointer to [**CreateMetricDrainRequestDrainConfiguration**](CreateMetricDrainRequestDrainConfiguration.md) |  | [optional] 
**Database** | Pointer to **string** | For Deploy-hosted InfluxDB instances | [optional] 
**DatabaseId** | Pointer to **int32** | For Deploy-hosted InfluxDB instances | [optional] 

## Methods

### NewCreateMetricDrainRequest

`func NewCreateMetricDrainRequest(handle string, drainType string, ) *CreateMetricDrainRequest`

NewCreateMetricDrainRequest instantiates a new CreateMetricDrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetricDrainRequestWithDefaults

`func NewCreateMetricDrainRequestWithDefaults() *CreateMetricDrainRequest`

NewCreateMetricDrainRequestWithDefaults instantiates a new CreateMetricDrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *CreateMetricDrainRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateMetricDrainRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateMetricDrainRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetDrainType

`func (o *CreateMetricDrainRequest) GetDrainType() string`

GetDrainType returns the DrainType field if non-nil, zero value otherwise.

### GetDrainTypeOk

`func (o *CreateMetricDrainRequest) GetDrainTypeOk() (*string, bool)`

GetDrainTypeOk returns a tuple with the DrainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainType

`func (o *CreateMetricDrainRequest) SetDrainType(v string)`

SetDrainType sets DrainType field to given value.


### GetDrainConfiguration

`func (o *CreateMetricDrainRequest) GetDrainConfiguration() CreateMetricDrainRequestDrainConfiguration`

GetDrainConfiguration returns the DrainConfiguration field if non-nil, zero value otherwise.

### GetDrainConfigurationOk

`func (o *CreateMetricDrainRequest) GetDrainConfigurationOk() (*CreateMetricDrainRequestDrainConfiguration, bool)`

GetDrainConfigurationOk returns a tuple with the DrainConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainConfiguration

`func (o *CreateMetricDrainRequest) SetDrainConfiguration(v CreateMetricDrainRequestDrainConfiguration)`

SetDrainConfiguration sets DrainConfiguration field to given value.

### HasDrainConfiguration

`func (o *CreateMetricDrainRequest) HasDrainConfiguration() bool`

HasDrainConfiguration returns a boolean if a field has been set.

### GetDatabase

`func (o *CreateMetricDrainRequest) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *CreateMetricDrainRequest) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *CreateMetricDrainRequest) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *CreateMetricDrainRequest) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDatabaseId

`func (o *CreateMetricDrainRequest) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *CreateMetricDrainRequest) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *CreateMetricDrainRequest) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *CreateMetricDrainRequest) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


