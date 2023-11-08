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

// BTPartMetadataInfo struct for BTPartMetadataInfo
type BTPartMetadataInfo struct {
	Appearance                       *BTPartAppearanceInfo `json:"appearance,omitempty"`
	BodyType                         *string               `json:"bodyType,omitempty"`
	ConfigurationId                  *string               `json:"configurationId,omitempty"`
	CustomProperties                 *map[string]string    `json:"customProperties,omitempty"`
	DefaultColorHash                 *string               `json:"defaultColorHash,omitempty"`
	Description                      *string               `json:"description,omitempty"`
	ElementId                        *string               `json:"elementId,omitempty"`
	Href                             *string               `json:"href,omitempty"`
	Id                               *string               `json:"id,omitempty"`
	IsFlattenedBody                  *bool                 `json:"isFlattenedBody,omitempty"`
	IsHidden                         *bool                 `json:"isHidden,omitempty"`
	IsMesh                           *bool                 `json:"isMesh,omitempty"`
	Material                         *BTPartMaterialInfo   `json:"material,omitempty"`
	MeshState                        *GBTMeshState         `json:"meshState,omitempty"`
	MicroversionId                   *string               `json:"microversionId,omitempty"`
	Name                             *string               `json:"name,omitempty"`
	Ordinal                          *int32                `json:"ordinal,omitempty"`
	PartId                           *string               `json:"partId,omitempty"`
	PartIdentity                     *string               `json:"partIdentity,omitempty"`
	PartNumber                       *string               `json:"partNumber,omitempty"`
	PartQuery                        *string               `json:"partQuery,omitempty"`
	ProductLine                      *string               `json:"productLine,omitempty"`
	Project                          *string               `json:"project,omitempty"`
	PropertySourceTypes              *map[string]int32     `json:"propertySourceTypes,omitempty"`
	ReferencingConfiguredPartNodeIds []string              `json:"referencingConfiguredPartNodeIds,omitempty"`
	Revision                         *string               `json:"revision,omitempty"`
	State                            *BTMetadataStateType  `json:"state,omitempty"`
	ThumbnailConfigurationId         *string               `json:"thumbnailConfigurationId,omitempty"`
	ThumbnailInfo                    *BTThumbnailInfo      `json:"thumbnailInfo,omitempty"`
	Title1                           *string               `json:"title1,omitempty"`
	Title2                           *string               `json:"title2,omitempty"`
	Title3                           *string               `json:"title3,omitempty"`
	UnflattenedPartId                *string               `json:"unflattenedPartId,omitempty"`
	Vendor                           *string               `json:"vendor,omitempty"`
}

// NewBTPartMetadataInfo instantiates a new BTPartMetadataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartMetadataInfo() *BTPartMetadataInfo {
	this := BTPartMetadataInfo{}
	return &this
}

