# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Handle** | **string** |  | 
**GitRepo** | **string** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Status** | **string** |  | 
**DeploymentMethod** | **NullableString** |  | 
**Embedded** | [**AppEmbedded**](AppEmbedded.md) |  | 
**Links** | Pointer to [**AppLinks**](AppLinks.md) |  | [optional] 

## Methods

### NewApp

`func NewApp(id int32, handle string, gitRepo string, metaType string, createdAt string, updatedAt string, status string, deploymentMethod NullableString, embedded AppEmbedded, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v int32)`

SetId sets Id field to given value.


### GetHandle

`func (o *App) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *App) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *App) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetGitRepo

`func (o *App) GetGitRepo() string`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *App) GetGitRepoOk() (*string, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *App) SetGitRepo(v string)`

SetGitRepo sets GitRepo field to given value.


### GetMetaType

`func (o *App) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *App) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *App) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *App) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *App) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *App) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *App) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *App) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *App) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *App) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *App) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *App) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeploymentMethod

`func (o *App) GetDeploymentMethod() string`

GetDeploymentMethod returns the DeploymentMethod field if non-nil, zero value otherwise.

### GetDeploymentMethodOk

`func (o *App) GetDeploymentMethodOk() (*string, bool)`

GetDeploymentMethodOk returns a tuple with the DeploymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMethod

`func (o *App) SetDeploymentMethod(v string)`

SetDeploymentMethod sets DeploymentMethod field to given value.


### SetDeploymentMethodNil

`func (o *App) SetDeploymentMethodNil(b bool)`

 SetDeploymentMethodNil sets the value for DeploymentMethod to be an explicit nil

### UnsetDeploymentMethod
`func (o *App) UnsetDeploymentMethod()`

UnsetDeploymentMethod ensures that no value is present for DeploymentMethod, not even an explicit nil
### GetEmbedded

`func (o *App) GetEmbedded() AppEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *App) GetEmbeddedOk() (*AppEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *App) SetEmbedded(v AppEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *App) GetLinks() AppLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *App) GetLinksOk() (*AppLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *App) SetLinks(v AppLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *App) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


