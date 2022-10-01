/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6730-405400b0583f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDebugGeometry2059 struct for BTDebugGeometry2059
type BTDebugGeometry2059 struct {
	BtType                           *string                    `json:"btType,omitempty"`
	Compressed                       *bool                      `json:"compressed,omitempty"`
	Decompressed                     *BTEntityGeometry35        `json:"decompressed,omitempty"`
	ErrorCode                        *int32                     `json:"errorCode,omitempty"`
	EstimatedMemoryUsageInBytes      *int32                     `json:"estimatedMemoryUsageInBytes,omitempty"`
	Face                             *bool                      `json:"face,omitempty"`
	HasTessellationError             *bool                      `json:"hasTessellationError,omitempty"`
	SettingIndex                     *int32                     `json:"settingIndex,omitempty"`
	Appearance                       *BTGraphicsAppearance1152  `json:"appearance,omitempty"`
	BelongsToFlattenedSheetMetalBody *bool                      `json:"belongsToFlattenedSheetMetalBody,omitempty"`
	BodyId                           *string                    `json:"bodyId,omitempty"`
	Color                            *string                    `json:"color,omitempty"`
	DeterministicId                  *string                    `json:"deterministicId,omitempty"`
	SheetMetalModelId                *string                    `json:"sheetMetalModelId,omitempty"`
	Style                            *string                    `json:"style,omitempty"`
	Tessellation                     *BTTessellatedGeometry2576 `json:"tessellation,omitempty"`
}

// NewBTDebugGeometry2059 instantiates a new BTDebugGeometry2059 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDebugGeometry2059() *BTDebugGeometry2059 {
	this := BTDebugGeometry2059{}
	return &this
}

// NewBTDebugGeometry2059WithDefaults instantiates a new BTDebugGeometry2059 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDebugGeometry2059WithDefaults() *BTDebugGeometry2059 {
	this := BTDebugGeometry2059{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDebugGeometry2059) SetBtType(v string) {
	o.BtType = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *BTDebugGeometry2059) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetDecompressed() BTEntityGeometry35 {
	if o == nil || o.Decompressed == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *BTDebugGeometry2059) SetDecompressed(v BTEntityGeometry35) {
	o.Decompressed = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTDebugGeometry2059) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetEstimatedMemoryUsageInBytes() int32 {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMemoryUsageInBytes
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		return nil, false
	}
	return o.EstimatedMemoryUsageInBytes, true
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasEstimatedMemoryUsageInBytes() bool {
	if o != nil && o.EstimatedMemoryUsageInBytes != nil {
		return true
	}

	return false
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *BTDebugGeometry2059) SetEstimatedMemoryUsageInBytes(v int32) {
	o.EstimatedMemoryUsageInBytes = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetFace() bool {
	if o == nil || o.Face == nil {
		var ret bool
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetFaceOk() (*bool, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *BTDebugGeometry2059) SetFace(v bool) {
	o.Face = &v
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetHasTessellationError() bool {
	if o == nil || o.HasTessellationError == nil {
		var ret bool
		return ret
	}
	return *o.HasTessellationError
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetHasTessellationErrorOk() (*bool, bool) {
	if o == nil || o.HasTessellationError == nil {
		return nil, false
	}
	return o.HasTessellationError, true
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasHasTessellationError() bool {
	if o != nil && o.HasTessellationError != nil {
		return true
	}

	return false
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *BTDebugGeometry2059) SetHasTessellationError(v bool) {
	o.HasTessellationError = &v
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetSettingIndex() int32 {
	if o == nil || o.SettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.SettingIndex
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetSettingIndexOk() (*int32, bool) {
	if o == nil || o.SettingIndex == nil {
		return nil, false
	}
	return o.SettingIndex, true
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasSettingIndex() bool {
	if o != nil && o.SettingIndex != nil {
		return true
	}

	return false
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *BTDebugGeometry2059) SetSettingIndex(v int32) {
	o.SettingIndex = &v
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTDebugGeometry2059) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBelongsToFlattenedSheetMetalBody returns the BelongsToFlattenedSheetMetalBody field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetBelongsToFlattenedSheetMetalBody() bool {
	if o == nil || o.BelongsToFlattenedSheetMetalBody == nil {
		var ret bool
		return ret
	}
	return *o.BelongsToFlattenedSheetMetalBody
}

// GetBelongsToFlattenedSheetMetalBodyOk returns a tuple with the BelongsToFlattenedSheetMetalBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetBelongsToFlattenedSheetMetalBodyOk() (*bool, bool) {
	if o == nil || o.BelongsToFlattenedSheetMetalBody == nil {
		return nil, false
	}
	return o.BelongsToFlattenedSheetMetalBody, true
}

// HasBelongsToFlattenedSheetMetalBody returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasBelongsToFlattenedSheetMetalBody() bool {
	if o != nil && o.BelongsToFlattenedSheetMetalBody != nil {
		return true
	}

	return false
}

// SetBelongsToFlattenedSheetMetalBody gets a reference to the given bool and assigns it to the BelongsToFlattenedSheetMetalBody field.
func (o *BTDebugGeometry2059) SetBelongsToFlattenedSheetMetalBody(v bool) {
	o.BelongsToFlattenedSheetMetalBody = &v
}

// GetBodyId returns the BodyId field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetBodyId() string {
	if o == nil || o.BodyId == nil {
		var ret string
		return ret
	}
	return *o.BodyId
}

// GetBodyIdOk returns a tuple with the BodyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetBodyIdOk() (*string, bool) {
	if o == nil || o.BodyId == nil {
		return nil, false
	}
	return o.BodyId, true
}

// HasBodyId returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasBodyId() bool {
	if o != nil && o.BodyId != nil {
		return true
	}

	return false
}

// SetBodyId gets a reference to the given string and assigns it to the BodyId field.
func (o *BTDebugGeometry2059) SetBodyId(v string) {
	o.BodyId = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *BTDebugGeometry2059) SetColor(v string) {
	o.Color = &v
}

// GetDeterministicId returns the DeterministicId field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetDeterministicId() string {
	if o == nil || o.DeterministicId == nil {
		var ret string
		return ret
	}
	return *o.DeterministicId
}

// GetDeterministicIdOk returns a tuple with the DeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetDeterministicIdOk() (*string, bool) {
	if o == nil || o.DeterministicId == nil {
		return nil, false
	}
	return o.DeterministicId, true
}

// HasDeterministicId returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasDeterministicId() bool {
	if o != nil && o.DeterministicId != nil {
		return true
	}

	return false
}

// SetDeterministicId gets a reference to the given string and assigns it to the DeterministicId field.
func (o *BTDebugGeometry2059) SetDeterministicId(v string) {
	o.DeterministicId = &v
}

// GetSheetMetalModelId returns the SheetMetalModelId field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetSheetMetalModelId() string {
	if o == nil || o.SheetMetalModelId == nil {
		var ret string
		return ret
	}
	return *o.SheetMetalModelId
}

