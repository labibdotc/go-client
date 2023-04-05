/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13995-cdd961a1a6ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRootDiffInfo struct for BTRootDiffInfo
type BTRootDiffInfo struct {
	Changes                *map[string]BTDiffInfo   `json:"changes,omitempty"`
	CollectionChanges      *map[string][]BTDiffInfo `json:"collectionChanges,omitempty"`
	EntityType             *string                  `json:"entityType,omitempty"`
	GeometryChangeMessages []string                 `json:"geometryChangeMessages,omitempty"`
	SourceConfiguration    *string                  `json:"sourceConfiguration,omitempty"`
	SourceId               *string                  `json:"sourceId,omitempty"`
	SourceMicroversionId   *string                  `json:"sourceMicroversionId,omitempty"`
	SourceValue            *string                  `json:"sourceValue,omitempty"`
	SourceVersionId        *string                  `json:"sourceVersionId,omitempty"`
	SourceWorkspaceId      *string                  `json:"sourceWorkspaceId,omitempty"`
	TargetConfiguration    *string                  `json:"targetConfiguration,omitempty"`
	TargetId               *string                  `json:"targetId,omitempty"`
	TargetMicroversionId   *string                  `json:"targetMicroversionId,omitempty"`
	TargetValue            *string                  `json:"targetValue,omitempty"`
	TargetVersionId        *string                  `json:"targetVersionId,omitempty"`
	TargetWorkspaceId      *string                  `json:"targetWorkspaceId,omitempty"`
	Type                   *string                  `json:"type,omitempty"`
}

// NewBTRootDiffInfo instantiates a new BTRootDiffInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRootDiffInfo() *BTRootDiffInfo {
	this := BTRootDiffInfo{}
	return &this
}

// NewBTRootDiffInfoWithDefaults instantiates a new BTRootDiffInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRootDiffInfoWithDefaults() *BTRootDiffInfo {
	this := BTRootDiffInfo{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetChanges() map[string]BTDiffInfo {
	if o == nil || o.Changes == nil {
		var ret map[string]BTDiffInfo
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetChangesOk() (*map[string]BTDiffInfo, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given map[string]BTDiffInfo and assigns it to the Changes field.
func (o *BTRootDiffInfo) SetChanges(v map[string]BTDiffInfo) {
	o.Changes = &v
}

// GetCollectionChanges returns the CollectionChanges field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetCollectionChanges() map[string][]BTDiffInfo {
	if o == nil || o.CollectionChanges == nil {
		var ret map[string][]BTDiffInfo
		return ret
	}
	return *o.CollectionChanges
}

// GetCollectionChangesOk returns a tuple with the CollectionChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetCollectionChangesOk() (*map[string][]BTDiffInfo, bool) {
	if o == nil || o.CollectionChanges == nil {
		return nil, false
	}
	return o.CollectionChanges, true
}

// HasCollectionChanges returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasCollectionChanges() bool {
	if o != nil && o.CollectionChanges != nil {
		return true
	}

	return false
}

// SetCollectionChanges gets a reference to the given map[string][]BTDiffInfo and assigns it to the CollectionChanges field.
func (o *BTRootDiffInfo) SetCollectionChanges(v map[string][]BTDiffInfo) {
	o.CollectionChanges = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *BTRootDiffInfo) SetEntityType(v string) {
	o.EntityType = &v
}

// GetGeometryChangeMessages returns the GeometryChangeMessages field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetGeometryChangeMessages() []string {
	if o == nil || o.GeometryChangeMessages == nil {
		var ret []string
		return ret
	}
	return o.GeometryChangeMessages
}

// GetGeometryChangeMessagesOk returns a tuple with the GeometryChangeMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetGeometryChangeMessagesOk() ([]string, bool) {
	if o == nil || o.GeometryChangeMessages == nil {
		return nil, false
	}
	return o.GeometryChangeMessages, true
}

// HasGeometryChangeMessages returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasGeometryChangeMessages() bool {
	if o != nil && o.GeometryChangeMessages != nil {
		return true
	}

	return false
}

// SetGeometryChangeMessages gets a reference to the given []string and assigns it to the GeometryChangeMessages field.
func (o *BTRootDiffInfo) SetGeometryChangeMessages(v []string) {
	o.GeometryChangeMessages = v
}

// GetSourceConfiguration returns the SourceConfiguration field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetSourceConfiguration() string {
	if o == nil || o.SourceConfiguration == nil {
		var ret string
		return ret
	}
	return *o.SourceConfiguration
}

// GetSourceConfigurationOk returns a tuple with the SourceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetSourceConfigurationOk() (*string, bool) {
	if o == nil || o.SourceConfiguration == nil {
		return nil, false
	}
	return o.SourceConfiguration, true
}

