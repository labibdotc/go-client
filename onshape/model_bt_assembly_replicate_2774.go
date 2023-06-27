/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18070-f42189da2c4f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyReplicate2774 struct for BTAssemblyReplicate2774
type BTAssemblyReplicate2774 struct {
	BtType                      *string                               `json:"btType,omitempty"`
	ImportMicroversion          *string                               `json:"importMicroversion,omitempty"`
	NodeId                      *string                               `json:"nodeId,omitempty"`
	AssemblyInstance            *bool                                 `json:"assemblyInstance,omitempty"`
	AssemblyPattern             *bool                                 `json:"assemblyPattern,omitempty"`
	AssemblyReplicate           *bool                                 `json:"assemblyReplicate,omitempty"`
	ClonedInstance              *bool                                 `json:"clonedInstance,omitempty"`
	CustomData                  *map[string]BTReferenceCustomData1551 `json:"customData,omitempty"`
	InstanceFolder              *bool                                 `json:"instanceFolder,omitempty"`
	InstanceName                *string                               `json:"instanceName,omitempty"`
	IsFlattenedPart             *bool                                 `json:"isFlattenedPart,omitempty"`
	Locked                      *bool                                 `json:"locked,omitempty"`
	MicroversionId              *BTMicroversionId366                  `json:"microversionId,omitempty"`
	ParametricInstance          *bool                                 `json:"parametricInstance,omitempty"`
	PartInstance                *bool                                 `json:"partInstance,omitempty"`
	Releasable                  *bool                                 `json:"releasable,omitempty"`
	RevisionCustomData          *BTRevisionCustomData2090             `json:"revisionCustomData,omitempty"`
	StandardContent             *bool                                 `json:"standardContent,omitempty"`
	StandardContentParametersId *string                               `json:"standardContentParametersId,omitempty"`
	Suppressed                  *bool                                 `json:"suppressed,omitempty"`
	SuppressedFieldIndex        *int32                                `json:"suppressedFieldIndex,omitempty"`
	SuppressionConfigured       *bool                                 `json:"suppressionConfigured,omitempty"`
	SuppressionState            *BTMSuppressionState1924              `json:"suppressionState,omitempty"`
	ValidRevisionReference      *bool                                 `json:"validRevisionReference,omitempty"`
	Version                     *int32                                `json:"version,omitempty"`
	Feature                     *BTMAssemblyFeature887                `json:"feature,omitempty"`
	FeatureId                   *string                               `json:"featureId,omitempty"`
	InstanceControlNodes        []BTInstanceControlNode750            `json:"instanceControlNodes,omitempty"`
	ReplicateFeature            *BTMAssemblyReplicateFeature1351      `json:"replicateFeature,omitempty"`
}

// NewBTAssemblyReplicate2774 instantiates a new BTAssemblyReplicate2774 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyReplicate2774() *BTAssemblyReplicate2774 {
	this := BTAssemblyReplicate2774{}
	return &this
}

// NewBTAssemblyReplicate2774WithDefaults instantiates a new BTAssemblyReplicate2774 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyReplicate2774WithDefaults() *BTAssemblyReplicate2774 {
	this := BTAssemblyReplicate2774{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAssemblyReplicate2774) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTAssemblyReplicate2774) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTAssemblyReplicate2774) SetNodeId(v string) {
	o.NodeId = &v
}

// GetAssemblyInstance returns the AssemblyInstance field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetAssemblyInstance() bool {
	if o == nil || o.AssemblyInstance == nil {
		var ret bool
		return ret
	}
	return *o.AssemblyInstance
}

// GetAssemblyInstanceOk returns a tuple with the AssemblyInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetAssemblyInstanceOk() (*bool, bool) {
	if o == nil || o.AssemblyInstance == nil {
		return nil, false
	}
	return o.AssemblyInstance, true
}

// HasAssemblyInstance returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasAssemblyInstance() bool {
	if o != nil && o.AssemblyInstance != nil {
		return true
	}

	return false
}

// SetAssemblyInstance gets a reference to the given bool and assigns it to the AssemblyInstance field.
func (o *BTAssemblyReplicate2774) SetAssemblyInstance(v bool) {
	o.AssemblyInstance = &v
}

// GetAssemblyPattern returns the AssemblyPattern field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetAssemblyPattern() bool {
	if o == nil || o.AssemblyPattern == nil {
		var ret bool
		return ret
	}
	return *o.AssemblyPattern
}

