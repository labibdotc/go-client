/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6730-405400b0583f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTLazilyParsedFeatureScript struct for BTLazilyParsedFeatureScript
type BTLazilyParsedFeatureScript struct {
	Checksum              *BTUiFeatureStudioChecksum2438  `json:"checksum,omitempty"`
	LanguageVersion       *int32                          `json:"languageVersion,omitempty"`
	Lines                 *Lines                          `json:"lines,omitempty"`
	Model                 *BTMModel141                    `json:"model,omitempty"`
	Module                *BTPModule234                   `json:"module,omitempty"`
	ModuleId              *BTPModuleId235                 `json:"moduleId,omitempty"`
	NoticeModuleIds       *BTPModuleId235                 `json:"noticeModuleIds,omitempty"`
	ParentLanguageVersion *int32                          `json:"parentLanguageVersion,omitempty"`
	References            *map[string]BTMicroversionId366 `json:"references,omitempty"`
	Source                *string                         `json:"source,omitempty"`
}

// NewBTLazilyParsedFeatureScript instantiates a new BTLazilyParsedFeatureScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLazilyParsedFeatureScript() *BTLazilyParsedFeatureScript {
	this := BTLazilyParsedFeatureScript{}
	return &this
}

// NewBTLazilyParsedFeatureScriptWithDefaults instantiates a new BTLazilyParsedFeatureScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLazilyParsedFeatureScriptWithDefaults() *BTLazilyParsedFeatureScript {
	this := BTLazilyParsedFeatureScript{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetChecksum() BTUiFeatureStudioChecksum2438 {
	if o == nil || o.Checksum == nil {
		var ret BTUiFeatureStudioChecksum2438
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetChecksumOk() (*BTUiFeatureStudioChecksum2438, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given BTUiFeatureStudioChecksum2438 and assigns it to the Checksum field.
func (o *BTLazilyParsedFeatureScript) SetChecksum(v BTUiFeatureStudioChecksum2438) {
	o.Checksum = &v
}

// GetLanguageVersion returns the LanguageVersion field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetLanguageVersion() int32 {
	if o == nil || o.LanguageVersion == nil {
		var ret int32
		return ret
	}
	return *o.LanguageVersion
}

// GetLanguageVersionOk returns a tuple with the LanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetLanguageVersionOk() (*int32, bool) {
	if o == nil || o.LanguageVersion == nil {
		return nil, false
	}
	return o.LanguageVersion, true
}

// HasLanguageVersion returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasLanguageVersion() bool {
	if o != nil && o.LanguageVersion != nil {
		return true
	}

	return false
}

// SetLanguageVersion gets a reference to the given int32 and assigns it to the LanguageVersion field.
func (o *BTLazilyParsedFeatureScript) SetLanguageVersion(v int32) {
	o.LanguageVersion = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetLines() Lines {
	if o == nil || o.Lines == nil {
		var ret Lines
		return ret
	}
	return *o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetLinesOk() (*Lines, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given Lines and assigns it to the Lines field.
func (o *BTLazilyParsedFeatureScript) SetLines(v Lines) {
	o.Lines = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetModel() BTMModel141 {
	if o == nil || o.Model == nil {
		var ret BTMModel141
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetModelOk() (*BTMModel141, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given BTMModel141 and assigns it to the Model field.
func (o *BTLazilyParsedFeatureScript) SetModel(v BTMModel141) {
	o.Model = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetModule() BTPModule234 {
	if o == nil || o.Module == nil {
		var ret BTPModule234
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetModuleOk() (*BTPModule234, bool) {
	if o == nil || o.Module == nil {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasModule() bool {
	if o != nil && o.Module != nil {
		return true
	}

	return false
}

// SetModule gets a reference to the given BTPModule234 and assigns it to the Module field.
func (o *BTLazilyParsedFeatureScript) SetModule(v BTPModule234) {
	o.Module = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetModuleId() BTPModuleId235 {
	if o == nil || o.ModuleId == nil {
		var ret BTPModuleId235
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetModuleIdOk() (*BTPModuleId235, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given BTPModuleId235 and assigns it to the ModuleId field.
func (o *BTLazilyParsedFeatureScript) SetModuleId(v BTPModuleId235) {
	o.ModuleId = &v
}

// GetNoticeModuleIds returns the NoticeModuleIds field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetNoticeModuleIds() BTPModuleId235 {
	if o == nil || o.NoticeModuleIds == nil {
		var ret BTPModuleId235
		return ret
	}
	return *o.NoticeModuleIds
}

// GetNoticeModuleIdsOk returns a tuple with the NoticeModuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetNoticeModuleIdsOk() (*BTPModuleId235, bool) {
	if o == nil || o.NoticeModuleIds == nil {
		return nil, false
	}
	return o.NoticeModuleIds, true
}

// HasNoticeModuleIds returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasNoticeModuleIds() bool {
	if o != nil && o.NoticeModuleIds != nil {
		return true
	}

	return false
}

// SetNoticeModuleIds gets a reference to the given BTPModuleId235 and assigns it to the NoticeModuleIds field.
func (o *BTLazilyParsedFeatureScript) SetNoticeModuleIds(v BTPModuleId235) {
	o.NoticeModuleIds = &v
}

// GetParentLanguageVersion returns the ParentLanguageVersion field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetParentLanguageVersion() int32 {
	if o == nil || o.ParentLanguageVersion == nil {
		var ret int32
		return ret
	}
	return *o.ParentLanguageVersion
}

// GetParentLanguageVersionOk returns a tuple with the ParentLanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetParentLanguageVersionOk() (*int32, bool) {
	if o == nil || o.ParentLanguageVersion == nil {
		return nil, false
	}
	return o.ParentLanguageVersion, true
}

// HasParentLanguageVersion returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasParentLanguageVersion() bool {
	if o != nil && o.ParentLanguageVersion != nil {
		return true
	}

	return false
}

// SetParentLanguageVersion gets a reference to the given int32 and assigns it to the ParentLanguageVersion field.
func (o *BTLazilyParsedFeatureScript) SetParentLanguageVersion(v int32) {
	o.ParentLanguageVersion = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetReferences() map[string]BTMicroversionId366 {
	if o == nil || o.References == nil {
		var ret map[string]BTMicroversionId366
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetReferencesOk() (*map[string]BTMicroversionId366, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given map[string]BTMicroversionId366 and assigns it to the References field.
func (o *BTLazilyParsedFeatureScript) SetReferences(v map[string]BTMicroversionId366) {
	o.References = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTLazilyParsedFeatureScript) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLazilyParsedFeatureScript) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTLazilyParsedFeatureScript) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *BTLazilyParsedFeatureScript) SetSource(v string) {
	o.Source = &v
}

func (o BTLazilyParsedFeatureScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.LanguageVersion != nil {
		toSerialize["languageVersion"] = o.LanguageVersion
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Module != nil {
		toSerialize["module"] = o.Module
	}
	if o.ModuleId != nil {
		toSerialize["moduleId"] = o.ModuleId
	}
	if o.NoticeModuleIds != nil {
		toSerialize["noticeModuleIds"] = o.NoticeModuleIds
	}
	if o.ParentLanguageVersion != nil {
		toSerialize["parentLanguageVersion"] = o.ParentLanguageVersion
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableBTLazilyParsedFeatureScript struct {
	value *BTLazilyParsedFeatureScript
	isSet bool
}

func (v NullableBTLazilyParsedFeatureScript) Get() *BTLazilyParsedFeatureScript {
	return v.value
}

func (v *NullableBTLazilyParsedFeatureScript) Set(val *BTLazilyParsedFeatureScript) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLazilyParsedFeatureScript) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLazilyParsedFeatureScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLazilyParsedFeatureScript(val *BTLazilyParsedFeatureScript) *NullableBTLazilyParsedFeatureScript {
	return &NullableBTLazilyParsedFeatureScript{value: val, isSet: true}
}

func (v NullableBTLazilyParsedFeatureScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLazilyParsedFeatureScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
