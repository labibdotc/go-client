/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTInsertableSketchDisplayData3775 struct for BTInsertableSketchDisplayData3775
type BTInsertableSketchDisplayData3775 struct {
	BtType                   *string                                     `json:"btType,omitempty"`
	FullElementId            *BTFullElementId756                         `json:"fullElementId,omitempty"`
	GraphicsBuffers          *map[string]map[string]BTGraphicsBuffer2668 `json:"graphicsBuffers,omitempty"`
	Id                       *string                                     `json:"id,omitempty"`
	Part                     *bool                                       `json:"part,omitempty"`
	SketchFeature            *bool                                       `json:"sketchFeature,omitempty"`
	TessellationSettingIndex *int32                                      `json:"tessellationSettingIndex,omitempty"`
	BodyIdToPartData         *map[string]BTPartData16                    `json:"bodyIdToPartData,omitempty"`
	SketchFeatureId          *string                                     `json:"sketchFeatureId,omitempty"`
}

// NewBTInsertableSketchDisplayData3775 instantiates a new BTInsertableSketchDisplayData3775 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInsertableSketchDisplayData3775() *BTInsertableSketchDisplayData3775 {
	this := BTInsertableSketchDisplayData3775{}
	return &this
}

// NewBTInsertableSketchDisplayData3775WithDefaults instantiates a new BTInsertableSketchDisplayData3775 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInsertableSketchDisplayData3775WithDefaults() *BTInsertableSketchDisplayData3775 {
	this := BTInsertableSketchDisplayData3775{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInsertableSketchDisplayData3775) SetBtType(v string) {
	o.BtType = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTInsertableSketchDisplayData3775) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetGraphicsBuffers returns the GraphicsBuffers field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetGraphicsBuffers() map[string]map[string]BTGraphicsBuffer2668 {
	if o == nil || o.GraphicsBuffers == nil {
		var ret map[string]map[string]BTGraphicsBuffer2668
		return ret
	}
	return *o.GraphicsBuffers
}

// GetGraphicsBuffersOk returns a tuple with the GraphicsBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetGraphicsBuffersOk() (*map[string]map[string]BTGraphicsBuffer2668, bool) {
	if o == nil || o.GraphicsBuffers == nil {
		return nil, false
	}
	return o.GraphicsBuffers, true
}

// HasGraphicsBuffers returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasGraphicsBuffers() bool {
	if o != nil && o.GraphicsBuffers != nil {
		return true
	}

	return false
}

// SetGraphicsBuffers gets a reference to the given map[string]map[string]BTGraphicsBuffer2668 and assigns it to the GraphicsBuffers field.
func (o *BTInsertableSketchDisplayData3775) SetGraphicsBuffers(v map[string]map[string]BTGraphicsBuffer2668) {
	o.GraphicsBuffers = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTInsertableSketchDisplayData3775) SetId(v string) {
	o.Id = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetPart() bool {
	if o == nil || o.Part == nil {
		var ret bool
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetPartOk() (*bool, bool) {
	if o == nil || o.Part == nil {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasPart() bool {
	if o != nil && o.Part != nil {
		return true
	}

	return false
}

// SetPart gets a reference to the given bool and assigns it to the Part field.
func (o *BTInsertableSketchDisplayData3775) SetPart(v bool) {
	o.Part = &v
}

// GetSketchFeature returns the SketchFeature field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetSketchFeature() bool {
	if o == nil || o.SketchFeature == nil {
		var ret bool
		return ret
	}
	return *o.SketchFeature
}

// GetSketchFeatureOk returns a tuple with the SketchFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetSketchFeatureOk() (*bool, bool) {
	if o == nil || o.SketchFeature == nil {
		return nil, false
	}
	return o.SketchFeature, true
}

// HasSketchFeature returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasSketchFeature() bool {
	if o != nil && o.SketchFeature != nil {
		return true
	}

	return false
}

// SetSketchFeature gets a reference to the given bool and assigns it to the SketchFeature field.
func (o *BTInsertableSketchDisplayData3775) SetSketchFeature(v bool) {
	o.SketchFeature = &v
}

// GetTessellationSettingIndex returns the TessellationSettingIndex field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetTessellationSettingIndex() int32 {
	if o == nil || o.TessellationSettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.TessellationSettingIndex
}

// GetTessellationSettingIndexOk returns a tuple with the TessellationSettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetTessellationSettingIndexOk() (*int32, bool) {
	if o == nil || o.TessellationSettingIndex == nil {
		return nil, false
	}
	return o.TessellationSettingIndex, true
}

