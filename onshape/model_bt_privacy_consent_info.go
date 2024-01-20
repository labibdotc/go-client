/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29320-74695940af99
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPrivacyConsentInfo struct for BTPrivacyConsentInfo
type BTPrivacyConsentInfo struct {
	CommunicationsOptInDate  *JSONTime `json:"communicationsOptInDate,omitempty"`
	CommunicationsOptOutDate *JSONTime `json:"communicationsOptOutDate,omitempty"`
	CommunicationsStatus     *bool     `json:"communicationsStatus,omitempty"`
	ConsentVersion           *string   `json:"consentVersion,omitempty"`
	DataProcessingOptInDate  *JSONTime `json:"dataProcessingOptInDate,omitempty"`
	DataProcessingOptOutDate *JSONTime `json:"dataProcessingOptOutDate,omitempty"`
	DataProcessingStatus     *bool     `json:"dataProcessingStatus,omitempty"`
	EulaVersion              *int64    `json:"eulaVersion,omitempty"`
	UserId                   *string   `json:"userId,omitempty"`
}

// NewBTPrivacyConsentInfo instantiates a new BTPrivacyConsentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPrivacyConsentInfo() *BTPrivacyConsentInfo {
	this := BTPrivacyConsentInfo{}
	return &this
}

// NewBTPrivacyConsentInfoWithDefaults instantiates a new BTPrivacyConsentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPrivacyConsentInfoWithDefaults() *BTPrivacyConsentInfo {
	this := BTPrivacyConsentInfo{}
	return &this
}

// GetCommunicationsOptInDate returns the CommunicationsOptInDate field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetCommunicationsOptInDate() JSONTime {
	if o == nil || o.CommunicationsOptInDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.CommunicationsOptInDate
}

// GetCommunicationsOptInDateOk returns a tuple with the CommunicationsOptInDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetCommunicationsOptInDateOk() (*JSONTime, bool) {
	if o == nil || o.CommunicationsOptInDate == nil {
		return nil, false
	}
	return o.CommunicationsOptInDate, true
}

// HasCommunicationsOptInDate returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasCommunicationsOptInDate() bool {
	if o != nil && o.CommunicationsOptInDate != nil {
		return true
	}

	return false
}

// SetCommunicationsOptInDate gets a reference to the given JSONTime and assigns it to the CommunicationsOptInDate field.
func (o *BTPrivacyConsentInfo) SetCommunicationsOptInDate(v JSONTime) {
	o.CommunicationsOptInDate = &v
}

// GetCommunicationsOptOutDate returns the CommunicationsOptOutDate field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetCommunicationsOptOutDate() JSONTime {
	if o == nil || o.CommunicationsOptOutDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.CommunicationsOptOutDate
}

// GetCommunicationsOptOutDateOk returns a tuple with the CommunicationsOptOutDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetCommunicationsOptOutDateOk() (*JSONTime, bool) {
	if o == nil || o.CommunicationsOptOutDate == nil {
		return nil, false
	}
	return o.CommunicationsOptOutDate, true
}

// HasCommunicationsOptOutDate returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasCommunicationsOptOutDate() bool {
	if o != nil && o.CommunicationsOptOutDate != nil {
		return true
	}

	return false
}

// SetCommunicationsOptOutDate gets a reference to the given JSONTime and assigns it to the CommunicationsOptOutDate field.
func (o *BTPrivacyConsentInfo) SetCommunicationsOptOutDate(v JSONTime) {
	o.CommunicationsOptOutDate = &v
}

// GetCommunicationsStatus returns the CommunicationsStatus field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetCommunicationsStatus() bool {
	if o == nil || o.CommunicationsStatus == nil {
		var ret bool
		return ret
	}
	return *o.CommunicationsStatus
}

// GetCommunicationsStatusOk returns a tuple with the CommunicationsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetCommunicationsStatusOk() (*bool, bool) {
	if o == nil || o.CommunicationsStatus == nil {
		return nil, false
	}
	return o.CommunicationsStatus, true
}

// HasCommunicationsStatus returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasCommunicationsStatus() bool {
	if o != nil && o.CommunicationsStatus != nil {
		return true
	}

	return false
}

// SetCommunicationsStatus gets a reference to the given bool and assigns it to the CommunicationsStatus field.
func (o *BTPrivacyConsentInfo) SetCommunicationsStatus(v bool) {
	o.CommunicationsStatus = &v
}

// GetConsentVersion returns the ConsentVersion field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetConsentVersion() string {
	if o == nil || o.ConsentVersion == nil {
		var ret string
		return ret
	}
	return *o.ConsentVersion
}

// GetConsentVersionOk returns a tuple with the ConsentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetConsentVersionOk() (*string, bool) {
	if o == nil || o.ConsentVersion == nil {
		return nil, false
	}
	return o.ConsentVersion, true
}

// HasConsentVersion returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasConsentVersion() bool {
	if o != nil && o.ConsentVersion != nil {
		return true
	}

	return false
}

// SetConsentVersion gets a reference to the given string and assigns it to the ConsentVersion field.
func (o *BTPrivacyConsentInfo) SetConsentVersion(v string) {
	o.ConsentVersion = &v
}

// GetDataProcessingOptInDate returns the DataProcessingOptInDate field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetDataProcessingOptInDate() JSONTime {
	if o == nil || o.DataProcessingOptInDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.DataProcessingOptInDate
}

