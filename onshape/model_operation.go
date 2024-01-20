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

// Operation struct for Operation
type Operation struct {
	Callbacks    *map[string]Callback              `json:"callbacks,omitempty"`
	Deprecated   *bool                             `json:"deprecated,omitempty"`
	Description  *string                           `json:"description,omitempty"`
	Extensions   map[string]map[string]interface{} `json:"extensions,omitempty"`
	ExternalDocs *ExternalDocumentation            `json:"externalDocs,omitempty"`
	OperationId  *string                           `json:"operationId,omitempty"`
	Parameters   []Parameter                       `json:"parameters,omitempty"`
	RequestBody  *RequestBody                      `json:"requestBody,omitempty"`
	Responses    *OperationResponses               `json:"responses,omitempty"`
	Security     []SecurityRequirement             `json:"security,omitempty"`
	Servers      []Server                          `json:"servers,omitempty"`
	Summary      *string                           `json:"summary,omitempty"`
	Tags         []string                          `json:"tags,omitempty"`
}

// NewOperation instantiates a new Operation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperation() *Operation {
	this := Operation{}
	return &this
}

// NewOperationWithDefaults instantiates a new Operation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationWithDefaults() *Operation {
	this := Operation{}
	return &this
}

// GetCallbacks returns the Callbacks field value if set, zero value otherwise.
func (o *Operation) GetCallbacks() map[string]Callback {
	if o == nil || o.Callbacks == nil {
		var ret map[string]Callback
		return ret
	}
	return *o.Callbacks
}

// GetCallbacksOk returns a tuple with the Callbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetCallbacksOk() (*map[string]Callback, bool) {
	if o == nil || o.Callbacks == nil {
		return nil, false
	}
	return o.Callbacks, true
}

// HasCallbacks returns a boolean if a field has been set.
func (o *Operation) HasCallbacks() bool {
	if o != nil && o.Callbacks != nil {
		return true
	}

	return false
}

// SetCallbacks gets a reference to the given map[string]Callback and assigns it to the Callbacks field.
func (o *Operation) SetCallbacks(v map[string]Callback) {
	o.Callbacks = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *Operation) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *Operation) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *Operation) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Operation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Operation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Operation) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Operation) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Operation) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Operation) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExternalDocs returns the ExternalDocs field value if set, zero value otherwise.
func (o *Operation) GetExternalDocs() ExternalDocumentation {
	if o == nil || o.ExternalDocs == nil {
		var ret ExternalDocumentation
		return ret
	}
	return *o.ExternalDocs
}

// GetExternalDocsOk returns a tuple with the ExternalDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetExternalDocsOk() (*ExternalDocumentation, bool) {
	if o == nil || o.ExternalDocs == nil {
		return nil, false
	}
	return o.ExternalDocs, true
}

// HasExternalDocs returns a boolean if a field has been set.
func (o *Operation) HasExternalDocs() bool {
	if o != nil && o.ExternalDocs != nil {
		return true
	}

	return false
}

// SetExternalDocs gets a reference to the given ExternalDocumentation and assigns it to the ExternalDocs field.
func (o *Operation) SetExternalDocs(v ExternalDocumentation) {
	o.ExternalDocs = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *Operation) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *Operation) HasOperationId() bool {
	if o != nil && o.OperationId != nil {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *Operation) SetOperationId(v string) {
	o.OperationId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Operation) GetParameters() []Parameter {
	if o == nil || o.Parameters == nil {
		var ret []Parameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetParametersOk() ([]Parameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Operation) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []Parameter and assigns it to the Parameters field.
func (o *Operation) SetParameters(v []Parameter) {
	o.Parameters = v
}

// GetRequestBody returns the RequestBody field value if set, zero value otherwise.
func (o *Operation) GetRequestBody() RequestBody {
	if o == nil || o.RequestBody == nil {
		var ret RequestBody
		return ret
	}
	return *o.RequestBody
}

// GetRequestBodyOk returns a tuple with the RequestBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetRequestBodyOk() (*RequestBody, bool) {
	if o == nil || o.RequestBody == nil {
		return nil, false
	}
	return o.RequestBody, true
}

// HasRequestBody returns a boolean if a field has been set.
func (o *Operation) HasRequestBody() bool {
	if o != nil && o.RequestBody != nil {
		return true
	}

	return false
}

// SetRequestBody gets a reference to the given RequestBody and assigns it to the RequestBody field.
func (o *Operation) SetRequestBody(v RequestBody) {
	o.RequestBody = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *Operation) GetResponses() OperationResponses {
	if o == nil || o.Responses == nil {
		var ret OperationResponses
		return ret
	}
	return *o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetResponsesOk() (*OperationResponses, bool) {
	if o == nil || o.Responses == nil {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *Operation) HasResponses() bool {
	if o != nil && o.Responses != nil {
		return true
	}

	return false
}

// SetResponses gets a reference to the given OperationResponses and assigns it to the Responses field.
func (o *Operation) SetResponses(v OperationResponses) {
	o.Responses = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *Operation) GetSecurity() []SecurityRequirement {
	if o == nil || o.Security == nil {
		var ret []SecurityRequirement
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetSecurityOk() ([]SecurityRequirement, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *Operation) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SecurityRequirement and assigns it to the Security field.
func (o *Operation) SetSecurity(v []SecurityRequirement) {
	o.Security = v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *Operation) GetServers() []Server {
	if o == nil || o.Servers == nil {
		var ret []Server
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetServersOk() ([]Server, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *Operation) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server and assigns it to the Servers field.
func (o *Operation) SetServers(v []Server) {
	o.Servers = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Operation) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Operation) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Operation) SetSummary(v string) {
	o.Summary = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Operation) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Operation) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Operation) SetTags(v []string) {
	o.Tags = v
}

func (o Operation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Callbacks != nil {
		toSerialize["callbacks"] = o.Callbacks
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.ExternalDocs != nil {
		toSerialize["externalDocs"] = o.ExternalDocs
	}
	if o.OperationId != nil {
		toSerialize["operationId"] = o.OperationId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.RequestBody != nil {
		toSerialize["requestBody"] = o.RequestBody
	}
	if o.Responses != nil {
		toSerialize["responses"] = o.Responses
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
