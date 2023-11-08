/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25541-b54c5b71ef45
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPModule234 struct for BTPModule234
type BTPModule234 struct {
	Atomic                 *bool                           `json:"atomic,omitempty"`
	BtType                 *string                         `json:"btType,omitempty"`
	DocumentationType      *GBTPDefinitionType             `json:"documentationType,omitempty"`
	EndSourceLocation      *int32                          `json:"endSourceLocation,omitempty"`
	NodeId                 *string                         `json:"nodeId,omitempty"`
	ShortDescriptor        *string                         `json:"shortDescriptor,omitempty"`
	SpaceAfter             *BTPSpace10                     `json:"spaceAfter,omitempty"`
	SpaceBefore            *BTPSpace10                     `json:"spaceBefore,omitempty"`
	SpaceDefault           *bool                           `json:"spaceDefault,omitempty"`
	StartSourceLocation    *int32                          `json:"startSourceLocation,omitempty"`
	DeepImports            *map[string][]BTImport          `json:"deepImports,omitempty"`
	Imports                []BTPTopLevelImport285          `json:"imports,omitempty"`
	IsBlob                 *bool                           `json:"isBlob,omitempty"`
	IsInternalModule       *bool                           `json:"isInternalModule,omitempty"`
	MayHaveImplicitImports *bool                           `json:"mayHaveImplicitImports,omitempty"`
	PathMap                *map[string]BTMicroversionId366 `json:"pathMap,omitempty"`
	ToBeParsed             *BTLazilyParsedFeatureScript    `json:"toBeParsed,omitempty"`
	TopLevel               []BTPTopLevelNode286            `json:"topLevel,omitempty"`
	Version                *BTPLiteralNumber258            `json:"version,omitempty"`
	VersionNumber          *int32                          `json:"versionNumber,omitempty"`
}

// NewBTPModule234 instantiates a new BTPModule234 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPModule234() *BTPModule234 {
	this := BTPModule234{}
	return &this
}

// NewBTPModule234WithDefaults instantiates a new BTPModule234 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPModule234WithDefaults() *BTPModule234 {
	this := BTPModule234{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPModule234) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPModule234) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPModule234) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPModule234) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPModule234) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPModule234) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPModule234) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPModule234) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPModule234) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPModule234) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPModule234) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPModule234) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPModule234) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPModule234) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPModule234) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPModule234) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPModule234) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPModule234) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPModule234) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPModule234) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPModule234) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPModule234) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPModule234) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPModule234) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPModule234) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPModule234) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPModule234) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPModule234) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPModule234) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPModule234) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetDeepImports returns the DeepImports field value if set, zero value otherwise.
func (o *BTPModule234) GetDeepImports() map[string][]BTImport {
	if o == nil || o.DeepImports == nil {
		var ret map[string][]BTImport
		return ret
	}
	return *o.DeepImports
}

// GetDeepImportsOk returns a tuple with the DeepImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetDeepImportsOk() (*map[string][]BTImport, bool) {
	if o == nil || o.DeepImports == nil {
		return nil, false
	}
	return o.DeepImports, true
}

// HasDeepImports returns a boolean if a field has been set.
func (o *BTPModule234) HasDeepImports() bool {
	if o != nil && o.DeepImports != nil {
		return true
	}

	return false
}

