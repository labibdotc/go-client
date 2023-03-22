/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13258-6b30d30337fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableColumnSpec1967 struct for BTTableColumnSpec1967
type BTTableColumnSpec1967 struct {
	DefaultCellSpec         *BTParameterSpec6 `json:"defaultCellSpec,omitempty"`
	DefaultColumnWidthUnits *string           `json:"defaultColumnWidthUnits,omitempty"`
	DefaultColumnWidthValue *int32            `json:"defaultColumnWidthValue,omitempty"`
	DefaultHeaderName       *string           `json:"defaultHeaderName,omitempty"`
	DefaultTextAlignment    *string           `json:"defaultTextAlignment,omitempty"`
	ReadOnly                *bool             `json:"readOnly,omitempty"`
}

// NewBTTableColumnSpec1967 instantiates a new BTTableColumnSpec1967 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableColumnSpec1967() *BTTableColumnSpec1967 {
	this := BTTableColumnSpec1967{}
	return &this
}

// NewBTTableColumnSpec1967WithDefaults instantiates a new BTTableColumnSpec1967 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableColumnSpec1967WithDefaults() *BTTableColumnSpec1967 {
	this := BTTableColumnSpec1967{}
	return &this
}

// GetDefaultCellSpec returns the DefaultCellSpec field value if set, zero value otherwise.
func (o *BTTableColumnSpec1967) GetDefaultCellSpec() BTParameterSpec6 {
	if o == nil || o.DefaultCellSpec == nil {
		var ret BTParameterSpec6
		return ret
	}
	return *o.DefaultCellSpec
}

// GetDefaultCellSpecOk returns a tuple with the DefaultCellSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnSpec1967) GetDefaultCellSpecOk() (*BTParameterSpec6, bool) {
	if o == nil || o.DefaultCellSpec == nil {
		return nil, false
	}
	return o.DefaultCellSpec, true
}

// HasDefaultCellSpec returns a boolean if a field has been set.
func (o *BTTableColumnSpec1967) HasDefaultCellSpec() bool {
	if o != nil && o.DefaultCellSpec != nil {
		return true
	}

	return false
}

// SetDefaultCellSpec gets a reference to the given BTParameterSpec6 and assigns it to the DefaultCellSpec field.
func (o *BTTableColumnSpec1967) SetDefaultCellSpec(v BTParameterSpec6) {
	o.DefaultCellSpec = &v
}

// GetDefaultColumnWidthUnits returns the DefaultColumnWidthUnits field value if set, zero value otherwise.
func (o *BTTableColumnSpec1967) GetDefaultColumnWidthUnits() string {
	if o == nil || o.DefaultColumnWidthUnits == nil {
		var ret string
		return ret
	}
	return *o.DefaultColumnWidthUnits
}

// GetDefaultColumnWidthUnitsOk returns a tuple with the DefaultColumnWidthUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnSpec1967) GetDefaultColumnWidthUnitsOk() (*string, bool) {
	if o == nil || o.DefaultColumnWidthUnits == nil {
		return nil, false
	}
	return o.DefaultColumnWidthUnits, true
}

// HasDefaultColumnWidthUnits returns a boolean if a field has been set.
func (o *BTTableColumnSpec1967) HasDefaultColumnWidthUnits() bool {
	if o != nil && o.DefaultColumnWidthUnits != nil {
		return true
	}

	return false
}

// SetDefaultColumnWidthUnits gets a reference to the given string and assigns it to the DefaultColumnWidthUnits field.
func (o *BTTableColumnSpec1967) SetDefaultColumnWidthUnits(v string) {
	o.DefaultColumnWidthUnits = &v
}

// GetDefaultColumnWidthValue returns the DefaultColumnWidthValue field value if set, zero value otherwise.
func (o *BTTableColumnSpec1967) GetDefaultColumnWidthValue() int32 {
	if o == nil || o.DefaultColumnWidthValue == nil {
		var ret int32
		return ret
	}
	return *o.DefaultColumnWidthValue
}

// GetDefaultColumnWidthValueOk returns a tuple with the DefaultColumnWidthValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnSpec1967) GetDefaultColumnWidthValueOk() (*int32, bool) {
	if o == nil || o.DefaultColumnWidthValue == nil {
		return nil, false
	}
	return o.DefaultColumnWidthValue, true
}

