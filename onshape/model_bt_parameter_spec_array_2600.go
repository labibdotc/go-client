/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18120-f464f720d215
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTParameterSpecArray2600 struct for BTParameterSpecArray2600
type BTParameterSpecArray2600 struct {
	AdditionalLocalizedStrings *int32                             `json:"additionalLocalizedStrings,omitempty"`
	BtType                     *string                            `json:"btType,omitempty"`
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
	UiHints                    []GBTUIHint                        `json:"uiHints,omitempty"`
	VisibilityCondition        *BTParameterVisibilityCondition177 `json:"visibilityCondition,omitempty"`
	DrivenQuery                *string                            `json:"drivenQuery,omitempty"`
	ItemLabelTemplate          *string                            `json:"itemLabelTemplate,omitempty"`
	ItemName                   *string                            `json:"itemName,omitempty"`
	MaxNumberOfPicks           *int32                             `json:"maxNumberOfPicks,omitempty"`
	Parameters                 []BTParameterSpec6                 `json:"parameters,omitempty"`
}

// NewBTParameterSpecArray2600 instantiates a new BTParameterSpecArray2600 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecArray2600() *BTParameterSpecArray2600 {
	this := BTParameterSpecArray2600{}
	return &this
}

// NewBTParameterSpecArray2600WithDefaults instantiates a new BTParameterSpecArray2600 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecArray2600WithDefaults() *BTParameterSpecArray2600 {
	this := BTParameterSpecArray2600{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterSpecArray2600) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecArray2600) SetBtType(v string) {
	o.BtType = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasColumnName() bool {
	if o != nil && o.ColumnName != nil {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *BTParameterSpecArray2600) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetDefaultValue() BTMParameter1 {
	if o == nil || o.DefaultValue == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetDefaultValueOk() (*BTMParameter1, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given BTMParameter1 and assigns it to the DefaultValue field.
func (o *BTParameterSpecArray2600) SetDefaultValue(v BTMParameter1) {
	o.DefaultValue = &v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetIconUri() string {
	if o == nil || o.IconUri == nil {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetIconUriOk() (*string, bool) {
	if o == nil || o.IconUri == nil {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasIconUri() bool {
	if o != nil && o.IconUri != nil {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *BTParameterSpecArray2600) SetIconUri(v string) {
	o.IconUri = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterSpecArray2600) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterSpecArray2600) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetParameterDescription returns the ParameterDescription field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetParameterDescription() string {
	if o == nil || o.ParameterDescription == nil {
		var ret string
		return ret
	}
	return *o.ParameterDescription
}

// GetParameterDescriptionOk returns a tuple with the ParameterDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetParameterDescriptionOk() (*string, bool) {
	if o == nil || o.ParameterDescription == nil {
		return nil, false
	}
	return o.ParameterDescription, true
}

// HasParameterDescription returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasParameterDescription() bool {
	if o != nil && o.ParameterDescription != nil {
		return true
	}

	return false
}

// SetParameterDescription gets a reference to the given string and assigns it to the ParameterDescription field.
func (o *BTParameterSpecArray2600) SetParameterDescription(v string) {
	o.ParameterDescription = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterSpecArray2600) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTParameterSpecArray2600) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterSpecArray2600) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

// GetUiHint returns the UiHint field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetUiHint() string {
	if o == nil || o.UiHint == nil {
		var ret string
		return ret
	}
	return *o.UiHint
}

// GetUiHintOk returns a tuple with the UiHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetUiHintOk() (*string, bool) {
	if o == nil || o.UiHint == nil {
		return nil, false
	}
	return o.UiHint, true
}

// HasUiHint returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasUiHint() bool {
	if o != nil && o.UiHint != nil {
		return true
	}

	return false
}

// SetUiHint gets a reference to the given string and assigns it to the UiHint field.
func (o *BTParameterSpecArray2600) SetUiHint(v string) {
	o.UiHint = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetUiHints() []GBTUIHint {
	if o == nil || o.UiHints == nil {
		var ret []GBTUIHint
		return ret
	}
	return o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetUiHintsOk() ([]GBTUIHint, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given []GBTUIHint and assigns it to the UiHints field.
func (o *BTParameterSpecArray2600) SetUiHints(v []GBTUIHint) {
	o.UiHints = v
}

// GetVisibilityCondition returns the VisibilityCondition field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetVisibilityCondition() BTParameterVisibilityCondition177 {
	if o == nil || o.VisibilityCondition == nil {
		var ret BTParameterVisibilityCondition177
		return ret
	}
	return *o.VisibilityCondition
}

// GetVisibilityConditionOk returns a tuple with the VisibilityCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetVisibilityConditionOk() (*BTParameterVisibilityCondition177, bool) {
	if o == nil || o.VisibilityCondition == nil {
		return nil, false
	}
	return o.VisibilityCondition, true
}

// HasVisibilityCondition returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasVisibilityCondition() bool {
	if o != nil && o.VisibilityCondition != nil {
		return true
	}

	return false
}

// SetVisibilityCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the VisibilityCondition field.
func (o *BTParameterSpecArray2600) SetVisibilityCondition(v BTParameterVisibilityCondition177) {
	o.VisibilityCondition = &v
}

// GetDrivenQuery returns the DrivenQuery field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetDrivenQuery() string {
	if o == nil || o.DrivenQuery == nil {
		var ret string
		return ret
	}
	return *o.DrivenQuery
}

// GetDrivenQueryOk returns a tuple with the DrivenQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetDrivenQueryOk() (*string, bool) {
	if o == nil || o.DrivenQuery == nil {
		return nil, false
	}
	return o.DrivenQuery, true
}

