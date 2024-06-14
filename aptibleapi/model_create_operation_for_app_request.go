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

// checks if the CreateOperationForAppRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOperationForAppRequest{}

// CreateOperationForAppRequest struct for CreateOperationForAppRequest
type CreateOperationForAppRequest struct {
	Type string `json:"type"`
	GitRef *string `json:"git_ref,omitempty"`
	DiskSize *int32 `json:"disk_size,omitempty"`
	ContainerCount *int32 `json:"container_count,omitempty"`
	ContainerSize *int32 `json:"container_size,omitempty"`
	Command *string `json:"command,omitempty"`
	Env map[string]interface{} `json:"env,omitempty"`
	Handle *string `json:"handle,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
	DestinationRegion *string `json:"destination_region,omitempty"`
	Interactive *bool `json:"interactive,omitempty"`
	DestinationAccount *string `json:"destination_account,omitempty"`
	DestinationAccountId *int32 `json:"destination_account_id,omitempty"`
	DockerRef *string `json:"docker_ref,omitempty"`
	Automated *bool `json:"automated,omitempty"`
	KeyArn *string `json:"key_arn,omitempty"`
	InstanceProfile *string `json:"instance_profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOperationForAppRequest CreateOperationForAppRequest

// NewCreateOperationForAppRequest instantiates a new CreateOperationForAppRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOperationForAppRequest(type_ string) *CreateOperationForAppRequest {
	this := CreateOperationForAppRequest{}
	this.Type = type_
	return &this
}

// NewCreateOperationForAppRequestWithDefaults instantiates a new CreateOperationForAppRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOperationForAppRequestWithDefaults() *CreateOperationForAppRequest {
	this := CreateOperationForAppRequest{}
	return &this
}

// GetType returns the Type field value
func (o *CreateOperationForAppRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOperationForAppRequest) SetType(v string) {
	o.Type = v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetGitRef() string {
	if o == nil || IsNil(o.GitRef) {
		var ret string
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetGitRefOk() (*string, bool) {
	if o == nil || IsNil(o.GitRef) {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasGitRef() bool {
	if o != nil && !IsNil(o.GitRef) {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given string and assigns it to the GitRef field.
func (o *CreateOperationForAppRequest) SetGitRef(v string) {
	o.GitRef = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetDiskSize() int32 {
	if o == nil || IsNil(o.DiskSize) {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *CreateOperationForAppRequest) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetContainerCount returns the ContainerCount field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetContainerCount() int32 {
	if o == nil || IsNil(o.ContainerCount) {
		var ret int32
		return ret
	}
	return *o.ContainerCount
}

// GetContainerCountOk returns a tuple with the ContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetContainerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ContainerCount) {
		return nil, false
	}
	return o.ContainerCount, true
}

// HasContainerCount returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasContainerCount() bool {
	if o != nil && !IsNil(o.ContainerCount) {
		return true
	}

	return false
}

// SetContainerCount gets a reference to the given int32 and assigns it to the ContainerCount field.
func (o *CreateOperationForAppRequest) SetContainerCount(v int32) {
	o.ContainerCount = &v
}

// GetContainerSize returns the ContainerSize field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetContainerSize() int32 {
	if o == nil || IsNil(o.ContainerSize) {
		var ret int32
		return ret
	}
	return *o.ContainerSize
}

// GetContainerSizeOk returns a tuple with the ContainerSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetContainerSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ContainerSize) {
		return nil, false
	}
	return o.ContainerSize, true
}

// HasContainerSize returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasContainerSize() bool {
	if o != nil && !IsNil(o.ContainerSize) {
		return true
	}

	return false
}

// SetContainerSize gets a reference to the given int32 and assigns it to the ContainerSize field.
func (o *CreateOperationForAppRequest) SetContainerSize(v int32) {
	o.ContainerSize = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *CreateOperationForAppRequest) SetCommand(v string) {
	o.Command = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetEnv() map[string]interface{} {
	if o == nil || IsNil(o.Env) {
		var ret map[string]interface{}
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetEnvOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Env) {
		return map[string]interface{}{}, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given map[string]interface{} and assigns it to the Env field.
func (o *CreateOperationForAppRequest) SetEnv(v map[string]interface{}) {
	o.Env = v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetHandle() string {
	if o == nil || IsNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetHandleOk() (*string, bool) {
	if o == nil || IsNil(o.Handle) {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasHandle() bool {
	if o != nil && !IsNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *CreateOperationForAppRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CreateOperationForAppRequest) SetCertificate(v string) {
	o.Certificate = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *CreateOperationForAppRequest) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetDestinationRegion returns the DestinationRegion field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetDestinationRegion() string {
	if o == nil || IsNil(o.DestinationRegion) {
		var ret string
		return ret
	}
	return *o.DestinationRegion
}

// GetDestinationRegionOk returns a tuple with the DestinationRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetDestinationRegionOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationRegion) {
		return nil, false
	}
	return o.DestinationRegion, true
}

// HasDestinationRegion returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasDestinationRegion() bool {
	if o != nil && !IsNil(o.DestinationRegion) {
		return true
	}

	return false
}

// SetDestinationRegion gets a reference to the given string and assigns it to the DestinationRegion field.
func (o *CreateOperationForAppRequest) SetDestinationRegion(v string) {
	o.DestinationRegion = &v
}

// GetInteractive returns the Interactive field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetInteractive() bool {
	if o == nil || IsNil(o.Interactive) {
		var ret bool
		return ret
	}
	return *o.Interactive
}

// GetInteractiveOk returns a tuple with the Interactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetInteractiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Interactive) {
		return nil, false
	}
	return o.Interactive, true
}

// HasInteractive returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasInteractive() bool {
	if o != nil && !IsNil(o.Interactive) {
		return true
	}

	return false
}

// SetInteractive gets a reference to the given bool and assigns it to the Interactive field.
func (o *CreateOperationForAppRequest) SetInteractive(v bool) {
	o.Interactive = &v
}

// GetDestinationAccount returns the DestinationAccount field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetDestinationAccount() string {
	if o == nil || IsNil(o.DestinationAccount) {
		var ret string
		return ret
	}
	return *o.DestinationAccount
}

// GetDestinationAccountOk returns a tuple with the DestinationAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetDestinationAccountOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationAccount) {
		return nil, false
	}
	return o.DestinationAccount, true
}

