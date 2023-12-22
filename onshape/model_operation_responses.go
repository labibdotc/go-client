/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28023-41481dfcfdcb
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// OperationResponses struct for OperationResponses
type OperationResponses struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Default    *ApiResponse                      `json:"default,omitempty"`
	Empty      *bool                             `json:"empty,omitempty"`
}

// NewOperationResponses instantiates a new OperationResponses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponses() *OperationResponses {
	this := OperationResponses{}
	return &this
}

// NewOperationResponsesWithDefaults instantiates a new OperationResponses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponsesWithDefaults() *OperationResponses {
	this := OperationResponses{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *OperationResponses) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponses) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *OperationResponses) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *OperationResponses) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *OperationResponses) GetDefault() ApiResponse {
	if o == nil || o.Default == nil {
		var ret ApiResponse
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponses) GetDefaultOk() (*ApiResponse, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *OperationResponses) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given ApiResponse and assigns it to the Default field.
func (o *OperationResponses) SetDefault(v ApiResponse) {
	o.Default = &v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *OperationResponses) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponses) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *OperationResponses) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *OperationResponses) SetEmpty(v bool) {
	o.Empty = &v
}

func (o OperationResponses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	return json.Marshal(toSerialize)
}

type NullableOperationResponses struct {
	value *OperationResponses
	isSet bool
}

func (v NullableOperationResponses) Get() *OperationResponses {
	return v.value
}

func (v *NullableOperationResponses) Set(val *OperationResponses) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponses) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponses(val *OperationResponses) *NullableOperationResponses {
	return &NullableOperationResponses{value: val, isSet: true}
}

func (v NullableOperationResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
