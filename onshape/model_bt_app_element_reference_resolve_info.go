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

// BTAppElementReferenceResolveInfo struct for BTAppElementReferenceResolveInfo
type BTAppElementReferenceResolveInfo struct {
	ChangeId *string `json:"changeId,omitempty"`
	// The numeric code identifying the error that occurred, if one occurred.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// A human-readable value for the error that occurred, if one occurred.
	ErrorDescription               *string                `json:"errorDescription,omitempty"`
	ErrorValue                     *BTAppElementErrorCode `json:"errorValue,omitempty"`
	IdTag                          *string                `json:"idTag,omitempty"`
	IdTagIsValid                   *bool                  `json:"idTagIsValid,omitempty"`
	IsConfigurable                 *bool                  `json:"isConfigurable,omitempty"`
	IsFlattenedPart                *bool                  `json:"isFlattenedPart,omitempty"`
	IsLocked                       *bool                  `json:"isLocked,omitempty"`
	IsSketchOnly                   *bool                  `json:"isSketchOnly,omitempty"`
	IsSurface                      *bool                  `json:"isSurface,omitempty"`
	LatestElementMicroversionId    *string                `json:"latestElementMicroversionId,omitempty"`
	PartNumber                     *string                `json:"partNumber,omitempty"`
	ReferenceId                    *string                `json:"referenceId,omitempty"`
	ReferenceType                  *int32                 `json:"referenceType,omitempty"`
	ResolvedDocumentMicroversionId *string                `json:"resolvedDocumentMicroversionId,omitempty"`
	ResolvedElementMicroversionId  *string                `json:"resolvedElementMicroversionId,omitempty"`
	Revision                       *string                `json:"revision,omitempty"`
	SketchIds                      []string               `json:"sketchIds,omitempty"`
	TargetConfiguration            *string                `json:"targetConfiguration,omitempty"`
	TargetDocumentId               *string                `json:"targetDocumentId,omitempty"`
	TargetDocumentMicroversionId   *string                `json:"targetDocumentMicroversionId,omitempty"`
	TargetElementId                *string                `json:"targetElementId,omitempty"`
	TargetElementMicroversionId    *string                `json:"targetElementMicroversionId,omitempty"`
	TargetVersionId                *string                `json:"targetVersionId,omitempty"`
	TrackNewVersions               *bool                  `json:"trackNewVersions,omitempty"`
}

// NewBTAppElementReferenceResolveInfo instantiates a new BTAppElementReferenceResolveInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementReferenceResolveInfo() *BTAppElementReferenceResolveInfo {
	this := BTAppElementReferenceResolveInfo{}
	return &this
}

// NewBTAppElementReferenceResolveInfoWithDefaults instantiates a new BTAppElementReferenceResolveInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementReferenceResolveInfoWithDefaults() *BTAppElementReferenceResolveInfo {
	this := BTAppElementReferenceResolveInfo{}
	return &this
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetChangeId() string {
	if o == nil || o.ChangeId == nil {
		var ret string
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetChangeIdOk() (*string, bool) {
	if o == nil || o.ChangeId == nil {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasChangeId() bool {
	if o != nil && o.ChangeId != nil {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given string and assigns it to the ChangeId field.
func (o *BTAppElementReferenceResolveInfo) SetChangeId(v string) {
	o.ChangeId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementReferenceResolveInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementReferenceResolveInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetErrorValue() BTAppElementErrorCode {
	if o == nil || o.ErrorValue == nil {
		var ret BTAppElementErrorCode
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given BTAppElementErrorCode and assigns it to the ErrorValue field.
func (o *BTAppElementReferenceResolveInfo) SetErrorValue(v BTAppElementErrorCode) {
	o.ErrorValue = &v
}

// GetIdTag returns the IdTag field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIdTag() string {
	if o == nil || o.IdTag == nil {
		var ret string
		return ret
	}
	return *o.IdTag
}

// GetIdTagOk returns a tuple with the IdTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIdTagOk() (*string, bool) {
	if o == nil || o.IdTag == nil {
		return nil, false
	}
	return o.IdTag, true
}

// HasIdTag returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIdTag() bool {
	if o != nil && o.IdTag != nil {
		return true
	}

	return false
}

// SetIdTag gets a reference to the given string and assigns it to the IdTag field.
func (o *BTAppElementReferenceResolveInfo) SetIdTag(v string) {
	o.IdTag = &v
}

// GetIdTagIsValid returns the IdTagIsValid field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIdTagIsValid() bool {
	if o == nil || o.IdTagIsValid == nil {
		var ret bool
		return ret
	}
	return *o.IdTagIsValid
}

// GetIdTagIsValidOk returns a tuple with the IdTagIsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIdTagIsValidOk() (*bool, bool) {
	if o == nil || o.IdTagIsValid == nil {
		return nil, false
	}
	return o.IdTagIsValid, true
}

// HasIdTagIsValid returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIdTagIsValid() bool {
	if o != nil && o.IdTagIsValid != nil {
		return true
	}

	return false
}

// SetIdTagIsValid gets a reference to the given bool and assigns it to the IdTagIsValid field.
func (o *BTAppElementReferenceResolveInfo) SetIdTagIsValid(v bool) {
	o.IdTagIsValid = &v
}

// GetIsConfigurable returns the IsConfigurable field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIsConfigurable() bool {
	if o == nil || o.IsConfigurable == nil {
		var ret bool
		return ret
	}
	return *o.IsConfigurable
}

// GetIsConfigurableOk returns a tuple with the IsConfigurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIsConfigurableOk() (*bool, bool) {
	if o == nil || o.IsConfigurable == nil {
		return nil, false
	}
	return o.IsConfigurable, true
}

// HasIsConfigurable returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIsConfigurable() bool {
	if o != nil && o.IsConfigurable != nil {
		return true
	}

	return false
}

