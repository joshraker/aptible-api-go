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

// checks if the OperationLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationLinks{}

// OperationLinks struct for OperationLinks
type OperationLinks struct {
	User *ListAccountsForStack200ResponseLinksStack `json:"user,omitempty"`
	Resource *ListAccountsForStack200ResponseLinksStack `json:"resource,omitempty"`
	Account *ListAccountsForStack200ResponseLinksStack `json:"account,omitempty"`
	DestinationAccount *ListAccountsForStack200ResponseLinksStack `json:"destination_account,omitempty"`
	SshPortalConnections *ListAccountsForStack200ResponseLinksStack `json:"ssh_portal_connections,omitempty"`
	EphemeralSessions *ListAccountsForStack200ResponseLinksStack `json:"ephemeral_sessions,omitempty"`
	CodeScanResult *ListAccountsForStack200ResponseLinksStack `json:"code_scan_result,omitempty"`
	Service *ListAccountsForStack200ResponseLinksStack `json:"service,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationLinks OperationLinks

// NewOperationLinks instantiates a new OperationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationLinks() *OperationLinks {
	this := OperationLinks{}
	return &this
}

// NewOperationLinksWithDefaults instantiates a new OperationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationLinksWithDefaults() *OperationLinks {
	this := OperationLinks{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OperationLinks) GetUser() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.User) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetUserOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OperationLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the User field.
func (o *OperationLinks) SetUser(v ListAccountsForStack200ResponseLinksStack) {
	o.User = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *OperationLinks) GetResource() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Resource) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetResourceOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *OperationLinks) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Resource field.
func (o *OperationLinks) SetResource(v ListAccountsForStack200ResponseLinksStack) {
	o.Resource = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *OperationLinks) GetAccount() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Account) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetAccountOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *OperationLinks) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Account field.
func (o *OperationLinks) SetAccount(v ListAccountsForStack200ResponseLinksStack) {
	o.Account = &v
}

// GetDestinationAccount returns the DestinationAccount field value if set, zero value otherwise.
func (o *OperationLinks) GetDestinationAccount() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.DestinationAccount) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.DestinationAccount
}

// GetDestinationAccountOk returns a tuple with the DestinationAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetDestinationAccountOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.DestinationAccount) {
		return nil, false
	}
	return o.DestinationAccount, true
}

// HasDestinationAccount returns a boolean if a field has been set.
func (o *OperationLinks) HasDestinationAccount() bool {
	if o != nil && !IsNil(o.DestinationAccount) {
		return true
	}

	return false
}

// SetDestinationAccount gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the DestinationAccount field.
func (o *OperationLinks) SetDestinationAccount(v ListAccountsForStack200ResponseLinksStack) {
	o.DestinationAccount = &v
}

// GetSshPortalConnections returns the SshPortalConnections field value if set, zero value otherwise.
func (o *OperationLinks) GetSshPortalConnections() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.SshPortalConnections) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.SshPortalConnections
}

// GetSshPortalConnectionsOk returns a tuple with the SshPortalConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetSshPortalConnectionsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.SshPortalConnections) {
		return nil, false
	}
	return o.SshPortalConnections, true
}

// HasSshPortalConnections returns a boolean if a field has been set.
func (o *OperationLinks) HasSshPortalConnections() bool {
	if o != nil && !IsNil(o.SshPortalConnections) {
		return true
	}

	return false
}

// SetSshPortalConnections gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the SshPortalConnections field.
func (o *OperationLinks) SetSshPortalConnections(v ListAccountsForStack200ResponseLinksStack) {
	o.SshPortalConnections = &v
}

// GetEphemeralSessions returns the EphemeralSessions field value if set, zero value otherwise.
func (o *OperationLinks) GetEphemeralSessions() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.EphemeralSessions) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.EphemeralSessions
}

// GetEphemeralSessionsOk returns a tuple with the EphemeralSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetEphemeralSessionsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.EphemeralSessions) {
		return nil, false
	}
	return o.EphemeralSessions, true
}

// HasEphemeralSessions returns a boolean if a field has been set.
func (o *OperationLinks) HasEphemeralSessions() bool {
	if o != nil && !IsNil(o.EphemeralSessions) {
		return true
	}

	return false
}

// SetEphemeralSessions gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the EphemeralSessions field.
func (o *OperationLinks) SetEphemeralSessions(v ListAccountsForStack200ResponseLinksStack) {
	o.EphemeralSessions = &v
}

// GetCodeScanResult returns the CodeScanResult field value if set, zero value otherwise.
func (o *OperationLinks) GetCodeScanResult() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.CodeScanResult) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.CodeScanResult
}

// GetCodeScanResultOk returns a tuple with the CodeScanResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetCodeScanResultOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.CodeScanResult) {
		return nil, false
	}
	return o.CodeScanResult, true
}

// HasCodeScanResult returns a boolean if a field has been set.
func (o *OperationLinks) HasCodeScanResult() bool {
	if o != nil && !IsNil(o.CodeScanResult) {
		return true
	}

	return false
}

// SetCodeScanResult gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the CodeScanResult field.
func (o *OperationLinks) SetCodeScanResult(v ListAccountsForStack200ResponseLinksStack) {
	o.CodeScanResult = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *OperationLinks) GetService() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Service) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetServiceOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *OperationLinks) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Service field.
func (o *OperationLinks) SetService(v ListAccountsForStack200ResponseLinksStack) {
	o.Service = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OperationLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OperationLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *OperationLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o OperationLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.DestinationAccount) {
		toSerialize["destination_account"] = o.DestinationAccount
	}
	if !IsNil(o.SshPortalConnections) {
		toSerialize["ssh_portal_connections"] = o.SshPortalConnections
	}
	if !IsNil(o.EphemeralSessions) {
		toSerialize["ephemeral_sessions"] = o.EphemeralSessions
	}
	if !IsNil(o.CodeScanResult) {
		toSerialize["code_scan_result"] = o.CodeScanResult
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OperationLinks) UnmarshalJSON(data []byte) (err error) {
	varOperationLinks := _OperationLinks{}

	err = json.Unmarshal(data, &varOperationLinks)

	if err != nil {
		return err
	}

	*o = OperationLinks(varOperationLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "account")
		delete(additionalProperties, "destination_account")
		delete(additionalProperties, "ssh_portal_connections")
		delete(additionalProperties, "ephemeral_sessions")
		delete(additionalProperties, "code_scan_result")
		delete(additionalProperties, "service")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperationLinks struct {
	value *OperationLinks
	isSet bool
}

func (v NullableOperationLinks) Get() *OperationLinks {
	return v.value
}

func (v *NullableOperationLinks) Set(val *OperationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationLinks(val *OperationLinks) *NullableOperationLinks {
	return &NullableOperationLinks{value: val, isSet: true}
}

func (v NullableOperationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


