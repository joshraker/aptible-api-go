# CreateBackupRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daily** | Pointer to **int32** |  | [optional] 
**Monthly** | Pointer to **int32** |  | [optional] 
**Yearly** | Pointer to **int32** |  | [optional] 
**KeepFinal** | Pointer to **bool** |  | [optional] 
**MakeCopy** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateBackupRetentionPolicyRequest

`func NewCreateBackupRetentionPolicyRequest() *CreateBackupRetentionPolicyRequest`

NewCreateBackupRetentionPolicyRequest instantiates a new CreateBackupRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupRetentionPolicyRequestWithDefaults

`func NewCreateBackupRetentionPolicyRequestWithDefaults() *CreateBackupRetentionPolicyRequest`

NewCreateBackupRetentionPolicyRequestWithDefaults instantiates a new CreateBackupRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaily

`func (o *CreateBackupRetentionPolicyRequest) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *CreateBackupRetentionPolicyRequest) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *CreateBackupRetentionPolicyRequest) SetDaily(v int32)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *CreateBackupRetentionPolicyRequest) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetMonthly

`func (o *CreateBackupRetentionPolicyRequest) GetMonthly() int32`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *CreateBackupRetentionPolicyRequest) GetMonthlyOk() (*int32, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *CreateBackupRetentionPolicyRequest) SetMonthly(v int32)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *CreateBackupRetentionPolicyRequest) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetYearly

`func (o *CreateBackupRetentionPolicyRequest) GetYearly() int32`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *CreateBackupRetentionPolicyRequest) GetYearlyOk() (*int32, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *CreateBackupRetentionPolicyRequest) SetYearly(v int32)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *CreateBackupRetentionPolicyRequest) HasYearly() bool`

HasYearly returns a boolean if a field has been set.

### GetKeepFinal

`func (o *CreateBackupRetentionPolicyRequest) GetKeepFinal() bool`

GetKeepFinal returns the KeepFinal field if non-nil, zero value otherwise.

### GetKeepFinalOk

`func (o *CreateBackupRetentionPolicyRequest) GetKeepFinalOk() (*bool, bool)`

GetKeepFinalOk returns a tuple with the KeepFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFinal

`func (o *CreateBackupRetentionPolicyRequest) SetKeepFinal(v bool)`

SetKeepFinal sets KeepFinal field to given value.

### HasKeepFinal

`func (o *CreateBackupRetentionPolicyRequest) HasKeepFinal() bool`

HasKeepFinal returns a boolean if a field has been set.

### GetMakeCopy

`func (o *CreateBackupRetentionPolicyRequest) GetMakeCopy() bool`

GetMakeCopy returns the MakeCopy field if non-nil, zero value otherwise.

### GetMakeCopyOk

`func (o *CreateBackupRetentionPolicyRequest) GetMakeCopyOk() (*bool, bool)`

GetMakeCopyOk returns a tuple with the MakeCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCopy

`func (o *CreateBackupRetentionPolicyRequest) SetMakeCopy(v bool)`

SetMakeCopy sets MakeCopy field to given value.

### HasMakeCopy

`func (o *CreateBackupRetentionPolicyRequest) HasMakeCopy() bool`

HasMakeCopy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


