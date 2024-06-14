# CodeScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DockerfilePresent** | **bool** |  | 
**ProcfilePresent** | **bool** |  | 
**AptibleYmlPresent** | **bool** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**CodeScanResultLinks**](CodeScanResultLinks.md) |  | [optional] 

## Methods

### NewCodeScanResult

`func NewCodeScanResult(id int32, dockerfilePresent bool, procfilePresent bool, aptibleYmlPresent bool, metaType string, createdAt string, updatedAt string, ) *CodeScanResult`

NewCodeScanResult instantiates a new CodeScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanResultWithDefaults

`func NewCodeScanResultWithDefaults() *CodeScanResult`

NewCodeScanResultWithDefaults instantiates a new CodeScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CodeScanResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeScanResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeScanResult) SetId(v int32)`

SetId sets Id field to given value.


### GetDockerfilePresent

`func (o *CodeScanResult) GetDockerfilePresent() bool`

GetDockerfilePresent returns the DockerfilePresent field if non-nil, zero value otherwise.

### GetDockerfilePresentOk

`func (o *CodeScanResult) GetDockerfilePresentOk() (*bool, bool)`

GetDockerfilePresentOk returns a tuple with the DockerfilePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePresent

`func (o *CodeScanResult) SetDockerfilePresent(v bool)`

SetDockerfilePresent sets DockerfilePresent field to given value.


### GetProcfilePresent

`func (o *CodeScanResult) GetProcfilePresent() bool`

GetProcfilePresent returns the ProcfilePresent field if non-nil, zero value otherwise.

### GetProcfilePresentOk

`func (o *CodeScanResult) GetProcfilePresentOk() (*bool, bool)`

GetProcfilePresentOk returns a tuple with the ProcfilePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcfilePresent

`func (o *CodeScanResult) SetProcfilePresent(v bool)`

SetProcfilePresent sets ProcfilePresent field to given value.


### GetAptibleYmlPresent

`func (o *CodeScanResult) GetAptibleYmlPresent() bool`

GetAptibleYmlPresent returns the AptibleYmlPresent field if non-nil, zero value otherwise.

### GetAptibleYmlPresentOk

`func (o *CodeScanResult) GetAptibleYmlPresentOk() (*bool, bool)`

GetAptibleYmlPresentOk returns a tuple with the AptibleYmlPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptibleYmlPresent

`func (o *CodeScanResult) SetAptibleYmlPresent(v bool)`

SetAptibleYmlPresent sets AptibleYmlPresent field to given value.


### GetMetaType

`func (o *CodeScanResult) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *CodeScanResult) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *CodeScanResult) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *CodeScanResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeScanResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeScanResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CodeScanResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CodeScanResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CodeScanResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *CodeScanResult) GetLinks() CodeScanResultLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CodeScanResult) GetLinksOk() (*CodeScanResultLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CodeScanResult) SetLinks(v CodeScanResultLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CodeScanResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


