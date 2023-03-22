/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13258-6b30d30337fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTEntityGeometry35 - struct for BTEntityGeometry35
type BTEntityGeometry35 struct {
	implBTEntityGeometry35 interface{}
}

// BTDebugGeometry2059AsBTEntityGeometry35 is a convenience function that returns BTDebugGeometry2059 wrapped in BTEntityGeometry35
func (o *BTDebugGeometry2059) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// BTEntityPoint29AsBTEntityGeometry35 is a convenience function that returns BTEntityPoint29 wrapped in BTEntityGeometry35
func (o *BTEntityPoint29) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// BTEntityDegenerateEdge1129AsBTEntityGeometry35 is a convenience function that returns BTEntityDegenerateEdge1129 wrapped in BTEntityGeometry35
func (o *BTEntityDegenerateEdge1129) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// BTEntityEdge30AsBTEntityGeometry35 is a convenience function that returns BTEntityEdge30 wrapped in BTEntityGeometry35
func (o *BTEntityEdge30) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// BTSimulationFace2147AsBTEntityGeometry35 is a convenience function that returns BTSimulationFace2147 wrapped in BTEntityGeometry35
func (o *BTSimulationFace2147) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// BTTessellatedGeometry2576AsBTEntityGeometry35 is a convenience function that returns BTTessellatedGeometry2576 wrapped in BTEntityGeometry35
func (o *BTTessellatedGeometry2576) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// BTEntityFace31AsBTEntityGeometry35 is a convenience function that returns BTEntityFace31 wrapped in BTEntityGeometry35
func (o *BTEntityFace31) AsBTEntityGeometry35() *BTEntityGeometry35 {
	return &BTEntityGeometry35{o}
}

// NewBTEntityGeometry35 instantiates a new BTEntityGeometry35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEntityGeometry35() *BTEntityGeometry35 {
	this := BTEntityGeometry35{Newbase_BTEntityGeometry35()}
	return &this
}