// HasDestinationAccount returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasDestinationAccount() bool {
	if o != nil && !IsNil(o.DestinationAccount) {
		return true
	}

	return false
}

// SetDestinationAccount gets a reference to the given string and assigns it to the DestinationAccount field.
func (o *CreateOperationForAppRequest) SetDestinationAccount(v string) {
	o.DestinationAccount = &v
}

// GetDestinationAccountId returns the DestinationAccountId field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetDestinationAccountId() int32 {
	if o == nil || IsNil(o.DestinationAccountId) {
		var ret int32
		return ret
	}
	return *o.DestinationAccountId
}

// GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetDestinationAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DestinationAccountId) {
		return nil, false
	}
	return o.DestinationAccountId, true
}

// HasDestinationAccountId returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasDestinationAccountId() bool {
	if o != nil && !IsNil(o.DestinationAccountId) {
		return true
	}

	return false
}

// SetDestinationAccountId gets a reference to the given int32 and assigns it to the DestinationAccountId field.
func (o *CreateOperationForAppRequest) SetDestinationAccountId(v int32) {
	o.DestinationAccountId = &v
}

// GetDockerRef returns the DockerRef field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetDockerRef() string {
	if o == nil || IsNil(o.DockerRef) {
		var ret string
		return ret
	}
	return *o.DockerRef
}

// GetDockerRefOk returns a tuple with the DockerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetDockerRefOk() (*string, bool) {
	if o == nil || IsNil(o.DockerRef) {
		return nil, false
	}
	return o.DockerRef, true
}

// HasDockerRef returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasDockerRef() bool {
	if o != nil && !IsNil(o.DockerRef) {
		return true
	}

	return false
}

// SetDockerRef gets a reference to the given string and assigns it to the DockerRef field.
func (o *CreateOperationForAppRequest) SetDockerRef(v string) {
	o.DockerRef = &v
}

// GetAutomated returns the Automated field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetAutomated() bool {
	if o == nil || IsNil(o.Automated) {
		var ret bool
		return ret
	}
	return *o.Automated
}

// GetAutomatedOk returns a tuple with the Automated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetAutomatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Automated) {
		return nil, false
	}
	return o.Automated, true
}

// HasAutomated returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasAutomated() bool {
	if o != nil && !IsNil(o.Automated) {
		return true
	}

	return false
}

// SetAutomated gets a reference to the given bool and assigns it to the Automated field.
func (o *CreateOperationForAppRequest) SetAutomated(v bool) {
	o.Automated = &v
}