// GetAssemblyPatternOk returns a tuple with the AssemblyPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetAssemblyPatternOk() (*bool, bool) {
	if o == nil || o.AssemblyPattern == nil {
		return nil, false
	}
	return o.AssemblyPattern, true
}

// HasAssemblyPattern returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasAssemblyPattern() bool {
	if o != nil && o.AssemblyPattern != nil {
		return true
	}

	return false
}

// SetAssemblyPattern gets a reference to the given bool and assigns it to the AssemblyPattern field.
func (o *BTAssemblyReplicate2774) SetAssemblyPattern(v bool) {
	o.AssemblyPattern = &v
}

// GetAssemblyReplicate returns the AssemblyReplicate field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetAssemblyReplicate() bool {
	if o == nil || o.AssemblyReplicate == nil {
		var ret bool
		return ret
	}
	return *o.AssemblyReplicate
}

// GetAssemblyReplicateOk returns a tuple with the AssemblyReplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetAssemblyReplicateOk() (*bool, bool) {
	if o == nil || o.AssemblyReplicate == nil {
		return nil, false
	}
	return o.AssemblyReplicate, true
}

// HasAssemblyReplicate returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasAssemblyReplicate() bool {
	if o != nil && o.AssemblyReplicate != nil {
		return true
	}

	return false
}

// SetAssemblyReplicate gets a reference to the given bool and assigns it to the AssemblyReplicate field.
func (o *BTAssemblyReplicate2774) SetAssemblyReplicate(v bool) {
	o.AssemblyReplicate = &v
}

// GetClonedInstance returns the ClonedInstance field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetClonedInstance() bool {
	if o == nil || o.ClonedInstance == nil {
		var ret bool
		return ret
	}
	return *o.ClonedInstance
}

// GetClonedInstanceOk returns a tuple with the ClonedInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetClonedInstanceOk() (*bool, bool) {
	if o == nil || o.ClonedInstance == nil {
		return nil, false
	}
	return o.ClonedInstance, true
}

// HasClonedInstance returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasClonedInstance() bool {
	if o != nil && o.ClonedInstance != nil {
		return true
	}

	return false
}

// SetClonedInstance gets a reference to the given bool and assigns it to the ClonedInstance field.
func (o *BTAssemblyReplicate2774) SetClonedInstance(v bool) {
	o.ClonedInstance = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetCustomData() map[string]BTReferenceCustomData1551 {
	if o == nil || o.CustomData == nil {
		var ret map[string]BTReferenceCustomData1551
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetCustomDataOk() (*map[string]BTReferenceCustomData1551, bool) {
	if o == nil || o.CustomData == nil {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasCustomData() bool {
	if o != nil && o.CustomData != nil {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]BTReferenceCustomData1551 and assigns it to the CustomData field.
func (o *BTAssemblyReplicate2774) SetCustomData(v map[string]BTReferenceCustomData1551) {
	o.CustomData = &v
}

// GetInstanceFolder returns the InstanceFolder field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetInstanceFolder() bool {
	if o == nil || o.InstanceFolder == nil {
		var ret bool
		return ret
	}
	return *o.InstanceFolder
}

// GetInstanceFolderOk returns a tuple with the InstanceFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetInstanceFolderOk() (*bool, bool) {
	if o == nil || o.InstanceFolder == nil {
		return nil, false
	}
	return o.InstanceFolder, true
}

// HasInstanceFolder returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasInstanceFolder() bool {
	if o != nil && o.InstanceFolder != nil {
		return true
	}

	return false
}

// SetInstanceFolder gets a reference to the given bool and assigns it to the InstanceFolder field.
func (o *BTAssemblyReplicate2774) SetInstanceFolder(v bool) {
	o.InstanceFolder = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasInstanceName() bool {
	if o != nil && o.InstanceName != nil {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *BTAssemblyReplicate2774) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetIsFlattenedPart returns the IsFlattenedPart field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetIsFlattenedPart() bool {
	if o == nil || o.IsFlattenedPart == nil {
		var ret bool
		return ret
	}
	return *o.IsFlattenedPart
}

// GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetIsFlattenedPartOk() (*bool, bool) {
	if o == nil || o.IsFlattenedPart == nil {
		return nil, false
	}
	return o.IsFlattenedPart, true
}

// HasIsFlattenedPart returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasIsFlattenedPart() bool {
	if o != nil && o.IsFlattenedPart != nil {
		return true
	}

	return false
}

// SetIsFlattenedPart gets a reference to the given bool and assigns it to the IsFlattenedPart field.
func (o *BTAssemblyReplicate2774) SetIsFlattenedPart(v bool) {
	o.IsFlattenedPart = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BTAssemblyReplicate2774) SetLocked(v bool) {
	o.Locked = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTAssemblyReplicate2774) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetParametricInstance returns the ParametricInstance field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetParametricInstance() bool {
	if o == nil || o.ParametricInstance == nil {
		var ret bool
		return ret
	}
	return *o.ParametricInstance
}

// GetParametricInstanceOk returns a tuple with the ParametricInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetParametricInstanceOk() (*bool, bool) {
	if o == nil || o.ParametricInstance == nil {
		return nil, false
	}
	return o.ParametricInstance, true
}

// HasParametricInstance returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasParametricInstance() bool {
	if o != nil && o.ParametricInstance != nil {
		return true
	}

	return false
}

// SetParametricInstance gets a reference to the given bool and assigns it to the ParametricInstance field.
func (o *BTAssemblyReplicate2774) SetParametricInstance(v bool) {
	o.ParametricInstance = &v
}

// GetPartInstance returns the PartInstance field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetPartInstance() bool {
	if o == nil || o.PartInstance == nil {
		var ret bool
		return ret
	}
	return *o.PartInstance
}

