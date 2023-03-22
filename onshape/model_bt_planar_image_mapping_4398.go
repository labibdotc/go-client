/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13258-6b30d30337fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPlanarImageMapping4398 struct for BTPlanarImageMapping4398
type BTPlanarImageMapping4398 struct {
	DeterministicIds []string               `json:"deterministicIds,omitempty"`
	UvTransform      *BTMatrix3x3340        `json:"uvTransform,omitempty"`
	BtType           *string                `json:"btType,omitempty"`
	PlaneSystem      *BTCoordinateSystem387 `json:"planeSystem,omitempty"`
}

// NewBTPlanarImageMapping4398 instantiates a new BTPlanarImageMapping4398 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPlanarImageMapping4398() *BTPlanarImageMapping4398 {
	this := BTPlanarImageMapping4398{}
	return &this
}

// NewBTPlanarImageMapping4398WithDefaults instantiates a new BTPlanarImageMapping4398 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPlanarImageMapping4398WithDefaults() *BTPlanarImageMapping4398 {
	this := BTPlanarImageMapping4398{}
	return &this
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTPlanarImageMapping4398) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanarImageMapping4398) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTPlanarImageMapping4398) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTPlanarImageMapping4398) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetUvTransform returns the UvTransform field value if set, zero value otherwise.
func (o *BTPlanarImageMapping4398) GetUvTransform() BTMatrix3x3340 {
	if o == nil || o.UvTransform == nil {
		var ret BTMatrix3x3340
		return ret
	}
	return *o.UvTransform
}

// GetUvTransformOk returns a tuple with the UvTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanarImageMapping4398) GetUvTransformOk() (*BTMatrix3x3340, bool) {
	if o == nil || o.UvTransform == nil {
		return nil, false
	}
	return o.UvTransform, true
}

// HasUvTransform returns a boolean if a field has been set.
func (o *BTPlanarImageMapping4398) HasUvTransform() bool {
	if o != nil && o.UvTransform != nil {
		return true
	}

	return false
}

// SetUvTransform gets a reference to the given BTMatrix3x3340 and assigns it to the UvTransform field.
func (o *BTPlanarImageMapping4398) SetUvTransform(v BTMatrix3x3340) {
	o.UvTransform = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPlanarImageMapping4398) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanarImageMapping4398) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPlanarImageMapping4398) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPlanarImageMapping4398) SetBtType(v string) {
	o.BtType = &v
}

// GetPlaneSystem returns the PlaneSystem field value if set, zero value otherwise.
func (o *BTPlanarImageMapping4398) GetPlaneSystem() BTCoordinateSystem387 {
	if o == nil || o.PlaneSystem == nil {
		var ret BTCoordinateSystem387
		return ret
	}
	return *o.PlaneSystem
}

// GetPlaneSystemOk returns a tuple with the PlaneSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanarImageMapping4398) GetPlaneSystemOk() (*BTCoordinateSystem387, bool) {
	if o == nil || o.PlaneSystem == nil {
		return nil, false
	}
	return o.PlaneSystem, true
}

// HasPlaneSystem returns a boolean if a field has been set.
func (o *BTPlanarImageMapping4398) HasPlaneSystem() bool {
	if o != nil && o.PlaneSystem != nil {
		return true
	}

	return false
}

// SetPlaneSystem gets a reference to the given BTCoordinateSystem387 and assigns it to the PlaneSystem field.
func (o *BTPlanarImageMapping4398) SetPlaneSystem(v BTCoordinateSystem387) {
	o.PlaneSystem = &v
}

func (o BTPlanarImageMapping4398) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.UvTransform != nil {
		toSerialize["uvTransform"] = o.UvTransform
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.PlaneSystem != nil {
		toSerialize["planeSystem"] = o.PlaneSystem
	}
	return json.Marshal(toSerialize)
}

type NullableBTPlanarImageMapping4398 struct {
	value *BTPlanarImageMapping4398
	isSet bool
}

func (v NullableBTPlanarImageMapping4398) Get() *BTPlanarImageMapping4398 {
	return v.value
}

func (v *NullableBTPlanarImageMapping4398) Set(val *BTPlanarImageMapping4398) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPlanarImageMapping4398) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPlanarImageMapping4398) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPlanarImageMapping4398(val *BTPlanarImageMapping4398) *NullableBTPlanarImageMapping4398 {
	return &NullableBTPlanarImageMapping4398{value: val, isSet: true}
}

func (v NullableBTPlanarImageMapping4398) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPlanarImageMapping4398) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
