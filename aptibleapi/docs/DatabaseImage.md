# DatabaseImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Type** | **string** |  | 
**Version** | **string** |  | 
**Discoverable** | **bool** |  | 
**Visible** | **bool** |  | 
**Default** | **bool** |  | 
**Description** | **string** |  | 
**DockerRepo** | **string** |  | 
**Links** | Pointer to [**DatabaseImageLinks**](DatabaseImageLinks.md) |  | [optional] 

## Methods

### NewDatabaseImage

`func NewDatabaseImage(id int32, metaType string, createdAt string, updatedAt string, type_ string, version string, discoverable bool, visible bool, default_ bool, description string, dockerRepo string, ) *DatabaseImage`

NewDatabaseImage instantiates a new DatabaseImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseImageWithDefaults

`func NewDatabaseImageWithDefaults() *DatabaseImage`

NewDatabaseImageWithDefaults instantiates a new DatabaseImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseImage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseImage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseImage) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *DatabaseImage) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *DatabaseImage) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *DatabaseImage) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *DatabaseImage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseImage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseImage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DatabaseImage) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabaseImage) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabaseImage) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *DatabaseImage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseImage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseImage) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *DatabaseImage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseImage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseImage) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDiscoverable

`func (o *DatabaseImage) GetDiscoverable() bool`

GetDiscoverable returns the Discoverable field if non-nil, zero value otherwise.

### GetDiscoverableOk

`func (o *DatabaseImage) GetDiscoverableOk() (*bool, bool)`

GetDiscoverableOk returns a tuple with the Discoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverable

`func (o *DatabaseImage) SetDiscoverable(v bool)`

SetDiscoverable sets Discoverable field to given value.


### GetVisible

`func (o *DatabaseImage) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *DatabaseImage) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *DatabaseImage) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetDefault

`func (o *DatabaseImage) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DatabaseImage) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DatabaseImage) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetDescription

`func (o *DatabaseImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseImage) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDockerRepo

`func (o *DatabaseImage) GetDockerRepo() string`

GetDockerRepo returns the DockerRepo field if non-nil, zero value otherwise.

### GetDockerRepoOk

`func (o *DatabaseImage) GetDockerRepoOk() (*string, bool)`

GetDockerRepoOk returns a tuple with the DockerRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepo

`func (o *DatabaseImage) SetDockerRepo(v string)`

SetDockerRepo sets DockerRepo field to given value.


### GetLinks

`func (o *DatabaseImage) GetLinks() DatabaseImageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DatabaseImage) GetLinksOk() (*DatabaseImageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DatabaseImage) SetLinks(v DatabaseImageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DatabaseImage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


