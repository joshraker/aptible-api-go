# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Status** | **string** |  | 
**DockerName** | **NullableString** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Layer** | **string** |  | 
**AwsInstanceId** | **NullableString** |  | 
**InstanceClass** | **NullableString** |  | 
**MemoryLimit** | **int32** |  | 
**PortMapping** | **[][]int32** |  | 
**Allocation** | **[]string** |  | 
**Mounts** | **[][]string** |  | 
**Links** | Pointer to [**ContainerLinks**](ContainerLinks.md) |  | [optional] 

## Methods

### NewContainer

`func NewContainer(id int32, host string, port int32, status string, dockerName NullableString, metaType string, createdAt string, updatedAt string, layer string, awsInstanceId NullableString, instanceClass NullableString, memoryLimit int32, portMapping [][]int32, allocation []string, mounts [][]string, ) *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Container) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Container) SetId(v int32)`

SetId sets Id field to given value.


### GetHost

`func (o *Container) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Container) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Container) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *Container) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Container) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Container) SetPort(v int32)`

SetPort sets Port field to given value.


### GetStatus

`func (o *Container) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Container) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Container) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDockerName

`func (o *Container) GetDockerName() string`

GetDockerName returns the DockerName field if non-nil, zero value otherwise.

### GetDockerNameOk

`func (o *Container) GetDockerNameOk() (*string, bool)`

GetDockerNameOk returns a tuple with the DockerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerName

`func (o *Container) SetDockerName(v string)`

SetDockerName sets DockerName field to given value.


### SetDockerNameNil

`func (o *Container) SetDockerNameNil(b bool)`

 SetDockerNameNil sets the value for DockerName to be an explicit nil

### UnsetDockerName
`func (o *Container) UnsetDockerName()`

UnsetDockerName ensures that no value is present for DockerName, not even an explicit nil
### GetMetaType

`func (o *Container) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Container) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Container) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Container) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Container) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Container) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Container) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Container) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Container) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLayer

`func (o *Container) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *Container) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *Container) SetLayer(v string)`

SetLayer sets Layer field to given value.


### GetAwsInstanceId

`func (o *Container) GetAwsInstanceId() string`

GetAwsInstanceId returns the AwsInstanceId field if non-nil, zero value otherwise.

### GetAwsInstanceIdOk

`func (o *Container) GetAwsInstanceIdOk() (*string, bool)`

GetAwsInstanceIdOk returns a tuple with the AwsInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceId

`func (o *Container) SetAwsInstanceId(v string)`

SetAwsInstanceId sets AwsInstanceId field to given value.


### SetAwsInstanceIdNil

`func (o *Container) SetAwsInstanceIdNil(b bool)`

 SetAwsInstanceIdNil sets the value for AwsInstanceId to be an explicit nil

### UnsetAwsInstanceId
`func (o *Container) UnsetAwsInstanceId()`

UnsetAwsInstanceId ensures that no value is present for AwsInstanceId, not even an explicit nil
### GetInstanceClass

`func (o *Container) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *Container) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *Container) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.


### SetInstanceClassNil

`func (o *Container) SetInstanceClassNil(b bool)`

 SetInstanceClassNil sets the value for InstanceClass to be an explicit nil

### UnsetInstanceClass
`func (o *Container) UnsetInstanceClass()`

UnsetInstanceClass ensures that no value is present for InstanceClass, not even an explicit nil
### GetMemoryLimit

`func (o *Container) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *Container) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *Container) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.


### GetPortMapping

`func (o *Container) GetPortMapping() [][]int32`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *Container) GetPortMappingOk() (*[][]int32, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *Container) SetPortMapping(v [][]int32)`

SetPortMapping sets PortMapping field to given value.


### GetAllocation

`func (o *Container) GetAllocation() []string`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *Container) GetAllocationOk() (*[]string, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *Container) SetAllocation(v []string)`

SetAllocation sets Allocation field to given value.


### GetMounts

`func (o *Container) GetMounts() [][]string`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *Container) GetMountsOk() (*[][]string, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *Container) SetMounts(v [][]string)`

SetMounts sets Mounts field to given value.


### GetLinks

`func (o *Container) GetLinks() ContainerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Container) GetLinksOk() (*ContainerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Container) SetLinks(v ContainerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Container) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


