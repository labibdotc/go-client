/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19740-5e8d8b0919a8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTIndividualTableProperties3760 struct for BTIndividualTableProperties3760
type BTIndividualTableProperties3760 struct {
	BtType        *string                   `json:"btType,omitempty"`
	HiddenColumns []BTStringNodeWrapper4224 `json:"hiddenColumns,omitempty"`
	NodeId        *string                   `json:"nodeId,omitempty"`
	Order         []BTStringNodeWrapper4224 `json:"order,omitempty"`
	TableNodeId   *string                   `json:"tableNodeId,omitempty"`
}

// NewBTIndividualTableProperties3760 instantiates a new BTIndividualTableProperties3760 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTIndividualTableProperties3760() *BTIndividualTableProperties3760 {
	this := BTIndividualTableProperties3760{}
	return &this
}

// NewBTIndividualTableProperties3760WithDefaults instantiates a new BTIndividualTableProperties3760 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTIndividualTableProperties3760WithDefaults() *BTIndividualTableProperties3760 {
	this := BTIndividualTableProperties3760{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTIndividualTableProperties3760) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIndividualTableProperties3760) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTIndividualTableProperties3760) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTIndividualTableProperties3760) SetBtType(v string) {
	o.BtType = &v
}

// GetHiddenColumns returns the HiddenColumns field value if set, zero value otherwise.
func (o *BTIndividualTableProperties3760) GetHiddenColumns() []BTStringNodeWrapper4224 {
	if o == nil || o.HiddenColumns == nil {
		var ret []BTStringNodeWrapper4224
		return ret
	}
	return o.HiddenColumns
}

// GetHiddenColumnsOk returns a tuple with the HiddenColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIndividualTableProperties3760) GetHiddenColumnsOk() ([]BTStringNodeWrapper4224, bool) {
	if o == nil || o.HiddenColumns == nil {
		return nil, false
	}
	return o.HiddenColumns, true
}

// HasHiddenColumns returns a boolean if a field has been set.
func (o *BTIndividualTableProperties3760) HasHiddenColumns() bool {
	if o != nil && o.HiddenColumns != nil {
		return true
	}

	return false
}

// SetHiddenColumns gets a reference to the given []BTStringNodeWrapper4224 and assigns it to the HiddenColumns field.
func (o *BTIndividualTableProperties3760) SetHiddenColumns(v []BTStringNodeWrapper4224) {
	o.HiddenColumns = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTIndividualTableProperties3760) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIndividualTableProperties3760) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTIndividualTableProperties3760) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTIndividualTableProperties3760) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *BTIndividualTableProperties3760) GetOrder() []BTStringNodeWrapper4224 {
	if o == nil || o.Order == nil {
		var ret []BTStringNodeWrapper4224
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIndividualTableProperties3760) GetOrderOk() ([]BTStringNodeWrapper4224, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *BTIndividualTableProperties3760) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given []BTStringNodeWrapper4224 and assigns it to the Order field.
func (o *BTIndividualTableProperties3760) SetOrder(v []BTStringNodeWrapper4224) {
	o.Order = v
}

// GetTableNodeId returns the TableNodeId field value if set, zero value otherwise.
func (o *BTIndividualTableProperties3760) GetTableNodeId() string {
	if o == nil || o.TableNodeId == nil {
		var ret string
		return ret
	}
	return *o.TableNodeId
}

// GetTableNodeIdOk returns a tuple with the TableNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIndividualTableProperties3760) GetTableNodeIdOk() (*string, bool) {
	if o == nil || o.TableNodeId == nil {
		return nil, false
	}
	return o.TableNodeId, true
}

// HasTableNodeId returns a boolean if a field has been set.
func (o *BTIndividualTableProperties3760) HasTableNodeId() bool {
	if o != nil && o.TableNodeId != nil {
		return true
	}

	return false
}

// SetTableNodeId gets a reference to the given string and assigns it to the TableNodeId field.
func (o *BTIndividualTableProperties3760) SetTableNodeId(v string) {
	o.TableNodeId = &v
}

func (o BTIndividualTableProperties3760) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.HiddenColumns != nil {
		toSerialize["hiddenColumns"] = o.HiddenColumns
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.TableNodeId != nil {
		toSerialize["tableNodeId"] = o.TableNodeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTIndividualTableProperties3760 struct {
	value *BTIndividualTableProperties3760
	isSet bool
}

func (v NullableBTIndividualTableProperties3760) Get() *BTIndividualTableProperties3760 {
	return v.value
}

func (v *NullableBTIndividualTableProperties3760) Set(val *BTIndividualTableProperties3760) {
	v.value = val
	v.isSet = true
}

func (v NullableBTIndividualTableProperties3760) IsSet() bool {
	return v.isSet
}

func (v *NullableBTIndividualTableProperties3760) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTIndividualTableProperties3760(val *BTIndividualTableProperties3760) *NullableBTIndividualTableProperties3760 {
	return &NullableBTIndividualTableProperties3760{value: val, isSet: true}
}

func (v NullableBTIndividualTableProperties3760) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTIndividualTableProperties3760) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