// SetDeepImports gets a reference to the given map[string][]BTImport and assigns it to the DeepImports field.
func (o *BTPModule234) SetDeepImports(v map[string][]BTImport) {
	o.DeepImports = &v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *BTPModule234) GetImports() []BTPTopLevelImport285 {
	if o == nil || o.Imports == nil {
		var ret []BTPTopLevelImport285
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetImportsOk() ([]BTPTopLevelImport285, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *BTPModule234) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []BTPTopLevelImport285 and assigns it to the Imports field.
func (o *BTPModule234) SetImports(v []BTPTopLevelImport285) {
	o.Imports = v
}

// GetIsBlob returns the IsBlob field value if set, zero value otherwise.
func (o *BTPModule234) GetIsBlob() bool {
	if o == nil || o.IsBlob == nil {
		var ret bool
		return ret
	}
	return *o.IsBlob
}

// GetIsBlobOk returns a tuple with the IsBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetIsBlobOk() (*bool, bool) {
	if o == nil || o.IsBlob == nil {
		return nil, false
	}
	return o.IsBlob, true
}

// HasIsBlob returns a boolean if a field has been set.
func (o *BTPModule234) HasIsBlob() bool {
	if o != nil && o.IsBlob != nil {
		return true
	}

	return false
}

// SetIsBlob gets a reference to the given bool and assigns it to the IsBlob field.
func (o *BTPModule234) SetIsBlob(v bool) {
	o.IsBlob = &v
}

// GetIsInternalModule returns the IsInternalModule field value if set, zero value otherwise.
func (o *BTPModule234) GetIsInternalModule() bool {
	if o == nil || o.IsInternalModule == nil {
		var ret bool
		return ret
	}
	return *o.IsInternalModule
}

// GetIsInternalModuleOk returns a tuple with the IsInternalModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetIsInternalModuleOk() (*bool, bool) {
	if o == nil || o.IsInternalModule == nil {
		return nil, false
	}
	return o.IsInternalModule, true
}

// HasIsInternalModule returns a boolean if a field has been set.
func (o *BTPModule234) HasIsInternalModule() bool {
	if o != nil && o.IsInternalModule != nil {
		return true
	}

	return false
}

// SetIsInternalModule gets a reference to the given bool and assigns it to the IsInternalModule field.
func (o *BTPModule234) SetIsInternalModule(v bool) {
	o.IsInternalModule = &v
}

// GetMayHaveImplicitImports returns the MayHaveImplicitImports field value if set, zero value otherwise.
func (o *BTPModule234) GetMayHaveImplicitImports() bool {
	if o == nil || o.MayHaveImplicitImports == nil {
		var ret bool
		return ret
	}
	return *o.MayHaveImplicitImports
}

// GetMayHaveImplicitImportsOk returns a tuple with the MayHaveImplicitImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetMayHaveImplicitImportsOk() (*bool, bool) {
	if o == nil || o.MayHaveImplicitImports == nil {
		return nil, false
	}
	return o.MayHaveImplicitImports, true
}

// HasMayHaveImplicitImports returns a boolean if a field has been set.
func (o *BTPModule234) HasMayHaveImplicitImports() bool {
	if o != nil && o.MayHaveImplicitImports != nil {
		return true
	}

	return false
}

// SetMayHaveImplicitImports gets a reference to the given bool and assigns it to the MayHaveImplicitImports field.
func (o *BTPModule234) SetMayHaveImplicitImports(v bool) {
	o.MayHaveImplicitImports = &v
}

// GetPathMap returns the PathMap field value if set, zero value otherwise.
func (o *BTPModule234) GetPathMap() map[string]BTMicroversionId366 {
	if o == nil || o.PathMap == nil {
		var ret map[string]BTMicroversionId366
		return ret
	}
	return *o.PathMap
}

// GetPathMapOk returns a tuple with the PathMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetPathMapOk() (*map[string]BTMicroversionId366, bool) {
	if o == nil || o.PathMap == nil {
		return nil, false
	}
	return o.PathMap, true
}

// HasPathMap returns a boolean if a field has been set.
func (o *BTPModule234) HasPathMap() bool {
	if o != nil && o.PathMap != nil {
		return true
	}

	return false
}

// SetPathMap gets a reference to the given map[string]BTMicroversionId366 and assigns it to the PathMap field.
func (o *BTPModule234) SetPathMap(v map[string]BTMicroversionId366) {
	o.PathMap = &v
}

// GetToBeParsed returns the ToBeParsed field value if set, zero value otherwise.
func (o *BTPModule234) GetToBeParsed() BTLazilyParsedFeatureScript {
	if o == nil || o.ToBeParsed == nil {
		var ret BTLazilyParsedFeatureScript
		return ret
	}
	return *o.ToBeParsed
}

