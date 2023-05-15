/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15808-38acf80dff96
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOnePartProperties230 struct for BTOnePartProperties230
type BTOnePartProperties230 struct {
	Appearance                   *BTGraphicsAppearance1152   `json:"appearance,omitempty"`
	AppearanceForNewCell         *BTGraphicsAppearance1152   `json:"appearanceForNewCell,omitempty"`
	BtType                       *string                     `json:"btType,omitempty"`
	ChangedPropertiesSet         []string                    `json:"changedPropertiesSet,omitempty"`
	CustomProperties             *BTPartCustomProperties1338 `json:"customProperties,omitempty"`
	Material                     *BTPartMaterial1445         `json:"material,omitempty"`
	MaterialForNewCell           *BTPartMaterial1445         `json:"materialForNewCell,omitempty"`
	Name                         *string                     `json:"name,omitempty"`
	NameForNewCell               *string                     `json:"nameForNewCell,omitempty"`
	NameIfNotNull                *BTOnePartProperties230     `json:"nameIfNotNull,omitempty"`
	NodeId                       *string                     `json:"nodeId,omitempty"`
	ParsedQuery                  *BTPFunctionDeclaration246  `json:"parsedQuery,omitempty"`
	Query                        *string                     `json:"query,omitempty"`
	QueryListParameter           *BTMParameterQueryList148   `json:"queryListParameter,omitempty"`
	SheetMetalBendOrder          []string                    `json:"sheetMetalBendOrder,omitempty"`
	SheetMetalBendOrderIfNotNull *BTOnePartProperties230     `json:"sheetMetalBendOrderIfNotNull,omitempty"`
	Visibility                   *GBTPartVisibility          `json:"visibility,omitempty"`
}

// NewBTOnePartProperties230 instantiates a new BTOnePartProperties230 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOnePartProperties230() *BTOnePartProperties230 {
	this := BTOnePartProperties230{}
	return &this
}

// NewBTOnePartProperties230WithDefaults instantiates a new BTOnePartProperties230 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOnePartProperties230WithDefaults() *BTOnePartProperties230 {
	this := BTOnePartProperties230{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTOnePartProperties230) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetAppearanceForNewCell returns the AppearanceForNewCell field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetAppearanceForNewCell() BTGraphicsAppearance1152 {
	if o == nil || o.AppearanceForNewCell == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.AppearanceForNewCell
}

// GetAppearanceForNewCellOk returns a tuple with the AppearanceForNewCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetAppearanceForNewCellOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.AppearanceForNewCell == nil {
		return nil, false
	}
	return o.AppearanceForNewCell, true
}

// HasAppearanceForNewCell returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasAppearanceForNewCell() bool {
	if o != nil && o.AppearanceForNewCell != nil {
		return true
	}

	return false
}

// SetAppearanceForNewCell gets a reference to the given BTGraphicsAppearance1152 and assigns it to the AppearanceForNewCell field.
func (o *BTOnePartProperties230) SetAppearanceForNewCell(v BTGraphicsAppearance1152) {
	o.AppearanceForNewCell = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOnePartProperties230) SetBtType(v string) {
	o.BtType = &v
}

// GetChangedPropertiesSet returns the ChangedPropertiesSet field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetChangedPropertiesSet() []string {
	if o == nil || o.ChangedPropertiesSet == nil {
		var ret []string
		return ret
	}
	return o.ChangedPropertiesSet
}

// GetChangedPropertiesSetOk returns a tuple with the ChangedPropertiesSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetChangedPropertiesSetOk() ([]string, bool) {
	if o == nil || o.ChangedPropertiesSet == nil {
		return nil, false
	}
	return o.ChangedPropertiesSet, true
}

// HasChangedPropertiesSet returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasChangedPropertiesSet() bool {
	if o != nil && o.ChangedPropertiesSet != nil {
		return true
	}

	return false
}

