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

// BTTableCellPropertyParameter2983 struct for BTTableCellPropertyParameter2983
type BTTableCellPropertyParameter2983 struct {
	BTTableCellParameter2399
	BtType                              *string                        `json:"btType,omitempty"`
	IsEverVisible                       *bool                          `json:"isEverVisible,omitempty"`
	IsReadOnly                          *bool                          `json:"isReadOnly,omitempty"`
	Error                               *string                        `json:"error,omitempty"`
	OverrideSpec                        *BTParameterSpec6              `json:"overrideSpec,omitempty"`
	Parameter                           *BTMParameter1                 `json:"parameter,omitempty"`
	AggregationSkippedFilteredOutValues *bool                          `json:"aggregationSkippedFilteredOutValues,omitempty"`
	IsUnchanged                         *bool                          `json:"isUnchanged,omitempty"`
	OverrideStatusType                  *GBTMetadataOverrideStatusType `json:"overrideStatusType,omitempty"`
	PropertySourceType                  *GBTMetadataSourceType         `json:"propertySourceType,omitempty"`
}

// NewBTTableCellPropertyParameter2983 instantiates a new BTTableCellPropertyParameter2983 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableCellPropertyParameter2983() *BTTableCellPropertyParameter2983 {
	this := BTTableCellPropertyParameter2983{}
	return &this
}

// NewBTTableCellPropertyParameter2983WithDefaults instantiates a new BTTableCellPropertyParameter2983 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableCellPropertyParameter2983WithDefaults() *BTTableCellPropertyParameter2983 {
	this := BTTableCellPropertyParameter2983{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableCellPropertyParameter2983) SetBtType(v string) {
	o.BtType = &v
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetIsEverVisible() bool {
	if o == nil || o.IsEverVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsEverVisible
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetIsEverVisibleOk() (*bool, bool) {
	if o == nil || o.IsEverVisible == nil {
		return nil, false
	}
	return o.IsEverVisible, true
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasIsEverVisible() bool {
	if o != nil && o.IsEverVisible != nil {
		return true
	}

	return false
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *BTTableCellPropertyParameter2983) SetIsEverVisible(v bool) {
	o.IsEverVisible = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTTableCellPropertyParameter2983) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTTableCellPropertyParameter2983) SetError(v string) {
	o.Error = &v
}

// GetOverrideSpec returns the OverrideSpec field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetOverrideSpec() BTParameterSpec6 {
	if o == nil || o.OverrideSpec == nil {
		var ret BTParameterSpec6
		return ret
	}
	return *o.OverrideSpec
}

// GetOverrideSpecOk returns a tuple with the OverrideSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetOverrideSpecOk() (*BTParameterSpec6, bool) {
	if o == nil || o.OverrideSpec == nil {
		return nil, false
	}
	return o.OverrideSpec, true
}

// HasOverrideSpec returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasOverrideSpec() bool {
	if o != nil && o.OverrideSpec != nil {
		return true
	}

	return false
}

// SetOverrideSpec gets a reference to the given BTParameterSpec6 and assigns it to the OverrideSpec field.
func (o *BTTableCellPropertyParameter2983) SetOverrideSpec(v BTParameterSpec6) {
	o.OverrideSpec = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetParameter() BTMParameter1 {
	if o == nil || o.Parameter == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetParameterOk() (*BTMParameter1, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given BTMParameter1 and assigns it to the Parameter field.
func (o *BTTableCellPropertyParameter2983) SetParameter(v BTMParameter1) {
	o.Parameter = &v
}

// GetAggregationSkippedFilteredOutValues returns the AggregationSkippedFilteredOutValues field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetAggregationSkippedFilteredOutValues() bool {
	if o == nil || o.AggregationSkippedFilteredOutValues == nil {
		var ret bool
		return ret
	}
	return *o.AggregationSkippedFilteredOutValues
}

// GetAggregationSkippedFilteredOutValuesOk returns a tuple with the AggregationSkippedFilteredOutValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetAggregationSkippedFilteredOutValuesOk() (*bool, bool) {
	if o == nil || o.AggregationSkippedFilteredOutValues == nil {
		return nil, false
	}
	return o.AggregationSkippedFilteredOutValues, true
}

// HasAggregationSkippedFilteredOutValues returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasAggregationSkippedFilteredOutValues() bool {
	if o != nil && o.AggregationSkippedFilteredOutValues != nil {
		return true
	}

	return false
}

// SetAggregationSkippedFilteredOutValues gets a reference to the given bool and assigns it to the AggregationSkippedFilteredOutValues field.
func (o *BTTableCellPropertyParameter2983) SetAggregationSkippedFilteredOutValues(v bool) {
	o.AggregationSkippedFilteredOutValues = &v
}

// GetIsUnchanged returns the IsUnchanged field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetIsUnchanged() bool {
	if o == nil || o.IsUnchanged == nil {
		var ret bool
		return ret
	}
	return *o.IsUnchanged
}

