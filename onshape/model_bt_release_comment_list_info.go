/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18120-f464f720d215
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTReleaseCommentListInfo struct for BTReleaseCommentListInfo
type BTReleaseCommentListInfo struct {
	Comments []BTCommentInfo `json:"comments,omitempty"`
	RpId     *string         `json:"rpId,omitempty"`
	RpName   *string         `json:"rpName,omitempty"`
}

// NewBTReleaseCommentListInfo instantiates a new BTReleaseCommentListInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleaseCommentListInfo() *BTReleaseCommentListInfo {
	this := BTReleaseCommentListInfo{}
	return &this
}

// NewBTReleaseCommentListInfoWithDefaults instantiates a new BTReleaseCommentListInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleaseCommentListInfoWithDefaults() *BTReleaseCommentListInfo {
	this := BTReleaseCommentListInfo{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BTReleaseCommentListInfo) GetComments() []BTCommentInfo {
	if o == nil || o.Comments == nil {
		var ret []BTCommentInfo
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseCommentListInfo) GetCommentsOk() ([]BTCommentInfo, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BTReleaseCommentListInfo) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []BTCommentInfo and assigns it to the Comments field.
func (o *BTReleaseCommentListInfo) SetComments(v []BTCommentInfo) {
	o.Comments = v
}

// GetRpId returns the RpId field value if set, zero value otherwise.
func (o *BTReleaseCommentListInfo) GetRpId() string {
	if o == nil || o.RpId == nil {
		var ret string
		return ret
	}
	return *o.RpId
}

// GetRpIdOk returns a tuple with the RpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseCommentListInfo) GetRpIdOk() (*string, bool) {
	if o == nil || o.RpId == nil {
		return nil, false
	}
	return o.RpId, true
}

// HasRpId returns a boolean if a field has been set.
func (o *BTReleaseCommentListInfo) HasRpId() bool {
	if o != nil && o.RpId != nil {
		return true
	}

	return false
}

// SetRpId gets a reference to the given string and assigns it to the RpId field.
func (o *BTReleaseCommentListInfo) SetRpId(v string) {
	o.RpId = &v
}

// GetRpName returns the RpName field value if set, zero value otherwise.
func (o *BTReleaseCommentListInfo) GetRpName() string {
	if o == nil || o.RpName == nil {
		var ret string
		return ret
	}
	return *o.RpName
}

// GetRpNameOk returns a tuple with the RpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseCommentListInfo) GetRpNameOk() (*string, bool) {
	if o == nil || o.RpName == nil {
		return nil, false
	}
	return o.RpName, true
}

// HasRpName returns a boolean if a field has been set.
func (o *BTReleaseCommentListInfo) HasRpName() bool {
	if o != nil && o.RpName != nil {
		return true
	}

	return false
}

// SetRpName gets a reference to the given string and assigns it to the RpName field.
func (o *BTReleaseCommentListInfo) SetRpName(v string) {
	o.RpName = &v
}

func (o BTReleaseCommentListInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.RpId != nil {
		toSerialize["rpId"] = o.RpId
	}
	if o.RpName != nil {
		toSerialize["rpName"] = o.RpName
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleaseCommentListInfo struct {
	value *BTReleaseCommentListInfo
	isSet bool
}

func (v NullableBTReleaseCommentListInfo) Get() *BTReleaseCommentListInfo {
	return v.value
}

func (v *NullableBTReleaseCommentListInfo) Set(val *BTReleaseCommentListInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleaseCommentListInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleaseCommentListInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleaseCommentListInfo(val *BTReleaseCommentListInfo) *NullableBTReleaseCommentListInfo {
	return &NullableBTReleaseCommentListInfo{value: val, isSet: true}
}

func (v NullableBTReleaseCommentListInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleaseCommentListInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
