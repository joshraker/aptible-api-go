# CreateLogDrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** |  | 
**DrainType** | **string** |  | 
**DrainHost** | Pointer to **string** | For syslog-tls-tcp drains | [optional] 
**DrainPort** | Pointer to **int32** | For syslog-tls-tcp drains | [optional] 
**DrainUsername** | Pointer to **string** | DEPRECATED: For legacy Elasticsearch drains | [optional] 
**DrainPassword** | Pointer to **string** | For Tail drains | [optional] 
**Database** | Pointer to **string** | For Deploy-hosted Elasticsearch drains | [optional] 
**DatabaseId** | Pointer to **int32** | For Deploy-hosted Elasticsearch drains | [optional] 
**Url** | Pointer to **string** | For HTTPs Post drains | [optional] 
**LoggingToken** | Pointer to **string** |  | [optional] 
**DrainApps** | Pointer to **bool** |  | [optional] 
**DrainDatabases** | Pointer to **bool** |  | [optional] 
**DrainEphemeralSessions** | Pointer to **bool** |  | [optional] 
**DrainProxies** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateLogDrainRequest

`func NewCreateLogDrainRequest(handle string, drainType string, ) *CreateLogDrainRequest`

NewCreateLogDrainRequest instantiates a new CreateLogDrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogDrainRequestWithDefaults

`func NewCreateLogDrainRequestWithDefaults() *CreateLogDrainRequest`

NewCreateLogDrainRequestWithDefaults instantiates a new CreateLogDrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *CreateLogDrainRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateLogDrainRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateLogDrainRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetDrainType

`func (o *CreateLogDrainRequest) GetDrainType() string`

GetDrainType returns the DrainType field if non-nil, zero value otherwise.

### GetDrainTypeOk

`func (o *CreateLogDrainRequest) GetDrainTypeOk() (*string, bool)`

GetDrainTypeOk returns a tuple with the DrainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainType

`func (o *CreateLogDrainRequest) SetDrainType(v string)`

SetDrainType sets DrainType field to given value.


### GetDrainHost

`func (o *CreateLogDrainRequest) GetDrainHost() string`

GetDrainHost returns the DrainHost field if non-nil, zero value otherwise.

### GetDrainHostOk

`func (o *CreateLogDrainRequest) GetDrainHostOk() (*string, bool)`

GetDrainHostOk returns a tuple with the DrainHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainHost

`func (o *CreateLogDrainRequest) SetDrainHost(v string)`

SetDrainHost sets DrainHost field to given value.

### HasDrainHost

`func (o *CreateLogDrainRequest) HasDrainHost() bool`

HasDrainHost returns a boolean if a field has been set.

### GetDrainPort

`func (o *CreateLogDrainRequest) GetDrainPort() int32`

GetDrainPort returns the DrainPort field if non-nil, zero value otherwise.

### GetDrainPortOk

`func (o *CreateLogDrainRequest) GetDrainPortOk() (*int32, bool)`

GetDrainPortOk returns a tuple with the DrainPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainPort

`func (o *CreateLogDrainRequest) SetDrainPort(v int32)`

SetDrainPort sets DrainPort field to given value.

### HasDrainPort

`func (o *CreateLogDrainRequest) HasDrainPort() bool`

HasDrainPort returns a boolean if a field has been set.

### GetDrainUsername

`func (o *CreateLogDrainRequest) GetDrainUsername() string`

GetDrainUsername returns the DrainUsername field if non-nil, zero value otherwise.

### GetDrainUsernameOk

`func (o *CreateLogDrainRequest) GetDrainUsernameOk() (*string, bool)`

GetDrainUsernameOk returns a tuple with the DrainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainUsername

`func (o *CreateLogDrainRequest) SetDrainUsername(v string)`

SetDrainUsername sets DrainUsername field to given value.

### HasDrainUsername

`func (o *CreateLogDrainRequest) HasDrainUsername() bool`

HasDrainUsername returns a boolean if a field has been set.

### GetDrainPassword

`func (o *CreateLogDrainRequest) GetDrainPassword() string`

GetDrainPassword returns the DrainPassword field if non-nil, zero value otherwise.

### GetDrainPasswordOk

`func (o *CreateLogDrainRequest) GetDrainPasswordOk() (*string, bool)`

GetDrainPasswordOk returns a tuple with the DrainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainPassword

