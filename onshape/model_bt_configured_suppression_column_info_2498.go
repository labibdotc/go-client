/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6159-a7ef074944fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConfiguredSuppressionColumnInfo2498 struct for BTConfiguredSuppressionColumnInfo2498
type BTConfiguredSuppressionColumnInfo2498 struct {
	Id            *string                `json:"id,omitempty"`
	NodeId        *string                `json:"nodeId,omitempty"`
	Specification *BTTableColumnSpec1967 `json:"specification,omitempty"`
	BtType        *string                `json:"btType,omitempty"`
	ParentId      *string                `json:"parentId,omitempty"`
	ParentName    *string                `json:"parentName,omitempty"`
	ParentType    *string                `json:"parentType,omitempty"`
}

// NewBTConfiguredSuppressionColumnInfo2498 instantiates a new BTConfiguredSuppressionColumnInfo2498 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredSuppressionColumnInfo2498() *BTConfiguredSuppressionColumnInfo2498 {
	this := BTConfiguredSuppressionColumnInfo2498{}
	return &this
}

// NewBTConfiguredSuppressionColumnInfo2498WithDefaults instantiates a new BTConfiguredSuppressionColumnInfo2498 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredSuppressionColumnInfo2498WithDefaults() *BTConfiguredSuppressionColumnInfo2498 {
	this := BTConfiguredSuppressionColumnInfo2498{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetBtType(v string) {
	o.BtType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetParentName() string {
	if o == nil || o.ParentName == nil {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetParentNameOk() (*string, bool) {
	if o == nil || o.ParentName == nil {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasParentName() bool {
	if o != nil && o.ParentName != nil {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetParentType() string {
	if o == nil || o.ParentType == nil {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetParentTypeOk() (*string, bool) {
	if o == nil || o.ParentType == nil {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasParentType() bool {
	if o != nil && o.ParentType != nil {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetParentType(v string) {
	o.ParentType = &v
}

func (o BTConfiguredSuppressionColumnInfo2498) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ParentName != nil {
		toSerialize["parentName"] = o.ParentName
	}
	if o.ParentType != nil {
		toSerialize["parentType"] = o.ParentType
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredSuppressionColumnInfo2498 struct {
	value *BTConfiguredSuppressionColumnInfo2498
	isSet bool
}

func (v NullableBTConfiguredSuppressionColumnInfo2498) Get() *BTConfiguredSuppressionColumnInfo2498 {
	return v.value
}

func (v *NullableBTConfiguredSuppressionColumnInfo2498) Set(val *BTConfiguredSuppressionColumnInfo2498) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredSuppressionColumnInfo2498) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredSuppressionColumnInfo2498) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredSuppressionColumnInfo2498(val *BTConfiguredSuppressionColumnInfo2498) *NullableBTConfiguredSuppressionColumnInfo2498 {
	return &NullableBTConfiguredSuppressionColumnInfo2498{value: val, isSet: true}
}

func (v NullableBTConfiguredSuppressionColumnInfo2498) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredSuppressionColumnInfo2498) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
