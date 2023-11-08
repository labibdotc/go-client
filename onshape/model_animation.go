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

// Animation struct for Animation
type Animation struct {
	Channels   []AnimationChannel                `json:"channels,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Samplers   []AnimationSampler                `json:"samplers,omitempty"`
}

// NewAnimation instantiates a new Animation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimation() *Animation {
	this := Animation{}
	return &this
}

// NewAnimationWithDefaults instantiates a new Animation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimationWithDefaults() *Animation {
	this := Animation{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *Animation) GetChannels() []AnimationChannel {
	if o == nil || o.Channels == nil {
		var ret []AnimationChannel
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetChannelsOk() ([]AnimationChannel, bool) {
	if o == nil || o.Channels == nil {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *Animation) HasChannels() bool {
	if o != nil && o.Channels != nil {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []AnimationChannel and assigns it to the Channels field.
func (o *Animation) SetChannels(v []AnimationChannel) {
	o.Channels = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Animation) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Animation) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Animation) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Animation) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Animation) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Animation) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Animation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Animation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Animation) SetName(v string) {
	o.Name = &v
}

// GetSamplers returns the Samplers field value if set, zero value otherwise.
func (o *Animation) GetSamplers() []AnimationSampler {
	if o == nil || o.Samplers == nil {
		var ret []AnimationSampler
		return ret
	}
	return o.Samplers
}

// GetSamplersOk returns a tuple with the Samplers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animation) GetSamplersOk() ([]AnimationSampler, bool) {
	if o == nil || o.Samplers == nil {
		return nil, false
	}
	return o.Samplers, true
}

// HasSamplers returns a boolean if a field has been set.
func (o *Animation) HasSamplers() bool {
	if o != nil && o.Samplers != nil {
		return true
	}

	return false
}

// SetSamplers gets a reference to the given []AnimationSampler and assigns it to the Samplers field.
func (o *Animation) SetSamplers(v []AnimationSampler) {
	o.Samplers = v
}

func (o Animation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Samplers != nil {
		toSerialize["samplers"] = o.Samplers
	}
	return json.Marshal(toSerialize)
}

type NullableAnimation struct {
	value *Animation
	isSet bool
}

func (v NullableAnimation) Get() *Animation {
	return v.value
}

func (v *NullableAnimation) Set(val *Animation) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimation) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimation(val *Animation) *NullableAnimation {
	return &NullableAnimation{value: val, isSet: true}
}

func (v NullableAnimation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
