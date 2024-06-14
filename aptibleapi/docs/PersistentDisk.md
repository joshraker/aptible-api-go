# PersistentDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaType** | **string** |  | 
**Id** | **int32** |  | 
**Handle** | **string** |  | 
**Status** | **string** |  | 
**DiskType** | **string** |  | 
**Size** | **int32** |  | 
**ProvisionedIops** | **int32** |  | 
**DiskStatus** | **NullableString** |  | 
**AwsResourceId** | **NullableString** |  | 
**KmsArn** | **NullableString** |  | 
**ModificationProgress** | **NullableInt32** |  | 
**AvailabilityZone** | **NullableString** |  | 
**Formatted** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ModifiedAt** | **NullableString** |  | 
**DeletedAt** | **NullableString** |  | 
**Links** | Pointer to [**PersistentDiskLinks**](PersistentDiskLinks.md) |  | [optional] 

## Methods

### NewPersistentDisk

`func NewPersistentDisk(metaType string, id int32, handle string, status string, diskType string, size int32, provisionedIops int32, diskStatus NullableString, awsResourceId NullableString, kmsArn NullableString, modificationProgress NullableInt32, availabilityZone NullableString, formatted bool, createdAt string, updatedAt string, modifiedAt NullableString, deletedAt NullableString, ) *PersistentDisk`

NewPersistentDisk instantiates a new PersistentDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentDiskWithDefaults

`func NewPersistentDiskWithDefaults() *PersistentDisk`

NewPersistentDiskWithDefaults instantiates a new PersistentDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaType

`func (o *PersistentDisk) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *PersistentDisk) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *PersistentDisk) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetId

`func (o *PersistentDisk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersistentDisk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersistentDisk) SetId(v int32)`

SetId sets Id field to given value.


### GetHandle

`func (o *PersistentDisk) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *PersistentDisk) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *PersistentDisk) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetStatus

`func (o *PersistentDisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PersistentDisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PersistentDisk) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDiskType

`func (o *PersistentDisk) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *PersistentDisk) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *PersistentDisk) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.


### GetSize

`func (o *PersistentDisk) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PersistentDisk) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PersistentDisk) SetSize(v int32)`

SetSize sets Size field to given value.


### GetProvisionedIops

`func (o *PersistentDisk) GetProvisionedIops() int32`

GetProvisionedIops returns the ProvisionedIops field if non-nil, zero value otherwise.

### GetProvisionedIopsOk

`func (o *PersistentDisk) GetProvisionedIopsOk() (*int32, bool)`

GetProvisionedIopsOk returns a tuple with the ProvisionedIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedIops

`func (o *PersistentDisk) SetProvisionedIops(v int32)`

SetProvisionedIops sets ProvisionedIops field to given value.


### GetDiskStatus

`func (o *PersistentDisk) GetDiskStatus() string`

GetDiskStatus returns the DiskStatus field if non-nil, zero value otherwise.

### GetDiskStatusOk

`func (o *PersistentDisk) GetDiskStatusOk() (*string, bool)`

GetDiskStatusOk returns a tuple with the DiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStatus

`func (o *PersistentDisk) SetDiskStatus(v string)`

SetDiskStatus sets DiskStatus field to given value.


### SetDiskStatusNil

`func (o *PersistentDisk) SetDiskStatusNil(b bool)`

 SetDiskStatusNil sets the value for DiskStatus to be an explicit nil

### UnsetDiskStatus
`func (o *PersistentDisk) UnsetDiskStatus()`

UnsetDiskStatus ensures that no value is present for DiskStatus, not even an explicit nil
### GetAwsResourceId

`func (o *PersistentDisk) GetAwsResourceId() string`

GetAwsResourceId returns the AwsResourceId field if non-nil, zero value otherwise.

### GetAwsResourceIdOk

`func (o *PersistentDisk) GetAwsResourceIdOk() (*string, bool)`

GetAwsResourceIdOk returns a tuple with the AwsResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsResourceId

`func (o *PersistentDisk) SetAwsResourceId(v string)`

SetAwsResourceId sets AwsResourceId field to given value.


### SetAwsResourceIdNil

`func (o *PersistentDisk) SetAwsResourceIdNil(b bool)`

 SetAwsResourceIdNil sets the value for AwsResourceId to be an explicit nil

### UnsetAwsResourceId
`func (o *PersistentDisk) UnsetAwsResourceId()`

UnsetAwsResourceId ensures that no value is present for AwsResourceId, not even an explicit nil
### GetKmsArn

`func (o *PersistentDisk) GetKmsArn() string`

GetKmsArn returns the KmsArn field if non-nil, zero value otherwise.

### GetKmsArnOk

`func (o *PersistentDisk) GetKmsArnOk() (*string, bool)`

GetKmsArnOk returns a tuple with the KmsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsArn

`func (o *PersistentDisk) SetKmsArn(v string)`

SetKmsArn sets KmsArn field to given value.


### SetKmsArnNil

`func (o *PersistentDisk) SetKmsArnNil(b bool)`

 SetKmsArnNil sets the value for KmsArn to be an explicit nil

### UnsetKmsArn
`func (o *PersistentDisk) UnsetKmsArn()`

UnsetKmsArn ensures that no value is present for KmsArn, not even an explicit nil
### GetModificationProgress

`func (o *PersistentDisk) GetModificationProgress() int32`

GetModificationProgress returns the ModificationProgress field if non-nil, zero value otherwise.

### GetModificationProgressOk

`func (o *PersistentDisk) GetModificationProgressOk() (*int32, bool)`

GetModificationProgressOk returns a tuple with the ModificationProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationProgress

`func (o *PersistentDisk) SetModificationProgress(v int32)`

SetModificationProgress sets ModificationProgress field to given value.


### SetModificationProgressNil

`func (o *PersistentDisk) SetModificationProgressNil(b bool)`

 SetModificationProgressNil sets the value for ModificationProgress to be an explicit nil

### UnsetModificationProgress
`func (o *PersistentDisk) UnsetModificationProgress()`

UnsetModificationProgress ensures that no value is present for ModificationProgress, not even an explicit nil
### GetAvailabilityZone

`func (o *PersistentDisk) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *PersistentDisk) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *PersistentDisk) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *PersistentDisk) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *PersistentDisk) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetFormatted

`func (o *PersistentDisk) GetFormatted() bool`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *PersistentDisk) GetFormattedOk() (*bool, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *PersistentDisk) SetFormatted(v bool)`

SetFormatted sets Formatted field to given value.


### GetCreatedAt

`func (o *PersistentDisk) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersistentDisk) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersistentDisk) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PersistentDisk) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PersistentDisk) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PersistentDisk) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetModifiedAt

`func (o *PersistentDisk) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PersistentDisk) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PersistentDisk) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *PersistentDisk) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PersistentDisk) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetDeletedAt

`func (o *PersistentDisk) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PersistentDisk) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PersistentDisk) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *PersistentDisk) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *PersistentDisk) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetLinks

`func (o *PersistentDisk) GetLinks() PersistentDiskLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PersistentDisk) GetLinksOk() (*PersistentDiskLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PersistentDisk) SetLinks(v PersistentDiskLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PersistentDisk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


