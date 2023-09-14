/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentElementProcessingInfo struct for BTDocumentElementProcessingInfo
type BTDocumentElementProcessingInfo struct {
	AccelerationUnits    *string                  `json:"accelerationUnits,omitempty"`
	AngleUnits           *string                  `json:"angleUnits,omitempty"`
	AngularVelocityUnits *string                  `json:"angularVelocityUnits,omitempty"`
	ApplicationTarget    *BTApplicationTargetInfo `json:"applicationTarget,omitempty"`
	AreaUnits            *string                  `json:"areaUnits,omitempty"`
	DataType             *string                  `json:"dataType,omitempty"`
	Deleted              *bool                    `json:"deleted,omitempty"`
	ElementType          *GBTElementType          `json:"elementType,omitempty"`
	EnergyUnits          *string                  `json:"energyUnits,omitempty"`
	Filename             *string                  `json:"filename,omitempty"`
	ForceUnits           *string                  `json:"forceUnits,omitempty"`
	ForeignDataId        *string                  `json:"foreignDataId,omitempty"`
	Id                   *string                  `json:"id,omitempty"`
	LengthUnits          *string                  `json:"lengthUnits,omitempty"`
	MassUnits            *string                  `json:"massUnits,omitempty"`
	MicroversionId       *string                  `json:"microversionId,omitempty"`
	MomentUnits          *string                  `json:"momentUnits,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	PressureUnits        *string                  `json:"pressureUnits,omitempty"`
	PrettyType           *string                  `json:"prettyType,omitempty"`
	SafeToShow           *bool                    `json:"safeToShow,omitempty"`
	SpecifiedUnit        *string                  `json:"specifiedUnit,omitempty"`
	ThumbnailInfo        *BTThumbnailInfo         `json:"thumbnailInfo,omitempty"`
	Thumbnails           *string                  `json:"thumbnails,omitempty"`
	TimeUnits            *string                  `json:"timeUnits,omitempty"`
	TranslationEventKey  *string                  `json:"translationEventKey,omitempty"`
	TranslationId        *string                  `json:"translationId,omitempty"`
	Type                 *string                  `json:"type,omitempty"`
	Unupdatable          *bool                    `json:"unupdatable,omitempty"`
	VolumeUnits          *string                  `json:"volumeUnits,omitempty"`
	Zip                  *BTZipFileInfo           `json:"zip,omitempty"`
}

// NewBTDocumentElementProcessingInfo instantiates a new BTDocumentElementProcessingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentElementProcessingInfo() *BTDocumentElementProcessingInfo {
	this := BTDocumentElementProcessingInfo{}
	return &this
}

// NewBTDocumentElementProcessingInfoWithDefaults instantiates a new BTDocumentElementProcessingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentElementProcessingInfoWithDefaults() *BTDocumentElementProcessingInfo {
	this := BTDocumentElementProcessingInfo{}
	return &this
}

// GetAccelerationUnits returns the AccelerationUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetAccelerationUnits() string {
	if o == nil || o.AccelerationUnits == nil {
		var ret string
		return ret
	}
	return *o.AccelerationUnits
}

// GetAccelerationUnitsOk returns a tuple with the AccelerationUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetAccelerationUnitsOk() (*string, bool) {
	if o == nil || o.AccelerationUnits == nil {
		return nil, false
	}
	return o.AccelerationUnits, true
}

// HasAccelerationUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasAccelerationUnits() bool {
	if o != nil && o.AccelerationUnits != nil {
		return true
	}

	return false
}

// SetAccelerationUnits gets a reference to the given string and assigns it to the AccelerationUnits field.
func (o *BTDocumentElementProcessingInfo) SetAccelerationUnits(v string) {
	o.AccelerationUnits = &v
}

// GetAngleUnits returns the AngleUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetAngleUnits() string {
	if o == nil || o.AngleUnits == nil {
		var ret string
		return ret
	}
	return *o.AngleUnits
}

// GetAngleUnitsOk returns a tuple with the AngleUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetAngleUnitsOk() (*string, bool) {
	if o == nil || o.AngleUnits == nil {
		return nil, false
	}
	return o.AngleUnits, true
}

// HasAngleUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasAngleUnits() bool {
	if o != nil && o.AngleUnits != nil {
		return true
	}

	return false
}

// SetAngleUnits gets a reference to the given string and assigns it to the AngleUnits field.
func (o *BTDocumentElementProcessingInfo) SetAngleUnits(v string) {
	o.AngleUnits = &v
}

// GetAngularVelocityUnits returns the AngularVelocityUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetAngularVelocityUnits() string {
	if o == nil || o.AngularVelocityUnits == nil {
		var ret string
		return ret
	}
	return *o.AngularVelocityUnits
}

// GetAngularVelocityUnitsOk returns a tuple with the AngularVelocityUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetAngularVelocityUnitsOk() (*string, bool) {
	if o == nil || o.AngularVelocityUnits == nil {
		return nil, false
	}
	return o.AngularVelocityUnits, true
}

// HasAngularVelocityUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasAngularVelocityUnits() bool {
	if o != nil && o.AngularVelocityUnits != nil {
		return true
	}

	return false
}

// SetAngularVelocityUnits gets a reference to the given string and assigns it to the AngularVelocityUnits field.
func (o *BTDocumentElementProcessingInfo) SetAngularVelocityUnits(v string) {
	o.AngularVelocityUnits = &v
}

// GetApplicationTarget returns the ApplicationTarget field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetApplicationTarget() BTApplicationTargetInfo {
	if o == nil || o.ApplicationTarget == nil {
		var ret BTApplicationTargetInfo
		return ret
	}
	return *o.ApplicationTarget
}

// GetApplicationTargetOk returns a tuple with the ApplicationTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetApplicationTargetOk() (*BTApplicationTargetInfo, bool) {
	if o == nil || o.ApplicationTarget == nil {
		return nil, false
	}
	return o.ApplicationTarget, true
}

// HasApplicationTarget returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasApplicationTarget() bool {
	if o != nil && o.ApplicationTarget != nil {
		return true
	}

	return false
}

// SetApplicationTarget gets a reference to the given BTApplicationTargetInfo and assigns it to the ApplicationTarget field.
func (o *BTDocumentElementProcessingInfo) SetApplicationTarget(v BTApplicationTargetInfo) {
	o.ApplicationTarget = &v
}

// GetAreaUnits returns the AreaUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetAreaUnits() string {
	if o == nil || o.AreaUnits == nil {
		var ret string
		return ret
	}
	return *o.AreaUnits
}

// GetAreaUnitsOk returns a tuple with the AreaUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetAreaUnitsOk() (*string, bool) {
	if o == nil || o.AreaUnits == nil {
		return nil, false
	}
	return o.AreaUnits, true
}

// HasAreaUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasAreaUnits() bool {
	if o != nil && o.AreaUnits != nil {
		return true
	}

	return false
}

// SetAreaUnits gets a reference to the given string and assigns it to the AreaUnits field.
func (o *BTDocumentElementProcessingInfo) SetAreaUnits(v string) {
	o.AreaUnits = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *BTDocumentElementProcessingInfo) SetDataType(v string) {
	o.DataType = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *BTDocumentElementProcessingInfo) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetElementType() GBTElementType {
	if o == nil || o.ElementType == nil {
		var ret GBTElementType
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetElementTypeOk() (*GBTElementType, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given GBTElementType and assigns it to the ElementType field.
func (o *BTDocumentElementProcessingInfo) SetElementType(v GBTElementType) {
	o.ElementType = &v
}

// GetEnergyUnits returns the EnergyUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetEnergyUnits() string {
	if o == nil || o.EnergyUnits == nil {
		var ret string
		return ret
	}
	return *o.EnergyUnits
}

// GetEnergyUnitsOk returns a tuple with the EnergyUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetEnergyUnitsOk() (*string, bool) {
	if o == nil || o.EnergyUnits == nil {
		return nil, false
	}
	return o.EnergyUnits, true
}

// HasEnergyUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasEnergyUnits() bool {
	if o != nil && o.EnergyUnits != nil {
		return true
	}

	return false
}

// SetEnergyUnits gets a reference to the given string and assigns it to the EnergyUnits field.
func (o *BTDocumentElementProcessingInfo) SetEnergyUnits(v string) {
	o.EnergyUnits = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *BTDocumentElementProcessingInfo) SetFilename(v string) {
	o.Filename = &v
}

// GetForceUnits returns the ForceUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetForceUnits() string {
	if o == nil || o.ForceUnits == nil {
		var ret string
		return ret
	}
	return *o.ForceUnits
}

// GetForceUnitsOk returns a tuple with the ForceUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetForceUnitsOk() (*string, bool) {
	if o == nil || o.ForceUnits == nil {
		return nil, false
	}
	return o.ForceUnits, true
}

// HasForceUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasForceUnits() bool {
	if o != nil && o.ForceUnits != nil {
		return true
	}

	return false
}

// SetForceUnits gets a reference to the given string and assigns it to the ForceUnits field.
func (o *BTDocumentElementProcessingInfo) SetForceUnits(v string) {
	o.ForceUnits = &v
}

// GetForeignDataId returns the ForeignDataId field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetForeignDataId() string {
	if o == nil || o.ForeignDataId == nil {
		var ret string
		return ret
	}
	return *o.ForeignDataId
}

// GetForeignDataIdOk returns a tuple with the ForeignDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetForeignDataIdOk() (*string, bool) {
	if o == nil || o.ForeignDataId == nil {
		return nil, false
	}
	return o.ForeignDataId, true
}

// HasForeignDataId returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasForeignDataId() bool {
	if o != nil && o.ForeignDataId != nil {
		return true
	}

	return false
}

// SetForeignDataId gets a reference to the given string and assigns it to the ForeignDataId field.
func (o *BTDocumentElementProcessingInfo) SetForeignDataId(v string) {
	o.ForeignDataId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTDocumentElementProcessingInfo) SetId(v string) {
	o.Id = &v
}

// GetLengthUnits returns the LengthUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetLengthUnits() string {
	if o == nil || o.LengthUnits == nil {
		var ret string
		return ret
	}
	return *o.LengthUnits
}

// GetLengthUnitsOk returns a tuple with the LengthUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetLengthUnitsOk() (*string, bool) {
	if o == nil || o.LengthUnits == nil {
		return nil, false
	}
	return o.LengthUnits, true
}

// HasLengthUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasLengthUnits() bool {
	if o != nil && o.LengthUnits != nil {
		return true
	}

	return false
}

// SetLengthUnits gets a reference to the given string and assigns it to the LengthUnits field.
func (o *BTDocumentElementProcessingInfo) SetLengthUnits(v string) {
	o.LengthUnits = &v
}

// GetMassUnits returns the MassUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetMassUnits() string {
	if o == nil || o.MassUnits == nil {
		var ret string
		return ret
	}
	return *o.MassUnits
}

// GetMassUnitsOk returns a tuple with the MassUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetMassUnitsOk() (*string, bool) {
	if o == nil || o.MassUnits == nil {
		return nil, false
	}
	return o.MassUnits, true
}

// HasMassUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasMassUnits() bool {
	if o != nil && o.MassUnits != nil {
		return true
	}

	return false
}

// SetMassUnits gets a reference to the given string and assigns it to the MassUnits field.
func (o *BTDocumentElementProcessingInfo) SetMassUnits(v string) {
	o.MassUnits = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTDocumentElementProcessingInfo) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetMomentUnits returns the MomentUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetMomentUnits() string {
	if o == nil || o.MomentUnits == nil {
		var ret string
		return ret
	}
	return *o.MomentUnits
}

// GetMomentUnitsOk returns a tuple with the MomentUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetMomentUnitsOk() (*string, bool) {
	if o == nil || o.MomentUnits == nil {
		return nil, false
	}
	return o.MomentUnits, true
}

// HasMomentUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasMomentUnits() bool {
	if o != nil && o.MomentUnits != nil {
		return true
	}

	return false
}

// SetMomentUnits gets a reference to the given string and assigns it to the MomentUnits field.
func (o *BTDocumentElementProcessingInfo) SetMomentUnits(v string) {
	o.MomentUnits = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDocumentElementProcessingInfo) SetName(v string) {
	o.Name = &v
}

// GetPressureUnits returns the PressureUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetPressureUnits() string {
	if o == nil || o.PressureUnits == nil {
		var ret string
		return ret
	}
	return *o.PressureUnits
}

// GetPressureUnitsOk returns a tuple with the PressureUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetPressureUnitsOk() (*string, bool) {
	if o == nil || o.PressureUnits == nil {
		return nil, false
	}
	return o.PressureUnits, true
}

// HasPressureUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasPressureUnits() bool {
	if o != nil && o.PressureUnits != nil {
		return true
	}

	return false
}

// SetPressureUnits gets a reference to the given string and assigns it to the PressureUnits field.
func (o *BTDocumentElementProcessingInfo) SetPressureUnits(v string) {
	o.PressureUnits = &v
}

// GetPrettyType returns the PrettyType field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetPrettyType() string {
	if o == nil || o.PrettyType == nil {
		var ret string
		return ret
	}
	return *o.PrettyType
}

// GetPrettyTypeOk returns a tuple with the PrettyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetPrettyTypeOk() (*string, bool) {
	if o == nil || o.PrettyType == nil {
		return nil, false
	}
	return o.PrettyType, true
}

// HasPrettyType returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasPrettyType() bool {
	if o != nil && o.PrettyType != nil {
		return true
	}

	return false
}

// SetPrettyType gets a reference to the given string and assigns it to the PrettyType field.
func (o *BTDocumentElementProcessingInfo) SetPrettyType(v string) {
	o.PrettyType = &v
}

// GetSafeToShow returns the SafeToShow field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetSafeToShow() bool {
	if o == nil || o.SafeToShow == nil {
		var ret bool
		return ret
	}
	return *o.SafeToShow
}

// GetSafeToShowOk returns a tuple with the SafeToShow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetSafeToShowOk() (*bool, bool) {
	if o == nil || o.SafeToShow == nil {
		return nil, false
	}
	return o.SafeToShow, true
}

// HasSafeToShow returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasSafeToShow() bool {
	if o != nil && o.SafeToShow != nil {
		return true
	}

	return false
}

// SetSafeToShow gets a reference to the given bool and assigns it to the SafeToShow field.
func (o *BTDocumentElementProcessingInfo) SetSafeToShow(v bool) {
	o.SafeToShow = &v
}

// GetSpecifiedUnit returns the SpecifiedUnit field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetSpecifiedUnit() string {
	if o == nil || o.SpecifiedUnit == nil {
		var ret string
		return ret
	}
	return *o.SpecifiedUnit
}

// GetSpecifiedUnitOk returns a tuple with the SpecifiedUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetSpecifiedUnitOk() (*string, bool) {
	if o == nil || o.SpecifiedUnit == nil {
		return nil, false
	}
	return o.SpecifiedUnit, true
}

// HasSpecifiedUnit returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasSpecifiedUnit() bool {
	if o != nil && o.SpecifiedUnit != nil {
		return true
	}

	return false
}

// SetSpecifiedUnit gets a reference to the given string and assigns it to the SpecifiedUnit field.
func (o *BTDocumentElementProcessingInfo) SetSpecifiedUnit(v string) {
	o.SpecifiedUnit = &v
}

// GetThumbnailInfo returns the ThumbnailInfo field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetThumbnailInfo() BTThumbnailInfo {
	if o == nil || o.ThumbnailInfo == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.ThumbnailInfo
}

// GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.ThumbnailInfo == nil {
		return nil, false
	}
	return o.ThumbnailInfo, true
}

// HasThumbnailInfo returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasThumbnailInfo() bool {
	if o != nil && o.ThumbnailInfo != nil {
		return true
	}

	return false
}

// SetThumbnailInfo gets a reference to the given BTThumbnailInfo and assigns it to the ThumbnailInfo field.
func (o *BTDocumentElementProcessingInfo) SetThumbnailInfo(v BTThumbnailInfo) {
	o.ThumbnailInfo = &v
}

// GetThumbnails returns the Thumbnails field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetThumbnails() string {
	if o == nil || o.Thumbnails == nil {
		var ret string
		return ret
	}
	return *o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetThumbnailsOk() (*string, bool) {
	if o == nil || o.Thumbnails == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// HasThumbnails returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasThumbnails() bool {
	if o != nil && o.Thumbnails != nil {
		return true
	}

	return false
}

// SetThumbnails gets a reference to the given string and assigns it to the Thumbnails field.
func (o *BTDocumentElementProcessingInfo) SetThumbnails(v string) {
	o.Thumbnails = &v
}

// GetTimeUnits returns the TimeUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetTimeUnits() string {
	if o == nil || o.TimeUnits == nil {
		var ret string
		return ret
	}
	return *o.TimeUnits
}

// GetTimeUnitsOk returns a tuple with the TimeUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetTimeUnitsOk() (*string, bool) {
	if o == nil || o.TimeUnits == nil {
		return nil, false
	}
	return o.TimeUnits, true
}

// HasTimeUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasTimeUnits() bool {
	if o != nil && o.TimeUnits != nil {
		return true
	}

	return false
}

// SetTimeUnits gets a reference to the given string and assigns it to the TimeUnits field.
func (o *BTDocumentElementProcessingInfo) SetTimeUnits(v string) {
	o.TimeUnits = &v
}

// GetTranslationEventKey returns the TranslationEventKey field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetTranslationEventKey() string {
	if o == nil || o.TranslationEventKey == nil {
		var ret string
		return ret
	}
	return *o.TranslationEventKey
}

// GetTranslationEventKeyOk returns a tuple with the TranslationEventKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetTranslationEventKeyOk() (*string, bool) {
	if o == nil || o.TranslationEventKey == nil {
		return nil, false
	}
	return o.TranslationEventKey, true
}

// HasTranslationEventKey returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasTranslationEventKey() bool {
	if o != nil && o.TranslationEventKey != nil {
		return true
	}

	return false
}

// SetTranslationEventKey gets a reference to the given string and assigns it to the TranslationEventKey field.
func (o *BTDocumentElementProcessingInfo) SetTranslationEventKey(v string) {
	o.TranslationEventKey = &v
}

// GetTranslationId returns the TranslationId field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetTranslationId() string {
	if o == nil || o.TranslationId == nil {
		var ret string
		return ret
	}
	return *o.TranslationId
}

// GetTranslationIdOk returns a tuple with the TranslationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetTranslationIdOk() (*string, bool) {
	if o == nil || o.TranslationId == nil {
		return nil, false
	}
	return o.TranslationId, true
}

// HasTranslationId returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasTranslationId() bool {
	if o != nil && o.TranslationId != nil {
		return true
	}

	return false
}

// SetTranslationId gets a reference to the given string and assigns it to the TranslationId field.
func (o *BTDocumentElementProcessingInfo) SetTranslationId(v string) {
	o.TranslationId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTDocumentElementProcessingInfo) SetType(v string) {
	o.Type = &v
}

// GetUnupdatable returns the Unupdatable field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetUnupdatable() bool {
	if o == nil || o.Unupdatable == nil {
		var ret bool
		return ret
	}
	return *o.Unupdatable
}

// GetUnupdatableOk returns a tuple with the Unupdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetUnupdatableOk() (*bool, bool) {
	if o == nil || o.Unupdatable == nil {
		return nil, false
	}
	return o.Unupdatable, true
}

// HasUnupdatable returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasUnupdatable() bool {
	if o != nil && o.Unupdatable != nil {
		return true
	}

	return false
}

// SetUnupdatable gets a reference to the given bool and assigns it to the Unupdatable field.
func (o *BTDocumentElementProcessingInfo) SetUnupdatable(v bool) {
	o.Unupdatable = &v
}

// GetVolumeUnits returns the VolumeUnits field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetVolumeUnits() string {
	if o == nil || o.VolumeUnits == nil {
		var ret string
		return ret
	}
	return *o.VolumeUnits
}

// GetVolumeUnitsOk returns a tuple with the VolumeUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetVolumeUnitsOk() (*string, bool) {
	if o == nil || o.VolumeUnits == nil {
		return nil, false
	}
	return o.VolumeUnits, true
}

// HasVolumeUnits returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasVolumeUnits() bool {
	if o != nil && o.VolumeUnits != nil {
		return true
	}

	return false
}

// SetVolumeUnits gets a reference to the given string and assigns it to the VolumeUnits field.
func (o *BTDocumentElementProcessingInfo) SetVolumeUnits(v string) {
	o.VolumeUnits = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *BTDocumentElementProcessingInfo) GetZip() BTZipFileInfo {
	if o == nil || o.Zip == nil {
		var ret BTZipFileInfo
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementProcessingInfo) GetZipOk() (*BTZipFileInfo, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *BTDocumentElementProcessingInfo) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given BTZipFileInfo and assigns it to the Zip field.
func (o *BTDocumentElementProcessingInfo) SetZip(v BTZipFileInfo) {
	o.Zip = &v
}

func (o BTDocumentElementProcessingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccelerationUnits != nil {
		toSerialize["accelerationUnits"] = o.AccelerationUnits
	}
	if o.AngleUnits != nil {
		toSerialize["angleUnits"] = o.AngleUnits
	}
	if o.AngularVelocityUnits != nil {
		toSerialize["angularVelocityUnits"] = o.AngularVelocityUnits
	}
	if o.ApplicationTarget != nil {
		toSerialize["applicationTarget"] = o.ApplicationTarget
	}
	if o.AreaUnits != nil {
		toSerialize["areaUnits"] = o.AreaUnits
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.EnergyUnits != nil {
		toSerialize["energyUnits"] = o.EnergyUnits
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.ForceUnits != nil {
		toSerialize["forceUnits"] = o.ForceUnits
	}
	if o.ForeignDataId != nil {
		toSerialize["foreignDataId"] = o.ForeignDataId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LengthUnits != nil {
		toSerialize["lengthUnits"] = o.LengthUnits
	}
	if o.MassUnits != nil {
		toSerialize["massUnits"] = o.MassUnits
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MomentUnits != nil {
		toSerialize["momentUnits"] = o.MomentUnits
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PressureUnits != nil {
		toSerialize["pressureUnits"] = o.PressureUnits
	}
	if o.PrettyType != nil {
		toSerialize["prettyType"] = o.PrettyType
	}
	if o.SafeToShow != nil {
		toSerialize["safeToShow"] = o.SafeToShow
	}
	if o.SpecifiedUnit != nil {
		toSerialize["specifiedUnit"] = o.SpecifiedUnit
	}
	if o.ThumbnailInfo != nil {
		toSerialize["thumbnailInfo"] = o.ThumbnailInfo
	}
	if o.Thumbnails != nil {
		toSerialize["thumbnails"] = o.Thumbnails
	}
	if o.TimeUnits != nil {
		toSerialize["timeUnits"] = o.TimeUnits
	}
	if o.TranslationEventKey != nil {
		toSerialize["translationEventKey"] = o.TranslationEventKey
	}
	if o.TranslationId != nil {
		toSerialize["translationId"] = o.TranslationId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Unupdatable != nil {
		toSerialize["unupdatable"] = o.Unupdatable
	}
	if o.VolumeUnits != nil {
		toSerialize["volumeUnits"] = o.VolumeUnits
	}
	if o.Zip != nil {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentElementProcessingInfo struct {
	value *BTDocumentElementProcessingInfo
	isSet bool
}

func (v NullableBTDocumentElementProcessingInfo) Get() *BTDocumentElementProcessingInfo {
	return v.value
}

func (v *NullableBTDocumentElementProcessingInfo) Set(val *BTDocumentElementProcessingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentElementProcessingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentElementProcessingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentElementProcessingInfo(val *BTDocumentElementProcessingInfo) *NullableBTDocumentElementProcessingInfo {
	return &NullableBTDocumentElementProcessingInfo{value: val, isSet: true}
}

func (v NullableBTDocumentElementProcessingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentElementProcessingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
