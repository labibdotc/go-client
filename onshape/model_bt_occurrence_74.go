/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14306-351f5b17f026
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTOccurrence74 - struct for BTOccurrence74
type BTOccurrence74 struct {
	implBTOccurrence74 interface{}
}

// BTOccurrenceWithFullPartIds1464AsBTOccurrence74 is a convenience function that returns BTOccurrenceWithFullPartIds1464 wrapped in BTOccurrence74
func (o *BTOccurrenceWithFullPartIds1464) AsBTOccurrence74() *BTOccurrence74 {
	return &BTOccurrence74{o}
}

// NewBTOccurrence74 instantiates a new BTOccurrence74 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOccurrence74() *BTOccurrence74 {
	this := BTOccurrence74{Newbase_BTOccurrence74()}
	return &this
}

// NewBTOccurrence74WithDefaults instantiates a new BTOccurrence74 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOccurrence74WithDefaults() *BTOccurrence74 {
	this := BTOccurrence74{Newbase_BTOccurrence74WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOccurrence74) GetBtType() string {
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
func (o *BTOccurrence74) GetBtTypeOk() (*string, bool) {
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
func (o *BTOccurrence74) HasBtType() bool {
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
func (o *BTOccurrence74) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTOccurrence74) GetFullPathAsString() string {
	type getResult interface {
		GetFullPathAsString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullPathAsString()
	} else {
		var de string
		return de
	}
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetFullPathAsStringOk() (*string, bool) {
	type getResult interface {
		GetFullPathAsStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullPathAsStringOk()
	} else {
		return nil, false
	}
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTOccurrence74) HasFullPathAsString() bool {
	type getResult interface {
		HasFullPathAsString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFullPathAsString()
	} else {
		return false
	}
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTOccurrence74) SetFullPathAsString(v string) {
	type getResult interface {
		SetFullPathAsString(v string)
	}

	o.GetActualInstance().(getResult).SetFullPathAsString(v)
}

// GetHeadInstanceId returns the HeadInstanceId field value if set, zero value otherwise.
func (o *BTOccurrence74) GetHeadInstanceId() string {
	type getResult interface {
		GetHeadInstanceId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHeadInstanceId()
	} else {
		var de string
		return de
	}
}

// GetHeadInstanceIdOk returns a tuple with the HeadInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetHeadInstanceIdOk() (*string, bool) {
	type getResult interface {
		GetHeadInstanceIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHeadInstanceIdOk()
	} else {
		return nil, false
	}
}

// HasHeadInstanceId returns a boolean if a field has been set.
func (o *BTOccurrence74) HasHeadInstanceId() bool {
	type getResult interface {
		HasHeadInstanceId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasHeadInstanceId()
	} else {
		return false
	}
}

// SetHeadInstanceId gets a reference to the given string and assigns it to the HeadInstanceId field.
func (o *BTOccurrence74) SetHeadInstanceId(v string) {
	type getResult interface {
		SetHeadInstanceId(v string)
	}

	o.GetActualInstance().(getResult).SetHeadInstanceId(v)
}

// GetOccurrenceWithoutHead returns the OccurrenceWithoutHead field value if set, zero value otherwise.
func (o *BTOccurrence74) GetOccurrenceWithoutHead() BTOccurrence74 {
	type getResult interface {
		GetOccurrenceWithoutHead() BTOccurrence74
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrenceWithoutHead()
	} else {
		var de BTOccurrence74
		return de
	}
}

// GetOccurrenceWithoutHeadOk returns a tuple with the OccurrenceWithoutHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetOccurrenceWithoutHeadOk() (*BTOccurrence74, bool) {
	type getResult interface {
		GetOccurrenceWithoutHeadOk() (*BTOccurrence74, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrenceWithoutHeadOk()
	} else {
		return nil, false
	}
}

// HasOccurrenceWithoutHead returns a boolean if a field has been set.
func (o *BTOccurrence74) HasOccurrenceWithoutHead() bool {
	type getResult interface {
		HasOccurrenceWithoutHead() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOccurrenceWithoutHead()
	} else {
		return false
	}
}

// SetOccurrenceWithoutHead gets a reference to the given BTOccurrence74 and assigns it to the OccurrenceWithoutHead field.
func (o *BTOccurrence74) SetOccurrenceWithoutHead(v BTOccurrence74) {
	type getResult interface {
		SetOccurrenceWithoutHead(v BTOccurrence74)
	}

	o.GetActualInstance().(getResult).SetOccurrenceWithoutHead(v)
}

// GetOccurrenceWithoutTail returns the OccurrenceWithoutTail field value if set, zero value otherwise.
func (o *BTOccurrence74) GetOccurrenceWithoutTail() BTOccurrence74 {
	type getResult interface {
		GetOccurrenceWithoutTail() BTOccurrence74
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrenceWithoutTail()
	} else {
		var de BTOccurrence74
		return de
	}
}

// GetOccurrenceWithoutTailOk returns a tuple with the OccurrenceWithoutTail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetOccurrenceWithoutTailOk() (*BTOccurrence74, bool) {
	type getResult interface {
		GetOccurrenceWithoutTailOk() (*BTOccurrence74, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrenceWithoutTailOk()
	} else {
		return nil, false
	}
}

// HasOccurrenceWithoutTail returns a boolean if a field has been set.
func (o *BTOccurrence74) HasOccurrenceWithoutTail() bool {
	type getResult interface {
		HasOccurrenceWithoutTail() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOccurrenceWithoutTail()
	} else {
		return false
	}
}

// SetOccurrenceWithoutTail gets a reference to the given BTOccurrence74 and assigns it to the OccurrenceWithoutTail field.
func (o *BTOccurrence74) SetOccurrenceWithoutTail(v BTOccurrence74) {
	type getResult interface {
		SetOccurrenceWithoutTail(v BTOccurrence74)
	}

	o.GetActualInstance().(getResult).SetOccurrenceWithoutTail(v)
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *BTOccurrence74) GetParent() BTOccurrence74 {
	type getResult interface {
		GetParent() BTOccurrence74
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParent()
	} else {
		var de BTOccurrence74
		return de
	}
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetParentOk() (*BTOccurrence74, bool) {
	type getResult interface {
		GetParentOk() (*BTOccurrence74, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParentOk()
	} else {
		return nil, false
	}
}

// HasParent returns a boolean if a field has been set.
func (o *BTOccurrence74) HasParent() bool {
	type getResult interface {
		HasParent() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParent()
	} else {
		return false
	}
}

// SetParent gets a reference to the given BTOccurrence74 and assigns it to the Parent field.
func (o *BTOccurrence74) SetParent(v BTOccurrence74) {
	type getResult interface {
		SetParent(v BTOccurrence74)
	}

	o.GetActualInstance().(getResult).SetParent(v)
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTOccurrence74) GetPath() []string {
	type getResult interface {
		GetPath() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPath()
	} else {
		var de []string
		return de
	}
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetPathOk() ([]string, bool) {
	type getResult interface {
		GetPathOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPathOk()
	} else {
		return nil, false
	}
}

// HasPath returns a boolean if a field has been set.
func (o *BTOccurrence74) HasPath() bool {
	type getResult interface {
		HasPath() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPath()
	} else {
		return false
	}
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTOccurrence74) SetPath(v []string) {
	type getResult interface {
		SetPath(v []string)
	}

	o.GetActualInstance().(getResult).SetPath(v)
}

// GetRootOccurrence returns the RootOccurrence field value if set, zero value otherwise.
func (o *BTOccurrence74) GetRootOccurrence() bool {
	type getResult interface {
		GetRootOccurrence() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRootOccurrence()
	} else {
		var de bool
		return de
	}
}

// GetRootOccurrenceOk returns a tuple with the RootOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetRootOccurrenceOk() (*bool, bool) {
	type getResult interface {
		GetRootOccurrenceOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRootOccurrenceOk()
	} else {
		return nil, false
	}
}

// HasRootOccurrence returns a boolean if a field has been set.
func (o *BTOccurrence74) HasRootOccurrence() bool {
	type getResult interface {
		HasRootOccurrence() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRootOccurrence()
	} else {
		return false
	}
}

// SetRootOccurrence gets a reference to the given bool and assigns it to the RootOccurrence field.
func (o *BTOccurrence74) SetRootOccurrence(v bool) {
	type getResult interface {
		SetRootOccurrence(v bool)
	}

	o.GetActualInstance().(getResult).SetRootOccurrence(v)
}

// GetTailInstanceId returns the TailInstanceId field value if set, zero value otherwise.
func (o *BTOccurrence74) GetTailInstanceId() string {
	type getResult interface {
		GetTailInstanceId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTailInstanceId()
	} else {
		var de string
		return de
	}
}

// GetTailInstanceIdOk returns a tuple with the TailInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrence74) GetTailInstanceIdOk() (*string, bool) {
	type getResult interface {
		GetTailInstanceIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTailInstanceIdOk()
	} else {
		return nil, false
	}
}

// HasTailInstanceId returns a boolean if a field has been set.
func (o *BTOccurrence74) HasTailInstanceId() bool {
	type getResult interface {
		HasTailInstanceId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTailInstanceId()
	} else {
		return false
	}
}

// SetTailInstanceId gets a reference to the given string and assigns it to the TailInstanceId field.
func (o *BTOccurrence74) SetTailInstanceId(v string) {
	type getResult interface {
		SetTailInstanceId(v string)
	}

	o.GetActualInstance().(getResult).SetTailInstanceId(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTOccurrence74) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTOccurrenceWithFullPartIds-1464'
	if jsonDict["btType"] == "BTOccurrenceWithFullPartIds-1464" {
		// try to unmarshal JSON data into BTOccurrenceWithFullPartIds1464
		var qr *BTOccurrenceWithFullPartIds1464
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTOccurrence74 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTOccurrence74 = nil
			return fmt.Errorf("Failed to unmarshal BTOccurrence74 as BTOccurrenceWithFullPartIds1464: %s", err.Error())
		}
	}

	var qtx *base_BTOccurrence74
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTOccurrence74 = qtx
		return nil // data stored in dst.base_BTOccurrence74, return on the first match
	} else {
		dst.implBTOccurrence74 = nil
		return fmt.Errorf("Failed to unmarshal BTOccurrence74 as base_BTOccurrence74: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTOccurrence74) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTOccurrence74) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTOccurrence74
}

type NullableBTOccurrence74 struct {
	value *BTOccurrence74
	isSet bool
}

func (v NullableBTOccurrence74) Get() *BTOccurrence74 {
	return v.value
}

func (v *NullableBTOccurrence74) Set(val *BTOccurrence74) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOccurrence74) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOccurrence74) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOccurrence74(val *BTOccurrence74) *NullableBTOccurrence74 {
	return &NullableBTOccurrence74{value: val, isSet: true}
}

func (v NullableBTOccurrence74) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOccurrence74) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTOccurrence74 struct {
	BtType                *string         `json:"btType,omitempty"`
	FullPathAsString      *string         `json:"fullPathAsString,omitempty"`
	HeadInstanceId        *string         `json:"headInstanceId,omitempty"`
	OccurrenceWithoutHead *BTOccurrence74 `json:"occurrenceWithoutHead,omitempty"`
	OccurrenceWithoutTail *BTOccurrence74 `json:"occurrenceWithoutTail,omitempty"`
	Parent                *BTOccurrence74 `json:"parent,omitempty"`
	Path                  []string        `json:"path,omitempty"`
	RootOccurrence        *bool           `json:"rootOccurrence,omitempty"`
	TailInstanceId        *string         `json:"tailInstanceId,omitempty"`
}