// HasSourceConfiguration returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasSourceConfiguration() bool {
	if o != nil && o.SourceConfiguration != nil {
		return true
	}

	return false
}

// SetSourceConfiguration gets a reference to the given string and assigns it to the SourceConfiguration field.
func (o *BTRootDiffInfo) SetSourceConfiguration(v string) {
	o.SourceConfiguration = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *BTRootDiffInfo) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceMicroversionId returns the SourceMicroversionId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetSourceMicroversionId() string {
	if o == nil || o.SourceMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversionId
}

// GetSourceMicroversionIdOk returns a tuple with the SourceMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetSourceMicroversionIdOk() (*string, bool) {
	if o == nil || o.SourceMicroversionId == nil {
		return nil, false
	}
	return o.SourceMicroversionId, true
}

// HasSourceMicroversionId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasSourceMicroversionId() bool {
	if o != nil && o.SourceMicroversionId != nil {
		return true
	}

	return false
}

// SetSourceMicroversionId gets a reference to the given string and assigns it to the SourceMicroversionId field.
func (o *BTRootDiffInfo) SetSourceMicroversionId(v string) {
	o.SourceMicroversionId = &v
}

// GetSourceValue returns the SourceValue field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetSourceValue() string {
	if o == nil || o.SourceValue == nil {
		var ret string
		return ret
	}
	return *o.SourceValue
}

// GetSourceValueOk returns a tuple with the SourceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetSourceValueOk() (*string, bool) {
	if o == nil || o.SourceValue == nil {
		return nil, false
	}
	return o.SourceValue, true
}

// HasSourceValue returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasSourceValue() bool {
	if o != nil && o.SourceValue != nil {
		return true
	}

	return false
}

// SetSourceValue gets a reference to the given string and assigns it to the SourceValue field.
func (o *BTRootDiffInfo) SetSourceValue(v string) {
	o.SourceValue = &v
}

// GetSourceVersionId returns the SourceVersionId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetSourceVersionId() string {
	if o == nil || o.SourceVersionId == nil {
		var ret string
		return ret
	}
	return *o.SourceVersionId
}

// GetSourceVersionIdOk returns a tuple with the SourceVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetSourceVersionIdOk() (*string, bool) {
	if o == nil || o.SourceVersionId == nil {
		return nil, false
	}
	return o.SourceVersionId, true
}

// HasSourceVersionId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasSourceVersionId() bool {
	if o != nil && o.SourceVersionId != nil {
		return true
	}

	return false
}

// SetSourceVersionId gets a reference to the given string and assigns it to the SourceVersionId field.
func (o *BTRootDiffInfo) SetSourceVersionId(v string) {
	o.SourceVersionId = &v
}

// GetSourceWorkspaceId returns the SourceWorkspaceId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetSourceWorkspaceId() string {
	if o == nil || o.SourceWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.SourceWorkspaceId
}

// GetSourceWorkspaceIdOk returns a tuple with the SourceWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetSourceWorkspaceIdOk() (*string, bool) {
	if o == nil || o.SourceWorkspaceId == nil {
		return nil, false
	}
	return o.SourceWorkspaceId, true
}

// HasSourceWorkspaceId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasSourceWorkspaceId() bool {
	if o != nil && o.SourceWorkspaceId != nil {
		return true
	}

	return false
}

// SetSourceWorkspaceId gets a reference to the given string and assigns it to the SourceWorkspaceId field.
func (o *BTRootDiffInfo) SetSourceWorkspaceId(v string) {
	o.SourceWorkspaceId = &v
}

// GetTargetConfiguration returns the TargetConfiguration field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetTargetConfiguration() string {
	if o == nil || o.TargetConfiguration == nil {
		var ret string
		return ret
	}
	return *o.TargetConfiguration
}

// GetTargetConfigurationOk returns a tuple with the TargetConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTargetConfigurationOk() (*string, bool) {
	if o == nil || o.TargetConfiguration == nil {
		return nil, false
	}
	return o.TargetConfiguration, true
}

// HasTargetConfiguration returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasTargetConfiguration() bool {
	if o != nil && o.TargetConfiguration != nil {
		return true
	}

	return false
}

// SetTargetConfiguration gets a reference to the given string and assigns it to the TargetConfiguration field.
func (o *BTRootDiffInfo) SetTargetConfiguration(v string) {
	o.TargetConfiguration = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *BTRootDiffInfo) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetMicroversionId returns the TargetMicroversionId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetTargetMicroversionId() string {
	if o == nil || o.TargetMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.TargetMicroversionId
}

// GetTargetMicroversionIdOk returns a tuple with the TargetMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTargetMicroversionIdOk() (*string, bool) {
	if o == nil || o.TargetMicroversionId == nil {
		return nil, false
	}
	return o.TargetMicroversionId, true
}

