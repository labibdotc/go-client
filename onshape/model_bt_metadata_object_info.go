/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14306-351f5b17f026
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMetadataObjectInfo - struct for BTMetadataObjectInfo
type BTMetadataObjectInfo struct {
	implBTMetadataObjectInfo interface{}
}

// BTMetadataPartInfoAsBTMetadataObjectInfo is a convenience function that returns BTMetadataPartInfo wrapped in BTMetadataObjectInfo
func (o *BTMetadataPartInfo) AsBTMetadataObjectInfo() *BTMetadataObjectInfo {
	return &BTMetadataObjectInfo{o}
}

// BTMetadataElementInfoAsBTMetadataObjectInfo is a convenience function that returns BTMetadataElementInfo wrapped in BTMetadataObjectInfo
func (o *BTMetadataElementInfo) AsBTMetadataObjectInfo() *BTMetadataObjectInfo {
	return &BTMetadataObjectInfo{o}
}

// NewBTMetadataObjectInfo instantiates a new BTMetadataObjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataObjectInfo(jsonType string) *BTMetadataObjectInfo {
	this := BTMetadataObjectInfo{Newbase_BTMetadataObjectInfo(jsonType)}
	return &this
}

// NewBTMetadataObjectInfoWithDefaults instantiates a new BTMetadataObjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataObjectInfoWithDefaults() *BTMetadataObjectInfo {
	this := BTMetadataObjectInfo{Newbase_BTMetadataObjectInfoWithDefaults()}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTMetadataObjectInfo) GetHref() string {
	type getResult interface {
		GetHref() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHref()
	} else {
		var de string
		return de
	}
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataObjectInfo) GetHrefOk() (*string, bool) {
	type getResult interface {
		GetHrefOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHrefOk()
	} else {
		return nil, false
	}
}

// HasHref returns a boolean if a field has been set.
func (o *BTMetadataObjectInfo) HasHref() bool {
	type getResult interface {
		HasHref() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasHref()
	} else {
		return false
	}
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTMetadataObjectInfo) SetHref(v string) {
	type getResult interface {
		SetHref(v string)
	}

	o.GetActualInstance().(getResult).SetHref(v)
}

// GetJsonType returns the JsonType field value
func (o *BTMetadataObjectInfo) GetJsonType() string {
	type getResult interface {
		GetJsonType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonType()
	} else {
		var de string
		return de
	}
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTMetadataObjectInfo) GetJsonTypeOk() (*string, bool) {
	type getResult interface {
		GetJsonTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonTypeOk()
	} else {
		return nil, false
	}
}

