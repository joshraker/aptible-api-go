# EphemeralSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AwsInstanceId** | **string** |  | 
**Host** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**MetaType** | **string** |  | 
**Links** | Pointer to [**EphemeralSessionLinks**](EphemeralSessionLinks.md) |  | [optional] 

## Methods

### NewEphemeralSession

`func NewEphemeralSession(id int32, awsInstanceId string, host string, name string, createdAt string, updatedAt string, metaType string, ) *EphemeralSession`

NewEphemeralSession instantiates a new EphemeralSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEphemeralSessionWithDefaults

`func NewEphemeralSessionWithDefaults() *EphemeralSession`

NewEphemeralSessionWithDefaults instantiates a new EphemeralSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EphemeralSession) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EphemeralSession) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EphemeralSession) SetId(v int32)`

SetId sets Id field to given value.


### GetAwsInstanceId

`func (o *EphemeralSession) GetAwsInstanceId() string`

GetAwsInstanceId returns the AwsInstanceId field if non-nil, zero value otherwise.

### GetAwsInstanceIdOk

`func (o *EphemeralSession) GetAwsInstanceIdOk() (*string, bool)`

GetAwsInstanceIdOk returns a tuple with the AwsInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceId

`func (o *EphemeralSession) SetAwsInstanceId(v string)`

SetAwsInstanceId sets AwsInstanceId field to given value.


### GetHost

`func (o *EphemeralSession) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EphemeralSession) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EphemeralSession) SetHost(v string)`

SetHost sets Host field to given value.


### GetName

`func (o *EphemeralSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EphemeralSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EphemeralSession) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *EphemeralSession) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EphemeralSession) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EphemeralSession) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EphemeralSession) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EphemeralSession) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EphemeralSession) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMetaType

`func (o *EphemeralSession) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *EphemeralSession) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *EphemeralSession) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetLinks

`func (o *EphemeralSession) GetLinks() EphemeralSessionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EphemeralSession) GetLinksOk() (*EphemeralSessionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EphemeralSession) SetLinks(v EphemeralSessionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EphemeralSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


