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

// BTDocumentParams struct for BTDocumentParams
type BTDocumentParams struct {
	BetaCapabilityIds       []string                              `json:"betaCapabilityIds,omitempty"`
	Description             *string                               `json:"description,omitempty"`
	Elements                []BTDocumentElementCreationDescriptor `json:"elements,omitempty"`
	ForceExportRules        *bool                                 `json:"forceExportRules,omitempty"`
	GenerateUnknownMessages *bool                                 `json:"generateUnknownMessages,omitempty"`
	IsEmptyContent          *bool                                 `json:"isEmptyContent,omitempty"`
	IsPublic                *bool                                 `json:"isPublic,omitempty"`
	Name                    *string                               `json:"name,omitempty"`
	NotRevisionManaged      *bool                                 `json:"notRevisionManaged,omitempty"`
	OwnerEmail              *string                               `json:"ownerEmail,omitempty"`
	OwnerId                 *string                               `json:"ownerId,omitempty"`
	OwnerType               *int32                                `json:"ownerType,omitempty"`
	ParentId                *string                               `json:"parentId,omitempty"`
	ProjectId               *string                               `json:"projectId,omitempty"`
	Tags                    []string                              `json:"tags,omitempty"`
}

// NewBTDocumentParams instantiates a new BTDocumentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentParams() *BTDocumentParams {
	this := BTDocumentParams{}
	return &this
}

// NewBTDocumentParamsWithDefaults instantiates a new BTDocumentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentParamsWithDefaults() *BTDocumentParams {
	this := BTDocumentParams{}
	return &this
}

// GetBetaCapabilityIds returns the BetaCapabilityIds field value if set, zero value otherwise.
func (o *BTDocumentParams) GetBetaCapabilityIds() []string {
	if o == nil || o.BetaCapabilityIds == nil {
		var ret []string
		return ret
	}
	return o.BetaCapabilityIds
}

// GetBetaCapabilityIdsOk returns a tuple with the BetaCapabilityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetBetaCapabilityIdsOk() ([]string, bool) {
	if o == nil || o.BetaCapabilityIds == nil {
		return nil, false
	}
	return o.BetaCapabilityIds, true
}

// HasBetaCapabilityIds returns a boolean if a field has been set.
func (o *BTDocumentParams) HasBetaCapabilityIds() bool {
	if o != nil && o.BetaCapabilityIds != nil {
		return true
	}

	return false
}

// SetBetaCapabilityIds gets a reference to the given []string and assigns it to the BetaCapabilityIds field.
func (o *BTDocumentParams) SetBetaCapabilityIds(v []string) {
	o.BetaCapabilityIds = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTDocumentParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTDocumentParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTDocumentParams) SetDescription(v string) {
	o.Description = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTDocumentParams) GetElements() []BTDocumentElementCreationDescriptor {
	if o == nil || o.Elements == nil {
		var ret []BTDocumentElementCreationDescriptor
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetElementsOk() ([]BTDocumentElementCreationDescriptor, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTDocumentParams) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []BTDocumentElementCreationDescriptor and assigns it to the Elements field.
func (o *BTDocumentParams) SetElements(v []BTDocumentElementCreationDescriptor) {
	o.Elements = v
}

// GetForceExportRules returns the ForceExportRules field value if set, zero value otherwise.
func (o *BTDocumentParams) GetForceExportRules() bool {
	if o == nil || o.ForceExportRules == nil {
		var ret bool
		return ret
	}
	return *o.ForceExportRules
}

// GetForceExportRulesOk returns a tuple with the ForceExportRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetForceExportRulesOk() (*bool, bool) {
	if o == nil || o.ForceExportRules == nil {
		return nil, false
	}
	return o.ForceExportRules, true
}

// HasForceExportRules returns a boolean if a field has been set.
func (o *BTDocumentParams) HasForceExportRules() bool {
	if o != nil && o.ForceExportRules != nil {
		return true
	}

	return false
}

// SetForceExportRules gets a reference to the given bool and assigns it to the ForceExportRules field.
func (o *BTDocumentParams) SetForceExportRules(v bool) {
	o.ForceExportRules = &v
}

// GetGenerateUnknownMessages returns the GenerateUnknownMessages field value if set, zero value otherwise.
func (o *BTDocumentParams) GetGenerateUnknownMessages() bool {
	if o == nil || o.GenerateUnknownMessages == nil {
		var ret bool
		return ret
	}
	return *o.GenerateUnknownMessages
}

// GetGenerateUnknownMessagesOk returns a tuple with the GenerateUnknownMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetGenerateUnknownMessagesOk() (*bool, bool) {
	if o == nil || o.GenerateUnknownMessages == nil {
		return nil, false
	}
	return o.GenerateUnknownMessages, true
}