// GetSheetMetalModelIdOk returns a tuple with the SheetMetalModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetSheetMetalModelIdOk() (*string, bool) {
	if o == nil || o.SheetMetalModelId == nil {
		return nil, false
	}
	return o.SheetMetalModelId, true
}

// HasSheetMetalModelId returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasSheetMetalModelId() bool {
	if o != nil && o.SheetMetalModelId != nil {
		return true
	}

	return false
}

// SetSheetMetalModelId gets a reference to the given string and assigns it to the SheetMetalModelId field.
func (o *BTDebugGeometry2059) SetSheetMetalModelId(v string) {
	o.SheetMetalModelId = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetStyle() string {
	if o == nil || o.Style == nil {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetStyleOk() (*string, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given string and assigns it to the Style field.
func (o *BTDebugGeometry2059) SetStyle(v string) {
	o.Style = &v
}

// GetTessellation returns the Tessellation field value if set, zero value otherwise.
func (o *BTDebugGeometry2059) GetTessellation() BTTessellatedGeometry2576 {
	if o == nil || o.Tessellation == nil {
		var ret BTTessellatedGeometry2576
		return ret
	}
	return *o.Tessellation
}

// GetTessellationOk returns a tuple with the Tessellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDebugGeometry2059) GetTessellationOk() (*BTTessellatedGeometry2576, bool) {
	if o == nil || o.Tessellation == nil {
		return nil, false
	}
	return o.Tessellation, true
}

// HasTessellation returns a boolean if a field has been set.
func (o *BTDebugGeometry2059) HasTessellation() bool {
	if o != nil && o.Tessellation != nil {
		return true
	}

	return false
}

// SetTessellation gets a reference to the given BTTessellatedGeometry2576 and assigns it to the Tessellation field.
func (o *BTDebugGeometry2059) SetTessellation(v BTTessellatedGeometry2576) {
	o.Tessellation = &v
}

func (o BTDebugGeometry2059) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Compressed != nil {
		toSerialize["compressed"] = o.Compressed
	}
	if o.Decompressed != nil {
		toSerialize["decompressed"] = o.Decompressed
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.EstimatedMemoryUsageInBytes != nil {
		toSerialize["estimatedMemoryUsageInBytes"] = o.EstimatedMemoryUsageInBytes
	}
	if o.Face != nil {
		toSerialize["face"] = o.Face
	}
	if o.HasTessellationError != nil {
		toSerialize["hasTessellationError"] = o.HasTessellationError
	}
	if o.SettingIndex != nil {
		toSerialize["settingIndex"] = o.SettingIndex
	}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BelongsToFlattenedSheetMetalBody != nil {
		toSerialize["belongsToFlattenedSheetMetalBody"] = o.BelongsToFlattenedSheetMetalBody
	}
	if o.BodyId != nil {
		toSerialize["bodyId"] = o.BodyId
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.DeterministicId != nil {
		toSerialize["deterministicId"] = o.DeterministicId
	}
	if o.SheetMetalModelId != nil {
		toSerialize["sheetMetalModelId"] = o.SheetMetalModelId
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.Tessellation != nil {
		toSerialize["tessellation"] = o.Tessellation
	}
	return json.Marshal(toSerialize)
}

type NullableBTDebugGeometry2059 struct {
	value *BTDebugGeometry2059
	isSet bool
}

func (v NullableBTDebugGeometry2059) Get() *BTDebugGeometry2059 {
	return v.value
}

func (v *NullableBTDebugGeometry2059) Set(val *BTDebugGeometry2059) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDebugGeometry2059) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDebugGeometry2059) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDebugGeometry2059(val *BTDebugGeometry2059) *NullableBTDebugGeometry2059 {
	return &NullableBTDebugGeometry2059{value: val, isSet: true}
}

func (v NullableBTDebugGeometry2059) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDebugGeometry2059) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
