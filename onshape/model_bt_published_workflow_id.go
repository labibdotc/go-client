/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7010-841c6a8f62e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPublishedWorkflowId struct for BTPublishedWorkflowId
type BTPublishedWorkflowId struct {
	CompanyId  *string `json:"companyId,omitempty"`
	VersionId  *string `json:"versionId,omitempty"`
	WorkflowId *string `json:"workflowId,omitempty"`
}

// NewBTPublishedWorkflowId instantiates a new BTPublishedWorkflowId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPublishedWorkflowId() *BTPublishedWorkflowId {
	this := BTPublishedWorkflowId{}
	return &this
}

// NewBTPublishedWorkflowIdWithDefaults instantiates a new BTPublishedWorkflowId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPublishedWorkflowIdWithDefaults() *BTPublishedWorkflowId {
	this := BTPublishedWorkflowId{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowId) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowId) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowId) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTPublishedWorkflowId) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowId) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowId) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowId) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTPublishedWorkflowId) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowId) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowId) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowId) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *BTPublishedWorkflowId) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

func (o BTPublishedWorkflowId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	return json.Marshal(toSerialize)
}

type NullableBTPublishedWorkflowId struct {
	value *BTPublishedWorkflowId
	isSet bool
}

func (v NullableBTPublishedWorkflowId) Get() *BTPublishedWorkflowId {
	return v.value
}

func (v *NullableBTPublishedWorkflowId) Set(val *BTPublishedWorkflowId) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPublishedWorkflowId) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPublishedWorkflowId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPublishedWorkflowId(val *BTPublishedWorkflowId) *NullableBTPublishedWorkflowId {
	return &NullableBTPublishedWorkflowId{value: val, isSet: true}
}

func (v NullableBTPublishedWorkflowId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPublishedWorkflowId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
