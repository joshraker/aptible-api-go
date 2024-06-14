# EphemeralContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Layer** | **string** |  | 
**DockerName** | **string** |  | 
**Interactive** | **bool** |  | 
**MemoryLimit** | **int32** |  | 
**Allocation** | **[]string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**MetaType** | **string** |  | 
**Links** | Pointer to [**EphemeralContainerLinks**](EphemeralContainerLinks.md) |  | [optional] 

## Methods

### NewEphemeralContainer

`func NewEphemeralContainer(id int32, layer string, dockerName string, interactive bool, memoryLimit int32, allocation []string, createdAt string, updatedAt string, metaType string, ) *EphemeralContainer`

NewEphemeralContainer instantiates a new EphemeralContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEphemeralContainerWithDefaults

`func NewEphemeralContainerWithDefaults() *EphemeralContainer`

NewEphemeralContainerWithDefaults instantiates a new EphemeralContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EphemeralContainer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EphemeralContainer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EphemeralContainer) SetId(v int32)`

SetId sets Id field to given value.


### GetLayer

`func (o *EphemeralContainer) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *EphemeralContainer) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *EphemeralContainer) SetLayer(v string)`

SetLayer sets Layer field to given value.


### GetDockerName

`func (o *EphemeralContainer) GetDockerName() string`

GetDockerName returns the DockerName field if non-nil, zero value otherwise.

### GetDockerNameOk

`func (o *EphemeralContainer) GetDockerNameOk() (*string, bool)`

GetDockerNameOk returns a tuple with the DockerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerName

`func (o *EphemeralContainer) SetDockerName(v string)`

SetDockerName sets DockerName field to given value.


### GetInteractive

`func (o *EphemeralContainer) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *EphemeralContainer) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *EphemeralContainer) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.


### GetMemoryLimit

`func (o *EphemeralContainer) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *EphemeralContainer) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *EphemeralContainer) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.


### GetAllocation

`func (o *EphemeralContainer) GetAllocation() []string`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *EphemeralContainer) GetAllocationOk() (*[]string, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *EphemeralContainer) SetAllocation(v []string)`

SetAllocation sets Allocation field to given value.


### GetCreatedAt

`func (o *EphemeralContainer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EphemeralContainer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EphemeralContainer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EphemeralContainer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EphemeralContainer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EphemeralContainer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMetaType

`func (o *EphemeralContainer) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *EphemeralContainer) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *EphemeralContainer) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetLinks

`func (o *EphemeralContainer) GetLinks() EphemeralContainerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EphemeralContainer) GetLinksOk() (*EphemeralContainerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EphemeralContainer) SetLinks(v EphemeralContainerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EphemeralContainer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


