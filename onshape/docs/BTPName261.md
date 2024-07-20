# BTPName261

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ForExport** | Pointer to **bool** |  | [optional] 
**GlobalNamespace** | Pointer to **bool** |  | [optional] 
**Identifier** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Element microversion that is being imported. | [optional] 
**Namespace** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 

## Methods

### NewBTPName261

`func NewBTPName261() *BTPName261`

NewBTPName261 instantiates a new BTPName261 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPName261WithDefaults

`func NewBTPName261WithDefaults() *BTPName261`

NewBTPName261WithDefaults instantiates a new BTPName261 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPName261) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPName261) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPName261) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPName261) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetForExport

`func (o *BTPName261) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *BTPName261) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *BTPName261) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *BTPName261) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetGlobalNamespace

`func (o *BTPName261) GetGlobalNamespace() bool`

GetGlobalNamespace returns the GlobalNamespace field if non-nil, zero value otherwise.

### GetGlobalNamespaceOk

`func (o *BTPName261) GetGlobalNamespaceOk() (*bool, bool)`

GetGlobalNamespaceOk returns a tuple with the GlobalNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalNamespace

`func (o *BTPName261) SetGlobalNamespace(v bool)`

SetGlobalNamespace sets GlobalNamespace field to given value.

### HasGlobalNamespace

`func (o *BTPName261) HasGlobalNamespace() bool`

HasGlobalNamespace returns a boolean if a field has been set.

### GetIdentifier

`func (o *BTPName261) GetIdentifier() BTPIdentifier8`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BTPName261) GetIdentifierOk() (*BTPIdentifier8, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BTPName261) SetIdentifier(v BTPIdentifier8)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BTPName261) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTPName261) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTPName261) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTPName261) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTPName261) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetNamespace

`func (o *BTPName261) GetNamespace() []BTPIdentifier8`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTPName261) GetNamespaceOk() (*[]BTPIdentifier8, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTPName261) SetNamespace(v []BTPIdentifier8)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTPName261) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


