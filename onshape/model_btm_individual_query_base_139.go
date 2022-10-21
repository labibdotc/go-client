/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7010-841c6a8f62e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMIndividualQueryBase139 - struct for BTMIndividualQueryBase139
type BTMIndividualQueryBase139 struct {
	implBTMIndividualQueryBase139 interface{}
}

// BTMIndividualCreatedByQuery137AsBTMIndividualQueryBase139 is a convenience function that returns BTMIndividualCreatedByQuery137 wrapped in BTMIndividualQueryBase139
func (o *BTMIndividualCreatedByQuery137) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// BTMInContextQuery2254AsBTMIndividualQueryBase139 is a convenience function that returns BTMInContextQuery2254 wrapped in BTMIndividualQueryBase139
func (o *BTMInContextQuery2254) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// BTMIndividualSketchRegionQuery140AsBTMIndividualQueryBase139 is a convenience function that returns BTMIndividualSketchRegionQuery140 wrapped in BTMIndividualQueryBase139
func (o *BTMIndividualSketchRegionQuery140) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// BTMIndividualQueryWithOccurrenceBase904AsBTMIndividualQueryBase139 is a convenience function that returns BTMIndividualQueryWithOccurrenceBase904 wrapped in BTMIndividualQueryBase139
func (o *BTMIndividualQueryWithOccurrenceBase904) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// BTMIndividualQuery138AsBTMIndividualQueryBase139 is a convenience function that returns BTMIndividualQuery138 wrapped in BTMIndividualQueryBase139
func (o *BTMIndividualQuery138) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// BTMIndividualCoEdgeQuery1332AsBTMIndividualQueryBase139 is a convenience function that returns BTMIndividualCoEdgeQuery1332 wrapped in BTMIndividualQueryBase139
func (o *BTMIndividualCoEdgeQuery1332) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// BTMIndividualSketchUniqueVerticesQuery1472AsBTMIndividualQueryBase139 is a convenience function that returns BTMIndividualSketchUniqueVerticesQuery1472 wrapped in BTMIndividualQueryBase139
func (o *BTMIndividualSketchUniqueVerticesQuery1472) AsBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	return &BTMIndividualQueryBase139{o}
}

// NewBTMIndividualQueryBase139 instantiates a new BTMIndividualQueryBase139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualQueryBase139() *BTMIndividualQueryBase139 {
	this := BTMIndividualQueryBase139{Newbase_BTMIndividualQueryBase139()}
	return &this
}

