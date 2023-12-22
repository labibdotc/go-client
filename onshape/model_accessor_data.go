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

// AccessorData struct for AccessorData
type AccessorData struct {
	NumComponentsPerElement *int32 `json:"numComponentsPerElement,omitempty"`
	NumElements             *int32 `json:"numElements,omitempty"`
	TotalNumComponents      *int32 `json:"totalNumComponents,omitempty"`
}

// NewAccessorData instantiates a new AccessorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessorData() *AccessorData {
	this := AccessorData{}
	return &this
}

// NewAccessorDataWithDefaults instantiates a new AccessorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessorDataWithDefaults() *AccessorData {
	this := AccessorData{}
	return &this
}

// GetNumComponentsPerElement returns the NumComponentsPerElement field value if set, zero value otherwise.
func (o *AccessorData) GetNumComponentsPerElement() int32 {
	if o == nil || o.NumComponentsPerElement == nil {
		var ret int32
		return ret
	}
	return *o.NumComponentsPerElement
}

// GetNumComponentsPerElementOk returns a tuple with the NumComponentsPerElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorData) GetNumComponentsPerElementOk() (*int32, bool) {
	if o == nil || o.NumComponentsPerElement == nil {
		return nil, false
	}
	return o.NumComponentsPerElement, true
}

// HasNumComponentsPerElement returns a boolean if a field has been set.
func (o *AccessorData) HasNumComponentsPerElement() bool {
	if o != nil && o.NumComponentsPerElement != nil {
		return true
	}

	return false
}

// SetNumComponentsPerElement gets a reference to the given int32 and assigns it to the NumComponentsPerElement field.
func (o *AccessorData) SetNumComponentsPerElement(v int32) {
	o.NumComponentsPerElement = &v
}

// GetNumElements returns the NumElements field value if set, zero value otherwise.
func (o *AccessorData) GetNumElements() int32 {
	if o == nil || o.NumElements == nil {
		var ret int32
		return ret
	}
	return *o.NumElements
}

// GetNumElementsOk returns a tuple with the NumElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorData) GetNumElementsOk() (*int32, bool) {
	if o == nil || o.NumElements == nil {
		return nil, false
	}
	return o.NumElements, true
}

// HasNumElements returns a boolean if a field has been set.
func (o *AccessorData) HasNumElements() bool {
	if o != nil && o.NumElements != nil {
		return true
	}

	return false
}

// SetNumElements gets a reference to the given int32 and assigns it to the NumElements field.
func (o *AccessorData) SetNumElements(v int32) {
	o.NumElements = &v
}

// GetTotalNumComponents returns the TotalNumComponents field value if set, zero value otherwise.
func (o *AccessorData) GetTotalNumComponents() int32 {
	if o == nil || o.TotalNumComponents == nil {
		var ret int32
		return ret
	}
	return *o.TotalNumComponents
}

// GetTotalNumComponentsOk returns a tuple with the TotalNumComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorData) GetTotalNumComponentsOk() (*int32, bool) {
	if o == nil || o.TotalNumComponents == nil {
		return nil, false
	}
	return o.TotalNumComponents, true
}

// HasTotalNumComponents returns a boolean if a field has been set.
func (o *AccessorData) HasTotalNumComponents() bool {
	if o != nil && o.TotalNumComponents != nil {
		return true
	}

	return false
}

// SetTotalNumComponents gets a reference to the given int32 and assigns it to the TotalNumComponents field.
func (o *AccessorData) SetTotalNumComponents(v int32) {
	o.TotalNumComponents = &v
}

func (o AccessorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumComponentsPerElement != nil {
		toSerialize["numComponentsPerElement"] = o.NumComponentsPerElement
	}
	if o.NumElements != nil {
		toSerialize["numElements"] = o.NumElements
	}
	if o.TotalNumComponents != nil {
		toSerialize["totalNumComponents"] = o.TotalNumComponents
	}
	return json.Marshal(toSerialize)
}

type NullableAccessorData struct {
	value *AccessorData
	isSet bool
}

func (v NullableAccessorData) Get() *AccessorData {
	return v.value
}

func (v *NullableAccessorData) Set(val *AccessorData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessorData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessorData(val *AccessorData) *NullableAccessorData {
	return &NullableAccessorData{value: val, isSet: true}
}

func (v NullableAccessorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