// Newbase_BTOccurrence74 instantiates a new base_BTOccurrence74 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTOccurrence74() *base_BTOccurrence74 {
	this := base_BTOccurrence74{}
	return &this
}

// Newbase_BTOccurrence74WithDefaults instantiates a new base_BTOccurrence74 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTOccurrence74WithDefaults() *base_BTOccurrence74 {
	this := base_BTOccurrence74{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTOccurrence74) SetBtType(v string) {
	o.BtType = &v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *base_BTOccurrence74) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetHeadInstanceId returns the HeadInstanceId field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetHeadInstanceId() string {
	if o == nil || o.HeadInstanceId == nil {
		var ret string
		return ret
	}
	return *o.HeadInstanceId
}

// GetHeadInstanceIdOk returns a tuple with the HeadInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetHeadInstanceIdOk() (*string, bool) {
	if o == nil || o.HeadInstanceId == nil {
		return nil, false
	}
	return o.HeadInstanceId, true
}

// HasHeadInstanceId returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasHeadInstanceId() bool {
	if o != nil && o.HeadInstanceId != nil {
		return true
	}

	return false
}

// SetHeadInstanceId gets a reference to the given string and assigns it to the HeadInstanceId field.
func (o *base_BTOccurrence74) SetHeadInstanceId(v string) {
	o.HeadInstanceId = &v
}

