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

// BTUnchangedReferenceInfo struct for BTUnchangedReferenceInfo
type BTUnchangedReferenceInfo struct {
	MetadataUnchanged *bool           `json:"metadataUnchanged,omitempty"`
	NodeIds           []string        `json:"nodeIds,omitempty"`
	ToRevision        *BTRevisionInfo `json:"toRevision,omitempty"`
}

// NewBTUnchangedReferenceInfo instantiates a new BTUnchangedReferenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUnchangedReferenceInfo() *BTUnchangedReferenceInfo {
	this := BTUnchangedReferenceInfo{}
	return &this
}

// NewBTUnchangedReferenceInfoWithDefaults instantiates a new BTUnchangedReferenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUnchangedReferenceInfoWithDefaults() *BTUnchangedReferenceInfo {
	this := BTUnchangedReferenceInfo{}
	return &this
}

// GetMetadataUnchanged returns the MetadataUnchanged field value if set, zero value otherwise.
func (o *BTUnchangedReferenceInfo) GetMetadataUnchanged() bool {
	if o == nil || o.MetadataUnchanged == nil {
		var ret bool
		return ret
	}
	return *o.MetadataUnchanged
}

// GetMetadataUnchangedOk returns a tuple with the MetadataUnchanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnchangedReferenceInfo) GetMetadataUnchangedOk() (*bool, bool) {
	if o == nil || o.MetadataUnchanged == nil {
		return nil, false
	}
	return o.MetadataUnchanged, true
}

// HasMetadataUnchanged returns a boolean if a field has been set.
func (o *BTUnchangedReferenceInfo) HasMetadataUnchanged() bool {
	if o != nil && o.MetadataUnchanged != nil {
		return true
	}

	return false
}

// SetMetadataUnchanged gets a reference to the given bool and assigns it to the MetadataUnchanged field.
func (o *BTUnchangedReferenceInfo) SetMetadataUnchanged(v bool) {
	o.MetadataUnchanged = &v
}

// GetNodeIds returns the NodeIds field value if set, zero value otherwise.
func (o *BTUnchangedReferenceInfo) GetNodeIds() []string {
	if o == nil || o.NodeIds == nil {
		var ret []string
		return ret
	}
	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnchangedReferenceInfo) GetNodeIdsOk() ([]string, bool) {
	if o == nil || o.NodeIds == nil {
		return nil, false
	}
	return o.NodeIds, true
}

// HasNodeIds returns a boolean if a field has been set.
func (o *BTUnchangedReferenceInfo) HasNodeIds() bool {
	if o != nil && o.NodeIds != nil {
		return true
	}

	return false
}

// SetNodeIds gets a reference to the given []string and assigns it to the NodeIds field.
func (o *BTUnchangedReferenceInfo) SetNodeIds(v []string) {
	o.NodeIds = v
}

// GetToRevision returns the ToRevision field value if set, zero value otherwise.
func (o *BTUnchangedReferenceInfo) GetToRevision() BTRevisionInfo {
	if o == nil || o.ToRevision == nil {
		var ret BTRevisionInfo
		return ret
	}
	return *o.ToRevision
}

// GetToRevisionOk returns a tuple with the ToRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnchangedReferenceInfo) GetToRevisionOk() (*BTRevisionInfo, bool) {
	if o == nil || o.ToRevision == nil {
		return nil, false
	}
	return o.ToRevision, true
}

// HasToRevision returns a boolean if a field has been set.
func (o *BTUnchangedReferenceInfo) HasToRevision() bool {
	if o != nil && o.ToRevision != nil {
		return true
	}

	return false
}

// SetToRevision gets a reference to the given BTRevisionInfo and assigns it to the ToRevision field.
func (o *BTUnchangedReferenceInfo) SetToRevision(v BTRevisionInfo) {
	o.ToRevision = &v
}

func (o BTUnchangedReferenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataUnchanged != nil {
		toSerialize["metadataUnchanged"] = o.MetadataUnchanged
	}
	if o.NodeIds != nil {
		toSerialize["nodeIds"] = o.NodeIds
	}
	if o.ToRevision != nil {
		toSerialize["toRevision"] = o.ToRevision
	}
	return json.Marshal(toSerialize)
}

type NullableBTUnchangedReferenceInfo struct {
	value *BTUnchangedReferenceInfo
	isSet bool
}

func (v NullableBTUnchangedReferenceInfo) Get() *BTUnchangedReferenceInfo {
	return v.value
}

func (v *NullableBTUnchangedReferenceInfo) Set(val *BTUnchangedReferenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUnchangedReferenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUnchangedReferenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUnchangedReferenceInfo(val *BTUnchangedReferenceInfo) *NullableBTUnchangedReferenceInfo {
	return &NullableBTUnchangedReferenceInfo{value: val, isSet: true}
}

func (v NullableBTUnchangedReferenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUnchangedReferenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