// GetPartInstanceOk returns a tuple with the PartInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetPartInstanceOk() (*bool, bool) {
	if o == nil || o.PartInstance == nil {
		return nil, false
	}
	return o.PartInstance, true
}

// HasPartInstance returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasPartInstance() bool {
	if o != nil && o.PartInstance != nil {
		return true
	}

	return false
}

// SetPartInstance gets a reference to the given bool and assigns it to the PartInstance field.
func (o *BTAssemblyReplicate2774) SetPartInstance(v bool) {
	o.PartInstance = &v
}

// GetReleasable returns the Releasable field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetReleasable() bool {
	if o == nil || o.Releasable == nil {
		var ret bool
		return ret
	}
	return *o.Releasable
}

// GetReleasableOk returns a tuple with the Releasable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetReleasableOk() (*bool, bool) {
	if o == nil || o.Releasable == nil {
		return nil, false
	}
	return o.Releasable, true
}

// HasReleasable returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasReleasable() bool {
	if o != nil && o.Releasable != nil {
		return true
	}

	return false
}

// SetReleasable gets a reference to the given bool and assigns it to the Releasable field.
func (o *BTAssemblyReplicate2774) SetReleasable(v bool) {
	o.Releasable = &v
}

// GetRevisionCustomData returns the RevisionCustomData field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetRevisionCustomData() BTRevisionCustomData2090 {
	if o == nil || o.RevisionCustomData == nil {
		var ret BTRevisionCustomData2090
		return ret
	}
	return *o.RevisionCustomData
}

// GetRevisionCustomDataOk returns a tuple with the RevisionCustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetRevisionCustomDataOk() (*BTRevisionCustomData2090, bool) {
	if o == nil || o.RevisionCustomData == nil {
		return nil, false
	}
	return o.RevisionCustomData, true
}

// HasRevisionCustomData returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasRevisionCustomData() bool {
	if o != nil && o.RevisionCustomData != nil {
		return true
	}

	return false
}

// SetRevisionCustomData gets a reference to the given BTRevisionCustomData2090 and assigns it to the RevisionCustomData field.
func (o *BTAssemblyReplicate2774) SetRevisionCustomData(v BTRevisionCustomData2090) {
	o.RevisionCustomData = &v
}

// GetStandardContent returns the StandardContent field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetStandardContent() bool {
	if o == nil || o.StandardContent == nil {
		var ret bool
		return ret
	}
	return *o.StandardContent
}

// GetStandardContentOk returns a tuple with the StandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetStandardContentOk() (*bool, bool) {
	if o == nil || o.StandardContent == nil {
		return nil, false
	}
	return o.StandardContent, true
}

// HasStandardContent returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasStandardContent() bool {
	if o != nil && o.StandardContent != nil {
		return true
	}

	return false
}

// SetStandardContent gets a reference to the given bool and assigns it to the StandardContent field.
func (o *BTAssemblyReplicate2774) SetStandardContent(v bool) {
	o.StandardContent = &v
}