// NewBTMIndividualQueryBase139WithDefaults instantiates a new BTMIndividualQueryBase139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualQueryBase139WithDefaults() *BTMIndividualQueryBase139 {
	this := BTMIndividualQueryBase139{Newbase_BTMIndividualQueryBase139WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetBtType() string {
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
func (o *BTMIndividualQueryBase139) GetBtTypeOk() (*string, bool) {
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
func (o *BTMIndividualQueryBase139) HasBtType() bool {
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
func (o *BTMIndividualQueryBase139) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetDeterministicIdList() BTMIndividualQueryBase139 {
	type getResult interface {
		GetDeterministicIdList() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdList()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryBase139) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdListOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualQueryBase139) HasDeterministicIdList() bool {
	type getResult interface {
		HasDeterministicIdList() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIdList()
	} else {
		return false
	}
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualQueryBase139) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetDeterministicIdList(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetDeterministicIdList(v)
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetDeterministicIds() []string {
	type getResult interface {
		GetDeterministicIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIds()
	} else {
		var de []string
		return de
	}
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryBase139) GetDeterministicIdsOk() ([]string, bool) {
	type getResult interface {
		GetDeterministicIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdsOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualQueryBase139) HasDeterministicIds() bool {
	type getResult interface {
		HasDeterministicIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIds()
	} else {
		return false
	}
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualQueryBase139) SetDeterministicIds(v []string) {
	type getResult interface {
		SetDeterministicIds(v []string)
	}

	o.GetActualInstance().(getResult).SetDeterministicIds(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryBase139) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualQueryBase139) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualQueryBase139) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetNodeId() string {
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
func (o *BTMIndividualQueryBase139) GetNodeIdOk() (*string, bool) {
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
func (o *BTMIndividualQueryBase139) HasNodeId() bool {
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
func (o *BTMIndividualQueryBase139) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetQuery() BTMIndividualQueryBase139 {
	type getResult interface {
		GetQuery() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQuery()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryBase139) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetQueryOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryOk()
	} else {
		return nil, false
	}
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualQueryBase139) HasQuery() bool {
	type getResult interface {
		HasQuery() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQuery()
	} else {
		return false
	}
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualQueryBase139) SetQuery(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetQuery(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetQuery(v)
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualQueryBase139) GetQueryString() string {
	type getResult interface {
		GetQueryString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryString()
	} else {
		var de string
		return de
	}
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryBase139) GetQueryStringOk() (*string, bool) {
	type getResult interface {
		GetQueryStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryStringOk()
	} else {
		return nil, false
	}
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualQueryBase139) HasQueryString() bool {
	type getResult interface {
		HasQueryString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQueryString()
	} else {
		return false
	}
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualQueryBase139) SetQueryString(v string) {
	type getResult interface {
		SetQueryString(v string)
	}

	o.GetActualInstance().(getResult).SetQueryString(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMIndividualQueryBase139) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMInContextQuery-2254'
	if jsonDict["btType"] == "BTMInContextQuery-2254" {
		// try to unmarshal JSON data into BTMInContextQuery2254
		var qr *BTMInContextQuery2254
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMInContextQuery2254: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualCoEdgeQuery-1332'
	if jsonDict["btType"] == "BTMIndividualCoEdgeQuery-1332" {
		// try to unmarshal JSON data into BTMIndividualCoEdgeQuery1332
		var qr *BTMIndividualCoEdgeQuery1332
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMIndividualCoEdgeQuery1332: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualCreatedByQuery-137'
	if jsonDict["btType"] == "BTMIndividualCreatedByQuery-137" {
		// try to unmarshal JSON data into BTMIndividualCreatedByQuery137
		var qr *BTMIndividualCreatedByQuery137
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMIndividualCreatedByQuery137: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualQuery-138'
	if jsonDict["btType"] == "BTMIndividualQuery-138" {
		// try to unmarshal JSON data into BTMIndividualQuery138
		var qr *BTMIndividualQuery138
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMIndividualQuery138: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualQueryWithOccurrenceBase-904'
	if jsonDict["btType"] == "BTMIndividualQueryWithOccurrenceBase-904" {
		// try to unmarshal JSON data into BTMIndividualQueryWithOccurrenceBase904
		var qr *BTMIndividualQueryWithOccurrenceBase904
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMIndividualQueryWithOccurrenceBase904: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualSketchRegionQuery-140'
	if jsonDict["btType"] == "BTMIndividualSketchRegionQuery-140" {
		// try to unmarshal JSON data into BTMIndividualSketchRegionQuery140
		var qr *BTMIndividualSketchRegionQuery140
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMIndividualSketchRegionQuery140: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualSketchUniqueVerticesQuery-1472'
	if jsonDict["btType"] == "BTMIndividualSketchUniqueVerticesQuery-1472" {
		// try to unmarshal JSON data into BTMIndividualSketchUniqueVerticesQuery1472
		var qr *BTMIndividualSketchUniqueVerticesQuery1472
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryBase139 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryBase139 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as BTMIndividualSketchUniqueVerticesQuery1472: %s", err.Error())
		}
	}

	var qtx *base_BTMIndividualQueryBase139
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMIndividualQueryBase139 = qtx
		return nil // data stored in dst.base_BTMIndividualQueryBase139, return on the first match
	} else {
		dst.implBTMIndividualQueryBase139 = nil
		return fmt.Errorf("Failed to unmarshal BTMIndividualQueryBase139 as base_BTMIndividualQueryBase139: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMIndividualQueryBase139) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMIndividualQueryBase139) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMIndividualQueryBase139
}

type NullableBTMIndividualQueryBase139 struct {
	value *BTMIndividualQueryBase139
	isSet bool
}

func (v NullableBTMIndividualQueryBase139) Get() *BTMIndividualQueryBase139 {
	return v.value
}

func (v *NullableBTMIndividualQueryBase139) Set(val *BTMIndividualQueryBase139) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualQueryBase139) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualQueryBase139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualQueryBase139(val *BTMIndividualQueryBase139) *NullableBTMIndividualQueryBase139 {
	return &NullableBTMIndividualQueryBase139{value: val, isSet: true}
}

func (v NullableBTMIndividualQueryBase139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualQueryBase139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMIndividualQueryBase139 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
}

// Newbase_BTMIndividualQueryBase139 instantiates a new base_BTMIndividualQueryBase139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMIndividualQueryBase139() *base_BTMIndividualQueryBase139 {
	this := base_BTMIndividualQueryBase139{}
	return &this
}

// Newbase_BTMIndividualQueryBase139WithDefaults instantiates a new base_BTMIndividualQueryBase139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMIndividualQueryBase139WithDefaults() *base_BTMIndividualQueryBase139 {
	this := base_BTMIndividualQueryBase139{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMIndividualQueryBase139) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *base_BTMIndividualQueryBase139) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *base_BTMIndividualQueryBase139) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMIndividualQueryBase139) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMIndividualQueryBase139) SetNodeId(v string) {
	o.NodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *base_BTMIndividualQueryBase139) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryBase139) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryBase139) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryBase139) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *base_BTMIndividualQueryBase139) SetQueryString(v string) {
	o.QueryString = &v
}

func (o base_BTMIndividualQueryBase139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	return json.Marshal(toSerialize)
}
