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

// BTLockedSubAssemblyOutputInfo3511 struct for BTLockedSubAssemblyOutputInfo3511
type BTLockedSubAssemblyOutputInfo3511 struct {
	BtType           *string                  `json:"btType,omitempty"`
	Locked           *bool                    `json:"locked,omitempty"`
	Rigid            *bool                    `json:"rigid,omitempty"`
	SyncedOutputMVID *BTMicroversionId366     `json:"syncedOutputMVID,omitempty"`
	LockInfo         *BTLockedSubAssembly4590 `json:"lockInfo,omitempty"`
}

// NewBTLockedSubAssemblyOutputInfo3511 instantiates a new BTLockedSubAssemblyOutputInfo3511 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLockedSubAssemblyOutputInfo3511() *BTLockedSubAssemblyOutputInfo3511 {
	this := BTLockedSubAssemblyOutputInfo3511{}
	return &this
}

// NewBTLockedSubAssemblyOutputInfo3511WithDefaults instantiates a new BTLockedSubAssemblyOutputInfo3511 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLockedSubAssemblyOutputInfo3511WithDefaults() *BTLockedSubAssemblyOutputInfo3511 {
	this := BTLockedSubAssemblyOutputInfo3511{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLockedSubAssemblyOutputInfo3511) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLockedSubAssemblyOutputInfo3511) SetBtType(v string) {
	o.BtType = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BTLockedSubAssemblyOutputInfo3511) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BTLockedSubAssemblyOutputInfo3511) SetLocked(v bool) {
	o.Locked = &v
}

// GetRigid returns the Rigid field value if set, zero value otherwise.
func (o *BTLockedSubAssemblyOutputInfo3511) GetRigid() bool {
	if o == nil || o.Rigid == nil {
		var ret bool
		return ret
	}
	return *o.Rigid
}

// GetRigidOk returns a tuple with the Rigid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) GetRigidOk() (*bool, bool) {
	if o == nil || o.Rigid == nil {
		return nil, false
	}
	return o.Rigid, true
}

// HasRigid returns a boolean if a field has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) HasRigid() bool {
	if o != nil && o.Rigid != nil {
		return true
	}

	return false
}

// SetRigid gets a reference to the given bool and assigns it to the Rigid field.
func (o *BTLockedSubAssemblyOutputInfo3511) SetRigid(v bool) {
	o.Rigid = &v
}

// GetSyncedOutputMVID returns the SyncedOutputMVID field value if set, zero value otherwise.
func (o *BTLockedSubAssemblyOutputInfo3511) GetSyncedOutputMVID() BTMicroversionId366 {
	if o == nil || o.SyncedOutputMVID == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.SyncedOutputMVID
}

// GetSyncedOutputMVIDOk returns a tuple with the SyncedOutputMVID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) GetSyncedOutputMVIDOk() (*BTMicroversionId366, bool) {
	if o == nil || o.SyncedOutputMVID == nil {
		return nil, false
	}
	return o.SyncedOutputMVID, true
}

// HasSyncedOutputMVID returns a boolean if a field has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) HasSyncedOutputMVID() bool {
	if o != nil && o.SyncedOutputMVID != nil {
		return true
	}

	return false
}

// SetSyncedOutputMVID gets a reference to the given BTMicroversionId366 and assigns it to the SyncedOutputMVID field.
func (o *BTLockedSubAssemblyOutputInfo3511) SetSyncedOutputMVID(v BTMicroversionId366) {
	o.SyncedOutputMVID = &v
}

// GetLockInfo returns the LockInfo field value if set, zero value otherwise.
func (o *BTLockedSubAssemblyOutputInfo3511) GetLockInfo() BTLockedSubAssembly4590 {
	if o == nil || o.LockInfo == nil {
		var ret BTLockedSubAssembly4590
		return ret
	}
	return *o.LockInfo
}

// GetLockInfoOk returns a tuple with the LockInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) GetLockInfoOk() (*BTLockedSubAssembly4590, bool) {
	if o == nil || o.LockInfo == nil {
		return nil, false
	}
	return o.LockInfo, true
}

// HasLockInfo returns a boolean if a field has been set.
func (o *BTLockedSubAssemblyOutputInfo3511) HasLockInfo() bool {
	if o != nil && o.LockInfo != nil {
		return true
	}

	return false
}

// SetLockInfo gets a reference to the given BTLockedSubAssembly4590 and assigns it to the LockInfo field.
func (o *BTLockedSubAssemblyOutputInfo3511) SetLockInfo(v BTLockedSubAssembly4590) {
	o.LockInfo = &v
}

func (o BTLockedSubAssemblyOutputInfo3511) MarshalJSON() ([]byte, error) {
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
	if o.LockInfo != nil {
		toSerialize["lockInfo"] = o.LockInfo
	}
	return json.Marshal(toSerialize)
}

type NullableBTLockedSubAssemblyOutputInfo3511 struct {
	value *BTLockedSubAssemblyOutputInfo3511
	isSet bool
}

func (v NullableBTLockedSubAssemblyOutputInfo3511) Get() *BTLockedSubAssemblyOutputInfo3511 {
	return v.value
}

func (v *NullableBTLockedSubAssemblyOutputInfo3511) Set(val *BTLockedSubAssemblyOutputInfo3511) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLockedSubAssemblyOutputInfo3511) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLockedSubAssemblyOutputInfo3511) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLockedSubAssemblyOutputInfo3511(val *BTLockedSubAssemblyOutputInfo3511) *NullableBTLockedSubAssemblyOutputInfo3511 {
	return &NullableBTLockedSubAssemblyOutputInfo3511{value: val, isSet: true}
}

func (v NullableBTLockedSubAssemblyOutputInfo3511) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLockedSubAssemblyOutputInfo3511) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}