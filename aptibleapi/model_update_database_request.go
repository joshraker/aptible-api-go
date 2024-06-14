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

// checks if the UpdateDatabaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDatabaseRequest{}

// UpdateDatabaseRequest struct for UpdateDatabaseRequest
type UpdateDatabaseRequest struct {
	CurrentConfiguration *int32 `json:"current_configuration,omitempty"`
	Handle *string `json:"handle,omitempty"`
	InitialDiskSize *int32 `json:"initial_disk_size,omitempty"`
	InitialContinerSize *int32 `json:"initial_continer_size,omitempty"`
	DatabaseImageId *int32 `json:"database_image_id,omitempty"`
	CurrentKmsArn *int32 `json:"current_kms_arn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDatabaseRequest UpdateDatabaseRequest

// NewUpdateDatabaseRequest instantiates a new UpdateDatabaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDatabaseRequest() *UpdateDatabaseRequest {
	this := UpdateDatabaseRequest{}
	return &this
}

// NewUpdateDatabaseRequestWithDefaults instantiates a new UpdateDatabaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDatabaseRequestWithDefaults() *UpdateDatabaseRequest {
	this := UpdateDatabaseRequest{}
	return &this
}

// GetCurrentConfiguration returns the CurrentConfiguration field value if set, zero value otherwise.
func (o *UpdateDatabaseRequest) GetCurrentConfiguration() int32 {
	if o == nil || IsNil(o.CurrentConfiguration) {
		var ret int32
		return ret
	}
	return *o.CurrentConfiguration
}

// GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseRequest) GetCurrentConfigurationOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentConfiguration) {
		return nil, false
	}
	return o.CurrentConfiguration, true
}

// HasCurrentConfiguration returns a boolean if a field has been set.
func (o *UpdateDatabaseRequest) HasCurrentConfiguration() bool {
	if o != nil && !IsNil(o.CurrentConfiguration) {
		return true
	}

	return false
}

// SetCurrentConfiguration gets a reference to the given int32 and assigns it to the CurrentConfiguration field.
func (o *UpdateDatabaseRequest) SetCurrentConfiguration(v int32) {
	o.CurrentConfiguration = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UpdateDatabaseRequest) GetHandle() string {
	if o == nil || IsNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseRequest) GetHandleOk() (*string, bool) {
	if o == nil || IsNil(o.Handle) {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UpdateDatabaseRequest) HasHandle() bool {
	if o != nil && !IsNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UpdateDatabaseRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetInitialDiskSize returns the InitialDiskSize field value if set, zero value otherwise.
func (o *UpdateDatabaseRequest) GetInitialDiskSize() int32 {
	if o == nil || IsNil(o.InitialDiskSize) {
		var ret int32
		return ret
	}
	return *o.InitialDiskSize
}

// GetInitialDiskSizeOk returns a tuple with the InitialDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseRequest) GetInitialDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.InitialDiskSize) {
		return nil, false
	}
	return o.InitialDiskSize, true
}

// HasInitialDiskSize returns a boolean if a field has been set.
func (o *UpdateDatabaseRequest) HasInitialDiskSize() bool {
	if o != nil && !IsNil(o.InitialDiskSize) {
		return true
	}

	return false
}

// SetInitialDiskSize gets a reference to the given int32 and assigns it to the InitialDiskSize field.
func (o *UpdateDatabaseRequest) SetInitialDiskSize(v int32) {
	o.InitialDiskSize = &v
}

// GetInitialContinerSize returns the InitialContinerSize field value if set, zero value otherwise.
func (o *UpdateDatabaseRequest) GetInitialContinerSize() int32 {
	if o == nil || IsNil(o.InitialContinerSize) {
		var ret int32
		return ret
	}
	return *o.InitialContinerSize
}

// GetInitialContinerSizeOk returns a tuple with the InitialContinerSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseRequest) GetInitialContinerSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.InitialContinerSize) {
		return nil, false
	}
	return o.InitialContinerSize, true
}

// HasInitialContinerSize returns a boolean if a field has been set.
func (o *UpdateDatabaseRequest) HasInitialContinerSize() bool {
	if o != nil && !IsNil(o.InitialContinerSize) {
		return true
	}

	return false
}

// SetInitialContinerSize gets a reference to the given int32 and assigns it to the InitialContinerSize field.
func (o *UpdateDatabaseRequest) SetInitialContinerSize(v int32) {
	o.InitialContinerSize = &v
}

// GetDatabaseImageId returns the DatabaseImageId field value if set, zero value otherwise.
func (o *UpdateDatabaseRequest) GetDatabaseImageId() int32 {
	if o == nil || IsNil(o.DatabaseImageId) {
		var ret int32
		return ret
	}
	return *o.DatabaseImageId
}

// GetDatabaseImageIdOk returns a tuple with the DatabaseImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseRequest) GetDatabaseImageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DatabaseImageId) {
		return nil, false
	}
	return o.DatabaseImageId, true
}

// HasDatabaseImageId returns a boolean if a field has been set.
func (o *UpdateDatabaseRequest) HasDatabaseImageId() bool {
	if o != nil && !IsNil(o.DatabaseImageId) {
		return true
	}

	return false
}

// SetDatabaseImageId gets a reference to the given int32 and assigns it to the DatabaseImageId field.
func (o *UpdateDatabaseRequest) SetDatabaseImageId(v int32) {
	o.DatabaseImageId = &v
}

// GetCurrentKmsArn returns the CurrentKmsArn field value if set, zero value otherwise.
func (o *UpdateDatabaseRequest) GetCurrentKmsArn() int32 {
	if o == nil || IsNil(o.CurrentKmsArn) {
		var ret int32
		return ret
	}
	return *o.CurrentKmsArn
}

// GetCurrentKmsArnOk returns a tuple with the CurrentKmsArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseRequest) GetCurrentKmsArnOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentKmsArn) {
		return nil, false
	}
	return o.CurrentKmsArn, true
}

// HasCurrentKmsArn returns a boolean if a field has been set.
func (o *UpdateDatabaseRequest) HasCurrentKmsArn() bool {
	if o != nil && !IsNil(o.CurrentKmsArn) {
		return true
	}

	return false
}

// SetCurrentKmsArn gets a reference to the given int32 and assigns it to the CurrentKmsArn field.
func (o *UpdateDatabaseRequest) SetCurrentKmsArn(v int32) {
	o.CurrentKmsArn = &v
}

func (o UpdateDatabaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDatabaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentConfiguration) {
		toSerialize["current_configuration"] = o.CurrentConfiguration
	}
	if !IsNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !IsNil(o.InitialDiskSize) {
		toSerialize["initial_disk_size"] = o.InitialDiskSize
	}
	if !IsNil(o.InitialContinerSize) {
		toSerialize["initial_continer_size"] = o.InitialContinerSize
	}
	if !IsNil(o.DatabaseImageId) {
		toSerialize["database_image_id"] = o.DatabaseImageId
	}
	if !IsNil(o.CurrentKmsArn) {
		toSerialize["current_kms_arn"] = o.CurrentKmsArn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDatabaseRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateDatabaseRequest := _UpdateDatabaseRequest{}

	err = json.Unmarshal(data, &varUpdateDatabaseRequest)

	if err != nil {
		return err
	}

	*o = UpdateDatabaseRequest(varUpdateDatabaseRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current_configuration")
		delete(additionalProperties, "handle")
		delete(additionalProperties, "initial_disk_size")
		delete(additionalProperties, "initial_continer_size")
		delete(additionalProperties, "database_image_id")
		delete(additionalProperties, "current_kms_arn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDatabaseRequest struct {
	value *UpdateDatabaseRequest
	isSet bool
}

func (v NullableUpdateDatabaseRequest) Get() *UpdateDatabaseRequest {
	return v.value
}

func (v *NullableUpdateDatabaseRequest) Set(val *UpdateDatabaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDatabaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDatabaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDatabaseRequest(val *UpdateDatabaseRequest) *NullableUpdateDatabaseRequest {
	return &NullableUpdateDatabaseRequest{value: val, isSet: true}
}

func (v NullableUpdateDatabaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDatabaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

