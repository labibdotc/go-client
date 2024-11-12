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

// BTInsertablePartDisplayData3103 struct for BTInsertablePartDisplayData3103
type BTInsertablePartDisplayData3103 struct {
	BTInsertableDisplayData2405
	BtType                   *string                                     `json:"btType,omitempty"`
	FullElementId            *BTFullElementId756                         `json:"fullElementId,omitempty"`
	GraphicsBuffers          *map[string]map[string]BTGraphicsBuffer2668 `json:"graphicsBuffers,omitempty"`
	Id                       *string                                     `json:"id,omitempty"`
	Part                     *bool                                       `json:"part,omitempty"`
	SketchFeature            *bool                                       `json:"sketchFeature,omitempty"`
	TessellationSettingIndex *int32                                      `json:"tessellationSettingIndex,omitempty"`
	PartData                 *BTPartData16                               `json:"partData,omitempty"`
	PartDisplayData          *BTPartDisplayData17                        `json:"partDisplayData,omitempty"`
	PartId                   *string                                     `json:"partId,omitempty"`
	TessellationSetting      *int32                                      `json:"tessellationSetting,omitempty"`
}

// NewBTInsertablePartDisplayData3103 instantiates a new BTInsertablePartDisplayData3103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInsertablePartDisplayData3103() *BTInsertablePartDisplayData3103 {
	this := BTInsertablePartDisplayData3103{}
	return &this
}

// NewBTInsertablePartDisplayData3103WithDefaults instantiates a new BTInsertablePartDisplayData3103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInsertablePartDisplayData3103WithDefaults() *BTInsertablePartDisplayData3103 {
	this := BTInsertablePartDisplayData3103{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInsertablePartDisplayData3103) SetBtType(v string) {
	o.BtType = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTInsertablePartDisplayData3103) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetGraphicsBuffers returns the GraphicsBuffers field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetGraphicsBuffers() map[string]map[string]BTGraphicsBuffer2668 {
	if o == nil || o.GraphicsBuffers == nil {
		var ret map[string]map[string]BTGraphicsBuffer2668
		return ret
	}
	return *o.GraphicsBuffers
}

// GetGraphicsBuffersOk returns a tuple with the GraphicsBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetGraphicsBuffersOk() (*map[string]map[string]BTGraphicsBuffer2668, bool) {
	if o == nil || o.GraphicsBuffers == nil {
		return nil, false
	}
	return o.GraphicsBuffers, true
}

// HasGraphicsBuffers returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasGraphicsBuffers() bool {
	if o != nil && o.GraphicsBuffers != nil {
		return true
	}

	return false
}

// SetGraphicsBuffers gets a reference to the given map[string]map[string]BTGraphicsBuffer2668 and assigns it to the GraphicsBuffers field.
func (o *BTInsertablePartDisplayData3103) SetGraphicsBuffers(v map[string]map[string]BTGraphicsBuffer2668) {
	o.GraphicsBuffers = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTInsertablePartDisplayData3103) SetId(v string) {
	o.Id = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetPart() bool {
	if o == nil || o.Part == nil {
		var ret bool
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetPartOk() (*bool, bool) {
	if o == nil || o.Part == nil {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasPart() bool {
	if o != nil && o.Part != nil {
		return true
	}

	return false
}

// SetPart gets a reference to the given bool and assigns it to the Part field.
func (o *BTInsertablePartDisplayData3103) SetPart(v bool) {
	o.Part = &v
}

// GetSketchFeature returns the SketchFeature field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetSketchFeature() bool {
	if o == nil || o.SketchFeature == nil {
		var ret bool
		return ret
	}
	return *o.SketchFeature
}

// GetSketchFeatureOk returns a tuple with the SketchFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetSketchFeatureOk() (*bool, bool) {
	if o == nil || o.SketchFeature == nil {
		return nil, false
	}
	return o.SketchFeature, true
}

// HasSketchFeature returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasSketchFeature() bool {
	if o != nil && o.SketchFeature != nil {
		return true
	}

	return false
}

// SetSketchFeature gets a reference to the given bool and assigns it to the SketchFeature field.
func (o *BTInsertablePartDisplayData3103) SetSketchFeature(v bool) {
	o.SketchFeature = &v
}

// GetTessellationSettingIndex returns the TessellationSettingIndex field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetTessellationSettingIndex() int32 {
	if o == nil || o.TessellationSettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.TessellationSettingIndex
}

// GetTessellationSettingIndexOk returns a tuple with the TessellationSettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetTessellationSettingIndexOk() (*int32, bool) {
	if o == nil || o.TessellationSettingIndex == nil {
		return nil, false
	}
	return o.TessellationSettingIndex, true
}

// HasTessellationSettingIndex returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasTessellationSettingIndex() bool {
	if o != nil && o.TessellationSettingIndex != nil {
		return true
	}

	return false
}

