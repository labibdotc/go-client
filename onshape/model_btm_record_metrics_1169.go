/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7489-3fe01473c812
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMRecordMetrics1169 struct for BTMRecordMetrics1169
type BTMRecordMetrics1169 struct {
	BtType             *string                 `json:"btType,omitempty"`
	ImportMicroversion *string                 `json:"importMicroversion,omitempty"`
	NodeId             *string                 `json:"nodeId,omitempty"`
	DoBodyValidation   *bool                   `json:"doBodyValidation,omitempty"`
	PreviousFeatureId  *string                 `json:"previousFeatureId,omitempty"`
	References         []BTMIndividualQuery138 `json:"references,omitempty"`
	UseLatestBehavior  *bool                   `json:"useLatestBehavior,omitempty"`
}

// NewBTMRecordMetrics1169 instantiates a new BTMRecordMetrics1169 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMRecordMetrics1169() *BTMRecordMetrics1169 {
	this := BTMRecordMetrics1169{}
	return &this
}

// NewBTMRecordMetrics1169WithDefaults instantiates a new BTMRecordMetrics1169 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMRecordMetrics1169WithDefaults() *BTMRecordMetrics1169 {
	this := BTMRecordMetrics1169{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMRecordMetrics1169) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMRecordMetrics1169) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMRecordMetrics1169) SetNodeId(v string) {
	o.NodeId = &v
}

// GetDoBodyValidation returns the DoBodyValidation field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetDoBodyValidation() bool {
	if o == nil || o.DoBodyValidation == nil {
		var ret bool
		return ret
	}
	return *o.DoBodyValidation
}

// GetDoBodyValidationOk returns a tuple with the DoBodyValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetDoBodyValidationOk() (*bool, bool) {
	if o == nil || o.DoBodyValidation == nil {
		return nil, false
	}
	return o.DoBodyValidation, true
}

// HasDoBodyValidation returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasDoBodyValidation() bool {
	if o != nil && o.DoBodyValidation != nil {
		return true
	}

	return false
}

// SetDoBodyValidation gets a reference to the given bool and assigns it to the DoBodyValidation field.
func (o *BTMRecordMetrics1169) SetDoBodyValidation(v bool) {
	o.DoBodyValidation = &v
}

// GetPreviousFeatureId returns the PreviousFeatureId field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetPreviousFeatureId() string {
	if o == nil || o.PreviousFeatureId == nil {
		var ret string
		return ret
	}
	return *o.PreviousFeatureId
}

// GetPreviousFeatureIdOk returns a tuple with the PreviousFeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetPreviousFeatureIdOk() (*string, bool) {
	if o == nil || o.PreviousFeatureId == nil {
		return nil, false
	}
	return o.PreviousFeatureId, true
}

// HasPreviousFeatureId returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasPreviousFeatureId() bool {
	if o != nil && o.PreviousFeatureId != nil {
		return true
	}

	return false
}

// SetPreviousFeatureId gets a reference to the given string and assigns it to the PreviousFeatureId field.
func (o *BTMRecordMetrics1169) SetPreviousFeatureId(v string) {
	o.PreviousFeatureId = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetReferences() []BTMIndividualQuery138 {
	if o == nil || o.References == nil {
		var ret []BTMIndividualQuery138
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetReferencesOk() ([]BTMIndividualQuery138, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []BTMIndividualQuery138 and assigns it to the References field.
func (o *BTMRecordMetrics1169) SetReferences(v []BTMIndividualQuery138) {
	o.References = v
}

// GetUseLatestBehavior returns the UseLatestBehavior field value if set, zero value otherwise.
func (o *BTMRecordMetrics1169) GetUseLatestBehavior() bool {
	if o == nil || o.UseLatestBehavior == nil {
		var ret bool
		return ret
	}
	return *o.UseLatestBehavior
}

// GetUseLatestBehaviorOk returns a tuple with the UseLatestBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMRecordMetrics1169) GetUseLatestBehaviorOk() (*bool, bool) {
	if o == nil || o.UseLatestBehavior == nil {
		return nil, false
	}
	return o.UseLatestBehavior, true
}

// HasUseLatestBehavior returns a boolean if a field has been set.
func (o *BTMRecordMetrics1169) HasUseLatestBehavior() bool {
	if o != nil && o.UseLatestBehavior != nil {
		return true
	}

	return false
}

// SetUseLatestBehavior gets a reference to the given bool and assigns it to the UseLatestBehavior field.
func (o *BTMRecordMetrics1169) SetUseLatestBehavior(v bool) {
	o.UseLatestBehavior = &v
}

func (o BTMRecordMetrics1169) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.DoBodyValidation != nil {
		toSerialize["doBodyValidation"] = o.DoBodyValidation
	}
	if o.PreviousFeatureId != nil {
		toSerialize["previousFeatureId"] = o.PreviousFeatureId
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.UseLatestBehavior != nil {
		toSerialize["useLatestBehavior"] = o.UseLatestBehavior
	}
	return json.Marshal(toSerialize)
}

type NullableBTMRecordMetrics1169 struct {
	value *BTMRecordMetrics1169
	isSet bool
}

func (v NullableBTMRecordMetrics1169) Get() *BTMRecordMetrics1169 {
	return v.value
}

func (v *NullableBTMRecordMetrics1169) Set(val *BTMRecordMetrics1169) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMRecordMetrics1169) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMRecordMetrics1169) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMRecordMetrics1169(val *BTMRecordMetrics1169) *NullableBTMRecordMetrics1169 {
	return &NullableBTMRecordMetrics1169{value: val, isSet: true}
}

func (v NullableBTMRecordMetrics1169) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMRecordMetrics1169) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
