/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14806-89d807e7089c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPPredicateDeclaration265 struct for BTPPredicateDeclaration265
type BTPPredicateDeclaration265 struct {
	Atomic                *bool                       `json:"atomic,omitempty"`
	BtType                *string                     `json:"btType,omitempty"`
	DocumentationType     *string                     `json:"documentationType,omitempty"`
	EndSourceLocation     *int32                      `json:"endSourceLocation,omitempty"`
	NodeId                *string                     `json:"nodeId,omitempty"`
	ShortDescriptor       *string                     `json:"shortDescriptor,omitempty"`
	SpaceAfter            *BTPSpace10                 `json:"spaceAfter,omitempty"`
	SpaceBefore           *BTPSpace10                 `json:"spaceBefore,omitempty"`
	SpaceDefault          *bool                       `json:"spaceDefault,omitempty"`
	StartSourceLocation   *int32                      `json:"startSourceLocation,omitempty"`
	Annotation            *BTPAnnotation231           `json:"annotation,omitempty"`
	ArgumentsToDocument   []BTPArgumentDeclaration232 `json:"argumentsToDocument,omitempty"`
	Deprecated            *bool                       `json:"deprecated,omitempty"`
	DeprecatedExplanation *string                     `json:"deprecatedExplanation,omitempty"`
	ForExport             *bool                       `json:"forExport,omitempty"`
	SpaceAfterExport      *BTPSpace10                 `json:"spaceAfterExport,omitempty"`
	SymbolName            *BTPIdentifier8             `json:"symbolName,omitempty"`
	Arguments             []BTPArgumentDeclaration232 `json:"arguments,omitempty"`
	Body                  *BTPStatementBlock271       `json:"body,omitempty"`
	Precondition          *BTPStatement269            `json:"precondition,omitempty"`
	ReturnType            *BTPTypeName290             `json:"returnType,omitempty"`
	SpaceAfterArglist     *BTPSpace10                 `json:"spaceAfterArglist,omitempty"`
	SpaceInEmptyList      *BTPSpace10                 `json:"spaceInEmptyList,omitempty"`
	Name                  *BTPIdentifier8             `json:"name,omitempty"`
}

// NewBTPPredicateDeclaration265 instantiates a new BTPPredicateDeclaration265 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPPredicateDeclaration265() *BTPPredicateDeclaration265 {
	this := BTPPredicateDeclaration265{}
	return &this
}

// NewBTPPredicateDeclaration265WithDefaults instantiates a new BTPPredicateDeclaration265 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPPredicateDeclaration265WithDefaults() *BTPPredicateDeclaration265 {
	this := BTPPredicateDeclaration265{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPPredicateDeclaration265) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPPredicateDeclaration265) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetDocumentationType() string {
	if o == nil || o.DocumentationType == nil {
		var ret string
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetDocumentationTypeOk() (*string, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *BTPPredicateDeclaration265) SetDocumentationType(v string) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPPredicateDeclaration265) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPPredicateDeclaration265) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPPredicateDeclaration265) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPPredicateDeclaration265) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPPredicateDeclaration265) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPPredicateDeclaration265) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPPredicateDeclaration265) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPPredicateDeclaration265) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetArgumentsToDocument returns the ArgumentsToDocument field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetArgumentsToDocument() []BTPArgumentDeclaration232 {
	if o == nil || o.ArgumentsToDocument == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return o.ArgumentsToDocument
}

// GetArgumentsToDocumentOk returns a tuple with the ArgumentsToDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetArgumentsToDocumentOk() ([]BTPArgumentDeclaration232, bool) {
	if o == nil || o.ArgumentsToDocument == nil {
		return nil, false
	}
	return o.ArgumentsToDocument, true
}

// HasArgumentsToDocument returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasArgumentsToDocument() bool {
	if o != nil && o.ArgumentsToDocument != nil {
		return true
	}

	return false
}

// SetArgumentsToDocument gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the ArgumentsToDocument field.
func (o *BTPPredicateDeclaration265) SetArgumentsToDocument(v []BTPArgumentDeclaration232) {
	o.ArgumentsToDocument = v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *BTPPredicateDeclaration265) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecatedExplanation returns the DeprecatedExplanation field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetDeprecatedExplanation() string {
	if o == nil || o.DeprecatedExplanation == nil {
		var ret string
		return ret
	}
	return *o.DeprecatedExplanation
}

