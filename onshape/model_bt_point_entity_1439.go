/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26754-ceeaad064d4a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPointEntity1439 struct for BTPointEntity1439
type BTPointEntity1439 struct {
	BtType                 *string                       `json:"btType,omitempty"`
	CopyWithoutGeometry    *BTBaseEntityData33           `json:"copyWithoutGeometry,omitempty"`
	Decompressed           *BTBaseEntityData33           `json:"decompressed,omitempty"`
	Deletion               *bool                         `json:"deletion,omitempty"`
	FeatureIds             []string                      `json:"featureIds,omitempty"`
	FromSketch             *bool                         `json:"fromSketch,omitempty"`
	Geometries             []BTEntityGeometry35          `json:"geometries,omitempty"`
	DomainSpecificMetadata []BTDomainSpecificMetadata961 `json:"domainSpecificMetadata,omitempty"`
	FirstGeometry          *BTEntityGeometry35           `json:"firstGeometry,omitempty"`
}

// NewBTPointEntity1439 instantiates a new BTPointEntity1439 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPointEntity1439() *BTPointEntity1439 {
	this := BTPointEntity1439{}
	return &this
}

// NewBTPointEntity1439WithDefaults instantiates a new BTPointEntity1439 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPointEntity1439WithDefaults() *BTPointEntity1439 {
	this := BTPointEntity1439{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPointEntity1439) SetBtType(v string) {
	o.BtType = &v
}

// GetCopyWithoutGeometry returns the CopyWithoutGeometry field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetCopyWithoutGeometry() BTBaseEntityData33 {
	if o == nil || o.CopyWithoutGeometry == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.CopyWithoutGeometry
}

// GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.CopyWithoutGeometry == nil {
		return nil, false
	}
	return o.CopyWithoutGeometry, true
}

// HasCopyWithoutGeometry returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasCopyWithoutGeometry() bool {
	if o != nil && o.CopyWithoutGeometry != nil {
		return true
	}

	return false
}

// SetCopyWithoutGeometry gets a reference to the given BTBaseEntityData33 and assigns it to the CopyWithoutGeometry field.
func (o *BTPointEntity1439) SetCopyWithoutGeometry(v BTBaseEntityData33) {
	o.CopyWithoutGeometry = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetDecompressed() BTBaseEntityData33 {
	if o == nil || o.Decompressed == nil {
		var ret BTBaseEntityData33
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetDecompressedOk() (*BTBaseEntityData33, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTBaseEntityData33 and assigns it to the Decompressed field.
func (o *BTPointEntity1439) SetDecompressed(v BTBaseEntityData33) {
	o.Decompressed = &v
}

// GetDeletion returns the Deletion field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetDeletion() bool {
	if o == nil || o.Deletion == nil {
		var ret bool
		return ret
	}
	return *o.Deletion
}

// GetDeletionOk returns a tuple with the Deletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetDeletionOk() (*bool, bool) {
	if o == nil || o.Deletion == nil {
		return nil, false
	}
	return o.Deletion, true
}

// HasDeletion returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasDeletion() bool {
	if o != nil && o.Deletion != nil {
		return true
	}

	return false
}

// SetDeletion gets a reference to the given bool and assigns it to the Deletion field.
func (o *BTPointEntity1439) SetDeletion(v bool) {
	o.Deletion = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasFeatureIds() bool {
	if o != nil && o.FeatureIds != nil {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *BTPointEntity1439) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetFromSketch returns the FromSketch field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetFromSketch() bool {
	if o == nil || o.FromSketch == nil {
		var ret bool
		return ret
	}
	return *o.FromSketch
}

// GetFromSketchOk returns a tuple with the FromSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetFromSketchOk() (*bool, bool) {
	if o == nil || o.FromSketch == nil {
		return nil, false
	}
	return o.FromSketch, true
}

// HasFromSketch returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasFromSketch() bool {
	if o != nil && o.FromSketch != nil {
		return true
	}

	return false
}

// SetFromSketch gets a reference to the given bool and assigns it to the FromSketch field.
func (o *BTPointEntity1439) SetFromSketch(v bool) {
	o.FromSketch = &v
}

// GetGeometries returns the Geometries field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetGeometries() []BTEntityGeometry35 {
	if o == nil || o.Geometries == nil {
		var ret []BTEntityGeometry35
		return ret
	}
	return o.Geometries
}

// GetGeometriesOk returns a tuple with the Geometries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetGeometriesOk() ([]BTEntityGeometry35, bool) {
	if o == nil || o.Geometries == nil {
		return nil, false
	}
	return o.Geometries, true
}

// HasGeometries returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasGeometries() bool {
	if o != nil && o.Geometries != nil {
		return true
	}

	return false
}

// SetGeometries gets a reference to the given []BTEntityGeometry35 and assigns it to the Geometries field.
func (o *BTPointEntity1439) SetGeometries(v []BTEntityGeometry35) {
	o.Geometries = v
}

// GetDomainSpecificMetadata returns the DomainSpecificMetadata field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetDomainSpecificMetadata() []BTDomainSpecificMetadata961 {
	if o == nil || o.DomainSpecificMetadata == nil {
		var ret []BTDomainSpecificMetadata961
		return ret
	}
	return o.DomainSpecificMetadata
}

// GetDomainSpecificMetadataOk returns a tuple with the DomainSpecificMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetDomainSpecificMetadataOk() ([]BTDomainSpecificMetadata961, bool) {
	if o == nil || o.DomainSpecificMetadata == nil {
		return nil, false
	}
	return o.DomainSpecificMetadata, true
}

// HasDomainSpecificMetadata returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasDomainSpecificMetadata() bool {
	if o != nil && o.DomainSpecificMetadata != nil {
		return true
	}

	return false
}

// SetDomainSpecificMetadata gets a reference to the given []BTDomainSpecificMetadata961 and assigns it to the DomainSpecificMetadata field.
func (o *BTPointEntity1439) SetDomainSpecificMetadata(v []BTDomainSpecificMetadata961) {
	o.DomainSpecificMetadata = v
}

// GetFirstGeometry returns the FirstGeometry field value if set, zero value otherwise.
func (o *BTPointEntity1439) GetFirstGeometry() BTEntityGeometry35 {
	if o == nil || o.FirstGeometry == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.FirstGeometry
}

// GetFirstGeometryOk returns a tuple with the FirstGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPointEntity1439) GetFirstGeometryOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.FirstGeometry == nil {
		return nil, false
	}
	return o.FirstGeometry, true
}

// HasFirstGeometry returns a boolean if a field has been set.
func (o *BTPointEntity1439) HasFirstGeometry() bool {
	if o != nil && o.FirstGeometry != nil {
		return true
	}

	return false
}

// SetFirstGeometry gets a reference to the given BTEntityGeometry35 and assigns it to the FirstGeometry field.
func (o *BTPointEntity1439) SetFirstGeometry(v BTEntityGeometry35) {
	o.FirstGeometry = &v
}

func (o BTPointEntity1439) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableBTPointEntity1439 struct {
	value *BTPointEntity1439
	isSet bool
}

func (v NullableBTPointEntity1439) Get() *BTPointEntity1439 {
	return v.value
}

func (v *NullableBTPointEntity1439) Set(val *BTPointEntity1439) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPointEntity1439) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPointEntity1439) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPointEntity1439(val *BTPointEntity1439) *NullableBTPointEntity1439 {
	return &NullableBTPointEntity1439{value: val, isSet: true}
}

func (v NullableBTPointEntity1439) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPointEntity1439) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
