# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Type** | **string** |  | 
**MetaType** | **string** |  | 
**Name** | **NullableString** |  | 
**Handle** | **string** |  | 
**OnboardingStatus** | **string** |  | 
**Number** | **string** |  | 
**Activated** | **bool** |  | 
**SyslogHost** | **NullableString** |  | 
**SyslogPort** | **NullableInt32** |  | 
**BastionHost** | **NullableString** |  | 
**BastionPort** | **int32** |  | 
**SshPortalPort** | **NullableInt32** |  | 
**ContainerCount** | **int32** |  | 
**DomainCount** | **int32** |  | 
**TotalDiskSize** | **int32** |  | 
**TotalAppCount** | **int32** |  | 
**AppContainerCount** | **int32** |  | 
**DatabaseContainerCount** | **int32** |  | 
**TotalDatabaseCount** | **int32** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**SweetnessStack** | **NullableString** |  | 
**GentlemanjerryEndpoint** | **NullableString** |  | 
**GentlemanjerryCertificate** | **NullableString** |  | 
**GentlemanjerryDockerName** | **NullableString** |  | 
**GentlemanjerryInstanceId** | **NullableString** |  | 
**SweetnessStackVersion** | **NullableString** |  | 
**CaBody** | **NullableString** |  | 
**CaPrivateKey** | Pointer to **NullableString** |  | [optional] 
**Embedded** | [**AccountEmbedded**](AccountEmbedded.md) |  | 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount(id int32, type_ string, metaType string, name NullableString, handle string, onboardingStatus string, number string, activated bool, syslogHost NullableString, syslogPort NullableInt32, bastionHost NullableString, bastionPort int32, sshPortalPort NullableInt32, containerCount int32, domainCount int32, totalDiskSize int32, totalAppCount int32, appContainerCount int32, databaseContainerCount int32, totalDatabaseCount int32, createdAt string, updatedAt string, sweetnessStack NullableString, gentlemanjerryEndpoint NullableString, gentlemanjerryCertificate NullableString, gentlemanjerryDockerName NullableString, gentlemanjerryInstanceId NullableString, sweetnessStackVersion NullableString, caBody NullableString, embedded AccountEmbedded, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.


### GetMetaType

`func (o *Account) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Account) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Account) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Account) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Account) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHandle

`func (o *Account) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Account) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Account) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetOnboardingStatus

`func (o *Account) GetOnboardingStatus() string`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *Account) GetOnboardingStatusOk() (*string, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingStatus

`func (o *Account) SetOnboardingStatus(v string)`

SetOnboardingStatus sets OnboardingStatus field to given value.


### GetNumber

`func (o *Account) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Account) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Account) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetActivated

`func (o *Account) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *Account) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *Account) SetActivated(v bool)`

SetActivated sets Activated field to given value.


### GetSyslogHost

`func (o *Account) GetSyslogHost() string`

GetSyslogHost returns the SyslogHost field if non-nil, zero value otherwise.

### GetSyslogHostOk

`func (o *Account) GetSyslogHostOk() (*string, bool)`

GetSyslogHostOk returns a tuple with the SyslogHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogHost

`func (o *Account) SetSyslogHost(v string)`

SetSyslogHost sets SyslogHost field to given value.


### SetSyslogHostNil

`func (o *Account) SetSyslogHostNil(b bool)`

 SetSyslogHostNil sets the value for SyslogHost to be an explicit nil

### UnsetSyslogHost
`func (o *Account) UnsetSyslogHost()`

UnsetSyslogHost ensures that no value is present for SyslogHost, not even an explicit nil
### GetSyslogPort

`func (o *Account) GetSyslogPort() int32`

GetSyslogPort returns the SyslogPort field if non-nil, zero value otherwise.

### GetSyslogPortOk

`func (o *Account) GetSyslogPortOk() (*int32, bool)`

GetSyslogPortOk returns a tuple with the SyslogPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogPort

`func (o *Account) SetSyslogPort(v int32)`

SetSyslogPort sets SyslogPort field to given value.


### SetSyslogPortNil

`func (o *Account) SetSyslogPortNil(b bool)`

 SetSyslogPortNil sets the value for SyslogPort to be an explicit nil

### UnsetSyslogPort
`func (o *Account) UnsetSyslogPort()`

UnsetSyslogPort ensures that no value is present for SyslogPort, not even an explicit nil
### GetBastionHost

`func (o *Account) GetBastionHost() string`

GetBastionHost returns the BastionHost field if non-nil, zero value otherwise.

### GetBastionHostOk

`func (o *Account) GetBastionHostOk() (*string, bool)`

GetBastionHostOk returns a tuple with the BastionHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionHost

`func (o *Account) SetBastionHost(v string)`

SetBastionHost sets BastionHost field to given value.


### SetBastionHostNil

`func (o *Account) SetBastionHostNil(b bool)`

 SetBastionHostNil sets the value for BastionHost to be an explicit nil

### UnsetBastionHost
`func (o *Account) UnsetBastionHost()`

UnsetBastionHost ensures that no value is present for BastionHost, not even an explicit nil
### GetBastionPort

`func (o *Account) GetBastionPort() int32`

GetBastionPort returns the BastionPort field if non-nil, zero value otherwise.

### GetBastionPortOk

`func (o *Account) GetBastionPortOk() (*int32, bool)`

GetBastionPortOk returns a tuple with the BastionPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionPort

`func (o *Account) SetBastionPort(v int32)`

SetBastionPort sets BastionPort field to given value.


### GetSshPortalPort

`func (o *Account) GetSshPortalPort() int32`

GetSshPortalPort returns the SshPortalPort field if non-nil, zero value otherwise.

### GetSshPortalPortOk

`func (o *Account) GetSshPortalPortOk() (*int32, bool)`

GetSshPortalPortOk returns a tuple with the SshPortalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortalPort

`func (o *Account) SetSshPortalPort(v int32)`

SetSshPortalPort sets SshPortalPort field to given value.


### SetSshPortalPortNil

`func (o *Account) SetSshPortalPortNil(b bool)`

 SetSshPortalPortNil sets the value for SshPortalPort to be an explicit nil

### UnsetSshPortalPort
`func (o *Account) UnsetSshPortalPort()`

UnsetSshPortalPort ensures that no value is present for SshPortalPort, not even an explicit nil
### GetContainerCount

`func (o *Account) GetContainerCount() int32`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *Account) GetContainerCountOk() (*int32, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *Account) SetContainerCount(v int32)`

