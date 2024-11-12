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

// BTDatumDisplayData3408 struct for BTDatumDisplayData3408
type BTDatumDisplayData3408 struct {
	BTAnnotationDisplayData3225
	AnnotationPlane         []float32 `json:"annotationPlane,omitempty"`
	BaseNormal              []float32 `json:"baseNormal,omitempty"`
	BasePoint               []float32 `json:"basePoint,omitempty"`
	BtType                  *string   `json:"btType,omitempty"`
	DeterministicId         *string   `json:"deterministicId,omitempty"`
	DxdySegments            []float32 `json:"dxdySegments,omitempty"`
	NumberOfLeaderSegements *int32    `json:"numberOfLeaderSegements,omitempty"`
	Name                    *string   `json:"name,omitempty"`
}

// NewBTDatumDisplayData3408 instantiates a new BTDatumDisplayData3408 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDatumDisplayData3408() *BTDatumDisplayData3408 {
	this := BTDatumDisplayData3408{}
	return &this
}

// NewBTDatumDisplayData3408WithDefaults instantiates a new BTDatumDisplayData3408 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDatumDisplayData3408WithDefaults() *BTDatumDisplayData3408 {
	this := BTDatumDisplayData3408{}
	return &this
}

// GetAnnotationPlane returns the AnnotationPlane field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetAnnotationPlane() []float32 {
	if o == nil || o.AnnotationPlane == nil {
		var ret []float32
		return ret
	}
	return o.AnnotationPlane
}

// GetAnnotationPlaneOk returns a tuple with the AnnotationPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetAnnotationPlaneOk() ([]float32, bool) {
	if o == nil || o.AnnotationPlane == nil {
		return nil, false
	}
	return o.AnnotationPlane, true
}

// HasAnnotationPlane returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasAnnotationPlane() bool {
	if o != nil && o.AnnotationPlane != nil {
		return true
	}

	return false
}

// SetAnnotationPlane gets a reference to the given []float32 and assigns it to the AnnotationPlane field.
func (o *BTDatumDisplayData3408) SetAnnotationPlane(v []float32) {
	o.AnnotationPlane = v
}

// GetBaseNormal returns the BaseNormal field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetBaseNormal() []float32 {
	if o == nil || o.BaseNormal == nil {
		var ret []float32
		return ret
	}
	return o.BaseNormal
}

// GetBaseNormalOk returns a tuple with the BaseNormal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetBaseNormalOk() ([]float32, bool) {
	if o == nil || o.BaseNormal == nil {
		return nil, false
	}
	return o.BaseNormal, true
}

// HasBaseNormal returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasBaseNormal() bool {
	if o != nil && o.BaseNormal != nil {
		return true
	}

	return false
}

// SetBaseNormal gets a reference to the given []float32 and assigns it to the BaseNormal field.
func (o *BTDatumDisplayData3408) SetBaseNormal(v []float32) {
	o.BaseNormal = v
}

// GetBasePoint returns the BasePoint field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetBasePoint() []float32 {
	if o == nil || o.BasePoint == nil {
		var ret []float32
		return ret
	}
	return o.BasePoint
}

// GetBasePointOk returns a tuple with the BasePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetBasePointOk() ([]float32, bool) {
	if o == nil || o.BasePoint == nil {
		return nil, false
	}
	return o.BasePoint, true
}

// HasBasePoint returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasBasePoint() bool {
	if o != nil && o.BasePoint != nil {
		return true
	}

	return false
}

// SetBasePoint gets a reference to the given []float32 and assigns it to the BasePoint field.
func (o *BTDatumDisplayData3408) SetBasePoint(v []float32) {
	o.BasePoint = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDatumDisplayData3408) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicId returns the DeterministicId field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetDeterministicId() string {
	if o == nil || o.DeterministicId == nil {
		var ret string
		return ret
	}
	return *o.DeterministicId
}

// GetDeterministicIdOk returns a tuple with the DeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetDeterministicIdOk() (*string, bool) {
	if o == nil || o.DeterministicId == nil {
		return nil, false
	}
	return o.DeterministicId, true
}

