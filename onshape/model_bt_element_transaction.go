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

// BTElementTransaction struct for BTElementTransaction
type BTElementTransaction struct {
	Description   *string `json:"description,omitempty"`
	DocumentId    *string `json:"documentId,omitempty"`
	ElementId     *string `json:"elementId,omitempty"`
	Id            *string `json:"id,omitempty"`
	MicrobranchId *string `json:"microbranchId,omitempty"`
	WorkspaceId   *string `json:"workspaceId,omitempty"`
}

// NewBTElementTransaction instantiates a new BTElementTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementTransaction() *BTElementTransaction {
	this := BTElementTransaction{}
	return &this
}

// NewBTElementTransactionWithDefaults instantiates a new BTElementTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementTransactionWithDefaults() *BTElementTransaction {
	this := BTElementTransaction{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTElementTransaction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTElementTransaction) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTElementTransaction) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTElementTransaction) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementTransaction) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTElementTransaction) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTElementTransaction) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTElementTransaction) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementTransaction) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTElementTransaction) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTElementTransaction) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTElementTransaction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementTransaction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTElementTransaction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTElementTransaction) SetId(v string) {
	o.Id = &v
}

// GetMicrobranchId returns the MicrobranchId field value if set, zero value otherwise.
func (o *BTElementTransaction) GetMicrobranchId() string {
	if o == nil || o.MicrobranchId == nil {
		var ret string
		return ret
	}
	return *o.MicrobranchId
}

// GetMicrobranchIdOk returns a tuple with the MicrobranchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementTransaction) GetMicrobranchIdOk() (*string, bool) {
	if o == nil || o.MicrobranchId == nil {
		return nil, false
	}
	return o.MicrobranchId, true
}

// HasMicrobranchId returns a boolean if a field has been set.
func (o *BTElementTransaction) HasMicrobranchId() bool {
	if o != nil && o.MicrobranchId != nil {
		return true
	}

	return false
}

// SetMicrobranchId gets a reference to the given string and assigns it to the MicrobranchId field.
func (o *BTElementTransaction) SetMicrobranchId(v string) {
	o.MicrobranchId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTElementTransaction) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementTransaction) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTElementTransaction) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTElementTransaction) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTElementTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MicrobranchId != nil {
		toSerialize["microbranchId"] = o.MicrobranchId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTElementTransaction struct {
	value *BTElementTransaction
	isSet bool
}

func (v NullableBTElementTransaction) Get() *BTElementTransaction {
	return v.value
}

func (v *NullableBTElementTransaction) Set(val *BTElementTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementTransaction(val *BTElementTransaction) *NullableBTElementTransaction {
	return &NullableBTElementTransaction{value: val, isSet: true}
}

func (v NullableBTElementTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
