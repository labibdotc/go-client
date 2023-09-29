/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23149-3610d8cd750e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTrialInfo struct for BTTrialInfo
type BTTrialInfo struct {
	PaidCustomerInPast *bool     `json:"paidCustomerInPast,omitempty"`
	PlanId             *string   `json:"planId,omitempty"`
	PlanInterval       *string   `json:"planInterval,omitempty"`
	Seats              *int64    `json:"seats,omitempty"`
	TrialDaysRemaining *int64    `json:"trialDaysRemaining,omitempty"`
	TrialEndDate       *JSONTime `json:"trialEndDate,omitempty"`
	TrialStartDate     *JSONTime `json:"trialStartDate,omitempty"`
}

// NewBTTrialInfo instantiates a new BTTrialInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTrialInfo() *BTTrialInfo {
	this := BTTrialInfo{}
	return &this
}

// NewBTTrialInfoWithDefaults instantiates a new BTTrialInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTrialInfoWithDefaults() *BTTrialInfo {
	this := BTTrialInfo{}
	return &this
}

// GetPaidCustomerInPast returns the PaidCustomerInPast field value if set, zero value otherwise.
func (o *BTTrialInfo) GetPaidCustomerInPast() bool {
	if o == nil || o.PaidCustomerInPast == nil {
		var ret bool
		return ret
	}
	return *o.PaidCustomerInPast
}

// GetPaidCustomerInPastOk returns a tuple with the PaidCustomerInPast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetPaidCustomerInPastOk() (*bool, bool) {
	if o == nil || o.PaidCustomerInPast == nil {
		return nil, false
	}
	return o.PaidCustomerInPast, true
}

// HasPaidCustomerInPast returns a boolean if a field has been set.
func (o *BTTrialInfo) HasPaidCustomerInPast() bool {
	if o != nil && o.PaidCustomerInPast != nil {
		return true
	}

	return false
}

// SetPaidCustomerInPast gets a reference to the given bool and assigns it to the PaidCustomerInPast field.
func (o *BTTrialInfo) SetPaidCustomerInPast(v bool) {
	o.PaidCustomerInPast = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *BTTrialInfo) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *BTTrialInfo) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *BTTrialInfo) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPlanInterval returns the PlanInterval field value if set, zero value otherwise.
func (o *BTTrialInfo) GetPlanInterval() string {
	if o == nil || o.PlanInterval == nil {
		var ret string
		return ret
	}
	return *o.PlanInterval
}

// GetPlanIntervalOk returns a tuple with the PlanInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetPlanIntervalOk() (*string, bool) {
	if o == nil || o.PlanInterval == nil {
		return nil, false
	}
	return o.PlanInterval, true
}

// HasPlanInterval returns a boolean if a field has been set.
func (o *BTTrialInfo) HasPlanInterval() bool {
	if o != nil && o.PlanInterval != nil {
		return true
	}

	return false
}

// SetPlanInterval gets a reference to the given string and assigns it to the PlanInterval field.
func (o *BTTrialInfo) SetPlanInterval(v string) {
	o.PlanInterval = &v
}

// GetSeats returns the Seats field value if set, zero value otherwise.
func (o *BTTrialInfo) GetSeats() int64 {
	if o == nil || o.Seats == nil {
		var ret int64
		return ret
	}
	return *o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetSeatsOk() (*int64, bool) {
	if o == nil || o.Seats == nil {
		return nil, false
	}
	return o.Seats, true
}

// HasSeats returns a boolean if a field has been set.
func (o *BTTrialInfo) HasSeats() bool {
	if o != nil && o.Seats != nil {
		return true
	}

	return false
}

// SetSeats gets a reference to the given int64 and assigns it to the Seats field.
func (o *BTTrialInfo) SetSeats(v int64) {
	o.Seats = &v
}

// GetTrialDaysRemaining returns the TrialDaysRemaining field value if set, zero value otherwise.
func (o *BTTrialInfo) GetTrialDaysRemaining() int64 {
	if o == nil || o.TrialDaysRemaining == nil {
		var ret int64
		return ret
	}
	return *o.TrialDaysRemaining
}

// GetTrialDaysRemainingOk returns a tuple with the TrialDaysRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetTrialDaysRemainingOk() (*int64, bool) {
	if o == nil || o.TrialDaysRemaining == nil {
		return nil, false
	}
	return o.TrialDaysRemaining, true
}

