/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7250-f7937557e62d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApiTableList1223 struct for BTApiTableList1223
type BTApiTableList1223 struct {
	Tables []BTApiTable2300 `json:"tables,omitempty"`
}

// NewBTApiTableList1223 instantiates a new BTApiTableList1223 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApiTableList1223() *BTApiTableList1223 {
	this := BTApiTableList1223{}
	return &this
}

// NewBTApiTableList1223WithDefaults instantiates a new BTApiTableList1223 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApiTableList1223WithDefaults() *BTApiTableList1223 {
	this := BTApiTableList1223{}
	return &this
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *BTApiTableList1223) GetTables() []BTApiTable2300 {
	if o == nil || o.Tables == nil {
		var ret []BTApiTable2300
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableList1223) GetTablesOk() ([]BTApiTable2300, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *BTApiTableList1223) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []BTApiTable2300 and assigns it to the Tables field.
func (o *BTApiTableList1223) SetTables(v []BTApiTable2300) {
	o.Tables = v
}

func (o BTApiTableList1223) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	return json.Marshal(toSerialize)
}

type NullableBTApiTableList1223 struct {
	value *BTApiTableList1223
	isSet bool
}

func (v NullableBTApiTableList1223) Get() *BTApiTableList1223 {
	return v.value
}

func (v *NullableBTApiTableList1223) Set(val *BTApiTableList1223) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiTableList1223) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiTableList1223) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiTableList1223(val *BTApiTableList1223) *NullableBTApiTableList1223 {
	return &NullableBTApiTableList1223{value: val, isSet: true}
}

func (v NullableBTApiTableList1223) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiTableList1223) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