// HasGenerateUnknownMessages returns a boolean if a field has been set.
func (o *BTDocumentParams) HasGenerateUnknownMessages() bool {
	if o != nil && o.GenerateUnknownMessages != nil {
		return true
	}

	return false
}

// SetGenerateUnknownMessages gets a reference to the given bool and assigns it to the GenerateUnknownMessages field.
func (o *BTDocumentParams) SetGenerateUnknownMessages(v bool) {
	o.GenerateUnknownMessages = &v
}

// GetIsEmptyContent returns the IsEmptyContent field value if set, zero value otherwise.
func (o *BTDocumentParams) GetIsEmptyContent() bool {
	if o == nil || o.IsEmptyContent == nil {
		var ret bool
		return ret
	}
	return *o.IsEmptyContent
}

// GetIsEmptyContentOk returns a tuple with the IsEmptyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetIsEmptyContentOk() (*bool, bool) {
	if o == nil || o.IsEmptyContent == nil {
		return nil, false
	}
	return o.IsEmptyContent, true
}

// HasIsEmptyContent returns a boolean if a field has been set.
func (o *BTDocumentParams) HasIsEmptyContent() bool {
	if o != nil && o.IsEmptyContent != nil {
		return true
	}

	return false
}

// SetIsEmptyContent gets a reference to the given bool and assigns it to the IsEmptyContent field.
func (o *BTDocumentParams) SetIsEmptyContent(v bool) {
	o.IsEmptyContent = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *BTDocumentParams) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *BTDocumentParams) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *BTDocumentParams) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDocumentParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDocumentParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDocumentParams) SetName(v string) {
	o.Name = &v
}

// GetNotRevisionManaged returns the NotRevisionManaged field value if set, zero value otherwise.
func (o *BTDocumentParams) GetNotRevisionManaged() bool {
	if o == nil || o.NotRevisionManaged == nil {
		var ret bool
		return ret
	}
	return *o.NotRevisionManaged
}

// GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetNotRevisionManagedOk() (*bool, bool) {
	if o == nil || o.NotRevisionManaged == nil {
		return nil, false
	}
	return o.NotRevisionManaged, true
}

// HasNotRevisionManaged returns a boolean if a field has been set.
func (o *BTDocumentParams) HasNotRevisionManaged() bool {
	if o != nil && o.NotRevisionManaged != nil {
		return true
	}

	return false
}

// SetNotRevisionManaged gets a reference to the given bool and assigns it to the NotRevisionManaged field.
func (o *BTDocumentParams) SetNotRevisionManaged(v bool) {
	o.NotRevisionManaged = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *BTDocumentParams) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *BTDocumentParams) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *BTDocumentParams) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTDocumentParams) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTDocumentParams) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTDocumentParams) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTDocumentParams) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTDocumentParams) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTDocumentParams) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTDocumentParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTDocumentParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTDocumentParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTDocumentParams) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTDocumentParams) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTDocumentParams) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BTDocumentParams) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BTDocumentParams) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BTDocumentParams) SetTags(v []string) {
	o.Tags = v
}

func (o BTDocumentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BetaCapabilityIds != nil {
		toSerialize["betaCapabilityIds"] = o.BetaCapabilityIds
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.ForceExportRules != nil {
		toSerialize["forceExportRules"] = o.ForceExportRules
	}
	if o.GenerateUnknownMessages != nil {
		toSerialize["generateUnknownMessages"] = o.GenerateUnknownMessages
	}
	if o.IsEmptyContent != nil {
		toSerialize["isEmptyContent"] = o.IsEmptyContent
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NotRevisionManaged != nil {
		toSerialize["notRevisionManaged"] = o.NotRevisionManaged
	}
	if o.OwnerEmail != nil {
		toSerialize["ownerEmail"] = o.OwnerEmail
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentParams struct {
	value *BTDocumentParams
	isSet bool
}

func (v NullableBTDocumentParams) Get() *BTDocumentParams {
	return v.value
}

func (v *NullableBTDocumentParams) Set(val *BTDocumentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentParams(val *BTDocumentParams) *NullableBTDocumentParams {
	return &NullableBTDocumentParams{value: val, isSet: true}
}

func (v NullableBTDocumentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
