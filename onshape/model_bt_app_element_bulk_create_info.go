/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6590-f8226b4e1789
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementBulkCreateInfo struct for BTAppElementBulkCreateInfo
type BTAppElementBulkCreateInfo struct {
	// The latest document microversion, after the edit was committed.
	DocumentMicroversionId string `json:"documentMicroversionId"`
	// The ids of the created elements.
	ElementIds []string `json:"elementIds,omitempty"`
	// The microversion ids of the created elements, at creation time.
	ElementMicroversions []string `json:"elementMicroversions,omitempty"`
	// The numeric code identifying the error that occurred, if one occurred.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// A human-readable value for the error that occurred, if one occurred.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	ErrorValue       *string `json:"errorValue,omitempty"`
}

// NewBTAppElementBulkCreateInfo instantiates a new BTAppElementBulkCreateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementBulkCreateInfo(documentMicroversionId string) *BTAppElementBulkCreateInfo {
	this := BTAppElementBulkCreateInfo{}
	this.DocumentMicroversionId = documentMicroversionId
	return &this
}

// NewBTAppElementBulkCreateInfoWithDefaults instantiates a new BTAppElementBulkCreateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementBulkCreateInfoWithDefaults() *BTAppElementBulkCreateInfo {
	this := BTAppElementBulkCreateInfo{}
	return &this
}

// GetDocumentMicroversionId returns the DocumentMicroversionId field value
func (o *BTAppElementBulkCreateInfo) GetDocumentMicroversionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentMicroversionId
}

// GetDocumentMicroversionIdOk returns a tuple with the DocumentMicroversionId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateInfo) GetDocumentMicroversionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentMicroversionId, true
}

// SetDocumentMicroversionId sets field value
func (o *BTAppElementBulkCreateInfo) SetDocumentMicroversionId(v string) {
	o.DocumentMicroversionId = v
}

// GetElementIds returns the ElementIds field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateInfo) GetElementIds() []string {
	if o == nil || o.ElementIds == nil {
		var ret []string
		return ret
	}
	return o.ElementIds
}

// GetElementIdsOk returns a tuple with the ElementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateInfo) GetElementIdsOk() ([]string, bool) {
	if o == nil || o.ElementIds == nil {
		return nil, false
	}
	return o.ElementIds, true
}

// HasElementIds returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateInfo) HasElementIds() bool {
	if o != nil && o.ElementIds != nil {
		return true
	}

	return false
}

// SetElementIds gets a reference to the given []string and assigns it to the ElementIds field.
func (o *BTAppElementBulkCreateInfo) SetElementIds(v []string) {
	o.ElementIds = v
}

// GetElementMicroversions returns the ElementMicroversions field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateInfo) GetElementMicroversions() []string {
	if o == nil || o.ElementMicroversions == nil {
		var ret []string
		return ret
	}
	return o.ElementMicroversions
}

// GetElementMicroversionsOk returns a tuple with the ElementMicroversions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateInfo) GetElementMicroversionsOk() ([]string, bool) {
	if o == nil || o.ElementMicroversions == nil {
		return nil, false
	}
	return o.ElementMicroversions, true
}

// HasElementMicroversions returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateInfo) HasElementMicroversions() bool {
	if o != nil && o.ElementMicroversions != nil {
		return true
	}

	return false
}

// SetElementMicroversions gets a reference to the given []string and assigns it to the ElementMicroversions field.
func (o *BTAppElementBulkCreateInfo) SetElementMicroversions(v []string) {
	o.ElementMicroversions = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementBulkCreateInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementBulkCreateInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateInfo) GetErrorValue() string {
	if o == nil || o.ErrorValue == nil {
		var ret string
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateInfo) GetErrorValueOk() (*string, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given string and assigns it to the ErrorValue field.
func (o *BTAppElementBulkCreateInfo) SetErrorValue(v string) {
	o.ErrorValue = &v
}

func (o BTAppElementBulkCreateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["documentMicroversionId"] = o.DocumentMicroversionId
	}
	if o.ElementIds != nil {
		toSerialize["elementIds"] = o.ElementIds
	}
	if o.ElementMicroversions != nil {
		toSerialize["elementMicroversions"] = o.ElementMicroversions
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementBulkCreateInfo struct {
	value *BTAppElementBulkCreateInfo
	isSet bool
}

func (v NullableBTAppElementBulkCreateInfo) Get() *BTAppElementBulkCreateInfo {
	return v.value
}

func (v *NullableBTAppElementBulkCreateInfo) Set(val *BTAppElementBulkCreateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementBulkCreateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementBulkCreateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementBulkCreateInfo(val *BTAppElementBulkCreateInfo) *NullableBTAppElementBulkCreateInfo {
	return &NullableBTAppElementBulkCreateInfo{value: val, isSet: true}
}

func (v NullableBTAppElementBulkCreateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementBulkCreateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