// GetStandardContentParametersId returns the StandardContentParametersId field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetStandardContentParametersId() string {
	if o == nil || o.StandardContentParametersId == nil {
		var ret string
		return ret
	}
	return *o.StandardContentParametersId
}

// GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetStandardContentParametersIdOk() (*string, bool) {
	if o == nil || o.StandardContentParametersId == nil {
		return nil, false
	}
	return o.StandardContentParametersId, true
}

// HasStandardContentParametersId returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasStandardContentParametersId() bool {
	if o != nil && o.StandardContentParametersId != nil {
		return true
	}

	return false
}

// SetStandardContentParametersId gets a reference to the given string and assigns it to the StandardContentParametersId field.
func (o *BTAssemblyReplicate2774) SetStandardContentParametersId(v string) {
	o.StandardContentParametersId = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTAssemblyReplicate2774) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressedFieldIndex returns the SuppressedFieldIndex field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetSuppressedFieldIndex() int32 {
	if o == nil || o.SuppressedFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.SuppressedFieldIndex
}

// GetSuppressedFieldIndexOk returns a tuple with the SuppressedFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetSuppressedFieldIndexOk() (*int32, bool) {
	if o == nil || o.SuppressedFieldIndex == nil {
		return nil, false
	}
	return o.SuppressedFieldIndex, true
}

// HasSuppressedFieldIndex returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasSuppressedFieldIndex() bool {
	if o != nil && o.SuppressedFieldIndex != nil {
		return true
	}

	return false
}

// SetSuppressedFieldIndex gets a reference to the given int32 and assigns it to the SuppressedFieldIndex field.
func (o *BTAssemblyReplicate2774) SetSuppressedFieldIndex(v int32) {
	o.SuppressedFieldIndex = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTAssemblyReplicate2774) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetSuppressionState returns the SuppressionState field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetSuppressionState() BTMSuppressionState1924 {
	if o == nil || o.SuppressionState == nil {
		var ret BTMSuppressionState1924
		return ret
	}
	return *o.SuppressionState
}

// GetSuppressionStateOk returns a tuple with the SuppressionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetSuppressionStateOk() (*BTMSuppressionState1924, bool) {
	if o == nil || o.SuppressionState == nil {
		return nil, false
	}
	return o.SuppressionState, true
}

// HasSuppressionState returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasSuppressionState() bool {
	if o != nil && o.SuppressionState != nil {
		return true
	}

	return false
}

// SetSuppressionState gets a reference to the given BTMSuppressionState1924 and assigns it to the SuppressionState field.
func (o *BTAssemblyReplicate2774) SetSuppressionState(v BTMSuppressionState1924) {
	o.SuppressionState = &v
}

// GetValidRevisionReference returns the ValidRevisionReference field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetValidRevisionReference() bool {
	if o == nil || o.ValidRevisionReference == nil {
		var ret bool
		return ret
	}
	return *o.ValidRevisionReference
}

// GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetValidRevisionReferenceOk() (*bool, bool) {
	if o == nil || o.ValidRevisionReference == nil {
		return nil, false
	}
	return o.ValidRevisionReference, true
}

// HasValidRevisionReference returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasValidRevisionReference() bool {
	if o != nil && o.ValidRevisionReference != nil {
		return true
	}

	return false
}

