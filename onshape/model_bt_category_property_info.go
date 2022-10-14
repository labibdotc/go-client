/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6905-ae59ed040327
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCategoryPropertyInfo struct for BTCategoryPropertyInfo
type BTCategoryPropertyInfo struct {
	Array                      *bool                           `json:"array,omitempty"`
	Assignable                 *bool                           `json:"assignable,omitempty"`
	BlobMimeType               *string                         `json:"blobMimeType,omitempty"`
	CategoryPropertyConfigInfo *BTCategoryPropertyConfigInfo   `json:"categoryPropertyConfigInfo,omitempty"`
	CategorySummaryInfoList    []BTMetadataCategorySummaryInfo `json:"categorySummaryInfoList,omitempty"`
	Description                *string                         `json:"description,omitempty"`
	EditableInMicroversion     *bool                           `json:"editableInMicroversion,omitempty"`
	EditableInVersion          *bool                           `json:"editableInVersion,omitempty"`
	Href                       *string                         `json:"href,omitempty"`
	Id                         *string                         `json:"id,omitempty"`
	Name                       *string                         `json:"name,omitempty"`
	ObjectDefName              *string                         `json:"objectDefName,omitempty"`
	OwnerId                    *string                         `json:"ownerId,omitempty"`
	OwnerType                  *int32                          `json:"ownerType,omitempty"`
	UiReadonlyInMicroversion   *bool                           `json:"uiReadonlyInMicroversion,omitempty"`
	UiReadonlyInVersion        *bool                           `json:"uiReadonlyInVersion,omitempty"`
	ValueType                  *int32                          `json:"valueType,omitempty"`
	ViewRef                    *string                         `json:"viewRef,omitempty"`
}

// NewBTCategoryPropertyInfo instantiates a new BTCategoryPropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCategoryPropertyInfo() *BTCategoryPropertyInfo {
	this := BTCategoryPropertyInfo{}
	return &this
}

// NewBTCategoryPropertyInfoWithDefaults instantiates a new BTCategoryPropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCategoryPropertyInfoWithDefaults() *BTCategoryPropertyInfo {
	this := BTCategoryPropertyInfo{}
	return &this
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetArray() bool {
	if o == nil || o.Array == nil {
		var ret bool
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetArrayOk() (*bool, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given bool and assigns it to the Array field.
func (o *BTCategoryPropertyInfo) SetArray(v bool) {
	o.Array = &v
}

// GetAssignable returns the Assignable field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetAssignable() bool {
	if o == nil || o.Assignable == nil {
		var ret bool
		return ret
	}
	return *o.Assignable
}

// GetAssignableOk returns a tuple with the Assignable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetAssignableOk() (*bool, bool) {
	if o == nil || o.Assignable == nil {
		return nil, false
	}
	return o.Assignable, true
}

// HasAssignable returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasAssignable() bool {
	if o != nil && o.Assignable != nil {
		return true
	}

	return false
}

// SetAssignable gets a reference to the given bool and assigns it to the Assignable field.
func (o *BTCategoryPropertyInfo) SetAssignable(v bool) {
	o.Assignable = &v
}

// GetBlobMimeType returns the BlobMimeType field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetBlobMimeType() string {
	if o == nil || o.BlobMimeType == nil {
		var ret string
		return ret
	}
	return *o.BlobMimeType
}

// GetBlobMimeTypeOk returns a tuple with the BlobMimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetBlobMimeTypeOk() (*string, bool) {
	if o == nil || o.BlobMimeType == nil {
		return nil, false
	}
	return o.BlobMimeType, true
}

// HasBlobMimeType returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasBlobMimeType() bool {
	if o != nil && o.BlobMimeType != nil {
		return true
	}

	return false
}

// SetBlobMimeType gets a reference to the given string and assigns it to the BlobMimeType field.
func (o *BTCategoryPropertyInfo) SetBlobMimeType(v string) {
	o.BlobMimeType = &v
}

// GetCategoryPropertyConfigInfo returns the CategoryPropertyConfigInfo field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetCategoryPropertyConfigInfo() BTCategoryPropertyConfigInfo {
	if o == nil || o.CategoryPropertyConfigInfo == nil {
		var ret BTCategoryPropertyConfigInfo
		return ret
	}
	return *o.CategoryPropertyConfigInfo
}

