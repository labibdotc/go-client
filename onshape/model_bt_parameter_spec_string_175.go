/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7250-f7937557e62d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTParameterSpecString175 struct for BTParameterSpecString175
type BTParameterSpecString175 struct {
	AdditionalLocalizedStrings *int32                             `json:"additionalLocalizedStrings,omitempty"`
	ColumnName                 *string                            `json:"columnName,omitempty"`
	DefaultValue               *BTMParameter1                     `json:"defaultValue,omitempty"`
	IconUri                    *string                            `json:"iconUri,omitempty"`
	LocalizableName            *string                            `json:"localizableName,omitempty"`
	LocalizedName              *string                            `json:"localizedName,omitempty"`
	ParameterDescription       *string                            `json:"parameterDescription,omitempty"`
	ParameterId                *string                            `json:"parameterId,omitempty"`
	ParameterName              *string                            `json:"parameterName,omitempty"`
	StringsToLocalize          []string                           `json:"stringsToLocalize,omitempty"`
	UiHint                     *string                            `json:"uiHint,omitempty"`
	UiHints                    []string                           `json:"uiHints,omitempty"`
	VisibilityCondition        *BTParameterVisibilityCondition177 `json:"visibilityCondition,omitempty"`
	BtType                     *string                            `json:"btType,omitempty"`
	DisallowsEmptyStrings      *string                            `json:"disallowsEmptyStrings,omitempty"`
	DisallowsHtml              *string                            `json:"disallowsHtml,omitempty"`
	FormatConditions           []BTStringFormatCondition683       `json:"formatConditions,omitempty"`
}

// NewBTParameterSpecString175 instantiates a new BTParameterSpecString175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecString175() *BTParameterSpecString175 {
	this := BTParameterSpecString175{}
	return &this
}

// NewBTParameterSpecString175WithDefaults instantiates a new BTParameterSpecString175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecString175WithDefaults() *BTParameterSpecString175 {
	this := BTParameterSpecString175{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterSpecString175) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasColumnName() bool {
	if o != nil && o.ColumnName != nil {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *BTParameterSpecString175) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetDefaultValue() BTMParameter1 {
	if o == nil || o.DefaultValue == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetDefaultValueOk() (*BTMParameter1, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given BTMParameter1 and assigns it to the DefaultValue field.
func (o *BTParameterSpecString175) SetDefaultValue(v BTMParameter1) {
	o.DefaultValue = &v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetIconUri() string {
	if o == nil || o.IconUri == nil {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetIconUriOk() (*string, bool) {
	if o == nil || o.IconUri == nil {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasIconUri() bool {
	if o != nil && o.IconUri != nil {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *BTParameterSpecString175) SetIconUri(v string) {
	o.IconUri = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterSpecString175) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterSpecString175) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetParameterDescription returns the ParameterDescription field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetParameterDescription() string {
	if o == nil || o.ParameterDescription == nil {
		var ret string
		return ret
	}
	return *o.ParameterDescription
}

// GetParameterDescriptionOk returns a tuple with the ParameterDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetParameterDescriptionOk() (*string, bool) {
	if o == nil || o.ParameterDescription == nil {
		return nil, false
	}
	return o.ParameterDescription, true
}

// HasParameterDescription returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasParameterDescription() bool {
	if o != nil && o.ParameterDescription != nil {
		return true
	}

	return false
}

// SetParameterDescription gets a reference to the given string and assigns it to the ParameterDescription field.
func (o *BTParameterSpecString175) SetParameterDescription(v string) {
	o.ParameterDescription = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterSpecString175) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTParameterSpecString175) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterSpecString175) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

// GetUiHint returns the UiHint field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetUiHint() string {
	if o == nil || o.UiHint == nil {
		var ret string
		return ret
	}
	return *o.UiHint
}

// GetUiHintOk returns a tuple with the UiHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetUiHintOk() (*string, bool) {
	if o == nil || o.UiHint == nil {
		return nil, false
	}
	return o.UiHint, true
}

// HasUiHint returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasUiHint() bool {
	if o != nil && o.UiHint != nil {
		return true
	}

	return false
}

// SetUiHint gets a reference to the given string and assigns it to the UiHint field.
func (o *BTParameterSpecString175) SetUiHint(v string) {
	o.UiHint = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetUiHints() []string {
	if o == nil || o.UiHints == nil {
		var ret []string
		return ret
	}
	return o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetUiHintsOk() ([]string, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given []string and assigns it to the UiHints field.
func (o *BTParameterSpecString175) SetUiHints(v []string) {
	o.UiHints = v
}

// GetVisibilityCondition returns the VisibilityCondition field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetVisibilityCondition() BTParameterVisibilityCondition177 {
	if o == nil || o.VisibilityCondition == nil {
		var ret BTParameterVisibilityCondition177
		return ret
	}
	return *o.VisibilityCondition
}

// GetVisibilityConditionOk returns a tuple with the VisibilityCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetVisibilityConditionOk() (*BTParameterVisibilityCondition177, bool) {
	if o == nil || o.VisibilityCondition == nil {
		return nil, false
	}
	return o.VisibilityCondition, true
}

// HasVisibilityCondition returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasVisibilityCondition() bool {
	if o != nil && o.VisibilityCondition != nil {
		return true
	}

	return false
}

// SetVisibilityCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the VisibilityCondition field.
func (o *BTParameterSpecString175) SetVisibilityCondition(v BTParameterVisibilityCondition177) {
	o.VisibilityCondition = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecString175) SetBtType(v string) {
	o.BtType = &v
}

// GetDisallowsEmptyStrings returns the DisallowsEmptyStrings field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetDisallowsEmptyStrings() string {
	if o == nil || o.DisallowsEmptyStrings == nil {
		var ret string
		return ret
	}
	return *o.DisallowsEmptyStrings
}