// NewBTEntityGeometry35WithDefaults instantiates a new BTEntityGeometry35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEntityGeometry35WithDefaults() *BTEntityGeometry35 {
	this := BTEntityGeometry35{Newbase_BTEntityGeometry35WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityGeometry35) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetCompressed() bool {
	type getResult interface {
		GetCompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCompressed()
	} else {
		var de bool
		return de
	}
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetCompressedOk() (*bool, bool) {
	type getResult interface {
		GetCompressedOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCompressedOk()
	} else {
		return nil, false
	}
}

// HasCompressed returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasCompressed() bool {
	type getResult interface {
		HasCompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCompressed()
	} else {
		return false
	}
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *BTEntityGeometry35) SetCompressed(v bool) {
	type getResult interface {
		SetCompressed(v bool)
	}

	o.GetActualInstance().(getResult).SetCompressed(v)
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetDecompressed() BTEntityGeometry35 {
	type getResult interface {
		GetDecompressed() BTEntityGeometry35
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDecompressed()
	} else {
		var de BTEntityGeometry35
		return de
	}
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	type getResult interface {
		GetDecompressedOk() (*BTEntityGeometry35, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDecompressedOk()
	} else {
		return nil, false
	}
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasDecompressed() bool {
	type getResult interface {
		HasDecompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDecompressed()
	} else {
		return false
	}
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *BTEntityGeometry35) SetDecompressed(v BTEntityGeometry35) {
	type getResult interface {
		SetDecompressed(v BTEntityGeometry35)
	}

	o.GetActualInstance().(getResult).SetDecompressed(v)
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetErrorCode() int32 {
	type getResult interface {
		GetErrorCode() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorCode()
	} else {
		var de int32
		return de
	}
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetErrorCodeOk() (*int32, bool) {
	type getResult interface {
		GetErrorCodeOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorCodeOk()
	} else {
		return nil, false
	}
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasErrorCode() bool {
	type getResult interface {
		HasErrorCode() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasErrorCode()
	} else {
		return false
	}
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTEntityGeometry35) SetErrorCode(v int32) {
	type getResult interface {
		SetErrorCode(v int32)
	}

	o.GetActualInstance().(getResult).SetErrorCode(v)
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetEstimatedMemoryUsageInBytes() int32 {
	type getResult interface {
		GetEstimatedMemoryUsageInBytes() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEstimatedMemoryUsageInBytes()
	} else {
		var de int32
		return de
	}
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	type getResult interface {
		GetEstimatedMemoryUsageInBytesOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEstimatedMemoryUsageInBytesOk()
	} else {
		return nil, false
	}
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasEstimatedMemoryUsageInBytes() bool {
	type getResult interface {
		HasEstimatedMemoryUsageInBytes() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEstimatedMemoryUsageInBytes()
	} else {
		return false
	}
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *BTEntityGeometry35) SetEstimatedMemoryUsageInBytes(v int32) {
	type getResult interface {
		SetEstimatedMemoryUsageInBytes(v int32)
	}

	o.GetActualInstance().(getResult).SetEstimatedMemoryUsageInBytes(v)
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetFace() bool {
	type getResult interface {
		GetFace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFace()
	} else {
		var de bool
		return de
	}
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetFaceOk() (*bool, bool) {
	type getResult interface {
		GetFaceOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFaceOk()
	} else {
		return nil, false
	}
}

// HasFace returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasFace() bool {
	type getResult interface {
		HasFace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFace()
	} else {
		return false
	}
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *BTEntityGeometry35) SetFace(v bool) {
	type getResult interface {
		SetFace(v bool)
	}

	o.GetActualInstance().(getResult).SetFace(v)
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetHasTessellationError() bool {
	type getResult interface {
		GetHasTessellationError() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHasTessellationError()
	} else {
		var de bool
		return de
	}
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetHasTessellationErrorOk() (*bool, bool) {
	type getResult interface {
		GetHasTessellationErrorOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHasTessellationErrorOk()
	} else {
		return nil, false
	}
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasHasTessellationError() bool {
	type getResult interface {
		HasHasTessellationError() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasHasTessellationError()
	} else {
		return false
	}
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *BTEntityGeometry35) SetHasTessellationError(v bool) {
	type getResult interface {
		SetHasTessellationError(v bool)
	}

	o.GetActualInstance().(getResult).SetHasTessellationError(v)
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *BTEntityGeometry35) GetSettingIndex() int32 {
	type getResult interface {
		GetSettingIndex() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSettingIndex()
	} else {
		var de int32
		return de
	}
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityGeometry35) GetSettingIndexOk() (*int32, bool) {
	type getResult interface {
		GetSettingIndexOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSettingIndexOk()
	} else {
		return nil, false
	}
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *BTEntityGeometry35) HasSettingIndex() bool {
	type getResult interface {
		HasSettingIndex() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSettingIndex()
	} else {
		return false
	}
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *BTEntityGeometry35) SetSettingIndex(v int32) {
	type getResult interface {
		SetSettingIndex(v int32)
	}

	o.GetActualInstance().(getResult).SetSettingIndex(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTEntityGeometry35) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTDebugGeometry-2059'
	if jsonDict["btType"] == "BTDebugGeometry-2059" {
		// try to unmarshal JSON data into BTDebugGeometry2059
		var qr *BTDebugGeometry2059
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTDebugGeometry2059: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityDegenerateEdge-1129'
	if jsonDict["btType"] == "BTEntityDegenerateEdge-1129" {
		// try to unmarshal JSON data into BTEntityDegenerateEdge1129
		var qr *BTEntityDegenerateEdge1129
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTEntityDegenerateEdge1129: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityEdge-30'
	if jsonDict["btType"] == "BTEntityEdge-30" {
		// try to unmarshal JSON data into BTEntityEdge30
		var qr *BTEntityEdge30
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTEntityEdge30: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityFace-31'
	if jsonDict["btType"] == "BTEntityFace-31" {
		// try to unmarshal JSON data into BTEntityFace31
		var qr *BTEntityFace31
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTEntityFace31: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityPoint-29'
	if jsonDict["btType"] == "BTEntityPoint-29" {
		// try to unmarshal JSON data into BTEntityPoint29
		var qr *BTEntityPoint29
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTEntityPoint29: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSimulationFace-2147'
	if jsonDict["btType"] == "BTSimulationFace-2147" {
		// try to unmarshal JSON data into BTSimulationFace2147
		var qr *BTSimulationFace2147
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTSimulationFace2147: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTessellatedGeometry-2576'
	if jsonDict["btType"] == "BTTessellatedGeometry-2576" {
		// try to unmarshal JSON data into BTTessellatedGeometry2576
		var qr *BTTessellatedGeometry2576
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEntityGeometry35 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEntityGeometry35 = nil
			return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as BTTessellatedGeometry2576: %s", err.Error())
		}
	}

	var qtx *base_BTEntityGeometry35
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTEntityGeometry35 = qtx
		return nil // data stored in dst.base_BTEntityGeometry35, return on the first match
	} else {
		dst.implBTEntityGeometry35 = nil
		return fmt.Errorf("Failed to unmarshal BTEntityGeometry35 as base_BTEntityGeometry35: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTEntityGeometry35) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTEntityGeometry35) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTEntityGeometry35
}

type NullableBTEntityGeometry35 struct {
	value *BTEntityGeometry35
	isSet bool
}

func (v NullableBTEntityGeometry35) Get() *BTEntityGeometry35 {
	return v.value
}

func (v *NullableBTEntityGeometry35) Set(val *BTEntityGeometry35) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEntityGeometry35) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEntityGeometry35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEntityGeometry35(val *BTEntityGeometry35) *NullableBTEntityGeometry35 {
	return &NullableBTEntityGeometry35{value: val, isSet: true}
}

func (v NullableBTEntityGeometry35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEntityGeometry35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTEntityGeometry35 struct {
	BtType                      *string             `json:"btType,omitempty"`
	Compressed                  *bool               `json:"compressed,omitempty"`
	Decompressed                *BTEntityGeometry35 `json:"decompressed,omitempty"`
	ErrorCode                   *int32              `json:"errorCode,omitempty"`
	EstimatedMemoryUsageInBytes *int32              `json:"estimatedMemoryUsageInBytes,omitempty"`
	Face                        *bool               `json:"face,omitempty"`
	HasTessellationError        *bool               `json:"hasTessellationError,omitempty"`
	SettingIndex                *int32              `json:"settingIndex,omitempty"`
}

// Newbase_BTEntityGeometry35 instantiates a new base_BTEntityGeometry35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTEntityGeometry35() *base_BTEntityGeometry35 {
	this := base_BTEntityGeometry35{}
	return &this
}

// Newbase_BTEntityGeometry35WithDefaults instantiates a new base_BTEntityGeometry35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTEntityGeometry35WithDefaults() *base_BTEntityGeometry35 {
	this := base_BTEntityGeometry35{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTEntityGeometry35) SetBtType(v string) {
	o.BtType = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *base_BTEntityGeometry35) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetDecompressed() BTEntityGeometry35 {
	if o == nil || o.Decompressed == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *base_BTEntityGeometry35) SetDecompressed(v BTEntityGeometry35) {
	o.Decompressed = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *base_BTEntityGeometry35) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetEstimatedMemoryUsageInBytes() int32 {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMemoryUsageInBytes
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		return nil, false
	}
	return o.EstimatedMemoryUsageInBytes, true
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasEstimatedMemoryUsageInBytes() bool {
	if o != nil && o.EstimatedMemoryUsageInBytes != nil {
		return true
	}

	return false
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *base_BTEntityGeometry35) SetEstimatedMemoryUsageInBytes(v int32) {
	o.EstimatedMemoryUsageInBytes = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetFace() bool {
	if o == nil || o.Face == nil {
		var ret bool
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetFaceOk() (*bool, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *base_BTEntityGeometry35) SetFace(v bool) {
	o.Face = &v
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetHasTessellationError() bool {
	if o == nil || o.HasTessellationError == nil {
		var ret bool
		return ret
	}
	return *o.HasTessellationError
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetHasTessellationErrorOk() (*bool, bool) {
	if o == nil || o.HasTessellationError == nil {
		return nil, false
	}
	return o.HasTessellationError, true
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasHasTessellationError() bool {
	if o != nil && o.HasTessellationError != nil {
		return true
	}

	return false
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *base_BTEntityGeometry35) SetHasTessellationError(v bool) {
	o.HasTessellationError = &v
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *base_BTEntityGeometry35) GetSettingIndex() int32 {
	if o == nil || o.SettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.SettingIndex
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEntityGeometry35) GetSettingIndexOk() (*int32, bool) {
	if o == nil || o.SettingIndex == nil {
		return nil, false
	}
	return o.SettingIndex, true
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *base_BTEntityGeometry35) HasSettingIndex() bool {
	if o != nil && o.SettingIndex != nil {
		return true
	}

	return false
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *base_BTEntityGeometry35) SetSettingIndex(v int32) {
	o.SettingIndex = &v
}

func (o base_BTEntityGeometry35) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}
