/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7108-42dac6f29840
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConstructionPlaneEntity27 struct for BTConstructionPlaneEntity27
type BTConstructionPlaneEntity27 struct {
	BtType                 *string                       `json:"btType,omitempty"`
	CopyWithoutGeometry    *BTBaseEntityData33           `json:"copyWithoutGeometry,omitempty"`
	Decompressed           *BTBaseEntityData33           `json:"decompressed,omitempty"`
	Deletion               *bool                         `json:"deletion,omitempty"`
	FeatureIds             []string                      `json:"featureIds,omitempty"`
	FromSketch             *bool                         `json:"fromSketch,omitempty"`
	Geometries             []BTEntityGeometry35          `json:"geometries,omitempty"`
	DomainSpecificMetadata []BTDomainSpecificMetadata961 `json:"domainSpecificMetadata,omitempty"`
	FirstGeometry          *BTEntityGeometry35           `json:"firstGeometry,omitempty"`
	IsDefault              *bool                         `json:"isDefault,omitempty"`
}

// NewBTConstructionPlaneEntity27 instantiates a new BTConstructionPlaneEntity27 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConstructionPlaneEntity27() *BTConstructionPlaneEntity27 {
	this := BTConstructionPlaneEntity27{}
	return &this
}

// NewBTConstructionPlaneEntity27WithDefaults instantiates a new BTConstructionPlaneEntity27 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConstructionPlaneEntity27WithDefaults() *BTConstructionPlaneEntity27 {
	this := BTConstructionPlaneEntity27{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConstructionPlaneEntity27) SetBtType(v string) {
	o.BtType = &v
}

// GetCopyWithoutGeometry returns the CopyWithoutGeometry field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetCopyWithoutGeometry() BTBaseEntityData33 {
	if o == nil || o.CopyWithoutGeometry == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.CopyWithoutGeometry
}

// GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.CopyWithoutGeometry == nil {
		return nil, false
	}
	return o.CopyWithoutGeometry, true
}

// HasCopyWithoutGeometry returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasCopyWithoutGeometry() bool {
	if o != nil && o.CopyWithoutGeometry != nil {
		return true
	}

	return false
}

// SetCopyWithoutGeometry gets a reference to the given BTBaseEntityData33 and assigns it to the CopyWithoutGeometry field.
func (o *BTConstructionPlaneEntity27) SetCopyWithoutGeometry(v BTBaseEntityData33) {
	o.CopyWithoutGeometry = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetDecompressed() BTBaseEntityData33 {
	if o == nil || o.Decompressed == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetDecompressedOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTBaseEntityData33 and assigns it to the Decompressed field.
func (o *BTConstructionPlaneEntity27) SetDecompressed(v BTBaseEntityData33) {
	o.Decompressed = &v
}

// GetDeletion returns the Deletion field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetDeletion() bool {
	if o == nil || o.Deletion == nil {
		var ret bool
		return ret
	}
	return *o.Deletion
}

// GetDeletionOk returns a tuple with the Deletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetDeletionOk() (*bool, bool) {
	if o == nil || o.Deletion == nil {
		return nil, false
	}
	return o.Deletion, true
}

// HasDeletion returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasDeletion() bool {
	if o != nil && o.Deletion != nil {
		return true
	}

	return false
}

// SetDeletion gets a reference to the given bool and assigns it to the Deletion field.
func (o *BTConstructionPlaneEntity27) SetDeletion(v bool) {
	o.Deletion = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasFeatureIds() bool {
	if o != nil && o.FeatureIds != nil {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *BTConstructionPlaneEntity27) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetFromSketch returns the FromSketch field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetFromSketch() bool {
	if o == nil || o.FromSketch == nil {
		var ret bool
		return ret
	}
	return *o.FromSketch
}

// GetFromSketchOk returns a tuple with the FromSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetFromSketchOk() (*bool, bool) {
	if o == nil || o.FromSketch == nil {
		return nil, false
	}
	return o.FromSketch, true
}

// HasFromSketch returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasFromSketch() bool {
	if o != nil && o.FromSketch != nil {
		return true
	}

	return false
}

// SetFromSketch gets a reference to the given bool and assigns it to the FromSketch field.
func (o *BTConstructionPlaneEntity27) SetFromSketch(v bool) {
	o.FromSketch = &v
}

// GetGeometries returns the Geometries field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetGeometries() []BTEntityGeometry35 {
	if o == nil || o.Geometries == nil {
		var ret []BTEntityGeometry35
		return ret
	}
	return o.Geometries
}

