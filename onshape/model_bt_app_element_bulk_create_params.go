/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25784-25d5580b0721
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementBulkCreateParams struct for BTAppElementBulkCreateParams
type BTAppElementBulkCreateParams struct {
	// The label that will appear in the document's edit history for this operation. If blank, a value will be auto-generated.
	Description *string `json:"description,omitempty"`
	// The data type of the application. This string allows an application to distinguish their elements from elements of another application.
	FormatId string                   `json:"formatId"`
	Location *BTElementLocationParams `json:"location,omitempty"`
	// The name of the element being created. If blank, a name will be auto-generated.
	Names []string `json:"names,omitempty"`
}

// NewBTAppElementBulkCreateParams instantiates a new BTAppElementBulkCreateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementBulkCreateParams(formatId string) *BTAppElementBulkCreateParams {
	this := BTAppElementBulkCreateParams{}
	this.FormatId = formatId
	return &this
}

// NewBTAppElementBulkCreateParamsWithDefaults instantiates a new BTAppElementBulkCreateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementBulkCreateParamsWithDefaults() *BTAppElementBulkCreateParams {
	this := BTAppElementBulkCreateParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementBulkCreateParams) SetDescription(v string) {
	o.Description = &v
}

// GetFormatId returns the FormatId field value
func (o *BTAppElementBulkCreateParams) GetFormatId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormatId
}

// GetFormatIdOk returns a tuple with the FormatId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateParams) GetFormatIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormatId, true
}

// SetFormatId sets field value
func (o *BTAppElementBulkCreateParams) SetFormatId(v string) {
	o.FormatId = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateParams) GetLocation() BTElementLocationParams {
	if o == nil || o.Location == nil {
		var ret BTElementLocationParams
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateParams) GetLocationOk() (*BTElementLocationParams, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateParams) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTElementLocationParams and assigns it to the Location field.
func (o *BTAppElementBulkCreateParams) SetLocation(v BTElementLocationParams) {
	o.Location = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *BTAppElementBulkCreateParams) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementBulkCreateParams) GetNamesOk() ([]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *BTAppElementBulkCreateParams) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *BTAppElementBulkCreateParams) SetNames(v []string) {
	o.Names = v
}

func (o BTAppElementBulkCreateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["formatId"] = o.FormatId
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementBulkCreateParams struct {
	value *BTAppElementBulkCreateParams
	isSet bool
}

func (v NullableBTAppElementBulkCreateParams) Get() *BTAppElementBulkCreateParams {
	return v.value
}

func (v *NullableBTAppElementBulkCreateParams) Set(val *BTAppElementBulkCreateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementBulkCreateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementBulkCreateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementBulkCreateParams(val *BTAppElementBulkCreateParams) *NullableBTAppElementBulkCreateParams {
	return &NullableBTAppElementBulkCreateParams{value: val, isSet: true}
}

func (v NullableBTAppElementBulkCreateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementBulkCreateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