// NewBTPartMetadataInfoWithDefaults instantiates a new BTPartMetadataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartMetadataInfoWithDefaults() *BTPartMetadataInfo {
	this := BTPartMetadataInfo{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetAppearance() BTPartAppearanceInfo {
	if o == nil || o.Appearance == nil {
		var ret BTPartAppearanceInfo
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetAppearanceOk() (*BTPartAppearanceInfo, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTPartAppearanceInfo and assigns it to the Appearance field.
func (o *BTPartMetadataInfo) SetAppearance(v BTPartAppearanceInfo) {
	o.Appearance = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *BTPartMetadataInfo) SetBodyType(v string) {
	o.BodyType = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetConfigurationId() string {
	if o == nil || o.ConfigurationId == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetConfigurationIdOk() (*string, bool) {
	if o == nil || o.ConfigurationId == nil {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId != nil {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *BTPartMetadataInfo) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetCustomProperties() map[string]string {
	if o == nil || o.CustomProperties == nil {
		var ret map[string]string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetCustomPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]string and assigns it to the CustomProperties field.
func (o *BTPartMetadataInfo) SetCustomProperties(v map[string]string) {
	o.CustomProperties = &v
}

// GetDefaultColorHash returns the DefaultColorHash field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetDefaultColorHash() string {
	if o == nil || o.DefaultColorHash == nil {
		var ret string
		return ret
	}
	return *o.DefaultColorHash
}

// GetDefaultColorHashOk returns a tuple with the DefaultColorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetDefaultColorHashOk() (*string, bool) {
	if o == nil || o.DefaultColorHash == nil {
		return nil, false
	}
	return o.DefaultColorHash, true
}

// HasDefaultColorHash returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasDefaultColorHash() bool {
	if o != nil && o.DefaultColorHash != nil {
		return true
	}

	return false
}

// SetDefaultColorHash gets a reference to the given string and assigns it to the DefaultColorHash field.
func (o *BTPartMetadataInfo) SetDefaultColorHash(v string) {
	o.DefaultColorHash = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTPartMetadataInfo) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTPartMetadataInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTPartMetadataInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPartMetadataInfo) SetId(v string) {
	o.Id = &v
}

// GetIsFlattenedBody returns the IsFlattenedBody field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetIsFlattenedBody() bool {
	if o == nil || o.IsFlattenedBody == nil {
		var ret bool
		return ret
	}
	return *o.IsFlattenedBody
}

// GetIsFlattenedBodyOk returns a tuple with the IsFlattenedBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetIsFlattenedBodyOk() (*bool, bool) {
	if o == nil || o.IsFlattenedBody == nil {
		return nil, false
	}
	return o.IsFlattenedBody, true
}

// HasIsFlattenedBody returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasIsFlattenedBody() bool {
	if o != nil && o.IsFlattenedBody != nil {
		return true
	}

	return false
}

// SetIsFlattenedBody gets a reference to the given bool and assigns it to the IsFlattenedBody field.
func (o *BTPartMetadataInfo) SetIsFlattenedBody(v bool) {
	o.IsFlattenedBody = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasIsHidden() bool {
	if o != nil && o.IsHidden != nil {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *BTPartMetadataInfo) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetIsMesh returns the IsMesh field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetIsMesh() bool {
	if o == nil || o.IsMesh == nil {
		var ret bool
		return ret
	}
	return *o.IsMesh
}

// GetIsMeshOk returns a tuple with the IsMesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetIsMeshOk() (*bool, bool) {
	if o == nil || o.IsMesh == nil {
		return nil, false
	}
	return o.IsMesh, true
}

// HasIsMesh returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasIsMesh() bool {
	if o != nil && o.IsMesh != nil {
		return true
	}

	return false
}

// SetIsMesh gets a reference to the given bool and assigns it to the IsMesh field.
func (o *BTPartMetadataInfo) SetIsMesh(v bool) {
	o.IsMesh = &v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetMaterial() BTPartMaterialInfo {
	if o == nil || o.Material == nil {
		var ret BTPartMaterialInfo
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetMaterialOk() (*BTPartMaterialInfo, bool) {
	if o == nil || o.Material == nil {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasMaterial() bool {
	if o != nil && o.Material != nil {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given BTPartMaterialInfo and assigns it to the Material field.
func (o *BTPartMetadataInfo) SetMaterial(v BTPartMaterialInfo) {
	o.Material = &v
}

// GetMeshState returns the MeshState field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetMeshState() GBTMeshState {
	if o == nil || o.MeshState == nil {
		var ret GBTMeshState
		return ret
	}
	return *o.MeshState
}

// GetMeshStateOk returns a tuple with the MeshState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetMeshStateOk() (*GBTMeshState, bool) {
	if o == nil || o.MeshState == nil {
		return nil, false
	}
	return o.MeshState, true
}

// HasMeshState returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasMeshState() bool {
	if o != nil && o.MeshState != nil {
		return true
	}

	return false
}

// SetMeshState gets a reference to the given GBTMeshState and assigns it to the MeshState field.
func (o *BTPartMetadataInfo) SetMeshState(v GBTMeshState) {
	o.MeshState = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTPartMetadataInfo) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPartMetadataInfo) SetName(v string) {
	o.Name = &v
}

// GetOrdinal returns the Ordinal field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetOrdinal() int32 {
	if o == nil || o.Ordinal == nil {
		var ret int32
		return ret
	}
	return *o.Ordinal
}

// GetOrdinalOk returns a tuple with the Ordinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetOrdinalOk() (*int32, bool) {
	if o == nil || o.Ordinal == nil {
		return nil, false
	}
	return o.Ordinal, true
}

// HasOrdinal returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasOrdinal() bool {
	if o != nil && o.Ordinal != nil {
		return true
	}

	return false
}

// SetOrdinal gets a reference to the given int32 and assigns it to the Ordinal field.
func (o *BTPartMetadataInfo) SetOrdinal(v int32) {
	o.Ordinal = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTPartMetadataInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetPartIdentity() string {
	if o == nil || o.PartIdentity == nil {
		var ret string
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetPartIdentityOk() (*string, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given string and assigns it to the PartIdentity field.
func (o *BTPartMetadataInfo) SetPartIdentity(v string) {
	o.PartIdentity = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTPartMetadataInfo) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPartQuery returns the PartQuery field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetPartQuery() string {
	if o == nil || o.PartQuery == nil {
		var ret string
		return ret
	}
	return *o.PartQuery
}

// GetPartQueryOk returns a tuple with the PartQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetPartQueryOk() (*string, bool) {
	if o == nil || o.PartQuery == nil {
		return nil, false
	}
	return o.PartQuery, true
}

// HasPartQuery returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasPartQuery() bool {
	if o != nil && o.PartQuery != nil {
		return true
	}

	return false
}

// SetPartQuery gets a reference to the given string and assigns it to the PartQuery field.
func (o *BTPartMetadataInfo) SetPartQuery(v string) {
	o.PartQuery = &v
}

// GetProductLine returns the ProductLine field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetProductLine() string {
	if o == nil || o.ProductLine == nil {
		var ret string
		return ret
	}
	return *o.ProductLine
}

// GetProductLineOk returns a tuple with the ProductLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetProductLineOk() (*string, bool) {
	if o == nil || o.ProductLine == nil {
		return nil, false
	}
	return o.ProductLine, true
}

// HasProductLine returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasProductLine() bool {
	if o != nil && o.ProductLine != nil {
		return true
	}

	return false
}

// SetProductLine gets a reference to the given string and assigns it to the ProductLine field.
func (o *BTPartMetadataInfo) SetProductLine(v string) {
	o.ProductLine = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *BTPartMetadataInfo) SetProject(v string) {
	o.Project = &v
}

// GetPropertySourceTypes returns the PropertySourceTypes field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetPropertySourceTypes() map[string]int32 {
	if o == nil || o.PropertySourceTypes == nil {
		var ret map[string]int32
		return ret
	}
	return *o.PropertySourceTypes
}