// GetOccurrenceWithoutHead returns the OccurrenceWithoutHead field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetOccurrenceWithoutHead() BTOccurrence74 {
	if o == nil || o.OccurrenceWithoutHead == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OccurrenceWithoutHead
}

// GetOccurrenceWithoutHeadOk returns a tuple with the OccurrenceWithoutHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetOccurrenceWithoutHeadOk() (*BTOccurrence74, bool) {
	if o == nil || o.OccurrenceWithoutHead == nil {
		return nil, false
	}
	return o.OccurrenceWithoutHead, true
}

// HasOccurrenceWithoutHead returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasOccurrenceWithoutHead() bool {
	if o != nil && o.OccurrenceWithoutHead != nil {
		return true
	}

	return false
}

// SetOccurrenceWithoutHead gets a reference to the given BTOccurrence74 and assigns it to the OccurrenceWithoutHead field.
func (o *base_BTOccurrence74) SetOccurrenceWithoutHead(v BTOccurrence74) {
	o.OccurrenceWithoutHead = &v
}

// GetOccurrenceWithoutTail returns the OccurrenceWithoutTail field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetOccurrenceWithoutTail() BTOccurrence74 {
	if o == nil || o.OccurrenceWithoutTail == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OccurrenceWithoutTail
}

