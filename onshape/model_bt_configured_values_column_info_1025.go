/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTConfiguredValuesColumnInfo1025 - struct for BTConfiguredValuesColumnInfo1025
type BTConfiguredValuesColumnInfo1025 struct {
	implBTConfiguredValuesColumnInfo1025 interface{}
}

// BTConfiguredParameterColumnInfo2900AsBTConfiguredValuesColumnInfo1025 is a convenience function that returns BTConfiguredParameterColumnInfo2900 wrapped in BTConfiguredValuesColumnInfo1025
func (o *BTConfiguredParameterColumnInfo2900) AsBTConfiguredValuesColumnInfo1025() *BTConfiguredValuesColumnInfo1025 {
	return &BTConfiguredValuesColumnInfo1025{o}
}

// BTConfiguredSuppressionColumnInfo2498AsBTConfiguredValuesColumnInfo1025 is a convenience function that returns BTConfiguredSuppressionColumnInfo2498 wrapped in BTConfiguredValuesColumnInfo1025
func (o *BTConfiguredSuppressionColumnInfo2498) AsBTConfiguredValuesColumnInfo1025() *BTConfiguredValuesColumnInfo1025 {
	return &BTConfiguredValuesColumnInfo1025{o}
}

// BTConfiguredFeatureColumnInfo1014AsBTConfiguredValuesColumnInfo1025 is a convenience function that returns BTConfiguredFeatureColumnInfo1014 wrapped in BTConfiguredValuesColumnInfo1025
func (o *BTConfiguredFeatureColumnInfo1014) AsBTConfiguredValuesColumnInfo1025() *BTConfiguredValuesColumnInfo1025 {
	return &BTConfiguredValuesColumnInfo1025{o}
}

// BTConfiguredDimensionColumnInfo2168AsBTConfiguredValuesColumnInfo1025 is a convenience function that returns BTConfiguredDimensionColumnInfo2168 wrapped in BTConfiguredValuesColumnInfo1025
func (o *BTConfiguredDimensionColumnInfo2168) AsBTConfiguredValuesColumnInfo1025() *BTConfiguredValuesColumnInfo1025 {
	return &BTConfiguredValuesColumnInfo1025{o}
}

// NewBTConfiguredValuesColumnInfo1025 instantiates a new BTConfiguredValuesColumnInfo1025 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredValuesColumnInfo1025() *BTConfiguredValuesColumnInfo1025 {
	this := BTConfiguredValuesColumnInfo1025{Newbase_BTConfiguredValuesColumnInfo1025()}
	return &this
}

