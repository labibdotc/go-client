/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26754-ceeaad064d4a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Accessor struct for Accessor
type Accessor struct {
	BufferView    *int32                            `json:"bufferView,omitempty"`
	ByteOffset    *int32                            `json:"byteOffset,omitempty"`
	ComponentType *int32                            `json:"componentType,omitempty"`
	Count         *int32                            `json:"count,omitempty"`
	Extensions    map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras        *map[string]interface{}           `json:"extras,omitempty"`
	Max           []float32                         `json:"max,omitempty"`
	Min           []float32                         `json:"min,omitempty"`
	Name          *string                           `json:"name,omitempty"`
	Normalized    *bool                             `json:"normalized,omitempty"`
	Sparse        *AccessorSparse                   `json:"sparse,omitempty"`
	Type          *string                           `json:"type,omitempty"`
}

// NewAccessor instantiates a new Accessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessor() *Accessor {
	this := Accessor{}
	return &this
}

// NewAccessorWithDefaults instantiates a new Accessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessorWithDefaults() *Accessor {
	this := Accessor{}
	return &this
}

// GetBufferView returns the BufferView field value if set, zero value otherwise.
func (o *Accessor) GetBufferView() int32 {
	if o == nil || o.BufferView == nil {
		var ret int32
		return ret
	}
	return *o.BufferView
}

// GetBufferViewOk returns a tuple with the BufferView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetBufferViewOk() (*int32, bool) {
	if o == nil || o.BufferView == nil {
		return nil, false
	}
	return o.BufferView, true
}

// HasBufferView returns a boolean if a field has been set.
func (o *Accessor) HasBufferView() bool {
	if o != nil && o.BufferView != nil {
		return true
	}

	return false
}

// SetBufferView gets a reference to the given int32 and assigns it to the BufferView field.
func (o *Accessor) SetBufferView(v int32) {
	o.BufferView = &v
}

// GetByteOffset returns the ByteOffset field value if set, zero value otherwise.
func (o *Accessor) GetByteOffset() int32 {
	if o == nil || o.ByteOffset == nil {
		var ret int32
		return ret
	}
	return *o.ByteOffset
}

// GetByteOffsetOk returns a tuple with the ByteOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetByteOffsetOk() (*int32, bool) {
	if o == nil || o.ByteOffset == nil {
		return nil, false
	}
	return o.ByteOffset, true
}

// HasByteOffset returns a boolean if a field has been set.
func (o *Accessor) HasByteOffset() bool {
	if o != nil && o.ByteOffset != nil {
		return true
	}

	return false
}

// SetByteOffset gets a reference to the given int32 and assigns it to the ByteOffset field.
func (o *Accessor) SetByteOffset(v int32) {
	o.ByteOffset = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *Accessor) GetComponentType() int32 {
	if o == nil || o.ComponentType == nil {
		var ret int32
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetComponentTypeOk() (*int32, bool) {
	if o == nil || o.ComponentType == nil {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *Accessor) HasComponentType() bool {
	if o != nil && o.ComponentType != nil {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given int32 and assigns it to the ComponentType field.
func (o *Accessor) SetComponentType(v int32) {
	o.ComponentType = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Accessor) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Accessor) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Accessor) SetCount(v int32) {
	o.Count = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Accessor) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Accessor) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Accessor) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Accessor) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Accessor) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Accessor) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Accessor) GetMax() []float32 {
	if o == nil || o.Max == nil {
		var ret []float32
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetMaxOk() ([]float32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Accessor) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given []float32 and assigns it to the Max field.
func (o *Accessor) SetMax(v []float32) {
	o.Max = v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Accessor) GetMin() []float32 {
	if o == nil || o.Min == nil {
		var ret []float32
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetMinOk() ([]float32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Accessor) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given []float32 and assigns it to the Min field.
func (o *Accessor) SetMin(v []float32) {
	o.Min = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Accessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Accessor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Accessor) SetName(v string) {
	o.Name = &v
}

// GetNormalized returns the Normalized field value if set, zero value otherwise.
func (o *Accessor) GetNormalized() bool {
	if o == nil || o.Normalized == nil {
		var ret bool
		return ret
	}
	return *o.Normalized
}

// GetNormalizedOk returns a tuple with the Normalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetNormalizedOk() (*bool, bool) {
	if o == nil || o.Normalized == nil {
		return nil, false
	}
	return o.Normalized, true
}

// HasNormalized returns a boolean if a field has been set.
func (o *Accessor) HasNormalized() bool {
	if o != nil && o.Normalized != nil {
		return true
	}

	return false
}

// SetNormalized gets a reference to the given bool and assigns it to the Normalized field.
func (o *Accessor) SetNormalized(v bool) {
	o.Normalized = &v
}

// GetSparse returns the Sparse field value if set, zero value otherwise.
func (o *Accessor) GetSparse() AccessorSparse {
	if o == nil || o.Sparse == nil {
		var ret AccessorSparse
		return ret
	}
	return *o.Sparse
}

// GetSparseOk returns a tuple with the Sparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetSparseOk() (*AccessorSparse, bool) {
	if o == nil || o.Sparse == nil {
		return nil, false
	}
	return o.Sparse, true
}

// HasSparse returns a boolean if a field has been set.
func (o *Accessor) HasSparse() bool {
	if o != nil && o.Sparse != nil {
		return true
	}

	return false
}

// SetSparse gets a reference to the given AccessorSparse and assigns it to the Sparse field.
func (o *Accessor) SetSparse(v AccessorSparse) {
	o.Sparse = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Accessor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Accessor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Accessor) SetType(v string) {
	o.Type = &v
}

func (o Accessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BufferView != nil {
		toSerialize["bufferView"] = o.BufferView
	}
	if o.ByteOffset != nil {
		toSerialize["byteOffset"] = o.ByteOffset
	}
	if o.ComponentType != nil {
		toSerialize["componentType"] = o.ComponentType
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Normalized != nil {
		toSerialize["normalized"] = o.Normalized
	}
	if o.Sparse != nil {
		toSerialize["sparse"] = o.Sparse
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAccessor struct {
	value *Accessor
	isSet bool
}

func (v NullableAccessor) Get() *Accessor {
	return v.value
}

func (v *NullableAccessor) Set(val *Accessor) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessor) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessor(val *Accessor) *NullableAccessor {
	return &NullableAccessor{value: val, isSet: true}
}

func (v NullableAccessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
