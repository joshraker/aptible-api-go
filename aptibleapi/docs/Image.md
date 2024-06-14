# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**GitRepo** | **NullableString** |  | 
**GitRef** | **NullableString** |  | 
**DockerRepo** | **NullableString** |  | 
**DockerRef** | **NullableString** |  | 
**DualstackHint** | **NullableInt32** |  | 
**Platform** | **NullableString** |  | 
**Release** | **NullableString** |  | 
**Scan** | **NullableString** |  | 
**ExposedPorts** | **[]int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**ImageLinks**](ImageLinks.md) |  | [optional] 

## Methods

### NewImage

`func NewImage(id int32, gitRepo NullableString, gitRef NullableString, dockerRepo NullableString, dockerRef NullableString, dualstackHint NullableInt32, platform NullableString, release NullableString, scan NullableString, exposedPorts []int32, metaType string, createdAt string, updatedAt string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v int32)`

SetId sets Id field to given value.


### GetGitRepo

`func (o *Image) GetGitRepo() string`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *Image) GetGitRepoOk() (*string, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *Image) SetGitRepo(v string)`

SetGitRepo sets GitRepo field to given value.


### SetGitRepoNil

`func (o *Image) SetGitRepoNil(b bool)`

 SetGitRepoNil sets the value for GitRepo to be an explicit nil

### UnsetGitRepo
`func (o *Image) UnsetGitRepo()`

UnsetGitRepo ensures that no value is present for GitRepo, not even an explicit nil
### GetGitRef

`func (o *Image) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *Image) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *Image) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### SetGitRefNil

`func (o *Image) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *Image) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetDockerRepo

`func (o *Image) GetDockerRepo() string`

GetDockerRepo returns the DockerRepo field if non-nil, zero value otherwise.

### GetDockerRepoOk

`func (o *Image) GetDockerRepoOk() (*string, bool)`

GetDockerRepoOk returns a tuple with the DockerRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepo

`func (o *Image) SetDockerRepo(v string)`

SetDockerRepo sets DockerRepo field to given value.


### SetDockerRepoNil

`func (o *Image) SetDockerRepoNil(b bool)`

 SetDockerRepoNil sets the value for DockerRepo to be an explicit nil

### UnsetDockerRepo
`func (o *Image) UnsetDockerRepo()`

UnsetDockerRepo ensures that no value is present for DockerRepo, not even an explicit nil
### GetDockerRef

`func (o *Image) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *Image) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *Image) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.


### SetDockerRefNil

`func (o *Image) SetDockerRefNil(b bool)`

 SetDockerRefNil sets the value for DockerRef to be an explicit nil

### UnsetDockerRef
`func (o *Image) UnsetDockerRef()`

UnsetDockerRef ensures that no value is present for DockerRef, not even an explicit nil
### GetDualstackHint

`func (o *Image) GetDualstackHint() int32`

GetDualstackHint returns the DualstackHint field if non-nil, zero value otherwise.

### GetDualstackHintOk

`func (o *Image) GetDualstackHintOk() (*int32, bool)`

GetDualstackHintOk returns a tuple with the DualstackHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualstackHint

`func (o *Image) SetDualstackHint(v int32)`

SetDualstackHint sets DualstackHint field to given value.


### SetDualstackHintNil

`func (o *Image) SetDualstackHintNil(b bool)`

 SetDualstackHintNil sets the value for DualstackHint to be an explicit nil

### UnsetDualstackHint
`func (o *Image) UnsetDualstackHint()`

UnsetDualstackHint ensures that no value is present for DualstackHint, not even an explicit nil
### GetPlatform

`func (o *Image) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Image) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Image) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### SetPlatformNil

`func (o *Image) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Image) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetRelease

`func (o *Image) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Image) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Image) SetRelease(v string)`

SetRelease sets Release field to given value.


### SetReleaseNil

`func (o *Image) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *Image) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetScan

`func (o *Image) GetScan() string`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *Image) GetScanOk() (*string, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *Image) SetScan(v string)`

SetScan sets Scan field to given value.


### SetScanNil

`func (o *Image) SetScanNil(b bool)`

 SetScanNil sets the value for Scan to be an explicit nil

### UnsetScan
`func (o *Image) UnsetScan()`

UnsetScan ensures that no value is present for Scan, not even an explicit nil
### GetExposedPorts

`func (o *Image) GetExposedPorts() []int32`

GetExposedPorts returns the ExposedPorts field if non-nil, zero value otherwise.

### GetExposedPortsOk

`func (o *Image) GetExposedPortsOk() (*[]int32, bool)`

GetExposedPortsOk returns a tuple with the ExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPorts

`func (o *Image) SetExposedPorts(v []int32)`

SetExposedPorts sets ExposedPorts field to given value.


### GetMetaType

`func (o *Image) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Image) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Image) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Image) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Image) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Image) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Image) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Image) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Image) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *Image) GetLinks() ImageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Image) GetLinksOk() (*ImageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Image) SetLinks(v ImageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Image) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