// HasTargetMicroversionId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasTargetMicroversionId() bool {
	if o != nil && o.TargetMicroversionId != nil {
		return true
	}

	return false
}

// SetTargetMicroversionId gets a reference to the given string and assigns it to the TargetMicroversionId field.
func (o *BTRootDiffInfo) SetTargetMicroversionId(v string) {
	o.TargetMicroversionId = &v
}

// GetTargetValue returns the TargetValue field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetTargetValue() string {
	if o == nil || o.TargetValue == nil {
		var ret string
		return ret
	}
	return *o.TargetValue
}

// GetTargetValueOk returns a tuple with the TargetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTargetValueOk() (*string, bool) {
	if o == nil || o.TargetValue == nil {
		return nil, false
	}
	return o.TargetValue, true
}

// HasTargetValue returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasTargetValue() bool {
	if o != nil && o.TargetValue != nil {
		return true
	}

	return false
}

// SetTargetValue gets a reference to the given string and assigns it to the TargetValue field.
func (o *BTRootDiffInfo) SetTargetValue(v string) {
	o.TargetValue = &v
}

// GetTargetVersionId returns the TargetVersionId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetTargetVersionId() string {
	if o == nil || o.TargetVersionId == nil {
		var ret string
		return ret
	}
	return *o.TargetVersionId
}

// GetTargetVersionIdOk returns a tuple with the TargetVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTargetVersionIdOk() (*string, bool) {
	if o == nil || o.TargetVersionId == nil {
		return nil, false
	}
	return o.TargetVersionId, true
}

// HasTargetVersionId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasTargetVersionId() bool {
	if o != nil && o.TargetVersionId != nil {
		return true
	}

	return false
}

// SetTargetVersionId gets a reference to the given string and assigns it to the TargetVersionId field.
func (o *BTRootDiffInfo) SetTargetVersionId(v string) {
	o.TargetVersionId = &v
}

// GetTargetWorkspaceId returns the TargetWorkspaceId field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetTargetWorkspaceId() string {
	if o == nil || o.TargetWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.TargetWorkspaceId
}

// GetTargetWorkspaceIdOk returns a tuple with the TargetWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTargetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.TargetWorkspaceId == nil {
		return nil, false
	}
	return o.TargetWorkspaceId, true
}

// HasTargetWorkspaceId returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasTargetWorkspaceId() bool {
	if o != nil && o.TargetWorkspaceId != nil {
		return true
	}

	return false
}

// SetTargetWorkspaceId gets a reference to the given string and assigns it to the TargetWorkspaceId field.
func (o *BTRootDiffInfo) SetTargetWorkspaceId(v string) {
	o.TargetWorkspaceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTRootDiffInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootDiffInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTRootDiffInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTRootDiffInfo) SetType(v string) {
	o.Type = &v
}

func (o BTRootDiffInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.CollectionChanges != nil {
		toSerialize["collectionChanges"] = o.CollectionChanges
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.GeometryChangeMessages != nil {
		toSerialize["geometryChangeMessages"] = o.GeometryChangeMessages
	}
	if o.SourceConfiguration != nil {
		toSerialize["sourceConfiguration"] = o.SourceConfiguration
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.SourceMicroversionId != nil {
		toSerialize["sourceMicroversionId"] = o.SourceMicroversionId
	}
	if o.SourceValue != nil {
		toSerialize["sourceValue"] = o.SourceValue
	}
	if o.SourceVersionId != nil {
		toSerialize["sourceVersionId"] = o.SourceVersionId
	}
	if o.SourceWorkspaceId != nil {
		toSerialize["sourceWorkspaceId"] = o.SourceWorkspaceId
	}
	if o.TargetConfiguration != nil {
		toSerialize["targetConfiguration"] = o.TargetConfiguration
	}
	if o.TargetId != nil {
		toSerialize["targetId"] = o.TargetId
	}
	if o.TargetMicroversionId != nil {
		toSerialize["targetMicroversionId"] = o.TargetMicroversionId
	}
	if o.TargetValue != nil {
		toSerialize["targetValue"] = o.TargetValue
	}
	if o.TargetVersionId != nil {
		toSerialize["targetVersionId"] = o.TargetVersionId
	}
	if o.TargetWorkspaceId != nil {
		toSerialize["targetWorkspaceId"] = o.TargetWorkspaceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTRootDiffInfo struct {
	value *BTRootDiffInfo
	isSet bool
}

func (v NullableBTRootDiffInfo) Get() *BTRootDiffInfo {
	return v.value
}

func (v *NullableBTRootDiffInfo) Set(val *BTRootDiffInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRootDiffInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRootDiffInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRootDiffInfo(val *BTRootDiffInfo) *NullableBTRootDiffInfo {
	return &NullableBTRootDiffInfo{value: val, isSet: true}
}

func (v NullableBTRootDiffInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRootDiffInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
