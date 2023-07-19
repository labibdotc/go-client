/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.19032-0b307c4b0d0e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPTopLevelImport285 struct for BTPTopLevelImport285
type BTPTopLevelImport285 struct {
	Atomic                          *bool                       `json:"atomic,omitempty"`
	BtType                          *string                     `json:"btType,omitempty"`
	DocumentationType               *GBTPDefinitionType         `json:"documentationType,omitempty"`
	EndSourceLocation               *int32                      `json:"endSourceLocation,omitempty"`
	NodeId                          *string                     `json:"nodeId,omitempty"`
	ShortDescriptor                 *string                     `json:"shortDescriptor,omitempty"`
	SpaceAfter                      *BTPSpace10                 `json:"spaceAfter,omitempty"`
	SpaceBefore                     *BTPSpace10                 `json:"spaceBefore,omitempty"`
	SpaceDefault                    *bool                       `json:"spaceDefault,omitempty"`
	StartSourceLocation             *int32                      `json:"startSourceLocation,omitempty"`
	Annotation                      *BTPAnnotation231           `json:"annotation,omitempty"`
	ArgumentsToDocument             []BTPArgumentDeclaration232 `json:"argumentsToDocument,omitempty"`
	Deprecated                      *bool                       `json:"deprecated,omitempty"`
	DeprecatedExplanation           *string                     `json:"deprecatedExplanation,omitempty"`
	ForExport                       *bool                       `json:"forExport,omitempty"`
	SpaceAfterExport                *BTPSpace10                 `json:"spaceAfterExport,omitempty"`
	SymbolName                      *BTPIdentifier8             `json:"symbolName,omitempty"`
	CombinedNamespacePathAndVersion *string                     `json:"combinedNamespacePathAndVersion,omitempty"`
	ImportMicroversion              *string                     `json:"importMicroversion,omitempty"`
	ModuleId                        *BTPModuleId235             `json:"moduleId,omitempty"`
	Namespace                       []BTPIdentifier8            `json:"namespace,omitempty"`
	NamespaceString                 *string                     `json:"namespaceString,omitempty"`
	SpaceBeforeImport               *BTPSpace10                 `json:"spaceBeforeImport,omitempty"`
}

// NewBTPTopLevelImport285 instantiates a new BTPTopLevelImport285 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPTopLevelImport285() *BTPTopLevelImport285 {
	this := BTPTopLevelImport285{}
	return &this
}

// NewBTPTopLevelImport285WithDefaults instantiates a new BTPTopLevelImport285 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPTopLevelImport285WithDefaults() *BTPTopLevelImport285 {
	this := BTPTopLevelImport285{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPTopLevelImport285) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPTopLevelImport285) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPTopLevelImport285) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPTopLevelImport285) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPTopLevelImport285) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPTopLevelImport285) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPTopLevelImport285) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPTopLevelImport285) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPTopLevelImport285) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPTopLevelImport285) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPTopLevelImport285) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetArgumentsToDocument returns the ArgumentsToDocument field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetArgumentsToDocument() []BTPArgumentDeclaration232 {
	if o == nil || o.ArgumentsToDocument == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return o.ArgumentsToDocument
}

// GetArgumentsToDocumentOk returns a tuple with the ArgumentsToDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetArgumentsToDocumentOk() ([]BTPArgumentDeclaration232, bool) {
	if o == nil || o.ArgumentsToDocument == nil {
		return nil, false
	}
	return o.ArgumentsToDocument, true
}

// HasArgumentsToDocument returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasArgumentsToDocument() bool {
	if o != nil && o.ArgumentsToDocument != nil {
		return true
	}

	return false
}

// SetArgumentsToDocument gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the ArgumentsToDocument field.
func (o *BTPTopLevelImport285) SetArgumentsToDocument(v []BTPArgumentDeclaration232) {
	o.ArgumentsToDocument = v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *BTPTopLevelImport285) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecatedExplanation returns the DeprecatedExplanation field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetDeprecatedExplanation() string {
	if o == nil || o.DeprecatedExplanation == nil {
		var ret string
		return ret
	}
	return *o.DeprecatedExplanation
}

// GetDeprecatedExplanationOk returns a tuple with the DeprecatedExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetDeprecatedExplanationOk() (*string, bool) {
	if o == nil || o.DeprecatedExplanation == nil {
		return nil, false
	}
	return o.DeprecatedExplanation, true
}

// HasDeprecatedExplanation returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasDeprecatedExplanation() bool {
	if o != nil && o.DeprecatedExplanation != nil {
		return true
	}

	return false
}

// SetDeprecatedExplanation gets a reference to the given string and assigns it to the DeprecatedExplanation field.
func (o *BTPTopLevelImport285) SetDeprecatedExplanation(v string) {
	o.DeprecatedExplanation = &v
}

// GetForExport returns the ForExport field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetForExport() bool {
	if o == nil || o.ForExport == nil {
		var ret bool
		return ret
	}
	return *o.ForExport
}

// GetForExportOk returns a tuple with the ForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetForExportOk() (*bool, bool) {
	if o == nil || o.ForExport == nil {
		return nil, false
	}
	return o.ForExport, true
}

// HasForExport returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasForExport() bool {
	if o != nil && o.ForExport != nil {
		return true
	}

	return false
}

// SetForExport gets a reference to the given bool and assigns it to the ForExport field.
func (o *BTPTopLevelImport285) SetForExport(v bool) {
	o.ForExport = &v
}

// GetSpaceAfterExport returns the SpaceAfterExport field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSpaceAfterExport() BTPSpace10 {
	if o == nil || o.SpaceAfterExport == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterExport
}

// GetSpaceAfterExportOk returns a tuple with the SpaceAfterExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSpaceAfterExportOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterExport == nil {
		return nil, false
	}
	return o.SpaceAfterExport, true
}

