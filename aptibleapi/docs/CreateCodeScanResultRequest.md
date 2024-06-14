# CreateCodeScanResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **int32** | Alternate name for &#x60;operation_id&#x60; | [optional] 
**OperationId** | **int32** |  | 
**DockerfilePresent** | Pointer to **bool** |  | [optional] 
**ProcfilePresent** | Pointer to **bool** |  | [optional] 
**AptibleYmlPresent** | Pointer to **bool** |  | [optional] 
**GitRef** | **string** |  | 
**GitCommit** | **string** |  | 

## Methods

### NewCreateCodeScanResultRequest

`func NewCreateCodeScanResultRequest(operationId int32, gitRef string, gitCommit string, ) *CreateCodeScanResultRequest`

NewCreateCodeScanResultRequest instantiates a new CreateCodeScanResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCodeScanResultRequestWithDefaults

`func NewCreateCodeScanResultRequestWithDefaults() *CreateCodeScanResultRequest`

NewCreateCodeScanResultRequestWithDefaults instantiates a new CreateCodeScanResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *CreateCodeScanResultRequest) GetOperation() int32`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateCodeScanResultRequest) GetOperationOk() (*int32, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateCodeScanResultRequest) SetOperation(v int32)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *CreateCodeScanResultRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOperationId

`func (o *CreateCodeScanResultRequest) GetOperationId() int32`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *CreateCodeScanResultRequest) GetOperationIdOk() (*int32, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *CreateCodeScanResultRequest) SetOperationId(v int32)`

SetOperationId sets OperationId field to given value.


### GetDockerfilePresent

`func (o *CreateCodeScanResultRequest) GetDockerfilePresent() bool`

GetDockerfilePresent returns the DockerfilePresent field if non-nil, zero value otherwise.

### GetDockerfilePresentOk

`func (o *CreateCodeScanResultRequest) GetDockerfilePresentOk() (*bool, bool)`

GetDockerfilePresentOk returns a tuple with the DockerfilePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePresent

`func (o *CreateCodeScanResultRequest) SetDockerfilePresent(v bool)`

SetDockerfilePresent sets DockerfilePresent field to given value.

### HasDockerfilePresent

`func (o *CreateCodeScanResultRequest) HasDockerfilePresent() bool`

HasDockerfilePresent returns a boolean if a field has been set.

### GetProcfilePresent

`func (o *CreateCodeScanResultRequest) GetProcfilePresent() bool`

GetProcfilePresent returns the ProcfilePresent field if non-nil, zero value otherwise.

### GetProcfilePresentOk

`func (o *CreateCodeScanResultRequest) GetProcfilePresentOk() (*bool, bool)`

GetProcfilePresentOk returns a tuple with the ProcfilePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcfilePresent

`func (o *CreateCodeScanResultRequest) SetProcfilePresent(v bool)`

SetProcfilePresent sets ProcfilePresent field to given value.

### HasProcfilePresent

`func (o *CreateCodeScanResultRequest) HasProcfilePresent() bool`

HasProcfilePresent returns a boolean if a field has been set.

### GetAptibleYmlPresent

`func (o *CreateCodeScanResultRequest) GetAptibleYmlPresent() bool`

GetAptibleYmlPresent returns the AptibleYmlPresent field if non-nil, zero value otherwise.

### GetAptibleYmlPresentOk

`func (o *CreateCodeScanResultRequest) GetAptibleYmlPresentOk() (*bool, bool)`

GetAptibleYmlPresentOk returns a tuple with the AptibleYmlPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptibleYmlPresent

`func (o *CreateCodeScanResultRequest) SetAptibleYmlPresent(v bool)`

SetAptibleYmlPresent sets AptibleYmlPresent field to given value.

### HasAptibleYmlPresent

`func (o *CreateCodeScanResultRequest) HasAptibleYmlPresent() bool`

HasAptibleYmlPresent returns a boolean if a field has been set.

### GetGitRef

`func (o *CreateCodeScanResultRequest) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *CreateCodeScanResultRequest) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *CreateCodeScanResultRequest) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitCommit

`func (o *CreateCodeScanResultRequest) GetGitCommit() string`

GetGitCommit returns the GitCommit field if non-nil, zero value otherwise.

### GetGitCommitOk

`func (o *CreateCodeScanResultRequest) GetGitCommitOk() (*string, bool)`

GetGitCommitOk returns a tuple with the GitCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommit

`func (o *CreateCodeScanResultRequest) SetGitCommit(v string)`

SetGitCommit sets GitCommit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


