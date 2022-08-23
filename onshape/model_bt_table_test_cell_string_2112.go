/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6159-a7ef074944fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableTestCellString2112 struct for BTTableTestCellString2112
type BTTableTestCellString2112 struct {
	IsEverVisible *bool   `json:"isEverVisible,omitempty"`
	IsReadOnly    *bool   `json:"isReadOnly,omitempty"`
	BtType        *string `json:"btType,omitempty"`
	CellValue     *string `json:"cellValue,omitempty"`
}

// NewBTTableTestCellString2112 instantiates a new BTTableTestCellString2112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableTestCellString2112() *BTTableTestCellString2112 {
	this := BTTableTestCellString2112{}
	return &this
}

// NewBTTableTestCellString2112WithDefaults instantiates a new BTTableTestCellString2112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableTestCellString2112WithDefaults() *BTTableTestCellString2112 {
	this := BTTableTestCellString2112{}
	return &this
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *BTTableTestCellString2112) GetIsEverVisible() bool {
	if o == nil || o.IsEverVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsEverVisible
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellString2112) GetIsEverVisibleOk() (*bool, bool) {
	if o == nil || o.IsEverVisible == nil {
		return nil, false
	}
	return o.IsEverVisible, true
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *BTTableTestCellString2112) HasIsEverVisible() bool {
	if o != nil && o.IsEverVisible != nil {
		return true
	}

	return false
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *BTTableTestCellString2112) SetIsEverVisible(v bool) {
	o.IsEverVisible = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTTableTestCellString2112) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellString2112) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTTableTestCellString2112) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTTableTestCellString2112) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableTestCellString2112) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellString2112) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableTestCellString2112) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableTestCellString2112) SetBtType(v string) {
	o.BtType = &v
}

// GetCellValue returns the CellValue field value if set, zero value otherwise.
func (o *BTTableTestCellString2112) GetCellValue() string {
	if o == nil || o.CellValue == nil {
		var ret string
		return ret
	}
	return *o.CellValue
}

// GetCellValueOk returns a tuple with the CellValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellString2112) GetCellValueOk() (*string, bool) {
	if o == nil || o.CellValue == nil {
		return nil, false
	}
	return o.CellValue, true
}

// HasCellValue returns a boolean if a field has been set.
func (o *BTTableTestCellString2112) HasCellValue() bool {
	if o != nil && o.CellValue != nil {
		return true
	}

	return false
}

// SetCellValue gets a reference to the given string and assigns it to the CellValue field.
func (o *BTTableTestCellString2112) SetCellValue(v string) {
	o.CellValue = &v
}

func (o BTTableTestCellString2112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEverVisible != nil {
		toSerialize["isEverVisible"] = o.IsEverVisible
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CellValue != nil {
		toSerialize["cellValue"] = o.CellValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableTestCellString2112 struct {
	value *BTTableTestCellString2112
	isSet bool
}

func (v NullableBTTableTestCellString2112) Get() *BTTableTestCellString2112 {
	return v.value
}

func (v *NullableBTTableTestCellString2112) Set(val *BTTableTestCellString2112) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableTestCellString2112) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableTestCellString2112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableTestCellString2112(val *BTTableTestCellString2112) *NullableBTTableTestCellString2112 {
	return &NullableBTTableTestCellString2112{value: val, isSet: true}
}

func (v NullableBTTableTestCellString2112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableTestCellString2112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