// GetPropertySourceTypesOk returns a tuple with the PropertySourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetPropertySourceTypesOk() (*map[string]int32, bool) {
	if o == nil || o.PropertySourceTypes == nil {
		return nil, false
	}
	return o.PropertySourceTypes, true
}

// HasPropertySourceTypes returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasPropertySourceTypes() bool {
	if o != nil && o.PropertySourceTypes != nil {
		return true
	}

	return false
}

// SetPropertySourceTypes gets a reference to the given map[string]int32 and assigns it to the PropertySourceTypes field.
func (o *BTPartMetadataInfo) SetPropertySourceTypes(v map[string]int32) {
	o.PropertySourceTypes = &v
}

// GetReferencingConfiguredPartNodeIds returns the ReferencingConfiguredPartNodeIds field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetReferencingConfiguredPartNodeIds() []string {
	if o == nil || o.ReferencingConfiguredPartNodeIds == nil {
		var ret []string
		return ret
	}
	return o.ReferencingConfiguredPartNodeIds
}

// GetReferencingConfiguredPartNodeIdsOk returns a tuple with the ReferencingConfiguredPartNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetReferencingConfiguredPartNodeIdsOk() ([]string, bool) {
	if o == nil || o.ReferencingConfiguredPartNodeIds == nil {
		return nil, false
	}
	return o.ReferencingConfiguredPartNodeIds, true
}

