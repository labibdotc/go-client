/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.20965-e01b1f4d96d1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyInsertTransformedInstancesResponse struct for BTAssemblyInsertTransformedInstancesResponse
type BTAssemblyInsertTransformedInstancesResponse struct {
	InsertInstanceResponses []BTAssemblyInstanceOccurrenceInfo `json:"insertInstanceResponses,omitempty"`
	InsertResponses         []BTAssemblyOccurrenceInfo         `json:"insertResponses,omitempty"`
}

// NewBTAssemblyInsertTransformedInstancesResponse instantiates a new BTAssemblyInsertTransformedInstancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyInsertTransformedInstancesResponse() *BTAssemblyInsertTransformedInstancesResponse {
	this := BTAssemblyInsertTransformedInstancesResponse{}
	return &this
}

// NewBTAssemblyInsertTransformedInstancesResponseWithDefaults instantiates a new BTAssemblyInsertTransformedInstancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyInsertTransformedInstancesResponseWithDefaults() *BTAssemblyInsertTransformedInstancesResponse {
	this := BTAssemblyInsertTransformedInstancesResponse{}
	return &this
}

// GetInsertInstanceResponses returns the InsertInstanceResponses field value if set, zero value otherwise.
func (o *BTAssemblyInsertTransformedInstancesResponse) GetInsertInstanceResponses() []BTAssemblyInstanceOccurrenceInfo {
	if o == nil || o.InsertInstanceResponses == nil {
		var ret []BTAssemblyInstanceOccurrenceInfo
		return ret
	}
	return o.InsertInstanceResponses
}

// GetInsertInstanceResponsesOk returns a tuple with the InsertInstanceResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInsertTransformedInstancesResponse) GetInsertInstanceResponsesOk() ([]BTAssemblyInstanceOccurrenceInfo, bool) {
	if o == nil || o.InsertInstanceResponses == nil {
		return nil, false
	}
	return o.InsertInstanceResponses, true
}

// HasInsertInstanceResponses returns a boolean if a field has been set.
func (o *BTAssemblyInsertTransformedInstancesResponse) HasInsertInstanceResponses() bool {
	if o != nil && o.InsertInstanceResponses != nil {
		return true
	}

	return false
}

// SetInsertInstanceResponses gets a reference to the given []BTAssemblyInstanceOccurrenceInfo and assigns it to the InsertInstanceResponses field.
func (o *BTAssemblyInsertTransformedInstancesResponse) SetInsertInstanceResponses(v []BTAssemblyInstanceOccurrenceInfo) {
	o.InsertInstanceResponses = v
}

// GetInsertResponses returns the InsertResponses field value if set, zero value otherwise.
func (o *BTAssemblyInsertTransformedInstancesResponse) GetInsertResponses() []BTAssemblyOccurrenceInfo {
	if o == nil || o.InsertResponses == nil {
		var ret []BTAssemblyOccurrenceInfo
		return ret
	}
	return o.InsertResponses
}

// GetInsertResponsesOk returns a tuple with the InsertResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInsertTransformedInstancesResponse) GetInsertResponsesOk() ([]BTAssemblyOccurrenceInfo, bool) {
	if o == nil || o.InsertResponses == nil {
		return nil, false
	}
	return o.InsertResponses, true
}

// HasInsertResponses returns a boolean if a field has been set.
func (o *BTAssemblyInsertTransformedInstancesResponse) HasInsertResponses() bool {
	if o != nil && o.InsertResponses != nil {
		return true
	}

	return false
}

// SetInsertResponses gets a reference to the given []BTAssemblyOccurrenceInfo and assigns it to the InsertResponses field.
func (o *BTAssemblyInsertTransformedInstancesResponse) SetInsertResponses(v []BTAssemblyOccurrenceInfo) {
	o.InsertResponses = v
}

func (o BTAssemblyInsertTransformedInstancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InsertInstanceResponses != nil {
		toSerialize["insertInstanceResponses"] = o.InsertInstanceResponses
	}
	if o.InsertResponses != nil {
		toSerialize["insertResponses"] = o.InsertResponses
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyInsertTransformedInstancesResponse struct {
	value *BTAssemblyInsertTransformedInstancesResponse
	isSet bool
}

func (v NullableBTAssemblyInsertTransformedInstancesResponse) Get() *BTAssemblyInsertTransformedInstancesResponse {
	return v.value
}

func (v *NullableBTAssemblyInsertTransformedInstancesResponse) Set(val *BTAssemblyInsertTransformedInstancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyInsertTransformedInstancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyInsertTransformedInstancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyInsertTransformedInstancesResponse(val *BTAssemblyInsertTransformedInstancesResponse) *NullableBTAssemblyInsertTransformedInstancesResponse {
	return &NullableBTAssemblyInsertTransformedInstancesResponse{value: val, isSet: true}
}

func (v NullableBTAssemblyInsertTransformedInstancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyInsertTransformedInstancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
