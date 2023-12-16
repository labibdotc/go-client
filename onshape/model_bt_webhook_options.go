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

// BTWebhookOptions struct for BTWebhookOptions
type BTWebhookOptions struct {
	CollapseEvents *bool `json:"collapseEvents,omitempty"`
}

// NewBTWebhookOptions instantiates a new BTWebhookOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebhookOptions() *BTWebhookOptions {
	this := BTWebhookOptions{}
	return &this
}

// NewBTWebhookOptionsWithDefaults instantiates a new BTWebhookOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebhookOptionsWithDefaults() *BTWebhookOptions {
	this := BTWebhookOptions{}
	return &this
}

// GetCollapseEvents returns the CollapseEvents field value if set, zero value otherwise.
func (o *BTWebhookOptions) GetCollapseEvents() bool {
	if o == nil || o.CollapseEvents == nil {
		var ret bool
		return ret
	}
	return *o.CollapseEvents
}

// GetCollapseEventsOk returns a tuple with the CollapseEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookOptions) GetCollapseEventsOk() (*bool, bool) {
	if o == nil || o.CollapseEvents == nil {
		return nil, false
	}
	return o.CollapseEvents, true
}

// HasCollapseEvents returns a boolean if a field has been set.
func (o *BTWebhookOptions) HasCollapseEvents() bool {
	if o != nil && o.CollapseEvents != nil {
		return true
	}

	return false
}

// SetCollapseEvents gets a reference to the given bool and assigns it to the CollapseEvents field.
func (o *BTWebhookOptions) SetCollapseEvents(v bool) {
	o.CollapseEvents = &v
}

func (o BTWebhookOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollapseEvents != nil {
		toSerialize["collapseEvents"] = o.CollapseEvents
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebhookOptions struct {
	value *BTWebhookOptions
	isSet bool
}

func (v NullableBTWebhookOptions) Get() *BTWebhookOptions {
	return v.value
}

func (v *NullableBTWebhookOptions) Set(val *BTWebhookOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebhookOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebhookOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebhookOptions(val *BTWebhookOptions) *NullableBTWebhookOptions {
	return &NullableBTWebhookOptions{value: val, isSet: true}
}

func (v NullableBTWebhookOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebhookOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
