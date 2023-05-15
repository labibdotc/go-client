/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15808-38acf80dff96
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTIdTranslationInfo struct for BTIdTranslationInfo
type BTIdTranslationInfo struct {
	DocumentId                 *string                     `json:"documentId,omitempty"`
	ElementId                  *string                     `json:"elementId,omitempty"`
	Ids                        []BTIdTranslationResultInfo `json:"ids,omitempty"`
	SourceDocumentMicroversion *string                     `json:"sourceDocumentMicroversion,omitempty"`
	TargetDocumentMicroversion *string                     `json:"targetDocumentMicroversion,omitempty"`
}

// NewBTIdTranslationInfo instantiates a new BTIdTranslationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTIdTranslationInfo() *BTIdTranslationInfo {
	this := BTIdTranslationInfo{}
	return &this
}

// NewBTIdTranslationInfoWithDefaults instantiates a new BTIdTranslationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTIdTranslationInfoWithDefaults() *BTIdTranslationInfo {
	this := BTIdTranslationInfo{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTIdTranslationInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTIdTranslationInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTIdTranslationInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTIdTranslationInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTIdTranslationInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTIdTranslationInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BTIdTranslationInfo) GetIds() []BTIdTranslationResultInfo {
	if o == nil || o.Ids == nil {
		var ret []BTIdTranslationResultInfo
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationInfo) GetIdsOk() ([]BTIdTranslationResultInfo, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BTIdTranslationInfo) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []BTIdTranslationResultInfo and assigns it to the Ids field.
func (o *BTIdTranslationInfo) SetIds(v []BTIdTranslationResultInfo) {
	o.Ids = v
}

// GetSourceDocumentMicroversion returns the SourceDocumentMicroversion field value if set, zero value otherwise.
func (o *BTIdTranslationInfo) GetSourceDocumentMicroversion() string {
	if o == nil || o.SourceDocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceDocumentMicroversion
}

// GetSourceDocumentMicroversionOk returns a tuple with the SourceDocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationInfo) GetSourceDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.SourceDocumentMicroversion == nil {
		return nil, false
	}
	return o.SourceDocumentMicroversion, true
}

// HasSourceDocumentMicroversion returns a boolean if a field has been set.
func (o *BTIdTranslationInfo) HasSourceDocumentMicroversion() bool {
	if o != nil && o.SourceDocumentMicroversion != nil {
		return true
	}

	return false
}

// SetSourceDocumentMicroversion gets a reference to the given string and assigns it to the SourceDocumentMicroversion field.
func (o *BTIdTranslationInfo) SetSourceDocumentMicroversion(v string) {
	o.SourceDocumentMicroversion = &v
}

// GetTargetDocumentMicroversion returns the TargetDocumentMicroversion field value if set, zero value otherwise.
func (o *BTIdTranslationInfo) GetTargetDocumentMicroversion() string {
	if o == nil || o.TargetDocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.TargetDocumentMicroversion
}

// GetTargetDocumentMicroversionOk returns a tuple with the TargetDocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationInfo) GetTargetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.TargetDocumentMicroversion == nil {
		return nil, false
	}
	return o.TargetDocumentMicroversion, true
}

// HasTargetDocumentMicroversion returns a boolean if a field has been set.
func (o *BTIdTranslationInfo) HasTargetDocumentMicroversion() bool {
	if o != nil && o.TargetDocumentMicroversion != nil {
		return true
	}

	return false
}

// SetTargetDocumentMicroversion gets a reference to the given string and assigns it to the TargetDocumentMicroversion field.
func (o *BTIdTranslationInfo) SetTargetDocumentMicroversion(v string) {
	o.TargetDocumentMicroversion = &v
}

func (o BTIdTranslationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.SourceDocumentMicroversion != nil {
		toSerialize["sourceDocumentMicroversion"] = o.SourceDocumentMicroversion
	}
	if o.TargetDocumentMicroversion != nil {
		toSerialize["targetDocumentMicroversion"] = o.TargetDocumentMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTIdTranslationInfo struct {
	value *BTIdTranslationInfo
	isSet bool
}

func (v NullableBTIdTranslationInfo) Get() *BTIdTranslationInfo {
	return v.value
}

func (v *NullableBTIdTranslationInfo) Set(val *BTIdTranslationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTIdTranslationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTIdTranslationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTIdTranslationInfo(val *BTIdTranslationInfo) *NullableBTIdTranslationInfo {
	return &NullableBTIdTranslationInfo{value: val, isSet: true}
}

func (v NullableBTIdTranslationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTIdTranslationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
