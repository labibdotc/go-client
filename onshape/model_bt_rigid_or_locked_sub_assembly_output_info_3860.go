/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTRigidOrLockedSubAssemblyOutputInfo3860 - struct for BTRigidOrLockedSubAssemblyOutputInfo3860
type BTRigidOrLockedSubAssemblyOutputInfo3860 struct {
	implBTRigidOrLockedSubAssemblyOutputInfo3860 interface{}
}

// BTLockedSubAssemblyOutputInfo3511AsBTRigidOrLockedSubAssemblyOutputInfo3860 is a convenience function that returns BTLockedSubAssemblyOutputInfo3511 wrapped in BTRigidOrLockedSubAssemblyOutputInfo3860
func (o *BTLockedSubAssemblyOutputInfo3511) AsBTRigidOrLockedSubAssemblyOutputInfo3860() *BTRigidOrLockedSubAssemblyOutputInfo3860 {
	return &BTRigidOrLockedSubAssemblyOutputInfo3860{o}
}

// NewBTRigidOrLockedSubAssemblyOutputInfo3860 instantiates a new BTRigidOrLockedSubAssemblyOutputInfo3860 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRigidOrLockedSubAssemblyOutputInfo3860() *BTRigidOrLockedSubAssemblyOutputInfo3860 {
	this := BTRigidOrLockedSubAssemblyOutputInfo3860{Newbase_BTRigidOrLockedSubAssemblyOutputInfo3860()}
	return &this
}

// NewBTRigidOrLockedSubAssemblyOutputInfo3860WithDefaults instantiates a new BTRigidOrLockedSubAssemblyOutputInfo3860 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRigidOrLockedSubAssemblyOutputInfo3860WithDefaults() *BTRigidOrLockedSubAssemblyOutputInfo3860 {
	this := BTRigidOrLockedSubAssemblyOutputInfo3860{Newbase_BTRigidOrLockedSubAssemblyOutputInfo3860WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetBtType() string {
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
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetBtTypeOk() (*string, bool) {
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
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) HasBtType() bool {
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
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetLocked() bool {
	type getResult interface {
		GetLocked() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocked()
	} else {
		var de bool
		return de
	}
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetLockedOk() (*bool, bool) {
	type getResult interface {
		GetLockedOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLockedOk()
	} else {
		return nil, false
	}
}

// HasLocked returns a boolean if a field has been set.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) HasLocked() bool {
	type getResult interface {
		HasLocked() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLocked()
	} else {
		return false
	}
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) SetLocked(v bool) {
	type getResult interface {
		SetLocked(v bool)
	}

	o.GetActualInstance().(getResult).SetLocked(v)
}

// GetRigid returns the Rigid field value if set, zero value otherwise.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetRigid() bool {
	type getResult interface {
		GetRigid() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRigid()
	} else {
		var de bool
		return de
	}
}

// GetRigidOk returns a tuple with the Rigid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetRigidOk() (*bool, bool) {
	type getResult interface {
		GetRigidOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRigidOk()
	} else {
		return nil, false
	}
}

// HasRigid returns a boolean if a field has been set.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) HasRigid() bool {
	type getResult interface {
		HasRigid() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRigid()
	} else {
		return false
	}
}

// SetRigid gets a reference to the given bool and assigns it to the Rigid field.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) SetRigid(v bool) {
	type getResult interface {
		SetRigid(v bool)
	}

	o.GetActualInstance().(getResult).SetRigid(v)
}

// GetSyncedOutputMVID returns the SyncedOutputMVID field value if set, zero value otherwise.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetSyncedOutputMVID() BTMicroversionId366 {
	type getResult interface {
		GetSyncedOutputMVID() BTMicroversionId366
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSyncedOutputMVID()
	} else {
		var de BTMicroversionId366
		return de
	}
}

// GetSyncedOutputMVIDOk returns a tuple with the SyncedOutputMVID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) GetSyncedOutputMVIDOk() (*BTMicroversionId366, bool) {
	type getResult interface {
		GetSyncedOutputMVIDOk() (*BTMicroversionId366, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSyncedOutputMVIDOk()
	} else {
		return nil, false
	}
}

// HasSyncedOutputMVID returns a boolean if a field has been set.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) HasSyncedOutputMVID() bool {
	type getResult interface {
		HasSyncedOutputMVID() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSyncedOutputMVID()
	} else {
		return false
	}
}