// HasDefaultColumnWidthValue returns a boolean if a field has been set.
func (o *BTTableColumnSpec1967) HasDefaultColumnWidthValue() bool {
	if o != nil && o.DefaultColumnWidthValue != nil {
		return true
	}

	return false
}

// SetDefaultColumnWidthValue gets a reference to the given int32 and assigns it to the DefaultColumnWidthValue field.
func (o *BTTableColumnSpec1967) SetDefaultColumnWidthValue(v int32) {
	o.DefaultColumnWidthValue = &v
}

// GetDefaultHeaderName returns the DefaultHeaderName field value if set, zero value otherwise.
func (o *BTTableColumnSpec1967) GetDefaultHeaderName() string {
	if o == nil || o.DefaultHeaderName == nil {
		var ret string
		return ret
	}
	return *o.DefaultHeaderName
}

// GetDefaultHeaderNameOk returns a tuple with the DefaultHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnSpec1967) GetDefaultHeaderNameOk() (*string, bool) {
	if o == nil || o.DefaultHeaderName == nil {
		return nil, false
	}
	return o.DefaultHeaderName, true
}

// HasDefaultHeaderName returns a boolean if a field has been set.
func (o *BTTableColumnSpec1967) HasDefaultHeaderName() bool {
	if o != nil && o.DefaultHeaderName != nil {
		return true
	}

	return false
}

// SetDefaultHeaderName gets a reference to the given string and assigns it to the DefaultHeaderName field.
func (o *BTTableColumnSpec1967) SetDefaultHeaderName(v string) {
	o.DefaultHeaderName = &v
}

// GetDefaultTextAlignment returns the DefaultTextAlignment field value if set, zero value otherwise.
func (o *BTTableColumnSpec1967) GetDefaultTextAlignment() string {
	if o == nil || o.DefaultTextAlignment == nil {
		var ret string
		return ret
	}
	return *o.DefaultTextAlignment
}

// GetDefaultTextAlignmentOk returns a tuple with the DefaultTextAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnSpec1967) GetDefaultTextAlignmentOk() (*string, bool) {
	if o == nil || o.DefaultTextAlignment == nil {
		return nil, false
	}
	return o.DefaultTextAlignment, true
}

// HasDefaultTextAlignment returns a boolean if a field has been set.
func (o *BTTableColumnSpec1967) HasDefaultTextAlignment() bool {
	if o != nil && o.DefaultTextAlignment != nil {
		return true
	}

	return false
}

// SetDefaultTextAlignment gets a reference to the given string and assigns it to the DefaultTextAlignment field.
func (o *BTTableColumnSpec1967) SetDefaultTextAlignment(v string) {
	o.DefaultTextAlignment = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BTTableColumnSpec1967) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnSpec1967) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BTTableColumnSpec1967) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BTTableColumnSpec1967) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o BTTableColumnSpec1967) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultCellSpec != nil {
		toSerialize["defaultCellSpec"] = o.DefaultCellSpec
	}
	if o.DefaultColumnWidthUnits != nil {
		toSerialize["defaultColumnWidthUnits"] = o.DefaultColumnWidthUnits
	}
	if o.DefaultColumnWidthValue != nil {
		toSerialize["defaultColumnWidthValue"] = o.DefaultColumnWidthValue
	}
	if o.DefaultHeaderName != nil {
		toSerialize["defaultHeaderName"] = o.DefaultHeaderName
	}
	if o.DefaultTextAlignment != nil {
		toSerialize["defaultTextAlignment"] = o.DefaultTextAlignment
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableColumnSpec1967 struct {
	value *BTTableColumnSpec1967
	isSet bool
}

func (v NullableBTTableColumnSpec1967) Get() *BTTableColumnSpec1967 {
	return v.value
}

func (v *NullableBTTableColumnSpec1967) Set(val *BTTableColumnSpec1967) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableColumnSpec1967) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableColumnSpec1967) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableColumnSpec1967(val *BTTableColumnSpec1967) *NullableBTTableColumnSpec1967 {
	return &NullableBTTableColumnSpec1967{value: val, isSet: true}
}

func (v NullableBTTableColumnSpec1967) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableColumnSpec1967) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