// SetIsConfigurable gets a reference to the given bool and assigns it to the IsConfigurable field.
func (o *BTAppElementReferenceResolveInfo) SetIsConfigurable(v bool) {
	o.IsConfigurable = &v
}

// GetIsFlattenedPart returns the IsFlattenedPart field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIsFlattenedPart() bool {
	if o == nil || o.IsFlattenedPart == nil {
		var ret bool
		return ret
	}
	return *o.IsFlattenedPart
}

// GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIsFlattenedPartOk() (*bool, bool) {
	if o == nil || o.IsFlattenedPart == nil {
		return nil, false
	}
	return o.IsFlattenedPart, true
}

// HasIsFlattenedPart returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIsFlattenedPart() bool {
	if o != nil && o.IsFlattenedPart != nil {
		return true
	}

	return false
}

// SetIsFlattenedPart gets a reference to the given bool and assigns it to the IsFlattenedPart field.
func (o *BTAppElementReferenceResolveInfo) SetIsFlattenedPart(v bool) {
	o.IsFlattenedPart = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIsLocked() bool {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIsLockedOk() (*bool, bool) {
	if o == nil || o.IsLocked == nil {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIsLocked() bool {
	if o != nil && o.IsLocked != nil {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *BTAppElementReferenceResolveInfo) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsSketchOnly returns the IsSketchOnly field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIsSketchOnly() bool {
	if o == nil || o.IsSketchOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsSketchOnly
}

// GetIsSketchOnlyOk returns a tuple with the IsSketchOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIsSketchOnlyOk() (*bool, bool) {
	if o == nil || o.IsSketchOnly == nil {
		return nil, false
	}
	return o.IsSketchOnly, true
}

// HasIsSketchOnly returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIsSketchOnly() bool {
	if o != nil && o.IsSketchOnly != nil {
		return true
	}

	return false
}

// SetIsSketchOnly gets a reference to the given bool and assigns it to the IsSketchOnly field.
func (o *BTAppElementReferenceResolveInfo) SetIsSketchOnly(v bool) {
	o.IsSketchOnly = &v
}

// GetIsSurface returns the IsSurface field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetIsSurface() bool {
	if o == nil || o.IsSurface == nil {
		var ret bool
		return ret
	}
	return *o.IsSurface
}

// GetIsSurfaceOk returns a tuple with the IsSurface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetIsSurfaceOk() (*bool, bool) {
	if o == nil || o.IsSurface == nil {
		return nil, false
	}
	return o.IsSurface, true
}

// HasIsSurface returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasIsSurface() bool {
	if o != nil && o.IsSurface != nil {
		return true
	}

	return false
}

// SetIsSurface gets a reference to the given bool and assigns it to the IsSurface field.
func (o *BTAppElementReferenceResolveInfo) SetIsSurface(v bool) {
	o.IsSurface = &v
}

