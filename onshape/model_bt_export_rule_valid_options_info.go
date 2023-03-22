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

// BTExportRuleValidOptionsInfo struct for BTExportRuleValidOptionsInfo
type BTExportRuleValidOptionsInfo struct {
	ConventionPlaceholder     *string                             `json:"conventionPlaceholder,omitempty"`
	HardcodedProperties       []BTExportRuleHardcodedPropertyInfo `json:"hardcodedProperties,omitempty"`
	PropertyContextDisplayMap *map[string]string                  `json:"propertyContextDisplayMap,omitempty"`
	ValidObjectTypes          []int32                             `json:"validObjectTypes,omitempty"`
}

// NewBTExportRuleValidOptionsInfo instantiates a new BTExportRuleValidOptionsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportRuleValidOptionsInfo() *BTExportRuleValidOptionsInfo {
	this := BTExportRuleValidOptionsInfo{}
	return &this
}

// NewBTExportRuleValidOptionsInfoWithDefaults instantiates a new BTExportRuleValidOptionsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportRuleValidOptionsInfoWithDefaults() *BTExportRuleValidOptionsInfo {
	this := BTExportRuleValidOptionsInfo{}
	return &this
}

// GetConventionPlaceholder returns the ConventionPlaceholder field value if set, zero value otherwise.
func (o *BTExportRuleValidOptionsInfo) GetConventionPlaceholder() string {
	if o == nil || o.ConventionPlaceholder == nil {
		var ret string
		return ret
	}
	return *o.ConventionPlaceholder
}

// GetConventionPlaceholderOk returns a tuple with the ConventionPlaceholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleValidOptionsInfo) GetConventionPlaceholderOk() (*string, bool) {
	if o == nil || o.ConventionPlaceholder == nil {
		return nil, false
	}
	return o.ConventionPlaceholder, true
}

// HasConventionPlaceholder returns a boolean if a field has been set.
func (o *BTExportRuleValidOptionsInfo) HasConventionPlaceholder() bool {
	if o != nil && o.ConventionPlaceholder != nil {
		return true
	}

	return false
}

// SetConventionPlaceholder gets a reference to the given string and assigns it to the ConventionPlaceholder field.
func (o *BTExportRuleValidOptionsInfo) SetConventionPlaceholder(v string) {
	o.ConventionPlaceholder = &v
}

// GetHardcodedProperties returns the HardcodedProperties field value if set, zero value otherwise.
func (o *BTExportRuleValidOptionsInfo) GetHardcodedProperties() []BTExportRuleHardcodedPropertyInfo {
	if o == nil || o.HardcodedProperties == nil {
		var ret []BTExportRuleHardcodedPropertyInfo
		return ret
	}
	return o.HardcodedProperties
}

// GetHardcodedPropertiesOk returns a tuple with the HardcodedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleValidOptionsInfo) GetHardcodedPropertiesOk() ([]BTExportRuleHardcodedPropertyInfo, bool) {
	if o == nil || o.HardcodedProperties == nil {
		return nil, false
	}
	return o.HardcodedProperties, true
}

// HasHardcodedProperties returns a boolean if a field has been set.
func (o *BTExportRuleValidOptionsInfo) HasHardcodedProperties() bool {
	if o != nil && o.HardcodedProperties != nil {
		return true
	}

	return false
}

// SetHardcodedProperties gets a reference to the given []BTExportRuleHardcodedPropertyInfo and assigns it to the HardcodedProperties field.
func (o *BTExportRuleValidOptionsInfo) SetHardcodedProperties(v []BTExportRuleHardcodedPropertyInfo) {
	o.HardcodedProperties = v
}

// GetPropertyContextDisplayMap returns the PropertyContextDisplayMap field value if set, zero value otherwise.
func (o *BTExportRuleValidOptionsInfo) GetPropertyContextDisplayMap() map[string]string {
	if o == nil || o.PropertyContextDisplayMap == nil {
		var ret map[string]string
		return ret
	}
	return *o.PropertyContextDisplayMap
}

// GetPropertyContextDisplayMapOk returns a tuple with the PropertyContextDisplayMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleValidOptionsInfo) GetPropertyContextDisplayMapOk() (*map[string]string, bool) {
	if o == nil || o.PropertyContextDisplayMap == nil {
		return nil, false
	}
	return o.PropertyContextDisplayMap, true
}

// HasPropertyContextDisplayMap returns a boolean if a field has been set.
func (o *BTExportRuleValidOptionsInfo) HasPropertyContextDisplayMap() bool {
	if o != nil && o.PropertyContextDisplayMap != nil {
		return true
	}

	return false
}

// SetPropertyContextDisplayMap gets a reference to the given map[string]string and assigns it to the PropertyContextDisplayMap field.
func (o *BTExportRuleValidOptionsInfo) SetPropertyContextDisplayMap(v map[string]string) {
	o.PropertyContextDisplayMap = &v
}

// GetValidObjectTypes returns the ValidObjectTypes field value if set, zero value otherwise.
func (o *BTExportRuleValidOptionsInfo) GetValidObjectTypes() []int32 {
	if o == nil || o.ValidObjectTypes == nil {
		var ret []int32
		return ret
	}
	return o.ValidObjectTypes
}

// GetValidObjectTypesOk returns a tuple with the ValidObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportRuleValidOptionsInfo) GetValidObjectTypesOk() ([]int32, bool) {
	if o == nil || o.ValidObjectTypes == nil {
		return nil, false
	}
	return o.ValidObjectTypes, true
}

// HasValidObjectTypes returns a boolean if a field has been set.
func (o *BTExportRuleValidOptionsInfo) HasValidObjectTypes() bool {
	if o != nil && o.ValidObjectTypes != nil {
		return true
	}

	return false
}

// SetValidObjectTypes gets a reference to the given []int32 and assigns it to the ValidObjectTypes field.
func (o *BTExportRuleValidOptionsInfo) SetValidObjectTypes(v []int32) {
	o.ValidObjectTypes = v
}

func (o BTExportRuleValidOptionsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConventionPlaceholder != nil {
		toSerialize["conventionPlaceholder"] = o.ConventionPlaceholder
	}
	if o.HardcodedProperties != nil {
		toSerialize["hardcodedProperties"] = o.HardcodedProperties
	}
	if o.PropertyContextDisplayMap != nil {
		toSerialize["propertyContextDisplayMap"] = o.PropertyContextDisplayMap
	}
	if o.ValidObjectTypes != nil {
		toSerialize["validObjectTypes"] = o.ValidObjectTypes
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportRuleValidOptionsInfo struct {
	value *BTExportRuleValidOptionsInfo
	isSet bool
}

func (v NullableBTExportRuleValidOptionsInfo) Get() *BTExportRuleValidOptionsInfo {
	return v.value
}

func (v *NullableBTExportRuleValidOptionsInfo) Set(val *BTExportRuleValidOptionsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportRuleValidOptionsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportRuleValidOptionsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportRuleValidOptionsInfo(val *BTExportRuleValidOptionsInfo) *NullableBTExportRuleValidOptionsInfo {
	return &NullableBTExportRuleValidOptionsInfo{value: val, isSet: true}
}

func (v NullableBTExportRuleValidOptionsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportRuleValidOptionsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