// GetDisallowsEmptyStringsOk returns a tuple with the DisallowsEmptyStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetDisallowsEmptyStringsOk() (*string, bool) {
	if o == nil || o.DisallowsEmptyStrings == nil {
		return nil, false
	}
	return o.DisallowsEmptyStrings, true
}

// HasDisallowsEmptyStrings returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasDisallowsEmptyStrings() bool {
	if o != nil && o.DisallowsEmptyStrings != nil {
		return true
	}

	return false
}

// SetDisallowsEmptyStrings gets a reference to the given string and assigns it to the DisallowsEmptyStrings field.
func (o *BTParameterSpecString175) SetDisallowsEmptyStrings(v string) {
	o.DisallowsEmptyStrings = &v
}

// GetDisallowsHtml returns the DisallowsHtml field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetDisallowsHtml() string {
	if o == nil || o.DisallowsHtml == nil {
		var ret string
		return ret
	}
	return *o.DisallowsHtml
}

// GetDisallowsHtmlOk returns a tuple with the DisallowsHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetDisallowsHtmlOk() (*string, bool) {
	if o == nil || o.DisallowsHtml == nil {
		return nil, false
	}
	return o.DisallowsHtml, true
}

// HasDisallowsHtml returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasDisallowsHtml() bool {
	if o != nil && o.DisallowsHtml != nil {
		return true
	}

	return false
}

// SetDisallowsHtml gets a reference to the given string and assigns it to the DisallowsHtml field.
func (o *BTParameterSpecString175) SetDisallowsHtml(v string) {
	o.DisallowsHtml = &v
}

// GetFormatConditions returns the FormatConditions field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetFormatConditions() []BTStringFormatCondition683 {
	if o == nil || o.FormatConditions == nil {
		var ret []BTStringFormatCondition683
		return ret
	}
	return o.FormatConditions
}

// GetFormatConditionsOk returns a tuple with the FormatConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetFormatConditionsOk() ([]BTStringFormatCondition683, bool) {
	if o == nil || o.FormatConditions == nil {
		return nil, false
	}
	return o.FormatConditions, true
}

// HasFormatConditions returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasFormatConditions() bool {
	if o != nil && o.FormatConditions != nil {
		return true
	}

	return false
}

// SetFormatConditions gets a reference to the given []BTStringFormatCondition683 and assigns it to the FormatConditions field.
func (o *BTParameterSpecString175) SetFormatConditions(v []BTStringFormatCondition683) {
	o.FormatConditions = v
}

func (o BTParameterSpecString175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.ColumnName != nil {
		toSerialize["columnName"] = o.ColumnName
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.IconUri != nil {
		toSerialize["iconUri"] = o.IconUri
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.ParameterDescription != nil {
		toSerialize["parameterDescription"] = o.ParameterDescription
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	if o.UiHint != nil {
		toSerialize["uiHint"] = o.UiHint
	}
	if o.UiHints != nil {
		toSerialize["uiHints"] = o.UiHints
	}
	if o.VisibilityCondition != nil {
		toSerialize["visibilityCondition"] = o.VisibilityCondition
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DisallowsEmptyStrings != nil {
		toSerialize["disallowsEmptyStrings"] = o.DisallowsEmptyStrings
	}
	if o.DisallowsHtml != nil {
		toSerialize["disallowsHtml"] = o.DisallowsHtml
	}
	if o.FormatConditions != nil {
		toSerialize["formatConditions"] = o.FormatConditions
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecString175 struct {
	value *BTParameterSpecString175
	isSet bool
}

func (v NullableBTParameterSpecString175) Get() *BTParameterSpecString175 {
	return v.value
}

func (v *NullableBTParameterSpecString175) Set(val *BTParameterSpecString175) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecString175) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecString175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecString175(val *BTParameterSpecString175) *NullableBTParameterSpecString175 {
	return &NullableBTParameterSpecString175{value: val, isSet: true}
}

func (v NullableBTParameterSpecString175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecString175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
