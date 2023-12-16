/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ServerVariable struct for ServerVariable
type ServerVariable struct {
	Default     *string                           `json:"default,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Enum        []string                          `json:"enum,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
}

// NewServerVariable instantiates a new ServerVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerVariable() *ServerVariable {
	this := ServerVariable{}
	return &this
}

// NewServerVariableWithDefaults instantiates a new ServerVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerVariableWithDefaults() *ServerVariable {
	this := ServerVariable{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ServerVariable) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVariable) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ServerVariable) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *ServerVariable) SetDefault(v string) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerVariable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVariable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerVariable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerVariable) SetDescription(v string) {
	o.Description = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *ServerVariable) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVariable) GetEnumOk() ([]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *ServerVariable) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *ServerVariable) SetEnum(v []string) {
	o.Enum = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ServerVariable) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVariable) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ServerVariable) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *ServerVariable) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

func (o ServerVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	return json.Marshal(toSerialize)
}

type NullableServerVariable struct {
	value *ServerVariable
	isSet bool
}

func (v NullableServerVariable) Get() *ServerVariable {
	return v.value
}

func (v *NullableServerVariable) Set(val *ServerVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableServerVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableServerVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerVariable(val *ServerVariable) *NullableServerVariable {
	return &NullableServerVariable{value: val, isSet: true}
}

func (v NullableServerVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
