# PatchAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchAccountRequest

`func NewPatchAccountRequest() *PatchAccountRequest`

NewPatchAccountRequest instantiates a new PatchAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountRequestWithDefaults

`func NewPatchAccountRequestWithDefaults() *PatchAccountRequest`

NewPatchAccountRequestWithDefaults instantiates a new PatchAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *PatchAccountRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *PatchAccountRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *PatchAccountRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *PatchAccountRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetName

`func (o *PatchAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