// GetToBeParsedOk returns a tuple with the ToBeParsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetToBeParsedOk() (*BTLazilyParsedFeatureScript, bool) {
	if o == nil || o.ToBeParsed == nil {
		return nil, false
	}
	return o.ToBeParsed, true
}

// HasToBeParsed returns a boolean if a field has been set.
func (o *BTPModule234) HasToBeParsed() bool {
	if o != nil && o.ToBeParsed != nil {
		return true
	}

	return false
}

// SetToBeParsed gets a reference to the given BTLazilyParsedFeatureScript and assigns it to the ToBeParsed field.
func (o *BTPModule234) SetToBeParsed(v BTLazilyParsedFeatureScript) {
	o.ToBeParsed = &v
}

// GetTopLevel returns the TopLevel field value if set, zero value otherwise.
func (o *BTPModule234) GetTopLevel() []BTPTopLevelNode286 {
	if o == nil || o.TopLevel == nil {
		var ret []BTPTopLevelNode286
		return ret
	}
	return o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetTopLevelOk() ([]BTPTopLevelNode286, bool) {
	if o == nil || o.TopLevel == nil {
		return nil, false
	}
	return o.TopLevel, true
}

// HasTopLevel returns a boolean if a field has been set.
func (o *BTPModule234) HasTopLevel() bool {
	if o != nil && o.TopLevel != nil {
		return true
	}

	return false
}

// SetTopLevel gets a reference to the given []BTPTopLevelNode286 and assigns it to the TopLevel field.
func (o *BTPModule234) SetTopLevel(v []BTPTopLevelNode286) {
	o.TopLevel = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTPModule234) GetVersion() BTPLiteralNumber258 {
	if o == nil || o.Version == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetVersionOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTPModule234) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given BTPLiteralNumber258 and assigns it to the Version field.
func (o *BTPModule234) SetVersion(v BTPLiteralNumber258) {
	o.Version = &v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *BTPModule234) GetVersionNumber() int32 {
	if o == nil || o.VersionNumber == nil {
		var ret int32
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetVersionNumberOk() (*int32, bool) {
	if o == nil || o.VersionNumber == nil {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *BTPModule234) HasVersionNumber() bool {
	if o != nil && o.VersionNumber != nil {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given int32 and assigns it to the VersionNumber field.
func (o *BTPModule234) SetVersionNumber(v int32) {
	o.VersionNumber = &v
}

func (o BTPModule234) MarshalJSON() ([]byte, error) {
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
	if o.DeepImports != nil {
		toSerialize["deepImports"] = o.DeepImports
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.IsBlob != nil {
		toSerialize["isBlob"] = o.IsBlob
	}
	if o.IsInternalModule != nil {
		toSerialize["isInternalModule"] = o.IsInternalModule
	}
	if o.MayHaveImplicitImports != nil {
		toSerialize["mayHaveImplicitImports"] = o.MayHaveImplicitImports
	}
	if o.PathMap != nil {
		toSerialize["pathMap"] = o.PathMap
	}
	if o.ToBeParsed != nil {
		toSerialize["toBeParsed"] = o.ToBeParsed
	}
	if o.TopLevel != nil {
		toSerialize["topLevel"] = o.TopLevel
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionNumber != nil {
		toSerialize["versionNumber"] = o.VersionNumber
	}
	return json.Marshal(toSerialize)
}

type NullableBTPModule234 struct {
	value *BTPModule234
	isSet bool
}

func (v NullableBTPModule234) Get() *BTPModule234 {
	return v.value
}

func (v *NullableBTPModule234) Set(val *BTPModule234) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPModule234) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPModule234) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPModule234(val *BTPModule234) *NullableBTPModule234 {
	return &NullableBTPModule234{value: val, isSet: true}
}

func (v NullableBTPModule234) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPModule234) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