// GetDeprecatedExplanationOk returns a tuple with the DeprecatedExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetDeprecatedExplanationOk() (*string, bool) {
	if o == nil || o.DeprecatedExplanation == nil {
		return nil, false
	}
	return o.DeprecatedExplanation, true
}

// HasDeprecatedExplanation returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasDeprecatedExplanation() bool {
	if o != nil && o.DeprecatedExplanation != nil {
		return true
	}

	return false
}

// SetDeprecatedExplanation gets a reference to the given string and assigns it to the DeprecatedExplanation field.
func (o *BTPPredicateDeclaration265) SetDeprecatedExplanation(v string) {
	o.DeprecatedExplanation = &v
}

// GetForExport returns the ForExport field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetForExport() bool {
	if o == nil || o.ForExport == nil {
		var ret bool
		return ret
	}
	return *o.ForExport
}

// GetForExportOk returns a tuple with the ForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetForExportOk() (*bool, bool) {
	if o == nil || o.ForExport == nil {
		return nil, false
	}
	return o.ForExport, true
}

// HasForExport returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasForExport() bool {
	if o != nil && o.ForExport != nil {
		return true
	}

	return false
}

// SetForExport gets a reference to the given bool and assigns it to the ForExport field.
func (o *BTPPredicateDeclaration265) SetForExport(v bool) {
	o.ForExport = &v
}

// GetSpaceAfterExport returns the SpaceAfterExport field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSpaceAfterExport() BTPSpace10 {
	if o == nil || o.SpaceAfterExport == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterExport
}

// GetSpaceAfterExportOk returns a tuple with the SpaceAfterExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSpaceAfterExportOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterExport == nil {
		return nil, false
	}
	return o.SpaceAfterExport, true
}

// HasSpaceAfterExport returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSpaceAfterExport() bool {
	if o != nil && o.SpaceAfterExport != nil {
		return true
	}

	return false
}

// SetSpaceAfterExport gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterExport field.
func (o *BTPPredicateDeclaration265) SetSpaceAfterExport(v BTPSpace10) {
	o.SpaceAfterExport = &v
}

// GetSymbolName returns the SymbolName field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSymbolName() BTPIdentifier8 {
	if o == nil || o.SymbolName == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.SymbolName
}

// GetSymbolNameOk returns a tuple with the SymbolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSymbolNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.SymbolName == nil {
		return nil, false
	}
	return o.SymbolName, true
}

// HasSymbolName returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSymbolName() bool {
	if o != nil && o.SymbolName != nil {
		return true
	}

	return false
}

