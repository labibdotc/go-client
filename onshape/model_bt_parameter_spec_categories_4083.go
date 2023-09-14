/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTParameterSpecCategories4083 struct for BTParameterSpecCategories4083
type BTParameterSpecCategories4083 struct {
	AdditionalLocalizedStrings     *int32                                        `json:"additionalLocalizedStrings,omitempty"`
	BtType                         *string                                       `json:"btType,omitempty"`
	ColumnName                     *string                                       `json:"columnName,omitempty"`
	DefaultValue                   *BTMParameter1                                `json:"defaultValue,omitempty"`
	IconUri                        *string                                       `json:"iconUri,omitempty"`
	LocalizableName                *string                                       `json:"localizableName,omitempty"`
	LocalizedName                  *string                                       `json:"localizedName,omitempty"`
	ParameterDescription           *string                                       `json:"parameterDescription,omitempty"`
	ParameterId                    *string                                       `json:"parameterId,omitempty"`
	ParameterName                  *string                                       `json:"parameterName,omitempty"`
	QuantityType                   *GBTQuantityType                              `json:"quantityType,omitempty"`
	StringsToLocalize              []string                                      `json:"stringsToLocalize,omitempty"`
	UiHint                         *string                                       `json:"uiHint,omitempty"`
	UiHints                        []GBTUIHint                                   `json:"uiHints,omitempty"`
	VisibilityCondition            *BTParameterVisibilityCondition177            `json:"visibilityCondition,omitempty"`
	EnumName                       *string                                       `json:"enumName,omitempty"`
	EnumValueToVisibilityCondition *map[string]BTParameterVisibilityCondition177 `json:"enumValueToVisibilityCondition,omitempty"`
	Multivalued                    *bool                                         `json:"multivalued,omitempty"`
	Namespace                      *string                                       `json:"namespace,omitempty"`
	OptionNames                    []string                                      `json:"optionNames,omitempty"`
	Options                        []string                                      `json:"options,omitempty"`
	CategoryIdToMetadataTypes      *map[string][]int32                           `json:"categoryIdToMetadataTypes,omitempty"`
}

// NewBTParameterSpecCategories4083 instantiates a new BTParameterSpecCategories4083 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecCategories4083() *BTParameterSpecCategories4083 {
	this := BTParameterSpecCategories4083{}
	return &this
}

