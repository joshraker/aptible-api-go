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

// checks if the LogDrain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogDrain{}

// LogDrain struct for LogDrain
type LogDrain struct {
	Id int32 `json:"id"`
	MetaType string `json:"_type"`
	Handle string `json:"handle"`
	DrainType string `json:"drain_type"`
	DrainHost string `json:"drain_host"`
	DrainPort int32 `json:"drain_port"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status string `json:"status"`
	DrainUsername NullableString `json:"drain_username"`
	DrainPassword NullableString `json:"drain_password"`
	DrainApps bool `json:"drain_apps"`
	DrainDatabases bool `json:"drain_databases"`
	DrainEphemeralSessions bool `json:"drain_ephemeral_sessions"`
	DrainProxies bool `json:"drain_proxies"`
	GentlemanjerryCertificate NullableString `json:"gentlemanjerry_certificate"`
	GentlemanjerryDockerName NullableString `json:"gentlemanjerry_docker_name"`
	GentlemanjerryInstanceId NullableString `json:"gentlemanjerry_instance_id"`
	GentlemanjerryHost NullableString `json:"gentlemanjerry_host"`
	GentlemanjerryPortMapping [][]int32 `json:"gentlemanjerry_port_mapping"`
	GentlemanjerryAllocation []string `json:"gentlemanjerry_allocation"`
	Url NullableString `json:"url"`
	LoggingToken NullableString `json:"logging_token"`
	Links *LogDrainLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogDrain LogDrain

// NewLogDrain instantiates a new LogDrain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDrain(id int32, metaType string, handle string, drainType string, drainHost string, drainPort int32, createdAt string, updatedAt string, status string, drainUsername NullableString, drainPassword NullableString, drainApps bool, drainDatabases bool, drainEphemeralSessions bool, drainProxies bool, gentlemanjerryCertificate NullableString, gentlemanjerryDockerName NullableString, gentlemanjerryInstanceId NullableString, gentlemanjerryHost NullableString, gentlemanjerryPortMapping [][]int32, gentlemanjerryAllocation []string, url NullableString, loggingToken NullableString) *LogDrain {
	this := LogDrain{}
	this.Id = id
	this.MetaType = metaType
	this.Handle = handle
	this.DrainType = drainType
	this.DrainHost = drainHost
	this.DrainPort = drainPort
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	this.DrainUsername = drainUsername
	this.DrainPassword = drainPassword
	this.DrainApps = drainApps
	this.DrainDatabases = drainDatabases
	this.DrainEphemeralSessions = drainEphemeralSessions
	this.DrainProxies = drainProxies
	this.GentlemanjerryCertificate = gentlemanjerryCertificate
	this.GentlemanjerryDockerName = gentlemanjerryDockerName
	this.GentlemanjerryInstanceId = gentlemanjerryInstanceId
	this.GentlemanjerryHost = gentlemanjerryHost
	this.GentlemanjerryPortMapping = gentlemanjerryPortMapping
	this.GentlemanjerryAllocation = gentlemanjerryAllocation
	this.Url = url
	this.LoggingToken = loggingToken
	return &this
}

// NewLogDrainWithDefaults instantiates a new LogDrain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDrainWithDefaults() *LogDrain {
	this := LogDrain{}
	return &this
}

// GetId returns the Id field value
func (o *LogDrain) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogDrain) SetId(v int32) {
	o.Id = v
}

// GetMetaType returns the MetaType field value
func (o *LogDrain) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *LogDrain) SetMetaType(v string) {
	o.MetaType = v
}

// GetHandle returns the Handle field value
func (o *LogDrain) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *LogDrain) SetHandle(v string) {
	o.Handle = v
}

// GetDrainType returns the DrainType field value
func (o *LogDrain) GetDrainType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DrainType
}

// GetDrainTypeOk returns a tuple with the DrainType field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainType, true
}

// SetDrainType sets field value
func (o *LogDrain) SetDrainType(v string) {
	o.DrainType = v
}

// GetDrainHost returns the DrainHost field value
func (o *LogDrain) GetDrainHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DrainHost
}

// GetDrainHostOk returns a tuple with the DrainHost field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainHost, true
}

// SetDrainHost sets field value
func (o *LogDrain) SetDrainHost(v string) {
	o.DrainHost = v
}

// GetDrainPort returns the DrainPort field value
func (o *LogDrain) GetDrainPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DrainPort
}

// GetDrainPortOk returns a tuple with the DrainPort field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainPort, true
}

// SetDrainPort sets field value
func (o *LogDrain) SetDrainPort(v int32) {
	o.DrainPort = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LogDrain) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LogDrain) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LogDrain) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LogDrain) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *LogDrain) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LogDrain) SetStatus(v string) {
	o.Status = v
}

// GetDrainUsername returns the DrainUsername field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetDrainUsername() string {
	if o == nil || o.DrainUsername.Get() == nil {
		var ret string
		return ret
	}

	return *o.DrainUsername.Get()
}

// GetDrainUsernameOk returns a tuple with the DrainUsername field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetDrainUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DrainUsername.Get(), o.DrainUsername.IsSet()
}

// SetDrainUsername sets field value
func (o *LogDrain) SetDrainUsername(v string) {
	o.DrainUsername.Set(&v)
}

// GetDrainPassword returns the DrainPassword field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetDrainPassword() string {
	if o == nil || o.DrainPassword.Get() == nil {
		var ret string
		return ret
	}

	return *o.DrainPassword.Get()
}

// GetDrainPasswordOk returns a tuple with the DrainPassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetDrainPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DrainPassword.Get(), o.DrainPassword.IsSet()
}

// SetDrainPassword sets field value
func (o *LogDrain) SetDrainPassword(v string) {
	o.DrainPassword.Set(&v)
}

// GetDrainApps returns the DrainApps field value
func (o *LogDrain) GetDrainApps() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DrainApps
}

// GetDrainAppsOk returns a tuple with the DrainApps field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainAppsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainApps, true
}

// SetDrainApps sets field value
func (o *LogDrain) SetDrainApps(v bool) {
	o.DrainApps = v
}

// GetDrainDatabases returns the DrainDatabases field value
func (o *LogDrain) GetDrainDatabases() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DrainDatabases
}

// GetDrainDatabasesOk returns a tuple with the DrainDatabases field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainDatabasesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainDatabases, true
}

// SetDrainDatabases sets field value
func (o *LogDrain) SetDrainDatabases(v bool) {
	o.DrainDatabases = v
}

// GetDrainEphemeralSessions returns the DrainEphemeralSessions field value
func (o *LogDrain) GetDrainEphemeralSessions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DrainEphemeralSessions
}

// GetDrainEphemeralSessionsOk returns a tuple with the DrainEphemeralSessions field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainEphemeralSessionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainEphemeralSessions, true
}

// SetDrainEphemeralSessions sets field value
func (o *LogDrain) SetDrainEphemeralSessions(v bool) {
	o.DrainEphemeralSessions = v
}

// GetDrainProxies returns the DrainProxies field value
func (o *LogDrain) GetDrainProxies() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DrainProxies
}

// GetDrainProxiesOk returns a tuple with the DrainProxies field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetDrainProxiesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainProxies, true
}

// SetDrainProxies sets field value
func (o *LogDrain) SetDrainProxies(v bool) {
	o.DrainProxies = v
}

// GetGentlemanjerryCertificate returns the GentlemanjerryCertificate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetGentlemanjerryCertificate() string {
	if o == nil || o.GentlemanjerryCertificate.Get() == nil {
		var ret string
		return ret
	}

	return *o.GentlemanjerryCertificate.Get()
}

// GetGentlemanjerryCertificateOk returns a tuple with the GentlemanjerryCertificate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetGentlemanjerryCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GentlemanjerryCertificate.Get(), o.GentlemanjerryCertificate.IsSet()
}

// SetGentlemanjerryCertificate sets field value
func (o *LogDrain) SetGentlemanjerryCertificate(v string) {
	o.GentlemanjerryCertificate.Set(&v)
}

// GetGentlemanjerryDockerName returns the GentlemanjerryDockerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetGentlemanjerryDockerName() string {
	if o == nil || o.GentlemanjerryDockerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.GentlemanjerryDockerName.Get()
}

// GetGentlemanjerryDockerNameOk returns a tuple with the GentlemanjerryDockerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetGentlemanjerryDockerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GentlemanjerryDockerName.Get(), o.GentlemanjerryDockerName.IsSet()
}

// SetGentlemanjerryDockerName sets field value
func (o *LogDrain) SetGentlemanjerryDockerName(v string) {
	o.GentlemanjerryDockerName.Set(&v)
}

// GetGentlemanjerryInstanceId returns the GentlemanjerryInstanceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetGentlemanjerryInstanceId() string {
	if o == nil || o.GentlemanjerryInstanceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GentlemanjerryInstanceId.Get()
}

// GetGentlemanjerryInstanceIdOk returns a tuple with the GentlemanjerryInstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetGentlemanjerryInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GentlemanjerryInstanceId.Get(), o.GentlemanjerryInstanceId.IsSet()
}

// SetGentlemanjerryInstanceId sets field value
func (o *LogDrain) SetGentlemanjerryInstanceId(v string) {
	o.GentlemanjerryInstanceId.Set(&v)
}

// GetGentlemanjerryHost returns the GentlemanjerryHost field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetGentlemanjerryHost() string {
	if o == nil || o.GentlemanjerryHost.Get() == nil {
		var ret string
		return ret
	}

	return *o.GentlemanjerryHost.Get()
}

// GetGentlemanjerryHostOk returns a tuple with the GentlemanjerryHost field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetGentlemanjerryHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GentlemanjerryHost.Get(), o.GentlemanjerryHost.IsSet()
}

// SetGentlemanjerryHost sets field value
func (o *LogDrain) SetGentlemanjerryHost(v string) {
	o.GentlemanjerryHost.Set(&v)
}

// GetGentlemanjerryPortMapping returns the GentlemanjerryPortMapping field value
func (o *LogDrain) GetGentlemanjerryPortMapping() [][]int32 {
	if o == nil {
		var ret [][]int32
		return ret
	}

	return o.GentlemanjerryPortMapping
}

// GetGentlemanjerryPortMappingOk returns a tuple with the GentlemanjerryPortMapping field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetGentlemanjerryPortMappingOk() ([][]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GentlemanjerryPortMapping, true
}

// SetGentlemanjerryPortMapping sets field value
func (o *LogDrain) SetGentlemanjerryPortMapping(v [][]int32) {
	o.GentlemanjerryPortMapping = v
}

// GetGentlemanjerryAllocation returns the GentlemanjerryAllocation field value
func (o *LogDrain) GetGentlemanjerryAllocation() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GentlemanjerryAllocation
}

// GetGentlemanjerryAllocationOk returns a tuple with the GentlemanjerryAllocation field value
// and a boolean to check if the value has been set.
func (o *LogDrain) GetGentlemanjerryAllocationOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GentlemanjerryAllocation, true
}

// SetGentlemanjerryAllocation sets field value
func (o *LogDrain) SetGentlemanjerryAllocation(v []string) {
	o.GentlemanjerryAllocation = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *LogDrain) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetLoggingToken returns the LoggingToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogDrain) GetLoggingToken() string {
	if o == nil || o.LoggingToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoggingToken.Get()
}

// GetLoggingTokenOk returns a tuple with the LoggingToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogDrain) GetLoggingTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoggingToken.Get(), o.LoggingToken.IsSet()
}

// SetLoggingToken sets field value
func (o *LogDrain) SetLoggingToken(v string) {
	o.LoggingToken.Set(&v)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LogDrain) GetLinks() LogDrainLinks {
	if o == nil || IsNil(o.Links) {
		var ret LogDrainLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDrain) GetLinksOk() (*LogDrainLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LogDrain) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LogDrainLinks and assigns it to the Links field.
func (o *LogDrain) SetLinks(v LogDrainLinks) {
	o.Links = &v
}

func (o LogDrain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogDrain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["_type"] = o.MetaType
	toSerialize["handle"] = o.Handle
	toSerialize["drain_type"] = o.DrainType
	toSerialize["drain_host"] = o.DrainHost
	toSerialize["drain_port"] = o.DrainPort
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["status"] = o.Status
	toSerialize["drain_username"] = o.DrainUsername.Get()
	toSerialize["drain_password"] = o.DrainPassword.Get()
	toSerialize["drain_apps"] = o.DrainApps
	toSerialize["drain_databases"] = o.DrainDatabases
	toSerialize["drain_ephemeral_sessions"] = o.DrainEphemeralSessions
	toSerialize["drain_proxies"] = o.DrainProxies
	toSerialize["gentlemanjerry_certificate"] = o.GentlemanjerryCertificate.Get()
	toSerialize["gentlemanjerry_docker_name"] = o.GentlemanjerryDockerName.Get()
	toSerialize["gentlemanjerry_instance_id"] = o.GentlemanjerryInstanceId.Get()
	toSerialize["gentlemanjerry_host"] = o.GentlemanjerryHost.Get()
	toSerialize["gentlemanjerry_port_mapping"] = o.GentlemanjerryPortMapping
	toSerialize["gentlemanjerry_allocation"] = o.GentlemanjerryAllocation
	toSerialize["url"] = o.Url.Get()
	toSerialize["logging_token"] = o.LoggingToken.Get()
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogDrain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"_type",
		"handle",
		"drain_type",
		"drain_host",
		"drain_port",
		"created_at",
		"updated_at",
		"status",
		"drain_username",
		"drain_password",
		"drain_apps",
		"drain_databases",
		"drain_ephemeral_sessions",
		"drain_proxies",
		"gentlemanjerry_certificate",
		"gentlemanjerry_docker_name",
		"gentlemanjerry_instance_id",
		"gentlemanjerry_host",
		"gentlemanjerry_port_mapping",
		"gentlemanjerry_allocation",
		"url",
		"logging_token",
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

	varLogDrain := _LogDrain{}

	err = json.Unmarshal(data, &varLogDrain)

	if err != nil {
		return err
	}

	*o = LogDrain(varLogDrain)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "handle")
		delete(additionalProperties, "drain_type")
		delete(additionalProperties, "drain_host")
		delete(additionalProperties, "drain_port")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "drain_username")
		delete(additionalProperties, "drain_password")
		delete(additionalProperties, "drain_apps")
		delete(additionalProperties, "drain_databases")
		delete(additionalProperties, "drain_ephemeral_sessions")
		delete(additionalProperties, "drain_proxies")
		delete(additionalProperties, "gentlemanjerry_certificate")
		delete(additionalProperties, "gentlemanjerry_docker_name")
		delete(additionalProperties, "gentlemanjerry_instance_id")
		delete(additionalProperties, "gentlemanjerry_host")
		delete(additionalProperties, "gentlemanjerry_port_mapping")
		delete(additionalProperties, "gentlemanjerry_allocation")
		delete(additionalProperties, "url")
		delete(additionalProperties, "logging_token")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogDrain struct {
	value *LogDrain
	isSet bool
}

func (v NullableLogDrain) Get() *LogDrain {
	return v.value
}

func (v *NullableLogDrain) Set(val *LogDrain) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDrain) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDrain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDrain(val *LogDrain) *NullableLogDrain {
	return &NullableLogDrain{value: val, isSet: true}
}

func (v NullableLogDrain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDrain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


