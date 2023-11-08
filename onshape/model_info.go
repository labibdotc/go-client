/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25541-b54c5b71ef45
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Info struct for Info
type Info struct {
	Contact        *Contact                          `json:"contact,omitempty"`
	Description    *string                           `json:"description,omitempty"`
	Extensions     map[string]map[string]interface{} `json:"extensions,omitempty"`
	License        *License                          `json:"license,omitempty"`
	Summary        *string                           `json:"summary,omitempty"`
	TermsOfService *string                           `json:"termsOfService,omitempty"`
	Title          *string                           `json:"title,omitempty"`
	Version        *string                           `json:"version,omitempty"`
}

// NewInfo instantiates a new Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfo() *Info {
	this := Info{}
	return &this
}

// NewInfoWithDefaults instantiates a new Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoWithDefaults() *Info {
	this := Info{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Info) GetContact() Contact {
	if o == nil || o.Contact == nil {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetContactOk() (*Contact, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Info) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *Info) SetContact(v Contact) {
	o.Contact = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Info) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Info) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Info) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Info) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Info) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Info) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *Info) GetLicense() License {
	if o == nil || o.License == nil {
		var ret License
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetLicenseOk() (*License, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *Info) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given License and assigns it to the License field.
func (o *Info) SetLicense(v License) {
	o.License = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Info) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Info) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Info) SetSummary(v string) {
	o.Summary = &v
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise.
func (o *Info) GetTermsOfService() string {
	if o == nil || o.TermsOfService == nil {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || o.TermsOfService == nil {
		return nil, false
	}
	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *Info) HasTermsOfService() bool {
	if o != nil && o.TermsOfService != nil {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *Info) SetTermsOfService(v string) {
	o.TermsOfService = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Info) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Info) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Info) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Info) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Info) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Info) SetVersion(v string) {
	o.Version = &v
}

func (o Info) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.TermsOfService != nil {
		toSerialize["termsOfService"] = o.TermsOfService
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInfo struct {
	value *Info
	isSet bool
}

func (v NullableInfo) Get() *Info {
	return v.value
}

func (v *NullableInfo) Set(val *Info) {
	v.value = val
	v.isSet = true
}

func (v NullableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfo(val *Info) *NullableInfo {
	return &NullableInfo{value: val, isSet: true}
}

func (v NullableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