// NewBTParameterSpecCategories4083WithDefaults instantiates a new BTParameterSpecCategories4083 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecCategories4083WithDefaults() *BTParameterSpecCategories4083 {
	this := BTParameterSpecCategories4083{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterSpecCategories4083) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecCategories4083) SetBtType(v string) {
	o.BtType = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasColumnName() bool {
	if o != nil && o.ColumnName != nil {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *BTParameterSpecCategories4083) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetDefaultValue() BTMParameter1 {
	if o == nil || o.DefaultValue == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetDefaultValueOk() (*BTMParameter1, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given BTMParameter1 and assigns it to the DefaultValue field.
func (o *BTParameterSpecCategories4083) SetDefaultValue(v BTMParameter1) {
	o.DefaultValue = &v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetIconUri() string {
	if o == nil || o.IconUri == nil {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetIconUriOk() (*string, bool) {
	if o == nil || o.IconUri == nil {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasIconUri() bool {
	if o != nil && o.IconUri != nil {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *BTParameterSpecCategories4083) SetIconUri(v string) {
	o.IconUri = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterSpecCategories4083) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterSpecCategories4083) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetParameterDescription returns the ParameterDescription field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetParameterDescription() string {
	if o == nil || o.ParameterDescription == nil {
		var ret string
		return ret
	}
	return *o.ParameterDescription
}

// GetParameterDescriptionOk returns a tuple with the ParameterDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetParameterDescriptionOk() (*string, bool) {
	if o == nil || o.ParameterDescription == nil {
		return nil, false
	}
	return o.ParameterDescription, true
}

// HasParameterDescription returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasParameterDescription() bool {
	if o != nil && o.ParameterDescription != nil {
		return true
	}

	return false
}

// SetParameterDescription gets a reference to the given string and assigns it to the ParameterDescription field.
func (o *BTParameterSpecCategories4083) SetParameterDescription(v string) {
	o.ParameterDescription = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterSpecCategories4083) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTParameterSpecCategories4083) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetQuantityType returns the QuantityType field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetQuantityType() GBTQuantityType {
	if o == nil || o.QuantityType == nil {
		var ret GBTQuantityType
		return ret
	}
	return *o.QuantityType
}

// GetQuantityTypeOk returns a tuple with the QuantityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetQuantityTypeOk() (*GBTQuantityType, bool) {
	if o == nil || o.QuantityType == nil {
		return nil, false
	}
	return o.QuantityType, true
}

// HasQuantityType returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasQuantityType() bool {
	if o != nil && o.QuantityType != nil {
		return true
	}

	return false
}

// SetQuantityType gets a reference to the given GBTQuantityType and assigns it to the QuantityType field.
func (o *BTParameterSpecCategories4083) SetQuantityType(v GBTQuantityType) {
	o.QuantityType = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterSpecCategories4083) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

// GetUiHint returns the UiHint field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetUiHint() string {
	if o == nil || o.UiHint == nil {
		var ret string
		return ret
	}
	return *o.UiHint
}

// GetUiHintOk returns a tuple with the UiHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetUiHintOk() (*string, bool) {
	if o == nil || o.UiHint == nil {
		return nil, false
	}
	return o.UiHint, true
}

// HasUiHint returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasUiHint() bool {
	if o != nil && o.UiHint != nil {
		return true
	}

	return false
}

// SetUiHint gets a reference to the given string and assigns it to the UiHint field.
func (o *BTParameterSpecCategories4083) SetUiHint(v string) {
	o.UiHint = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetUiHints() []GBTUIHint {
	if o == nil || o.UiHints == nil {
		var ret []GBTUIHint
		return ret
	}
	return o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetUiHintsOk() ([]GBTUIHint, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given []GBTUIHint and assigns it to the UiHints field.
func (o *BTParameterSpecCategories4083) SetUiHints(v []GBTUIHint) {
	o.UiHints = v
}

// GetVisibilityCondition returns the VisibilityCondition field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetVisibilityCondition() BTParameterVisibilityCondition177 {
	if o == nil || o.VisibilityCondition == nil {
		var ret BTParameterVisibilityCondition177
		return ret
	}
	return *o.VisibilityCondition
}

// GetVisibilityConditionOk returns a tuple with the VisibilityCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetVisibilityConditionOk() (*BTParameterVisibilityCondition177, bool) {
	if o == nil || o.VisibilityCondition == nil {
		return nil, false
	}
	return o.VisibilityCondition, true
}

// HasVisibilityCondition returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasVisibilityCondition() bool {
	if o != nil && o.VisibilityCondition != nil {
		return true
	}

	return false
}

// SetVisibilityCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the VisibilityCondition field.
func (o *BTParameterSpecCategories4083) SetVisibilityCondition(v BTParameterVisibilityCondition177) {
	o.VisibilityCondition = &v
}

// GetEnumName returns the EnumName field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetEnumName() string {
	if o == nil || o.EnumName == nil {
		var ret string
		return ret
	}
	return *o.EnumName
}

// GetEnumNameOk returns a tuple with the EnumName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetEnumNameOk() (*string, bool) {
	if o == nil || o.EnumName == nil {
		return nil, false
	}
	return o.EnumName, true
}

// HasEnumName returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasEnumName() bool {
	if o != nil && o.EnumName != nil {
		return true
	}

	return false
}

// SetEnumName gets a reference to the given string and assigns it to the EnumName field.
func (o *BTParameterSpecCategories4083) SetEnumName(v string) {
	o.EnumName = &v
}

// GetEnumValueToVisibilityCondition returns the EnumValueToVisibilityCondition field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetEnumValueToVisibilityCondition() map[string]BTParameterVisibilityCondition177 {
	if o == nil || o.EnumValueToVisibilityCondition == nil {
		var ret map[string]BTParameterVisibilityCondition177
		return ret
	}
	return *o.EnumValueToVisibilityCondition
}

// GetEnumValueToVisibilityConditionOk returns a tuple with the EnumValueToVisibilityCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetEnumValueToVisibilityConditionOk() (*map[string]BTParameterVisibilityCondition177, bool) {
	if o == nil || o.EnumValueToVisibilityCondition == nil {
		return nil, false
	}
	return o.EnumValueToVisibilityCondition, true
}

// HasEnumValueToVisibilityCondition returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasEnumValueToVisibilityCondition() bool {
	if o != nil && o.EnumValueToVisibilityCondition != nil {
		return true
	}

	return false
}

