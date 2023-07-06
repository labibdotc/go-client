/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18417-1bd990c6fbaa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTObjectWorkflowInfo An workflowable object like Release or Task that supports states and transitions.
type BTObjectWorkflowInfo struct {
	// Whether workflowable object can be discarded.
	CanBeDiscarded *bool `json:"canBeDiscarded,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Whether workflowable object has been discarded.
	IsDiscarded *bool `json:"isDiscarded,omitempty"`
	// Whether workflowable object has reached terminal state and is frozen.
	IsFrozen      *bool                `json:"isFrozen,omitempty"`
	MetadataState *BTMetadataStateType `json:"metadataState,omitempty"`
	// Name of the resource.
	Name       *string                `json:"name,omitempty"`
	ObjectType *BTAPIWorkflowableType `json:"objectType,omitempty"`
	// The current state of object like SETUP, REJECTED etc. Custom workflows can have any declared state.
	StateId *string `json:"stateId,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
	// The workflow definition id that governs this object's states and transitions.
	WorkflowId *string `json:"workflowId,omitempty"`
}

// NewBTObjectWorkflowInfo instantiates a new BTObjectWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTObjectWorkflowInfo() *BTObjectWorkflowInfo {
	this := BTObjectWorkflowInfo{}
	return &this
}

// NewBTObjectWorkflowInfoWithDefaults instantiates a new BTObjectWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTObjectWorkflowInfoWithDefaults() *BTObjectWorkflowInfo {
	this := BTObjectWorkflowInfo{}
	return &this
}

// GetCanBeDiscarded returns the CanBeDiscarded field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetCanBeDiscarded() bool {
	if o == nil || o.CanBeDiscarded == nil {
		var ret bool
		return ret
	}
	return *o.CanBeDiscarded
}

// GetCanBeDiscardedOk returns a tuple with the CanBeDiscarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetCanBeDiscardedOk() (*bool, bool) {
	if o == nil || o.CanBeDiscarded == nil {
		return nil, false
	}
	return o.CanBeDiscarded, true
}

// HasCanBeDiscarded returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasCanBeDiscarded() bool {
	if o != nil && o.CanBeDiscarded != nil {
		return true
	}

	return false
}

// SetCanBeDiscarded gets a reference to the given bool and assigns it to the CanBeDiscarded field.
func (o *BTObjectWorkflowInfo) SetCanBeDiscarded(v bool) {
	o.CanBeDiscarded = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTObjectWorkflowInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTObjectWorkflowInfo) SetId(v string) {
	o.Id = &v
}

// GetIsDiscarded returns the IsDiscarded field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetIsDiscarded() bool {
	if o == nil || o.IsDiscarded == nil {
		var ret bool
		return ret
	}
	return *o.IsDiscarded
}

// GetIsDiscardedOk returns a tuple with the IsDiscarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetIsDiscardedOk() (*bool, bool) {
	if o == nil || o.IsDiscarded == nil {
		return nil, false
	}
	return o.IsDiscarded, true
}

// HasIsDiscarded returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasIsDiscarded() bool {
	if o != nil && o.IsDiscarded != nil {
		return true
	}

	return false
}

// SetIsDiscarded gets a reference to the given bool and assigns it to the IsDiscarded field.
func (o *BTObjectWorkflowInfo) SetIsDiscarded(v bool) {
	o.IsDiscarded = &v
}

// GetIsFrozen returns the IsFrozen field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetIsFrozen() bool {
	if o == nil || o.IsFrozen == nil {
		var ret bool
		return ret
	}
	return *o.IsFrozen
}

// GetIsFrozenOk returns a tuple with the IsFrozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetIsFrozenOk() (*bool, bool) {
	if o == nil || o.IsFrozen == nil {
		return nil, false
	}
	return o.IsFrozen, true
}

// HasIsFrozen returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasIsFrozen() bool {
	if o != nil && o.IsFrozen != nil {
		return true
	}

	return false
}

// SetIsFrozen gets a reference to the given bool and assigns it to the IsFrozen field.
func (o *BTObjectWorkflowInfo) SetIsFrozen(v bool) {
	o.IsFrozen = &v
}

// GetMetadataState returns the MetadataState field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetMetadataState() BTMetadataStateType {
	if o == nil || o.MetadataState == nil {
		var ret BTMetadataStateType
		return ret
	}
	return *o.MetadataState
}

// GetMetadataStateOk returns a tuple with the MetadataState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetMetadataStateOk() (*BTMetadataStateType, bool) {
	if o == nil || o.MetadataState == nil {
		return nil, false
	}
	return o.MetadataState, true
}

// HasMetadataState returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasMetadataState() bool {
	if o != nil && o.MetadataState != nil {
		return true
	}

	return false
}

// SetMetadataState gets a reference to the given BTMetadataStateType and assigns it to the MetadataState field.
func (o *BTObjectWorkflowInfo) SetMetadataState(v BTMetadataStateType) {
	o.MetadataState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTObjectWorkflowInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetObjectType() BTAPIWorkflowableType {
	if o == nil || o.ObjectType == nil {
		var ret BTAPIWorkflowableType
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetObjectTypeOk() (*BTAPIWorkflowableType, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given BTAPIWorkflowableType and assigns it to the ObjectType field.
func (o *BTObjectWorkflowInfo) SetObjectType(v BTAPIWorkflowableType) {
	o.ObjectType = &v
}

// GetStateId returns the StateId field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetStateId() string {
	if o == nil || o.StateId == nil {
		var ret string
		return ret
	}
	return *o.StateId
}

// GetStateIdOk returns a tuple with the StateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetStateIdOk() (*string, bool) {
	if o == nil || o.StateId == nil {
		return nil, false
	}
	return o.StateId, true
}

// HasStateId returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasStateId() bool {
	if o != nil && o.StateId != nil {
		return true
	}

	return false
}

// SetStateId gets a reference to the given string and assigns it to the StateId field.
func (o *BTObjectWorkflowInfo) SetStateId(v string) {
	o.StateId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTObjectWorkflowInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *BTObjectWorkflowInfo) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTObjectWorkflowInfo) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *BTObjectWorkflowInfo) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *BTObjectWorkflowInfo) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

func (o BTObjectWorkflowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanBeDiscarded != nil {
		toSerialize["canBeDiscarded"] = o.CanBeDiscarded
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDiscarded != nil {
		toSerialize["isDiscarded"] = o.IsDiscarded
	}
	if o.IsFrozen != nil {
		toSerialize["isFrozen"] = o.IsFrozen
	}
	if o.MetadataState != nil {
		toSerialize["metadataState"] = o.MetadataState
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.StateId != nil {
		toSerialize["stateId"] = o.StateId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	return json.Marshal(toSerialize)
}

type NullableBTObjectWorkflowInfo struct {
	value *BTObjectWorkflowInfo
	isSet bool
}

func (v NullableBTObjectWorkflowInfo) Get() *BTObjectWorkflowInfo {
	return v.value
}

func (v *NullableBTObjectWorkflowInfo) Set(val *BTObjectWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTObjectWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTObjectWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTObjectWorkflowInfo(val *BTObjectWorkflowInfo) *NullableBTObjectWorkflowInfo {
	return &NullableBTObjectWorkflowInfo{value: val, isSet: true}
}

func (v NullableBTObjectWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTObjectWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
