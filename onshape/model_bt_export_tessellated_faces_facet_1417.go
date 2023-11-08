/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25541-b54c5b71ef45
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportTessellatedFacesFacet1417 struct for BTExportTessellatedFacesFacet1417
type BTExportTessellatedFacesFacet1417 struct {
	BtType             *string          `json:"btType,omitempty"`
	Indices            []int32          `json:"indices,omitempty"`
	Normal             *BTVector3d389   `json:"normal,omitempty"`
	Normals            []BTVector3d389  `json:"normals,omitempty"`
	TextureCoordinates []BTVector2d1812 `json:"textureCoordinates,omitempty"`
	Vertices           []BTVector3d389  `json:"vertices,omitempty"`
}

// NewBTExportTessellatedFacesFacet1417 instantiates a new BTExportTessellatedFacesFacet1417 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedFacesFacet1417() *BTExportTessellatedFacesFacet1417 {
	this := BTExportTessellatedFacesFacet1417{}
	return &this
}

// NewBTExportTessellatedFacesFacet1417WithDefaults instantiates a new BTExportTessellatedFacesFacet1417 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedFacesFacet1417WithDefaults() *BTExportTessellatedFacesFacet1417 {
	this := BTExportTessellatedFacesFacet1417{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFacet1417) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFacet1417) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFacet1417) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportTessellatedFacesFacet1417) SetBtType(v string) {
	o.BtType = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFacet1417) GetIndices() []int32 {
	if o == nil || o.Indices == nil {
		var ret []int32
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFacet1417) GetIndicesOk() ([]int32, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFacet1417) HasIndices() bool {
	if o != nil && o.Indices != nil {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []int32 and assigns it to the Indices field.
func (o *BTExportTessellatedFacesFacet1417) SetIndices(v []int32) {
	o.Indices = v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFacet1417) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFacet1417) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFacet1417) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTExportTessellatedFacesFacet1417) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

// GetNormals returns the Normals field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFacet1417) GetNormals() []BTVector3d389 {
	if o == nil || o.Normals == nil {
		var ret []BTVector3d389
		return ret
	}
	return o.Normals
}

// GetNormalsOk returns a tuple with the Normals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFacet1417) GetNormalsOk() ([]BTVector3d389, bool) {
	if o == nil || o.Normals == nil {
		return nil, false
	}
	return o.Normals, true
}

// HasNormals returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFacet1417) HasNormals() bool {
	if o != nil && o.Normals != nil {
		return true
	}

	return false
}

// SetNormals gets a reference to the given []BTVector3d389 and assigns it to the Normals field.
func (o *BTExportTessellatedFacesFacet1417) SetNormals(v []BTVector3d389) {
	o.Normals = v
}

// GetTextureCoordinates returns the TextureCoordinates field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFacet1417) GetTextureCoordinates() []BTVector2d1812 {
	if o == nil || o.TextureCoordinates == nil {
		var ret []BTVector2d1812
		return ret
	}
	return o.TextureCoordinates
}

// GetTextureCoordinatesOk returns a tuple with the TextureCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFacet1417) GetTextureCoordinatesOk() ([]BTVector2d1812, bool) {
	if o == nil || o.TextureCoordinates == nil {
		return nil, false
	}
	return o.TextureCoordinates, true
}

// HasTextureCoordinates returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFacet1417) HasTextureCoordinates() bool {
	if o != nil && o.TextureCoordinates != nil {
		return true
	}

	return false
}

// SetTextureCoordinates gets a reference to the given []BTVector2d1812 and assigns it to the TextureCoordinates field.
func (o *BTExportTessellatedFacesFacet1417) SetTextureCoordinates(v []BTVector2d1812) {
	o.TextureCoordinates = v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFacet1417) GetVertices() []BTVector3d389 {
	if o == nil || o.Vertices == nil {
		var ret []BTVector3d389
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFacet1417) GetVerticesOk() ([]BTVector3d389, bool) {
	if o == nil || o.Vertices == nil {
		return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFacet1417) HasVertices() bool {
	if o != nil && o.Vertices != nil {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []BTVector3d389 and assigns it to the Vertices field.
func (o *BTExportTessellatedFacesFacet1417) SetVertices(v []BTVector3d389) {
	o.Vertices = v
}

func (o BTExportTessellatedFacesFacet1417) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Indices != nil {
		toSerialize["indices"] = o.Indices
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	if o.Normals != nil {
		toSerialize["normals"] = o.Normals
	}
	if o.TextureCoordinates != nil {
		toSerialize["textureCoordinates"] = o.TextureCoordinates
	}
	if o.Vertices != nil {
		toSerialize["vertices"] = o.Vertices
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedFacesFacet1417 struct {
	value *BTExportTessellatedFacesFacet1417
	isSet bool
}

func (v NullableBTExportTessellatedFacesFacet1417) Get() *BTExportTessellatedFacesFacet1417 {
	return v.value
}

func (v *NullableBTExportTessellatedFacesFacet1417) Set(val *BTExportTessellatedFacesFacet1417) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedFacesFacet1417) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedFacesFacet1417) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedFacesFacet1417(val *BTExportTessellatedFacesFacet1417) *NullableBTExportTessellatedFacesFacet1417 {
	return &NullableBTExportTessellatedFacesFacet1417{value: val, isSet: true}
}

func (v NullableBTExportTessellatedFacesFacet1417) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedFacesFacet1417) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
