/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5686-86cede96e73b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExternalElementReferenceInfo struct for BTExternalElementReferenceInfo
type BTExternalElementReferenceInfo struct {
	DocumentId            *string `json:"documentId,omitempty"`
	ElementId             *string `json:"elementId,omitempty"`
	ElementMicroversionId *string `json:"elementMicroversionId,omitempty"`
	VersionId             *string `json:"versionId,omitempty"`
}

// NewBTExternalElementReferenceInfo instantiates a new BTExternalElementReferenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExternalElementReferenceInfo() *BTExternalElementReferenceInfo {
	this := BTExternalElementReferenceInfo{}
	return &this
}

// NewBTExternalElementReferenceInfoWithDefaults instantiates a new BTExternalElementReferenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExternalElementReferenceInfoWithDefaults() *BTExternalElementReferenceInfo {
	this := BTExternalElementReferenceInfo{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTExternalElementReferenceInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalElementReferenceInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTExternalElementReferenceInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTExternalElementReferenceInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTExternalElementReferenceInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalElementReferenceInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTExternalElementReferenceInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTExternalElementReferenceInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementMicroversionId returns the ElementMicroversionId field value if set, zero value otherwise.
func (o *BTExternalElementReferenceInfo) GetElementMicroversionId() string {
	if o == nil || o.ElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ElementMicroversionId
}

// GetElementMicroversionIdOk returns a tuple with the ElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalElementReferenceInfo) GetElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.ElementMicroversionId == nil {
		return nil, false
	}
	return o.ElementMicroversionId, true
}

// HasElementMicroversionId returns a boolean if a field has been set.
func (o *BTExternalElementReferenceInfo) HasElementMicroversionId() bool {
	if o != nil && o.ElementMicroversionId != nil {
		return true
	}

	return false
}

// SetElementMicroversionId gets a reference to the given string and assigns it to the ElementMicroversionId field.
func (o *BTExternalElementReferenceInfo) SetElementMicroversionId(v string) {
	o.ElementMicroversionId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTExternalElementReferenceInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalElementReferenceInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTExternalElementReferenceInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTExternalElementReferenceInfo) SetVersionId(v string) {
	o.VersionId = &v
}

func (o BTExternalElementReferenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementMicroversionId != nil {
		toSerialize["elementMicroversionId"] = o.ElementMicroversionId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTExternalElementReferenceInfo struct {
	value *BTExternalElementReferenceInfo
	isSet bool
}

func (v NullableBTExternalElementReferenceInfo) Get() *BTExternalElementReferenceInfo {
	return v.value
}

func (v *NullableBTExternalElementReferenceInfo) Set(val *BTExternalElementReferenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExternalElementReferenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExternalElementReferenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExternalElementReferenceInfo(val *BTExternalElementReferenceInfo) *NullableBTExternalElementReferenceInfo {
	return &NullableBTExternalElementReferenceInfo{value: val, isSet: true}
}

func (v NullableBTExternalElementReferenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExternalElementReferenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
