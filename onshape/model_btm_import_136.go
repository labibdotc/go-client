/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMImport136 struct for BTMImport136
type BTMImport136 struct {
	BtType *string `json:"btType,omitempty"`
	// Element microversion that is being imported.
	ImportMicroversion         *string `json:"importMicroversion,omitempty"`
	NodeId                     *string `json:"nodeId,omitempty"`
	ElementImport              *bool   `json:"elementImport,omitempty"`
	ImportedExternalDocumentId *string `json:"importedExternalDocumentId,omitempty"`
	Namespace                  *string `json:"namespace,omitempty"`
	Path                       *string `json:"path,omitempty"`
	Version                    *string `json:"version,omitempty"`
}

// NewBTMImport136 instantiates a new BTMImport136 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMImport136() *BTMImport136 {
	this := BTMImport136{}
	return &this
}

// NewBTMImport136WithDefaults instantiates a new BTMImport136 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMImport136WithDefaults() *BTMImport136 {
	this := BTMImport136{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMImport136) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMImport136) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMImport136) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMImport136) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMImport136) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMImport136) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMImport136) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMImport136) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMImport136) SetNodeId(v string) {
	o.NodeId = &v
}

// GetElementImport returns the ElementImport field value if set, zero value otherwise.
func (o *BTMImport136) GetElementImport() bool {
	if o == nil || o.ElementImport == nil {
		var ret bool
		return ret
	}
	return *o.ElementImport
}

// GetElementImportOk returns a tuple with the ElementImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetElementImportOk() (*bool, bool) {
	if o == nil || o.ElementImport == nil {
		return nil, false
	}
	return o.ElementImport, true
}

// HasElementImport returns a boolean if a field has been set.
func (o *BTMImport136) HasElementImport() bool {
	if o != nil && o.ElementImport != nil {
		return true
	}

	return false
}

// SetElementImport gets a reference to the given bool and assigns it to the ElementImport field.
func (o *BTMImport136) SetElementImport(v bool) {
	o.ElementImport = &v
}

// GetImportedExternalDocumentId returns the ImportedExternalDocumentId field value if set, zero value otherwise.
func (o *BTMImport136) GetImportedExternalDocumentId() string {
	if o == nil || o.ImportedExternalDocumentId == nil {
		var ret string
		return ret
	}
	return *o.ImportedExternalDocumentId
}

// GetImportedExternalDocumentIdOk returns a tuple with the ImportedExternalDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetImportedExternalDocumentIdOk() (*string, bool) {
	if o == nil || o.ImportedExternalDocumentId == nil {
		return nil, false
	}
	return o.ImportedExternalDocumentId, true
}

// HasImportedExternalDocumentId returns a boolean if a field has been set.
func (o *BTMImport136) HasImportedExternalDocumentId() bool {
	if o != nil && o.ImportedExternalDocumentId != nil {
		return true
	}

	return false
}

// SetImportedExternalDocumentId gets a reference to the given string and assigns it to the ImportedExternalDocumentId field.
func (o *BTMImport136) SetImportedExternalDocumentId(v string) {
	o.ImportedExternalDocumentId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMImport136) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMImport136) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMImport136) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMImport136) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTMImport136) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BTMImport136) SetPath(v string) {
	o.Path = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTMImport136) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMImport136) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTMImport136) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BTMImport136) SetVersion(v string) {
	o.Version = &v
}

func (o BTMImport136) MarshalJSON() ([]byte, error) {
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
	if o.ElementImport != nil {
		toSerialize["elementImport"] = o.ElementImport
	}
	if o.ImportedExternalDocumentId != nil {
		toSerialize["importedExternalDocumentId"] = o.ImportedExternalDocumentId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBTMImport136 struct {
	value *BTMImport136
	isSet bool
}

func (v NullableBTMImport136) Get() *BTMImport136 {
	return v.value
}

func (v *NullableBTMImport136) Set(val *BTMImport136) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMImport136) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMImport136) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMImport136(val *BTMImport136) *NullableBTMImport136 {
	return &NullableBTMImport136{value: val, isSet: true}
}

func (v NullableBTMImport136) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMImport136) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
