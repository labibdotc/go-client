/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.20169-88260985a0b6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsSourceInfo struct for BTBillOfMaterialsSourceInfo
type BTBillOfMaterialsSourceInfo struct {
	Document             *BTBillOfMaterialsObjectWithPropertiesInfo `json:"document,omitempty"`
	DocumentMicroversion *BTBillOfMaterialsObjectWithPropertiesInfo `json:"documentMicroversion,omitempty"`
	Element              *BTBillOfMaterialsElementInfo              `json:"element,omitempty"`
	Href                 *string                                    `json:"href,omitempty"`
	ThumbnailInfo        *BTThumbnailInfo                           `json:"thumbnailInfo,omitempty"`
	Version              *BTBillOfMaterialsObjectWithPropertiesInfo `json:"version,omitempty"`
	ViewHref             *string                                    `json:"viewHref,omitempty"`
	Workspace            *BTBillOfMaterialsObjectWithPropertiesInfo `json:"workspace,omitempty"`
}

// NewBTBillOfMaterialsSourceInfo instantiates a new BTBillOfMaterialsSourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsSourceInfo() *BTBillOfMaterialsSourceInfo {
	this := BTBillOfMaterialsSourceInfo{}
	return &this
}

// NewBTBillOfMaterialsSourceInfoWithDefaults instantiates a new BTBillOfMaterialsSourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsSourceInfoWithDefaults() *BTBillOfMaterialsSourceInfo {
	this := BTBillOfMaterialsSourceInfo{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetDocument() BTBillOfMaterialsObjectWithPropertiesInfo {
	if o == nil || o.Document == nil {
		var ret BTBillOfMaterialsObjectWithPropertiesInfo
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetDocumentOk() (*BTBillOfMaterialsObjectWithPropertiesInfo, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasDocument() bool {
	if o != nil && o.Document != nil {
		return true
	}

	return false
}

// SetDocument gets a reference to the given BTBillOfMaterialsObjectWithPropertiesInfo and assigns it to the Document field.
func (o *BTBillOfMaterialsSourceInfo) SetDocument(v BTBillOfMaterialsObjectWithPropertiesInfo) {
	o.Document = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetDocumentMicroversion() BTBillOfMaterialsObjectWithPropertiesInfo {
	if o == nil || o.DocumentMicroversion == nil {
		var ret BTBillOfMaterialsObjectWithPropertiesInfo
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetDocumentMicroversionOk() (*BTBillOfMaterialsObjectWithPropertiesInfo, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given BTBillOfMaterialsObjectWithPropertiesInfo and assigns it to the DocumentMicroversion field.
func (o *BTBillOfMaterialsSourceInfo) SetDocumentMicroversion(v BTBillOfMaterialsObjectWithPropertiesInfo) {
	o.DocumentMicroversion = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetElement() BTBillOfMaterialsElementInfo {
	if o == nil || o.Element == nil {
		var ret BTBillOfMaterialsElementInfo
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetElementOk() (*BTBillOfMaterialsElementInfo, bool) {
	if o == nil || o.Element == nil {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasElement() bool {
	if o != nil && o.Element != nil {
		return true
	}

	return false
}

// SetElement gets a reference to the given BTBillOfMaterialsElementInfo and assigns it to the Element field.
func (o *BTBillOfMaterialsSourceInfo) SetElement(v BTBillOfMaterialsElementInfo) {
	o.Element = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsSourceInfo) SetHref(v string) {
	o.Href = &v
}

// GetThumbnailInfo returns the ThumbnailInfo field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetThumbnailInfo() BTThumbnailInfo {
	if o == nil || o.ThumbnailInfo == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.ThumbnailInfo
}

// GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.ThumbnailInfo == nil {
		return nil, false
	}
	return o.ThumbnailInfo, true
}

// HasThumbnailInfo returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasThumbnailInfo() bool {
	if o != nil && o.ThumbnailInfo != nil {
		return true
	}

	return false
}

// SetThumbnailInfo gets a reference to the given BTThumbnailInfo and assigns it to the ThumbnailInfo field.
func (o *BTBillOfMaterialsSourceInfo) SetThumbnailInfo(v BTThumbnailInfo) {
	o.ThumbnailInfo = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetVersion() BTBillOfMaterialsObjectWithPropertiesInfo {
	if o == nil || o.Version == nil {
		var ret BTBillOfMaterialsObjectWithPropertiesInfo
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetVersionOk() (*BTBillOfMaterialsObjectWithPropertiesInfo, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given BTBillOfMaterialsObjectWithPropertiesInfo and assigns it to the Version field.
func (o *BTBillOfMaterialsSourceInfo) SetVersion(v BTBillOfMaterialsObjectWithPropertiesInfo) {
	o.Version = &v
}

// GetViewHref returns the ViewHref field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetViewHref() string {
	if o == nil || o.ViewHref == nil {
		var ret string
		return ret
	}
	return *o.ViewHref
}

// GetViewHrefOk returns a tuple with the ViewHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetViewHrefOk() (*string, bool) {
	if o == nil || o.ViewHref == nil {
		return nil, false
	}
	return o.ViewHref, true
}

// HasViewHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasViewHref() bool {
	if o != nil && o.ViewHref != nil {
		return true
	}

	return false
}

// SetViewHref gets a reference to the given string and assigns it to the ViewHref field.
func (o *BTBillOfMaterialsSourceInfo) SetViewHref(v string) {
	o.ViewHref = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *BTBillOfMaterialsSourceInfo) GetWorkspace() BTBillOfMaterialsObjectWithPropertiesInfo {
	if o == nil || o.Workspace == nil {
		var ret BTBillOfMaterialsObjectWithPropertiesInfo
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsSourceInfo) GetWorkspaceOk() (*BTBillOfMaterialsObjectWithPropertiesInfo, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *BTBillOfMaterialsSourceInfo) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given BTBillOfMaterialsObjectWithPropertiesInfo and assigns it to the Workspace field.
func (o *BTBillOfMaterialsSourceInfo) SetWorkspace(v BTBillOfMaterialsObjectWithPropertiesInfo) {
	o.Workspace = &v
}

func (o BTBillOfMaterialsSourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Document != nil {
		toSerialize["document"] = o.Document
	}
	if o.DocumentMicroversion != nil {
		toSerialize["documentMicroversion"] = o.DocumentMicroversion
	}
	if o.Element != nil {
		toSerialize["element"] = o.Element
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.ThumbnailInfo != nil {
		toSerialize["thumbnailInfo"] = o.ThumbnailInfo
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ViewHref != nil {
		toSerialize["viewHref"] = o.ViewHref
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsSourceInfo struct {
	value *BTBillOfMaterialsSourceInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsSourceInfo) Get() *BTBillOfMaterialsSourceInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsSourceInfo) Set(val *BTBillOfMaterialsSourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsSourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsSourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsSourceInfo(val *BTBillOfMaterialsSourceInfo) *NullableBTBillOfMaterialsSourceInfo {
	return &NullableBTBillOfMaterialsSourceInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsSourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsSourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
