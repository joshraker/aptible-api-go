/*
Deploy API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aptibleapi

import (
	"encoding/json"
)

// checks if the ClaimsTypePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimsTypePostRequest{}

// ClaimsTypePostRequest struct for ClaimsTypePostRequest
type ClaimsTypePostRequest struct {
	Handle *string `json:"handle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClaimsTypePostRequest ClaimsTypePostRequest

// NewClaimsTypePostRequest instantiates a new ClaimsTypePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimsTypePostRequest() *ClaimsTypePostRequest {
	this := ClaimsTypePostRequest{}
	return &this
}

// NewClaimsTypePostRequestWithDefaults instantiates a new ClaimsTypePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimsTypePostRequestWithDefaults() *ClaimsTypePostRequest {
	this := ClaimsTypePostRequest{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *ClaimsTypePostRequest) GetHandle() string {
	if o == nil || IsNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimsTypePostRequest) GetHandleOk() (*string, bool) {
	if o == nil || IsNil(o.Handle) {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *ClaimsTypePostRequest) HasHandle() bool {
	if o != nil && !IsNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *ClaimsTypePostRequest) SetHandle(v string) {
	o.Handle = &v
}

func (o ClaimsTypePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimsTypePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClaimsTypePostRequest) UnmarshalJSON(data []byte) (err error) {
	varClaimsTypePostRequest := _ClaimsTypePostRequest{}

	err = json.Unmarshal(data, &varClaimsTypePostRequest)

	if err != nil {
		return err
	}

	*o = ClaimsTypePostRequest(varClaimsTypePostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "handle")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClaimsTypePostRequest struct {
	value *ClaimsTypePostRequest
	isSet bool
}

func (v NullableClaimsTypePostRequest) Get() *ClaimsTypePostRequest {
	return v.value
}

func (v *NullableClaimsTypePostRequest) Set(val *ClaimsTypePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimsTypePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimsTypePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimsTypePostRequest(val *ClaimsTypePostRequest) *NullableClaimsTypePostRequest {
	return &NullableClaimsTypePostRequest{value: val, isSet: true}
}

func (v NullableClaimsTypePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimsTypePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

