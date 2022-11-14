/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7328-9aa102855752
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMIndividualCoEdgeQuery1332 struct for BTMIndividualCoEdgeQuery1332
type BTMIndividualCoEdgeQuery1332 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
	PersistentQuery     *BTPStatement269           `json:"persistentQuery,omitempty"`
	QueryStatement      *BTPStatement269           `json:"queryStatement,omitempty"`
	VariableName        *BTMIndividualQuery138     `json:"variableName,omitempty"`
	EdgeQuery           *BTMIndividualQuery138     `json:"edgeQuery,omitempty"`
	FaceQuery           *BTMIndividualQuery138     `json:"faceQuery,omitempty"`
}

// NewBTMIndividualCoEdgeQuery1332 instantiates a new BTMIndividualCoEdgeQuery1332 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualCoEdgeQuery1332() *BTMIndividualCoEdgeQuery1332 {
	this := BTMIndividualCoEdgeQuery1332{}
	return &this
}

// NewBTMIndividualCoEdgeQuery1332WithDefaults instantiates a new BTMIndividualCoEdgeQuery1332 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualCoEdgeQuery1332WithDefaults() *BTMIndividualCoEdgeQuery1332 {
	this := BTMIndividualCoEdgeQuery1332{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualCoEdgeQuery1332) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualCoEdgeQuery1332) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualCoEdgeQuery1332) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualCoEdgeQuery1332) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMIndividualCoEdgeQuery1332) SetNodeId(v string) {
	o.NodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualCoEdgeQuery1332) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualCoEdgeQuery1332) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPersistentQuery returns the PersistentQuery field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetPersistentQuery() BTPStatement269 {
	if o == nil || o.PersistentQuery == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.PersistentQuery
}

// GetPersistentQueryOk returns a tuple with the PersistentQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetPersistentQueryOk() (*BTPStatement269, bool) {
	if o == nil || o.PersistentQuery == nil {
		return nil, false
	}
	return o.PersistentQuery, true
}

// HasPersistentQuery returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasPersistentQuery() bool {
	if o != nil && o.PersistentQuery != nil {
		return true
	}

	return false
}

// SetPersistentQuery gets a reference to the given BTPStatement269 and assigns it to the PersistentQuery field.
func (o *BTMIndividualCoEdgeQuery1332) SetPersistentQuery(v BTPStatement269) {
	o.PersistentQuery = &v
}

// GetQueryStatement returns the QueryStatement field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetQueryStatement() BTPStatement269 {
	if o == nil || o.QueryStatement == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.QueryStatement
}

// GetQueryStatementOk returns a tuple with the QueryStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetQueryStatementOk() (*BTPStatement269, bool) {
	if o == nil || o.QueryStatement == nil {
		return nil, false
	}
	return o.QueryStatement, true
}

// HasQueryStatement returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasQueryStatement() bool {
	if o != nil && o.QueryStatement != nil {
		return true
	}

	return false
}

// SetQueryStatement gets a reference to the given BTPStatement269 and assigns it to the QueryStatement field.
func (o *BTMIndividualCoEdgeQuery1332) SetQueryStatement(v BTPStatement269) {
	o.QueryStatement = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetVariableName() BTMIndividualQuery138 {
	if o == nil || o.VariableName == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetVariableNameOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given BTMIndividualQuery138 and assigns it to the VariableName field.
func (o *BTMIndividualCoEdgeQuery1332) SetVariableName(v BTMIndividualQuery138) {
	o.VariableName = &v
}

// GetEdgeQuery returns the EdgeQuery field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetEdgeQuery() BTMIndividualQuery138 {
	if o == nil || o.EdgeQuery == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.EdgeQuery
}

// GetEdgeQueryOk returns a tuple with the EdgeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetEdgeQueryOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.EdgeQuery == nil {
		return nil, false
	}
	return o.EdgeQuery, true
}

// HasEdgeQuery returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasEdgeQuery() bool {
	if o != nil && o.EdgeQuery != nil {
		return true
	}

	return false
}

// SetEdgeQuery gets a reference to the given BTMIndividualQuery138 and assigns it to the EdgeQuery field.
func (o *BTMIndividualCoEdgeQuery1332) SetEdgeQuery(v BTMIndividualQuery138) {
	o.EdgeQuery = &v
}

// GetFaceQuery returns the FaceQuery field value if set, zero value otherwise.
func (o *BTMIndividualCoEdgeQuery1332) GetFaceQuery() BTMIndividualQuery138 {
	if o == nil || o.FaceQuery == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.FaceQuery
}

// GetFaceQueryOk returns a tuple with the FaceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualCoEdgeQuery1332) GetFaceQueryOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.FaceQuery == nil {
		return nil, false
	}
	return o.FaceQuery, true
}

// HasFaceQuery returns a boolean if a field has been set.
func (o *BTMIndividualCoEdgeQuery1332) HasFaceQuery() bool {
	if o != nil && o.FaceQuery != nil {
		return true
	}

	return false
}

// SetFaceQuery gets a reference to the given BTMIndividualQuery138 and assigns it to the FaceQuery field.
func (o *BTMIndividualCoEdgeQuery1332) SetFaceQuery(v BTMIndividualQuery138) {
	o.FaceQuery = &v
}

func (o BTMIndividualCoEdgeQuery1332) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.PersistentQuery != nil {
		toSerialize["persistentQuery"] = o.PersistentQuery
	}
	if o.QueryStatement != nil {
		toSerialize["queryStatement"] = o.QueryStatement
	}
	if o.VariableName != nil {
		toSerialize["variableName"] = o.VariableName
	}
	if o.EdgeQuery != nil {
		toSerialize["edgeQuery"] = o.EdgeQuery
	}
	if o.FaceQuery != nil {
		toSerialize["faceQuery"] = o.FaceQuery
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualCoEdgeQuery1332 struct {
	value *BTMIndividualCoEdgeQuery1332
	isSet bool
}

func (v NullableBTMIndividualCoEdgeQuery1332) Get() *BTMIndividualCoEdgeQuery1332 {
	return v.value
}

func (v *NullableBTMIndividualCoEdgeQuery1332) Set(val *BTMIndividualCoEdgeQuery1332) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualCoEdgeQuery1332) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualCoEdgeQuery1332) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualCoEdgeQuery1332(val *BTMIndividualCoEdgeQuery1332) *NullableBTMIndividualCoEdgeQuery1332 {
	return &NullableBTMIndividualCoEdgeQuery1332{value: val, isSet: true}
}

func (v NullableBTMIndividualCoEdgeQuery1332) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualCoEdgeQuery1332) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
