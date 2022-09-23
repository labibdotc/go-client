/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6590-f8226b4e1789
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSingleAssemblyReferenceDisplayData1557 struct for BTSingleAssemblyReferenceDisplayData1557
type BTSingleAssemblyReferenceDisplayData1557 struct {
	ContextWorkspaceId   *string                      `json:"contextWorkspaceId,omitempty"`
	DocumentId           *string                      `json:"documentId,omitempty"`
	Error                *string                      `json:"error,omitempty"`
	ErrorMessage         *string                      `json:"errorMessage,omitempty"`
	IsTransient          *bool                        `json:"isTransient,omitempty"`
	Name                 *string                      `json:"name,omitempty"`
	ReferenceName        *string                      `json:"referenceName,omitempty"`
	ReferenceNodeId      *string                      `json:"referenceNodeId,omitempty"`
	Visibility           *string                      `json:"visibility,omitempty"`
	AssemblyDisplayData  *BTRootAssemblyDisplayData96 `json:"assemblyDisplayData,omitempty"`
	BtType               *string                      `json:"btType,omitempty"`
	OccurrencesToExclude []BTOccurrence74             `json:"occurrencesToExclude,omitempty"`
	Transform            *BTBSMatrix386               `json:"transform,omitempty"`
}

// NewBTSingleAssemblyReferenceDisplayData1557 instantiates a new BTSingleAssemblyReferenceDisplayData1557 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSingleAssemblyReferenceDisplayData1557() *BTSingleAssemblyReferenceDisplayData1557 {
	this := BTSingleAssemblyReferenceDisplayData1557{}
	return &this
}

// NewBTSingleAssemblyReferenceDisplayData1557WithDefaults instantiates a new BTSingleAssemblyReferenceDisplayData1557 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSingleAssemblyReferenceDisplayData1557WithDefaults() *BTSingleAssemblyReferenceDisplayData1557 {
	this := BTSingleAssemblyReferenceDisplayData1557{}
	return &this
}

// GetContextWorkspaceId returns the ContextWorkspaceId field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetContextWorkspaceId() string {
	if o == nil || o.ContextWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.ContextWorkspaceId
}

// GetContextWorkspaceIdOk returns a tuple with the ContextWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetContextWorkspaceIdOk() (*string, bool) {
	if o == nil || o.ContextWorkspaceId == nil {
		return nil, false
	}
	return o.ContextWorkspaceId, true
}

// HasContextWorkspaceId returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasContextWorkspaceId() bool {
	if o != nil && o.ContextWorkspaceId != nil {
		return true
	}

	return false
}

// SetContextWorkspaceId gets a reference to the given string and assigns it to the ContextWorkspaceId field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetContextWorkspaceId(v string) {
	o.ContextWorkspaceId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetError(v string) {
	o.Error = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsTransient returns the IsTransient field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetIsTransient() bool {
	if o == nil || o.IsTransient == nil {
		var ret bool
		return ret
	}
	return *o.IsTransient
}

// GetIsTransientOk returns a tuple with the IsTransient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetIsTransientOk() (*bool, bool) {
	if o == nil || o.IsTransient == nil {
		return nil, false
	}
	return o.IsTransient, true
}

// HasIsTransient returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasIsTransient() bool {
	if o != nil && o.IsTransient != nil {
		return true
	}

	return false
}