// SetValidRevisionReference gets a reference to the given bool and assigns it to the ValidRevisionReference field.
func (o *BTAssemblyReplicate2774) SetValidRevisionReference(v bool) {
	o.ValidRevisionReference = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTAssemblyReplicate2774) SetVersion(v int32) {
	o.Version = &v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetFeature() BTMAssemblyFeature887 {
	if o == nil || o.Feature == nil {
		var ret BTMAssemblyFeature887
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetFeatureOk() (*BTMAssemblyFeature887, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given BTMAssemblyFeature887 and assigns it to the Feature field.
func (o *BTAssemblyReplicate2774) SetFeature(v BTMAssemblyFeature887) {
	o.Feature = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTAssemblyReplicate2774) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetInstanceControlNodes returns the InstanceControlNodes field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetInstanceControlNodes() []BTInstanceControlNode750 {
	if o == nil || o.InstanceControlNodes == nil {
		var ret []BTInstanceControlNode750
		return ret
	}
	return o.InstanceControlNodes
}

// GetInstanceControlNodesOk returns a tuple with the InstanceControlNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetInstanceControlNodesOk() ([]BTInstanceControlNode750, bool) {
	if o == nil || o.InstanceControlNodes == nil {
		return nil, false
	}
	return o.InstanceControlNodes, true
}

// HasInstanceControlNodes returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasInstanceControlNodes() bool {
	if o != nil && o.InstanceControlNodes != nil {
		return true
	}

	return false
}

// SetInstanceControlNodes gets a reference to the given []BTInstanceControlNode750 and assigns it to the InstanceControlNodes field.
func (o *BTAssemblyReplicate2774) SetInstanceControlNodes(v []BTInstanceControlNode750) {
	o.InstanceControlNodes = v
}

// GetReplicateFeature returns the ReplicateFeature field value if set, zero value otherwise.
func (o *BTAssemblyReplicate2774) GetReplicateFeature() BTMAssemblyReplicateFeature1351 {
	if o == nil || o.ReplicateFeature == nil {
		var ret BTMAssemblyReplicateFeature1351
		return ret
	}
	return *o.ReplicateFeature
}

// GetReplicateFeatureOk returns a tuple with the ReplicateFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReplicate2774) GetReplicateFeatureOk() (*BTMAssemblyReplicateFeature1351, bool) {
	if o == nil || o.ReplicateFeature == nil {
		return nil, false
	}
	return o.ReplicateFeature, true
}

// HasReplicateFeature returns a boolean if a field has been set.
func (o *BTAssemblyReplicate2774) HasReplicateFeature() bool {
	if o != nil && o.ReplicateFeature != nil {
		return true
	}

	return false
}

// SetReplicateFeature gets a reference to the given BTMAssemblyReplicateFeature1351 and assigns it to the ReplicateFeature field.
func (o *BTAssemblyReplicate2774) SetReplicateFeature(v BTMAssemblyReplicateFeature1351) {
	o.ReplicateFeature = &v
}

func (o BTAssemblyReplicate2774) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.AssemblyInstance != nil {
		toSerialize["assemblyInstance"] = o.AssemblyInstance
	}
	if o.AssemblyPattern != nil {
		toSerialize["assemblyPattern"] = o.AssemblyPattern
	}
	if o.AssemblyReplicate != nil {
		toSerialize["assemblyReplicate"] = o.AssemblyReplicate
	}
	if o.ClonedInstance != nil {
		toSerialize["clonedInstance"] = o.ClonedInstance
	}
	if o.CustomData != nil {
		toSerialize["customData"] = o.CustomData
	}
	if o.InstanceFolder != nil {
		toSerialize["instanceFolder"] = o.InstanceFolder
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.IsFlattenedPart != nil {
		toSerialize["isFlattenedPart"] = o.IsFlattenedPart
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.ParametricInstance != nil {
		toSerialize["parametricInstance"] = o.ParametricInstance
	}
	if o.PartInstance != nil {
		toSerialize["partInstance"] = o.PartInstance
	}
	if o.Releasable != nil {
		toSerialize["releasable"] = o.Releasable
	}
	if o.RevisionCustomData != nil {
		toSerialize["revisionCustomData"] = o.RevisionCustomData
	}
	if o.StandardContent != nil {
		toSerialize["standardContent"] = o.StandardContent
	}
	if o.StandardContentParametersId != nil {
		toSerialize["standardContentParametersId"] = o.StandardContentParametersId
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressedFieldIndex != nil {
		toSerialize["suppressedFieldIndex"] = o.SuppressedFieldIndex
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.SuppressionState != nil {
		toSerialize["suppressionState"] = o.SuppressionState
	}
	if o.ValidRevisionReference != nil {
		toSerialize["validRevisionReference"] = o.ValidRevisionReference
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Feature != nil {
		toSerialize["feature"] = o.Feature
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.InstanceControlNodes != nil {
		toSerialize["instanceControlNodes"] = o.InstanceControlNodes
	}
	if o.ReplicateFeature != nil {
		toSerialize["replicateFeature"] = o.ReplicateFeature
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyReplicate2774 struct {
	value *BTAssemblyReplicate2774
	isSet bool
}

func (v NullableBTAssemblyReplicate2774) Get() *BTAssemblyReplicate2774 {
	return v.value
}

func (v *NullableBTAssemblyReplicate2774) Set(val *BTAssemblyReplicate2774) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyReplicate2774) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyReplicate2774) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyReplicate2774(val *BTAssemblyReplicate2774) *NullableBTAssemblyReplicate2774 {
	return &NullableBTAssemblyReplicate2774{value: val, isSet: true}
}

func (v NullableBTAssemblyReplicate2774) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyReplicate2774) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