// NewBTConfiguredValuesColumnInfo1025WithDefaults instantiates a new BTConfiguredValuesColumnInfo1025 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredValuesColumnInfo1025WithDefaults() *BTConfiguredValuesColumnInfo1025 {
	this := BTConfiguredValuesColumnInfo1025{Newbase_BTConfiguredValuesColumnInfo1025WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetBtType() string {
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
func (o *BTConfiguredValuesColumnInfo1025) GetBtTypeOk() (*string, bool) {
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
func (o *BTConfiguredValuesColumnInfo1025) HasBtType() bool {
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
func (o *BTConfiguredValuesColumnInfo1025) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTConfiguredValuesColumnInfo1025) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTConfiguredValuesColumnInfo1025) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetSpecification() BTTableColumnSpec1967 {
	type getResult interface {
		GetSpecification() BTTableColumnSpec1967
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpecification()
	} else {
		var de BTTableColumnSpec1967
		return de
	}
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	type getResult interface {
		GetSpecificationOk() (*BTTableColumnSpec1967, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpecificationOk()
	} else {
		return nil, false
	}
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasSpecification() bool {
	type getResult interface {
		HasSpecification() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSpecification()
	} else {
		return false
	}
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTConfiguredValuesColumnInfo1025) SetSpecification(v BTTableColumnSpec1967) {
	type getResult interface {
		SetSpecification(v BTTableColumnSpec1967)
	}

	o.GetActualInstance().(getResult).SetSpecification(v)
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetParentId() string {
	type getResult interface {
		GetParentId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentId()
	} else {
		var de string
		return de
	}
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetParentIdOk() (*string, bool) {
	type getResult interface {
		GetParentIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentIdOk()
	} else {
		return nil, false
	}
}

// HasParentId returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasParentId() bool {
	type getResult interface {
		HasParentId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParentId()
	} else {
		return false
	}
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTConfiguredValuesColumnInfo1025) SetParentId(v string) {
	type getResult interface {
		SetParentId(v string)
	}

	o.GetActualInstance().(getResult).SetParentId(v)
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetParentName() string {
	type getResult interface {
		GetParentName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentName()
	} else {
		var de string
		return de
	}
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetParentNameOk() (*string, bool) {
	type getResult interface {
		GetParentNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentNameOk()
	} else {
		return nil, false
	}
}

// HasParentName returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasParentName() bool {
	type getResult interface {
		HasParentName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParentName()
	} else {
		return false
	}
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *BTConfiguredValuesColumnInfo1025) SetParentName(v string) {
	type getResult interface {
		SetParentName(v string)
	}

	o.GetActualInstance().(getResult).SetParentName(v)
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetParentType() GBTConfiguredParentType {
	type getResult interface {
		GetParentType() GBTConfiguredParentType
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentType()
	} else {
		var de GBTConfiguredParentType
		return de
	}
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetParentTypeOk() (*GBTConfiguredParentType, bool) {
	type getResult interface {
		GetParentTypeOk() (*GBTConfiguredParentType, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentTypeOk()
	} else {
		return nil, false
	}
}

// HasParentType returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasParentType() bool {
	type getResult interface {
		HasParentType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParentType()
	} else {
		return false
	}
}

// SetParentType gets a reference to the given GBTConfiguredParentType and assigns it to the ParentType field.
func (o *BTConfiguredValuesColumnInfo1025) SetParentType(v GBTConfiguredParentType) {
	type getResult interface {
		SetParentType(v GBTConfiguredParentType)
	}

	o.GetActualInstance().(getResult).SetParentType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTConfiguredValuesColumnInfo1025) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTConfiguredDimensionColumnInfo-2168'
	if jsonDict["btType"] == "BTConfiguredDimensionColumnInfo-2168" {
		// try to unmarshal JSON data into BTConfiguredDimensionColumnInfo2168
		var qr *BTConfiguredDimensionColumnInfo2168
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTConfiguredValuesColumnInfo1025 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTConfiguredValuesColumnInfo1025 = nil
			return fmt.Errorf("Failed to unmarshal BTConfiguredValuesColumnInfo1025 as BTConfiguredDimensionColumnInfo2168: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredFeatureColumnInfo-1014'
	if jsonDict["btType"] == "BTConfiguredFeatureColumnInfo-1014" {
		// try to unmarshal JSON data into BTConfiguredFeatureColumnInfo1014
		var qr *BTConfiguredFeatureColumnInfo1014
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTConfiguredValuesColumnInfo1025 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTConfiguredValuesColumnInfo1025 = nil
			return fmt.Errorf("Failed to unmarshal BTConfiguredValuesColumnInfo1025 as BTConfiguredFeatureColumnInfo1014: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredParameterColumnInfo-2900'
	if jsonDict["btType"] == "BTConfiguredParameterColumnInfo-2900" {
		// try to unmarshal JSON data into BTConfiguredParameterColumnInfo2900
		var qr *BTConfiguredParameterColumnInfo2900
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTConfiguredValuesColumnInfo1025 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTConfiguredValuesColumnInfo1025 = nil
			return fmt.Errorf("Failed to unmarshal BTConfiguredValuesColumnInfo1025 as BTConfiguredParameterColumnInfo2900: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredSuppressionColumnInfo-2498'
	if jsonDict["btType"] == "BTConfiguredSuppressionColumnInfo-2498" {
		// try to unmarshal JSON data into BTConfiguredSuppressionColumnInfo2498
		var qr *BTConfiguredSuppressionColumnInfo2498
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTConfiguredValuesColumnInfo1025 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTConfiguredValuesColumnInfo1025 = nil
			return fmt.Errorf("Failed to unmarshal BTConfiguredValuesColumnInfo1025 as BTConfiguredSuppressionColumnInfo2498: %s", err.Error())
		}
	}

	var qtx *base_BTConfiguredValuesColumnInfo1025
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTConfiguredValuesColumnInfo1025 = qtx
		return nil // data stored in dst.base_BTConfiguredValuesColumnInfo1025, return on the first match
	} else {
		dst.implBTConfiguredValuesColumnInfo1025 = nil
		return fmt.Errorf("Failed to unmarshal BTConfiguredValuesColumnInfo1025 as base_BTConfiguredValuesColumnInfo1025: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTConfiguredValuesColumnInfo1025) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTConfiguredValuesColumnInfo1025) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTConfiguredValuesColumnInfo1025
}

type NullableBTConfiguredValuesColumnInfo1025 struct {
	value *BTConfiguredValuesColumnInfo1025
	isSet bool
}

func (v NullableBTConfiguredValuesColumnInfo1025) Get() *BTConfiguredValuesColumnInfo1025 {
	return v.value
}

func (v *NullableBTConfiguredValuesColumnInfo1025) Set(val *BTConfiguredValuesColumnInfo1025) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredValuesColumnInfo1025) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredValuesColumnInfo1025) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredValuesColumnInfo1025(val *BTConfiguredValuesColumnInfo1025) *NullableBTConfiguredValuesColumnInfo1025 {
	return &NullableBTConfiguredValuesColumnInfo1025{value: val, isSet: true}
}

func (v NullableBTConfiguredValuesColumnInfo1025) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredValuesColumnInfo1025) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTConfiguredValuesColumnInfo1025 struct {
	BtType        *string                  `json:"btType,omitempty"`
	Id            *string                  `json:"id,omitempty"`
	NodeId        *string                  `json:"nodeId,omitempty"`
	Specification *BTTableColumnSpec1967   `json:"specification,omitempty"`
	ParentId      *string                  `json:"parentId,omitempty"`
	ParentName    *string                  `json:"parentName,omitempty"`
	ParentType    *GBTConfiguredParentType `json:"parentType,omitempty"`
}

// Newbase_BTConfiguredValuesColumnInfo1025 instantiates a new base_BTConfiguredValuesColumnInfo1025 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTConfiguredValuesColumnInfo1025() *base_BTConfiguredValuesColumnInfo1025 {
	this := base_BTConfiguredValuesColumnInfo1025{}
	return &this
}

// Newbase_BTConfiguredValuesColumnInfo1025WithDefaults instantiates a new base_BTConfiguredValuesColumnInfo1025 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTConfiguredValuesColumnInfo1025WithDefaults() *base_BTConfiguredValuesColumnInfo1025 {
	this := base_BTConfiguredValuesColumnInfo1025{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetParentName() string {
	if o == nil || o.ParentName == nil {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetParentNameOk() (*string, bool) {
	if o == nil || o.ParentName == nil {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasParentName() bool {
	if o != nil && o.ParentName != nil {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *base_BTConfiguredValuesColumnInfo1025) GetParentType() GBTConfiguredParentType {
	if o == nil || o.ParentType == nil {
		var ret GBTConfiguredParentType
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) GetParentTypeOk() (*GBTConfiguredParentType, bool) {
	if o == nil || o.ParentType == nil {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *base_BTConfiguredValuesColumnInfo1025) HasParentType() bool {
	if o != nil && o.ParentType != nil {
		return true
	}

	return false
}

// SetParentType gets a reference to the given GBTConfiguredParentType and assigns it to the ParentType field.
func (o *base_BTConfiguredValuesColumnInfo1025) SetParentType(v GBTConfiguredParentType) {
	o.ParentType = &v
}

func (o base_BTConfiguredValuesColumnInfo1025) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
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