// GetCategoryPropertyConfigInfoOk returns a tuple with the CategoryPropertyConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetCategoryPropertyConfigInfoOk() (*BTCategoryPropertyConfigInfo, bool) {
	if o == nil || o.CategoryPropertyConfigInfo == nil {
		return nil, false
	}
	return o.CategoryPropertyConfigInfo, true
}

// HasCategoryPropertyConfigInfo returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasCategoryPropertyConfigInfo() bool {
	if o != nil && o.CategoryPropertyConfigInfo != nil {
		return true
	}

	return false
}

// SetCategoryPropertyConfigInfo gets a reference to the given BTCategoryPropertyConfigInfo and assigns it to the CategoryPropertyConfigInfo field.
func (o *BTCategoryPropertyInfo) SetCategoryPropertyConfigInfo(v BTCategoryPropertyConfigInfo) {
	o.CategoryPropertyConfigInfo = &v
}

// GetCategorySummaryInfoList returns the CategorySummaryInfoList field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetCategorySummaryInfoList() []BTMetadataCategorySummaryInfo {
	if o == nil || o.CategorySummaryInfoList == nil {
		var ret []BTMetadataCategorySummaryInfo
		return ret
	}
	return o.CategorySummaryInfoList
}

// GetCategorySummaryInfoListOk returns a tuple with the CategorySummaryInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetCategorySummaryInfoListOk() ([]BTMetadataCategorySummaryInfo, bool) {
	if o == nil || o.CategorySummaryInfoList == nil {
		return nil, false
	}
	return o.CategorySummaryInfoList, true
}

// HasCategorySummaryInfoList returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasCategorySummaryInfoList() bool {
	if o != nil && o.CategorySummaryInfoList != nil {
		return true
	}

	return false
}

// SetCategorySummaryInfoList gets a reference to the given []BTMetadataCategorySummaryInfo and assigns it to the CategorySummaryInfoList field.
func (o *BTCategoryPropertyInfo) SetCategorySummaryInfoList(v []BTMetadataCategorySummaryInfo) {
	o.CategorySummaryInfoList = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTCategoryPropertyInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEditableInMicroversion returns the EditableInMicroversion field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetEditableInMicroversion() bool {
	if o == nil || o.EditableInMicroversion == nil {
		var ret bool
		return ret
	}
	return *o.EditableInMicroversion
}

// GetEditableInMicroversionOk returns a tuple with the EditableInMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetEditableInMicroversionOk() (*bool, bool) {
	if o == nil || o.EditableInMicroversion == nil {
		return nil, false
	}
	return o.EditableInMicroversion, true
}

// HasEditableInMicroversion returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasEditableInMicroversion() bool {
	if o != nil && o.EditableInMicroversion != nil {
		return true
	}

	return false
}

// SetEditableInMicroversion gets a reference to the given bool and assigns it to the EditableInMicroversion field.
func (o *BTCategoryPropertyInfo) SetEditableInMicroversion(v bool) {
	o.EditableInMicroversion = &v
}

// GetEditableInVersion returns the EditableInVersion field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetEditableInVersion() bool {
	if o == nil || o.EditableInVersion == nil {
		var ret bool
		return ret
	}
	return *o.EditableInVersion
}

// GetEditableInVersionOk returns a tuple with the EditableInVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetEditableInVersionOk() (*bool, bool) {
	if o == nil || o.EditableInVersion == nil {
		return nil, false
	}
	return o.EditableInVersion, true
}

// HasEditableInVersion returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasEditableInVersion() bool {
	if o != nil && o.EditableInVersion != nil {
		return true
	}

	return false
}

// SetEditableInVersion gets a reference to the given bool and assigns it to the EditableInVersion field.
func (o *BTCategoryPropertyInfo) SetEditableInVersion(v bool) {
	o.EditableInVersion = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCategoryPropertyInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCategoryPropertyInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCategoryPropertyInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectDefName returns the ObjectDefName field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetObjectDefName() string {
	if o == nil || o.ObjectDefName == nil {
		var ret string
		return ret
	}
	return *o.ObjectDefName
}

// GetObjectDefNameOk returns a tuple with the ObjectDefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetObjectDefNameOk() (*string, bool) {
	if o == nil || o.ObjectDefName == nil {
		return nil, false
	}
	return o.ObjectDefName, true
}

// HasObjectDefName returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasObjectDefName() bool {
	if o != nil && o.ObjectDefName != nil {
		return true
	}

	return false
}

