/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6373-48cecdd7807c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMoveElementInfo struct for BTMoveElementInfo
type BTMoveElementInfo struct {
	ElementOriginalToNewMap *map[string]string `json:"elementOriginalToNewMap,omitempty"`
	ErrorMessage            *string            `json:"errorMessage,omitempty"`
	IsNewDocument           *bool              `json:"isNewDocument,omitempty"`
	NewDocumentId           *string            `json:"newDocumentId,omitempty"`
	NewDocumentName         *string            `json:"newDocumentName,omitempty"`
	NewDocumentVersionId    *string            `json:"newDocumentVersionId,omitempty"`
	NewWorkspaceId          *string            `json:"newWorkspaceId,omitempty"`
}

// NewBTMoveElementInfo instantiates a new BTMoveElementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMoveElementInfo() *BTMoveElementInfo {
	this := BTMoveElementInfo{}
	return &this
}

// NewBTMoveElementInfoWithDefaults instantiates a new BTMoveElementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMoveElementInfoWithDefaults() *BTMoveElementInfo {
	this := BTMoveElementInfo{}
	return &this
}

// GetElementOriginalToNewMap returns the ElementOriginalToNewMap field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetElementOriginalToNewMap() map[string]string {
	if o == nil || o.ElementOriginalToNewMap == nil {
		var ret map[string]string
		return ret
	}
	return *o.ElementOriginalToNewMap
}

// GetElementOriginalToNewMapOk returns a tuple with the ElementOriginalToNewMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetElementOriginalToNewMapOk() (*map[string]string, bool) {
	if o == nil || o.ElementOriginalToNewMap == nil {
		return nil, false
	}
	return o.ElementOriginalToNewMap, true
}

// HasElementOriginalToNewMap returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasElementOriginalToNewMap() bool {
	if o != nil && o.ElementOriginalToNewMap != nil {
		return true
	}

	return false
}

// SetElementOriginalToNewMap gets a reference to the given map[string]string and assigns it to the ElementOriginalToNewMap field.
func (o *BTMoveElementInfo) SetElementOriginalToNewMap(v map[string]string) {
	o.ElementOriginalToNewMap = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTMoveElementInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsNewDocument returns the IsNewDocument field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetIsNewDocument() bool {
	if o == nil || o.IsNewDocument == nil {
		var ret bool
		return ret
	}
	return *o.IsNewDocument
}

// GetIsNewDocumentOk returns a tuple with the IsNewDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetIsNewDocumentOk() (*bool, bool) {
	if o == nil || o.IsNewDocument == nil {
		return nil, false
	}
	return o.IsNewDocument, true
}

// HasIsNewDocument returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasIsNewDocument() bool {
	if o != nil && o.IsNewDocument != nil {
		return true
	}

	return false
}

// SetIsNewDocument gets a reference to the given bool and assigns it to the IsNewDocument field.
func (o *BTMoveElementInfo) SetIsNewDocument(v bool) {
	o.IsNewDocument = &v
}

// GetNewDocumentId returns the NewDocumentId field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetNewDocumentId() string {
	if o == nil || o.NewDocumentId == nil {
		var ret string
		return ret
	}
	return *o.NewDocumentId
}

// GetNewDocumentIdOk returns a tuple with the NewDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetNewDocumentIdOk() (*string, bool) {
	if o == nil || o.NewDocumentId == nil {
		return nil, false
	}
	return o.NewDocumentId, true
}

// HasNewDocumentId returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasNewDocumentId() bool {
	if o != nil && o.NewDocumentId != nil {
		return true
	}

	return false
}

// SetNewDocumentId gets a reference to the given string and assigns it to the NewDocumentId field.
func (o *BTMoveElementInfo) SetNewDocumentId(v string) {
	o.NewDocumentId = &v
}

// GetNewDocumentName returns the NewDocumentName field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetNewDocumentName() string {
	if o == nil || o.NewDocumentName == nil {
		var ret string
		return ret
	}
	return *o.NewDocumentName
}

// GetNewDocumentNameOk returns a tuple with the NewDocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetNewDocumentNameOk() (*string, bool) {
	if o == nil || o.NewDocumentName == nil {
		return nil, false
	}
	return o.NewDocumentName, true
}

// HasNewDocumentName returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasNewDocumentName() bool {
	if o != nil && o.NewDocumentName != nil {
		return true
	}

	return false
}

// SetNewDocumentName gets a reference to the given string and assigns it to the NewDocumentName field.
func (o *BTMoveElementInfo) SetNewDocumentName(v string) {
	o.NewDocumentName = &v
}

// GetNewDocumentVersionId returns the NewDocumentVersionId field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetNewDocumentVersionId() string {
	if o == nil || o.NewDocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.NewDocumentVersionId
}

// GetNewDocumentVersionIdOk returns a tuple with the NewDocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetNewDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.NewDocumentVersionId == nil {
		return nil, false
	}
	return o.NewDocumentVersionId, true
}

// HasNewDocumentVersionId returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasNewDocumentVersionId() bool {
	if o != nil && o.NewDocumentVersionId != nil {
		return true
	}

	return false
}

// SetNewDocumentVersionId gets a reference to the given string and assigns it to the NewDocumentVersionId field.
func (o *BTMoveElementInfo) SetNewDocumentVersionId(v string) {
	o.NewDocumentVersionId = &v
}

// GetNewWorkspaceId returns the NewWorkspaceId field value if set, zero value otherwise.
func (o *BTMoveElementInfo) GetNewWorkspaceId() string {
	if o == nil || o.NewWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.NewWorkspaceId
}

// GetNewWorkspaceIdOk returns a tuple with the NewWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementInfo) GetNewWorkspaceIdOk() (*string, bool) {
	if o == nil || o.NewWorkspaceId == nil {
		return nil, false
	}
	return o.NewWorkspaceId, true
}

// HasNewWorkspaceId returns a boolean if a field has been set.
func (o *BTMoveElementInfo) HasNewWorkspaceId() bool {
	if o != nil && o.NewWorkspaceId != nil {
		return true
	}

	return false
}

// SetNewWorkspaceId gets a reference to the given string and assigns it to the NewWorkspaceId field.
func (o *BTMoveElementInfo) SetNewWorkspaceId(v string) {
	o.NewWorkspaceId = &v
}

func (o BTMoveElementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ElementOriginalToNewMap != nil {
		toSerialize["elementOriginalToNewMap"] = o.ElementOriginalToNewMap
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.IsNewDocument != nil {
		toSerialize["isNewDocument"] = o.IsNewDocument
	}
	if o.NewDocumentId != nil {
		toSerialize["newDocumentId"] = o.NewDocumentId
	}
	if o.NewDocumentName != nil {
		toSerialize["newDocumentName"] = o.NewDocumentName
	}
	if o.NewDocumentVersionId != nil {
		toSerialize["newDocumentVersionId"] = o.NewDocumentVersionId
	}
	if o.NewWorkspaceId != nil {
		toSerialize["newWorkspaceId"] = o.NewWorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMoveElementInfo struct {
	value *BTMoveElementInfo
	isSet bool
}

func (v NullableBTMoveElementInfo) Get() *BTMoveElementInfo {
	return v.value
}

func (v *NullableBTMoveElementInfo) Set(val *BTMoveElementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMoveElementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMoveElementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMoveElementInfo(val *BTMoveElementInfo) *NullableBTMoveElementInfo {
	return &NullableBTMoveElementInfo{value: val, isSet: true}
}

func (v NullableBTMoveElementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMoveElementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