// SetChangedPropertiesSet gets a reference to the given []string and assigns it to the ChangedPropertiesSet field.
func (o *BTOnePartProperties230) SetChangedPropertiesSet(v []string) {
	o.ChangedPropertiesSet = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetCustomProperties() BTPartCustomProperties1338 {
	if o == nil || o.CustomProperties == nil {
		var ret BTPartCustomProperties1338
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetCustomPropertiesOk() (*BTPartCustomProperties1338, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given BTPartCustomProperties1338 and assigns it to the CustomProperties field.
func (o *BTOnePartProperties230) SetCustomProperties(v BTPartCustomProperties1338) {
	o.CustomProperties = &v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetMaterial() BTPartMaterial1445 {
	if o == nil || o.Material == nil {
		var ret BTPartMaterial1445
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetMaterialOk() (*BTPartMaterial1445, bool) {
	if o == nil || o.Material == nil {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasMaterial() bool {
	if o != nil && o.Material != nil {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given BTPartMaterial1445 and assigns it to the Material field.
func (o *BTOnePartProperties230) SetMaterial(v BTPartMaterial1445) {
	o.Material = &v
}

// GetMaterialForNewCell returns the MaterialForNewCell field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetMaterialForNewCell() BTPartMaterial1445 {
	if o == nil || o.MaterialForNewCell == nil {
		var ret BTPartMaterial1445
		return ret
	}
	return *o.MaterialForNewCell
}

// GetMaterialForNewCellOk returns a tuple with the MaterialForNewCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetMaterialForNewCellOk() (*BTPartMaterial1445, bool) {
	if o == nil || o.MaterialForNewCell == nil {
		return nil, false
	}
	return o.MaterialForNewCell, true
}

// HasMaterialForNewCell returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasMaterialForNewCell() bool {
	if o != nil && o.MaterialForNewCell != nil {
		return true
	}

	return false
}

// SetMaterialForNewCell gets a reference to the given BTPartMaterial1445 and assigns it to the MaterialForNewCell field.
func (o *BTOnePartProperties230) SetMaterialForNewCell(v BTPartMaterial1445) {
	o.MaterialForNewCell = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTOnePartProperties230) SetName(v string) {
	o.Name = &v
}

// GetNameForNewCell returns the NameForNewCell field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetNameForNewCell() string {
	if o == nil || o.NameForNewCell == nil {
		var ret string
		return ret
	}
	return *o.NameForNewCell
}

// GetNameForNewCellOk returns a tuple with the NameForNewCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetNameForNewCellOk() (*string, bool) {
	if o == nil || o.NameForNewCell == nil {
		return nil, false
	}
	return o.NameForNewCell, true
}

// HasNameForNewCell returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasNameForNewCell() bool {
	if o != nil && o.NameForNewCell != nil {
		return true
	}

	return false
}

// SetNameForNewCell gets a reference to the given string and assigns it to the NameForNewCell field.
func (o *BTOnePartProperties230) SetNameForNewCell(v string) {
	o.NameForNewCell = &v
}

// GetNameIfNotNull returns the NameIfNotNull field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetNameIfNotNull() BTOnePartProperties230 {
	if o == nil || o.NameIfNotNull == nil {
		var ret BTOnePartProperties230
		return ret
	}
	return *o.NameIfNotNull
}

// GetNameIfNotNullOk returns a tuple with the NameIfNotNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetNameIfNotNullOk() (*BTOnePartProperties230, bool) {
	if o == nil || o.NameIfNotNull == nil {
		return nil, false
	}
	return o.NameIfNotNull, true
}

// HasNameIfNotNull returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasNameIfNotNull() bool {
	if o != nil && o.NameIfNotNull != nil {
		return true
	}

	return false
}

// SetNameIfNotNull gets a reference to the given BTOnePartProperties230 and assigns it to the NameIfNotNull field.
func (o *BTOnePartProperties230) SetNameIfNotNull(v BTOnePartProperties230) {
	o.NameIfNotNull = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTOnePartProperties230) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParsedQuery returns the ParsedQuery field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetParsedQuery() BTPFunctionDeclaration246 {
	if o == nil || o.ParsedQuery == nil {
		var ret BTPFunctionDeclaration246
		return ret
	}
	return *o.ParsedQuery
}

// GetParsedQueryOk returns a tuple with the ParsedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetParsedQueryOk() (*BTPFunctionDeclaration246, bool) {
	if o == nil || o.ParsedQuery == nil {
		return nil, false
	}
	return o.ParsedQuery, true
}

// HasParsedQuery returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasParsedQuery() bool {
	if o != nil && o.ParsedQuery != nil {
		return true
	}

	return false
}

// SetParsedQuery gets a reference to the given BTPFunctionDeclaration246 and assigns it to the ParsedQuery field.
func (o *BTOnePartProperties230) SetParsedQuery(v BTPFunctionDeclaration246) {
	o.ParsedQuery = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *BTOnePartProperties230) SetQuery(v string) {
	o.Query = &v
}

// GetQueryListParameter returns the QueryListParameter field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetQueryListParameter() BTMParameterQueryList148 {
	if o == nil || o.QueryListParameter == nil {
		var ret BTMParameterQueryList148
		return ret
	}
	return *o.QueryListParameter
}

// GetQueryListParameterOk returns a tuple with the QueryListParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetQueryListParameterOk() (*BTMParameterQueryList148, bool) {
	if o == nil || o.QueryListParameter == nil {
		return nil, false
	}
	return o.QueryListParameter, true
}

// HasQueryListParameter returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasQueryListParameter() bool {
	if o != nil && o.QueryListParameter != nil {
		return true
	}

	return false
}

