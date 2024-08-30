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

// BTCreateTaskParams struct for BTCreateTaskParams
type BTCreateTaskParams struct {
	// Id of the company that owns the task.
	CompanyId string `json:"companyId"`
	// Description of the task.
	Description           *string `json:"description,omitempty"`
	DescriptionParamValue *string `json:"descriptionParamValue,omitempty"`
	// Id of a document to associate the task to.
	DocumentId *string `json:"documentId,omitempty"`
	// References to include in the task.
	ItemParams []BTTaskItemParams `json:"itemParams,omitempty"`
	// Name of the task.
	Name           *string `json:"name,omitempty"`
	NameParamValue *string `json:"nameParamValue,omitempty"`
	// Set Task metadata properties.
	PropertyValues []BTPropertyValueParam `json:"propertyValues,omitempty"`
}

// NewBTCreateTaskParams instantiates a new BTCreateTaskParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCreateTaskParams(companyId string) *BTCreateTaskParams {
	this := BTCreateTaskParams{}
	this.CompanyId = companyId
	return &this
}

// NewBTCreateTaskParamsWithDefaults instantiates a new BTCreateTaskParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCreateTaskParamsWithDefaults() *BTCreateTaskParams {
	this := BTCreateTaskParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *BTCreateTaskParams) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *BTCreateTaskParams) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTCreateTaskParams) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionParamValue returns the DescriptionParamValue field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetDescriptionParamValue() string {
	if o == nil || o.DescriptionParamValue == nil {
		var ret string
		return ret
	}
	return *o.DescriptionParamValue
}

// GetDescriptionParamValueOk returns a tuple with the DescriptionParamValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetDescriptionParamValueOk() (*string, bool) {
	if o == nil || o.DescriptionParamValue == nil {
		return nil, false
	}
	return o.DescriptionParamValue, true
}

// HasDescriptionParamValue returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasDescriptionParamValue() bool {
	if o != nil && o.DescriptionParamValue != nil {
		return true
	}

	return false
}

// SetDescriptionParamValue gets a reference to the given string and assigns it to the DescriptionParamValue field.
func (o *BTCreateTaskParams) SetDescriptionParamValue(v string) {
	o.DescriptionParamValue = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTCreateTaskParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetItemParams returns the ItemParams field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetItemParams() []BTTaskItemParams {
	if o == nil || o.ItemParams == nil {
		var ret []BTTaskItemParams
		return ret
	}
	return o.ItemParams
}

// GetItemParamsOk returns a tuple with the ItemParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetItemParamsOk() ([]BTTaskItemParams, bool) {
	if o == nil || o.ItemParams == nil {
		return nil, false
	}
	return o.ItemParams, true
}

// HasItemParams returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasItemParams() bool {
	if o != nil && o.ItemParams != nil {
		return true
	}

	return false
}

// SetItemParams gets a reference to the given []BTTaskItemParams and assigns it to the ItemParams field.
func (o *BTCreateTaskParams) SetItemParams(v []BTTaskItemParams) {
	o.ItemParams = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCreateTaskParams) SetName(v string) {
	o.Name = &v
}

// GetNameParamValue returns the NameParamValue field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetNameParamValue() string {
	if o == nil || o.NameParamValue == nil {
		var ret string
		return ret
	}
	return *o.NameParamValue
}

// GetNameParamValueOk returns a tuple with the NameParamValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetNameParamValueOk() (*string, bool) {
	if o == nil || o.NameParamValue == nil {
		return nil, false
	}
	return o.NameParamValue, true
}

// HasNameParamValue returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasNameParamValue() bool {
	if o != nil && o.NameParamValue != nil {
		return true
	}

	return false
}

// SetNameParamValue gets a reference to the given string and assigns it to the NameParamValue field.
func (o *BTCreateTaskParams) SetNameParamValue(v string) {
	o.NameParamValue = &v
}

// GetPropertyValues returns the PropertyValues field value if set, zero value otherwise.
func (o *BTCreateTaskParams) GetPropertyValues() []BTPropertyValueParam {
	if o == nil || o.PropertyValues == nil {
		var ret []BTPropertyValueParam
		return ret
	}
	return o.PropertyValues
}

// GetPropertyValuesOk returns a tuple with the PropertyValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCreateTaskParams) GetPropertyValuesOk() ([]BTPropertyValueParam, bool) {
	if o == nil || o.PropertyValues == nil {
		return nil, false
	}
	return o.PropertyValues, true
}

// HasPropertyValues returns a boolean if a field has been set.
func (o *BTCreateTaskParams) HasPropertyValues() bool {
	if o != nil && o.PropertyValues != nil {
		return true
	}

	return false
}

// SetPropertyValues gets a reference to the given []BTPropertyValueParam and assigns it to the PropertyValues field.
func (o *BTCreateTaskParams) SetPropertyValues(v []BTPropertyValueParam) {
	o.PropertyValues = v
}

func (o BTCreateTaskParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DescriptionParamValue != nil {
		toSerialize["descriptionParamValue"] = o.DescriptionParamValue
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ItemParams != nil {
		toSerialize["itemParams"] = o.ItemParams
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameParamValue != nil {
		toSerialize["nameParamValue"] = o.NameParamValue
	}
	if o.PropertyValues != nil {
		toSerialize["propertyValues"] = o.PropertyValues
	}
	return json.Marshal(toSerialize)
}

type NullableBTCreateTaskParams struct {
	value *BTCreateTaskParams
	isSet bool
}

func (v NullableBTCreateTaskParams) Get() *BTCreateTaskParams {
	return v.value
}

func (v *NullableBTCreateTaskParams) Set(val *BTCreateTaskParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCreateTaskParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCreateTaskParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCreateTaskParams(val *BTCreateTaskParams) *NullableBTCreateTaskParams {
	return &NullableBTCreateTaskParams{value: val, isSet: true}
}

func (v NullableBTCreateTaskParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCreateTaskParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
