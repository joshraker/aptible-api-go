# PrereleaseCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Command** | **string** |  | 
**Index** | **int32** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**PrereleaseCommandLinks**](PrereleaseCommandLinks.md) |  | [optional] 

## Methods

### NewPrereleaseCommand

`func NewPrereleaseCommand(id int32, metaType string, command string, index int32, createdAt string, updatedAt string, ) *PrereleaseCommand`

NewPrereleaseCommand instantiates a new PrereleaseCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrereleaseCommandWithDefaults

`func NewPrereleaseCommandWithDefaults() *PrereleaseCommand`

NewPrereleaseCommandWithDefaults instantiates a new PrereleaseCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrereleaseCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrereleaseCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrereleaseCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *PrereleaseCommand) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *PrereleaseCommand) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *PrereleaseCommand) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCommand

`func (o *PrereleaseCommand) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PrereleaseCommand) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PrereleaseCommand) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetIndex

`func (o *PrereleaseCommand) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PrereleaseCommand) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PrereleaseCommand) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetCreatedAt

`func (o *PrereleaseCommand) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrereleaseCommand) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrereleaseCommand) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PrereleaseCommand) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrereleaseCommand) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrereleaseCommand) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *PrereleaseCommand) GetLinks() PrereleaseCommandLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrereleaseCommand) GetLinksOk() (*PrereleaseCommandLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrereleaseCommand) SetLinks(v PrereleaseCommandLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PrereleaseCommand) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


