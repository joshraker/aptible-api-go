# VpcPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ConnectionId** | **string** |  | 
**ConnectionStatus** | **string** |  | 
**PeerAccountId** | **string** |  | 
**PeerVpcId** | **string** |  | 
**Description** | **string** |  | 
**Links** | Pointer to [**VpcPeerLinks**](VpcPeerLinks.md) |  | [optional] 

## Methods

### NewVpcPeer

`func NewVpcPeer(id int32, metaType string, createdAt string, updatedAt string, connectionId string, connectionStatus string, peerAccountId string, peerVpcId string, description string, ) *VpcPeer`

NewVpcPeer instantiates a new VpcPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeerWithDefaults

`func NewVpcPeerWithDefaults() *VpcPeer`

NewVpcPeerWithDefaults instantiates a new VpcPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpcPeer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcPeer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcPeer) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *VpcPeer) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *VpcPeer) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *VpcPeer) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *VpcPeer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpcPeer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpcPeer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VpcPeer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VpcPeer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VpcPeer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetConnectionId

`func (o *VpcPeer) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *VpcPeer) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *VpcPeer) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnectionStatus

`func (o *VpcPeer) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *VpcPeer) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *VpcPeer) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetPeerAccountId

`func (o *VpcPeer) GetPeerAccountId() string`

GetPeerAccountId returns the PeerAccountId field if non-nil, zero value otherwise.

### GetPeerAccountIdOk

`func (o *VpcPeer) GetPeerAccountIdOk() (*string, bool)`

GetPeerAccountIdOk returns a tuple with the PeerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAccountId

`func (o *VpcPeer) SetPeerAccountId(v string)`

SetPeerAccountId sets PeerAccountId field to given value.


### GetPeerVpcId

`func (o *VpcPeer) GetPeerVpcId() string`

GetPeerVpcId returns the PeerVpcId field if non-nil, zero value otherwise.

### GetPeerVpcIdOk

`func (o *VpcPeer) GetPeerVpcIdOk() (*string, bool)`

GetPeerVpcIdOk returns a tuple with the PeerVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerVpcId

`func (o *VpcPeer) SetPeerVpcId(v string)`

SetPeerVpcId sets PeerVpcId field to given value.


### GetDescription

`func (o *VpcPeer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcPeer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcPeer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLinks

`func (o *VpcPeer) GetLinks() VpcPeerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VpcPeer) GetLinksOk() (*VpcPeerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VpcPeer) SetLinks(v VpcPeerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VpcPeer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