// SetJsonType sets field value
func (o *BTMetadataObjectInfo) SetJsonType(v string) {
	type getResult interface {
		SetJsonType(v string)
	}

	o.GetActualInstance().(getResult).SetJsonType(v)
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTMetadataObjectInfo) GetProperties() []BTMetadataPropertyInfo {
	type getResult interface {
		GetProperties() []BTMetadataPropertyInfo
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetProperties()
	} else {
		var de []BTMetadataPropertyInfo
		return de
	}
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataObjectInfo) GetPropertiesOk() ([]BTMetadataPropertyInfo, bool) {
	type getResult interface {
		GetPropertiesOk() ([]BTMetadataPropertyInfo, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPropertiesOk()
	} else {
		return nil, false
	}
}

// HasProperties returns a boolean if a field has been set.
func (o *BTMetadataObjectInfo) HasProperties() bool {
	type getResult interface {
		HasProperties() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasProperties()
	} else {
		return false
	}
}

// SetProperties gets a reference to the given []BTMetadataPropertyInfo and assigns it to the Properties field.
func (o *BTMetadataObjectInfo) SetProperties(v []BTMetadataPropertyInfo) {
	type getResult interface {
		SetProperties(v []BTMetadataPropertyInfo)
	}

	o.GetActualInstance().(getResult).SetProperties(v)
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *BTMetadataObjectInfo) GetThumbnail() BTThumbnailInfo {
	type getResult interface {
		GetThumbnail() BTThumbnailInfo
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetThumbnail()
	} else {
		var de BTThumbnailInfo
		return de
	}
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataObjectInfo) GetThumbnailOk() (*BTThumbnailInfo, bool) {
	type getResult interface {
		GetThumbnailOk() (*BTThumbnailInfo, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetThumbnailOk()
	} else {
		return nil, false
	}
}

// HasThumbnail returns a boolean if a field has been set.
func (o *BTMetadataObjectInfo) HasThumbnail() bool {
	type getResult interface {
		HasThumbnail() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasThumbnail()
	} else {
		return false
	}
}

// SetThumbnail gets a reference to the given BTThumbnailInfo and assigns it to the Thumbnail field.
func (o *BTMetadataObjectInfo) SetThumbnail(v BTThumbnailInfo) {
	type getResult interface {
		SetThumbnail(v BTThumbnailInfo)
	}

	o.GetActualInstance().(getResult).SetThumbnail(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMetadataObjectInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMetadataElementInfo'
	if jsonDict["jsonType"] == "BTMetadataElementInfo" {
		// try to unmarshal JSON data into BTMetadataElementInfo
		var qr *BTMetadataElementInfo
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMetadataObjectInfo = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMetadataObjectInfo = nil
			return fmt.Errorf("Failed to unmarshal BTMetadataObjectInfo as BTMetadataElementInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMetadataPartInfo'
	if jsonDict["jsonType"] == "BTMetadataPartInfo" {
		// try to unmarshal JSON data into BTMetadataPartInfo
		var qr *BTMetadataPartInfo
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMetadataObjectInfo = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMetadataObjectInfo = nil
			return fmt.Errorf("Failed to unmarshal BTMetadataObjectInfo as BTMetadataPartInfo: %s", err.Error())
		}
	}

	var qtx *base_BTMetadataObjectInfo
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMetadataObjectInfo = qtx
		return nil // data stored in dst.base_BTMetadataObjectInfo, return on the first match
	} else {
		dst.implBTMetadataObjectInfo = nil
		return fmt.Errorf("Failed to unmarshal BTMetadataObjectInfo as base_BTMetadataObjectInfo: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMetadataObjectInfo) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMetadataObjectInfo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMetadataObjectInfo
}

type NullableBTMetadataObjectInfo struct {
	value *BTMetadataObjectInfo
	isSet bool
}

func (v NullableBTMetadataObjectInfo) Get() *BTMetadataObjectInfo {
	return v.value
}

func (v *NullableBTMetadataObjectInfo) Set(val *BTMetadataObjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataObjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataObjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataObjectInfo(val *BTMetadataObjectInfo) *NullableBTMetadataObjectInfo {
	return &NullableBTMetadataObjectInfo{value: val, isSet: true}
}

func (v NullableBTMetadataObjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataObjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMetadataObjectInfo struct {
	Href       *string                  `json:"href,omitempty"`
	JsonType   string                   `json:"jsonType"`
	Properties []BTMetadataPropertyInfo `json:"properties,omitempty"`
	Thumbnail  *BTThumbnailInfo         `json:"thumbnail,omitempty"`
}

// Newbase_BTMetadataObjectInfo instantiates a new base_BTMetadataObjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMetadataObjectInfo(jsonType string) *base_BTMetadataObjectInfo {
	this := base_BTMetadataObjectInfo{}
	this.JsonType = jsonType
	return &this
}

// Newbase_BTMetadataObjectInfoWithDefaults instantiates a new base_BTMetadataObjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMetadataObjectInfoWithDefaults() *base_BTMetadataObjectInfo {
	this := base_BTMetadataObjectInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *base_BTMetadataObjectInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMetadataObjectInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *base_BTMetadataObjectInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *base_BTMetadataObjectInfo) SetHref(v string) {
	o.Href = &v
}

// GetJsonType returns the JsonType field value
func (o *base_BTMetadataObjectInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *base_BTMetadataObjectInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *base_BTMetadataObjectInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *base_BTMetadataObjectInfo) GetProperties() []BTMetadataPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTMetadataPropertyInfo
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMetadataObjectInfo) GetPropertiesOk() ([]BTMetadataPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *base_BTMetadataObjectInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTMetadataPropertyInfo and assigns it to the Properties field.
func (o *base_BTMetadataObjectInfo) SetProperties(v []BTMetadataPropertyInfo) {
	o.Properties = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *base_BTMetadataObjectInfo) GetThumbnail() BTThumbnailInfo {
	if o == nil || o.Thumbnail == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMetadataObjectInfo) GetThumbnailOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *base_BTMetadataObjectInfo) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given BTThumbnailInfo and assigns it to the Thumbnail field.
func (o *base_BTMetadataObjectInfo) SetThumbnail(v BTThumbnailInfo) {
	o.Thumbnail = &v
}

func (o base_BTMetadataObjectInfo) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}
