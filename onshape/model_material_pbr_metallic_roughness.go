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

// MaterialPbrMetallicRoughness struct for MaterialPbrMetallicRoughness
type MaterialPbrMetallicRoughness struct {
	BaseColorFactor          []float32                         `json:"baseColorFactor,omitempty"`
	BaseColorTexture         *TextureInfo                      `json:"baseColorTexture,omitempty"`
	Extensions               map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras                   *map[string]interface{}           `json:"extras,omitempty"`
	MetallicFactor           *float32                          `json:"metallicFactor,omitempty"`
	MetallicRoughnessTexture *TextureInfo                      `json:"metallicRoughnessTexture,omitempty"`
	RoughnessFactor          *float32                          `json:"roughnessFactor,omitempty"`
}

// NewMaterialPbrMetallicRoughness instantiates a new MaterialPbrMetallicRoughness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterialPbrMetallicRoughness() *MaterialPbrMetallicRoughness {
	this := MaterialPbrMetallicRoughness{}
	return &this
}

// NewMaterialPbrMetallicRoughnessWithDefaults instantiates a new MaterialPbrMetallicRoughness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialPbrMetallicRoughnessWithDefaults() *MaterialPbrMetallicRoughness {
	this := MaterialPbrMetallicRoughness{}
	return &this
}

// GetBaseColorFactor returns the BaseColorFactor field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetBaseColorFactor() []float32 {
	if o == nil || o.BaseColorFactor == nil {
		var ret []float32
		return ret
	}
	return o.BaseColorFactor
}

// GetBaseColorFactorOk returns a tuple with the BaseColorFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetBaseColorFactorOk() ([]float32, bool) {
	if o == nil || o.BaseColorFactor == nil {
		return nil, false
	}
	return o.BaseColorFactor, true
}

// HasBaseColorFactor returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasBaseColorFactor() bool {
	if o != nil && o.BaseColorFactor != nil {
		return true
	}

	return false
}

// SetBaseColorFactor gets a reference to the given []float32 and assigns it to the BaseColorFactor field.
func (o *MaterialPbrMetallicRoughness) SetBaseColorFactor(v []float32) {
	o.BaseColorFactor = v
}

// GetBaseColorTexture returns the BaseColorTexture field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetBaseColorTexture() TextureInfo {
	if o == nil || o.BaseColorTexture == nil {
		var ret TextureInfo
		return ret
	}
	return *o.BaseColorTexture
}

// GetBaseColorTextureOk returns a tuple with the BaseColorTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetBaseColorTextureOk() (*TextureInfo, bool) {
	if o == nil || o.BaseColorTexture == nil {
		return nil, false
	}
	return o.BaseColorTexture, true
}

// HasBaseColorTexture returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasBaseColorTexture() bool {
	if o != nil && o.BaseColorTexture != nil {
		return true
	}

	return false
}

// SetBaseColorTexture gets a reference to the given TextureInfo and assigns it to the BaseColorTexture field.
func (o *MaterialPbrMetallicRoughness) SetBaseColorTexture(v TextureInfo) {
	o.BaseColorTexture = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *MaterialPbrMetallicRoughness) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *MaterialPbrMetallicRoughness) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetMetallicFactor returns the MetallicFactor field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetMetallicFactor() float32 {
	if o == nil || o.MetallicFactor == nil {
		var ret float32
		return ret
	}
	return *o.MetallicFactor
}

// GetMetallicFactorOk returns a tuple with the MetallicFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetMetallicFactorOk() (*float32, bool) {
	if o == nil || o.MetallicFactor == nil {
		return nil, false
	}
	return o.MetallicFactor, true
}

// HasMetallicFactor returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasMetallicFactor() bool {
	if o != nil && o.MetallicFactor != nil {
		return true
	}

	return false
}

// SetMetallicFactor gets a reference to the given float32 and assigns it to the MetallicFactor field.
func (o *MaterialPbrMetallicRoughness) SetMetallicFactor(v float32) {
	o.MetallicFactor = &v
}

// GetMetallicRoughnessTexture returns the MetallicRoughnessTexture field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetMetallicRoughnessTexture() TextureInfo {
	if o == nil || o.MetallicRoughnessTexture == nil {
		var ret TextureInfo
		return ret
	}
	return *o.MetallicRoughnessTexture
}

// GetMetallicRoughnessTextureOk returns a tuple with the MetallicRoughnessTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetMetallicRoughnessTextureOk() (*TextureInfo, bool) {
	if o == nil || o.MetallicRoughnessTexture == nil {
		return nil, false
	}
	return o.MetallicRoughnessTexture, true
}

// HasMetallicRoughnessTexture returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasMetallicRoughnessTexture() bool {
	if o != nil && o.MetallicRoughnessTexture != nil {
		return true
	}

	return false
}

// SetMetallicRoughnessTexture gets a reference to the given TextureInfo and assigns it to the MetallicRoughnessTexture field.
func (o *MaterialPbrMetallicRoughness) SetMetallicRoughnessTexture(v TextureInfo) {
	o.MetallicRoughnessTexture = &v
}

// GetRoughnessFactor returns the RoughnessFactor field value if set, zero value otherwise.
func (o *MaterialPbrMetallicRoughness) GetRoughnessFactor() float32 {
	if o == nil || o.RoughnessFactor == nil {
		var ret float32
		return ret
	}
	return *o.RoughnessFactor
}

// GetRoughnessFactorOk returns a tuple with the RoughnessFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialPbrMetallicRoughness) GetRoughnessFactorOk() (*float32, bool) {
	if o == nil || o.RoughnessFactor == nil {
		return nil, false
	}
	return o.RoughnessFactor, true
}

// HasRoughnessFactor returns a boolean if a field has been set.
func (o *MaterialPbrMetallicRoughness) HasRoughnessFactor() bool {
	if o != nil && o.RoughnessFactor != nil {
		return true
	}

	return false
}

// SetRoughnessFactor gets a reference to the given float32 and assigns it to the RoughnessFactor field.
func (o *MaterialPbrMetallicRoughness) SetRoughnessFactor(v float32) {
	o.RoughnessFactor = &v
}

func (o MaterialPbrMetallicRoughness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseColorFactor != nil {
		toSerialize["baseColorFactor"] = o.BaseColorFactor
	}
	if o.BaseColorTexture != nil {
		toSerialize["baseColorTexture"] = o.BaseColorTexture
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.MetallicFactor != nil {
		toSerialize["metallicFactor"] = o.MetallicFactor
	}
	if o.MetallicRoughnessTexture != nil {
		toSerialize["metallicRoughnessTexture"] = o.MetallicRoughnessTexture
	}
	if o.RoughnessFactor != nil {
		toSerialize["roughnessFactor"] = o.RoughnessFactor
	}
	return json.Marshal(toSerialize)
}

type NullableMaterialPbrMetallicRoughness struct {
	value *MaterialPbrMetallicRoughness
	isSet bool
}

func (v NullableMaterialPbrMetallicRoughness) Get() *MaterialPbrMetallicRoughness {
	return v.value
}

func (v *NullableMaterialPbrMetallicRoughness) Set(val *MaterialPbrMetallicRoughness) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterialPbrMetallicRoughness) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterialPbrMetallicRoughness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterialPbrMetallicRoughness(val *MaterialPbrMetallicRoughness) *NullableMaterialPbrMetallicRoughness {
	return &NullableMaterialPbrMetallicRoughness{value: val, isSet: true}
}

func (v NullableMaterialPbrMetallicRoughness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterialPbrMetallicRoughness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
