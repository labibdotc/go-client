/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7108-42dac6f29840
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartProperties293 struct for BTPartProperties293
type BTPartProperties293 struct {
	BtType                 *string                         `json:"btType,omitempty"`
	ConfiguredParts        *BTConfiguredPartProperties2645 `json:"configuredParts,omitempty"`
	IdentityIdToQueryIndex *map[string]int32               `json:"identityIdToQueryIndex,omitempty"`
	NodeId                 *string                         `json:"nodeId,omitempty"`
	Parts                  []BTOnePartProperties230        `json:"parts,omitempty"`
	RoughBytesEstimate     *int64                          `json:"roughBytesEstimate,omitempty"`
	TessellationProperties *BTTessellationProperties927    `json:"tessellationProperties,omitempty"`
}

// NewBTPartProperties293 instantiates a new BTPartProperties293 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartProperties293() *BTPartProperties293 {
	this := BTPartProperties293{}
	return &this
}

// NewBTPartProperties293WithDefaults instantiates a new BTPartProperties293 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartProperties293WithDefaults() *BTPartProperties293 {
	this := BTPartProperties293{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartProperties293) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartProperties293) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartProperties293) SetBtType(v string) {
	o.BtType = &v
}

// GetConfiguredParts returns the ConfiguredParts field value if set, zero value otherwise.
func (o *BTPartProperties293) GetConfiguredParts() BTConfiguredPartProperties2645 {
	if o == nil || o.ConfiguredParts == nil {
		var ret BTConfiguredPartProperties2645
		return ret
	}
	return *o.ConfiguredParts
}

// GetConfiguredPartsOk returns a tuple with the ConfiguredParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetConfiguredPartsOk() (*BTConfiguredPartProperties2645, bool) {
	if o == nil || o.ConfiguredParts == nil {
		return nil, false
	}
	return o.ConfiguredParts, true
}

// HasConfiguredParts returns a boolean if a field has been set.
func (o *BTPartProperties293) HasConfiguredParts() bool {
	if o != nil && o.ConfiguredParts != nil {
		return true
	}

	return false
}

// SetConfiguredParts gets a reference to the given BTConfiguredPartProperties2645 and assigns it to the ConfiguredParts field.
func (o *BTPartProperties293) SetConfiguredParts(v BTConfiguredPartProperties2645) {
	o.ConfiguredParts = &v
}

// GetIdentityIdToQueryIndex returns the IdentityIdToQueryIndex field value if set, zero value otherwise.
func (o *BTPartProperties293) GetIdentityIdToQueryIndex() map[string]int32 {
	if o == nil || o.IdentityIdToQueryIndex == nil {
		var ret map[string]int32
		return ret
	}
	return *o.IdentityIdToQueryIndex
}

// GetIdentityIdToQueryIndexOk returns a tuple with the IdentityIdToQueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetIdentityIdToQueryIndexOk() (*map[string]int32, bool) {
	if o == nil || o.IdentityIdToQueryIndex == nil {
		return nil, false
	}
	return o.IdentityIdToQueryIndex, true
}

// HasIdentityIdToQueryIndex returns a boolean if a field has been set.
func (o *BTPartProperties293) HasIdentityIdToQueryIndex() bool {
	if o != nil && o.IdentityIdToQueryIndex != nil {
		return true
	}

	return false
}

// SetIdentityIdToQueryIndex gets a reference to the given map[string]int32 and assigns it to the IdentityIdToQueryIndex field.
func (o *BTPartProperties293) SetIdentityIdToQueryIndex(v map[string]int32) {
	o.IdentityIdToQueryIndex = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPartProperties293) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPartProperties293) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPartProperties293) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *BTPartProperties293) GetParts() []BTOnePartProperties230 {
	if o == nil || o.Parts == nil {
		var ret []BTOnePartProperties230
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetPartsOk() ([]BTOnePartProperties230, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *BTPartProperties293) HasParts() bool {
	if o != nil && o.Parts != nil {
		return true
	}

	return false
}

// SetParts gets a reference to the given []BTOnePartProperties230 and assigns it to the Parts field.
func (o *BTPartProperties293) SetParts(v []BTOnePartProperties230) {
	o.Parts = v
}

// GetRoughBytesEstimate returns the RoughBytesEstimate field value if set, zero value otherwise.
func (o *BTPartProperties293) GetRoughBytesEstimate() int64 {
	if o == nil || o.RoughBytesEstimate == nil {
		var ret int64
		return ret
	}
	return *o.RoughBytesEstimate
}

// GetRoughBytesEstimateOk returns a tuple with the RoughBytesEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetRoughBytesEstimateOk() (*int64, bool) {
	if o == nil || o.RoughBytesEstimate == nil {
		return nil, false
	}
	return o.RoughBytesEstimate, true
}

// HasRoughBytesEstimate returns a boolean if a field has been set.
func (o *BTPartProperties293) HasRoughBytesEstimate() bool {
	if o != nil && o.RoughBytesEstimate != nil {
		return true
	}

	return false
}

// SetRoughBytesEstimate gets a reference to the given int64 and assigns it to the RoughBytesEstimate field.
func (o *BTPartProperties293) SetRoughBytesEstimate(v int64) {
	o.RoughBytesEstimate = &v
}

// GetTessellationProperties returns the TessellationProperties field value if set, zero value otherwise.
func (o *BTPartProperties293) GetTessellationProperties() BTTessellationProperties927 {
	if o == nil || o.TessellationProperties == nil {
		var ret BTTessellationProperties927
		return ret
	}
	return *o.TessellationProperties
}

// GetTessellationPropertiesOk returns a tuple with the TessellationProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartProperties293) GetTessellationPropertiesOk() (*BTTessellationProperties927, bool) {
	if o == nil || o.TessellationProperties == nil {
		return nil, false
	}
	return o.TessellationProperties, true
}

// HasTessellationProperties returns a boolean if a field has been set.
func (o *BTPartProperties293) HasTessellationProperties() bool {
	if o != nil && o.TessellationProperties != nil {
		return true
	}

	return false
}

// SetTessellationProperties gets a reference to the given BTTessellationProperties927 and assigns it to the TessellationProperties field.
func (o *BTPartProperties293) SetTessellationProperties(v BTTessellationProperties927) {
	o.TessellationProperties = &v
}

func (o BTPartProperties293) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfiguredParts != nil {
		toSerialize["configuredParts"] = o.ConfiguredParts
	}
	if o.IdentityIdToQueryIndex != nil {
		toSerialize["identityIdToQueryIndex"] = o.IdentityIdToQueryIndex
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
	}
	if o.RoughBytesEstimate != nil {
		toSerialize["roughBytesEstimate"] = o.RoughBytesEstimate
	}
	if o.TessellationProperties != nil {
		toSerialize["tessellationProperties"] = o.TessellationProperties
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartProperties293 struct {
	value *BTPartProperties293
	isSet bool
}

func (v NullableBTPartProperties293) Get() *BTPartProperties293 {
	return v.value
}

func (v *NullableBTPartProperties293) Set(val *BTPartProperties293) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartProperties293) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartProperties293) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartProperties293(val *BTPartProperties293) *NullableBTPartProperties293 {
	return &NullableBTPartProperties293{value: val, isSet: true}
}

func (v NullableBTPartProperties293) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartProperties293) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
