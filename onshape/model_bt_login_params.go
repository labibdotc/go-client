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

// BTLoginParams struct for BTLoginParams
type BTLoginParams struct {
	DeviceId                       *string                                    `json:"deviceId,omitempty"`
	Email                          *string                                    `json:"email,omitempty"`
	EnableTotp                     *bool                                      `json:"enableTotp,omitempty"`
	IsRecaptchaV3                  *bool                                      `json:"isRecaptchaV3,omitempty"`
	Password                       *string                                    `json:"password,omitempty"`
	RandomToken                    *string                                    `json:"randomToken,omitempty"`
	RecaptchaToken                 *string                                    `json:"recaptchaToken,omitempty"`
	RememberTotp                   *bool                                      `json:"rememberTotp,omitempty"`
	RendererPerformanceMeasurement *BTWebRendererPerformanceMeasurementParams `json:"rendererPerformanceMeasurement,omitempty"`
	Totp                           *string                                    `json:"totp,omitempty"`
	WebClientCapabilities          *BTWebClientCapabilitiesParams             `json:"webClientCapabilities,omitempty"`
}

// NewBTLoginParams instantiates a new BTLoginParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLoginParams() *BTLoginParams {
	this := BTLoginParams{}
	return &this
}

// NewBTLoginParamsWithDefaults instantiates a new BTLoginParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLoginParamsWithDefaults() *BTLoginParams {
	this := BTLoginParams{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *BTLoginParams) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *BTLoginParams) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *BTLoginParams) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTLoginParams) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTLoginParams) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTLoginParams) SetEmail(v string) {
	o.Email = &v
}

// GetEnableTotp returns the EnableTotp field value if set, zero value otherwise.
func (o *BTLoginParams) GetEnableTotp() bool {
	if o == nil || o.EnableTotp == nil {
		var ret bool
		return ret
	}
	return *o.EnableTotp
}

// GetEnableTotpOk returns a tuple with the EnableTotp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetEnableTotpOk() (*bool, bool) {
	if o == nil || o.EnableTotp == nil {
		return nil, false
	}
	return o.EnableTotp, true
}

// HasEnableTotp returns a boolean if a field has been set.
func (o *BTLoginParams) HasEnableTotp() bool {
	if o != nil && o.EnableTotp != nil {
		return true
	}

	return false
}

// SetEnableTotp gets a reference to the given bool and assigns it to the EnableTotp field.
func (o *BTLoginParams) SetEnableTotp(v bool) {
	o.EnableTotp = &v
}

// GetIsRecaptchaV3 returns the IsRecaptchaV3 field value if set, zero value otherwise.
func (o *BTLoginParams) GetIsRecaptchaV3() bool {
	if o == nil || o.IsRecaptchaV3 == nil {
		var ret bool
		return ret
	}
	return *o.IsRecaptchaV3
}

// GetIsRecaptchaV3Ok returns a tuple with the IsRecaptchaV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetIsRecaptchaV3Ok() (*bool, bool) {
	if o == nil || o.IsRecaptchaV3 == nil {
		return nil, false
	}
	return o.IsRecaptchaV3, true
}

// HasIsRecaptchaV3 returns a boolean if a field has been set.
func (o *BTLoginParams) HasIsRecaptchaV3() bool {
	if o != nil && o.IsRecaptchaV3 != nil {
		return true
	}

	return false
}

// SetIsRecaptchaV3 gets a reference to the given bool and assigns it to the IsRecaptchaV3 field.
func (o *BTLoginParams) SetIsRecaptchaV3(v bool) {
	o.IsRecaptchaV3 = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *BTLoginParams) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *BTLoginParams) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *BTLoginParams) SetPassword(v string) {
	o.Password = &v
}

// GetRandomToken returns the RandomToken field value if set, zero value otherwise.
func (o *BTLoginParams) GetRandomToken() string {
	if o == nil || o.RandomToken == nil {
		var ret string
		return ret
	}
	return *o.RandomToken
}

// GetRandomTokenOk returns a tuple with the RandomToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetRandomTokenOk() (*string, bool) {
	if o == nil || o.RandomToken == nil {
		return nil, false
	}
	return o.RandomToken, true
}

// HasRandomToken returns a boolean if a field has been set.
func (o *BTLoginParams) HasRandomToken() bool {
	if o != nil && o.RandomToken != nil {
		return true
	}

	return false
}

// SetRandomToken gets a reference to the given string and assigns it to the RandomToken field.
func (o *BTLoginParams) SetRandomToken(v string) {
	o.RandomToken = &v
}

// GetRecaptchaToken returns the RecaptchaToken field value if set, zero value otherwise.
func (o *BTLoginParams) GetRecaptchaToken() string {
	if o == nil || o.RecaptchaToken == nil {
		var ret string
		return ret
	}
	return *o.RecaptchaToken
}

// GetRecaptchaTokenOk returns a tuple with the RecaptchaToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetRecaptchaTokenOk() (*string, bool) {
	if o == nil || o.RecaptchaToken == nil {
		return nil, false
	}
	return o.RecaptchaToken, true
}

// HasRecaptchaToken returns a boolean if a field has been set.
func (o *BTLoginParams) HasRecaptchaToken() bool {
	if o != nil && o.RecaptchaToken != nil {
		return true
	}

	return false
}

// SetRecaptchaToken gets a reference to the given string and assigns it to the RecaptchaToken field.
func (o *BTLoginParams) SetRecaptchaToken(v string) {
	o.RecaptchaToken = &v
}

