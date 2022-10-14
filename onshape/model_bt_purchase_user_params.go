/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6905-ae59ed040327
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPurchaseUserParams struct for BTPurchaseUserParams
type BTPurchaseUserParams struct {
	ConsumedQuantity *int32  `json:"consumedQuantity,omitempty"`
	PurchaseId       *string `json:"purchaseId,omitempty"`
	UserId           *string `json:"userId,omitempty"`
}

// NewBTPurchaseUserParams instantiates a new BTPurchaseUserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPurchaseUserParams() *BTPurchaseUserParams {
	this := BTPurchaseUserParams{}
	return &this
}

// NewBTPurchaseUserParamsWithDefaults instantiates a new BTPurchaseUserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPurchaseUserParamsWithDefaults() *BTPurchaseUserParams {
	this := BTPurchaseUserParams{}
	return &this
}

// GetConsumedQuantity returns the ConsumedQuantity field value if set, zero value otherwise.
func (o *BTPurchaseUserParams) GetConsumedQuantity() int32 {
	if o == nil || o.ConsumedQuantity == nil {
		var ret int32
		return ret
	}
	return *o.ConsumedQuantity
}

// GetConsumedQuantityOk returns a tuple with the ConsumedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPurchaseUserParams) GetConsumedQuantityOk() (*int32, bool) {
	if o == nil || o.ConsumedQuantity == nil {
		return nil, false
	}
	return o.ConsumedQuantity, true
}

// HasConsumedQuantity returns a boolean if a field has been set.
func (o *BTPurchaseUserParams) HasConsumedQuantity() bool {
	if o != nil && o.ConsumedQuantity != nil {
		return true
	}

	return false
}

// SetConsumedQuantity gets a reference to the given int32 and assigns it to the ConsumedQuantity field.
func (o *BTPurchaseUserParams) SetConsumedQuantity(v int32) {
	o.ConsumedQuantity = &v
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *BTPurchaseUserParams) GetPurchaseId() string {
	if o == nil || o.PurchaseId == nil {
		var ret string
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPurchaseUserParams) GetPurchaseIdOk() (*string, bool) {
	if o == nil || o.PurchaseId == nil {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *BTPurchaseUserParams) HasPurchaseId() bool {
	if o != nil && o.PurchaseId != nil {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given string and assigns it to the PurchaseId field.
func (o *BTPurchaseUserParams) SetPurchaseId(v string) {
	o.PurchaseId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTPurchaseUserParams) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPurchaseUserParams) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTPurchaseUserParams) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTPurchaseUserParams) SetUserId(v string) {
	o.UserId = &v
}

func (o BTPurchaseUserParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumedQuantity != nil {
		toSerialize["consumedQuantity"] = o.ConsumedQuantity
	}
	if o.PurchaseId != nil {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableBTPurchaseUserParams struct {
	value *BTPurchaseUserParams
	isSet bool
}

func (v NullableBTPurchaseUserParams) Get() *BTPurchaseUserParams {
	return v.value
}

func (v *NullableBTPurchaseUserParams) Set(val *BTPurchaseUserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPurchaseUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPurchaseUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPurchaseUserParams(val *BTPurchaseUserParams) *NullableBTPurchaseUserParams {
	return &NullableBTPurchaseUserParams{value: val, isSet: true}
}

func (v NullableBTPurchaseUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPurchaseUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
