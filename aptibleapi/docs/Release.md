# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**DockerRepo** | **NullableString** |  | 
**DockerRef** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**ReleaseLinks**](ReleaseLinks.md) |  | [optional] 

## Methods

### NewRelease

`func NewRelease(id int32, metaType string, dockerRepo NullableString, dockerRef NullableString, createdAt string, updatedAt string, ) *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Release) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Release) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Release) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Release) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Release) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Release) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetDockerRepo

`func (o *Release) GetDockerRepo() string`

GetDockerRepo returns the DockerRepo field if non-nil, zero value otherwise.

### GetDockerRepoOk

`func (o *Release) GetDockerRepoOk() (*string, bool)`

GetDockerRepoOk returns a tuple with the DockerRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepo

`func (o *Release) SetDockerRepo(v string)`

SetDockerRepo sets DockerRepo field to given value.


### SetDockerRepoNil

`func (o *Release) SetDockerRepoNil(b bool)`

 SetDockerRepoNil sets the value for DockerRepo to be an explicit nil

### UnsetDockerRepo
`func (o *Release) UnsetDockerRepo()`

UnsetDockerRepo ensures that no value is present for DockerRepo, not even an explicit nil
### GetDockerRef

`func (o *Release) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *Release) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *Release) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.


### SetDockerRefNil

`func (o *Release) SetDockerRefNil(b bool)`

 SetDockerRefNil sets the value for DockerRef to be an explicit nil

### UnsetDockerRef
`func (o *Release) UnsetDockerRef()`

UnsetDockerRef ensures that no value is present for DockerRef, not even an explicit nil
### GetCreatedAt

`func (o *Release) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Release) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Release) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Release) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Release) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Release) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *Release) GetLinks() ReleaseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Release) GetLinksOk() (*ReleaseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Release) SetLinks(v ReleaseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Release) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