// GetRememberTotp returns the RememberTotp field value if set, zero value otherwise.
func (o *BTLoginParams) GetRememberTotp() bool {
	if o == nil || o.RememberTotp == nil {
		var ret bool
		return ret
	}
	return *o.RememberTotp
}

// GetRememberTotpOk returns a tuple with the RememberTotp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetRememberTotpOk() (*bool, bool) {
	if o == nil || o.RememberTotp == nil {
		return nil, false
	}
	return o.RememberTotp, true
}

// HasRememberTotp returns a boolean if a field has been set.
func (o *BTLoginParams) HasRememberTotp() bool {
	if o != nil && o.RememberTotp != nil {
		return true
	}

	return false
}

// SetRememberTotp gets a reference to the given bool and assigns it to the RememberTotp field.
func (o *BTLoginParams) SetRememberTotp(v bool) {
	o.RememberTotp = &v
}

// GetRendererPerformanceMeasurement returns the RendererPerformanceMeasurement field value if set, zero value otherwise.
func (o *BTLoginParams) GetRendererPerformanceMeasurement() BTWebRendererPerformanceMeasurementParams {
	if o == nil || o.RendererPerformanceMeasurement == nil {
		var ret BTWebRendererPerformanceMeasurementParams
		return ret
	}
	return *o.RendererPerformanceMeasurement
}

// GetRendererPerformanceMeasurementOk returns a tuple with the RendererPerformanceMeasurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetRendererPerformanceMeasurementOk() (*BTWebRendererPerformanceMeasurementParams, bool) {
	if o == nil || o.RendererPerformanceMeasurement == nil {
		return nil, false
	}
	return o.RendererPerformanceMeasurement, true
}

// HasRendererPerformanceMeasurement returns a boolean if a field has been set.
func (o *BTLoginParams) HasRendererPerformanceMeasurement() bool {
	if o != nil && o.RendererPerformanceMeasurement != nil {
		return true
	}

	return false
}

// SetRendererPerformanceMeasurement gets a reference to the given BTWebRendererPerformanceMeasurementParams and assigns it to the RendererPerformanceMeasurement field.
func (o *BTLoginParams) SetRendererPerformanceMeasurement(v BTWebRendererPerformanceMeasurementParams) {
	o.RendererPerformanceMeasurement = &v
}

// GetTotp returns the Totp field value if set, zero value otherwise.
func (o *BTLoginParams) GetTotp() string {
	if o == nil || o.Totp == nil {
		var ret string
		return ret
	}
	return *o.Totp
}

// GetTotpOk returns a tuple with the Totp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetTotpOk() (*string, bool) {
	if o == nil || o.Totp == nil {
		return nil, false
	}
	return o.Totp, true
}

// HasTotp returns a boolean if a field has been set.
func (o *BTLoginParams) HasTotp() bool {
	if o != nil && o.Totp != nil {
		return true
	}

	return false
}

// SetTotp gets a reference to the given string and assigns it to the Totp field.
func (o *BTLoginParams) SetTotp(v string) {
	o.Totp = &v
}

// GetWebClientCapabilities returns the WebClientCapabilities field value if set, zero value otherwise.
func (o *BTLoginParams) GetWebClientCapabilities() BTWebClientCapabilitiesParams {
	if o == nil || o.WebClientCapabilities == nil {
		var ret BTWebClientCapabilitiesParams
		return ret
	}
	return *o.WebClientCapabilities
}

// GetWebClientCapabilitiesOk returns a tuple with the WebClientCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLoginParams) GetWebClientCapabilitiesOk() (*BTWebClientCapabilitiesParams, bool) {
	if o == nil || o.WebClientCapabilities == nil {
		return nil, false
	}
	return o.WebClientCapabilities, true
}

// HasWebClientCapabilities returns a boolean if a field has been set.
func (o *BTLoginParams) HasWebClientCapabilities() bool {
	if o != nil && o.WebClientCapabilities != nil {
		return true
	}

	return false
}

// SetWebClientCapabilities gets a reference to the given BTWebClientCapabilitiesParams and assigns it to the WebClientCapabilities field.
func (o *BTLoginParams) SetWebClientCapabilities(v BTWebClientCapabilitiesParams) {
	o.WebClientCapabilities = &v
}

func (o BTLoginParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EnableTotp != nil {
		toSerialize["enableTotp"] = o.EnableTotp
	}
	if o.IsRecaptchaV3 != nil {
		toSerialize["isRecaptchaV3"] = o.IsRecaptchaV3
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RandomToken != nil {
		toSerialize["randomToken"] = o.RandomToken
	}
	if o.RecaptchaToken != nil {
		toSerialize["recaptchaToken"] = o.RecaptchaToken
	}
	if o.RememberTotp != nil {
		toSerialize["rememberTotp"] = o.RememberTotp
	}
	if o.RendererPerformanceMeasurement != nil {
		toSerialize["rendererPerformanceMeasurement"] = o.RendererPerformanceMeasurement
	}
	if o.Totp != nil {
		toSerialize["totp"] = o.Totp
	}
	if o.WebClientCapabilities != nil {
		toSerialize["webClientCapabilities"] = o.WebClientCapabilities
	}
	return json.Marshal(toSerialize)
}

type NullableBTLoginParams struct {
	value *BTLoginParams
	isSet bool
}

func (v NullableBTLoginParams) Get() *BTLoginParams {
	return v.value
}

func (v *NullableBTLoginParams) Set(val *BTLoginParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLoginParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLoginParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLoginParams(val *BTLoginParams) *NullableBTLoginParams {
	return &NullableBTLoginParams{value: val, isSet: true}
}

func (v NullableBTLoginParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLoginParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