// HasReferencingConfiguredPartNodeIds returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasReferencingConfiguredPartNodeIds() bool {
	if o != nil && o.ReferencingConfiguredPartNodeIds != nil {
		return true
	}

	return false
}

// SetReferencingConfiguredPartNodeIds gets a reference to the given []string and assigns it to the ReferencingConfiguredPartNodeIds field.
func (o *BTPartMetadataInfo) SetReferencingConfiguredPartNodeIds(v []string) {
	o.ReferencingConfiguredPartNodeIds = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTPartMetadataInfo) SetRevision(v string) {
	o.Revision = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetState() BTMetadataStateType {
	if o == nil || o.State == nil {
		var ret BTMetadataStateType
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetStateOk() (*BTMetadataStateType, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given BTMetadataStateType and assigns it to the State field.
func (o *BTPartMetadataInfo) SetState(v BTMetadataStateType) {
	o.State = &v
}

// GetThumbnailConfigurationId returns the ThumbnailConfigurationId field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetThumbnailConfigurationId() string {
	if o == nil || o.ThumbnailConfigurationId == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailConfigurationId
}

// GetThumbnailConfigurationIdOk returns a tuple with the ThumbnailConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetThumbnailConfigurationIdOk() (*string, bool) {
	if o == nil || o.ThumbnailConfigurationId == nil {
		return nil, false
	}
	return o.ThumbnailConfigurationId, true
}

// HasThumbnailConfigurationId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasThumbnailConfigurationId() bool {
	if o != nil && o.ThumbnailConfigurationId != nil {
		return true
	}

	return false
}

// SetThumbnailConfigurationId gets a reference to the given string and assigns it to the ThumbnailConfigurationId field.
func (o *BTPartMetadataInfo) SetThumbnailConfigurationId(v string) {
	o.ThumbnailConfigurationId = &v
}

// GetThumbnailInfo returns the ThumbnailInfo field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetThumbnailInfo() BTThumbnailInfo {
	if o == nil || o.ThumbnailInfo == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.ThumbnailInfo
}

// GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.ThumbnailInfo == nil {
		return nil, false
	}
	return o.ThumbnailInfo, true
}

// HasThumbnailInfo returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasThumbnailInfo() bool {
	if o != nil && o.ThumbnailInfo != nil {
		return true
	}

	return false
}

// SetThumbnailInfo gets a reference to the given BTThumbnailInfo and assigns it to the ThumbnailInfo field.
func (o *BTPartMetadataInfo) SetThumbnailInfo(v BTThumbnailInfo) {
	o.ThumbnailInfo = &v
}

// GetTitle1 returns the Title1 field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetTitle1() string {
	if o == nil || o.Title1 == nil {
		var ret string
		return ret
	}
	return *o.Title1
}

// GetTitle1Ok returns a tuple with the Title1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetTitle1Ok() (*string, bool) {
	if o == nil || o.Title1 == nil {
		return nil, false
	}
	return o.Title1, true
}

// HasTitle1 returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasTitle1() bool {
	if o != nil && o.Title1 != nil {
		return true
	}

	return false
}

// SetTitle1 gets a reference to the given string and assigns it to the Title1 field.
func (o *BTPartMetadataInfo) SetTitle1(v string) {
	o.Title1 = &v
}

// GetTitle2 returns the Title2 field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetTitle2() string {
	if o == nil || o.Title2 == nil {
		var ret string
		return ret
	}
	return *o.Title2
}

// GetTitle2Ok returns a tuple with the Title2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetTitle2Ok() (*string, bool) {
	if o == nil || o.Title2 == nil {
		return nil, false
	}
	return o.Title2, true
}

// HasTitle2 returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasTitle2() bool {
	if o != nil && o.Title2 != nil {
		return true
	}

	return false
}

// SetTitle2 gets a reference to the given string and assigns it to the Title2 field.
func (o *BTPartMetadataInfo) SetTitle2(v string) {
	o.Title2 = &v
}

// GetTitle3 returns the Title3 field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetTitle3() string {
	if o == nil || o.Title3 == nil {
		var ret string
		return ret
	}
	return *o.Title3
}