// SetEnumValueToVisibilityCondition gets a reference to the given map[string]BTParameterVisibilityCondition177 and assigns it to the EnumValueToVisibilityCondition field.
func (o *BTParameterSpecCategories4083) SetEnumValueToVisibilityCondition(v map[string]BTParameterVisibilityCondition177) {
	o.EnumValueToVisibilityCondition = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetMultivalued() bool {
	if o == nil || o.Multivalued == nil {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetMultivaluedOk() (*bool, bool) {
	if o == nil || o.Multivalued == nil {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasMultivalued() bool {
	if o != nil && o.Multivalued != nil {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *BTParameterSpecCategories4083) SetMultivalued(v bool) {
	o.Multivalued = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTParameterSpecCategories4083) SetNamespace(v string) {
	o.Namespace = &v
}

// GetOptionNames returns the OptionNames field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetOptionNames() []string {
	if o == nil || o.OptionNames == nil {
		var ret []string
		return ret
	}
	return o.OptionNames
}

// GetOptionNamesOk returns a tuple with the OptionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetOptionNamesOk() ([]string, bool) {
	if o == nil || o.OptionNames == nil {
		return nil, false
	}
	return o.OptionNames, true
}

// HasOptionNames returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasOptionNames() bool {
	if o != nil && o.OptionNames != nil {
		return true
	}

	return false
}

// SetOptionNames gets a reference to the given []string and assigns it to the OptionNames field.
func (o *BTParameterSpecCategories4083) SetOptionNames(v []string) {
	o.OptionNames = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetOptions() []string {
	if o == nil || o.Options == nil {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetOptionsOk() ([]string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *BTParameterSpecCategories4083) SetOptions(v []string) {
	o.Options = v
}

// GetCategoryIdToMetadataTypes returns the CategoryIdToMetadataTypes field value if set, zero value otherwise.
func (o *BTParameterSpecCategories4083) GetCategoryIdToMetadataTypes() map[string][]int32 {
	if o == nil || o.CategoryIdToMetadataTypes == nil {
		var ret map[string][]int32
		return ret
	}
	return *o.CategoryIdToMetadataTypes
}

// GetCategoryIdToMetadataTypesOk returns a tuple with the CategoryIdToMetadataTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecCategories4083) GetCategoryIdToMetadataTypesOk() (*map[string][]int32, bool) {
	if o == nil || o.CategoryIdToMetadataTypes == nil {
		return nil, false
	}
	return o.CategoryIdToMetadataTypes, true
}

// HasCategoryIdToMetadataTypes returns a boolean if a field has been set.
func (o *BTParameterSpecCategories4083) HasCategoryIdToMetadataTypes() bool {
	if o != nil && o.CategoryIdToMetadataTypes != nil {
		return true
	}

	return false
}

// SetCategoryIdToMetadataTypes gets a reference to the given map[string][]int32 and assigns it to the CategoryIdToMetadataTypes field.
func (o *BTParameterSpecCategories4083) SetCategoryIdToMetadataTypes(v map[string][]int32) {
	o.CategoryIdToMetadataTypes = &v
}

func (o BTParameterSpecCategories4083) MarshalJSON() ([]byte, error) {
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
	if o.QuantityType != nil {
		toSerialize["quantityType"] = o.QuantityType
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
	if o.EnumName != nil {
		toSerialize["enumName"] = o.EnumName
	}
	if o.EnumValueToVisibilityCondition != nil {
		toSerialize["enumValueToVisibilityCondition"] = o.EnumValueToVisibilityCondition
	}
	if o.Multivalued != nil {
		toSerialize["multivalued"] = o.Multivalued
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.OptionNames != nil {
		toSerialize["optionNames"] = o.OptionNames
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.CategoryIdToMetadataTypes != nil {
		toSerialize["categoryIdToMetadataTypes"] = o.CategoryIdToMetadataTypes
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecCategories4083 struct {
	value *BTParameterSpecCategories4083
	isSet bool
}

func (v NullableBTParameterSpecCategories4083) Get() *BTParameterSpecCategories4083 {
	return v.value
}

func (v *NullableBTParameterSpecCategories4083) Set(val *BTParameterSpecCategories4083) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecCategories4083) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecCategories4083) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecCategories4083(val *BTParameterSpecCategories4083) *NullableBTParameterSpecCategories4083 {
	return &NullableBTParameterSpecCategories4083{value: val, isSet: true}
}

func (v NullableBTParameterSpecCategories4083) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecCategories4083) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
