# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**AwsSnapshotId** | **string** |  | 
**AwsRegion** | **string** |  | 
**DatabaseHandle** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**KmsArn** | **NullableString** |  | 
**Size** | **NullableInt32** |  | 
**Manual** | **NullableBool** |  | 
**Embedded** | [**BackupEmbedded**](BackupEmbedded.md) |  | 
**Links** | Pointer to [**BackupLinks**](BackupLinks.md) |  | [optional] 

## Methods

### NewBackup

`func NewBackup(id int32, metaType string, awsSnapshotId string, awsRegion string, databaseHandle string, createdAt string, updatedAt string, kmsArn NullableString, size NullableInt32, manual NullableBool, embedded BackupEmbedded, ) *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Backup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Backup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Backup) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Backup) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Backup) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Backup) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetAwsSnapshotId

`func (o *Backup) GetAwsSnapshotId() string`

GetAwsSnapshotId returns the AwsSnapshotId field if non-nil, zero value otherwise.

### GetAwsSnapshotIdOk

`func (o *Backup) GetAwsSnapshotIdOk() (*string, bool)`

GetAwsSnapshotIdOk returns a tuple with the AwsSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSnapshotId

`func (o *Backup) SetAwsSnapshotId(v string)`

SetAwsSnapshotId sets AwsSnapshotId field to given value.


### GetAwsRegion

`func (o *Backup) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *Backup) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *Backup) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetDatabaseHandle

`func (o *Backup) GetDatabaseHandle() string`

GetDatabaseHandle returns the DatabaseHandle field if non-nil, zero value otherwise.

### GetDatabaseHandleOk

`func (o *Backup) GetDatabaseHandleOk() (*string, bool)`

GetDatabaseHandleOk returns a tuple with the DatabaseHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseHandle

`func (o *Backup) SetDatabaseHandle(v string)`

SetDatabaseHandle sets DatabaseHandle field to given value.


### GetCreatedAt

`func (o *Backup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Backup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Backup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Backup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Backup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Backup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKmsArn

`func (o *Backup) GetKmsArn() string`

GetKmsArn returns the KmsArn field if non-nil, zero value otherwise.

### GetKmsArnOk

`func (o *Backup) GetKmsArnOk() (*string, bool)`

GetKmsArnOk returns a tuple with the KmsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsArn

`func (o *Backup) SetKmsArn(v string)`

SetKmsArn sets KmsArn field to given value.


### SetKmsArnNil

`func (o *Backup) SetKmsArnNil(b bool)`

 SetKmsArnNil sets the value for KmsArn to be an explicit nil

### UnsetKmsArn
`func (o *Backup) UnsetKmsArn()`

UnsetKmsArn ensures that no value is present for KmsArn, not even an explicit nil
### GetSize

`func (o *Backup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Backup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Backup) SetSize(v int32)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *Backup) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Backup) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetManual

`func (o *Backup) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *Backup) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *Backup) SetManual(v bool)`

SetManual sets Manual field to given value.


### SetManualNil

`func (o *Backup) SetManualNil(b bool)`

 SetManualNil sets the value for Manual to be an explicit nil

### UnsetManual
`func (o *Backup) UnsetManual()`

UnsetManual ensures that no value is present for Manual, not even an explicit nil
### GetEmbedded

`func (o *Backup) GetEmbedded() BackupEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Backup) GetEmbeddedOk() (*BackupEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Backup) SetEmbedded(v BackupEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *Backup) GetLinks() BackupLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Backup) GetLinksOk() (*BackupLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Backup) SetLinks(v BackupLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Backup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