// GetTitle3Ok returns a tuple with the Title3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetTitle3Ok() (*string, bool) {
	if o == nil || o.Title3 == nil {
		return nil, false
	}
	return o.Title3, true
}

// HasTitle3 returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasTitle3() bool {
	if o != nil && o.Title3 != nil {
		return true
	}

	return false
}

// SetTitle3 gets a reference to the given string and assigns it to the Title3 field.
func (o *BTPartMetadataInfo) SetTitle3(v string) {
	o.Title3 = &v
}

// GetUnflattenedPartId returns the UnflattenedPartId field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetUnflattenedPartId() string {
	if o == nil || o.UnflattenedPartId == nil {
		var ret string
		return ret
	}
	return *o.UnflattenedPartId
}

// GetUnflattenedPartIdOk returns a tuple with the UnflattenedPartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetUnflattenedPartIdOk() (*string, bool) {
	if o == nil || o.UnflattenedPartId == nil {
		return nil, false
	}
	return o.UnflattenedPartId, true
}

// HasUnflattenedPartId returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasUnflattenedPartId() bool {
	if o != nil && o.UnflattenedPartId != nil {
		return true
	}

	return false
}

// SetUnflattenedPartId gets a reference to the given string and assigns it to the UnflattenedPartId field.
func (o *BTPartMetadataInfo) SetUnflattenedPartId(v string) {
	o.UnflattenedPartId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BTPartMetadataInfo) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataInfo) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BTPartMetadataInfo) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BTPartMetadataInfo) SetVendor(v string) {
	o.Vendor = &v
}

func (o BTPartMetadataInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.ConfigurationId != nil {
		toSerialize["configurationId"] = o.ConfigurationId
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if o.DefaultColorHash != nil {
		toSerialize["defaultColorHash"] = o.DefaultColorHash
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsFlattenedBody != nil {
		toSerialize["isFlattenedBody"] = o.IsFlattenedBody
	}
	if o.IsHidden != nil {
		toSerialize["isHidden"] = o.IsHidden
	}
	if o.IsMesh != nil {
		toSerialize["isMesh"] = o.IsMesh
	}
	if o.Material != nil {
		toSerialize["material"] = o.Material
	}
	if o.MeshState != nil {
		toSerialize["meshState"] = o.MeshState
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ordinal != nil {
		toSerialize["ordinal"] = o.Ordinal
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartIdentity != nil {
		toSerialize["partIdentity"] = o.PartIdentity
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.PartQuery != nil {
		toSerialize["partQuery"] = o.PartQuery
	}
	if o.ProductLine != nil {
		toSerialize["productLine"] = o.ProductLine
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.PropertySourceTypes != nil {
		toSerialize["propertySourceTypes"] = o.PropertySourceTypes
	}
	if o.ReferencingConfiguredPartNodeIds != nil {
		toSerialize["referencingConfiguredPartNodeIds"] = o.ReferencingConfiguredPartNodeIds
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ThumbnailConfigurationId != nil {
		toSerialize["thumbnailConfigurationId"] = o.ThumbnailConfigurationId
	}
	if o.ThumbnailInfo != nil {
		toSerialize["thumbnailInfo"] = o.ThumbnailInfo
	}
	if o.Title1 != nil {
		toSerialize["title1"] = o.Title1
	}
	if o.Title2 != nil {
		toSerialize["title2"] = o.Title2
	}
	if o.Title3 != nil {
		toSerialize["title3"] = o.Title3
	}
	if o.UnflattenedPartId != nil {
		toSerialize["unflattenedPartId"] = o.UnflattenedPartId
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartMetadataInfo struct {
	value *BTPartMetadataInfo
	isSet bool
}

func (v NullableBTPartMetadataInfo) Get() *BTPartMetadataInfo {
	return v.value
}

func (v *NullableBTPartMetadataInfo) Set(val *BTPartMetadataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartMetadataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartMetadataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartMetadataInfo(val *BTPartMetadataInfo) *NullableBTPartMetadataInfo {
	return &NullableBTPartMetadataInfo{value: val, isSet: true}
}

func (v NullableBTPartMetadataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartMetadataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
