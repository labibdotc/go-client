/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7108-42dac6f29840
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartCustomProperties1338 struct for BTPartCustomProperties1338
type BTPartCustomProperties1338 struct {
	BtType              *string            `json:"btType,omitempty"`
	Description         *string            `json:"description,omitempty"`
	PartNumber          *string            `json:"partNumber,omitempty"`
	ProductLine         *string            `json:"productLine,omitempty"`
	Project             *string            `json:"project,omitempty"`
	Properties          *map[string]string `json:"properties,omitempty"`
	Revision            *string            `json:"revision,omitempty"`
	TessellationSetting *string            `json:"tessellationSetting,omitempty"`
	Title1              *string            `json:"title1,omitempty"`
	Title2              *string            `json:"title2,omitempty"`
	Title3              *string            `json:"title3,omitempty"`
	Vendor              *string            `json:"vendor,omitempty"`
}

// NewBTPartCustomProperties1338 instantiates a new BTPartCustomProperties1338 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartCustomProperties1338() *BTPartCustomProperties1338 {
	this := BTPartCustomProperties1338{}
	return &this
}

// NewBTPartCustomProperties1338WithDefaults instantiates a new BTPartCustomProperties1338 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartCustomProperties1338WithDefaults() *BTPartCustomProperties1338 {
	this := BTPartCustomProperties1338{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartCustomProperties1338) SetBtType(v string) {
	o.BtType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTPartCustomProperties1338) SetDescription(v string) {
	o.Description = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTPartCustomProperties1338) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductLine returns the ProductLine field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetProductLine() string {
	if o == nil || o.ProductLine == nil {
		var ret string
		return ret
	}
	return *o.ProductLine
}

// GetProductLineOk returns a tuple with the ProductLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetProductLineOk() (*string, bool) {
	if o == nil || o.ProductLine == nil {
		return nil, false
	}
	return o.ProductLine, true
}

// HasProductLine returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasProductLine() bool {
	if o != nil && o.ProductLine != nil {
		return true
	}

	return false
}

// SetProductLine gets a reference to the given string and assigns it to the ProductLine field.
func (o *BTPartCustomProperties1338) SetProductLine(v string) {
	o.ProductLine = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *BTPartCustomProperties1338) SetProject(v string) {
	o.Project = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *BTPartCustomProperties1338) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTPartCustomProperties1338) SetRevision(v string) {
	o.Revision = &v
}

// GetTessellationSetting returns the TessellationSetting field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetTessellationSetting() string {
	if o == nil || o.TessellationSetting == nil {
		var ret string
		return ret
	}
	return *o.TessellationSetting
}

// GetTessellationSettingOk returns a tuple with the TessellationSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetTessellationSettingOk() (*string, bool) {
	if o == nil || o.TessellationSetting == nil {
		return nil, false
	}
	return o.TessellationSetting, true
}

// HasTessellationSetting returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasTessellationSetting() bool {
	if o != nil && o.TessellationSetting != nil {
		return true
	}

	return false
}

// SetTessellationSetting gets a reference to the given string and assigns it to the TessellationSetting field.
func (o *BTPartCustomProperties1338) SetTessellationSetting(v string) {
	o.TessellationSetting = &v
}

// GetTitle1 returns the Title1 field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetTitle1() string {
	if o == nil || o.Title1 == nil {
		var ret string
		return ret
	}
	return *o.Title1
}

// GetTitle1Ok returns a tuple with the Title1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetTitle1Ok() (*string, bool) {
	if o == nil || o.Title1 == nil {
		return nil, false
	}
	return o.Title1, true
}

// HasTitle1 returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasTitle1() bool {
	if o != nil && o.Title1 != nil {
		return true
	}

	return false
}

// SetTitle1 gets a reference to the given string and assigns it to the Title1 field.
func (o *BTPartCustomProperties1338) SetTitle1(v string) {
	o.Title1 = &v
}

// GetTitle2 returns the Title2 field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetTitle2() string {
	if o == nil || o.Title2 == nil {
		var ret string
		return ret
	}
	return *o.Title2
}

// GetTitle2Ok returns a tuple with the Title2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetTitle2Ok() (*string, bool) {
	if o == nil || o.Title2 == nil {
		return nil, false
	}
	return o.Title2, true
}

// HasTitle2 returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasTitle2() bool {
	if o != nil && o.Title2 != nil {
		return true
	}

	return false
}

// SetTitle2 gets a reference to the given string and assigns it to the Title2 field.
func (o *BTPartCustomProperties1338) SetTitle2(v string) {
	o.Title2 = &v
}

// GetTitle3 returns the Title3 field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetTitle3() string {
	if o == nil || o.Title3 == nil {
		var ret string
		return ret
	}
	return *o.Title3
}

// GetTitle3Ok returns a tuple with the Title3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetTitle3Ok() (*string, bool) {
	if o == nil || o.Title3 == nil {
		return nil, false
	}
	return o.Title3, true
}

// HasTitle3 returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasTitle3() bool {
	if o != nil && o.Title3 != nil {
		return true
	}

	return false
}

// SetTitle3 gets a reference to the given string and assigns it to the Title3 field.
func (o *BTPartCustomProperties1338) SetTitle3(v string) {
	o.Title3 = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BTPartCustomProperties1338) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartCustomProperties1338) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BTPartCustomProperties1338) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BTPartCustomProperties1338) SetVendor(v string) {
	o.Vendor = &v
}

func (o BTPartCustomProperties1338) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.ProductLine != nil {
		toSerialize["productLine"] = o.ProductLine
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.TessellationSetting != nil {
		toSerialize["tessellationSetting"] = o.TessellationSetting
	}
	if o.Title1 != nil {
		toSerialize["title1"] = o.Title1
	}
	if o.Title2 != nil {
		toSerialize["title2"] = o.Title2
	}
	if o.Title3 != nil {
		toSerialize["title3"] = o.Title3
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartCustomProperties1338 struct {
	value *BTPartCustomProperties1338
	isSet bool
}

func (v NullableBTPartCustomProperties1338) Get() *BTPartCustomProperties1338 {
	return v.value
}

func (v *NullableBTPartCustomProperties1338) Set(val *BTPartCustomProperties1338) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartCustomProperties1338) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartCustomProperties1338) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartCustomProperties1338(val *BTPartCustomProperties1338) *NullableBTPartCustomProperties1338 {
	return &NullableBTPartCustomProperties1338{value: val, isSet: true}
}

func (v NullableBTPartCustomProperties1338) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartCustomProperties1338) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