// HasTessellationSettingIndex returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasTessellationSettingIndex() bool {
	if o != nil && o.TessellationSettingIndex != nil {
		return true
	}

	return false
}

// SetTessellationSettingIndex gets a reference to the given int32 and assigns it to the TessellationSettingIndex field.
func (o *BTInsertableSketchDisplayData3775) SetTessellationSettingIndex(v int32) {
	o.TessellationSettingIndex = &v
}

// GetBodyIdToPartData returns the BodyIdToPartData field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetBodyIdToPartData() map[string]BTPartData16 {
	if o == nil || o.BodyIdToPartData == nil {
		var ret map[string]BTPartData16
		return ret
	}
	return *o.BodyIdToPartData
}

// GetBodyIdToPartDataOk returns a tuple with the BodyIdToPartData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetBodyIdToPartDataOk() (*map[string]BTPartData16, bool) {
	if o == nil || o.BodyIdToPartData == nil {
		return nil, false
	}
	return o.BodyIdToPartData, true
}

// HasBodyIdToPartData returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasBodyIdToPartData() bool {
	if o != nil && o.BodyIdToPartData != nil {
		return true
	}

	return false
}

// SetBodyIdToPartData gets a reference to the given map[string]BTPartData16 and assigns it to the BodyIdToPartData field.
func (o *BTInsertableSketchDisplayData3775) SetBodyIdToPartData(v map[string]BTPartData16) {
	o.BodyIdToPartData = &v
}

// GetSketchFeatureId returns the SketchFeatureId field value if set, zero value otherwise.
func (o *BTInsertableSketchDisplayData3775) GetSketchFeatureId() string {
	if o == nil || o.SketchFeatureId == nil {
		var ret string
		return ret
	}
	return *o.SketchFeatureId
}

// GetSketchFeatureIdOk returns a tuple with the SketchFeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableSketchDisplayData3775) GetSketchFeatureIdOk() (*string, bool) {
	if o == nil || o.SketchFeatureId == nil {
		return nil, false
	}
	return o.SketchFeatureId, true
}

// HasSketchFeatureId returns a boolean if a field has been set.
func (o *BTInsertableSketchDisplayData3775) HasSketchFeatureId() bool {
	if o != nil && o.SketchFeatureId != nil {
		return true
	}

	return false
}

// SetSketchFeatureId gets a reference to the given string and assigns it to the SketchFeatureId field.
func (o *BTInsertableSketchDisplayData3775) SetSketchFeatureId(v string) {
	o.SketchFeatureId = &v
}

func (o BTInsertableSketchDisplayData3775) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.GraphicsBuffers != nil {
		toSerialize["graphicsBuffers"] = o.GraphicsBuffers
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Part != nil {
		toSerialize["part"] = o.Part
	}
	if o.SketchFeature != nil {
		toSerialize["sketchFeature"] = o.SketchFeature
	}
	if o.TessellationSettingIndex != nil {
		toSerialize["tessellationSettingIndex"] = o.TessellationSettingIndex
	}
	if o.BodyIdToPartData != nil {
		toSerialize["bodyIdToPartData"] = o.BodyIdToPartData
	}
	if o.SketchFeatureId != nil {
		toSerialize["sketchFeatureId"] = o.SketchFeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableBTInsertableSketchDisplayData3775 struct {
	value *BTInsertableSketchDisplayData3775
	isSet bool
}

func (v NullableBTInsertableSketchDisplayData3775) Get() *BTInsertableSketchDisplayData3775 {
	return v.value
}

func (v *NullableBTInsertableSketchDisplayData3775) Set(val *BTInsertableSketchDisplayData3775) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInsertableSketchDisplayData3775) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInsertableSketchDisplayData3775) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInsertableSketchDisplayData3775(val *BTInsertableSketchDisplayData3775) *NullableBTInsertableSketchDisplayData3775 {
	return &NullableBTInsertableSketchDisplayData3775{value: val, isSet: true}
}

func (v NullableBTInsertableSketchDisplayData3775) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInsertableSketchDisplayData3775) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
