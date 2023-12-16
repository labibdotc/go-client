/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConfiguredPartProperties2645 struct for BTConfiguredPartProperties2645
type BTConfiguredPartProperties2645 struct {
	BtType                       *string                                        `json:"btType,omitempty"`
	NodeId                       *string                                        `json:"nodeId,omitempty"`
	Parts                        []BTPartWithConfiguredProperties2163           `json:"parts,omitempty"`
	PropertyIdToConfiguredTable  *map[string]BTPartWithConfiguredProperties2163 `json:"propertyIdToConfiguredTable,omitempty"`
	SynchronizeToSingleEnumInput *bool                                          `json:"synchronizeToSingleEnumInput,omitempty"`
}

// NewBTConfiguredPartProperties2645 instantiates a new BTConfiguredPartProperties2645 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredPartProperties2645() *BTConfiguredPartProperties2645 {
	this := BTConfiguredPartProperties2645{}
	return &this
}

// NewBTConfiguredPartProperties2645WithDefaults instantiates a new BTConfiguredPartProperties2645 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredPartProperties2645WithDefaults() *BTConfiguredPartProperties2645 {
	this := BTConfiguredPartProperties2645{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredPartProperties2645) SetBtType(v string) {
	o.BtType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTConfiguredPartProperties2645) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetParts() []BTPartWithConfiguredProperties2163 {
	if o == nil || o.Parts == nil {
		var ret []BTPartWithConfiguredProperties2163
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetPartsOk() ([]BTPartWithConfiguredProperties2163, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasParts() bool {
	if o != nil && o.Parts != nil {
		return true
	}

	return false
}

// SetParts gets a reference to the given []BTPartWithConfiguredProperties2163 and assigns it to the Parts field.
func (o *BTConfiguredPartProperties2645) SetParts(v []BTPartWithConfiguredProperties2163) {
	o.Parts = v
}

// GetPropertyIdToConfiguredTable returns the PropertyIdToConfiguredTable field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetPropertyIdToConfiguredTable() map[string]BTPartWithConfiguredProperties2163 {
	if o == nil || o.PropertyIdToConfiguredTable == nil {
		var ret map[string]BTPartWithConfiguredProperties2163
		return ret
	}
	return *o.PropertyIdToConfiguredTable
}

// GetPropertyIdToConfiguredTableOk returns a tuple with the PropertyIdToConfiguredTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetPropertyIdToConfiguredTableOk() (*map[string]BTPartWithConfiguredProperties2163, bool) {
	if o == nil || o.PropertyIdToConfiguredTable == nil {
		return nil, false
	}
	return o.PropertyIdToConfiguredTable, true
}

// HasPropertyIdToConfiguredTable returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasPropertyIdToConfiguredTable() bool {
	if o != nil && o.PropertyIdToConfiguredTable != nil {
		return true
	}

	return false
}

// SetPropertyIdToConfiguredTable gets a reference to the given map[string]BTPartWithConfiguredProperties2163 and assigns it to the PropertyIdToConfiguredTable field.
func (o *BTConfiguredPartProperties2645) SetPropertyIdToConfiguredTable(v map[string]BTPartWithConfiguredProperties2163) {
	o.PropertyIdToConfiguredTable = &v
}

// GetSynchronizeToSingleEnumInput returns the SynchronizeToSingleEnumInput field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetSynchronizeToSingleEnumInput() bool {
	if o == nil || o.SynchronizeToSingleEnumInput == nil {
		var ret bool
		return ret
	}
	return *o.SynchronizeToSingleEnumInput
}

// GetSynchronizeToSingleEnumInputOk returns a tuple with the SynchronizeToSingleEnumInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetSynchronizeToSingleEnumInputOk() (*bool, bool) {
	if o == nil || o.SynchronizeToSingleEnumInput == nil {
		return nil, false
	}
	return o.SynchronizeToSingleEnumInput, true
}

// HasSynchronizeToSingleEnumInput returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasSynchronizeToSingleEnumInput() bool {
	if o != nil && o.SynchronizeToSingleEnumInput != nil {
		return true
	}

	return false
}

// SetSynchronizeToSingleEnumInput gets a reference to the given bool and assigns it to the SynchronizeToSingleEnumInput field.
func (o *BTConfiguredPartProperties2645) SetSynchronizeToSingleEnumInput(v bool) {
	o.SynchronizeToSingleEnumInput = &v
}

func (o BTConfiguredPartProperties2645) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
	}
	if o.PropertyIdToConfiguredTable != nil {
		toSerialize["propertyIdToConfiguredTable"] = o.PropertyIdToConfiguredTable
	}
	if o.SynchronizeToSingleEnumInput != nil {
		toSerialize["synchronizeToSingleEnumInput"] = o.SynchronizeToSingleEnumInput
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredPartProperties2645 struct {
	value *BTConfiguredPartProperties2645
	isSet bool
}

func (v NullableBTConfiguredPartProperties2645) Get() *BTConfiguredPartProperties2645 {
	return v.value
}

func (v *NullableBTConfiguredPartProperties2645) Set(val *BTConfiguredPartProperties2645) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredPartProperties2645) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredPartProperties2645) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredPartProperties2645(val *BTConfiguredPartProperties2645) *NullableBTConfiguredPartProperties2645 {
	return &NullableBTConfiguredPartProperties2645{value: val, isSet: true}
}

func (v NullableBTConfiguredPartProperties2645) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredPartProperties2645) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
