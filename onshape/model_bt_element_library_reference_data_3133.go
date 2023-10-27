/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24804-920f3dc76f2b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTElementLibraryReferenceData3133 struct for BTElementLibraryReferenceData3133
type BTElementLibraryReferenceData3133 struct {
	BtType                      *string      `json:"btType,omitempty"`
	ElementLibraryId            *string      `json:"elementLibraryId,omitempty"`
	ElementLibraryIdRaw         *BTObjectId  `json:"elementLibraryIdRaw,omitempty"`
	ElementLibrarySelectionPath []BTObjectId `json:"elementLibrarySelectionPath,omitempty"`
	ElementLibraryVersion       *string      `json:"elementLibraryVersion,omitempty"`
	ElementLibraryVersionRaw    *BTObjectId  `json:"elementLibraryVersionRaw,omitempty"`
}

// NewBTElementLibraryReferenceData3133 instantiates a new BTElementLibraryReferenceData3133 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementLibraryReferenceData3133() *BTElementLibraryReferenceData3133 {
	this := BTElementLibraryReferenceData3133{}
	return &this
}

// NewBTElementLibraryReferenceData3133WithDefaults instantiates a new BTElementLibraryReferenceData3133 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementLibraryReferenceData3133WithDefaults() *BTElementLibraryReferenceData3133 {
	this := BTElementLibraryReferenceData3133{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTElementLibraryReferenceData3133) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibraryReferenceData3133) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTElementLibraryReferenceData3133) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTElementLibraryReferenceData3133) SetBtType(v string) {
	o.BtType = &v
}

// GetElementLibraryId returns the ElementLibraryId field value if set, zero value otherwise.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryId() string {
	if o == nil || o.ElementLibraryId == nil {
		var ret string
		return ret
	}
	return *o.ElementLibraryId
}

// GetElementLibraryIdOk returns a tuple with the ElementLibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryIdOk() (*string, bool) {
	if o == nil || o.ElementLibraryId == nil {
		return nil, false
	}
	return o.ElementLibraryId, true
}

// HasElementLibraryId returns a boolean if a field has been set.
func (o *BTElementLibraryReferenceData3133) HasElementLibraryId() bool {
	if o != nil && o.ElementLibraryId != nil {
		return true
	}

	return false
}

// SetElementLibraryId gets a reference to the given string and assigns it to the ElementLibraryId field.
func (o *BTElementLibraryReferenceData3133) SetElementLibraryId(v string) {
	o.ElementLibraryId = &v
}

// GetElementLibraryIdRaw returns the ElementLibraryIdRaw field value if set, zero value otherwise.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryIdRaw() BTObjectId {
	if o == nil || o.ElementLibraryIdRaw == nil {
		var ret BTObjectId
		return ret
	}
	return *o.ElementLibraryIdRaw
}

// GetElementLibraryIdRawOk returns a tuple with the ElementLibraryIdRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryIdRawOk() (*BTObjectId, bool) {
	if o == nil || o.ElementLibraryIdRaw == nil {
		return nil, false
	}
	return o.ElementLibraryIdRaw, true
}

// HasElementLibraryIdRaw returns a boolean if a field has been set.
func (o *BTElementLibraryReferenceData3133) HasElementLibraryIdRaw() bool {
	if o != nil && o.ElementLibraryIdRaw != nil {
		return true
	}

	return false
}

// SetElementLibraryIdRaw gets a reference to the given BTObjectId and assigns it to the ElementLibraryIdRaw field.
func (o *BTElementLibraryReferenceData3133) SetElementLibraryIdRaw(v BTObjectId) {
	o.ElementLibraryIdRaw = &v
}

// GetElementLibrarySelectionPath returns the ElementLibrarySelectionPath field value if set, zero value otherwise.
func (o *BTElementLibraryReferenceData3133) GetElementLibrarySelectionPath() []BTObjectId {
	if o == nil || o.ElementLibrarySelectionPath == nil {
		var ret []BTObjectId
		return ret
	}
	return o.ElementLibrarySelectionPath
}

// GetElementLibrarySelectionPathOk returns a tuple with the ElementLibrarySelectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibraryReferenceData3133) GetElementLibrarySelectionPathOk() ([]BTObjectId, bool) {
	if o == nil || o.ElementLibrarySelectionPath == nil {
		return nil, false
	}
	return o.ElementLibrarySelectionPath, true
}

