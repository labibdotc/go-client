/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6905-ae59ed040327
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsInfo struct for BTBillOfMaterialsInfo
type BTBillOfMaterialsInfo struct {
	BomSource     *BTBillOfMaterialsSourceInfo  `json:"bomSource,omitempty"`
	CreatedAt     *string                       `json:"createdAt,omitempty"`
	FormatVersion *string                       `json:"formatVersion,omitempty"`
	Headers       []BTBillOfMaterialsHeaderInfo `json:"headers,omitempty"`
	Href          *string                       `json:"href,omitempty"`
	Id            *string                       `json:"id,omitempty"`
	Name          *string                       `json:"name,omitempty"`
	Rows          []BTBillOfMaterialsRowInfo    `json:"rows,omitempty"`
	TemplateId    *string                       `json:"templateId,omitempty"`
	Type          *string                       `json:"type,omitempty"`
	ViewRef       *string                       `json:"viewRef,omitempty"`
}

// NewBTBillOfMaterialsInfo instantiates a new BTBillOfMaterialsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsInfo() *BTBillOfMaterialsInfo {
	this := BTBillOfMaterialsInfo{}
	return &this
}

// NewBTBillOfMaterialsInfoWithDefaults instantiates a new BTBillOfMaterialsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsInfoWithDefaults() *BTBillOfMaterialsInfo {
	this := BTBillOfMaterialsInfo{}
	return &this
}

// GetBomSource returns the BomSource field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetBomSource() BTBillOfMaterialsSourceInfo {
	if o == nil || o.BomSource == nil {
		var ret BTBillOfMaterialsSourceInfo
		return ret
	}
	return *o.BomSource
}

// GetBomSourceOk returns a tuple with the BomSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetBomSourceOk() (*BTBillOfMaterialsSourceInfo, bool) {
	if o == nil || o.BomSource == nil {
		return nil, false
	}
	return o.BomSource, true
}

// HasBomSource returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasBomSource() bool {
	if o != nil && o.BomSource != nil {
		return true
	}

	return false
}

// SetBomSource gets a reference to the given BTBillOfMaterialsSourceInfo and assigns it to the BomSource field.
func (o *BTBillOfMaterialsInfo) SetBomSource(v BTBillOfMaterialsSourceInfo) {
	o.BomSource = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *BTBillOfMaterialsInfo) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *BTBillOfMaterialsInfo) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetHeaders() []BTBillOfMaterialsHeaderInfo {
	if o == nil || o.Headers == nil {
		var ret []BTBillOfMaterialsHeaderInfo
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetHeadersOk() ([]BTBillOfMaterialsHeaderInfo, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []BTBillOfMaterialsHeaderInfo and assigns it to the Headers field.
func (o *BTBillOfMaterialsInfo) SetHeaders(v []BTBillOfMaterialsHeaderInfo) {
	o.Headers = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillOfMaterialsInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTBillOfMaterialsInfo) SetName(v string) {
	o.Name = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetRows() []BTBillOfMaterialsRowInfo {
	if o == nil || o.Rows == nil {
		var ret []BTBillOfMaterialsRowInfo
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetRowsOk() ([]BTBillOfMaterialsRowInfo, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []BTBillOfMaterialsRowInfo and assigns it to the Rows field.
func (o *BTBillOfMaterialsInfo) SetRows(v []BTBillOfMaterialsRowInfo) {
	o.Rows = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *BTBillOfMaterialsInfo) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTBillOfMaterialsInfo) SetType(v string) {
	o.Type = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTBillOfMaterialsInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTBillOfMaterialsInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTBillOfMaterialsInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTBillOfMaterialsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BomSource != nil {
		toSerialize["bomSource"] = o.BomSource
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.FormatVersion != nil {
		toSerialize["formatVersion"] = o.FormatVersion
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsInfo struct {
	value *BTBillOfMaterialsInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsInfo) Get() *BTBillOfMaterialsInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsInfo) Set(val *BTBillOfMaterialsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsInfo(val *BTBillOfMaterialsInfo) *NullableBTBillOfMaterialsInfo {
	return &NullableBTBillOfMaterialsInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