// GetDataProcessingOptInDateOk returns a tuple with the DataProcessingOptInDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetDataProcessingOptInDateOk() (*JSONTime, bool) {
	if o == nil || o.DataProcessingOptInDate == nil {
		return nil, false
	}
	return o.DataProcessingOptInDate, true
}

// HasDataProcessingOptInDate returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasDataProcessingOptInDate() bool {
	if o != nil && o.DataProcessingOptInDate != nil {
		return true
	}

	return false
}

// SetDataProcessingOptInDate gets a reference to the given JSONTime and assigns it to the DataProcessingOptInDate field.
func (o *BTPrivacyConsentInfo) SetDataProcessingOptInDate(v JSONTime) {
	o.DataProcessingOptInDate = &v
}

// GetDataProcessingOptOutDate returns the DataProcessingOptOutDate field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetDataProcessingOptOutDate() JSONTime {
	if o == nil || o.DataProcessingOptOutDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.DataProcessingOptOutDate
}

// GetDataProcessingOptOutDateOk returns a tuple with the DataProcessingOptOutDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetDataProcessingOptOutDateOk() (*JSONTime, bool) {
	if o == nil || o.DataProcessingOptOutDate == nil {
		return nil, false
	}
	return o.DataProcessingOptOutDate, true
}

// HasDataProcessingOptOutDate returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasDataProcessingOptOutDate() bool {
	if o != nil && o.DataProcessingOptOutDate != nil {
		return true
	}

	return false
}

// SetDataProcessingOptOutDate gets a reference to the given JSONTime and assigns it to the DataProcessingOptOutDate field.
func (o *BTPrivacyConsentInfo) SetDataProcessingOptOutDate(v JSONTime) {
	o.DataProcessingOptOutDate = &v
}

// GetDataProcessingStatus returns the DataProcessingStatus field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetDataProcessingStatus() bool {
	if o == nil || o.DataProcessingStatus == nil {
		var ret bool
		return ret
	}
	return *o.DataProcessingStatus
}

// GetDataProcessingStatusOk returns a tuple with the DataProcessingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetDataProcessingStatusOk() (*bool, bool) {
	if o == nil || o.DataProcessingStatus == nil {
		return nil, false
	}
	return o.DataProcessingStatus, true
}

// HasDataProcessingStatus returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasDataProcessingStatus() bool {
	if o != nil && o.DataProcessingStatus != nil {
		return true
	}

	return false
}

// SetDataProcessingStatus gets a reference to the given bool and assigns it to the DataProcessingStatus field.
func (o *BTPrivacyConsentInfo) SetDataProcessingStatus(v bool) {
	o.DataProcessingStatus = &v
}

// GetEulaVersion returns the EulaVersion field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetEulaVersion() int64 {
	if o == nil || o.EulaVersion == nil {
		var ret int64
		return ret
	}
	return *o.EulaVersion
}

// GetEulaVersionOk returns a tuple with the EulaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetEulaVersionOk() (*int64, bool) {
	if o == nil || o.EulaVersion == nil {
		return nil, false
	}
	return o.EulaVersion, true
}

// HasEulaVersion returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasEulaVersion() bool {
	if o != nil && o.EulaVersion != nil {
		return true
	}

	return false
}

// SetEulaVersion gets a reference to the given int64 and assigns it to the EulaVersion field.
func (o *BTPrivacyConsentInfo) SetEulaVersion(v int64) {
	o.EulaVersion = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTPrivacyConsentInfo) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPrivacyConsentInfo) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTPrivacyConsentInfo) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTPrivacyConsentInfo) SetUserId(v string) {
	o.UserId = &v
}

func (o BTPrivacyConsentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommunicationsOptInDate != nil {
		toSerialize["communicationsOptInDate"] = o.CommunicationsOptInDate
	}
	if o.CommunicationsOptOutDate != nil {
		toSerialize["communicationsOptOutDate"] = o.CommunicationsOptOutDate
	}
	if o.CommunicationsStatus != nil {
		toSerialize["communicationsStatus"] = o.CommunicationsStatus
	}
	if o.ConsentVersion != nil {
		toSerialize["consentVersion"] = o.ConsentVersion
	}
	if o.DataProcessingOptInDate != nil {
		toSerialize["dataProcessingOptInDate"] = o.DataProcessingOptInDate
	}
	if o.DataProcessingOptOutDate != nil {
		toSerialize["dataProcessingOptOutDate"] = o.DataProcessingOptOutDate
	}
	if o.DataProcessingStatus != nil {
		toSerialize["dataProcessingStatus"] = o.DataProcessingStatus
	}
	if o.EulaVersion != nil {
		toSerialize["eulaVersion"] = o.EulaVersion
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableBTPrivacyConsentInfo struct {
	value *BTPrivacyConsentInfo
	isSet bool
}

func (v NullableBTPrivacyConsentInfo) Get() *BTPrivacyConsentInfo {
	return v.value
}

func (v *NullableBTPrivacyConsentInfo) Set(val *BTPrivacyConsentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPrivacyConsentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPrivacyConsentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPrivacyConsentInfo(val *BTPrivacyConsentInfo) *NullableBTPrivacyConsentInfo {
	return &NullableBTPrivacyConsentInfo{value: val, isSet: true}
}

func (v NullableBTPrivacyConsentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPrivacyConsentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