// GetLatestElementMicroversionId returns the LatestElementMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetLatestElementMicroversionId() string {
	if o == nil || o.LatestElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.LatestElementMicroversionId
}

// GetLatestElementMicroversionIdOk returns a tuple with the LatestElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetLatestElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.LatestElementMicroversionId == nil {
		return nil, false
	}
	return o.LatestElementMicroversionId, true
}

// HasLatestElementMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasLatestElementMicroversionId() bool {
	if o != nil && o.LatestElementMicroversionId != nil {
		return true
	}

	return false
}

// SetLatestElementMicroversionId gets a reference to the given string and assigns it to the LatestElementMicroversionId field.
func (o *BTAppElementReferenceResolveInfo) SetLatestElementMicroversionId(v string) {
	o.LatestElementMicroversionId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTAppElementReferenceResolveInfo) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetReferenceIdOk() (*string, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *BTAppElementReferenceResolveInfo) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetReferenceType() int32 {
	if o == nil || o.ReferenceType == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetReferenceTypeOk() (*int32, bool) {
	if o == nil || o.ReferenceType == nil {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasReferenceType() bool {
	if o != nil && o.ReferenceType != nil {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given int32 and assigns it to the ReferenceType field.
func (o *BTAppElementReferenceResolveInfo) SetReferenceType(v int32) {
	o.ReferenceType = &v
}

// GetResolvedDocumentMicroversionId returns the ResolvedDocumentMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetResolvedDocumentMicroversionId() string {
	if o == nil || o.ResolvedDocumentMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ResolvedDocumentMicroversionId
}

// GetResolvedDocumentMicroversionIdOk returns a tuple with the ResolvedDocumentMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetResolvedDocumentMicroversionIdOk() (*string, bool) {
	if o == nil || o.ResolvedDocumentMicroversionId == nil {
		return nil, false
	}
	return o.ResolvedDocumentMicroversionId, true
}

// HasResolvedDocumentMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasResolvedDocumentMicroversionId() bool {
	if o != nil && o.ResolvedDocumentMicroversionId != nil {
		return true
	}

	return false
}

// SetResolvedDocumentMicroversionId gets a reference to the given string and assigns it to the ResolvedDocumentMicroversionId field.
func (o *BTAppElementReferenceResolveInfo) SetResolvedDocumentMicroversionId(v string) {
	o.ResolvedDocumentMicroversionId = &v
}

// GetResolvedElementMicroversionId returns the ResolvedElementMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetResolvedElementMicroversionId() string {
	if o == nil || o.ResolvedElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ResolvedElementMicroversionId
}

// GetResolvedElementMicroversionIdOk returns a tuple with the ResolvedElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetResolvedElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.ResolvedElementMicroversionId == nil {
		return nil, false
	}
	return o.ResolvedElementMicroversionId, true
}

// HasResolvedElementMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasResolvedElementMicroversionId() bool {
	if o != nil && o.ResolvedElementMicroversionId != nil {
		return true
	}

	return false
}

// SetResolvedElementMicroversionId gets a reference to the given string and assigns it to the ResolvedElementMicroversionId field.
func (o *BTAppElementReferenceResolveInfo) SetResolvedElementMicroversionId(v string) {
	o.ResolvedElementMicroversionId = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAppElementReferenceResolveInfo) SetRevision(v string) {
	o.Revision = &v
}

// GetSketchIds returns the SketchIds field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetSketchIds() []string {
	if o == nil || o.SketchIds == nil {
		var ret []string
		return ret
	}
	return o.SketchIds
}

// GetSketchIdsOk returns a tuple with the SketchIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetSketchIdsOk() ([]string, bool) {
	if o == nil || o.SketchIds == nil {
		return nil, false
	}
	return o.SketchIds, true
}

// HasSketchIds returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasSketchIds() bool {
	if o != nil && o.SketchIds != nil {
		return true
	}

	return false
}

// SetSketchIds gets a reference to the given []string and assigns it to the SketchIds field.
func (o *BTAppElementReferenceResolveInfo) SetSketchIds(v []string) {
	o.SketchIds = v
}

// GetTargetConfiguration returns the TargetConfiguration field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTargetConfiguration() string {
	if o == nil || o.TargetConfiguration == nil {
		var ret string
		return ret
	}
	return *o.TargetConfiguration
}