// GetOccurrenceWithoutTailOk returns a tuple with the OccurrenceWithoutTail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetOccurrenceWithoutTailOk() (*BTOccurrence74, bool) {
	if o == nil || o.OccurrenceWithoutTail == nil {
		return nil, false
	}
	return o.OccurrenceWithoutTail, true
}

// HasOccurrenceWithoutTail returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasOccurrenceWithoutTail() bool {
	if o != nil && o.OccurrenceWithoutTail != nil {
		return true
	}

	return false
}

// SetOccurrenceWithoutTail gets a reference to the given BTOccurrence74 and assigns it to the OccurrenceWithoutTail field.
func (o *base_BTOccurrence74) SetOccurrenceWithoutTail(v BTOccurrence74) {
	o.OccurrenceWithoutTail = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetParent() BTOccurrence74 {
	if o == nil || o.Parent == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetParentOk() (*BTOccurrence74, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given BTOccurrence74 and assigns it to the Parent field.
func (o *base_BTOccurrence74) SetParent(v BTOccurrence74) {
	o.Parent = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *base_BTOccurrence74) SetPath(v []string) {
	o.Path = v
}

// GetRootOccurrence returns the RootOccurrence field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetRootOccurrence() bool {
	if o == nil || o.RootOccurrence == nil {
		var ret bool
		return ret
	}
	return *o.RootOccurrence
}

// GetRootOccurrenceOk returns a tuple with the RootOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetRootOccurrenceOk() (*bool, bool) {
	if o == nil || o.RootOccurrence == nil {
		return nil, false
	}
	return o.RootOccurrence, true
}

// HasRootOccurrence returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasRootOccurrence() bool {
	if o != nil && o.RootOccurrence != nil {
		return true
	}

	return false
}

// SetRootOccurrence gets a reference to the given bool and assigns it to the RootOccurrence field.
func (o *base_BTOccurrence74) SetRootOccurrence(v bool) {
	o.RootOccurrence = &v
}

// GetTailInstanceId returns the TailInstanceId field value if set, zero value otherwise.
func (o *base_BTOccurrence74) GetTailInstanceId() string {
	if o == nil || o.TailInstanceId == nil {
		var ret string
		return ret
	}
	return *o.TailInstanceId
}

// GetTailInstanceIdOk returns a tuple with the TailInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTOccurrence74) GetTailInstanceIdOk() (*string, bool) {
	if o == nil || o.TailInstanceId == nil {
		return nil, false
	}
	return o.TailInstanceId, true
}

// HasTailInstanceId returns a boolean if a field has been set.
func (o *base_BTOccurrence74) HasTailInstanceId() bool {
	if o != nil && o.TailInstanceId != nil {
		return true
	}

	return false
}

// SetTailInstanceId gets a reference to the given string and assigns it to the TailInstanceId field.
func (o *base_BTOccurrence74) SetTailInstanceId(v string) {
	o.TailInstanceId = &v
}

func (o base_BTOccurrence74) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.HeadInstanceId != nil {
		toSerialize["headInstanceId"] = o.HeadInstanceId
	}
	if o.OccurrenceWithoutHead != nil {
		toSerialize["occurrenceWithoutHead"] = o.OccurrenceWithoutHead
	}
	if o.OccurrenceWithoutTail != nil {
		toSerialize["occurrenceWithoutTail"] = o.OccurrenceWithoutTail
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.RootOccurrence != nil {
		toSerialize["rootOccurrence"] = o.RootOccurrence
	}
	if o.TailInstanceId != nil {
		toSerialize["tailInstanceId"] = o.TailInstanceId
	}
	return json.Marshal(toSerialize)
}