// GetIsUnchangedOk returns a tuple with the IsUnchanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetIsUnchangedOk() (*bool, bool) {
	if o == nil || o.IsUnchanged == nil {
		return nil, false
	}
	return o.IsUnchanged, true
}

// HasIsUnchanged returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasIsUnchanged() bool {
	if o != nil && o.IsUnchanged != nil {
		return true
	}

	return false
}

// SetIsUnchanged gets a reference to the given bool and assigns it to the IsUnchanged field.
func (o *BTTableCellPropertyParameter2983) SetIsUnchanged(v bool) {
	o.IsUnchanged = &v
}

// GetOverrideStatusType returns the OverrideStatusType field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetOverrideStatusType() GBTMetadataOverrideStatusType {
	if o == nil || o.OverrideStatusType == nil {
		var ret GBTMetadataOverrideStatusType
		return ret
	}
	return *o.OverrideStatusType
}

// GetOverrideStatusTypeOk returns a tuple with the OverrideStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetOverrideStatusTypeOk() (*GBTMetadataOverrideStatusType, bool) {
	if o == nil || o.OverrideStatusType == nil {
		return nil, false
	}
	return o.OverrideStatusType, true
}

// HasOverrideStatusType returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasOverrideStatusType() bool {
	if o != nil && o.OverrideStatusType != nil {
		return true
	}

	return false
}

// SetOverrideStatusType gets a reference to the given GBTMetadataOverrideStatusType and assigns it to the OverrideStatusType field.
func (o *BTTableCellPropertyParameter2983) SetOverrideStatusType(v GBTMetadataOverrideStatusType) {
	o.OverrideStatusType = &v
}

// GetPropertySourceType returns the PropertySourceType field value if set, zero value otherwise.
func (o *BTTableCellPropertyParameter2983) GetPropertySourceType() GBTMetadataSourceType {
	if o == nil || o.PropertySourceType == nil {
		var ret GBTMetadataSourceType
		return ret
	}
	return *o.PropertySourceType
}

// GetPropertySourceTypeOk returns a tuple with the PropertySourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellPropertyParameter2983) GetPropertySourceTypeOk() (*GBTMetadataSourceType, bool) {
	if o == nil || o.PropertySourceType == nil {
		return nil, false
	}
	return o.PropertySourceType, true
}

// HasPropertySourceType returns a boolean if a field has been set.
func (o *BTTableCellPropertyParameter2983) HasPropertySourceType() bool {
	if o != nil && o.PropertySourceType != nil {
		return true
	}

	return false
}

// SetPropertySourceType gets a reference to the given GBTMetadataSourceType and assigns it to the PropertySourceType field.
func (o *BTTableCellPropertyParameter2983) SetPropertySourceType(v GBTMetadataSourceType) {
	o.PropertySourceType = &v
}

func (o BTTableCellPropertyParameter2983) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTableCellParameter2399, errBTTableCellParameter2399 := json.Marshal(o.BTTableCellParameter2399)
	if errBTTableCellParameter2399 != nil {
		return []byte{}, errBTTableCellParameter2399
	}
	errBTTableCellParameter2399 = json.Unmarshal([]byte(serializedBTTableCellParameter2399), &toSerialize)
	if errBTTableCellParameter2399 != nil {
		return []byte{}, errBTTableCellParameter2399
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsEverVisible != nil {
		toSerialize["isEverVisible"] = o.IsEverVisible
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OverrideSpec != nil {
		toSerialize["overrideSpec"] = o.OverrideSpec
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	if o.AggregationSkippedFilteredOutValues != nil {
		toSerialize["aggregationSkippedFilteredOutValues"] = o.AggregationSkippedFilteredOutValues
	}
	if o.IsUnchanged != nil {
		toSerialize["isUnchanged"] = o.IsUnchanged
	}
	if o.OverrideStatusType != nil {
		toSerialize["overrideStatusType"] = o.OverrideStatusType
	}
	if o.PropertySourceType != nil {
		toSerialize["propertySourceType"] = o.PropertySourceType
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableCellPropertyParameter2983 struct {
	value *BTTableCellPropertyParameter2983
	isSet bool
}

func (v NullableBTTableCellPropertyParameter2983) Get() *BTTableCellPropertyParameter2983 {
	return v.value
}

func (v *NullableBTTableCellPropertyParameter2983) Set(val *BTTableCellPropertyParameter2983) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableCellPropertyParameter2983) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableCellPropertyParameter2983) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableCellPropertyParameter2983(val *BTTableCellPropertyParameter2983) *NullableBTTableCellPropertyParameter2983 {
	return &NullableBTTableCellPropertyParameter2983{value: val, isSet: true}
}

func (v NullableBTTableCellPropertyParameter2983) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableCellPropertyParameter2983) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
