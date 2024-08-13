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

// BTLockedAtNamedPosition3201 struct for BTLockedAtNamedPosition3201
type BTLockedAtNamedPosition3201 struct {
	BtType                      *string                                   `json:"btType,omitempty"`
	LockType                    *GBTSubAssemblyLockType                   `json:"lockType,omitempty"`
	LockedSubAssemblyOutputInfo *BTRigidOrLockedSubAssemblyOutputInfo3860 `json:"lockedSubAssemblyOutputInfo,omitempty"`
	NamedPositionId             *string                                   `json:"namedPositionId,omitempty"`
}

// NewBTLockedAtNamedPosition3201 instantiates a new BTLockedAtNamedPosition3201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLockedAtNamedPosition3201() *BTLockedAtNamedPosition3201 {
	this := BTLockedAtNamedPosition3201{}
	return &this
}

// NewBTLockedAtNamedPosition3201WithDefaults instantiates a new BTLockedAtNamedPosition3201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLockedAtNamedPosition3201WithDefaults() *BTLockedAtNamedPosition3201 {
	this := BTLockedAtNamedPosition3201{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLockedAtNamedPosition3201) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtNamedPosition3201) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLockedAtNamedPosition3201) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLockedAtNamedPosition3201) SetBtType(v string) {
	o.BtType = &v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *BTLockedAtNamedPosition3201) GetLockType() GBTSubAssemblyLockType {
	if o == nil || o.LockType == nil {
		var ret GBTSubAssemblyLockType
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtNamedPosition3201) GetLockTypeOk() (*GBTSubAssemblyLockType, bool) {
	if o == nil || o.LockType == nil {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *BTLockedAtNamedPosition3201) HasLockType() bool {
	if o != nil && o.LockType != nil {
		return true
	}

	return false
}

// SetLockType gets a reference to the given GBTSubAssemblyLockType and assigns it to the LockType field.
func (o *BTLockedAtNamedPosition3201) SetLockType(v GBTSubAssemblyLockType) {
	o.LockType = &v
}

// GetLockedSubAssemblyOutputInfo returns the LockedSubAssemblyOutputInfo field value if set, zero value otherwise.
func (o *BTLockedAtNamedPosition3201) GetLockedSubAssemblyOutputInfo() BTRigidOrLockedSubAssemblyOutputInfo3860 {
	if o == nil || o.LockedSubAssemblyOutputInfo == nil {
		var ret BTRigidOrLockedSubAssemblyOutputInfo3860
		return ret
	}
	return *o.LockedSubAssemblyOutputInfo
}

// GetLockedSubAssemblyOutputInfoOk returns a tuple with the LockedSubAssemblyOutputInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtNamedPosition3201) GetLockedSubAssemblyOutputInfoOk() (*BTRigidOrLockedSubAssemblyOutputInfo3860, bool) {
	if o == nil || o.LockedSubAssemblyOutputInfo == nil {
		return nil, false
	}
	return o.LockedSubAssemblyOutputInfo, true
}

// HasLockedSubAssemblyOutputInfo returns a boolean if a field has been set.
func (o *BTLockedAtNamedPosition3201) HasLockedSubAssemblyOutputInfo() bool {
	if o != nil && o.LockedSubAssemblyOutputInfo != nil {
		return true
	}

	return false
}

// SetLockedSubAssemblyOutputInfo gets a reference to the given BTRigidOrLockedSubAssemblyOutputInfo3860 and assigns it to the LockedSubAssemblyOutputInfo field.
func (o *BTLockedAtNamedPosition3201) SetLockedSubAssemblyOutputInfo(v BTRigidOrLockedSubAssemblyOutputInfo3860) {
	o.LockedSubAssemblyOutputInfo = &v
}

// GetNamedPositionId returns the NamedPositionId field value if set, zero value otherwise.
func (o *BTLockedAtNamedPosition3201) GetNamedPositionId() string {
	if o == nil || o.NamedPositionId == nil {
		var ret string
		return ret
	}
	return *o.NamedPositionId
}

// GetNamedPositionIdOk returns a tuple with the NamedPositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtNamedPosition3201) GetNamedPositionIdOk() (*string, bool) {
	if o == nil || o.NamedPositionId == nil {
		return nil, false
	}
	return o.NamedPositionId, true
}

// HasNamedPositionId returns a boolean if a field has been set.
func (o *BTLockedAtNamedPosition3201) HasNamedPositionId() bool {
	if o != nil && o.NamedPositionId != nil {
		return true
	}

	return false
}

// SetNamedPositionId gets a reference to the given string and assigns it to the NamedPositionId field.
func (o *BTLockedAtNamedPosition3201) SetNamedPositionId(v string) {
	o.NamedPositionId = &v
}

func (o BTLockedAtNamedPosition3201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.LockType != nil {
		toSerialize["lockType"] = o.LockType
	}
	if o.LockedSubAssemblyOutputInfo != nil {
		toSerialize["lockedSubAssemblyOutputInfo"] = o.LockedSubAssemblyOutputInfo
	}
	if o.NamedPositionId != nil {
		toSerialize["namedPositionId"] = o.NamedPositionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTLockedAtNamedPosition3201 struct {
	value *BTLockedAtNamedPosition3201
	isSet bool
}

func (v NullableBTLockedAtNamedPosition3201) Get() *BTLockedAtNamedPosition3201 {
	return v.value
}

func (v *NullableBTLockedAtNamedPosition3201) Set(val *BTLockedAtNamedPosition3201) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLockedAtNamedPosition3201) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLockedAtNamedPosition3201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLockedAtNamedPosition3201(val *BTLockedAtNamedPosition3201) *NullableBTLockedAtNamedPosition3201 {
	return &NullableBTLockedAtNamedPosition3201{value: val, isSet: true}
}

func (v NullableBTLockedAtNamedPosition3201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLockedAtNamedPosition3201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
