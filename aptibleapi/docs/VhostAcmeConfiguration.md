# VhostAcmeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **[]string** |  | [optional] 
**Challenges** | Pointer to [**[]VhostAcmeConfigurationChallengesInner**](VhostAcmeConfigurationChallengesInner.md) |  | [optional] 

## Methods

### NewVhostAcmeConfiguration

`func NewVhostAcmeConfiguration() *VhostAcmeConfiguration`

NewVhostAcmeConfiguration instantiates a new VhostAcmeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhostAcmeConfigurationWithDefaults

`func NewVhostAcmeConfigurationWithDefaults() *VhostAcmeConfiguration`

NewVhostAcmeConfigurationWithDefaults instantiates a new VhostAcmeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *VhostAcmeConfiguration) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VhostAcmeConfiguration) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VhostAcmeConfiguration) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VhostAcmeConfiguration) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetChallenges

`func (o *VhostAcmeConfiguration) GetChallenges() []VhostAcmeConfigurationChallengesInner`

GetChallenges returns the Challenges field if non-nil, zero value otherwise.

### GetChallengesOk

`func (o *VhostAcmeConfiguration) GetChallengesOk() (*[]VhostAcmeConfigurationChallengesInner, bool)`

GetChallengesOk returns a tuple with the Challenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenges

`func (o *VhostAcmeConfiguration) SetChallenges(v []VhostAcmeConfigurationChallengesInner)`

SetChallenges sets Challenges field to given value.

### HasChallenges

`func (o *VhostAcmeConfiguration) HasChallenges() bool`

HasChallenges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


