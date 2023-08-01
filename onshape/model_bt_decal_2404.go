/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19740-5e8d8b0919a8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDecal2404 struct for BTDecal2404
type BTDecal2404 struct {
	BtType        *string              `json:"btType,omitempty"`
	ImageSourceId *string              `json:"imageSourceId,omitempty"`
	IsDeletion    *bool                `json:"isDeletion,omitempty"`
	Mappings      []BTImageMapping3821 `json:"mappings,omitempty"`
}

// NewBTDecal2404 instantiates a new BTDecal2404 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDecal2404() *BTDecal2404 {
	this := BTDecal2404{}
	return &this
}

// NewBTDecal2404WithDefaults instantiates a new BTDecal2404 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDecal2404WithDefaults() *BTDecal2404 {
	this := BTDecal2404{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDecal2404) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDecal2404) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDecal2404) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDecal2404) SetBtType(v string) {
	o.BtType = &v
}

// GetImageSourceId returns the ImageSourceId field value if set, zero value otherwise.
func (o *BTDecal2404) GetImageSourceId() string {
	if o == nil || o.ImageSourceId == nil {
		var ret string
		return ret
	}
	return *o.ImageSourceId
}

// GetImageSourceIdOk returns a tuple with the ImageSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDecal2404) GetImageSourceIdOk() (*string, bool) {
	if o == nil || o.ImageSourceId == nil {
		return nil, false
	}
	return o.ImageSourceId, true
}

// HasImageSourceId returns a boolean if a field has been set.
func (o *BTDecal2404) HasImageSourceId() bool {
	if o != nil && o.ImageSourceId != nil {
		return true
	}

	return false
}

// SetImageSourceId gets a reference to the given string and assigns it to the ImageSourceId field.
func (o *BTDecal2404) SetImageSourceId(v string) {
	o.ImageSourceId = &v
}

// GetIsDeletion returns the IsDeletion field value if set, zero value otherwise.
func (o *BTDecal2404) GetIsDeletion() bool {
	if o == nil || o.IsDeletion == nil {
		var ret bool
		return ret
	}
	return *o.IsDeletion
}

// GetIsDeletionOk returns a tuple with the IsDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDecal2404) GetIsDeletionOk() (*bool, bool) {
	if o == nil || o.IsDeletion == nil {
		return nil, false
	}
	return o.IsDeletion, true
}

// HasIsDeletion returns a boolean if a field has been set.
func (o *BTDecal2404) HasIsDeletion() bool {
	if o != nil && o.IsDeletion != nil {
		return true
	}

	return false
}

// SetIsDeletion gets a reference to the given bool and assigns it to the IsDeletion field.
func (o *BTDecal2404) SetIsDeletion(v bool) {
	o.IsDeletion = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *BTDecal2404) GetMappings() []BTImageMapping3821 {
	if o == nil || o.Mappings == nil {
		var ret []BTImageMapping3821
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDecal2404) GetMappingsOk() ([]BTImageMapping3821, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *BTDecal2404) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []BTImageMapping3821 and assigns it to the Mappings field.
func (o *BTDecal2404) SetMappings(v []BTImageMapping3821) {
	o.Mappings = v
}

func (o BTDecal2404) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImageSourceId != nil {
		toSerialize["imageSourceId"] = o.ImageSourceId
	}
	if o.IsDeletion != nil {
		toSerialize["isDeletion"] = o.IsDeletion
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableBTDecal2404 struct {
	value *BTDecal2404
	isSet bool
}

func (v NullableBTDecal2404) Get() *BTDecal2404 {
	return v.value
}

func (v *NullableBTDecal2404) Set(val *BTDecal2404) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDecal2404) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDecal2404) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDecal2404(val *BTDecal2404) *NullableBTDecal2404 {
	return &NullableBTDecal2404{value: val, isSet: true}
}

func (v NullableBTDecal2404) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDecal2404) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