// SetQueryListParameter gets a reference to the given BTMParameterQueryList148 and assigns it to the QueryListParameter field.
func (o *BTOnePartProperties230) SetQueryListParameter(v BTMParameterQueryList148) {
	o.QueryListParameter = &v
}

// GetSheetMetalBendOrder returns the SheetMetalBendOrder field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetSheetMetalBendOrder() []string {
	if o == nil || o.SheetMetalBendOrder == nil {
		var ret []string
		return ret
	}
	return o.SheetMetalBendOrder
}

// GetSheetMetalBendOrderOk returns a tuple with the SheetMetalBendOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetSheetMetalBendOrderOk() ([]string, bool) {
	if o == nil || o.SheetMetalBendOrder == nil {
		return nil, false
	}
	return o.SheetMetalBendOrder, true
}

// HasSheetMetalBendOrder returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasSheetMetalBendOrder() bool {
	if o != nil && o.SheetMetalBendOrder != nil {
		return true
	}

	return false
}

// SetSheetMetalBendOrder gets a reference to the given []string and assigns it to the SheetMetalBendOrder field.
func (o *BTOnePartProperties230) SetSheetMetalBendOrder(v []string) {
	o.SheetMetalBendOrder = v
}

// GetSheetMetalBendOrderIfNotNull returns the SheetMetalBendOrderIfNotNull field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetSheetMetalBendOrderIfNotNull() BTOnePartProperties230 {
	if o == nil || o.SheetMetalBendOrderIfNotNull == nil {
		var ret BTOnePartProperties230
		return ret
	}
	return *o.SheetMetalBendOrderIfNotNull
}

// GetSheetMetalBendOrderIfNotNullOk returns a tuple with the SheetMetalBendOrderIfNotNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetSheetMetalBendOrderIfNotNullOk() (*BTOnePartProperties230, bool) {
	if o == nil || o.SheetMetalBendOrderIfNotNull == nil {
		return nil, false
	}
	return o.SheetMetalBendOrderIfNotNull, true
}

// HasSheetMetalBendOrderIfNotNull returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasSheetMetalBendOrderIfNotNull() bool {
	if o != nil && o.SheetMetalBendOrderIfNotNull != nil {
		return true
	}

	return false
}

// SetSheetMetalBendOrderIfNotNull gets a reference to the given BTOnePartProperties230 and assigns it to the SheetMetalBendOrderIfNotNull field.
func (o *BTOnePartProperties230) SetSheetMetalBendOrderIfNotNull(v BTOnePartProperties230) {
	o.SheetMetalBendOrderIfNotNull = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTOnePartProperties230) GetVisibility() GBTPartVisibility {
	if o == nil || o.Visibility == nil {
		var ret GBTPartVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOnePartProperties230) GetVisibilityOk() (*GBTPartVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTOnePartProperties230) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given GBTPartVisibility and assigns it to the Visibility field.
func (o *BTOnePartProperties230) SetVisibility(v GBTPartVisibility) {
	o.Visibility = &v
}

func (o BTOnePartProperties230) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.AppearanceForNewCell != nil {
		toSerialize["appearanceForNewCell"] = o.AppearanceForNewCell
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ChangedPropertiesSet != nil {
		toSerialize["changedPropertiesSet"] = o.ChangedPropertiesSet
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if o.Material != nil {
		toSerialize["material"] = o.Material
	}
	if o.MaterialForNewCell != nil {
		toSerialize["materialForNewCell"] = o.MaterialForNewCell
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameForNewCell != nil {
		toSerialize["nameForNewCell"] = o.NameForNewCell
	}
	if o.NameIfNotNull != nil {
		toSerialize["nameIfNotNull"] = o.NameIfNotNull
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParsedQuery != nil {
		toSerialize["parsedQuery"] = o.ParsedQuery
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryListParameter != nil {
		toSerialize["queryListParameter"] = o.QueryListParameter
	}
	if o.SheetMetalBendOrder != nil {
		toSerialize["sheetMetalBendOrder"] = o.SheetMetalBendOrder
	}
	if o.SheetMetalBendOrderIfNotNull != nil {
		toSerialize["sheetMetalBendOrderIfNotNull"] = o.SheetMetalBendOrderIfNotNull
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableBTOnePartProperties230 struct {
	value *BTOnePartProperties230
	isSet bool
}

func (v NullableBTOnePartProperties230) Get() *BTOnePartProperties230 {
	return v.value
}

func (v *NullableBTOnePartProperties230) Set(val *BTOnePartProperties230) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOnePartProperties230) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOnePartProperties230) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOnePartProperties230(val *BTOnePartProperties230) *NullableBTOnePartProperties230 {
	return &NullableBTOnePartProperties230{value: val, isSet: true}
}

func (v NullableBTOnePartProperties230) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOnePartProperties230) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