// HasSpaceAfterExport returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSpaceAfterExport() bool {
	if o != nil && o.SpaceAfterExport != nil {
		return true
	}

	return false
}

// SetSpaceAfterExport gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterExport field.
func (o *BTPTopLevelImport285) SetSpaceAfterExport(v BTPSpace10) {
	o.SpaceAfterExport = &v
}

// GetSymbolName returns the SymbolName field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSymbolName() BTPIdentifier8 {
	if o == nil || o.SymbolName == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.SymbolName
}

// GetSymbolNameOk returns a tuple with the SymbolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSymbolNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.SymbolName == nil {
		return nil, false
	}
	return o.SymbolName, true
}

// HasSymbolName returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSymbolName() bool {
	if o != nil && o.SymbolName != nil {
		return true
	}

	return false
}

// SetSymbolName gets a reference to the given BTPIdentifier8 and assigns it to the SymbolName field.
func (o *BTPTopLevelImport285) SetSymbolName(v BTPIdentifier8) {
	o.SymbolName = &v
}

// GetCombinedNamespacePathAndVersion returns the CombinedNamespacePathAndVersion field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetCombinedNamespacePathAndVersion() string {
	if o == nil || o.CombinedNamespacePathAndVersion == nil {
		var ret string
		return ret
	}
	return *o.CombinedNamespacePathAndVersion
}

// GetCombinedNamespacePathAndVersionOk returns a tuple with the CombinedNamespacePathAndVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetCombinedNamespacePathAndVersionOk() (*string, bool) {
	if o == nil || o.CombinedNamespacePathAndVersion == nil {
		return nil, false
	}
	return o.CombinedNamespacePathAndVersion, true
}

// HasCombinedNamespacePathAndVersion returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasCombinedNamespacePathAndVersion() bool {
	if o != nil && o.CombinedNamespacePathAndVersion != nil {
		return true
	}

	return false
}

// SetCombinedNamespacePathAndVersion gets a reference to the given string and assigns it to the CombinedNamespacePathAndVersion field.
func (o *BTPTopLevelImport285) SetCombinedNamespacePathAndVersion(v string) {
	o.CombinedNamespacePathAndVersion = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTPTopLevelImport285) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetModuleId() BTPModuleId235 {
	if o == nil || o.ModuleId == nil {
		var ret BTPModuleId235
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetModuleIdOk() (*BTPModuleId235, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given BTPModuleId235 and assigns it to the ModuleId field.
func (o *BTPTopLevelImport285) SetModuleId(v BTPModuleId235) {
	o.ModuleId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetNamespace() []BTPIdentifier8 {
	if o == nil || o.Namespace == nil {
		var ret []BTPIdentifier8
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetNamespaceOk() ([]BTPIdentifier8, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given []BTPIdentifier8 and assigns it to the Namespace field.
func (o *BTPTopLevelImport285) SetNamespace(v []BTPIdentifier8) {
	o.Namespace = v
}

// GetNamespaceString returns the NamespaceString field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetNamespaceString() string {
	if o == nil || o.NamespaceString == nil {
		var ret string
		return ret
	}
	return *o.NamespaceString
}

// GetNamespaceStringOk returns a tuple with the NamespaceString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetNamespaceStringOk() (*string, bool) {
	if o == nil || o.NamespaceString == nil {
		return nil, false
	}
	return o.NamespaceString, true
}

// HasNamespaceString returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasNamespaceString() bool {
	if o != nil && o.NamespaceString != nil {
		return true
	}

	return false
}

// SetNamespaceString gets a reference to the given string and assigns it to the NamespaceString field.
func (o *BTPTopLevelImport285) SetNamespaceString(v string) {
	o.NamespaceString = &v
}

// GetSpaceBeforeImport returns the SpaceBeforeImport field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSpaceBeforeImport() BTPSpace10 {
	if o == nil || o.SpaceBeforeImport == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeImport
}

// GetSpaceBeforeImportOk returns a tuple with the SpaceBeforeImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSpaceBeforeImportOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeImport == nil {
		return nil, false
	}
	return o.SpaceBeforeImport, true
}

// HasSpaceBeforeImport returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSpaceBeforeImport() bool {
	if o != nil && o.SpaceBeforeImport != nil {
		return true
	}

	return false
}

// SetSpaceBeforeImport gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeImport field.
func (o *BTPTopLevelImport285) SetSpaceBeforeImport(v BTPSpace10) {
	o.SpaceBeforeImport = &v
}

func (o BTPTopLevelImport285) MarshalJSON() ([]byte, error) {
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
	if o.CombinedNamespacePathAndVersion != nil {
		toSerialize["combinedNamespacePathAndVersion"] = o.CombinedNamespacePathAndVersion
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.ModuleId != nil {
		toSerialize["moduleId"] = o.ModuleId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NamespaceString != nil {
		toSerialize["namespaceString"] = o.NamespaceString
	}
	if o.SpaceBeforeImport != nil {
		toSerialize["spaceBeforeImport"] = o.SpaceBeforeImport
	}
	return json.Marshal(toSerialize)
}

type NullableBTPTopLevelImport285 struct {
	value *BTPTopLevelImport285
	isSet bool
}

func (v NullableBTPTopLevelImport285) Get() *BTPTopLevelImport285 {
	return v.value
}

func (v *NullableBTPTopLevelImport285) Set(val *BTPTopLevelImport285) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPTopLevelImport285) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPTopLevelImport285) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPTopLevelImport285(val *BTPTopLevelImport285) *NullableBTPTopLevelImport285 {
	return &NullableBTPTopLevelImport285{value: val, isSet: true}
}

func (v NullableBTPTopLevelImport285) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPTopLevelImport285) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