// SetTessellationSettingIndex gets a reference to the given int32 and assigns it to the TessellationSettingIndex field.
func (o *BTInsertablePartDisplayData3103) SetTessellationSettingIndex(v int32) {
	o.TessellationSettingIndex = &v
}

// GetPartData returns the PartData field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetPartData() BTPartData16 {
	if o == nil || o.PartData == nil {
		var ret BTPartData16
		return ret
	}
	return *o.PartData
}

// GetPartDataOk returns a tuple with the PartData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetPartDataOk() (*BTPartData16, bool) {
	if o == nil || o.PartData == nil {
		return nil, false
	}
	return o.PartData, true
}

// HasPartData returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasPartData() bool {
	if o != nil && o.PartData != nil {
		return true
	}

	return false
}

// SetPartData gets a reference to the given BTPartData16 and assigns it to the PartData field.
func (o *BTInsertablePartDisplayData3103) SetPartData(v BTPartData16) {
	o.PartData = &v
}

// GetPartDisplayData returns the PartDisplayData field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetPartDisplayData() BTPartDisplayData17 {
	if o == nil || o.PartDisplayData == nil {
		var ret BTPartDisplayData17
		return ret
	}
	return *o.PartDisplayData
}

// GetPartDisplayDataOk returns a tuple with the PartDisplayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetPartDisplayDataOk() (*BTPartDisplayData17, bool) {
	if o == nil || o.PartDisplayData == nil {
		return nil, false
	}
	return o.PartDisplayData, true
}

// HasPartDisplayData returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasPartDisplayData() bool {
	if o != nil && o.PartDisplayData != nil {
		return true
	}

	return false
}

// SetPartDisplayData gets a reference to the given BTPartDisplayData17 and assigns it to the PartDisplayData field.
func (o *BTInsertablePartDisplayData3103) SetPartDisplayData(v BTPartDisplayData17) {
	o.PartDisplayData = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTInsertablePartDisplayData3103) SetPartId(v string) {
	o.PartId = &v
}

// GetTessellationSetting returns the TessellationSetting field value if set, zero value otherwise.
func (o *BTInsertablePartDisplayData3103) GetTessellationSetting() int32 {
	if o == nil || o.TessellationSetting == nil {
		var ret int32
		return ret
	}
	return *o.TessellationSetting
}

// GetTessellationSettingOk returns a tuple with the TessellationSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablePartDisplayData3103) GetTessellationSettingOk() (*int32, bool) {
	if o == nil || o.TessellationSetting == nil {
		return nil, false
	}
	return o.TessellationSetting, true
}

// HasTessellationSetting returns a boolean if a field has been set.
func (o *BTInsertablePartDisplayData3103) HasTessellationSetting() bool {
	if o != nil && o.TessellationSetting != nil {
		return true
	}

	return false
}

// SetTessellationSetting gets a reference to the given int32 and assigns it to the TessellationSetting field.
func (o *BTInsertablePartDisplayData3103) SetTessellationSetting(v int32) {
	o.TessellationSetting = &v
}

func (o BTInsertablePartDisplayData3103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTInsertableDisplayData2405, errBTInsertableDisplayData2405 := json.Marshal(o.BTInsertableDisplayData2405)
	if errBTInsertableDisplayData2405 != nil {
		return []byte{}, errBTInsertableDisplayData2405
	}
	errBTInsertableDisplayData2405 = json.Unmarshal([]byte(serializedBTInsertableDisplayData2405), &toSerialize)
	if errBTInsertableDisplayData2405 != nil {
		return []byte{}, errBTInsertableDisplayData2405
	}
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
	if o.PartData != nil {
		toSerialize["partData"] = o.PartData
	}
	if o.PartDisplayData != nil {
		toSerialize["partDisplayData"] = o.PartDisplayData
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.TessellationSetting != nil {
		toSerialize["tessellationSetting"] = o.TessellationSetting
	}
	return json.Marshal(toSerialize)
}

type NullableBTInsertablePartDisplayData3103 struct {
	value *BTInsertablePartDisplayData3103
	isSet bool
}

func (v NullableBTInsertablePartDisplayData3103) Get() *BTInsertablePartDisplayData3103 {
	return v.value
}

func (v *NullableBTInsertablePartDisplayData3103) Set(val *BTInsertablePartDisplayData3103) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInsertablePartDisplayData3103) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInsertablePartDisplayData3103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInsertablePartDisplayData3103(val *BTInsertablePartDisplayData3103) *NullableBTInsertablePartDisplayData3103 {
	return &NullableBTInsertablePartDisplayData3103{value: val, isSet: true}
}

func (v NullableBTInsertablePartDisplayData3103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInsertablePartDisplayData3103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
