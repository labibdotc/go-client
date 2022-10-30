/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7108-42dac6f29840
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementUpdateParams struct for BTAppElementUpdateParams
type BTAppElementUpdateParams struct {
	// Edits to be applied to the element's subelement data.
	Changes []BTAppElementChangeParams `json:"changes,omitempty"`
	// The label that will appear in the document's edit history for this operation. If blank, a value will be auto-generated.
	Description  *string      `json:"description,omitempty"`
	JsonTreeEdit *BTJEdit3734 `json:"jsonTreeEdit,omitempty"`
	// The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint.
	ParentChangeId *string `json:"parentChangeId,omitempty"`
	// Edits to be applied to the element's metadata.
	PropertyUpdates []BTMetadataPropertyUpdateParams `json:"propertyUpdates,omitempty"`
	// If true, errors in request processing will be returned in a 200 response with a meaningful description. Otherwise, errors will result in a relevant HTTP error response.
	ReturnError *bool `json:"returnError,omitempty"`
	// If specified, and jsonTreeEdit is non-empty, the json difference will be returned in this format, otherwise no json difference will be returned.
	ReturnJsonDifferenceFormat *string `json:"returnJsonDifferenceFormat,omitempty"`
	// The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API.
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewBTAppElementUpdateParams instantiates a new BTAppElementUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementUpdateParams() *BTAppElementUpdateParams {
	this := BTAppElementUpdateParams{}
	return &this
}

// NewBTAppElementUpdateParamsWithDefaults instantiates a new BTAppElementUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementUpdateParamsWithDefaults() *BTAppElementUpdateParams {
	this := BTAppElementUpdateParams{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetChanges() []BTAppElementChangeParams {
	if o == nil || o.Changes == nil {
		var ret []BTAppElementChangeParams
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetChangesOk() ([]BTAppElementChangeParams, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []BTAppElementChangeParams and assigns it to the Changes field.
func (o *BTAppElementUpdateParams) SetChanges(v []BTAppElementChangeParams) {
	o.Changes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementUpdateParams) SetDescription(v string) {
	o.Description = &v
}

// GetJsonTreeEdit returns the JsonTreeEdit field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetJsonTreeEdit() BTJEdit3734 {
	if o == nil || o.JsonTreeEdit == nil {
		var ret BTJEdit3734
		return ret
	}
	return *o.JsonTreeEdit
}

// GetJsonTreeEditOk returns a tuple with the JsonTreeEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetJsonTreeEditOk() (*BTJEdit3734, bool) {
	if o == nil || o.JsonTreeEdit == nil {
		return nil, false
	}
	return o.JsonTreeEdit, true
}

// HasJsonTreeEdit returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasJsonTreeEdit() bool {
	if o != nil && o.JsonTreeEdit != nil {
		return true
	}

	return false
}

// SetJsonTreeEdit gets a reference to the given BTJEdit3734 and assigns it to the JsonTreeEdit field.
func (o *BTAppElementUpdateParams) SetJsonTreeEdit(v BTJEdit3734) {
	o.JsonTreeEdit = &v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementUpdateParams) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

// GetPropertyUpdates returns the PropertyUpdates field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetPropertyUpdates() []BTMetadataPropertyUpdateParams {
	if o == nil || o.PropertyUpdates == nil {
		var ret []BTMetadataPropertyUpdateParams
		return ret
	}
	return o.PropertyUpdates
}

// GetPropertyUpdatesOk returns a tuple with the PropertyUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetPropertyUpdatesOk() ([]BTMetadataPropertyUpdateParams, bool) {
	if o == nil || o.PropertyUpdates == nil {
		return nil, false
	}
	return o.PropertyUpdates, true
}

// HasPropertyUpdates returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasPropertyUpdates() bool {
	if o != nil && o.PropertyUpdates != nil {
		return true
	}

	return false
}

// SetPropertyUpdates gets a reference to the given []BTMetadataPropertyUpdateParams and assigns it to the PropertyUpdates field.
func (o *BTAppElementUpdateParams) SetPropertyUpdates(v []BTMetadataPropertyUpdateParams) {
	o.PropertyUpdates = v
}

// GetReturnError returns the ReturnError field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetReturnError() bool {
	if o == nil || o.ReturnError == nil {
		var ret bool
		return ret
	}
	return *o.ReturnError
}

// GetReturnErrorOk returns a tuple with the ReturnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetReturnErrorOk() (*bool, bool) {
	if o == nil || o.ReturnError == nil {
		return nil, false
	}
	return o.ReturnError, true
}

// HasReturnError returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasReturnError() bool {
	if o != nil && o.ReturnError != nil {
		return true
	}

	return false
}

// SetReturnError gets a reference to the given bool and assigns it to the ReturnError field.
func (o *BTAppElementUpdateParams) SetReturnError(v bool) {
	o.ReturnError = &v
}

// GetReturnJsonDifferenceFormat returns the ReturnJsonDifferenceFormat field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetReturnJsonDifferenceFormat() string {
	if o == nil || o.ReturnJsonDifferenceFormat == nil {
		var ret string
		return ret
	}
	return *o.ReturnJsonDifferenceFormat
}

// GetReturnJsonDifferenceFormatOk returns a tuple with the ReturnJsonDifferenceFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetReturnJsonDifferenceFormatOk() (*string, bool) {
	if o == nil || o.ReturnJsonDifferenceFormat == nil {
		return nil, false
	}
	return o.ReturnJsonDifferenceFormat, true
}

// HasReturnJsonDifferenceFormat returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasReturnJsonDifferenceFormat() bool {
	if o != nil && o.ReturnJsonDifferenceFormat != nil {
		return true
	}

	return false
}

// SetReturnJsonDifferenceFormat gets a reference to the given string and assigns it to the ReturnJsonDifferenceFormat field.
func (o *BTAppElementUpdateParams) SetReturnJsonDifferenceFormat(v string) {
	o.ReturnJsonDifferenceFormat = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BTAppElementUpdateParams) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementUpdateParams) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BTAppElementUpdateParams) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *BTAppElementUpdateParams) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o BTAppElementUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.JsonTreeEdit != nil {
		toSerialize["jsonTreeEdit"] = o.JsonTreeEdit
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	if o.PropertyUpdates != nil {
		toSerialize["propertyUpdates"] = o.PropertyUpdates
	}
	if o.ReturnError != nil {
		toSerialize["returnError"] = o.ReturnError
	}
	if o.ReturnJsonDifferenceFormat != nil {
		toSerialize["returnJsonDifferenceFormat"] = o.ReturnJsonDifferenceFormat
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementUpdateParams struct {
	value *BTAppElementUpdateParams
	isSet bool
}

func (v NullableBTAppElementUpdateParams) Get() *BTAppElementUpdateParams {
	return v.value
}

func (v *NullableBTAppElementUpdateParams) Set(val *BTAppElementUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementUpdateParams(val *BTAppElementUpdateParams) *NullableBTAppElementUpdateParams {
	return &NullableBTAppElementUpdateParams{value: val, isSet: true}
}

func (v NullableBTAppElementUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
