# AccountEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDrains** | Pointer to [**[]LogDrain**](LogDrain.md) |  | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) |  | [optional] 

## Methods

### NewAccountEmbedded

`func NewAccountEmbedded() *AccountEmbedded`

NewAccountEmbedded instantiates a new AccountEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEmbeddedWithDefaults

`func NewAccountEmbeddedWithDefaults() *AccountEmbedded`

NewAccountEmbeddedWithDefaults instantiates a new AccountEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDrains

`func (o *AccountEmbedded) GetLogDrains() []LogDrain`

GetLogDrains returns the LogDrains field if non-nil, zero value otherwise.

### GetLogDrainsOk

`func (o *AccountEmbedded) GetLogDrainsOk() (*[]LogDrain, bool)`

GetLogDrainsOk returns a tuple with the LogDrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDrains

`func (o *AccountEmbedded) SetLogDrains(v []LogDrain)`

SetLogDrains sets LogDrains field to given value.

### HasLogDrains

`func (o *AccountEmbedded) HasLogDrains() bool`

HasLogDrains returns a boolean if a field has been set.

### GetPermissions

`func (o *AccountEmbedded) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccountEmbedded) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccountEmbedded) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccountEmbedded) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


