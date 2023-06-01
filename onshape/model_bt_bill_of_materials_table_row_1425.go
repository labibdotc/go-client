/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16643-ef813b2da145
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsTableRow1425 struct for BTBillOfMaterialsTableRow1425
type BTBillOfMaterialsTableRow1425 struct {
	BtType                 *string                            `json:"btType,omitempty"`
	ColumnIdToCell         *map[string]BTTableCell1114        `json:"columnIdToCell,omitempty"`
	Id                     *string                            `json:"id,omitempty"`
	MetaData               *BTTreeNode20                      `json:"metaData,omitempty"`
	NodeId                 *string                            `json:"nodeId,omitempty"`
	RowMetadata            *BTTableBaseRowMetadata3181        `json:"rowMetadata,omitempty"`
	ExcludeIsEditable      *bool                              `json:"excludeIsEditable,omitempty"`
	ExclusionStatus        *GBTBillOfMaterialsExclusionStatus `json:"exclusionStatus,omitempty"`
	ExpansionStatus        *GBTBillOfMaterialsExpansionStatus `json:"expansionStatus,omitempty"`
	IndentLevel            *int32                             `json:"indentLevel,omitempty"`
	MetadataUpdateHref     *string                            `json:"metadataUpdateHref,omitempty"`
	RelatedOccurrencePaths []string                           `json:"relatedOccurrencePaths,omitempty"`
	UniqueItemId           *BTBillOfMaterialsUniqueItemId2029 `json:"uniqueItemId,omitempty"`
}

// NewBTBillOfMaterialsTableRow1425 instantiates a new BTBillOfMaterialsTableRow1425 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsTableRow1425() *BTBillOfMaterialsTableRow1425 {
	this := BTBillOfMaterialsTableRow1425{}
	return &this
}

// NewBTBillOfMaterialsTableRow1425WithDefaults instantiates a new BTBillOfMaterialsTableRow1425 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsTableRow1425WithDefaults() *BTBillOfMaterialsTableRow1425 {
	this := BTBillOfMaterialsTableRow1425{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsTableRow1425) SetBtType(v string) {
	o.BtType = &v
}

// GetColumnIdToCell returns the ColumnIdToCell field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetColumnIdToCell() map[string]BTTableCell1114 {
	if o == nil || o.ColumnIdToCell == nil {
		var ret map[string]BTTableCell1114
		return ret
	}
	return *o.ColumnIdToCell
}

// GetColumnIdToCellOk returns a tuple with the ColumnIdToCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetColumnIdToCellOk() (*map[string]BTTableCell1114, bool) {
	if o == nil || o.ColumnIdToCell == nil {
		return nil, false
	}
	return o.ColumnIdToCell, true
}

// HasColumnIdToCell returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasColumnIdToCell() bool {
	if o != nil && o.ColumnIdToCell != nil {
		return true
	}

	return false
}

// SetColumnIdToCell gets a reference to the given map[string]BTTableCell1114 and assigns it to the ColumnIdToCell field.
func (o *BTBillOfMaterialsTableRow1425) SetColumnIdToCell(v map[string]BTTableCell1114) {
	o.ColumnIdToCell = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillOfMaterialsTableRow1425) SetId(v string) {
	o.Id = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetMetaData() BTTreeNode20 {
	if o == nil || o.MetaData == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetMetaDataOk() (*BTTreeNode20, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given BTTreeNode20 and assigns it to the MetaData field.
func (o *BTBillOfMaterialsTableRow1425) SetMetaData(v BTTreeNode20) {
	o.MetaData = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTBillOfMaterialsTableRow1425) SetNodeId(v string) {
	o.NodeId = &v
}

// GetRowMetadata returns the RowMetadata field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetRowMetadata() BTTableBaseRowMetadata3181 {
	if o == nil || o.RowMetadata == nil {
		var ret BTTableBaseRowMetadata3181
		return ret
	}
	return *o.RowMetadata
}

// GetRowMetadataOk returns a tuple with the RowMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetRowMetadataOk() (*BTTableBaseRowMetadata3181, bool) {
	if o == nil || o.RowMetadata == nil {
		return nil, false
	}
	return o.RowMetadata, true
}

// HasRowMetadata returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasRowMetadata() bool {
	if o != nil && o.RowMetadata != nil {
		return true
	}

	return false
}