// SetSyncedOutputMVID gets a reference to the given BTMicroversionId366 and assigns it to the SyncedOutputMVID field.
func (o *BTRigidOrLockedSubAssemblyOutputInfo3860) SetSyncedOutputMVID(v BTMicroversionId366) {
	type getResult interface {
		SetSyncedOutputMVID(v BTMicroversionId366)
	}

	o.GetActualInstance().(getResult).SetSyncedOutputMVID(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTRigidOrLockedSubAssemblyOutputInfo3860) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTLockedSubAssemblyOutputInfo-3511'
	if jsonDict["btType"] == "BTLockedSubAssemblyOutputInfo-3511" {
		// try to unmarshal JSON data into BTLockedSubAssemblyOutputInfo3511
		var qr *BTLockedSubAssemblyOutputInfo3511
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTRigidOrLockedSubAssemblyOutputInfo3860 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTRigidOrLockedSubAssemblyOutputInfo3860 = nil
			return fmt.Errorf("failed to unmarshal BTRigidOrLockedSubAssemblyOutputInfo3860 as BTLockedSubAssemblyOutputInfo3511: %s", err.Error())
		}
	}

	var qtx *base_BTRigidOrLockedSubAssemblyOutputInfo3860
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTRigidOrLockedSubAssemblyOutputInfo3860 = qtx
		return nil // data stored in dst.base_BTRigidOrLockedSubAssemblyOutputInfo3860, return on the first match
	} else {
		dst.implBTRigidOrLockedSubAssemblyOutputInfo3860 = nil
		return fmt.Errorf("failed to unmarshal BTRigidOrLockedSubAssemblyOutputInfo3860 as base_BTRigidOrLockedSubAssemblyOutputInfo3860: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTRigidOrLockedSubAssemblyOutputInfo3860) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTRigidOrLockedSubAssemblyOutputInfo3860) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTRigidOrLockedSubAssemblyOutputInfo3860
}

type NullableBTRigidOrLockedSubAssemblyOutputInfo3860 struct {
	value *BTRigidOrLockedSubAssemblyOutputInfo3860
	isSet bool
}

func (v NullableBTRigidOrLockedSubAssemblyOutputInfo3860) Get() *BTRigidOrLockedSubAssemblyOutputInfo3860 {
	return v.value
}

func (v *NullableBTRigidOrLockedSubAssemblyOutputInfo3860) Set(val *BTRigidOrLockedSubAssemblyOutputInfo3860) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRigidOrLockedSubAssemblyOutputInfo3860) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRigidOrLockedSubAssemblyOutputInfo3860) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRigidOrLockedSubAssemblyOutputInfo3860(val *BTRigidOrLockedSubAssemblyOutputInfo3860) *NullableBTRigidOrLockedSubAssemblyOutputInfo3860 {
	return &NullableBTRigidOrLockedSubAssemblyOutputInfo3860{value: val, isSet: true}
}

func (v NullableBTRigidOrLockedSubAssemblyOutputInfo3860) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRigidOrLockedSubAssemblyOutputInfo3860) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTRigidOrLockedSubAssemblyOutputInfo3860 struct {
	// Type of JSON object.
	BtType           *string              `json:"btType,omitempty"`
	Locked           *bool                `json:"locked,omitempty"`
	Rigid            *bool                `json:"rigid,omitempty"`
	SyncedOutputMVID *BTMicroversionId366 `json:"syncedOutputMVID,omitempty"`
}

// Newbase_BTRigidOrLockedSubAssemblyOutputInfo3860 instantiates a new base_BTRigidOrLockedSubAssemblyOutputInfo3860 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTRigidOrLockedSubAssemblyOutputInfo3860() *base_BTRigidOrLockedSubAssemblyOutputInfo3860 {
	this := base_BTRigidOrLockedSubAssemblyOutputInfo3860{}
	return &this
}

// Newbase_BTRigidOrLockedSubAssemblyOutputInfo3860WithDefaults instantiates a new base_BTRigidOrLockedSubAssemblyOutputInfo3860 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTRigidOrLockedSubAssemblyOutputInfo3860WithDefaults() *base_BTRigidOrLockedSubAssemblyOutputInfo3860 {
	this := base_BTRigidOrLockedSubAssemblyOutputInfo3860{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) SetBtType(v string) {
	o.BtType = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) SetLocked(v bool) {
	o.Locked = &v
}

// GetRigid returns the Rigid field value if set, zero value otherwise.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetRigid() bool {
	if o == nil || o.Rigid == nil {
		var ret bool
		return ret
	}
	return *o.Rigid
}

// GetRigidOk returns a tuple with the Rigid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetRigidOk() (*bool, bool) {
	if o == nil || o.Rigid == nil {
		return nil, false
	}
	return o.Rigid, true
}

// HasRigid returns a boolean if a field has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) HasRigid() bool {
	if o != nil && o.Rigid != nil {
		return true
	}

	return false
}

// SetRigid gets a reference to the given bool and assigns it to the Rigid field.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) SetRigid(v bool) {
	o.Rigid = &v
}

// GetSyncedOutputMVID returns the SyncedOutputMVID field value if set, zero value otherwise.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetSyncedOutputMVID() BTMicroversionId366 {
	if o == nil || o.SyncedOutputMVID == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.SyncedOutputMVID
}

// GetSyncedOutputMVIDOk returns a tuple with the SyncedOutputMVID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) GetSyncedOutputMVIDOk() (*BTMicroversionId366, bool) {
	if o == nil || o.SyncedOutputMVID == nil {
		return nil, false
	}
	return o.SyncedOutputMVID, true
}

// HasSyncedOutputMVID returns a boolean if a field has been set.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) HasSyncedOutputMVID() bool {
	if o != nil && o.SyncedOutputMVID != nil {
		return true
	}

	return false
}

// SetSyncedOutputMVID gets a reference to the given BTMicroversionId366 and assigns it to the SyncedOutputMVID field.
func (o *base_BTRigidOrLockedSubAssemblyOutputInfo3860) SetSyncedOutputMVID(v BTMicroversionId366) {
	o.SyncedOutputMVID = &v
}

func (o base_BTRigidOrLockedSubAssemblyOutputInfo3860) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Rigid != nil {
		toSerialize["rigid"] = o.Rigid
	}
	if o.SyncedOutputMVID != nil {
		toSerialize["syncedOutputMVID"] = o.SyncedOutputMVID
	}
	return json.Marshal(toSerialize)
}