// GetGeometriesOk returns a tuple with the Geometries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetGeometriesOk() ([]BTEntityGeometry35, bool) {
	if o == nil || o.Geometries == nil {
		return nil, false
	}
	return o.Geometries, true
}

// HasGeometries returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasGeometries() bool {
	if o != nil && o.Geometries != nil {
		return true
	}

	return false
}

// SetGeometries gets a reference to the given []BTEntityGeometry35 and assigns it to the Geometries field.
func (o *BTConstructionPlaneEntity27) SetGeometries(v []BTEntityGeometry35) {
	o.Geometries = v
}

// GetDomainSpecificMetadata returns the DomainSpecificMetadata field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetDomainSpecificMetadata() []BTDomainSpecificMetadata961 {
	if o == nil || o.DomainSpecificMetadata == nil {
		var ret []BTDomainSpecificMetadata961
		return ret
	}
	return o.DomainSpecificMetadata
}

// GetDomainSpecificMetadataOk returns a tuple with the DomainSpecificMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetDomainSpecificMetadataOk() ([]BTDomainSpecificMetadata961, bool) {
	if o == nil || o.DomainSpecificMetadata == nil {
		return nil, false
	}
	return o.DomainSpecificMetadata, true
}

// HasDomainSpecificMetadata returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasDomainSpecificMetadata() bool {
	if o != nil && o.DomainSpecificMetadata != nil {
		return true
	}

	return false
}

// SetDomainSpecificMetadata gets a reference to the given []BTDomainSpecificMetadata961 and assigns it to the DomainSpecificMetadata field.
func (o *BTConstructionPlaneEntity27) SetDomainSpecificMetadata(v []BTDomainSpecificMetadata961) {
	o.DomainSpecificMetadata = v
}

// GetFirstGeometry returns the FirstGeometry field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetFirstGeometry() BTEntityGeometry35 {
	if o == nil || o.FirstGeometry == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.FirstGeometry
}

// GetFirstGeometryOk returns a tuple with the FirstGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetFirstGeometryOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.FirstGeometry == nil {
		return nil, false
	}
	return o.FirstGeometry, true
}

// HasFirstGeometry returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasFirstGeometry() bool {
	if o != nil && o.FirstGeometry != nil {
		return true
	}

	return false
}

// SetFirstGeometry gets a reference to the given BTEntityGeometry35 and assigns it to the FirstGeometry field.
func (o *BTConstructionPlaneEntity27) SetFirstGeometry(v BTEntityGeometry35) {
	o.FirstGeometry = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *BTConstructionPlaneEntity27) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionPlaneEntity27) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *BTConstructionPlaneEntity27) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *BTConstructionPlaneEntity27) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o BTConstructionPlaneEntity27) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CopyWithoutGeometry != nil {
		toSerialize["copyWithoutGeometry"] = o.CopyWithoutGeometry
	}
	if o.Decompressed != nil {
		toSerialize["decompressed"] = o.Decompressed
	}
	if o.Deletion != nil {
		toSerialize["deletion"] = o.Deletion
	}
	if o.FeatureIds != nil {
		toSerialize["featureIds"] = o.FeatureIds
	}
	if o.FromSketch != nil {
		toSerialize["fromSketch"] = o.FromSketch
	}
	if o.Geometries != nil {
		toSerialize["geometries"] = o.Geometries
	}
	if o.DomainSpecificMetadata != nil {
		toSerialize["domainSpecificMetadata"] = o.DomainSpecificMetadata
	}
	if o.FirstGeometry != nil {
		toSerialize["firstGeometry"] = o.FirstGeometry
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	return json.Marshal(toSerialize)
}

type NullableBTConstructionPlaneEntity27 struct {
	value *BTConstructionPlaneEntity27
	isSet bool
}

func (v NullableBTConstructionPlaneEntity27) Get() *BTConstructionPlaneEntity27 {
	return v.value
}

func (v *NullableBTConstructionPlaneEntity27) Set(val *BTConstructionPlaneEntity27) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConstructionPlaneEntity27) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConstructionPlaneEntity27) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConstructionPlaneEntity27(val *BTConstructionPlaneEntity27) *NullableBTConstructionPlaneEntity27 {
	return &NullableBTConstructionPlaneEntity27{value: val, isSet: true}
}

func (v NullableBTConstructionPlaneEntity27) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConstructionPlaneEntity27) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
