/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCompanyUserInfo struct for BTCompanyUserInfo
type BTCompanyUserInfo struct {
	Admin                     *bool                 `json:"admin,omitempty"`
	Company                   *BTCompanySummaryInfo `json:"company,omitempty"`
	DocumentationNameOverride *string               `json:"documentationNameOverride,omitempty"`
	Guest                     *bool                 `json:"guest,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id            *string   `json:"id,omitempty"`
	LastLoginTime *JSONTime `json:"lastLoginTime,omitempty"`
	Light         *bool     `json:"light,omitempty"`
	// Name of the resource.
	Name  *string                 `json:"name,omitempty"`
	State *int32                  `json:"state,omitempty"`
	User  *BTUserBasicSummaryInfo `json:"user,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTCompanyUserInfo instantiates a new BTCompanyUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCompanyUserInfo() *BTCompanyUserInfo {
	this := BTCompanyUserInfo{}
	return &this
}

// NewBTCompanyUserInfoWithDefaults instantiates a new BTCompanyUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCompanyUserInfoWithDefaults() *BTCompanyUserInfo {
	this := BTCompanyUserInfo{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTCompanyUserInfo) SetAdmin(v bool) {
	o.Admin = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetCompany() BTCompanySummaryInfo {
	if o == nil || o.Company == nil {
		var ret BTCompanySummaryInfo
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given BTCompanySummaryInfo and assigns it to the Company field.
func (o *BTCompanyUserInfo) SetCompany(v BTCompanySummaryInfo) {
	o.Company = &v
}

// GetDocumentationNameOverride returns the DocumentationNameOverride field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetDocumentationNameOverride() string {
	if o == nil || o.DocumentationNameOverride == nil {
		var ret string
		return ret
	}
	return *o.DocumentationNameOverride
}

// GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetDocumentationNameOverrideOk() (*string, bool) {
	if o == nil || o.DocumentationNameOverride == nil {
		return nil, false
	}
	return o.DocumentationNameOverride, true
}

// HasDocumentationNameOverride returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasDocumentationNameOverride() bool {
	if o != nil && o.DocumentationNameOverride != nil {
		return true
	}

	return false
}

// SetDocumentationNameOverride gets a reference to the given string and assigns it to the DocumentationNameOverride field.
func (o *BTCompanyUserInfo) SetDocumentationNameOverride(v string) {
	o.DocumentationNameOverride = &v
}

// GetGuest returns the Guest field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetGuest() bool {
	if o == nil || o.Guest == nil {
		var ret bool
		return ret
	}
	return *o.Guest
}

// GetGuestOk returns a tuple with the Guest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetGuestOk() (*bool, bool) {
	if o == nil || o.Guest == nil {
		return nil, false
	}
	return o.Guest, true
}

// HasGuest returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasGuest() bool {
	if o != nil && o.Guest != nil {
		return true
	}

	return false
}

// SetGuest gets a reference to the given bool and assigns it to the Guest field.
func (o *BTCompanyUserInfo) SetGuest(v bool) {
	o.Guest = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCompanyUserInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCompanyUserInfo) SetId(v string) {
	o.Id = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetLastLoginTime() JSONTime {
	if o == nil || o.LastLoginTime == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetLastLoginTimeOk() (*JSONTime, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given JSONTime and assigns it to the LastLoginTime field.
func (o *BTCompanyUserInfo) SetLastLoginTime(v JSONTime) {
	o.LastLoginTime = &v
}

// GetLight returns the Light field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetLight() bool {
	if o == nil || o.Light == nil {
		var ret bool
		return ret
	}
	return *o.Light
}

// GetLightOk returns a tuple with the Light field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetLightOk() (*bool, bool) {
	if o == nil || o.Light == nil {
		return nil, false
	}
	return o.Light, true
}

// HasLight returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasLight() bool {
	if o != nil && o.Light != nil {
		return true
	}

	return false
}

// SetLight gets a reference to the given bool and assigns it to the Light field.
func (o *BTCompanyUserInfo) SetLight(v bool) {
	o.Light = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCompanyUserInfo) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTCompanyUserInfo) SetState(v int32) {
	o.State = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetUser() BTUserBasicSummaryInfo {
	if o == nil || o.User == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetUserOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given BTUserBasicSummaryInfo and assigns it to the User field.
func (o *BTCompanyUserInfo) SetUser(v BTUserBasicSummaryInfo) {
	o.User = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCompanyUserInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCompanyUserInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCompanyUserInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTCompanyUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.DocumentationNameOverride != nil {
		toSerialize["documentationNameOverride"] = o.DocumentationNameOverride
	}
	if o.Guest != nil {
		toSerialize["guest"] = o.Guest
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastLoginTime != nil {
		toSerialize["lastLoginTime"] = o.LastLoginTime
	}
	if o.Light != nil {
		toSerialize["light"] = o.Light
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTCompanyUserInfo struct {
	value *BTCompanyUserInfo
	isSet bool
}

func (v NullableBTCompanyUserInfo) Get() *BTCompanyUserInfo {
	return v.value
}

func (v *NullableBTCompanyUserInfo) Set(val *BTCompanyUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCompanyUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCompanyUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCompanyUserInfo(val *BTCompanyUserInfo) *NullableBTCompanyUserInfo {
	return &NullableBTCompanyUserInfo{value: val, isSet: true}
}

func (v NullableBTCompanyUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCompanyUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
