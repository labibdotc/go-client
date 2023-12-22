/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28023-41481dfcfdcb
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTJEdit3734 - An edit that will be applied to the application element's json data.
type BTJEdit3734 struct {
	implBTJEdit3734 interface{}
}

// BTJEditMove3245AsBTJEdit3734 is a convenience function that returns BTJEditMove3245 wrapped in BTJEdit3734
func (o *BTJEditMove3245) AsBTJEdit3734() *BTJEdit3734 {
	return &BTJEdit3734{o}
}

// BTJEditList2707AsBTJEdit3734 is a convenience function that returns BTJEditList2707 wrapped in BTJEdit3734
func (o *BTJEditList2707) AsBTJEdit3734() *BTJEdit3734 {
	return &BTJEdit3734{o}
}

// BTJEditDelete1992AsBTJEdit3734 is a convenience function that returns BTJEditDelete1992 wrapped in BTJEdit3734
func (o *BTJEditDelete1992) AsBTJEdit3734() *BTJEdit3734 {
	return &BTJEdit3734{o}
}

// BTJEditChange2636AsBTJEdit3734 is a convenience function that returns BTJEditChange2636 wrapped in BTJEdit3734
func (o *BTJEditChange2636) AsBTJEdit3734() *BTJEdit3734 {
	return &BTJEdit3734{o}
}

// BTJEditInsert2523AsBTJEdit3734 is a convenience function that returns BTJEditInsert2523 wrapped in BTJEdit3734
func (o *BTJEditInsert2523) AsBTJEdit3734() *BTJEdit3734 {
	return &BTJEdit3734{o}
}

// NewBTJEdit3734 instantiates a new BTJEdit3734 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJEdit3734() *BTJEdit3734 {
	this := BTJEdit3734{Newbase_BTJEdit3734()}
	return &this
}

// NewBTJEdit3734WithDefaults instantiates a new BTJEdit3734 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJEdit3734WithDefaults() *BTJEdit3734 {
	this := BTJEdit3734{Newbase_BTJEdit3734WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJEdit3734) GetBtType() string {
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
func (o *BTJEdit3734) GetBtTypeOk() (*string, bool) {
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
func (o *BTJEdit3734) HasBtType() bool {
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
func (o *BTJEdit3734) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTJEdit3734) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTJEditChange-2636'
	if jsonDict["btType"] == "BTJEditChange-2636" {
		// try to unmarshal JSON data into BTJEditChange2636
		var qr *BTJEditChange2636
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJEdit3734 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJEdit3734 = nil
			return fmt.Errorf("Failed to unmarshal BTJEdit3734 as BTJEditChange2636: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTJEditDelete-1992'
	if jsonDict["btType"] == "BTJEditDelete-1992" {
		// try to unmarshal JSON data into BTJEditDelete1992
		var qr *BTJEditDelete1992
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJEdit3734 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJEdit3734 = nil
			return fmt.Errorf("Failed to unmarshal BTJEdit3734 as BTJEditDelete1992: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTJEditInsert-2523'
	if jsonDict["btType"] == "BTJEditInsert-2523" {
		// try to unmarshal JSON data into BTJEditInsert2523
		var qr *BTJEditInsert2523
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJEdit3734 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJEdit3734 = nil
			return fmt.Errorf("Failed to unmarshal BTJEdit3734 as BTJEditInsert2523: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTJEditList-2707'
	if jsonDict["btType"] == "BTJEditList-2707" {
		// try to unmarshal JSON data into BTJEditList2707
		var qr *BTJEditList2707
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJEdit3734 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJEdit3734 = nil
			return fmt.Errorf("Failed to unmarshal BTJEdit3734 as BTJEditList2707: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTJEditMove-3245'
	if jsonDict["btType"] == "BTJEditMove-3245" {
		// try to unmarshal JSON data into BTJEditMove3245
		var qr *BTJEditMove3245
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTJEdit3734 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTJEdit3734 = nil
			return fmt.Errorf("Failed to unmarshal BTJEdit3734 as BTJEditMove3245: %s", err.Error())
		}
	}

	var qtx *base_BTJEdit3734
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTJEdit3734 = qtx
		return nil // data stored in dst.base_BTJEdit3734, return on the first match
	} else {
		dst.implBTJEdit3734 = nil
		return fmt.Errorf("Failed to unmarshal BTJEdit3734 as base_BTJEdit3734: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTJEdit3734) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTJEdit3734) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTJEdit3734
}

type NullableBTJEdit3734 struct {
	value *BTJEdit3734
	isSet bool
}

func (v NullableBTJEdit3734) Get() *BTJEdit3734 {
	return v.value
}

func (v *NullableBTJEdit3734) Set(val *BTJEdit3734) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJEdit3734) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJEdit3734) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJEdit3734(val *BTJEdit3734) *NullableBTJEdit3734 {
	return &NullableBTJEdit3734{value: val, isSet: true}
}

func (v NullableBTJEdit3734) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJEdit3734) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTJEdit3734 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTJEdit3734 instantiates a new base_BTJEdit3734 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTJEdit3734() *base_BTJEdit3734 {
	this := base_BTJEdit3734{}
	return &this
}

// Newbase_BTJEdit3734WithDefaults instantiates a new base_BTJEdit3734 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTJEdit3734WithDefaults() *base_BTJEdit3734 {
	this := base_BTJEdit3734{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTJEdit3734) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTJEdit3734) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTJEdit3734) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTJEdit3734) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTJEdit3734) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
