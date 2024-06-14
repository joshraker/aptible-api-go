# CreateOperationForAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**GitRef** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int32** |  | [optional] 
**ContainerCount** | Pointer to **int32** |  | [optional] 
**ContainerSize** | Pointer to **int32** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **map[string]interface{}** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**DestinationRegion** | Pointer to **string** |  | [optional] 
**Interactive** | Pointer to **bool** |  | [optional] 
**DestinationAccount** | Pointer to **string** |  | [optional] 
**DestinationAccountId** | Pointer to **int32** |  | [optional] 
**DockerRef** | Pointer to **string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**KeyArn** | Pointer to **string** |  | [optional] 
**InstanceProfile** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOperationForAppRequest

`func NewCreateOperationForAppRequest(type_ string, ) *CreateOperationForAppRequest`

NewCreateOperationForAppRequest instantiates a new CreateOperationForAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOperationForAppRequestWithDefaults

`func NewCreateOperationForAppRequestWithDefaults() *CreateOperationForAppRequest`

NewCreateOperationForAppRequestWithDefaults instantiates a new CreateOperationForAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOperationForAppRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOperationForAppRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOperationForAppRequest) SetType(v string)`

SetType sets Type field to given value.


### GetGitRef

`func (o *CreateOperationForAppRequest) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *CreateOperationForAppRequest) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *CreateOperationForAppRequest) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *CreateOperationForAppRequest) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetDiskSize

`func (o *CreateOperationForAppRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *CreateOperationForAppRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *CreateOperationForAppRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *CreateOperationForAppRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetContainerCount

`func (o *CreateOperationForAppRequest) GetContainerCount() int32`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *CreateOperationForAppRequest) GetContainerCountOk() (*int32, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *CreateOperationForAppRequest) SetContainerCount(v int32)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *CreateOperationForAppRequest) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetContainerSize

`func (o *CreateOperationForAppRequest) GetContainerSize() int32`

GetContainerSize returns the ContainerSize field if non-nil, zero value otherwise.

### GetContainerSizeOk

`func (o *CreateOperationForAppRequest) GetContainerSizeOk() (*int32, bool)`

GetContainerSizeOk returns a tuple with the ContainerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSize

`func (o *CreateOperationForAppRequest) SetContainerSize(v int32)`

SetContainerSize sets ContainerSize field to given value.

### HasContainerSize

`func (o *CreateOperationForAppRequest) HasContainerSize() bool`

HasContainerSize returns a boolean if a field has been set.

### GetCommand

`func (o *CreateOperationForAppRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateOperationForAppRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateOperationForAppRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CreateOperationForAppRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *CreateOperationForAppRequest) GetEnv() map[string]interface{}`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreateOperationForAppRequest) GetEnvOk() (*map[string]interface{}, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreateOperationForAppRequest) SetEnv(v map[string]interface{})`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreateOperationForAppRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHandle

`func (o *CreateOperationForAppRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateOperationForAppRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateOperationForAppRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CreateOperationForAppRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetCertificate

`func (o *CreateOperationForAppRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateOperationForAppRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateOperationForAppRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CreateOperationForAppRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CreateOperationForAppRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateOperationForAppRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateOperationForAppRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreateOperationForAppRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetDestinationRegion

`func (o *CreateOperationForAppRequest) GetDestinationRegion() string`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *CreateOperationForAppRequest) GetDestinationRegionOk() (*string, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *CreateOperationForAppRequest) SetDestinationRegion(v string)`

SetDestinationRegion sets DestinationRegion field to given value.

### HasDestinationRegion

`func (o *CreateOperationForAppRequest) HasDestinationRegion() bool`

HasDestinationRegion returns a boolean if a field has been set.

### GetInteractive

`func (o *CreateOperationForAppRequest) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *CreateOperationForAppRequest) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *CreateOperationForAppRequest) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *CreateOperationForAppRequest) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetDestinationAccount

`func (o *CreateOperationForAppRequest) GetDestinationAccount() string`

GetDestinationAccount returns the DestinationAccount field if non-nil, zero value otherwise.

### GetDestinationAccountOk

`func (o *CreateOperationForAppRequest) GetDestinationAccountOk() (*string, bool)`

GetDestinationAccountOk returns a tuple with the DestinationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccount

`func (o *CreateOperationForAppRequest) SetDestinationAccount(v string)`

SetDestinationAccount sets DestinationAccount field to given value.

### HasDestinationAccount

`func (o *CreateOperationForAppRequest) HasDestinationAccount() bool`

HasDestinationAccount returns a boolean if a field has been set.

### GetDestinationAccountId

`func (o *CreateOperationForAppRequest) GetDestinationAccountId() int32`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CreateOperationForAppRequest) GetDestinationAccountIdOk() (*int32, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CreateOperationForAppRequest) SetDestinationAccountId(v int32)`

SetDestinationAccountId sets DestinationAccountId field to given value.

### HasDestinationAccountId

`func (o *CreateOperationForAppRequest) HasDestinationAccountId() bool`

HasDestinationAccountId returns a boolean if a field has been set.

### GetDockerRef

`func (o *CreateOperationForAppRequest) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *CreateOperationForAppRequest) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *CreateOperationForAppRequest) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.

### HasDockerRef

`func (o *CreateOperationForAppRequest) HasDockerRef() bool`

HasDockerRef returns a boolean if a field has been set.

### GetAutomated

`func (o *CreateOperationForAppRequest) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *CreateOperationForAppRequest) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *CreateOperationForAppRequest) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *CreateOperationForAppRequest) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetKeyArn

`func (o *CreateOperationForAppRequest) GetKeyArn() string`

GetKeyArn returns the KeyArn field if non-nil, zero value otherwise.

### GetKeyArnOk

`func (o *CreateOperationForAppRequest) GetKeyArnOk() (*string, bool)`

GetKeyArnOk returns a tuple with the KeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArn

`func (o *CreateOperationForAppRequest) SetKeyArn(v string)`

SetKeyArn sets KeyArn field to given value.

### HasKeyArn

`func (o *CreateOperationForAppRequest) HasKeyArn() bool`

HasKeyArn returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *CreateOperationForAppRequest) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *CreateOperationForAppRequest) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *CreateOperationForAppRequest) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *CreateOperationForAppRequest) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


