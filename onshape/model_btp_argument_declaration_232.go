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

// BTPArgumentDeclaration232 struct for BTPArgumentDeclaration232
type BTPArgumentDeclaration232 struct {
	Atomic              *bool           `json:"atomic,omitempty"`
	BtType              *string         `json:"btType,omitempty"`
	DocumentationType   *string         `json:"documentationType,omitempty"`
	EndSourceLocation   *int32          `json:"endSourceLocation,omitempty"`
	Identifier          *BTPIdentifier8 `json:"identifier,omitempty"`
	Name                *BTPIdentifier8 `json:"name,omitempty"`
	NodeId              *string         `json:"nodeId,omitempty"`
	ShortDescriptor     *string         `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10     `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10     `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool           `json:"spaceDefault,omitempty"`
	StandardType        *string         `json:"standardType,omitempty"`
	StartSourceLocation *int32          `json:"startSourceLocation,omitempty"`
	Type                *BTPTypeName290 `json:"type,omitempty"`
	TypeName            *string         `json:"typeName,omitempty"`
}

// NewBTPArgumentDeclaration232 instantiates a new BTPArgumentDeclaration232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPArgumentDeclaration232() *BTPArgumentDeclaration232 {
	this := BTPArgumentDeclaration232{}
	return &this
}

// NewBTPArgumentDeclaration232WithDefaults instantiates a new BTPArgumentDeclaration232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPArgumentDeclaration232WithDefaults() *BTPArgumentDeclaration232 {
	this := BTPArgumentDeclaration232{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPArgumentDeclaration232) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPArgumentDeclaration232) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetDocumentationType() string {
	if o == nil || o.DocumentationType == nil {
		var ret string
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetDocumentationTypeOk() (*string, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *BTPArgumentDeclaration232) SetDocumentationType(v string) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPArgumentDeclaration232) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetIdentifier() BTPIdentifier8 {
	if o == nil || o.Identifier == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetIdentifierOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given BTPIdentifier8 and assigns it to the Identifier field.
func (o *BTPArgumentDeclaration232) SetIdentifier(v BTPIdentifier8) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetName() BTPIdentifier8 {
	if o == nil || o.Name == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPIdentifier8 and assigns it to the Name field.
func (o *BTPArgumentDeclaration232) SetName(v BTPIdentifier8) {
	o.Name = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPArgumentDeclaration232) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPArgumentDeclaration232) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPArgumentDeclaration232) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPArgumentDeclaration232) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPArgumentDeclaration232) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStandardType returns the StandardType field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetStandardType() string {
	if o == nil || o.StandardType == nil {
		var ret string
		return ret
	}
	return *o.StandardType
}

// GetStandardTypeOk returns a tuple with the StandardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetStandardTypeOk() (*string, bool) {
	if o == nil || o.StandardType == nil {
		return nil, false
	}
	return o.StandardType, true
}

// HasStandardType returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasStandardType() bool {
	if o != nil && o.StandardType != nil {
		return true
	}

	return false
}

// SetStandardType gets a reference to the given string and assigns it to the StandardType field.
func (o *BTPArgumentDeclaration232) SetStandardType(v string) {
	o.StandardType = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPArgumentDeclaration232) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetType() BTPTypeName290 {
	if o == nil || o.Type == nil {
		var ret BTPTypeName290
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetTypeOk() (*BTPTypeName290, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given BTPTypeName290 and assigns it to the Type field.
func (o *BTPArgumentDeclaration232) SetType(v BTPTypeName290) {
	o.Type = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *BTPArgumentDeclaration232) GetTypeName() string {
	if o == nil || o.TypeName == nil {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPArgumentDeclaration232) GetTypeNameOk() (*string, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BTPArgumentDeclaration232) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *BTPArgumentDeclaration232) SetTypeName(v string) {
	o.TypeName = &v
}

func (o BTPArgumentDeclaration232) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentationType != nil {
		toSerialize["documentationType"] = o.DocumentationType
	}
	if o.EndSourceLocation != nil {
		toSerialize["endSourceLocation"] = o.EndSourceLocation
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ShortDescriptor != nil {
		toSerialize["shortDescriptor"] = o.ShortDescriptor
	}
	if o.SpaceAfter != nil {
		toSerialize["spaceAfter"] = o.SpaceAfter
	}
	if o.SpaceBefore != nil {
		toSerialize["spaceBefore"] = o.SpaceBefore
	}
	if o.SpaceDefault != nil {
		toSerialize["spaceDefault"] = o.SpaceDefault
	}
	if o.StandardType != nil {
		toSerialize["standardType"] = o.StandardType
	}
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableBTPArgumentDeclaration232 struct {
	value *BTPArgumentDeclaration232
	isSet bool
}

func (v NullableBTPArgumentDeclaration232) Get() *BTPArgumentDeclaration232 {
	return v.value
}

func (v *NullableBTPArgumentDeclaration232) Set(val *BTPArgumentDeclaration232) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPArgumentDeclaration232) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPArgumentDeclaration232) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPArgumentDeclaration232(val *BTPArgumentDeclaration232) *NullableBTPArgumentDeclaration232 {
	return &NullableBTPArgumentDeclaration232{value: val, isSet: true}
}

func (v NullableBTPArgumentDeclaration232) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPArgumentDeclaration232) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
