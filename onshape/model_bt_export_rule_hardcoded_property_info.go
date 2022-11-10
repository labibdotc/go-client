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

// BTExportRuleHardcodedPropertyInfo struct for BTExportRuleHardcodedPropertyInfo
type BTExportRuleHardcodedPropertyInfo struct {
	Context     *int32  `json:"context,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ObjectTypes []int32 `json:"objectTypes,omitempty"`
}

// NewBTExportRuleHardcodedPropertyInfo instantiates a new BTExportRuleHardcodedPropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportRuleHardcodedPropertyInfo() *BTExportRuleHardcodedPropertyInfo {
	this := BTExportRuleHardcodedPropertyInfo{}
	return &this
}

// NewBTExportRuleHardcodedPropertyInfoWithDefaults instantiates a new BTExportRuleHardcodedPropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportRuleHardcodedPropertyInfoWithDefaults() *BTExportRuleHardcodedPropertyInfo {
	this := BTExportRuleHardcodedPropertyInfo{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *BTExportRuleHardcodedPropertyInfo) GetContext() int32 {
	if o == nil || o.Context == nil {
		var ret int32
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleHardcodedPropertyInfo) GetContextOk() (*int32, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *BTExportRuleHardcodedPropertyInfo) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given int32 and assigns it to the Context field.
func (o *BTExportRuleHardcodedPropertyInfo) SetContext(v int32) {
	o.Context = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportRuleHardcodedPropertyInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleHardcodedPropertyInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportRuleHardcodedPropertyInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportRuleHardcodedPropertyInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTExportRuleHardcodedPropertyInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleHardcodedPropertyInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTExportRuleHardcodedPropertyInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTExportRuleHardcodedPropertyInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *BTExportRuleHardcodedPropertyInfo) GetObjectTypes() []int32 {
	if o == nil || o.ObjectTypes == nil {
		var ret []int32
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleHardcodedPropertyInfo) GetObjectTypesOk() ([]int32, bool) {
	if o == nil || o.ObjectTypes == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *BTExportRuleHardcodedPropertyInfo) HasObjectTypes() bool {
	if o != nil && o.ObjectTypes != nil {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []int32 and assigns it to the ObjectTypes field.
func (o *BTExportRuleHardcodedPropertyInfo) SetObjectTypes(v []int32) {
	o.ObjectTypes = v
}

func (o BTExportRuleHardcodedPropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectTypes != nil {
		toSerialize["objectTypes"] = o.ObjectTypes
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportRuleHardcodedPropertyInfo struct {
	value *BTExportRuleHardcodedPropertyInfo
	isSet bool
}

func (v NullableBTExportRuleHardcodedPropertyInfo) Get() *BTExportRuleHardcodedPropertyInfo {
	return v.value
}

func (v *NullableBTExportRuleHardcodedPropertyInfo) Set(val *BTExportRuleHardcodedPropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportRuleHardcodedPropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportRuleHardcodedPropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportRuleHardcodedPropertyInfo(val *BTExportRuleHardcodedPropertyInfo) *NullableBTExportRuleHardcodedPropertyInfo {
	return &NullableBTExportRuleHardcodedPropertyInfo{value: val, isSet: true}
}

func (v NullableBTExportRuleHardcodedPropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportRuleHardcodedPropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