// SetIsTransient gets a reference to the given bool and assigns it to the IsTransient field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetIsTransient(v bool) {
	o.IsTransient = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetName(v string) {
	o.Name = &v
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetReferenceName() string {
	if o == nil || o.ReferenceName == nil {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetReferenceNameOk() (*string, bool) {
	if o == nil || o.ReferenceName == nil {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasReferenceName() bool {
	if o != nil && o.ReferenceName != nil {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetReferenceNodeId returns the ReferenceNodeId field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetReferenceNodeId() string {
	if o == nil || o.ReferenceNodeId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceNodeId
}

// GetReferenceNodeIdOk returns a tuple with the ReferenceNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetReferenceNodeIdOk() (*string, bool) {
	if o == nil || o.ReferenceNodeId == nil {
		return nil, false
	}
	return o.ReferenceNodeId, true
}

// HasReferenceNodeId returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasReferenceNodeId() bool {
	if o != nil && o.ReferenceNodeId != nil {
		return true
	}

	return false
}

// SetReferenceNodeId gets a reference to the given string and assigns it to the ReferenceNodeId field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetReferenceNodeId(v string) {
	o.ReferenceNodeId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetVisibility(v string) {
	o.Visibility = &v
}

// GetAssemblyDisplayData returns the AssemblyDisplayData field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetAssemblyDisplayData() BTRootAssemblyDisplayData96 {
	if o == nil || o.AssemblyDisplayData == nil {
		var ret BTRootAssemblyDisplayData96
		return ret
	}
	return *o.AssemblyDisplayData
}

// GetAssemblyDisplayDataOk returns a tuple with the AssemblyDisplayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetAssemblyDisplayDataOk() (*BTRootAssemblyDisplayData96, bool) {
	if o == nil || o.AssemblyDisplayData == nil {
		return nil, false
	}
	return o.AssemblyDisplayData, true
}

// HasAssemblyDisplayData returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasAssemblyDisplayData() bool {
	if o != nil && o.AssemblyDisplayData != nil {
		return true
	}

	return false
}

// SetAssemblyDisplayData gets a reference to the given BTRootAssemblyDisplayData96 and assigns it to the AssemblyDisplayData field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetAssemblyDisplayData(v BTRootAssemblyDisplayData96) {
	o.AssemblyDisplayData = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetBtType(v string) {
	o.BtType = &v
}

// GetOccurrencesToExclude returns the OccurrencesToExclude field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetOccurrencesToExclude() []BTOccurrence74 {
	if o == nil || o.OccurrencesToExclude == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.OccurrencesToExclude
}

// GetOccurrencesToExcludeOk returns a tuple with the OccurrencesToExclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetOccurrencesToExcludeOk() ([]BTOccurrence74, bool) {
	if o == nil || o.OccurrencesToExclude == nil {
		return nil, false
	}
	return o.OccurrencesToExclude, true
}

// HasOccurrencesToExclude returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasOccurrencesToExclude() bool {
	if o != nil && o.OccurrencesToExclude != nil {
		return true
	}

	return false
}

// SetOccurrencesToExclude gets a reference to the given []BTOccurrence74 and assigns it to the OccurrencesToExclude field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetOccurrencesToExclude(v []BTOccurrence74) {
	o.OccurrencesToExclude = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetTransform() BTBSMatrix386 {
	if o == nil || o.Transform == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) GetTransformOk() (*BTBSMatrix386, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *BTSingleAssemblyReferenceDisplayData1557) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given BTBSMatrix386 and assigns it to the Transform field.
func (o *BTSingleAssemblyReferenceDisplayData1557) SetTransform(v BTBSMatrix386) {
	o.Transform = &v
}

func (o BTSingleAssemblyReferenceDisplayData1557) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContextWorkspaceId != nil {
		toSerialize["contextWorkspaceId"] = o.ContextWorkspaceId
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.IsTransient != nil {
		toSerialize["isTransient"] = o.IsTransient
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ReferenceName != nil {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if o.ReferenceNodeId != nil {
		toSerialize["referenceNodeId"] = o.ReferenceNodeId
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.AssemblyDisplayData != nil {
		toSerialize["assemblyDisplayData"] = o.AssemblyDisplayData
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.OccurrencesToExclude != nil {
		toSerialize["occurrencesToExclude"] = o.OccurrencesToExclude
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableBTSingleAssemblyReferenceDisplayData1557 struct {
	value *BTSingleAssemblyReferenceDisplayData1557
	isSet bool
}

func (v NullableBTSingleAssemblyReferenceDisplayData1557) Get() *BTSingleAssemblyReferenceDisplayData1557 {
	return v.value
}

func (v *NullableBTSingleAssemblyReferenceDisplayData1557) Set(val *BTSingleAssemblyReferenceDisplayData1557) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSingleAssemblyReferenceDisplayData1557) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSingleAssemblyReferenceDisplayData1557) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSingleAssemblyReferenceDisplayData1557(val *BTSingleAssemblyReferenceDisplayData1557) *NullableBTSingleAssemblyReferenceDisplayData1557 {
	return &NullableBTSingleAssemblyReferenceDisplayData1557{value: val, isSet: true}
}

func (v NullableBTSingleAssemblyReferenceDisplayData1557) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSingleAssemblyReferenceDisplayData1557) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