// GetTargetConfigurationOk returns a tuple with the TargetConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTargetConfigurationOk() (*string, bool) {
	if o == nil || o.TargetConfiguration == nil {
		return nil, false
	}
	return o.TargetConfiguration, true
}

// HasTargetConfiguration returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTargetConfiguration() bool {
	if o != nil && o.TargetConfiguration != nil {
		return true
	}

	return false
}

// SetTargetConfiguration gets a reference to the given string and assigns it to the TargetConfiguration field.
func (o *BTAppElementReferenceResolveInfo) SetTargetConfiguration(v string) {
	o.TargetConfiguration = &v
}

// GetTargetDocumentId returns the TargetDocumentId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentId() string {
	if o == nil || o.TargetDocumentId == nil {
		var ret string
		return ret
	}
	return *o.TargetDocumentId
}

// GetTargetDocumentIdOk returns a tuple with the TargetDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentIdOk() (*string, bool) {
	if o == nil || o.TargetDocumentId == nil {
		return nil, false
	}
	return o.TargetDocumentId, true
}

// HasTargetDocumentId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTargetDocumentId() bool {
	if o != nil && o.TargetDocumentId != nil {
		return true
	}

	return false
}

// SetTargetDocumentId gets a reference to the given string and assigns it to the TargetDocumentId field.
func (o *BTAppElementReferenceResolveInfo) SetTargetDocumentId(v string) {
	o.TargetDocumentId = &v
}

// GetTargetDocumentMicroversionId returns the TargetDocumentMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentMicroversionId() string {
	if o == nil || o.TargetDocumentMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.TargetDocumentMicroversionId
}

// GetTargetDocumentMicroversionIdOk returns a tuple with the TargetDocumentMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentMicroversionIdOk() (*string, bool) {
	if o == nil || o.TargetDocumentMicroversionId == nil {
		return nil, false
	}
	return o.TargetDocumentMicroversionId, true
}

// HasTargetDocumentMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTargetDocumentMicroversionId() bool {
	if o != nil && o.TargetDocumentMicroversionId != nil {
		return true
	}

	return false
}

// SetTargetDocumentMicroversionId gets a reference to the given string and assigns it to the TargetDocumentMicroversionId field.
func (o *BTAppElementReferenceResolveInfo) SetTargetDocumentMicroversionId(v string) {
	o.TargetDocumentMicroversionId = &v
}

// GetTargetElementId returns the TargetElementId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTargetElementId() string {
	if o == nil || o.TargetElementId == nil {
		var ret string
		return ret
	}
	return *o.TargetElementId
}

// GetTargetElementIdOk returns a tuple with the TargetElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTargetElementIdOk() (*string, bool) {
	if o == nil || o.TargetElementId == nil {
		return nil, false
	}
	return o.TargetElementId, true
}

// HasTargetElementId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTargetElementId() bool {
	if o != nil && o.TargetElementId != nil {
		return true
	}

	return false
}

// SetTargetElementId gets a reference to the given string and assigns it to the TargetElementId field.
func (o *BTAppElementReferenceResolveInfo) SetTargetElementId(v string) {
	o.TargetElementId = &v
}

// GetTargetElementMicroversionId returns the TargetElementMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTargetElementMicroversionId() string {
	if o == nil || o.TargetElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.TargetElementMicroversionId
}

// GetTargetElementMicroversionIdOk returns a tuple with the TargetElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTargetElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.TargetElementMicroversionId == nil {
		return nil, false
	}
	return o.TargetElementMicroversionId, true
}

// HasTargetElementMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTargetElementMicroversionId() bool {
	if o != nil && o.TargetElementMicroversionId != nil {
		return true
	}

	return false
}

// SetTargetElementMicroversionId gets a reference to the given string and assigns it to the TargetElementMicroversionId field.
func (o *BTAppElementReferenceResolveInfo) SetTargetElementMicroversionId(v string) {
	o.TargetElementMicroversionId = &v
}

// GetTargetVersionId returns the TargetVersionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTargetVersionId() string {
	if o == nil || o.TargetVersionId == nil {
		var ret string
		return ret
	}
	return *o.TargetVersionId
}

// GetTargetVersionIdOk returns a tuple with the TargetVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTargetVersionIdOk() (*string, bool) {
	if o == nil || o.TargetVersionId == nil {
		return nil, false
	}
	return o.TargetVersionId, true
}

// HasTargetVersionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTargetVersionId() bool {
	if o != nil && o.TargetVersionId != nil {
		return true
	}

	return false
}