`func (o *CreateLogDrainRequest) SetDrainPassword(v string)`

SetDrainPassword sets DrainPassword field to given value.

### HasDrainPassword

`func (o *CreateLogDrainRequest) HasDrainPassword() bool`

HasDrainPassword returns a boolean if a field has been set.

### GetDatabase

`func (o *CreateLogDrainRequest) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *CreateLogDrainRequest) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *CreateLogDrainRequest) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *CreateLogDrainRequest) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDatabaseId

`func (o *CreateLogDrainRequest) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *CreateLogDrainRequest) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *CreateLogDrainRequest) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *CreateLogDrainRequest) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetUrl

`func (o *CreateLogDrainRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateLogDrainRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateLogDrainRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateLogDrainRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLoggingToken

`func (o *CreateLogDrainRequest) GetLoggingToken() string`

GetLoggingToken returns the LoggingToken field if non-nil, zero value otherwise.

### GetLoggingTokenOk

`func (o *CreateLogDrainRequest) GetLoggingTokenOk() (*string, bool)`

GetLoggingTokenOk returns a tuple with the LoggingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingToken

`func (o *CreateLogDrainRequest) SetLoggingToken(v string)`

SetLoggingToken sets LoggingToken field to given value.

### HasLoggingToken

`func (o *CreateLogDrainRequest) HasLoggingToken() bool`

HasLoggingToken returns a boolean if a field has been set.

### GetDrainApps

`func (o *CreateLogDrainRequest) GetDrainApps() bool`

GetDrainApps returns the DrainApps field if non-nil, zero value otherwise.

### GetDrainAppsOk

`func (o *CreateLogDrainRequest) GetDrainAppsOk() (*bool, bool)`

GetDrainAppsOk returns a tuple with the DrainApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainApps

`func (o *CreateLogDrainRequest) SetDrainApps(v bool)`

SetDrainApps sets DrainApps field to given value.

### HasDrainApps

`func (o *CreateLogDrainRequest) HasDrainApps() bool`

HasDrainApps returns a boolean if a field has been set.

### GetDrainDatabases

`func (o *CreateLogDrainRequest) GetDrainDatabases() bool`

GetDrainDatabases returns the DrainDatabases field if non-nil, zero value otherwise.

### GetDrainDatabasesOk

`func (o *CreateLogDrainRequest) GetDrainDatabasesOk() (*bool, bool)`

GetDrainDatabasesOk returns a tuple with the DrainDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainDatabases

`func (o *CreateLogDrainRequest) SetDrainDatabases(v bool)`

SetDrainDatabases sets DrainDatabases field to given value.

### HasDrainDatabases

`func (o *CreateLogDrainRequest) HasDrainDatabases() bool`

HasDrainDatabases returns a boolean if a field has been set.

### GetDrainEphemeralSessions

`func (o *CreateLogDrainRequest) GetDrainEphemeralSessions() bool`

GetDrainEphemeralSessions returns the DrainEphemeralSessions field if non-nil, zero value otherwise.

### GetDrainEphemeralSessionsOk

`func (o *CreateLogDrainRequest) GetDrainEphemeralSessionsOk() (*bool, bool)`

GetDrainEphemeralSessionsOk returns a tuple with the DrainEphemeralSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainEphemeralSessions

`func (o *CreateLogDrainRequest) SetDrainEphemeralSessions(v bool)`

SetDrainEphemeralSessions sets DrainEphemeralSessions field to given value.

### HasDrainEphemeralSessions

`func (o *CreateLogDrainRequest) HasDrainEphemeralSessions() bool`

HasDrainEphemeralSessions returns a boolean if a field has been set.

### GetDrainProxies

`func (o *CreateLogDrainRequest) GetDrainProxies() bool`

GetDrainProxies returns the DrainProxies field if non-nil, zero value otherwise.

### GetDrainProxiesOk

`func (o *CreateLogDrainRequest) GetDrainProxiesOk() (*bool, bool)`

GetDrainProxiesOk returns a tuple with the DrainProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainProxies

`func (o *CreateLogDrainRequest) SetDrainProxies(v bool)`

SetDrainProxies sets DrainProxies field to given value.

### HasDrainProxies

`func (o *CreateLogDrainRequest) HasDrainProxies() bool`

HasDrainProxies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


