# CreateSshPortalConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPublicKey** | **string** |  | 
**Command** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSshPortalConnectionRequest

`func NewCreateSshPortalConnectionRequest(sshPublicKey string, ) *CreateSshPortalConnectionRequest`

NewCreateSshPortalConnectionRequest instantiates a new CreateSshPortalConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSshPortalConnectionRequestWithDefaults

`func NewCreateSshPortalConnectionRequestWithDefaults() *CreateSshPortalConnectionRequest`

NewCreateSshPortalConnectionRequestWithDefaults instantiates a new CreateSshPortalConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshPublicKey

`func (o *CreateSshPortalConnectionRequest) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *CreateSshPortalConnectionRequest) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *CreateSshPortalConnectionRequest) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.


### GetCommand

`func (o *CreateSshPortalConnectionRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateSshPortalConnectionRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateSshPortalConnectionRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CreateSshPortalConnectionRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