// SetTargetVersionId gets a reference to the given string and assigns it to the TargetVersionId field.
func (o *BTAppElementReferenceResolveInfo) SetTargetVersionId(v string) {
	o.TargetVersionId = &v
}

// GetTrackNewVersions returns the TrackNewVersions field value if set, zero value otherwise.
func (o *BTAppElementReferenceResolveInfo) GetTrackNewVersions() bool {
	if o == nil || o.TrackNewVersions == nil {
		var ret bool
		return ret
	}
	return *o.TrackNewVersions
}

// GetTrackNewVersionsOk returns a tuple with the TrackNewVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceResolveInfo) GetTrackNewVersionsOk() (*bool, bool) {
	if o == nil || o.TrackNewVersions == nil {
		return nil, false
	}
	return o.TrackNewVersions, true
}

// HasTrackNewVersions returns a boolean if a field has been set.
func (o *BTAppElementReferenceResolveInfo) HasTrackNewVersions() bool {
	if o != nil && o.TrackNewVersions != nil {
		return true
	}

	return false
}

// SetTrackNewVersions gets a reference to the given bool and assigns it to the TrackNewVersions field.
func (o *BTAppElementReferenceResolveInfo) SetTrackNewVersions(v bool) {
	o.TrackNewVersions = &v
}

func (o BTAppElementReferenceResolveInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeId != nil {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	if o.IdTag != nil {
		toSerialize["idTag"] = o.IdTag
	}
	if o.IdTagIsValid != nil {
		toSerialize["idTagIsValid"] = o.IdTagIsValid
	}
	if o.IsConfigurable != nil {
		toSerialize["isConfigurable"] = o.IsConfigurable
	}
	if o.IsFlattenedPart != nil {
		toSerialize["isFlattenedPart"] = o.IsFlattenedPart
	}
	if o.IsLocked != nil {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.IsSketchOnly != nil {
		toSerialize["isSketchOnly"] = o.IsSketchOnly
	}
	if o.IsSurface != nil {
		toSerialize["isSurface"] = o.IsSurface
	}
	if o.LatestElementMicroversionId != nil {
		toSerialize["latestElementMicroversionId"] = o.LatestElementMicroversionId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.ReferenceId != nil {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if o.ReferenceType != nil {
		toSerialize["referenceType"] = o.ReferenceType
	}
	if o.ResolvedDocumentMicroversionId != nil {
		toSerialize["resolvedDocumentMicroversionId"] = o.ResolvedDocumentMicroversionId
	}
	if o.ResolvedElementMicroversionId != nil {
		toSerialize["resolvedElementMicroversionId"] = o.ResolvedElementMicroversionId
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.SketchIds != nil {
		toSerialize["sketchIds"] = o.SketchIds
	}
	if o.TargetConfiguration != nil {
		toSerialize["targetConfiguration"] = o.TargetConfiguration
	}
	if o.TargetDocumentId != nil {
		toSerialize["targetDocumentId"] = o.TargetDocumentId
	}
	if o.TargetDocumentMicroversionId != nil {
		toSerialize["targetDocumentMicroversionId"] = o.TargetDocumentMicroversionId
	}
	if o.TargetElementId != nil {
		toSerialize["targetElementId"] = o.TargetElementId
	}
	if o.TargetElementMicroversionId != nil {
		toSerialize["targetElementMicroversionId"] = o.TargetElementMicroversionId
	}
	if o.TargetVersionId != nil {
		toSerialize["targetVersionId"] = o.TargetVersionId
	}
	if o.TrackNewVersions != nil {
		toSerialize["trackNewVersions"] = o.TrackNewVersions
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementReferenceResolveInfo struct {
	value *BTAppElementReferenceResolveInfo
	isSet bool
}

func (v NullableBTAppElementReferenceResolveInfo) Get() *BTAppElementReferenceResolveInfo {
	return v.value
}

func (v *NullableBTAppElementReferenceResolveInfo) Set(val *BTAppElementReferenceResolveInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementReferenceResolveInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementReferenceResolveInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementReferenceResolveInfo(val *BTAppElementReferenceResolveInfo) *NullableBTAppElementReferenceResolveInfo {
	return &NullableBTAppElementReferenceResolveInfo{value: val, isSet: true}
}

func (v NullableBTAppElementReferenceResolveInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementReferenceResolveInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
