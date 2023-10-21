/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24440-f37a82fd6450
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// XML struct for XML
type XML struct {
	Attribute  *bool                             `json:"attribute,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Namespace  *string                           `json:"namespace,omitempty"`
	Prefix     *string                           `json:"prefix,omitempty"`
	Wrapped    *bool                             `json:"wrapped,omitempty"`
}

// NewXML instantiates a new XML object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXML() *XML {
	this := XML{}
	return &this
}

// NewXMLWithDefaults instantiates a new XML object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXMLWithDefaults() *XML {
	this := XML{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *XML) GetAttribute() bool {
	if o == nil || o.Attribute == nil {
		var ret bool
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XML) GetAttributeOk() (*bool, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *XML) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given bool and assigns it to the Attribute field.
func (o *XML) SetAttribute(v bool) {
	o.Attribute = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *XML) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XML) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *XML) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *XML) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XML) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XML) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XML) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XML) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *XML) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XML) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *XML) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *XML) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *XML) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XML) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *XML) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *XML) SetPrefix(v string) {
	o.Prefix = &v
}

// GetWrapped returns the Wrapped field value if set, zero value otherwise.
func (o *XML) GetWrapped() bool {
	if o == nil || o.Wrapped == nil {
		var ret bool
		return ret
	}
	return *o.Wrapped
}

// GetWrappedOk returns a tuple with the Wrapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XML) GetWrappedOk() (*bool, bool) {
	if o == nil || o.Wrapped == nil {
		return nil, false
	}
	return o.Wrapped, true
}

// HasWrapped returns a boolean if a field has been set.
func (o *XML) HasWrapped() bool {
	if o != nil && o.Wrapped != nil {
		return true
	}

	return false
}

// SetWrapped gets a reference to the given bool and assigns it to the Wrapped field.
func (o *XML) SetWrapped(v bool) {
	o.Wrapped = &v
}

func (o XML) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Wrapped != nil {
		toSerialize["wrapped"] = o.Wrapped
	}
	return json.Marshal(toSerialize)
}

type NullableXML struct {
	value *XML
	isSet bool
}

func (v NullableXML) Get() *XML {
	return v.value
}

func (v *NullableXML) Set(val *XML) {
	v.value = val
	v.isSet = true
}

func (v NullableXML) IsSet() bool {
	return v.isSet
}

func (v *NullableXML) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXML(val *XML) *NullableXML {
	return &NullableXML{value: val, isSet: true}
}

func (v NullableXML) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXML) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
