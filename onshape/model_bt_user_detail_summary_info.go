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

// BTUserDetailSummaryInfo struct for BTUserDetailSummaryInfo
type BTUserDetailSummaryInfo struct {
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef           *string `json:"viewRef,omitempty"`
	Image             *string `json:"image,omitempty"`
	State             *int32  `json:"state,omitempty"`
	DocumentationName *string `json:"documentationName,omitempty"`
	Email             *string `json:"email,omitempty"`
	FirstName         *string `json:"firstName,omitempty"`
	LastName          *string `json:"lastName,omitempty"`
}

// NewBTUserDetailSummaryInfo instantiates a new BTUserDetailSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserDetailSummaryInfo() *BTUserDetailSummaryInfo {
	this := BTUserDetailSummaryInfo{}
	return &this
}

// NewBTUserDetailSummaryInfoWithDefaults instantiates a new BTUserDetailSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserDetailSummaryInfoWithDefaults() *BTUserDetailSummaryInfo {
	this := BTUserDetailSummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserDetailSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserDetailSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserDetailSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserDetailSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTUserDetailSummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTUserDetailSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetDocumentationName returns the DocumentationName field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetDocumentationName() string {
	if o == nil || o.DocumentationName == nil {
		var ret string
		return ret
	}
	return *o.DocumentationName
}

// GetDocumentationNameOk returns a tuple with the DocumentationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetDocumentationNameOk() (*string, bool) {
	if o == nil || o.DocumentationName == nil {
		return nil, false
	}
	return o.DocumentationName, true
}

// HasDocumentationName returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasDocumentationName() bool {
	if o != nil && o.DocumentationName != nil {
		return true
	}

	return false
}

// SetDocumentationName gets a reference to the given string and assigns it to the DocumentationName field.
func (o *BTUserDetailSummaryInfo) SetDocumentationName(v string) {
	o.DocumentationName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTUserDetailSummaryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTUserDetailSummaryInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTUserDetailSummaryInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserDetailSummaryInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTUserDetailSummaryInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTUserDetailSummaryInfo) SetLastName(v string) {
	o.LastName = &v
}

func (o BTUserDetailSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.DocumentationName != nil {
		toSerialize["documentationName"] = o.DocumentationName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserDetailSummaryInfo struct {
	value *BTUserDetailSummaryInfo
	isSet bool
}

func (v NullableBTUserDetailSummaryInfo) Get() *BTUserDetailSummaryInfo {
	return v.value
}

func (v *NullableBTUserDetailSummaryInfo) Set(val *BTUserDetailSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserDetailSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserDetailSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserDetailSummaryInfo(val *BTUserDetailSummaryInfo) *NullableBTUserDetailSummaryInfo {
	return &NullableBTUserDetailSummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserDetailSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserDetailSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