// HasDrivenQuery returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasDrivenQuery() bool {
	if o != nil && o.DrivenQuery != nil {
		return true
	}

	return false
}

// SetDrivenQuery gets a reference to the given string and assigns it to the DrivenQuery field.
func (o *BTParameterSpecArray2600) SetDrivenQuery(v string) {
	o.DrivenQuery = &v
}

// GetItemLabelTemplate returns the ItemLabelTemplate field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetItemLabelTemplate() string {
	if o == nil || o.ItemLabelTemplate == nil {
		var ret string
		return ret
	}
	return *o.ItemLabelTemplate
}

// GetItemLabelTemplateOk returns a tuple with the ItemLabelTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetItemLabelTemplateOk() (*string, bool) {
	if o == nil || o.ItemLabelTemplate == nil {
		return nil, false
	}
	return o.ItemLabelTemplate, true
}

// HasItemLabelTemplate returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasItemLabelTemplate() bool {
	if o != nil && o.ItemLabelTemplate != nil {
		return true
	}

	return false
}

// SetItemLabelTemplate gets a reference to the given string and assigns it to the ItemLabelTemplate field.
func (o *BTParameterSpecArray2600) SetItemLabelTemplate(v string) {
	o.ItemLabelTemplate = &v
}

// GetItemName returns the ItemName field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetItemName() string {
	if o == nil || o.ItemName == nil {
		var ret string
		return ret
	}
	return *o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetItemNameOk() (*string, bool) {
	if o == nil || o.ItemName == nil {
		return nil, false
	}
	return o.ItemName, true
}

// HasItemName returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasItemName() bool {
	if o != nil && o.ItemName != nil {
		return true
	}

	return false
}

// SetItemName gets a reference to the given string and assigns it to the ItemName field.
func (o *BTParameterSpecArray2600) SetItemName(v string) {
	o.ItemName = &v
}

// GetMaxNumberOfPicks returns the MaxNumberOfPicks field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetMaxNumberOfPicks() int32 {
	if o == nil || o.MaxNumberOfPicks == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPicks
}

// GetMaxNumberOfPicksOk returns a tuple with the MaxNumberOfPicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetMaxNumberOfPicksOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfPicks == nil {
		return nil, false
	}
	return o.MaxNumberOfPicks, true
}

// HasMaxNumberOfPicks returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasMaxNumberOfPicks() bool {
	if o != nil && o.MaxNumberOfPicks != nil {
		return true
	}

	return false
}

// SetMaxNumberOfPicks gets a reference to the given int32 and assigns it to the MaxNumberOfPicks field.
func (o *BTParameterSpecArray2600) SetMaxNumberOfPicks(v int32) {
	o.MaxNumberOfPicks = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTParameterSpecArray2600) GetParameters() []BTParameterSpec6 {
	if o == nil || o.Parameters == nil {
		var ret []BTParameterSpec6
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecArray2600) GetParametersOk() ([]BTParameterSpec6, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTParameterSpecArray2600) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTParameterSpec6 and assigns it to the Parameters field.
func (o *BTParameterSpecArray2600) SetParameters(v []BTParameterSpec6) {
	o.Parameters = v
}

func (o BTParameterSpecArray2600) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
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
	if o.DrivenQuery != nil {
		toSerialize["drivenQuery"] = o.DrivenQuery
	}
	if o.ItemLabelTemplate != nil {
		toSerialize["itemLabelTemplate"] = o.ItemLabelTemplate
	}
	if o.ItemName != nil {
		toSerialize["itemName"] = o.ItemName
	}
	if o.MaxNumberOfPicks != nil {
		toSerialize["maxNumberOfPicks"] = o.MaxNumberOfPicks
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecArray2600 struct {
	value *BTParameterSpecArray2600
	isSet bool
}

func (v NullableBTParameterSpecArray2600) Get() *BTParameterSpecArray2600 {
	return v.value
}

func (v *NullableBTParameterSpecArray2600) Set(val *BTParameterSpecArray2600) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecArray2600) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecArray2600) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecArray2600(val *BTParameterSpecArray2600) *NullableBTParameterSpecArray2600 {
	return &NullableBTParameterSpecArray2600{value: val, isSet: true}
}

func (v NullableBTParameterSpecArray2600) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecArray2600) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