// SetRowMetadata gets a reference to the given BTTableBaseRowMetadata3181 and assigns it to the RowMetadata field.
func (o *BTBillOfMaterialsTableRow1425) SetRowMetadata(v BTTableBaseRowMetadata3181) {
	o.RowMetadata = &v
}

// GetExcludeIsEditable returns the ExcludeIsEditable field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetExcludeIsEditable() bool {
	if o == nil || o.ExcludeIsEditable == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeIsEditable
}

// GetExcludeIsEditableOk returns a tuple with the ExcludeIsEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetExcludeIsEditableOk() (*bool, bool) {
	if o == nil || o.ExcludeIsEditable == nil {
		return nil, false
	}
	return o.ExcludeIsEditable, true
}

// HasExcludeIsEditable returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasExcludeIsEditable() bool {
	if o != nil && o.ExcludeIsEditable != nil {
		return true
	}

	return false
}

// SetExcludeIsEditable gets a reference to the given bool and assigns it to the ExcludeIsEditable field.
func (o *BTBillOfMaterialsTableRow1425) SetExcludeIsEditable(v bool) {
	o.ExcludeIsEditable = &v
}

// GetExclusionStatus returns the ExclusionStatus field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetExclusionStatus() GBTBillOfMaterialsExclusionStatus {
	if o == nil || o.ExclusionStatus == nil {
		var ret GBTBillOfMaterialsExclusionStatus
		return ret
	}
	return *o.ExclusionStatus
}

// GetExclusionStatusOk returns a tuple with the ExclusionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetExclusionStatusOk() (*GBTBillOfMaterialsExclusionStatus, bool) {
	if o == nil || o.ExclusionStatus == nil {
		return nil, false
	}
	return o.ExclusionStatus, true
}

// HasExclusionStatus returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasExclusionStatus() bool {
	if o != nil && o.ExclusionStatus != nil {
		return true
	}

	return false
}

// SetExclusionStatus gets a reference to the given GBTBillOfMaterialsExclusionStatus and assigns it to the ExclusionStatus field.
func (o *BTBillOfMaterialsTableRow1425) SetExclusionStatus(v GBTBillOfMaterialsExclusionStatus) {
	o.ExclusionStatus = &v
}

// GetExpansionStatus returns the ExpansionStatus field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetExpansionStatus() GBTBillOfMaterialsExpansionStatus {
	if o == nil || o.ExpansionStatus == nil {
		var ret GBTBillOfMaterialsExpansionStatus
		return ret
	}
	return *o.ExpansionStatus
}

// GetExpansionStatusOk returns a tuple with the ExpansionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetExpansionStatusOk() (*GBTBillOfMaterialsExpansionStatus, bool) {
	if o == nil || o.ExpansionStatus == nil {
		return nil, false
	}
	return o.ExpansionStatus, true
}

// HasExpansionStatus returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasExpansionStatus() bool {
	if o != nil && o.ExpansionStatus != nil {
		return true
	}

	return false
}

// SetExpansionStatus gets a reference to the given GBTBillOfMaterialsExpansionStatus and assigns it to the ExpansionStatus field.
func (o *BTBillOfMaterialsTableRow1425) SetExpansionStatus(v GBTBillOfMaterialsExpansionStatus) {
	o.ExpansionStatus = &v
}

// GetIndentLevel returns the IndentLevel field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetIndentLevel() int32 {
	if o == nil || o.IndentLevel == nil {
		var ret int32
		return ret
	}
	return *o.IndentLevel
}

// GetIndentLevelOk returns a tuple with the IndentLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetIndentLevelOk() (*int32, bool) {
	if o == nil || o.IndentLevel == nil {
		return nil, false
	}
	return o.IndentLevel, true
}

// HasIndentLevel returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasIndentLevel() bool {
	if o != nil && o.IndentLevel != nil {
		return true
	}

	return false
}

// SetIndentLevel gets a reference to the given int32 and assigns it to the IndentLevel field.
func (o *BTBillOfMaterialsTableRow1425) SetIndentLevel(v int32) {
	o.IndentLevel = &v
}

// GetMetadataUpdateHref returns the MetadataUpdateHref field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetMetadataUpdateHref() string {
	if o == nil || o.MetadataUpdateHref == nil {
		var ret string
		return ret
	}
	return *o.MetadataUpdateHref
}

// GetMetadataUpdateHrefOk returns a tuple with the MetadataUpdateHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetMetadataUpdateHrefOk() (*string, bool) {
	if o == nil || o.MetadataUpdateHref == nil {
		return nil, false
	}
	return o.MetadataUpdateHref, true
}

// HasMetadataUpdateHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasMetadataUpdateHref() bool {
	if o != nil && o.MetadataUpdateHref != nil {
		return true
	}

	return false
}

// SetMetadataUpdateHref gets a reference to the given string and assigns it to the MetadataUpdateHref field.
func (o *BTBillOfMaterialsTableRow1425) SetMetadataUpdateHref(v string) {
	o.MetadataUpdateHref = &v
}

// GetRelatedOccurrencePaths returns the RelatedOccurrencePaths field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetRelatedOccurrencePaths() []string {
	if o == nil || o.RelatedOccurrencePaths == nil {
		var ret []string
		return ret
	}
	return o.RelatedOccurrencePaths
}

// GetRelatedOccurrencePathsOk returns a tuple with the RelatedOccurrencePaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetRelatedOccurrencePathsOk() ([]string, bool) {
	if o == nil || o.RelatedOccurrencePaths == nil {
		return nil, false
	}
	return o.RelatedOccurrencePaths, true
}

// HasRelatedOccurrencePaths returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasRelatedOccurrencePaths() bool {
	if o != nil && o.RelatedOccurrencePaths != nil {
		return true
	}

	return false
}

// SetRelatedOccurrencePaths gets a reference to the given []string and assigns it to the RelatedOccurrencePaths field.
func (o *BTBillOfMaterialsTableRow1425) SetRelatedOccurrencePaths(v []string) {
	o.RelatedOccurrencePaths = v
}

// GetUniqueItemId returns the UniqueItemId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRow1425) GetUniqueItemId() BTBillOfMaterialsUniqueItemId2029 {
	if o == nil || o.UniqueItemId == nil {
		var ret BTBillOfMaterialsUniqueItemId2029
		return ret
	}
	return *o.UniqueItemId
}

// GetUniqueItemIdOk returns a tuple with the UniqueItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRow1425) GetUniqueItemIdOk() (*BTBillOfMaterialsUniqueItemId2029, bool) {
	if o == nil || o.UniqueItemId == nil {
		return nil, false
	}
	return o.UniqueItemId, true
}

// HasUniqueItemId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRow1425) HasUniqueItemId() bool {
	if o != nil && o.UniqueItemId != nil {
		return true
	}

	return false
}

// SetUniqueItemId gets a reference to the given BTBillOfMaterialsUniqueItemId2029 and assigns it to the UniqueItemId field.
func (o *BTBillOfMaterialsTableRow1425) SetUniqueItemId(v BTBillOfMaterialsUniqueItemId2029) {
	o.UniqueItemId = &v
}

func (o BTBillOfMaterialsTableRow1425) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ColumnIdToCell != nil {
		toSerialize["columnIdToCell"] = o.ColumnIdToCell
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetaData != nil {
		toSerialize["metaData"] = o.MetaData
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.RowMetadata != nil {
		toSerialize["rowMetadata"] = o.RowMetadata
	}
	if o.ExcludeIsEditable != nil {
		toSerialize["excludeIsEditable"] = o.ExcludeIsEditable
	}
	if o.ExclusionStatus != nil {
		toSerialize["exclusionStatus"] = o.ExclusionStatus
	}
	if o.ExpansionStatus != nil {
		toSerialize["expansionStatus"] = o.ExpansionStatus
	}
	if o.IndentLevel != nil {
		toSerialize["indentLevel"] = o.IndentLevel
	}
	if o.MetadataUpdateHref != nil {
		toSerialize["metadataUpdateHref"] = o.MetadataUpdateHref
	}
	if o.RelatedOccurrencePaths != nil {
		toSerialize["relatedOccurrencePaths"] = o.RelatedOccurrencePaths
	}
	if o.UniqueItemId != nil {
		toSerialize["uniqueItemId"] = o.UniqueItemId
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsTableRow1425 struct {
	value *BTBillOfMaterialsTableRow1425
	isSet bool
}

func (v NullableBTBillOfMaterialsTableRow1425) Get() *BTBillOfMaterialsTableRow1425 {
	return v.value
}

func (v *NullableBTBillOfMaterialsTableRow1425) Set(val *BTBillOfMaterialsTableRow1425) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsTableRow1425) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsTableRow1425) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsTableRow1425(val *BTBillOfMaterialsTableRow1425) *NullableBTBillOfMaterialsTableRow1425 {
	return &NullableBTBillOfMaterialsTableRow1425{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsTableRow1425) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsTableRow1425) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
