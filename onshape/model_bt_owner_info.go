/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6268-501b86d3d653
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOwnerInfo struct for BTOwnerInfo
type BTOwnerInfo struct {
	Href                      *string `json:"href,omitempty"`
	Id                        *string `json:"id,omitempty"`
	Image                     *string `json:"image,omitempty"`
	IsEnterpriseOwnedResource *bool   `json:"isEnterpriseOwnedResource,omitempty"`
	Name                      *string `json:"name,omitempty"`
	Type                      *int32  `json:"type,omitempty"`
	ViewRef                   *string `json:"viewRef,omitempty"`
}

// NewBTOwnerInfo instantiates a new BTOwnerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOwnerInfo() *BTOwnerInfo {
	this := BTOwnerInfo{}
	return &this
}

// NewBTOwnerInfoWithDefaults instantiates a new BTOwnerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOwnerInfoWithDefaults() *BTOwnerInfo {
	this := BTOwnerInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTOwnerInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTOwnerInfo) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTOwnerInfo) SetImage(v string) {
	o.Image = &v
}

// GetIsEnterpriseOwnedResource returns the IsEnterpriseOwnedResource field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetIsEnterpriseOwnedResource() bool {
	if o == nil || o.IsEnterpriseOwnedResource == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwnedResource
}

// GetIsEnterpriseOwnedResourceOk returns a tuple with the IsEnterpriseOwnedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetIsEnterpriseOwnedResourceOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwnedResource == nil {
		return nil, false
	}
	return o.IsEnterpriseOwnedResource, true
}

// HasIsEnterpriseOwnedResource returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasIsEnterpriseOwnedResource() bool {
	if o != nil && o.IsEnterpriseOwnedResource != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwnedResource gets a reference to the given bool and assigns it to the IsEnterpriseOwnedResource field.
func (o *BTOwnerInfo) SetIsEnterpriseOwnedResource(v bool) {
	o.IsEnterpriseOwnedResource = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTOwnerInfo) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *BTOwnerInfo) SetType(v int32) {
	o.Type = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTOwnerInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwnerInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTOwnerInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTOwnerInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTOwnerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.IsEnterpriseOwnedResource != nil {
		toSerialize["isEnterpriseOwnedResource"] = o.IsEnterpriseOwnedResource
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTOwnerInfo struct {
	value *BTOwnerInfo
	isSet bool
}

func (v NullableBTOwnerInfo) Get() *BTOwnerInfo {
	return v.value
}

func (v *NullableBTOwnerInfo) Set(val *BTOwnerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOwnerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOwnerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOwnerInfo(val *BTOwnerInfo) *NullableBTOwnerInfo {
	return &NullableBTOwnerInfo{value: val, isSet: true}
}

func (v NullableBTOwnerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOwnerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