// HasTrialDaysRemaining returns a boolean if a field has been set.
func (o *BTTrialInfo) HasTrialDaysRemaining() bool {
	if o != nil && o.TrialDaysRemaining != nil {
		return true
	}

	return false
}

// SetTrialDaysRemaining gets a reference to the given int64 and assigns it to the TrialDaysRemaining field.
func (o *BTTrialInfo) SetTrialDaysRemaining(v int64) {
	o.TrialDaysRemaining = &v
}

// GetTrialEndDate returns the TrialEndDate field value if set, zero value otherwise.
func (o *BTTrialInfo) GetTrialEndDate() JSONTime {
	if o == nil || o.TrialEndDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.TrialEndDate
}

// GetTrialEndDateOk returns a tuple with the TrialEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetTrialEndDateOk() (*JSONTime, bool) {
	if o == nil || o.TrialEndDate == nil {
		return nil, false
	}
	return o.TrialEndDate, true
}

// HasTrialEndDate returns a boolean if a field has been set.
func (o *BTTrialInfo) HasTrialEndDate() bool {
	if o != nil && o.TrialEndDate != nil {
		return true
	}

	return false
}

// SetTrialEndDate gets a reference to the given JSONTime and assigns it to the TrialEndDate field.
func (o *BTTrialInfo) SetTrialEndDate(v JSONTime) {
	o.TrialEndDate = &v
}

// GetTrialStartDate returns the TrialStartDate field value if set, zero value otherwise.
func (o *BTTrialInfo) GetTrialStartDate() JSONTime {
	if o == nil || o.TrialStartDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.TrialStartDate
}

// GetTrialStartDateOk returns a tuple with the TrialStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTrialInfo) GetTrialStartDateOk() (*JSONTime, bool) {
	if o == nil || o.TrialStartDate == nil {
		return nil, false
	}
	return o.TrialStartDate, true
}

// HasTrialStartDate returns a boolean if a field has been set.
func (o *BTTrialInfo) HasTrialStartDate() bool {
	if o != nil && o.TrialStartDate != nil {
		return true
	}

	return false
}

// SetTrialStartDate gets a reference to the given JSONTime and assigns it to the TrialStartDate field.
func (o *BTTrialInfo) SetTrialStartDate(v JSONTime) {
	o.TrialStartDate = &v
}

func (o BTTrialInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaidCustomerInPast != nil {
		toSerialize["paidCustomerInPast"] = o.PaidCustomerInPast
	}
	if o.PlanId != nil {
		toSerialize["planId"] = o.PlanId
	}
	if o.PlanInterval != nil {
		toSerialize["planInterval"] = o.PlanInterval
	}
	if o.Seats != nil {
		toSerialize["seats"] = o.Seats
	}
	if o.TrialDaysRemaining != nil {
		toSerialize["trialDaysRemaining"] = o.TrialDaysRemaining
	}
	if o.TrialEndDate != nil {
		toSerialize["trialEndDate"] = o.TrialEndDate
	}
	if o.TrialStartDate != nil {
		toSerialize["trialStartDate"] = o.TrialStartDate
	}
	return json.Marshal(toSerialize)
}

type NullableBTTrialInfo struct {
	value *BTTrialInfo
	isSet bool
}

func (v NullableBTTrialInfo) Get() *BTTrialInfo {
	return v.value
}

func (v *NullableBTTrialInfo) Set(val *BTTrialInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTrialInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTrialInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTrialInfo(val *BTTrialInfo) *NullableBTTrialInfo {
	return &NullableBTTrialInfo{value: val, isSet: true}
}

func (v NullableBTTrialInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTrialInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}