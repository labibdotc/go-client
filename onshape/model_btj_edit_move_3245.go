/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7328-9aa102855752
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTJEditMove3245 Move an existing node from one path to another.
type BTJEditMove3245 struct {
	BtType          *string      `json:"btType,omitempty"`
	DestinationPath *BTJPath3073 `json:"destinationPath,omitempty"`
	SourcePath      *BTJPath3073 `json:"sourcePath,omitempty"`
}

// NewBTJEditMove3245 instantiates a new BTJEditMove3245 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJEditMove3245() *BTJEditMove3245 {
	this := BTJEditMove3245{}
	return &this
}

// NewBTJEditMove3245WithDefaults instantiates a new BTJEditMove3245 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJEditMove3245WithDefaults() *BTJEditMove3245 {
	this := BTJEditMove3245{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJEditMove3245) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditMove3245) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJEditMove3245) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJEditMove3245) SetBtType(v string) {
	o.BtType = &v
}

// GetDestinationPath returns the DestinationPath field value if set, zero value otherwise.
func (o *BTJEditMove3245) GetDestinationPath() BTJPath3073 {
	if o == nil || o.DestinationPath == nil {
		var ret BTJPath3073
		return ret
	}
	return *o.DestinationPath
}

// GetDestinationPathOk returns a tuple with the DestinationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditMove3245) GetDestinationPathOk() (*BTJPath3073, bool) {
	if o == nil || o.DestinationPath == nil {
		return nil, false
	}
	return o.DestinationPath, true
}

// HasDestinationPath returns a boolean if a field has been set.
func (o *BTJEditMove3245) HasDestinationPath() bool {
	if o != nil && o.DestinationPath != nil {
		return true
	}

	return false
}

// SetDestinationPath gets a reference to the given BTJPath3073 and assigns it to the DestinationPath field.
func (o *BTJEditMove3245) SetDestinationPath(v BTJPath3073) {
	o.DestinationPath = &v
}

// GetSourcePath returns the SourcePath field value if set, zero value otherwise.
func (o *BTJEditMove3245) GetSourcePath() BTJPath3073 {
	if o == nil || o.SourcePath == nil {
		var ret BTJPath3073
		return ret
	}
	return *o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditMove3245) GetSourcePathOk() (*BTJPath3073, bool) {
	if o == nil || o.SourcePath == nil {
		return nil, false
	}
	return o.SourcePath, true
}

// HasSourcePath returns a boolean if a field has been set.
func (o *BTJEditMove3245) HasSourcePath() bool {
	if o != nil && o.SourcePath != nil {
		return true
	}

	return false
}

// SetSourcePath gets a reference to the given BTJPath3073 and assigns it to the SourcePath field.
func (o *BTJEditMove3245) SetSourcePath(v BTJPath3073) {
	o.SourcePath = &v
}

func (o BTJEditMove3245) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DestinationPath != nil {
		toSerialize["destinationPath"] = o.DestinationPath
	}
	if o.SourcePath != nil {
		toSerialize["sourcePath"] = o.SourcePath
	}
	return json.Marshal(toSerialize)
}

type NullableBTJEditMove3245 struct {
	value *BTJEditMove3245
	isSet bool
}

func (v NullableBTJEditMove3245) Get() *BTJEditMove3245 {
	return v.value
}

func (v *NullableBTJEditMove3245) Set(val *BTJEditMove3245) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJEditMove3245) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJEditMove3245) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJEditMove3245(val *BTJEditMove3245) *NullableBTJEditMove3245 {
	return &NullableBTJEditMove3245{value: val, isSet: true}
}

func (v NullableBTJEditMove3245) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJEditMove3245) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
