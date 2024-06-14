# DatabaseEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastOperation** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Disk** | Pointer to [**Disk**](Disk.md) |  | [optional] 
**DatabaseCredentials** | Pointer to [**[]DatabaseCredential**](DatabaseCredential.md) |  | [optional] 

## Methods

### NewDatabaseEmbedded

`func NewDatabaseEmbedded() *DatabaseEmbedded`

NewDatabaseEmbedded instantiates a new DatabaseEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseEmbeddedWithDefaults

`func NewDatabaseEmbeddedWithDefaults() *DatabaseEmbedded`

NewDatabaseEmbeddedWithDefaults instantiates a new DatabaseEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastOperation

`func (o *DatabaseEmbedded) GetLastOperation() Operation`

GetLastOperation returns the LastOperation field if non-nil, zero value otherwise.

### GetLastOperationOk

`func (o *DatabaseEmbedded) GetLastOperationOk() (*Operation, bool)`

GetLastOperationOk returns a tuple with the LastOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOperation

`func (o *DatabaseEmbedded) SetLastOperation(v Operation)`

SetLastOperation sets LastOperation field to given value.

### HasLastOperation

`func (o *DatabaseEmbedded) HasLastOperation() bool`

HasLastOperation returns a boolean if a field has been set.

### GetDisk

`func (o *DatabaseEmbedded) GetDisk() Disk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *DatabaseEmbedded) GetDiskOk() (*Disk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *DatabaseEmbedded) SetDisk(v Disk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *DatabaseEmbedded) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDatabaseCredentials

`func (o *DatabaseEmbedded) GetDatabaseCredentials() []DatabaseCredential`

GetDatabaseCredentials returns the DatabaseCredentials field if non-nil, zero value otherwise.

### GetDatabaseCredentialsOk

`func (o *DatabaseEmbedded) GetDatabaseCredentialsOk() (*[]DatabaseCredential, bool)`

GetDatabaseCredentialsOk returns a tuple with the DatabaseCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCredentials

`func (o *DatabaseEmbedded) SetDatabaseCredentials(v []DatabaseCredential)`

SetDatabaseCredentials sets DatabaseCredentials field to given value.

### HasDatabaseCredentials

`func (o *DatabaseEmbedded) HasDatabaseCredentials() bool`

HasDatabaseCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


