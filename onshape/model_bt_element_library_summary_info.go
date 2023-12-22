/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28023-41481dfcfdcb
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTElementLibrarySummaryInfo Element library metadata
type BTElementLibrarySummaryInfo struct {
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// The Id of the library -- unique across Onshape.
	LibraryId *string `json:"libraryId,omitempty"`
	// The current version Id of the library.
	LibraryVersion *string `json:"libraryVersion,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// The owner Id of an element library (either Onshape, company, or user).
	OwnerId *string `json:"ownerId,omitempty"`
	// The type of library owner, Onshape, user, or company
	OwnerType *int32 `json:"ownerType,omitempty"`
	// The purpose string identifying the type of element library.
	Purpose *string `json:"purpose,omitempty"`
	// The id of the root folder of the library
	SourceFolderId *string `json:"sourceFolderId,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTElementLibrarySummaryInfo instantiates a new BTElementLibrarySummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementLibrarySummaryInfo() *BTElementLibrarySummaryInfo {
	this := BTElementLibrarySummaryInfo{}
	return &this
}

// NewBTElementLibrarySummaryInfoWithDefaults instantiates a new BTElementLibrarySummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementLibrarySummaryInfoWithDefaults() *BTElementLibrarySummaryInfo {
	this := BTElementLibrarySummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTElementLibrarySummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTElementLibrarySummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetLibraryId returns the LibraryId field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetLibraryId() string {
	if o == nil || o.LibraryId == nil {
		var ret string
		return ret
	}
	return *o.LibraryId
}

// GetLibraryIdOk returns a tuple with the LibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetLibraryIdOk() (*string, bool) {
	if o == nil || o.LibraryId == nil {
		return nil, false
	}
	return o.LibraryId, true
}

// HasLibraryId returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasLibraryId() bool {
	if o != nil && o.LibraryId != nil {
		return true
	}

	return false
}

// SetLibraryId gets a reference to the given string and assigns it to the LibraryId field.
func (o *BTElementLibrarySummaryInfo) SetLibraryId(v string) {
	o.LibraryId = &v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetLibraryVersion() string {
	if o == nil || o.LibraryVersion == nil {
		var ret string
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetLibraryVersionOk() (*string, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given string and assigns it to the LibraryVersion field.
func (o *BTElementLibrarySummaryInfo) SetLibraryVersion(v string) {
	o.LibraryVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTElementLibrarySummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTElementLibrarySummaryInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTElementLibrarySummaryInfo) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *BTElementLibrarySummaryInfo) SetPurpose(v string) {
	o.Purpose = &v
}

// GetSourceFolderId returns the SourceFolderId field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetSourceFolderId() string {
	if o == nil || o.SourceFolderId == nil {
		var ret string
		return ret
	}
	return *o.SourceFolderId
}

// GetSourceFolderIdOk returns a tuple with the SourceFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetSourceFolderIdOk() (*string, bool) {
	if o == nil || o.SourceFolderId == nil {
		return nil, false
	}
	return o.SourceFolderId, true
}

// HasSourceFolderId returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasSourceFolderId() bool {
	if o != nil && o.SourceFolderId != nil {
		return true
	}

	return false
}

// SetSourceFolderId gets a reference to the given string and assigns it to the SourceFolderId field.
func (o *BTElementLibrarySummaryInfo) SetSourceFolderId(v string) {
	o.SourceFolderId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTElementLibrarySummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibrarySummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTElementLibrarySummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTElementLibrarySummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTElementLibrarySummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LibraryId != nil {
		toSerialize["libraryId"] = o.LibraryId
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.SourceFolderId != nil {
		toSerialize["sourceFolderId"] = o.SourceFolderId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTElementLibrarySummaryInfo struct {
	value *BTElementLibrarySummaryInfo
	isSet bool
}

func (v NullableBTElementLibrarySummaryInfo) Get() *BTElementLibrarySummaryInfo {
	return v.value
}

func (v *NullableBTElementLibrarySummaryInfo) Set(val *BTElementLibrarySummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementLibrarySummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementLibrarySummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementLibrarySummaryInfo(val *BTElementLibrarySummaryInfo) *NullableBTElementLibrarySummaryInfo {
	return &NullableBTElementLibrarySummaryInfo{value: val, isSet: true}
}

func (v NullableBTElementLibrarySummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementLibrarySummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
