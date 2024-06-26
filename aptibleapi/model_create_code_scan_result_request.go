/*
Deploy API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aptibleapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateCodeScanResultRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCodeScanResultRequest{}

// CreateCodeScanResultRequest struct for CreateCodeScanResultRequest
type CreateCodeScanResultRequest struct {
	// Alternate name for `operation_id`
	Operation *int32 `json:"operation,omitempty"`
	OperationId int32 `json:"operation_id"`
	DockerfilePresent *bool `json:"dockerfile_present,omitempty"`
	ProcfilePresent *bool `json:"procfile_present,omitempty"`
	AptibleYmlPresent *bool `json:"aptible_yml_present,omitempty"`
	GitRef string `json:"git_ref"`
	GitCommit string `json:"git_commit"`
	AdditionalProperties map[string]interface{}
}

type _CreateCodeScanResultRequest CreateCodeScanResultRequest

// NewCreateCodeScanResultRequest instantiates a new CreateCodeScanResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCodeScanResultRequest(operationId int32, gitRef string, gitCommit string) *CreateCodeScanResultRequest {
	this := CreateCodeScanResultRequest{}
	this.OperationId = operationId
	this.GitRef = gitRef
	this.GitCommit = gitCommit
	return &this
}

// NewCreateCodeScanResultRequestWithDefaults instantiates a new CreateCodeScanResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCodeScanResultRequestWithDefaults() *CreateCodeScanResultRequest {
	this := CreateCodeScanResultRequest{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *CreateCodeScanResultRequest) GetOperation() int32 {
	if o == nil || IsNil(o.Operation) {
		var ret int32
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetOperationOk() (*int32, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *CreateCodeScanResultRequest) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given int32 and assigns it to the Operation field.
func (o *CreateCodeScanResultRequest) SetOperation(v int32) {
	o.Operation = &v
}

// GetOperationId returns the OperationId field value
func (o *CreateCodeScanResultRequest) GetOperationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetOperationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *CreateCodeScanResultRequest) SetOperationId(v int32) {
	o.OperationId = v
}

// GetDockerfilePresent returns the DockerfilePresent field value if set, zero value otherwise.
func (o *CreateCodeScanResultRequest) GetDockerfilePresent() bool {
	if o == nil || IsNil(o.DockerfilePresent) {
		var ret bool
		return ret
	}
	return *o.DockerfilePresent
}

// GetDockerfilePresentOk returns a tuple with the DockerfilePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetDockerfilePresentOk() (*bool, bool) {
	if o == nil || IsNil(o.DockerfilePresent) {
		return nil, false
	}
	return o.DockerfilePresent, true
}

// HasDockerfilePresent returns a boolean if a field has been set.
func (o *CreateCodeScanResultRequest) HasDockerfilePresent() bool {
	if o != nil && !IsNil(o.DockerfilePresent) {
		return true
	}

	return false
}

// SetDockerfilePresent gets a reference to the given bool and assigns it to the DockerfilePresent field.
func (o *CreateCodeScanResultRequest) SetDockerfilePresent(v bool) {
	o.DockerfilePresent = &v
}

// GetProcfilePresent returns the ProcfilePresent field value if set, zero value otherwise.
func (o *CreateCodeScanResultRequest) GetProcfilePresent() bool {
	if o == nil || IsNil(o.ProcfilePresent) {
		var ret bool
		return ret
	}
	return *o.ProcfilePresent
}

// GetProcfilePresentOk returns a tuple with the ProcfilePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetProcfilePresentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProcfilePresent) {
		return nil, false
	}
	return o.ProcfilePresent, true
}

// HasProcfilePresent returns a boolean if a field has been set.
func (o *CreateCodeScanResultRequest) HasProcfilePresent() bool {
	if o != nil && !IsNil(o.ProcfilePresent) {
		return true
	}

	return false
}

// SetProcfilePresent gets a reference to the given bool and assigns it to the ProcfilePresent field.
func (o *CreateCodeScanResultRequest) SetProcfilePresent(v bool) {
	o.ProcfilePresent = &v
}

// GetAptibleYmlPresent returns the AptibleYmlPresent field value if set, zero value otherwise.
func (o *CreateCodeScanResultRequest) GetAptibleYmlPresent() bool {
	if o == nil || IsNil(o.AptibleYmlPresent) {
		var ret bool
		return ret
	}
	return *o.AptibleYmlPresent
}

// GetAptibleYmlPresentOk returns a tuple with the AptibleYmlPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetAptibleYmlPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.AptibleYmlPresent) {
		return nil, false
	}
	return o.AptibleYmlPresent, true
}

// HasAptibleYmlPresent returns a boolean if a field has been set.
func (o *CreateCodeScanResultRequest) HasAptibleYmlPresent() bool {
	if o != nil && !IsNil(o.AptibleYmlPresent) {
		return true
	}

	return false
}

// SetAptibleYmlPresent gets a reference to the given bool and assigns it to the AptibleYmlPresent field.
func (o *CreateCodeScanResultRequest) SetAptibleYmlPresent(v bool) {
	o.AptibleYmlPresent = &v
}

// GetGitRef returns the GitRef field value
func (o *CreateCodeScanResultRequest) GetGitRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetGitRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitRef, true
}

// SetGitRef sets field value
func (o *CreateCodeScanResultRequest) SetGitRef(v string) {
	o.GitRef = v
}

// GetGitCommit returns the GitCommit field value
func (o *CreateCodeScanResultRequest) GetGitCommit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitCommit
}

// GetGitCommitOk returns a tuple with the GitCommit field value
// and a boolean to check if the value has been set.
func (o *CreateCodeScanResultRequest) GetGitCommitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitCommit, true
}

// SetGitCommit sets field value
func (o *CreateCodeScanResultRequest) SetGitCommit(v string) {
	o.GitCommit = v
}

func (o CreateCodeScanResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCodeScanResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	toSerialize["operation_id"] = o.OperationId
	if !IsNil(o.DockerfilePresent) {
		toSerialize["dockerfile_present"] = o.DockerfilePresent
	}
	if !IsNil(o.ProcfilePresent) {
		toSerialize["procfile_present"] = o.ProcfilePresent
	}
	if !IsNil(o.AptibleYmlPresent) {
		toSerialize["aptible_yml_present"] = o.AptibleYmlPresent
	}
	toSerialize["git_ref"] = o.GitRef
	toSerialize["git_commit"] = o.GitCommit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCodeScanResultRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operation_id",
		"git_ref",
		"git_commit",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateCodeScanResultRequest := _CreateCodeScanResultRequest{}

	err = json.Unmarshal(data, &varCreateCodeScanResultRequest)

	if err != nil {
		return err
	}

	*o = CreateCodeScanResultRequest(varCreateCodeScanResultRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "operation_id")
		delete(additionalProperties, "dockerfile_present")
		delete(additionalProperties, "procfile_present")
		delete(additionalProperties, "aptible_yml_present")
		delete(additionalProperties, "git_ref")
		delete(additionalProperties, "git_commit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCodeScanResultRequest struct {
	value *CreateCodeScanResultRequest
	isSet bool
}

func (v NullableCreateCodeScanResultRequest) Get() *CreateCodeScanResultRequest {
	return v.value
}

func (v *NullableCreateCodeScanResultRequest) Set(val *CreateCodeScanResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCodeScanResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCodeScanResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCodeScanResultRequest(val *CreateCodeScanResultRequest) *NullableCreateCodeScanResultRequest {
	return &NullableCreateCodeScanResultRequest{value: val, isSet: true}
}

func (v NullableCreateCodeScanResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCodeScanResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


