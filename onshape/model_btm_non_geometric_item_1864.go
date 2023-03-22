/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13258-6b30d30337fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMNonGeometricItem1864 struct for BTMNonGeometricItem1864
type BTMNonGeometricItem1864 struct {
	FeatureId                              *string                                   `json:"featureId,omitempty"`
	FeatureType                            *string                                   `json:"featureType,omitempty"`
	ImportMicroversion                     *string                                   `json:"importMicroversion,omitempty"`
	Name                                   *string                                   `json:"name,omitempty"`
	Namespace                              *string                                   `json:"namespace,omitempty"`
	NodeId                                 *string                                   `json:"nodeId,omitempty"`
	Parameters                             []BTMParameter1                           `json:"parameters,omitempty"`
	ReturnAfterSubfeatures                 *bool                                     `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures                            []BTMFeature134                           `json:"subFeatures,omitempty"`
	Suppressed                             *bool                                     `json:"suppressed,omitempty"`
	SuppressionConfigured                  *bool                                     `json:"suppressionConfigured,omitempty"`
	VariableStudioReference                *bool                                     `json:"variableStudioReference,omitempty"`
	AuxiliaryTreeFeature                   *bool                                     `json:"auxiliaryTreeFeature,omitempty"`
	BtType                                 *string                                   `json:"btType,omitempty"`
	FeatureFolder                          *bool                                     `json:"featureFolder,omitempty"`
	FeatureListFieldIndex                  *int32                                    `json:"featureListFieldIndex,omitempty"`
	FieldIndexForOwnedMateConnectors       *int32                                    `json:"fieldIndexForOwnedMateConnectors,omitempty"`
	OccurrenceQueriesFromAllConfigurations []BTMIndividualQueryWithOccurrenceBase904 `json:"occurrenceQueriesFromAllConfigurations,omitempty"`
	ParametricInstanceFeature              *bool                                     `json:"parametricInstanceFeature,omitempty"`
	Version                                *int32                                    `json:"version,omitempty"`
	ItemDefinitionId                       *string                                   `json:"itemDefinitionId,omitempty"`
}

// NewBTMNonGeometricItem1864 instantiates a new BTMNonGeometricItem1864 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMNonGeometricItem1864() *BTMNonGeometricItem1864 {
	this := BTMNonGeometricItem1864{}
	return &this
}

// NewBTMNonGeometricItem1864WithDefaults instantiates a new BTMNonGeometricItem1864 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMNonGeometricItem1864WithDefaults() *BTMNonGeometricItem1864 {
	this := BTMNonGeometricItem1864{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMNonGeometricItem1864) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTMNonGeometricItem1864) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMNonGeometricItem1864) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMNonGeometricItem1864) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMNonGeometricItem1864) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMNonGeometricItem1864) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMNonGeometricItem1864) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetReturnAfterSubfeatures() bool {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		var ret bool
		return ret
	}
	return *o.ReturnAfterSubfeatures
}

// GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetReturnAfterSubfeaturesOk() (*bool, bool) {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		return nil, false
	}
	return o.ReturnAfterSubfeatures, true
}

// HasReturnAfterSubfeatures returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasReturnAfterSubfeatures() bool {
	if o != nil && o.ReturnAfterSubfeatures != nil {
		return true
	}

	return false
}

// SetReturnAfterSubfeatures gets a reference to the given bool and assigns it to the ReturnAfterSubfeatures field.
func (o *BTMNonGeometricItem1864) SetReturnAfterSubfeatures(v bool) {
	o.ReturnAfterSubfeatures = &v
}

// GetSubFeatures returns the SubFeatures field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetSubFeatures() []BTMFeature134 {
	if o == nil || o.SubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.SubFeatures == nil {
		return nil, false
	}
	return o.SubFeatures, true
}

// HasSubFeatures returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasSubFeatures() bool {
	if o != nil && o.SubFeatures != nil {
		return true
	}

	return false
}

// SetSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the SubFeatures field.
func (o *BTMNonGeometricItem1864) SetSubFeatures(v []BTMFeature134) {
	o.SubFeatures = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTMNonGeometricItem1864) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTMNonGeometricItem1864) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetVariableStudioReference returns the VariableStudioReference field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetVariableStudioReference() bool {
	if o == nil || o.VariableStudioReference == nil {
		var ret bool
		return ret
	}
	return *o.VariableStudioReference
}

// GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetVariableStudioReferenceOk() (*bool, bool) {
	if o == nil || o.VariableStudioReference == nil {
		return nil, false
	}
	return o.VariableStudioReference, true
}

// HasVariableStudioReference returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasVariableStudioReference() bool {
	if o != nil && o.VariableStudioReference != nil {
		return true
	}

	return false
}

// SetVariableStudioReference gets a reference to the given bool and assigns it to the VariableStudioReference field.
func (o *BTMNonGeometricItem1864) SetVariableStudioReference(v bool) {
	o.VariableStudioReference = &v
}

// GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetAuxiliaryTreeFeature() bool {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		var ret bool
		return ret
	}
	return *o.AuxiliaryTreeFeature
}

// GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetAuxiliaryTreeFeatureOk() (*bool, bool) {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		return nil, false
	}
	return o.AuxiliaryTreeFeature, true
}

// HasAuxiliaryTreeFeature returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasAuxiliaryTreeFeature() bool {
	if o != nil && o.AuxiliaryTreeFeature != nil {
		return true
	}

	return false
}

// SetAuxiliaryTreeFeature gets a reference to the given bool and assigns it to the AuxiliaryTreeFeature field.
func (o *BTMNonGeometricItem1864) SetAuxiliaryTreeFeature(v bool) {
	o.AuxiliaryTreeFeature = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMNonGeometricItem1864) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureFolder returns the FeatureFolder field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetFeatureFolder() bool {
	if o == nil || o.FeatureFolder == nil {
		var ret bool
		return ret
	}
	return *o.FeatureFolder
}

// GetFeatureFolderOk returns a tuple with the FeatureFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetFeatureFolderOk() (*bool, bool) {
	if o == nil || o.FeatureFolder == nil {
		return nil, false
	}
	return o.FeatureFolder, true
}

// HasFeatureFolder returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasFeatureFolder() bool {
	if o != nil && o.FeatureFolder != nil {
		return true
	}

	return false
}

// SetFeatureFolder gets a reference to the given bool and assigns it to the FeatureFolder field.
func (o *BTMNonGeometricItem1864) SetFeatureFolder(v bool) {
	o.FeatureFolder = &v
}

// GetFeatureListFieldIndex returns the FeatureListFieldIndex field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetFeatureListFieldIndex() int32 {
	if o == nil || o.FeatureListFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.FeatureListFieldIndex
}

// GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetFeatureListFieldIndexOk() (*int32, bool) {
	if o == nil || o.FeatureListFieldIndex == nil {
		return nil, false
	}
	return o.FeatureListFieldIndex, true
}

// HasFeatureListFieldIndex returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasFeatureListFieldIndex() bool {
	if o != nil && o.FeatureListFieldIndex != nil {
		return true
	}

	return false
}

// SetFeatureListFieldIndex gets a reference to the given int32 and assigns it to the FeatureListFieldIndex field.
func (o *BTMNonGeometricItem1864) SetFeatureListFieldIndex(v int32) {
	o.FeatureListFieldIndex = &v
}

// GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetFieldIndexForOwnedMateConnectors() int32 {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		var ret int32
		return ret
	}
	return *o.FieldIndexForOwnedMateConnectors
}

// GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool) {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		return nil, false
	}
	return o.FieldIndexForOwnedMateConnectors, true
}

// HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasFieldIndexForOwnedMateConnectors() bool {
	if o != nil && o.FieldIndexForOwnedMateConnectors != nil {
		return true
	}

	return false
}

// SetFieldIndexForOwnedMateConnectors gets a reference to the given int32 and assigns it to the FieldIndexForOwnedMateConnectors field.
func (o *BTMNonGeometricItem1864) SetFieldIndexForOwnedMateConnectors(v int32) {
	o.FieldIndexForOwnedMateConnectors = &v
}

// GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904 {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		var ret []BTMIndividualQueryWithOccurrenceBase904
		return ret
	}
	return o.OccurrenceQueriesFromAllConfigurations
}

// GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetOccurrenceQueriesFromAllConfigurationsOk() ([]BTMIndividualQueryWithOccurrenceBase904, bool) {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		return nil, false
	}
	return o.OccurrenceQueriesFromAllConfigurations, true
}

// HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasOccurrenceQueriesFromAllConfigurations() bool {
	if o != nil && o.OccurrenceQueriesFromAllConfigurations != nil {
		return true
	}

	return false
}

// SetOccurrenceQueriesFromAllConfigurations gets a reference to the given []BTMIndividualQueryWithOccurrenceBase904 and assigns it to the OccurrenceQueriesFromAllConfigurations field.
func (o *BTMNonGeometricItem1864) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904) {
	o.OccurrenceQueriesFromAllConfigurations = v
}

// GetParametricInstanceFeature returns the ParametricInstanceFeature field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetParametricInstanceFeature() bool {
	if o == nil || o.ParametricInstanceFeature == nil {
		var ret bool
		return ret
	}
	return *o.ParametricInstanceFeature
}

// GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetParametricInstanceFeatureOk() (*bool, bool) {
	if o == nil || o.ParametricInstanceFeature == nil {
		return nil, false
	}
	return o.ParametricInstanceFeature, true
}

// HasParametricInstanceFeature returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasParametricInstanceFeature() bool {
	if o != nil && o.ParametricInstanceFeature != nil {
		return true
	}

	return false
}

// SetParametricInstanceFeature gets a reference to the given bool and assigns it to the ParametricInstanceFeature field.
func (o *BTMNonGeometricItem1864) SetParametricInstanceFeature(v bool) {
	o.ParametricInstanceFeature = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTMNonGeometricItem1864) SetVersion(v int32) {
	o.Version = &v
}

// GetItemDefinitionId returns the ItemDefinitionId field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864) GetItemDefinitionId() string {
	if o == nil || o.ItemDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.ItemDefinitionId
}

// GetItemDefinitionIdOk returns a tuple with the ItemDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864) GetItemDefinitionIdOk() (*string, bool) {
	if o == nil || o.ItemDefinitionId == nil {
		return nil, false
	}
	return o.ItemDefinitionId, true
}

// HasItemDefinitionId returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864) HasItemDefinitionId() bool {
	if o != nil && o.ItemDefinitionId != nil {
		return true
	}

	return false
}

// SetItemDefinitionId gets a reference to the given string and assigns it to the ItemDefinitionId field.
func (o *BTMNonGeometricItem1864) SetItemDefinitionId(v string) {
	o.ItemDefinitionId = &v
}

func (o BTMNonGeometricItem1864) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.ReturnAfterSubfeatures != nil {
		toSerialize["returnAfterSubfeatures"] = o.ReturnAfterSubfeatures
	}
	if o.SubFeatures != nil {
		toSerialize["subFeatures"] = o.SubFeatures
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.VariableStudioReference != nil {
		toSerialize["variableStudioReference"] = o.VariableStudioReference
	}
	if o.AuxiliaryTreeFeature != nil {
		toSerialize["auxiliaryTreeFeature"] = o.AuxiliaryTreeFeature
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureFolder != nil {
		toSerialize["featureFolder"] = o.FeatureFolder
	}
	if o.FeatureListFieldIndex != nil {
		toSerialize["featureListFieldIndex"] = o.FeatureListFieldIndex
	}
	if o.FieldIndexForOwnedMateConnectors != nil {
		toSerialize["fieldIndexForOwnedMateConnectors"] = o.FieldIndexForOwnedMateConnectors
	}
	if o.OccurrenceQueriesFromAllConfigurations != nil {
		toSerialize["occurrenceQueriesFromAllConfigurations"] = o.OccurrenceQueriesFromAllConfigurations
	}
	if o.ParametricInstanceFeature != nil {
		toSerialize["parametricInstanceFeature"] = o.ParametricInstanceFeature
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ItemDefinitionId != nil {
		toSerialize["itemDefinitionId"] = o.ItemDefinitionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMNonGeometricItem1864 struct {
	value *BTMNonGeometricItem1864
	isSet bool
}

func (v NullableBTMNonGeometricItem1864) Get() *BTMNonGeometricItem1864 {
	return v.value
}

func (v *NullableBTMNonGeometricItem1864) Set(val *BTMNonGeometricItem1864) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMNonGeometricItem1864) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMNonGeometricItem1864) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMNonGeometricItem1864(val *BTMNonGeometricItem1864) *NullableBTMNonGeometricItem1864 {
	return &NullableBTMNonGeometricItem1864{value: val, isSet: true}
}

func (v NullableBTMNonGeometricItem1864) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMNonGeometricItem1864) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