// GetKeyArn returns the KeyArn field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetKeyArn() string {
	if o == nil || IsNil(o.KeyArn) {
		var ret string
		return ret
	}
	return *o.KeyArn
}

// GetKeyArnOk returns a tuple with the KeyArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetKeyArnOk() (*string, bool) {
	if o == nil || IsNil(o.KeyArn) {
		return nil, false
	}
	return o.KeyArn, true
}

// HasKeyArn returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasKeyArn() bool {
	if o != nil && !IsNil(o.KeyArn) {
		return true
	}

	return false
}

// SetKeyArn gets a reference to the given string and assigns it to the KeyArn field.
func (o *CreateOperationForAppRequest) SetKeyArn(v string) {
	o.KeyArn = &v
}

// GetInstanceProfile returns the InstanceProfile field value if set, zero value otherwise.
func (o *CreateOperationForAppRequest) GetInstanceProfile() string {
	if o == nil || IsNil(o.InstanceProfile) {
		var ret string
		return ret
	}
	return *o.InstanceProfile
}

// GetInstanceProfileOk returns a tuple with the InstanceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOperationForAppRequest) GetInstanceProfileOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceProfile) {
		return nil, false
	}
	return o.InstanceProfile, true
}

// HasInstanceProfile returns a boolean if a field has been set.
func (o *CreateOperationForAppRequest) HasInstanceProfile() bool {
	if o != nil && !IsNil(o.InstanceProfile) {
		return true
	}

	return false
}

// SetInstanceProfile gets a reference to the given string and assigns it to the InstanceProfile field.
func (o *CreateOperationForAppRequest) SetInstanceProfile(v string) {
	o.InstanceProfile = &v
}

func (o CreateOperationForAppRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOperationForAppRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.GitRef) {
		toSerialize["git_ref"] = o.GitRef
	}
	if !IsNil(o.DiskSize) {
		toSerialize["disk_size"] = o.DiskSize
	}
	if !IsNil(o.ContainerCount) {
		toSerialize["container_count"] = o.ContainerCount
	}
	if !IsNil(o.ContainerSize) {
		toSerialize["container_size"] = o.ContainerSize
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["private_key"] = o.PrivateKey
	}
	if !IsNil(o.DestinationRegion) {
		toSerialize["destination_region"] = o.DestinationRegion
	}
	if !IsNil(o.Interactive) {
		toSerialize["interactive"] = o.Interactive
	}
	if !IsNil(o.DestinationAccount) {
		toSerialize["destination_account"] = o.DestinationAccount
	}
	if !IsNil(o.DestinationAccountId) {
		toSerialize["destination_account_id"] = o.DestinationAccountId
	}
	if !IsNil(o.DockerRef) {
		toSerialize["docker_ref"] = o.DockerRef
	}
	if !IsNil(o.Automated) {
		toSerialize["automated"] = o.Automated
	}
	if !IsNil(o.KeyArn) {
		toSerialize["key_arn"] = o.KeyArn
	}
	if !IsNil(o.InstanceProfile) {
		toSerialize["instance_profile"] = o.InstanceProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOperationForAppRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCreateOperationForAppRequest := _CreateOperationForAppRequest{}

	err = json.Unmarshal(data, &varCreateOperationForAppRequest)

	if err != nil {
		return err
	}

	*o = CreateOperationForAppRequest(varCreateOperationForAppRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "git_ref")
		delete(additionalProperties, "disk_size")
		delete(additionalProperties, "container_count")
		delete(additionalProperties, "container_size")
		delete(additionalProperties, "command")
		delete(additionalProperties, "env")
		delete(additionalProperties, "handle")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "private_key")
		delete(additionalProperties, "destination_region")
		delete(additionalProperties, "interactive")
		delete(additionalProperties, "destination_account")
		delete(additionalProperties, "destination_account_id")
		delete(additionalProperties, "docker_ref")
		delete(additionalProperties, "automated")
		delete(additionalProperties, "key_arn")
		delete(additionalProperties, "instance_profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOperationForAppRequest struct {
	value *CreateOperationForAppRequest
	isSet bool
}

func (v NullableCreateOperationForAppRequest) Get() *CreateOperationForAppRequest {
	return v.value
}

func (v *NullableCreateOperationForAppRequest) Set(val *CreateOperationForAppRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOperationForAppRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOperationForAppRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOperationForAppRequest(val *CreateOperationForAppRequest) *NullableCreateOperationForAppRequest {
	return &NullableCreateOperationForAppRequest{value: val, isSet: true}
}

func (v NullableCreateOperationForAppRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOperationForAppRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

