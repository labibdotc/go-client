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

// BTPlanSubscriberInfo struct for BTPlanSubscriberInfo
type BTPlanSubscriberInfo struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id       *string `json:"id,omitempty"`
	Image    *string `json:"image,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	// Name of the resource.
	Name  *string `json:"name,omitempty"`
	State *int32  `json:"state,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTPlanSubscriberInfo instantiates a new BTPlanSubscriberInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPlanSubscriberInfo() *BTPlanSubscriberInfo {
	this := BTPlanSubscriberInfo{}
	return &this
}

// NewBTPlanSubscriberInfoWithDefaults instantiates a new BTPlanSubscriberInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPlanSubscriberInfoWithDefaults() *BTPlanSubscriberInfo {
	this := BTPlanSubscriberInfo{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTPlanSubscriberInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTPlanSubscriberInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTPlanSubscriberInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPlanSubscriberInfo) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTPlanSubscriberInfo) SetImage(v string) {
	o.Image = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTPlanSubscriberInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPlanSubscriberInfo) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTPlanSubscriberInfo) SetState(v int32) {
	o.State = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTPlanSubscriberInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlanSubscriberInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTPlanSubscriberInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTPlanSubscriberInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTPlanSubscriberInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTPlanSubscriberInfo struct {
	value *BTPlanSubscriberInfo
	isSet bool
}

func (v NullableBTPlanSubscriberInfo) Get() *BTPlanSubscriberInfo {
	return v.value
}

func (v *NullableBTPlanSubscriberInfo) Set(val *BTPlanSubscriberInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPlanSubscriberInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPlanSubscriberInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPlanSubscriberInfo(val *BTPlanSubscriberInfo) *NullableBTPlanSubscriberInfo {
	return &NullableBTPlanSubscriberInfo{value: val, isSet: true}
}

func (v NullableBTPlanSubscriberInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPlanSubscriberInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
