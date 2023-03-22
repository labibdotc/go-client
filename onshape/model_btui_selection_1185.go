/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13258-6b30d30337fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUiSelection1185 struct for BTUiSelection1185
type BTUiSelection1185 struct {
	DeterministicIdList []string        `json:"deterministicIdList,omitempty"`
	Id                  *string         `json:"id,omitempty"`
	Occurrence          *BTOccurrence74 `json:"occurrence,omitempty"`
	TableRowId          *string         `json:"tableRowId,omitempty"`
	Type                *string         `json:"type,omitempty"`
}

// NewBTUiSelection1185 instantiates a new BTUiSelection1185 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUiSelection1185() *BTUiSelection1185 {
	this := BTUiSelection1185{}
	return &this
}

// NewBTUiSelection1185WithDefaults instantiates a new BTUiSelection1185 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUiSelection1185WithDefaults() *BTUiSelection1185 {
	this := BTUiSelection1185{}
	return &this
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTUiSelection1185) GetDeterministicIdList() []string {
	if o == nil || o.DeterministicIdList == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiSelection1185) GetDeterministicIdListOk() ([]string, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTUiSelection1185) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given []string and assigns it to the DeterministicIdList field.
func (o *BTUiSelection1185) SetDeterministicIdList(v []string) {
	o.DeterministicIdList = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUiSelection1185) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiSelection1185) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUiSelection1185) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUiSelection1185) SetId(v string) {
	o.Id = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTUiSelection1185) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiSelection1185) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTUiSelection1185) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTUiSelection1185) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetTableRowId returns the TableRowId field value if set, zero value otherwise.
func (o *BTUiSelection1185) GetTableRowId() string {
	if o == nil || o.TableRowId == nil {
		var ret string
		return ret
	}
	return *o.TableRowId
}

// GetTableRowIdOk returns a tuple with the TableRowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiSelection1185) GetTableRowIdOk() (*string, bool) {
	if o == nil || o.TableRowId == nil {
		return nil, false
	}
	return o.TableRowId, true
}

// HasTableRowId returns a boolean if a field has been set.
func (o *BTUiSelection1185) HasTableRowId() bool {
	if o != nil && o.TableRowId != nil {
		return true
	}

	return false
}

// SetTableRowId gets a reference to the given string and assigns it to the TableRowId field.
func (o *BTUiSelection1185) SetTableRowId(v string) {
	o.TableRowId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTUiSelection1185) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiSelection1185) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTUiSelection1185) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTUiSelection1185) SetType(v string) {
	o.Type = &v
}

func (o BTUiSelection1185) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.TableRowId != nil {
		toSerialize["tableRowId"] = o.TableRowId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTUiSelection1185 struct {
	value *BTUiSelection1185
	isSet bool
}

func (v NullableBTUiSelection1185) Get() *BTUiSelection1185 {
	return v.value
}

func (v *NullableBTUiSelection1185) Set(val *BTUiSelection1185) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUiSelection1185) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUiSelection1185) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUiSelection1185(val *BTUiSelection1185) *NullableBTUiSelection1185 {
	return &NullableBTUiSelection1185{value: val, isSet: true}
}

func (v NullableBTUiSelection1185) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUiSelection1185) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
