/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13995-cdd961a1a6ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataElementInfo struct for BTMetadataElementInfo
type BTMetadataElementInfo struct {
	Href        *string                                     `json:"href,omitempty"`
	JsonType    string                                      `json:"jsonType"`
	Properties  []BTMetadataPropertyInfo                    `json:"properties,omitempty"`
	Thumbnail   *BTThumbnailInfo                            `json:"thumbnail,omitempty"`
	ElementId   *string                                     `json:"elementId,omitempty"`
	ElementType *int32                                      `json:"elementType,omitempty"`
	MimeType    *string                                     `json:"mimeType,omitempty"`
	Parts       *BTMetadataObjectListInfoBTMetadataPartInfo `json:"parts,omitempty"`
}

// NewBTMetadataElementInfo instantiates a new BTMetadataElementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataElementInfo(jsonType string) *BTMetadataElementInfo {
	this := BTMetadataElementInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTMetadataElementInfoWithDefaults instantiates a new BTMetadataElementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataElementInfoWithDefaults() *BTMetadataElementInfo {
	this := BTMetadataElementInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTMetadataElementInfo) SetHref(v string) {
	o.Href = &v
}

// GetJsonType returns the JsonType field value
func (o *BTMetadataElementInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTMetadataElementInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetProperties() []BTMetadataPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTMetadataPropertyInfo
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetPropertiesOk() ([]BTMetadataPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTMetadataPropertyInfo and assigns it to the Properties field.
func (o *BTMetadataElementInfo) SetProperties(v []BTMetadataPropertyInfo) {
	o.Properties = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetThumbnail() BTThumbnailInfo {
	if o == nil || o.Thumbnail == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetThumbnailOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given BTThumbnailInfo and assigns it to the Thumbnail field.
func (o *BTMetadataElementInfo) SetThumbnail(v BTThumbnailInfo) {
	o.Thumbnail = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTMetadataElementInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetElementType() int32 {
	if o == nil || o.ElementType == nil {
		var ret int32
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetElementTypeOk() (*int32, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given int32 and assigns it to the ElementType field.
func (o *BTMetadataElementInfo) SetElementType(v int32) {
	o.ElementType = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *BTMetadataElementInfo) SetMimeType(v string) {
	o.MimeType = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *BTMetadataElementInfo) GetParts() BTMetadataObjectListInfoBTMetadataPartInfo {
	if o == nil || o.Parts == nil {
		var ret BTMetadataObjectListInfoBTMetadataPartInfo
		return ret
	}
	return *o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataElementInfo) GetPartsOk() (*BTMetadataObjectListInfoBTMetadataPartInfo, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *BTMetadataElementInfo) HasParts() bool {
	if o != nil && o.Parts != nil {
		return true
	}

	return false
}

// SetParts gets a reference to the given BTMetadataObjectListInfoBTMetadataPartInfo and assigns it to the Parts field.
func (o *BTMetadataElementInfo) SetParts(v BTMetadataObjectListInfoBTMetadataPartInfo) {
	o.Parts = &v
}

func (o BTMetadataElementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Thumbnail != nil {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataElementInfo struct {
	value *BTMetadataElementInfo
	isSet bool
}

func (v NullableBTMetadataElementInfo) Get() *BTMetadataElementInfo {
	return v.value
}

func (v *NullableBTMetadataElementInfo) Set(val *BTMetadataElementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataElementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataElementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataElementInfo(val *BTMetadataElementInfo) *NullableBTMetadataElementInfo {
	return &NullableBTMetadataElementInfo{value: val, isSet: true}
}

func (v NullableBTMetadataElementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataElementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