// SetSymbolName gets a reference to the given BTPIdentifier8 and assigns it to the SymbolName field.
func (o *BTPPredicateDeclaration265) SetSymbolName(v BTPIdentifier8) {
	o.SymbolName = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetArguments() []BTPArgumentDeclaration232 {
	if o == nil || o.Arguments == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetArgumentsOk() ([]BTPArgumentDeclaration232, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the Arguments field.
func (o *BTPPredicateDeclaration265) SetArguments(v []BTPArgumentDeclaration232) {
	o.Arguments = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetBody() BTPStatementBlock271 {
	if o == nil || o.Body == nil {
		var ret BTPStatementBlock271
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetBodyOk() (*BTPStatementBlock271, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatementBlock271 and assigns it to the Body field.
func (o *BTPPredicateDeclaration265) SetBody(v BTPStatementBlock271) {
	o.Body = &v
}

// GetPrecondition returns the Precondition field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetPrecondition() BTPStatement269 {
	if o == nil || o.Precondition == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Precondition
}

// GetPreconditionOk returns a tuple with the Precondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetPreconditionOk() (*BTPStatement269, bool) {
	if o == nil || o.Precondition == nil {
		return nil, false
	}
	return o.Precondition, true
}

// HasPrecondition returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasPrecondition() bool {
	if o != nil && o.Precondition != nil {
		return true
	}

	return false
}

// SetPrecondition gets a reference to the given BTPStatement269 and assigns it to the Precondition field.
func (o *BTPPredicateDeclaration265) SetPrecondition(v BTPStatement269) {
	o.Precondition = &v
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetReturnType() BTPTypeName290 {
	if o == nil || o.ReturnType == nil {
		var ret BTPTypeName290
		return ret
	}
	return *o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetReturnTypeOk() (*BTPTypeName290, bool) {
	if o == nil || o.ReturnType == nil {
		return nil, false
	}
	return o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasReturnType() bool {
	if o != nil && o.ReturnType != nil {
		return true
	}

	return false
}

// SetReturnType gets a reference to the given BTPTypeName290 and assigns it to the ReturnType field.
func (o *BTPPredicateDeclaration265) SetReturnType(v BTPTypeName290) {
	o.ReturnType = &v
}

// GetSpaceAfterArglist returns the SpaceAfterArglist field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSpaceAfterArglist() BTPSpace10 {
	if o == nil || o.SpaceAfterArglist == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterArglist
}

// GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSpaceAfterArglistOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterArglist == nil {
		return nil, false
	}
	return o.SpaceAfterArglist, true
}

// HasSpaceAfterArglist returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSpaceAfterArglist() bool {
	if o != nil && o.SpaceAfterArglist != nil {
		return true
	}

	return false
}

// SetSpaceAfterArglist gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterArglist field.
func (o *BTPPredicateDeclaration265) SetSpaceAfterArglist(v BTPSpace10) {
	o.SpaceAfterArglist = &v
}

// GetSpaceInEmptyList returns the SpaceInEmptyList field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetSpaceInEmptyList() BTPSpace10 {
	if o == nil || o.SpaceInEmptyList == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInEmptyList
}

// GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetSpaceInEmptyListOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInEmptyList == nil {
		return nil, false
	}
	return o.SpaceInEmptyList, true
}

// HasSpaceInEmptyList returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasSpaceInEmptyList() bool {
	if o != nil && o.SpaceInEmptyList != nil {
		return true
	}

	return false
}

// SetSpaceInEmptyList gets a reference to the given BTPSpace10 and assigns it to the SpaceInEmptyList field.
func (o *BTPPredicateDeclaration265) SetSpaceInEmptyList(v BTPSpace10) {
	o.SpaceInEmptyList = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetName() BTPIdentifier8 {
	if o == nil || o.Name == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPIdentifier8 and assigns it to the Name field.
func (o *BTPPredicateDeclaration265) SetName(v BTPIdentifier8) {
	o.Name = &v
}

func (o BTPPredicateDeclaration265) MarshalJSON() ([]byte, error) {
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
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.Annotation != nil {
		toSerialize["annotation"] = o.Annotation
	}
	if o.ArgumentsToDocument != nil {
		toSerialize["argumentsToDocument"] = o.ArgumentsToDocument
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.DeprecatedExplanation != nil {
		toSerialize["deprecatedExplanation"] = o.DeprecatedExplanation
	}
	if o.ForExport != nil {
		toSerialize["forExport"] = o.ForExport
	}
	if o.SpaceAfterExport != nil {
		toSerialize["spaceAfterExport"] = o.SpaceAfterExport
	}
	if o.SymbolName != nil {
		toSerialize["symbolName"] = o.SymbolName
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Precondition != nil {
		toSerialize["precondition"] = o.Precondition
	}
	if o.ReturnType != nil {
		toSerialize["returnType"] = o.ReturnType
	}
	if o.SpaceAfterArglist != nil {
		toSerialize["spaceAfterArglist"] = o.SpaceAfterArglist
	}
	if o.SpaceInEmptyList != nil {
		toSerialize["spaceInEmptyList"] = o.SpaceInEmptyList
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTPPredicateDeclaration265 struct {
	value *BTPPredicateDeclaration265
	isSet bool
}

func (v NullableBTPPredicateDeclaration265) Get() *BTPPredicateDeclaration265 {
	return v.value
}

func (v *NullableBTPPredicateDeclaration265) Set(val *BTPPredicateDeclaration265) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPPredicateDeclaration265) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPPredicateDeclaration265) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPPredicateDeclaration265(val *BTPPredicateDeclaration265) *NullableBTPPredicateDeclaration265 {
	return &NullableBTPPredicateDeclaration265{value: val, isSet: true}
}

func (v NullableBTPPredicateDeclaration265) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPPredicateDeclaration265) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
