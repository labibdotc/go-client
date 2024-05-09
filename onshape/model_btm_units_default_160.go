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

// BTMUnitsDefault160 struct for BTMUnitsDefault160
type BTMUnitsDefault160 struct {
	BTMNode19
	BtType *string `json:"btType,omitempty"`
	// Microversion that resulted from the import.
	ImportMicroversion *string            `json:"importMicroversion,omitempty"`
	NodeId             *string            `json:"nodeId,omitempty"`
	Units              *map[string]string `json:"units,omitempty"`
}

// NewBTMUnitsDefault160 instantiates a new BTMUnitsDefault160 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMUnitsDefault160() *BTMUnitsDefault160 {
	this := BTMUnitsDefault160{}
	return &this
}

// NewBTMUnitsDefault160WithDefaults instantiates a new BTMUnitsDefault160 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMUnitsDefault160WithDefaults() *BTMUnitsDefault160 {
	this := BTMUnitsDefault160{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMUnitsDefault160) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUnitsDefault160) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMUnitsDefault160) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMUnitsDefault160) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMUnitsDefault160) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUnitsDefault160) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMUnitsDefault160) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMUnitsDefault160) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMUnitsDefault160) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUnitsDefault160) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMUnitsDefault160) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMUnitsDefault160) SetNodeId(v string) {
	o.NodeId = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTMUnitsDefault160) GetUnits() map[string]string {
	if o == nil || o.Units == nil {
		var ret map[string]string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMUnitsDefault160) GetUnitsOk() (*map[string]string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTMUnitsDefault160) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given map[string]string and assigns it to the Units field.
func (o *BTMUnitsDefault160) SetUnits(v map[string]string) {
	o.Units = &v
}

func (o BTMUnitsDefault160) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMNode19, errBTMNode19 := json.Marshal(o.BTMNode19)
	if errBTMNode19 != nil {
		return []byte{}, errBTMNode19
	}
	errBTMNode19 = json.Unmarshal([]byte(serializedBTMNode19), &toSerialize)
	if errBTMNode19 != nil {
		return []byte{}, errBTMNode19
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableBTMUnitsDefault160 struct {
	value *BTMUnitsDefault160
	isSet bool
}

func (v NullableBTMUnitsDefault160) Get() *BTMUnitsDefault160 {
	return v.value
}

func (v *NullableBTMUnitsDefault160) Set(val *BTMUnitsDefault160) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMUnitsDefault160) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMUnitsDefault160) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMUnitsDefault160(val *BTMUnitsDefault160) *NullableBTMUnitsDefault160 {
	return &NullableBTMUnitsDefault160{value: val, isSet: true}
}

func (v NullableBTMUnitsDefault160) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMUnitsDefault160) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