SetContainerCount sets ContainerCount field to given value.


### GetDomainCount

`func (o *Account) GetDomainCount() int32`

GetDomainCount returns the DomainCount field if non-nil, zero value otherwise.

### GetDomainCountOk

`func (o *Account) GetDomainCountOk() (*int32, bool)`

GetDomainCountOk returns a tuple with the DomainCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCount

`func (o *Account) SetDomainCount(v int32)`

SetDomainCount sets DomainCount field to given value.


### GetTotalDiskSize

`func (o *Account) GetTotalDiskSize() int32`

GetTotalDiskSize returns the TotalDiskSize field if non-nil, zero value otherwise.

### GetTotalDiskSizeOk

`func (o *Account) GetTotalDiskSizeOk() (*int32, bool)`

GetTotalDiskSizeOk returns a tuple with the TotalDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiskSize

`func (o *Account) SetTotalDiskSize(v int32)`

SetTotalDiskSize sets TotalDiskSize field to given value.


### GetTotalAppCount

`func (o *Account) GetTotalAppCount() int32`

GetTotalAppCount returns the TotalAppCount field if non-nil, zero value otherwise.

### GetTotalAppCountOk

`func (o *Account) GetTotalAppCountOk() (*int32, bool)`

GetTotalAppCountOk returns a tuple with the TotalAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAppCount

`func (o *Account) SetTotalAppCount(v int32)`

SetTotalAppCount sets TotalAppCount field to given value.


### GetAppContainerCount

`func (o *Account) GetAppContainerCount() int32`

GetAppContainerCount returns the AppContainerCount field if non-nil, zero value otherwise.

### GetAppContainerCountOk

`func (o *Account) GetAppContainerCountOk() (*int32, bool)`

GetAppContainerCountOk returns a tuple with the AppContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContainerCount

`func (o *Account) SetAppContainerCount(v int32)`

SetAppContainerCount sets AppContainerCount field to given value.


### GetDatabaseContainerCount

`func (o *Account) GetDatabaseContainerCount() int32`

GetDatabaseContainerCount returns the DatabaseContainerCount field if non-nil, zero value otherwise.

### GetDatabaseContainerCountOk

`func (o *Account) GetDatabaseContainerCountOk() (*int32, bool)`

GetDatabaseContainerCountOk returns a tuple with the DatabaseContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseContainerCount

`func (o *Account) SetDatabaseContainerCount(v int32)`

SetDatabaseContainerCount sets DatabaseContainerCount field to given value.


### GetTotalDatabaseCount

`func (o *Account) GetTotalDatabaseCount() int32`

GetTotalDatabaseCount returns the TotalDatabaseCount field if non-nil, zero value otherwise.

### GetTotalDatabaseCountOk

`func (o *Account) GetTotalDatabaseCountOk() (*int32, bool)`

GetTotalDatabaseCountOk returns a tuple with the TotalDatabaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDatabaseCount

`func (o *Account) SetTotalDatabaseCount(v int32)`

SetTotalDatabaseCount sets TotalDatabaseCount field to given value.


### GetCreatedAt

`func (o *Account) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Account) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSweetnessStack

`func (o *Account) GetSweetnessStack() string`

GetSweetnessStack returns the SweetnessStack field if non-nil, zero value otherwise.

### GetSweetnessStackOk

`func (o *Account) GetSweetnessStackOk() (*string, bool)`

GetSweetnessStackOk returns a tuple with the SweetnessStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweetnessStack

`func (o *Account) SetSweetnessStack(v string)`

SetSweetnessStack sets SweetnessStack field to given value.


### SetSweetnessStackNil

`func (o *Account) SetSweetnessStackNil(b bool)`

 SetSweetnessStackNil sets the value for SweetnessStack to be an explicit nil

### UnsetSweetnessStack
`func (o *Account) UnsetSweetnessStack()`

UnsetSweetnessStack ensures that no value is present for SweetnessStack, not even an explicit nil
### GetGentlemanjerryEndpoint

`func (o *Account) GetGentlemanjerryEndpoint() string`

GetGentlemanjerryEndpoint returns the GentlemanjerryEndpoint field if non-nil, zero value otherwise.

### GetGentlemanjerryEndpointOk

`func (o *Account) GetGentlemanjerryEndpointOk() (*string, bool)`

GetGentlemanjerryEndpointOk returns a tuple with the GentlemanjerryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryEndpoint

`func (o *Account) SetGentlemanjerryEndpoint(v string)`

SetGentlemanjerryEndpoint sets GentlemanjerryEndpoint field to given value.


### SetGentlemanjerryEndpointNil

`func (o *Account) SetGentlemanjerryEndpointNil(b bool)`

 SetGentlemanjerryEndpointNil sets the value for GentlemanjerryEndpoint to be an explicit nil

### UnsetGentlemanjerryEndpoint
`func (o *Account) UnsetGentlemanjerryEndpoint()`

UnsetGentlemanjerryEndpoint ensures that no value is present for GentlemanjerryEndpoint, not even an explicit nil
### GetGentlemanjerryCertificate

`func (o *Account) GetGentlemanjerryCertificate() string`

GetGentlemanjerryCertificate returns the GentlemanjerryCertificate field if non-nil, zero value otherwise.

### GetGentlemanjerryCertificateOk

`func (o *Account) GetGentlemanjerryCertificateOk() (*string, bool)`

GetGentlemanjerryCertificateOk returns a tuple with the GentlemanjerryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryCertificate

`func (o *Account) SetGentlemanjerryCertificate(v string)`

SetGentlemanjerryCertificate sets GentlemanjerryCertificate field to given value.


### SetGentlemanjerryCertificateNil

`func (o *Account) SetGentlemanjerryCertificateNil(b bool)`

 SetGentlemanjerryCertificateNil sets the value for GentlemanjerryCertificate to be an explicit nil

### UnsetGentlemanjerryCertificate
`func (o *Account) UnsetGentlemanjerryCertificate()`

UnsetGentlemanjerryCertificate ensures that no value is present for GentlemanjerryCertificate, not even an explicit nil
### GetGentlemanjerryDockerName

`func (o *Account) GetGentlemanjerryDockerName() string`

GetGentlemanjerryDockerName returns the GentlemanjerryDockerName field if non-nil, zero value otherwise.

### GetGentlemanjerryDockerNameOk

`func (o *Account) GetGentlemanjerryDockerNameOk() (*string, bool)`

GetGentlemanjerryDockerNameOk returns a tuple with the GentlemanjerryDockerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryDockerName

`func (o *Account) SetGentlemanjerryDockerName(v string)`

SetGentlemanjerryDockerName sets GentlemanjerryDockerName field to given value.


### SetGentlemanjerryDockerNameNil

`func (o *Account) SetGentlemanjerryDockerNameNil(b bool)`

 SetGentlemanjerryDockerNameNil sets the value for GentlemanjerryDockerName to be an explicit nil

### UnsetGentlemanjerryDockerName
`func (o *Account) UnsetGentlemanjerryDockerName()`

UnsetGentlemanjerryDockerName ensures that no value is present for GentlemanjerryDockerName, not even an explicit nil
### GetGentlemanjerryInstanceId

`func (o *Account) GetGentlemanjerryInstanceId() string`

GetGentlemanjerryInstanceId returns the GentlemanjerryInstanceId field if non-nil, zero value otherwise.

### GetGentlemanjerryInstanceIdOk

`func (o *Account) GetGentlemanjerryInstanceIdOk() (*string, bool)`

GetGentlemanjerryInstanceIdOk returns a tuple with the GentlemanjerryInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryInstanceId

`func (o *Account) SetGentlemanjerryInstanceId(v string)`

SetGentlemanjerryInstanceId sets GentlemanjerryInstanceId field to given value.


### SetGentlemanjerryInstanceIdNil

`func (o *Account) SetGentlemanjerryInstanceIdNil(b bool)`

 SetGentlemanjerryInstanceIdNil sets the value for GentlemanjerryInstanceId to be an explicit nil

### UnsetGentlemanjerryInstanceId
`func (o *Account) UnsetGentlemanjerryInstanceId()`

UnsetGentlemanjerryInstanceId ensures that no value is present for GentlemanjerryInstanceId, not even an explicit nil
### GetSweetnessStackVersion

`func (o *Account) GetSweetnessStackVersion() string`

GetSweetnessStackVersion returns the SweetnessStackVersion field if non-nil, zero value otherwise.

### GetSweetnessStackVersionOk

`func (o *Account) GetSweetnessStackVersionOk() (*string, bool)`

GetSweetnessStackVersionOk returns a tuple with the SweetnessStackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweetnessStackVersion

`func (o *Account) SetSweetnessStackVersion(v string)`

SetSweetnessStackVersion sets SweetnessStackVersion field to given value.


### SetSweetnessStackVersionNil

`func (o *Account) SetSweetnessStackVersionNil(b bool)`

 SetSweetnessStackVersionNil sets the value for SweetnessStackVersion to be an explicit nil

### UnsetSweetnessStackVersion
`func (o *Account) UnsetSweetnessStackVersion()`

UnsetSweetnessStackVersion ensures that no value is present for SweetnessStackVersion, not even an explicit nil
### GetCaBody

`func (o *Account) GetCaBody() string`

GetCaBody returns the CaBody field if non-nil, zero value otherwise.

### GetCaBodyOk

`func (o *Account) GetCaBodyOk() (*string, bool)`

GetCaBodyOk returns a tuple with the CaBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBody

`func (o *Account) SetCaBody(v string)`

SetCaBody sets CaBody field to given value.


### SetCaBodyNil

`func (o *Account) SetCaBodyNil(b bool)`

 SetCaBodyNil sets the value for CaBody to be an explicit nil

### UnsetCaBody
`func (o *Account) UnsetCaBody()`

UnsetCaBody ensures that no value is present for CaBody, not even an explicit nil
### GetCaPrivateKey

`func (o *Account) GetCaPrivateKey() string`

GetCaPrivateKey returns the CaPrivateKey field if non-nil, zero value otherwise.

### GetCaPrivateKeyOk

`func (o *Account) GetCaPrivateKeyOk() (*string, bool)`

GetCaPrivateKeyOk returns a tuple with the CaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPrivateKey

`func (o *Account) SetCaPrivateKey(v string)`

SetCaPrivateKey sets CaPrivateKey field to given value.

### HasCaPrivateKey

`func (o *Account) HasCaPrivateKey() bool`

HasCaPrivateKey returns a boolean if a field has been set.

### SetCaPrivateKeyNil

`func (o *Account) SetCaPrivateKeyNil(b bool)`

 SetCaPrivateKeyNil sets the value for CaPrivateKey to be an explicit nil

### UnsetCaPrivateKey
`func (o *Account) UnsetCaPrivateKey()`

UnsetCaPrivateKey ensures that no value is present for CaPrivateKey, not even an explicit nil
### GetEmbedded

`func (o *Account) GetEmbedded() AccountEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Account) GetEmbeddedOk() (*AccountEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Account) SetEmbedded(v AccountEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *Account) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Account) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Account) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Account) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


