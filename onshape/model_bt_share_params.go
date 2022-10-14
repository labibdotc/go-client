/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6905-ae59ed040327
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTShareParams struct for BTShareParams
type BTShareParams struct {
	DocumentId           *string              `json:"documentId,omitempty"`
	ElementId            *string              `json:"elementId,omitempty"`
	EncodedConfiguration *string              `json:"encodedConfiguration,omitempty"`
	Entries              []BTShareEntryParams `json:"entries,omitempty"`
	FolderId             *string              `json:"folderId,omitempty"`
	Message              *string              `json:"message,omitempty"`
	Permission           *int64               `json:"permission,omitempty"`
	PermissionSet        []string             `json:"permissionSet,omitempty"`
	Update               *bool                `json:"update,omitempty"`
	WorkspaceId          *string              `json:"workspaceId,omitempty"`
}

// NewBTShareParams instantiates a new BTShareParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTShareParams() *BTShareParams {
	this := BTShareParams{}
	return &this
}

// NewBTShareParamsWithDefaults instantiates a new BTShareParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTShareParamsWithDefaults() *BTShareParams {
	this := BTShareParams{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTShareParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTShareParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTShareParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTShareParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTShareParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTShareParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetEncodedConfiguration returns the EncodedConfiguration field value if set, zero value otherwise.
func (o *BTShareParams) GetEncodedConfiguration() string {
	if o == nil || o.EncodedConfiguration == nil {
		var ret string
		return ret
	}
	return *o.EncodedConfiguration
}

// GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetEncodedConfigurationOk() (*string, bool) {
	if o == nil || o.EncodedConfiguration == nil {
		return nil, false
	}
	return o.EncodedConfiguration, true
}

// HasEncodedConfiguration returns a boolean if a field has been set.
func (o *BTShareParams) HasEncodedConfiguration() bool {
	if o != nil && o.EncodedConfiguration != nil {
		return true
	}

	return false
}

// SetEncodedConfiguration gets a reference to the given string and assigns it to the EncodedConfiguration field.
func (o *BTShareParams) SetEncodedConfiguration(v string) {
	o.EncodedConfiguration = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTShareParams) GetEntries() []BTShareEntryParams {
	if o == nil || o.Entries == nil {
		var ret []BTShareEntryParams
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetEntriesOk() ([]BTShareEntryParams, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTShareParams) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BTShareEntryParams and assigns it to the Entries field.
func (o *BTShareParams) SetEntries(v []BTShareEntryParams) {
	o.Entries = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *BTShareParams) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *BTShareParams) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *BTShareParams) SetFolderId(v string) {
	o.FolderId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BTShareParams) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BTShareParams) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BTShareParams) SetMessage(v string) {
	o.Message = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *BTShareParams) GetPermission() int64 {
	if o == nil || o.Permission == nil {
		var ret int64
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetPermissionOk() (*int64, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *BTShareParams) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given int64 and assigns it to the Permission field.
func (o *BTShareParams) SetPermission(v int64) {
	o.Permission = &v
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *BTShareParams) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetPermissionSetOk() ([]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *BTShareParams) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *BTShareParams) SetPermissionSet(v []string) {
	o.PermissionSet = v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *BTShareParams) GetUpdate() bool {
	if o == nil || o.Update == nil {
		var ret bool
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetUpdateOk() (*bool, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *BTShareParams) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given bool and assigns it to the Update field.
func (o *BTShareParams) SetUpdate(v bool) {
	o.Update = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTShareParams) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareParams) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTShareParams) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTShareParams) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTShareParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.EncodedConfiguration != nil {
		toSerialize["encodedConfiguration"] = o.EncodedConfiguration
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.FolderId != nil {
		toSerialize["folderId"] = o.FolderId
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTShareParams struct {
	value *BTShareParams
	isSet bool
}

func (v NullableBTShareParams) Get() *BTShareParams {
	return v.value
}

func (v *NullableBTShareParams) Set(val *BTShareParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTShareParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTShareParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTShareParams(val *BTShareParams) *NullableBTShareParams {
	return &NullableBTShareParams{value: val, isSet: true}
}

func (v NullableBTShareParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTShareParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
