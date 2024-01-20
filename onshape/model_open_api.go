/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29320-74695940af99
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// OpenAPI struct for OpenAPI
type OpenAPI struct {
	Components        *Components                       `json:"components,omitempty"`
	Extensions        map[string]map[string]interface{} `json:"extensions,omitempty"`
	ExternalDocs      *ExternalDocumentation            `json:"externalDocs,omitempty"`
	Info              *Info                             `json:"info,omitempty"`
	JsonSchemaDialect *string                           `json:"jsonSchemaDialect,omitempty"`
	Openapi           *string                           `json:"openapi,omitempty"`
	Paths             *OpenAPIPaths                     `json:"paths,omitempty"`
	Security          []SecurityRequirement             `json:"security,omitempty"`
	Servers           []Server                          `json:"servers,omitempty"`
	Tags              []Tag                             `json:"tags,omitempty"`
	Webhooks          *map[string]PathItem              `json:"webhooks,omitempty"`
}

// NewOpenAPI instantiates a new OpenAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAPI() *OpenAPI {
	this := OpenAPI{}
	return &this
}

// NewOpenAPIWithDefaults instantiates a new OpenAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAPIWithDefaults() *OpenAPI {
	this := OpenAPI{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *OpenAPI) GetComponents() Components {
	if o == nil || o.Components == nil {
		var ret Components
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetComponentsOk() (*Components, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *OpenAPI) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given Components and assigns it to the Components field.
func (o *OpenAPI) SetComponents(v Components) {
	o.Components = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *OpenAPI) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *OpenAPI) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *OpenAPI) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExternalDocs returns the ExternalDocs field value if set, zero value otherwise.
func (o *OpenAPI) GetExternalDocs() ExternalDocumentation {
	if o == nil || o.ExternalDocs == nil {
		var ret ExternalDocumentation
		return ret
	}
	return *o.ExternalDocs
}

// GetExternalDocsOk returns a tuple with the ExternalDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetExternalDocsOk() (*ExternalDocumentation, bool) {
	if o == nil || o.ExternalDocs == nil {
		return nil, false
	}
	return o.ExternalDocs, true
}

// HasExternalDocs returns a boolean if a field has been set.
func (o *OpenAPI) HasExternalDocs() bool {
	if o != nil && o.ExternalDocs != nil {
		return true
	}

	return false
}

// SetExternalDocs gets a reference to the given ExternalDocumentation and assigns it to the ExternalDocs field.
func (o *OpenAPI) SetExternalDocs(v ExternalDocumentation) {
	o.ExternalDocs = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *OpenAPI) GetInfo() Info {
	if o == nil || o.Info == nil {
		var ret Info
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetInfoOk() (*Info, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OpenAPI) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given Info and assigns it to the Info field.
func (o *OpenAPI) SetInfo(v Info) {
	o.Info = &v
}

// GetJsonSchemaDialect returns the JsonSchemaDialect field value if set, zero value otherwise.
func (o *OpenAPI) GetJsonSchemaDialect() string {
	if o == nil || o.JsonSchemaDialect == nil {
		var ret string
		return ret
	}
	return *o.JsonSchemaDialect
}

// GetJsonSchemaDialectOk returns a tuple with the JsonSchemaDialect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetJsonSchemaDialectOk() (*string, bool) {
	if o == nil || o.JsonSchemaDialect == nil {
		return nil, false
	}
	return o.JsonSchemaDialect, true
}

// HasJsonSchemaDialect returns a boolean if a field has been set.
func (o *OpenAPI) HasJsonSchemaDialect() bool {
	if o != nil && o.JsonSchemaDialect != nil {
		return true
	}

	return false
}

// SetJsonSchemaDialect gets a reference to the given string and assigns it to the JsonSchemaDialect field.
func (o *OpenAPI) SetJsonSchemaDialect(v string) {
	o.JsonSchemaDialect = &v
}

// GetOpenapi returns the Openapi field value if set, zero value otherwise.
func (o *OpenAPI) GetOpenapi() string {
	if o == nil || o.Openapi == nil {
		var ret string
		return ret
	}
	return *o.Openapi
}

// GetOpenapiOk returns a tuple with the Openapi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetOpenapiOk() (*string, bool) {
	if o == nil || o.Openapi == nil {
		return nil, false
	}
	return o.Openapi, true
}

// HasOpenapi returns a boolean if a field has been set.
func (o *OpenAPI) HasOpenapi() bool {
	if o != nil && o.Openapi != nil {
		return true
	}

	return false
}

// SetOpenapi gets a reference to the given string and assigns it to the Openapi field.
func (o *OpenAPI) SetOpenapi(v string) {
	o.Openapi = &v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *OpenAPI) GetPaths() OpenAPIPaths {
	if o == nil || o.Paths == nil {
		var ret OpenAPIPaths
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetPathsOk() (*OpenAPIPaths, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *OpenAPI) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given OpenAPIPaths and assigns it to the Paths field.
func (o *OpenAPI) SetPaths(v OpenAPIPaths) {
	o.Paths = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *OpenAPI) GetSecurity() []SecurityRequirement {
	if o == nil || o.Security == nil {
		var ret []SecurityRequirement
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetSecurityOk() ([]SecurityRequirement, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *OpenAPI) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SecurityRequirement and assigns it to the Security field.
func (o *OpenAPI) SetSecurity(v []SecurityRequirement) {
	o.Security = v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *OpenAPI) GetServers() []Server {
	if o == nil || o.Servers == nil {
		var ret []Server
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetServersOk() ([]Server, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *OpenAPI) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server and assigns it to the Servers field.
func (o *OpenAPI) SetServers(v []Server) {
	o.Servers = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OpenAPI) GetTags() []Tag {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetTagsOk() ([]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OpenAPI) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *OpenAPI) SetTags(v []Tag) {
	o.Tags = v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *OpenAPI) GetWebhooks() map[string]PathItem {
	if o == nil || o.Webhooks == nil {
		var ret map[string]PathItem
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPI) GetWebhooksOk() (*map[string]PathItem, bool) {
	if o == nil || o.Webhooks == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *OpenAPI) HasWebhooks() bool {
	if o != nil && o.Webhooks != nil {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given map[string]PathItem and assigns it to the Webhooks field.
func (o *OpenAPI) SetWebhooks(v map[string]PathItem) {
	o.Webhooks = &v
}

func (o OpenAPI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.ExternalDocs != nil {
		toSerialize["externalDocs"] = o.ExternalDocs
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.JsonSchemaDialect != nil {
		toSerialize["jsonSchemaDialect"] = o.JsonSchemaDialect
	}
	if o.Openapi != nil {
		toSerialize["openapi"] = o.Openapi
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Webhooks != nil {
		toSerialize["webhooks"] = o.Webhooks
	}
	return json.Marshal(toSerialize)
}

type NullableOpenAPI struct {
	value *OpenAPI
	isSet bool
}

func (v NullableOpenAPI) Get() *OpenAPI {
	return v.value
}

func (v *NullableOpenAPI) Set(val *OpenAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAPI(val *OpenAPI) *NullableOpenAPI {
	return &NullableOpenAPI{value: val, isSet: true}
}

func (v NullableOpenAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
