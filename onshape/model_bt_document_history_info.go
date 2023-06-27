/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18070-f42189da2c4f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentHistoryInfo struct for BTDocumentHistoryInfo
type BTDocumentHistoryInfo struct {
	CanBeRestored      *bool     `json:"canBeRestored,omitempty"`
	Date               *JSONTime `json:"date,omitempty"`
	Description        *string   `json:"description,omitempty"`
	MicroversionId     *string   `json:"microversionId,omitempty"`
	NextMicroversionId *string   `json:"nextMicroversionId,omitempty"`
	UserId             *string   `json:"userId,omitempty"`
	Username           *string   `json:"username,omitempty"`
}

// NewBTDocumentHistoryInfo instantiates a new BTDocumentHistoryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentHistoryInfo() *BTDocumentHistoryInfo {
	this := BTDocumentHistoryInfo{}
	return &this
}

// NewBTDocumentHistoryInfoWithDefaults instantiates a new BTDocumentHistoryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentHistoryInfoWithDefaults() *BTDocumentHistoryInfo {
	this := BTDocumentHistoryInfo{}
	return &this
}

// GetCanBeRestored returns the CanBeRestored field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetCanBeRestored() bool {
	if o == nil || o.CanBeRestored == nil {
		var ret bool
		return ret
	}
	return *o.CanBeRestored
}

// GetCanBeRestoredOk returns a tuple with the CanBeRestored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetCanBeRestoredOk() (*bool, bool) {
	if o == nil || o.CanBeRestored == nil {
		return nil, false
	}
	return o.CanBeRestored, true
}

// HasCanBeRestored returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasCanBeRestored() bool {
	if o != nil && o.CanBeRestored != nil {
		return true
	}

	return false
}

// SetCanBeRestored gets a reference to the given bool and assigns it to the CanBeRestored field.
func (o *BTDocumentHistoryInfo) SetCanBeRestored(v bool) {
	o.CanBeRestored = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetDate() JSONTime {
	if o == nil || o.Date == nil {
		var ret JSONTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetDateOk() (*JSONTime, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given JSONTime and assigns it to the Date field.
func (o *BTDocumentHistoryInfo) SetDate(v JSONTime) {
	o.Date = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTDocumentHistoryInfo) SetDescription(v string) {
	o.Description = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTDocumentHistoryInfo) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetNextMicroversionId returns the NextMicroversionId field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetNextMicroversionId() string {
	if o == nil || o.NextMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.NextMicroversionId
}

// GetNextMicroversionIdOk returns a tuple with the NextMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetNextMicroversionIdOk() (*string, bool) {
	if o == nil || o.NextMicroversionId == nil {
		return nil, false
	}
	return o.NextMicroversionId, true
}

// HasNextMicroversionId returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasNextMicroversionId() bool {
	if o != nil && o.NextMicroversionId != nil {
		return true
	}

	return false
}

// SetNextMicroversionId gets a reference to the given string and assigns it to the NextMicroversionId field.
func (o *BTDocumentHistoryInfo) SetNextMicroversionId(v string) {
	o.NextMicroversionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTDocumentHistoryInfo) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *BTDocumentHistoryInfo) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentHistoryInfo) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *BTDocumentHistoryInfo) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *BTDocumentHistoryInfo) SetUsername(v string) {
	o.Username = &v
}

func (o BTDocumentHistoryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanBeRestored != nil {
		toSerialize["canBeRestored"] = o.CanBeRestored
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.NextMicroversionId != nil {
		toSerialize["nextMicroversionId"] = o.NextMicroversionId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentHistoryInfo struct {
	value *BTDocumentHistoryInfo
	isSet bool
}

func (v NullableBTDocumentHistoryInfo) Get() *BTDocumentHistoryInfo {
	return v.value
}

func (v *NullableBTDocumentHistoryInfo) Set(val *BTDocumentHistoryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentHistoryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentHistoryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentHistoryInfo(val *BTDocumentHistoryInfo) *NullableBTDocumentHistoryInfo {
	return &NullableBTDocumentHistoryInfo{value: val, isSet: true}
}

func (v NullableBTDocumentHistoryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentHistoryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
