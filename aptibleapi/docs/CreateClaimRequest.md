# CreateClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Handle** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateClaimRequest

`func NewCreateClaimRequest(type_ string, ) *CreateClaimRequest`

NewCreateClaimRequest instantiates a new CreateClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClaimRequestWithDefaults

`func NewCreateClaimRequestWithDefaults() *CreateClaimRequest`

NewCreateClaimRequestWithDefaults instantiates a new CreateClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateClaimRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateClaimRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateClaimRequest) SetType(v string)`

SetType sets Type field to given value.


### GetHandle

`func (o *CreateClaimRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateClaimRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateClaimRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CreateClaimRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


