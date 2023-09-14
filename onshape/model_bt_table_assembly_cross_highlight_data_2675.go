/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableAssemblyCrossHighlightData2675 struct for BTTableAssemblyCrossHighlightData2675
type BTTableAssemblyCrossHighlightData2675 struct {
	AssemblyCrossHighlightItems []BTTableAssemblyCrossHighlightDataItem2659 `json:"assemblyCrossHighlightItems,omitempty"`
	BtType                      *string                                     `json:"btType,omitempty"`
}

// NewBTTableAssemblyCrossHighlightData2675 instantiates a new BTTableAssemblyCrossHighlightData2675 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableAssemblyCrossHighlightData2675() *BTTableAssemblyCrossHighlightData2675 {
	this := BTTableAssemblyCrossHighlightData2675{}
	return &this
}

// NewBTTableAssemblyCrossHighlightData2675WithDefaults instantiates a new BTTableAssemblyCrossHighlightData2675 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableAssemblyCrossHighlightData2675WithDefaults() *BTTableAssemblyCrossHighlightData2675 {
	this := BTTableAssemblyCrossHighlightData2675{}
	return &this
}

// GetAssemblyCrossHighlightItems returns the AssemblyCrossHighlightItems field value if set, zero value otherwise.
func (o *BTTableAssemblyCrossHighlightData2675) GetAssemblyCrossHighlightItems() []BTTableAssemblyCrossHighlightDataItem2659 {
	if o == nil || o.AssemblyCrossHighlightItems == nil {
		var ret []BTTableAssemblyCrossHighlightDataItem2659
		return ret
	}
	return o.AssemblyCrossHighlightItems
}

// GetAssemblyCrossHighlightItemsOk returns a tuple with the AssemblyCrossHighlightItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableAssemblyCrossHighlightData2675) GetAssemblyCrossHighlightItemsOk() ([]BTTableAssemblyCrossHighlightDataItem2659, bool) {
	if o == nil || o.AssemblyCrossHighlightItems == nil {
		return nil, false
	}
	return o.AssemblyCrossHighlightItems, true
}

// HasAssemblyCrossHighlightItems returns a boolean if a field has been set.
func (o *BTTableAssemblyCrossHighlightData2675) HasAssemblyCrossHighlightItems() bool {
	if o != nil && o.AssemblyCrossHighlightItems != nil {
		return true
	}

	return false
}

// SetAssemblyCrossHighlightItems gets a reference to the given []BTTableAssemblyCrossHighlightDataItem2659 and assigns it to the AssemblyCrossHighlightItems field.
func (o *BTTableAssemblyCrossHighlightData2675) SetAssemblyCrossHighlightItems(v []BTTableAssemblyCrossHighlightDataItem2659) {
	o.AssemblyCrossHighlightItems = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableAssemblyCrossHighlightData2675) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableAssemblyCrossHighlightData2675) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableAssemblyCrossHighlightData2675) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableAssemblyCrossHighlightData2675) SetBtType(v string) {
	o.BtType = &v
}

func (o BTTableAssemblyCrossHighlightData2675) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyCrossHighlightItems != nil {
		toSerialize["assemblyCrossHighlightItems"] = o.AssemblyCrossHighlightItems
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableAssemblyCrossHighlightData2675 struct {
	value *BTTableAssemblyCrossHighlightData2675
	isSet bool
}

func (v NullableBTTableAssemblyCrossHighlightData2675) Get() *BTTableAssemblyCrossHighlightData2675 {
	return v.value
}

func (v *NullableBTTableAssemblyCrossHighlightData2675) Set(val *BTTableAssemblyCrossHighlightData2675) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableAssemblyCrossHighlightData2675) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableAssemblyCrossHighlightData2675) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableAssemblyCrossHighlightData2675(val *BTTableAssemblyCrossHighlightData2675) *NullableBTTableAssemblyCrossHighlightData2675 {
	return &NullableBTTableAssemblyCrossHighlightData2675{value: val, isSet: true}
}

func (v NullableBTTableAssemblyCrossHighlightData2675) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableAssemblyCrossHighlightData2675) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
