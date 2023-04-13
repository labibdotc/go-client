/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14462-13ace71ec1df
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPropertiesTableTemplateParams struct for BTPropertiesTableTemplateParams
type BTPropertiesTableTemplateParams struct {
	CompanyId       *string  `json:"companyId,omitempty"`
	IsAllCaps       *bool    `json:"isAllCaps,omitempty"`
	Name            *string  `json:"name,omitempty"`
	PropertyColumns []string `json:"propertyColumns,omitempty"`
	TableType       *string  `json:"tableType,omitempty"`
}

// NewBTPropertiesTableTemplateParams instantiates a new BTPropertiesTableTemplateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPropertiesTableTemplateParams() *BTPropertiesTableTemplateParams {
	this := BTPropertiesTableTemplateParams{}
	return &this
}

// NewBTPropertiesTableTemplateParamsWithDefaults instantiates a new BTPropertiesTableTemplateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPropertiesTableTemplateParamsWithDefaults() *BTPropertiesTableTemplateParams {
	this := BTPropertiesTableTemplateParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateParams) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateParams) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTPropertiesTableTemplateParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetIsAllCaps returns the IsAllCaps field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateParams) GetIsAllCaps() bool {
	if o == nil || o.IsAllCaps == nil {
		var ret bool
		return ret
	}
	return *o.IsAllCaps
}

// GetIsAllCapsOk returns a tuple with the IsAllCaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateParams) GetIsAllCapsOk() (*bool, bool) {
	if o == nil || o.IsAllCaps == nil {
		return nil, false
	}
	return o.IsAllCaps, true
}

// HasIsAllCaps returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateParams) HasIsAllCaps() bool {
	if o != nil && o.IsAllCaps != nil {
		return true
	}

	return false
}

// SetIsAllCaps gets a reference to the given bool and assigns it to the IsAllCaps field.
func (o *BTPropertiesTableTemplateParams) SetIsAllCaps(v bool) {
	o.IsAllCaps = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPropertiesTableTemplateParams) SetName(v string) {
	o.Name = &v
}

// GetPropertyColumns returns the PropertyColumns field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateParams) GetPropertyColumns() []string {
	if o == nil || o.PropertyColumns == nil {
		var ret []string
		return ret
	}
	return o.PropertyColumns
}

// GetPropertyColumnsOk returns a tuple with the PropertyColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateParams) GetPropertyColumnsOk() ([]string, bool) {
	if o == nil || o.PropertyColumns == nil {
		return nil, false
	}
	return o.PropertyColumns, true
}

// HasPropertyColumns returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateParams) HasPropertyColumns() bool {
	if o != nil && o.PropertyColumns != nil {
		return true
	}

	return false
}

// SetPropertyColumns gets a reference to the given []string and assigns it to the PropertyColumns field.
func (o *BTPropertiesTableTemplateParams) SetPropertyColumns(v []string) {
	o.PropertyColumns = v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateParams) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateParams) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateParams) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *BTPropertiesTableTemplateParams) SetTableType(v string) {
	o.TableType = &v
}

func (o BTPropertiesTableTemplateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.IsAllCaps != nil {
		toSerialize["isAllCaps"] = o.IsAllCaps
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PropertyColumns != nil {
		toSerialize["propertyColumns"] = o.PropertyColumns
	}
	if o.TableType != nil {
		toSerialize["tableType"] = o.TableType
	}
	return json.Marshal(toSerialize)
}

type NullableBTPropertiesTableTemplateParams struct {
	value *BTPropertiesTableTemplateParams
	isSet bool
}

func (v NullableBTPropertiesTableTemplateParams) Get() *BTPropertiesTableTemplateParams {
	return v.value
}

func (v *NullableBTPropertiesTableTemplateParams) Set(val *BTPropertiesTableTemplateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPropertiesTableTemplateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPropertiesTableTemplateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPropertiesTableTemplateParams(val *BTPropertiesTableTemplateParams) *NullableBTPropertiesTableTemplateParams {
	return &NullableBTPropertiesTableTemplateParams{value: val, isSet: true}
}

func (v NullableBTPropertiesTableTemplateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPropertiesTableTemplateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
