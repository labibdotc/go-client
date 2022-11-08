/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7215-bae57a997b00
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ExternalAccount struct for ExternalAccount
type ExternalAccount struct {
	Account     *string            `json:"account,omitempty"`
	Customer    *string            `json:"customer,omitempty"`
	Id          *string            `json:"id,omitempty"`
	InstanceURL *string            `json:"instanceURL,omitempty"`
	Metadata    *map[string]string `json:"metadata,omitempty"`
	Object      *string            `json:"object,omitempty"`
}

// NewExternalAccount instantiates a new ExternalAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccount() *ExternalAccount {
	this := ExternalAccount{}
	return &this
}

// NewExternalAccountWithDefaults instantiates a new ExternalAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountWithDefaults() *ExternalAccount {
	this := ExternalAccount{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ExternalAccount) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ExternalAccount) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ExternalAccount) SetAccount(v string) {
	o.Account = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ExternalAccount) GetCustomer() string {
	if o == nil || o.Customer == nil {
		var ret string
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetCustomerOk() (*string, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ExternalAccount) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given string and assigns it to the Customer field.
func (o *ExternalAccount) SetCustomer(v string) {
	o.Customer = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalAccount) SetId(v string) {
	o.Id = &v
}

// GetInstanceURL returns the InstanceURL field value if set, zero value otherwise.
func (o *ExternalAccount) GetInstanceURL() string {
	if o == nil || o.InstanceURL == nil {
		var ret string
		return ret
	}
	return *o.InstanceURL
}

// GetInstanceURLOk returns a tuple with the InstanceURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetInstanceURLOk() (*string, bool) {
	if o == nil || o.InstanceURL == nil {
		return nil, false
	}
	return o.InstanceURL, true
}

// HasInstanceURL returns a boolean if a field has been set.
func (o *ExternalAccount) HasInstanceURL() bool {
	if o != nil && o.InstanceURL != nil {
		return true
	}

	return false
}

// SetInstanceURL gets a reference to the given string and assigns it to the InstanceURL field.
func (o *ExternalAccount) SetInstanceURL(v string) {
	o.InstanceURL = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ExternalAccount) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ExternalAccount) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ExternalAccount) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ExternalAccount) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ExternalAccount) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ExternalAccount) SetObject(v string) {
	o.Object = &v
}

func (o ExternalAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InstanceURL != nil {
		toSerialize["instanceURL"] = o.InstanceURL
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccount struct {
	value *ExternalAccount
	isSet bool
}

func (v NullableExternalAccount) Get() *ExternalAccount {
	return v.value
}

func (v *NullableExternalAccount) Set(val *ExternalAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccount(val *ExternalAccount) *NullableExternalAccount {
	return &NullableExternalAccount{value: val, isSet: true}
}

func (v NullableExternalAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
