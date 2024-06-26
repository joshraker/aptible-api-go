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

// checks if the AppEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEmbedded{}

// AppEmbedded struct for AppEmbedded
type AppEmbedded struct {
	Services []Service `json:"services,omitempty"`
	LastOperation *Operation `json:"last_operation,omitempty"`
	LastDeployOperation *Operation `json:"last_deploy_operation,omitempty"`
	CurrentImage *Image `json:"current_image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppEmbedded AppEmbedded

// NewAppEmbedded instantiates a new AppEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEmbedded() *AppEmbedded {
	this := AppEmbedded{}
	return &this
}

// NewAppEmbeddedWithDefaults instantiates a new AppEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEmbeddedWithDefaults() *AppEmbedded {
	this := AppEmbedded{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AppEmbedded) GetServices() []Service {
	if o == nil || IsNil(o.Services) {
		var ret []Service
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEmbedded) GetServicesOk() ([]Service, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AppEmbedded) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []Service and assigns it to the Services field.
func (o *AppEmbedded) SetServices(v []Service) {
	o.Services = v
}

// GetLastOperation returns the LastOperation field value if set, zero value otherwise.
func (o *AppEmbedded) GetLastOperation() Operation {
	if o == nil || IsNil(o.LastOperation) {
		var ret Operation
		return ret
	}
	return *o.LastOperation
}

// GetLastOperationOk returns a tuple with the LastOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEmbedded) GetLastOperationOk() (*Operation, bool) {
	if o == nil || IsNil(o.LastOperation) {
		return nil, false
	}
	return o.LastOperation, true
}

// HasLastOperation returns a boolean if a field has been set.
func (o *AppEmbedded) HasLastOperation() bool {
	if o != nil && !IsNil(o.LastOperation) {
		return true
	}

	return false
}

// SetLastOperation gets a reference to the given Operation and assigns it to the LastOperation field.
func (o *AppEmbedded) SetLastOperation(v Operation) {
	o.LastOperation = &v
}

// GetLastDeployOperation returns the LastDeployOperation field value if set, zero value otherwise.
func (o *AppEmbedded) GetLastDeployOperation() Operation {
	if o == nil || IsNil(o.LastDeployOperation) {
		var ret Operation
		return ret
	}
	return *o.LastDeployOperation
}

// GetLastDeployOperationOk returns a tuple with the LastDeployOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEmbedded) GetLastDeployOperationOk() (*Operation, bool) {
	if o == nil || IsNil(o.LastDeployOperation) {
		return nil, false
	}
	return o.LastDeployOperation, true
}

// HasLastDeployOperation returns a boolean if a field has been set.
func (o *AppEmbedded) HasLastDeployOperation() bool {
	if o != nil && !IsNil(o.LastDeployOperation) {
		return true
	}

	return false
}

// SetLastDeployOperation gets a reference to the given Operation and assigns it to the LastDeployOperation field.
func (o *AppEmbedded) SetLastDeployOperation(v Operation) {
	o.LastDeployOperation = &v
}

// GetCurrentImage returns the CurrentImage field value if set, zero value otherwise.
func (o *AppEmbedded) GetCurrentImage() Image {
	if o == nil || IsNil(o.CurrentImage) {
		var ret Image
		return ret
	}
	return *o.CurrentImage
}

// GetCurrentImageOk returns a tuple with the CurrentImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEmbedded) GetCurrentImageOk() (*Image, bool) {
	if o == nil || IsNil(o.CurrentImage) {
		return nil, false
	}
	return o.CurrentImage, true
}

// HasCurrentImage returns a boolean if a field has been set.
func (o *AppEmbedded) HasCurrentImage() bool {
	if o != nil && !IsNil(o.CurrentImage) {
		return true
	}

	return false
}

// SetCurrentImage gets a reference to the given Image and assigns it to the CurrentImage field.
func (o *AppEmbedded) SetCurrentImage(v Image) {
	o.CurrentImage = &v
}

func (o AppEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.LastOperation) {
		toSerialize["last_operation"] = o.LastOperation
	}
	if !IsNil(o.LastDeployOperation) {
		toSerialize["last_deploy_operation"] = o.LastDeployOperation
	}
	if !IsNil(o.CurrentImage) {
		toSerialize["current_image"] = o.CurrentImage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppEmbedded) UnmarshalJSON(data []byte) (err error) {
	varAppEmbedded := _AppEmbedded{}

	err = json.Unmarshal(data, &varAppEmbedded)

	if err != nil {
		return err
	}

	*o = AppEmbedded(varAppEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "services")
		delete(additionalProperties, "last_operation")
		delete(additionalProperties, "last_deploy_operation")
		delete(additionalProperties, "current_image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppEmbedded struct {
	value *AppEmbedded
	isSet bool
}

func (v NullableAppEmbedded) Get() *AppEmbedded {
	return v.value
}

func (v *NullableAppEmbedded) Set(val *AppEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEmbedded(val *AppEmbedded) *NullableAppEmbedded {
	return &NullableAppEmbedded{value: val, isSet: true}
}

func (v NullableAppEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


