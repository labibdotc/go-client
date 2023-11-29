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

// MeshPrimitive struct for MeshPrimitive
type MeshPrimitive struct {
	Attributes *map[string]int32                 `json:"attributes,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Indices    *int32                            `json:"indices,omitempty"`
	Material   *int32                            `json:"material,omitempty"`
	Mode       *int32                            `json:"mode,omitempty"`
	Targets    []map[string]int32                `json:"targets,omitempty"`
}

// NewMeshPrimitive instantiates a new MeshPrimitive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeshPrimitive() *MeshPrimitive {
	this := MeshPrimitive{}
	return &this
}

// NewMeshPrimitiveWithDefaults instantiates a new MeshPrimitive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeshPrimitiveWithDefaults() *MeshPrimitive {
	this := MeshPrimitive{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MeshPrimitive) GetAttributes() map[string]int32 {
	if o == nil || o.Attributes == nil {
		var ret map[string]int32
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetAttributesOk() (*map[string]int32, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MeshPrimitive) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]int32 and assigns it to the Attributes field.
func (o *MeshPrimitive) SetAttributes(v map[string]int32) {
	o.Attributes = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *MeshPrimitive) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MeshPrimitive) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *MeshPrimitive) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *MeshPrimitive) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *MeshPrimitive) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *MeshPrimitive) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *MeshPrimitive) GetIndices() int32 {
	if o == nil || o.Indices == nil {
		var ret int32
		return ret
	}
	return *o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetIndicesOk() (*int32, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *MeshPrimitive) HasIndices() bool {
	if o != nil && o.Indices != nil {
		return true
	}

	return false
}

// SetIndices gets a reference to the given int32 and assigns it to the Indices field.
func (o *MeshPrimitive) SetIndices(v int32) {
	o.Indices = &v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *MeshPrimitive) GetMaterial() int32 {
	if o == nil || o.Material == nil {
		var ret int32
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetMaterialOk() (*int32, bool) {
	if o == nil || o.Material == nil {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *MeshPrimitive) HasMaterial() bool {
	if o != nil && o.Material != nil {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given int32 and assigns it to the Material field.
func (o *MeshPrimitive) SetMaterial(v int32) {
	o.Material = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *MeshPrimitive) GetMode() int32 {
	if o == nil || o.Mode == nil {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetModeOk() (*int32, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *MeshPrimitive) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *MeshPrimitive) SetMode(v int32) {
	o.Mode = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *MeshPrimitive) GetTargets() []map[string]int32 {
	if o == nil || o.Targets == nil {
		var ret []map[string]int32
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeshPrimitive) GetTargetsOk() ([]map[string]int32, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *MeshPrimitive) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []map[string]int32 and assigns it to the Targets field.
func (o *MeshPrimitive) SetTargets(v []map[string]int32) {
	o.Targets = v
}

func (o MeshPrimitive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Indices != nil {
		toSerialize["indices"] = o.Indices
	}
	if o.Material != nil {
		toSerialize["material"] = o.Material
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	return json.Marshal(toSerialize)
}

type NullableMeshPrimitive struct {
	value *MeshPrimitive
	isSet bool
}

func (v NullableMeshPrimitive) Get() *MeshPrimitive {
	return v.value
}

func (v *NullableMeshPrimitive) Set(val *MeshPrimitive) {
	v.value = val
	v.isSet = true
}

func (v NullableMeshPrimitive) IsSet() bool {
	return v.isSet
}

func (v *NullableMeshPrimitive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeshPrimitive(val *MeshPrimitive) *NullableMeshPrimitive {
	return &NullableMeshPrimitive{value: val, isSet: true}
}

func (v NullableMeshPrimitive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeshPrimitive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
