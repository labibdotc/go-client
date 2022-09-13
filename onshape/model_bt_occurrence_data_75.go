/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6415-48a6b2252b8c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOccurrenceData75 struct for BTOccurrenceData75
type BTOccurrenceData75 struct {
	BtType                          *string                                `json:"btType,omitempty"`
	FeatureData                     *map[string]BTFeatureOccurrenceData775 `json:"featureData,omitempty"`
	ForceHighestQualityTessellation *bool                                  `json:"forceHighestQualityTessellation,omitempty"`
	Hidden                          *bool                                  `json:"hidden,omitempty"`
	IsFixed                         *bool                                  `json:"isFixed,omitempty"`
	IsHidden                        *bool                                  `json:"isHidden,omitempty"`
	NodeId                          *string                                `json:"nodeId,omitempty"`
	Occurrence                      *BTOccurrence74                        `json:"occurrence,omitempty"`
	Transform                       *BTBSMatrix386                         `json:"transform,omitempty"`
}

// NewBTOccurrenceData75 instantiates a new BTOccurrenceData75 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOccurrenceData75() *BTOccurrenceData75 {
	this := BTOccurrenceData75{}
	return &this
}

// NewBTOccurrenceData75WithDefaults instantiates a new BTOccurrenceData75 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOccurrenceData75WithDefaults() *BTOccurrenceData75 {
	this := BTOccurrenceData75{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOccurrenceData75) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureData returns the FeatureData field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetFeatureData() map[string]BTFeatureOccurrenceData775 {
	if o == nil || o.FeatureData == nil {
		var ret map[string]BTFeatureOccurrenceData775
		return ret
	}
	return *o.FeatureData
}

// GetFeatureDataOk returns a tuple with the FeatureData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetFeatureDataOk() (*map[string]BTFeatureOccurrenceData775, bool) {
	if o == nil || o.FeatureData == nil {
		return nil, false
	}
	return o.FeatureData, true
}

// HasFeatureData returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasFeatureData() bool {
	if o != nil && o.FeatureData != nil {
		return true
	}

	return false
}

// SetFeatureData gets a reference to the given map[string]BTFeatureOccurrenceData775 and assigns it to the FeatureData field.
func (o *BTOccurrenceData75) SetFeatureData(v map[string]BTFeatureOccurrenceData775) {
	o.FeatureData = &v
}

// GetForceHighestQualityTessellation returns the ForceHighestQualityTessellation field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetForceHighestQualityTessellation() bool {
	if o == nil || o.ForceHighestQualityTessellation == nil {
		var ret bool
		return ret
	}
	return *o.ForceHighestQualityTessellation
}

// GetForceHighestQualityTessellationOk returns a tuple with the ForceHighestQualityTessellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetForceHighestQualityTessellationOk() (*bool, bool) {
	if o == nil || o.ForceHighestQualityTessellation == nil {
		return nil, false
	}
	return o.ForceHighestQualityTessellation, true
}

// HasForceHighestQualityTessellation returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasForceHighestQualityTessellation() bool {
	if o != nil && o.ForceHighestQualityTessellation != nil {
		return true
	}

	return false
}

// SetForceHighestQualityTessellation gets a reference to the given bool and assigns it to the ForceHighestQualityTessellation field.
func (o *BTOccurrenceData75) SetForceHighestQualityTessellation(v bool) {
	o.ForceHighestQualityTessellation = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTOccurrenceData75) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIsFixed returns the IsFixed field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetIsFixed() bool {
	if o == nil || o.IsFixed == nil {
		var ret bool
		return ret
	}
	return *o.IsFixed
}

// GetIsFixedOk returns a tuple with the IsFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetIsFixedOk() (*bool, bool) {
	if o == nil || o.IsFixed == nil {
		return nil, false
	}
	return o.IsFixed, true
}

// HasIsFixed returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasIsFixed() bool {
	if o != nil && o.IsFixed != nil {
		return true
	}

	return false
}

// SetIsFixed gets a reference to the given bool and assigns it to the IsFixed field.
func (o *BTOccurrenceData75) SetIsFixed(v bool) {
	o.IsFixed = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasIsHidden() bool {
	if o != nil && o.IsHidden != nil {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *BTOccurrenceData75) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTOccurrenceData75) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTOccurrenceData75) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *BTOccurrenceData75) GetTransform() BTBSMatrix386 {
	if o == nil || o.Transform == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceData75) GetTransformOk() (*BTBSMatrix386, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *BTOccurrenceData75) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given BTBSMatrix386 and assigns it to the Transform field.
func (o *BTOccurrenceData75) SetTransform(v BTBSMatrix386) {
	o.Transform = &v
}

func (o BTOccurrenceData75) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureData != nil {
		toSerialize["featureData"] = o.FeatureData
	}
	if o.ForceHighestQualityTessellation != nil {
		toSerialize["forceHighestQualityTessellation"] = o.ForceHighestQualityTessellation
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.IsFixed != nil {
		toSerialize["isFixed"] = o.IsFixed
	}
	if o.IsHidden != nil {
		toSerialize["isHidden"] = o.IsHidden
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableBTOccurrenceData75 struct {
	value *BTOccurrenceData75
	isSet bool
}

func (v NullableBTOccurrenceData75) Get() *BTOccurrenceData75 {
	return v.value
}

func (v *NullableBTOccurrenceData75) Set(val *BTOccurrenceData75) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOccurrenceData75) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOccurrenceData75) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOccurrenceData75(val *BTOccurrenceData75) *NullableBTOccurrenceData75 {
	return &NullableBTOccurrenceData75{value: val, isSet: true}
}

func (v NullableBTOccurrenceData75) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOccurrenceData75) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}