// SetObjectDefName gets a reference to the given string and assigns it to the ObjectDefName field.
func (o *BTCategoryPropertyInfo) SetObjectDefName(v string) {
	o.ObjectDefName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTCategoryPropertyInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTCategoryPropertyInfo) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetUiReadonlyInMicroversion returns the UiReadonlyInMicroversion field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetUiReadonlyInMicroversion() bool {
	if o == nil || o.UiReadonlyInMicroversion == nil {
		var ret bool
		return ret
	}
	return *o.UiReadonlyInMicroversion
}

// GetUiReadonlyInMicroversionOk returns a tuple with the UiReadonlyInMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetUiReadonlyInMicroversionOk() (*bool, bool) {
	if o == nil || o.UiReadonlyInMicroversion == nil {
		return nil, false
	}
	return o.UiReadonlyInMicroversion, true
}

// HasUiReadonlyInMicroversion returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasUiReadonlyInMicroversion() bool {
	if o != nil && o.UiReadonlyInMicroversion != nil {
		return true
	}

	return false
}

// SetUiReadonlyInMicroversion gets a reference to the given bool and assigns it to the UiReadonlyInMicroversion field.
func (o *BTCategoryPropertyInfo) SetUiReadonlyInMicroversion(v bool) {
	o.UiReadonlyInMicroversion = &v
}

// GetUiReadonlyInVersion returns the UiReadonlyInVersion field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetUiReadonlyInVersion() bool {
	if o == nil || o.UiReadonlyInVersion == nil {
		var ret bool
		return ret
	}
	return *o.UiReadonlyInVersion
}

// GetUiReadonlyInVersionOk returns a tuple with the UiReadonlyInVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetUiReadonlyInVersionOk() (*bool, bool) {
	if o == nil || o.UiReadonlyInVersion == nil {
		return nil, false
	}
	return o.UiReadonlyInVersion, true
}

// HasUiReadonlyInVersion returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasUiReadonlyInVersion() bool {
	if o != nil && o.UiReadonlyInVersion != nil {
		return true
	}

	return false
}

// SetUiReadonlyInVersion gets a reference to the given bool and assigns it to the UiReadonlyInVersion field.
func (o *BTCategoryPropertyInfo) SetUiReadonlyInVersion(v bool) {
	o.UiReadonlyInVersion = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetValueType() int32 {
	if o == nil || o.ValueType == nil {
		var ret int32
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetValueTypeOk() (*int32, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given int32 and assigns it to the ValueType field.
func (o *BTCategoryPropertyInfo) SetValueType(v int32) {
	o.ValueType = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCategoryPropertyInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCategoryPropertyInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCategoryPropertyInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTCategoryPropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Array != nil {
		toSerialize["array"] = o.Array
	}
	if o.Assignable != nil {
		toSerialize["assignable"] = o.Assignable
	}
	if o.BlobMimeType != nil {
		toSerialize["blobMimeType"] = o.BlobMimeType
	}
	if o.CategoryPropertyConfigInfo != nil {
		toSerialize["categoryPropertyConfigInfo"] = o.CategoryPropertyConfigInfo
	}
	if o.CategorySummaryInfoList != nil {
		toSerialize["categorySummaryInfoList"] = o.CategorySummaryInfoList
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EditableInMicroversion != nil {
		toSerialize["editableInMicroversion"] = o.EditableInMicroversion
	}
	if o.EditableInVersion != nil {
		toSerialize["editableInVersion"] = o.EditableInVersion
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectDefName != nil {
		toSerialize["objectDefName"] = o.ObjectDefName
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.UiReadonlyInMicroversion != nil {
		toSerialize["uiReadonlyInMicroversion"] = o.UiReadonlyInMicroversion
	}
	if o.UiReadonlyInVersion != nil {
		toSerialize["uiReadonlyInVersion"] = o.UiReadonlyInVersion
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTCategoryPropertyInfo struct {
	value *BTCategoryPropertyInfo
	isSet bool
}

func (v NullableBTCategoryPropertyInfo) Get() *BTCategoryPropertyInfo {
	return v.value
}

func (v *NullableBTCategoryPropertyInfo) Set(val *BTCategoryPropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCategoryPropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCategoryPropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCategoryPropertyInfo(val *BTCategoryPropertyInfo) *NullableBTCategoryPropertyInfo {
	return &NullableBTCategoryPropertyInfo{value: val, isSet: true}
}

func (v NullableBTCategoryPropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCategoryPropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
