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

// checks if the ListBackupsForAccount200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBackupsForAccount200ResponseLinks{}

// ListBackupsForAccount200ResponseLinks struct for ListBackupsForAccount200ResponseLinks
type ListBackupsForAccount200ResponseLinks struct {
	Account *ListAccountsForStack200ResponseLinksStack `json:"account,omitempty"`
	Database *ListAccountsForStack200ResponseLinksStack `json:"database,omitempty"`
	Next *ListAccountsForStack200ResponseLinksStack `json:"next,omitempty"`
	Prev *ListAccountsForStack200ResponseLinksStack `json:"prev,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListBackupsForAccount200ResponseLinks ListBackupsForAccount200ResponseLinks

// NewListBackupsForAccount200ResponseLinks instantiates a new ListBackupsForAccount200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBackupsForAccount200ResponseLinks() *ListBackupsForAccount200ResponseLinks {
	this := ListBackupsForAccount200ResponseLinks{}
	return &this
}

// NewListBackupsForAccount200ResponseLinksWithDefaults instantiates a new ListBackupsForAccount200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBackupsForAccount200ResponseLinksWithDefaults() *ListBackupsForAccount200ResponseLinks {
	this := ListBackupsForAccount200ResponseLinks{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ListBackupsForAccount200ResponseLinks) GetAccount() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Account) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsForAccount200ResponseLinks) GetAccountOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ListBackupsForAccount200ResponseLinks) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Account field.
func (o *ListBackupsForAccount200ResponseLinks) SetAccount(v ListAccountsForStack200ResponseLinksStack) {
	o.Account = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ListBackupsForAccount200ResponseLinks) GetDatabase() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Database) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsForAccount200ResponseLinks) GetDatabaseOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ListBackupsForAccount200ResponseLinks) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Database field.
func (o *ListBackupsForAccount200ResponseLinks) SetDatabase(v ListAccountsForStack200ResponseLinksStack) {
	o.Database = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListBackupsForAccount200ResponseLinks) GetNext() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Next) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsForAccount200ResponseLinks) GetNextOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListBackupsForAccount200ResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Next field.
func (o *ListBackupsForAccount200ResponseLinks) SetNext(v ListAccountsForStack200ResponseLinksStack) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *ListBackupsForAccount200ResponseLinks) GetPrev() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Prev) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsForAccount200ResponseLinks) GetPrevOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *ListBackupsForAccount200ResponseLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Prev field.
func (o *ListBackupsForAccount200ResponseLinks) SetPrev(v ListAccountsForStack200ResponseLinksStack) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListBackupsForAccount200ResponseLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsForAccount200ResponseLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListBackupsForAccount200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ListBackupsForAccount200ResponseLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ListBackupsForAccount200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBackupsForAccount200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListBackupsForAccount200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varListBackupsForAccount200ResponseLinks := _ListBackupsForAccount200ResponseLinks{}

	err = json.Unmarshal(data, &varListBackupsForAccount200ResponseLinks)

	if err != nil {
		return err
	}

	*o = ListBackupsForAccount200ResponseLinks(varListBackupsForAccount200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "database")
		delete(additionalProperties, "next")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListBackupsForAccount200ResponseLinks struct {
	value *ListBackupsForAccount200ResponseLinks
	isSet bool
}

func (v NullableListBackupsForAccount200ResponseLinks) Get() *ListBackupsForAccount200ResponseLinks {
	return v.value
}

func (v *NullableListBackupsForAccount200ResponseLinks) Set(val *ListBackupsForAccount200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListBackupsForAccount200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListBackupsForAccount200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBackupsForAccount200ResponseLinks(val *ListBackupsForAccount200ResponseLinks) *NullableListBackupsForAccount200ResponseLinks {
	return &NullableListBackupsForAccount200ResponseLinks{value: val, isSet: true}
}

func (v NullableListBackupsForAccount200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBackupsForAccount200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


