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

// checks if the ListCodeScanResultsForApp200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCodeScanResultsForApp200ResponseEmbedded{}

// ListCodeScanResultsForApp200ResponseEmbedded struct for ListCodeScanResultsForApp200ResponseEmbedded
type ListCodeScanResultsForApp200ResponseEmbedded struct {
	CodeScanResults []CodeScanResult `json:"code_scan_results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListCodeScanResultsForApp200ResponseEmbedded ListCodeScanResultsForApp200ResponseEmbedded

// NewListCodeScanResultsForApp200ResponseEmbedded instantiates a new ListCodeScanResultsForApp200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCodeScanResultsForApp200ResponseEmbedded() *ListCodeScanResultsForApp200ResponseEmbedded {
	this := ListCodeScanResultsForApp200ResponseEmbedded{}
	return &this
}

// NewListCodeScanResultsForApp200ResponseEmbeddedWithDefaults instantiates a new ListCodeScanResultsForApp200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCodeScanResultsForApp200ResponseEmbeddedWithDefaults() *ListCodeScanResultsForApp200ResponseEmbedded {
	this := ListCodeScanResultsForApp200ResponseEmbedded{}
	return &this
}

// GetCodeScanResults returns the CodeScanResults field value if set, zero value otherwise.
func (o *ListCodeScanResultsForApp200ResponseEmbedded) GetCodeScanResults() []CodeScanResult {
	if o == nil || IsNil(o.CodeScanResults) {
		var ret []CodeScanResult
		return ret
	}
	return o.CodeScanResults
}

// GetCodeScanResultsOk returns a tuple with the CodeScanResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCodeScanResultsForApp200ResponseEmbedded) GetCodeScanResultsOk() ([]CodeScanResult, bool) {
	if o == nil || IsNil(o.CodeScanResults) {
		return nil, false
	}
	return o.CodeScanResults, true
}

// HasCodeScanResults returns a boolean if a field has been set.
func (o *ListCodeScanResultsForApp200ResponseEmbedded) HasCodeScanResults() bool {
	if o != nil && !IsNil(o.CodeScanResults) {
		return true
	}

	return false
}

// SetCodeScanResults gets a reference to the given []CodeScanResult and assigns it to the CodeScanResults field.
func (o *ListCodeScanResultsForApp200ResponseEmbedded) SetCodeScanResults(v []CodeScanResult) {
	o.CodeScanResults = v
}

func (o ListCodeScanResultsForApp200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCodeScanResultsForApp200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeScanResults) {
		toSerialize["code_scan_results"] = o.CodeScanResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCodeScanResultsForApp200ResponseEmbedded) UnmarshalJSON(data []byte) (err error) {
	varListCodeScanResultsForApp200ResponseEmbedded := _ListCodeScanResultsForApp200ResponseEmbedded{}

	err = json.Unmarshal(data, &varListCodeScanResultsForApp200ResponseEmbedded)

	if err != nil {
		return err
	}

	*o = ListCodeScanResultsForApp200ResponseEmbedded(varListCodeScanResultsForApp200ResponseEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code_scan_results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCodeScanResultsForApp200ResponseEmbedded struct {
	value *ListCodeScanResultsForApp200ResponseEmbedded
	isSet bool
}

func (v NullableListCodeScanResultsForApp200ResponseEmbedded) Get() *ListCodeScanResultsForApp200ResponseEmbedded {
	return v.value
}

func (v *NullableListCodeScanResultsForApp200ResponseEmbedded) Set(val *ListCodeScanResultsForApp200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListCodeScanResultsForApp200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListCodeScanResultsForApp200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCodeScanResultsForApp200ResponseEmbedded(val *ListCodeScanResultsForApp200ResponseEmbedded) *NullableListCodeScanResultsForApp200ResponseEmbedded {
	return &NullableListCodeScanResultsForApp200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListCodeScanResultsForApp200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCodeScanResultsForApp200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


