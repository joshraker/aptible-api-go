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

// checks if the GetRoot200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoot200ResponseLinks{}

// GetRoot200ResponseLinks struct for GetRoot200ResponseLinks
type GetRoot200ResponseLinks struct {
	Accounts *ListAccountsForStack200ResponseLinksStack `json:"accounts,omitempty"`
	Apps *ListAccountsForStack200ResponseLinksStack `json:"apps,omitempty"`
	Databases *ListAccountsForStack200ResponseLinksStack `json:"databases,omitempty"`
	DatabaseImages *ListAccountsForStack200ResponseLinksStack `json:"database_images,omitempty"`
	Stacks *ListAccountsForStack200ResponseLinksStack `json:"stacks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRoot200ResponseLinks GetRoot200ResponseLinks

// NewGetRoot200ResponseLinks instantiates a new GetRoot200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoot200ResponseLinks() *GetRoot200ResponseLinks {
	this := GetRoot200ResponseLinks{}
	return &this
}

// NewGetRoot200ResponseLinksWithDefaults instantiates a new GetRoot200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoot200ResponseLinksWithDefaults() *GetRoot200ResponseLinks {
	this := GetRoot200ResponseLinks{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *GetRoot200ResponseLinks) GetAccounts() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Accounts) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoot200ResponseLinks) GetAccountsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *GetRoot200ResponseLinks) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Accounts field.
func (o *GetRoot200ResponseLinks) SetAccounts(v ListAccountsForStack200ResponseLinksStack) {
	o.Accounts = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *GetRoot200ResponseLinks) GetApps() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Apps) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoot200ResponseLinks) GetAppsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *GetRoot200ResponseLinks) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Apps field.
func (o *GetRoot200ResponseLinks) SetApps(v ListAccountsForStack200ResponseLinksStack) {
	o.Apps = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *GetRoot200ResponseLinks) GetDatabases() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Databases) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoot200ResponseLinks) GetDatabasesOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *GetRoot200ResponseLinks) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Databases field.
func (o *GetRoot200ResponseLinks) SetDatabases(v ListAccountsForStack200ResponseLinksStack) {
	o.Databases = &v
}

// GetDatabaseImages returns the DatabaseImages field value if set, zero value otherwise.
func (o *GetRoot200ResponseLinks) GetDatabaseImages() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.DatabaseImages) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.DatabaseImages
}

// GetDatabaseImagesOk returns a tuple with the DatabaseImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoot200ResponseLinks) GetDatabaseImagesOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.DatabaseImages) {
		return nil, false
	}
	return o.DatabaseImages, true
}

// HasDatabaseImages returns a boolean if a field has been set.
func (o *GetRoot200ResponseLinks) HasDatabaseImages() bool {
	if o != nil && !IsNil(o.DatabaseImages) {
		return true
	}

	return false
}

// SetDatabaseImages gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the DatabaseImages field.
func (o *GetRoot200ResponseLinks) SetDatabaseImages(v ListAccountsForStack200ResponseLinksStack) {
	o.DatabaseImages = &v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *GetRoot200ResponseLinks) GetStacks() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Stacks) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoot200ResponseLinks) GetStacksOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Stacks) {
		return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *GetRoot200ResponseLinks) HasStacks() bool {
	if o != nil && !IsNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Stacks field.
func (o *GetRoot200ResponseLinks) SetStacks(v ListAccountsForStack200ResponseLinksStack) {
	o.Stacks = &v
}

func (o GetRoot200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoot200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Databases) {
		toSerialize["databases"] = o.Databases
	}
	if !IsNil(o.DatabaseImages) {
		toSerialize["database_images"] = o.DatabaseImages
	}
	if !IsNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRoot200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varGetRoot200ResponseLinks := _GetRoot200ResponseLinks{}

	err = json.Unmarshal(data, &varGetRoot200ResponseLinks)

	if err != nil {
		return err
	}

	*o = GetRoot200ResponseLinks(varGetRoot200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "databases")
		delete(additionalProperties, "database_images")
		delete(additionalProperties, "stacks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRoot200ResponseLinks struct {
	value *GetRoot200ResponseLinks
	isSet bool
}

func (v NullableGetRoot200ResponseLinks) Get() *GetRoot200ResponseLinks {
	return v.value
}

func (v *NullableGetRoot200ResponseLinks) Set(val *GetRoot200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoot200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoot200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoot200ResponseLinks(val *GetRoot200ResponseLinks) *NullableGetRoot200ResponseLinks {
	return &NullableGetRoot200ResponseLinks{value: val, isSet: true}
}

func (v NullableGetRoot200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoot200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

