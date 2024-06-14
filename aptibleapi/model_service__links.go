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

// checks if the ServiceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceLinks{}

// ServiceLinks struct for ServiceLinks
type ServiceLinks struct {
	Account *ListAccountsForStack200ResponseLinksStack `json:"account,omitempty"`
	CurrentRelease *ListAccountsForStack200ResponseLinksStack `json:"current_release,omitempty"`
	App *ListAccountsForStack200ResponseLinksStack `json:"app,omitempty"`
	Database *ListAccountsForStack200ResponseLinksStack `json:"database,omitempty"`
	Operations *ListAccountsForStack200ResponseLinksStack `json:"operations,omitempty"`
	Releases *ListAccountsForStack200ResponseLinksStack `json:"releases,omitempty"`
	Vhosts *ListAccountsForStack200ResponseLinksStack `json:"vhosts,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceLinks ServiceLinks

// NewServiceLinks instantiates a new ServiceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLinks() *ServiceLinks {
	this := ServiceLinks{}
	return &this
}

// NewServiceLinksWithDefaults instantiates a new ServiceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLinksWithDefaults() *ServiceLinks {
	this := ServiceLinks{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ServiceLinks) GetAccount() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Account) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetAccountOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ServiceLinks) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Account field.
func (o *ServiceLinks) SetAccount(v ListAccountsForStack200ResponseLinksStack) {
	o.Account = &v
}

// GetCurrentRelease returns the CurrentRelease field value if set, zero value otherwise.
func (o *ServiceLinks) GetCurrentRelease() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.CurrentRelease) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.CurrentRelease
}

// GetCurrentReleaseOk returns a tuple with the CurrentRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetCurrentReleaseOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.CurrentRelease) {
		return nil, false
	}
	return o.CurrentRelease, true
}

// HasCurrentRelease returns a boolean if a field has been set.
func (o *ServiceLinks) HasCurrentRelease() bool {
	if o != nil && !IsNil(o.CurrentRelease) {
		return true
	}

	return false
}

// SetCurrentRelease gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the CurrentRelease field.
func (o *ServiceLinks) SetCurrentRelease(v ListAccountsForStack200ResponseLinksStack) {
	o.CurrentRelease = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ServiceLinks) GetApp() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.App) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetAppOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ServiceLinks) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the App field.
func (o *ServiceLinks) SetApp(v ListAccountsForStack200ResponseLinksStack) {
	o.App = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ServiceLinks) GetDatabase() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Database) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetDatabaseOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ServiceLinks) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Database field.
func (o *ServiceLinks) SetDatabase(v ListAccountsForStack200ResponseLinksStack) {
	o.Database = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *ServiceLinks) GetOperations() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Operations) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetOperationsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ServiceLinks) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Operations field.
func (o *ServiceLinks) SetOperations(v ListAccountsForStack200ResponseLinksStack) {
	o.Operations = &v
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *ServiceLinks) GetReleases() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Releases) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetReleasesOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *ServiceLinks) HasReleases() bool {
	if o != nil && !IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Releases field.
func (o *ServiceLinks) SetReleases(v ListAccountsForStack200ResponseLinksStack) {
	o.Releases = &v
}

// GetVhosts returns the Vhosts field value if set, zero value otherwise.
func (o *ServiceLinks) GetVhosts() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Vhosts) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Vhosts
}

// GetVhostsOk returns a tuple with the Vhosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetVhostsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Vhosts) {
		return nil, false
	}
	return o.Vhosts, true
}

// HasVhosts returns a boolean if a field has been set.
func (o *ServiceLinks) HasVhosts() bool {
	if o != nil && !IsNil(o.Vhosts) {
		return true
	}

	return false
}

// SetVhosts gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Vhosts field.
func (o *ServiceLinks) SetVhosts(v ListAccountsForStack200ResponseLinksStack) {
	o.Vhosts = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ServiceLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ServiceLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ServiceLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ServiceLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.CurrentRelease) {
		toSerialize["current_release"] = o.CurrentRelease
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Releases) {
		toSerialize["releases"] = o.Releases
	}
	if !IsNil(o.Vhosts) {
		toSerialize["vhosts"] = o.Vhosts
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceLinks) UnmarshalJSON(data []byte) (err error) {
	varServiceLinks := _ServiceLinks{}

	err = json.Unmarshal(data, &varServiceLinks)

	if err != nil {
		return err
	}

	*o = ServiceLinks(varServiceLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "current_release")
		delete(additionalProperties, "app")
		delete(additionalProperties, "database")
		delete(additionalProperties, "operations")
		delete(additionalProperties, "releases")
		delete(additionalProperties, "vhosts")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceLinks struct {
	value *ServiceLinks
	isSet bool
}

func (v NullableServiceLinks) Get() *ServiceLinks {
	return v.value
}

func (v *NullableServiceLinks) Set(val *ServiceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLinks(val *ServiceLinks) *NullableServiceLinks {
	return &NullableServiceLinks{value: val, isSet: true}
}

func (v NullableServiceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