// HasDeterministicId returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasDeterministicId() bool {
	if o != nil && o.DeterministicId != nil {
		return true
	}

	return false
}

// SetDeterministicId gets a reference to the given string and assigns it to the DeterministicId field.
func (o *BTDatumDisplayData3408) SetDeterministicId(v string) {
	o.DeterministicId = &v
}

// GetDxdySegments returns the DxdySegments field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetDxdySegments() []float32 {
	if o == nil || o.DxdySegments == nil {
		var ret []float32
		return ret
	}
	return o.DxdySegments
}

// GetDxdySegmentsOk returns a tuple with the DxdySegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetDxdySegmentsOk() ([]float32, bool) {
	if o == nil || o.DxdySegments == nil {
		return nil, false
	}
	return o.DxdySegments, true
}

// HasDxdySegments returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasDxdySegments() bool {
	if o != nil && o.DxdySegments != nil {
		return true
	}

	return false
}

// SetDxdySegments gets a reference to the given []float32 and assigns it to the DxdySegments field.
func (o *BTDatumDisplayData3408) SetDxdySegments(v []float32) {
	o.DxdySegments = v
}

// GetNumberOfLeaderSegements returns the NumberOfLeaderSegements field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetNumberOfLeaderSegements() int32 {
	if o == nil || o.NumberOfLeaderSegements == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfLeaderSegements
}

// GetNumberOfLeaderSegementsOk returns a tuple with the NumberOfLeaderSegements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetNumberOfLeaderSegementsOk() (*int32, bool) {
	if o == nil || o.NumberOfLeaderSegements == nil {
		return nil, false
	}
	return o.NumberOfLeaderSegements, true
}

// HasNumberOfLeaderSegements returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasNumberOfLeaderSegements() bool {
	if o != nil && o.NumberOfLeaderSegements != nil {
		return true
	}

	return false
}

// SetNumberOfLeaderSegements gets a reference to the given int32 and assigns it to the NumberOfLeaderSegements field.
func (o *BTDatumDisplayData3408) SetNumberOfLeaderSegements(v int32) {
	o.NumberOfLeaderSegements = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDatumDisplayData3408) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDatumDisplayData3408) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDatumDisplayData3408) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDatumDisplayData3408) SetName(v string) {
	o.Name = &v
}

func (o BTDatumDisplayData3408) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTAnnotationDisplayData3225, errBTAnnotationDisplayData3225 := json.Marshal(o.BTAnnotationDisplayData3225)
	if errBTAnnotationDisplayData3225 != nil {
		return []byte{}, errBTAnnotationDisplayData3225
	}
	errBTAnnotationDisplayData3225 = json.Unmarshal([]byte(serializedBTAnnotationDisplayData3225), &toSerialize)
	if errBTAnnotationDisplayData3225 != nil {
		return []byte{}, errBTAnnotationDisplayData3225
	}
	if o.AnnotationPlane != nil {
		toSerialize["annotationPlane"] = o.AnnotationPlane
	}
	if o.BaseNormal != nil {
		toSerialize["baseNormal"] = o.BaseNormal
	}
	if o.BasePoint != nil {
		toSerialize["basePoint"] = o.BasePoint
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicId != nil {
		toSerialize["deterministicId"] = o.DeterministicId
	}
	if o.DxdySegments != nil {
		toSerialize["dxdySegments"] = o.DxdySegments
	}
	if o.NumberOfLeaderSegements != nil {
		toSerialize["numberOfLeaderSegements"] = o.NumberOfLeaderSegements
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTDatumDisplayData3408 struct {
	value *BTDatumDisplayData3408
	isSet bool
}

func (v NullableBTDatumDisplayData3408) Get() *BTDatumDisplayData3408 {
	return v.value
}

func (v *NullableBTDatumDisplayData3408) Set(val *BTDatumDisplayData3408) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDatumDisplayData3408) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDatumDisplayData3408) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDatumDisplayData3408(val *BTDatumDisplayData3408) *NullableBTDatumDisplayData3408 {
	return &NullableBTDatumDisplayData3408{value: val, isSet: true}
}

func (v NullableBTDatumDisplayData3408) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDatumDisplayData3408) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