// HasElementLibrarySelectionPath returns a boolean if a field has been set.
func (o *BTElementLibraryReferenceData3133) HasElementLibrarySelectionPath() bool {
	if o != nil && o.ElementLibrarySelectionPath != nil {
		return true
	}

	return false
}

// SetElementLibrarySelectionPath gets a reference to the given []BTObjectId and assigns it to the ElementLibrarySelectionPath field.
func (o *BTElementLibraryReferenceData3133) SetElementLibrarySelectionPath(v []BTObjectId) {
	o.ElementLibrarySelectionPath = v
}

// GetElementLibraryVersion returns the ElementLibraryVersion field value if set, zero value otherwise.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryVersion() string {
	if o == nil || o.ElementLibraryVersion == nil {
		var ret string
		return ret
	}
	return *o.ElementLibraryVersion
}

// GetElementLibraryVersionOk returns a tuple with the ElementLibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryVersionOk() (*string, bool) {
	if o == nil || o.ElementLibraryVersion == nil {
		return nil, false
	}
	return o.ElementLibraryVersion, true
}

// HasElementLibraryVersion returns a boolean if a field has been set.
func (o *BTElementLibraryReferenceData3133) HasElementLibraryVersion() bool {
	if o != nil && o.ElementLibraryVersion != nil {
		return true
	}

	return false
}

// SetElementLibraryVersion gets a reference to the given string and assigns it to the ElementLibraryVersion field.
func (o *BTElementLibraryReferenceData3133) SetElementLibraryVersion(v string) {
	o.ElementLibraryVersion = &v
}

// GetElementLibraryVersionRaw returns the ElementLibraryVersionRaw field value if set, zero value otherwise.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryVersionRaw() BTObjectId {
	if o == nil || o.ElementLibraryVersionRaw == nil {
		var ret BTObjectId
		return ret
	}
	return *o.ElementLibraryVersionRaw
}

// GetElementLibraryVersionRawOk returns a tuple with the ElementLibraryVersionRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLibraryReferenceData3133) GetElementLibraryVersionRawOk() (*BTObjectId, bool) {
	if o == nil || o.ElementLibraryVersionRaw == nil {
		return nil, false
	}
	return o.ElementLibraryVersionRaw, true
}

// HasElementLibraryVersionRaw returns a boolean if a field has been set.
func (o *BTElementLibraryReferenceData3133) HasElementLibraryVersionRaw() bool {
	if o != nil && o.ElementLibraryVersionRaw != nil {
		return true
	}

	return false
}

// SetElementLibraryVersionRaw gets a reference to the given BTObjectId and assigns it to the ElementLibraryVersionRaw field.
func (o *BTElementLibraryReferenceData3133) SetElementLibraryVersionRaw(v BTObjectId) {
	o.ElementLibraryVersionRaw = &v
}

func (o BTElementLibraryReferenceData3133) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ElementLibraryId != nil {
		toSerialize["elementLibraryId"] = o.ElementLibraryId
	}
	if o.ElementLibraryIdRaw != nil {
		toSerialize["elementLibraryIdRaw"] = o.ElementLibraryIdRaw
	}
	if o.ElementLibrarySelectionPath != nil {
		toSerialize["elementLibrarySelectionPath"] = o.ElementLibrarySelectionPath
	}
	if o.ElementLibraryVersion != nil {
		toSerialize["elementLibraryVersion"] = o.ElementLibraryVersion
	}
	if o.ElementLibraryVersionRaw != nil {
		toSerialize["elementLibraryVersionRaw"] = o.ElementLibraryVersionRaw
	}
	return json.Marshal(toSerialize)
}

type NullableBTElementLibraryReferenceData3133 struct {
	value *BTElementLibraryReferenceData3133
	isSet bool
}

func (v NullableBTElementLibraryReferenceData3133) Get() *BTElementLibraryReferenceData3133 {
	return v.value
}

func (v *NullableBTElementLibraryReferenceData3133) Set(val *BTElementLibraryReferenceData3133) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementLibraryReferenceData3133) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementLibraryReferenceData3133) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementLibraryReferenceData3133(val *BTElementLibraryReferenceData3133) *NullableBTElementLibraryReferenceData3133 {
	return &NullableBTElementLibraryReferenceData3133{value: val, isSet: true}
}

func (v NullableBTElementLibraryReferenceData3133) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementLibraryReferenceData3133) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
