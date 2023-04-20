/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14806-89d807e7089c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportModelBody1272 struct for BTExportModelBody1272
type BTExportModelBody1272 struct {
	// If type == COMPOSITE, indicates whether it is open or closed.
	Closed             *bool    `json:"closed,omitempty"`
	ConstituentBodyIds []string `json:"constituentBodyIds,omitempty"`
	// Indicates if there is a closed composite that consumes this body.
	ConsumedByComposite *bool                       `json:"consumedByComposite,omitempty"`
	Edges               []BTExportModelEdge1782     `json:"edges,omitempty"`
	Faces               []BTExportModelFace1363     `json:"faces,omitempty"`
	Id                  *string                     `json:"id,omitempty"`
	Properties          *BTExportBodyProperties3559 `json:"properties,omitempty"`
	Type                *string                     `json:"type,omitempty"`
	Vertices            []BTExportModelVertex858    `json:"vertices,omitempty"`
}

// NewBTExportModelBody1272 instantiates a new BTExportModelBody1272 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelBody1272() *BTExportModelBody1272 {
	this := BTExportModelBody1272{}
	return &this
}

// NewBTExportModelBody1272WithDefaults instantiates a new BTExportModelBody1272 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelBody1272WithDefaults() *BTExportModelBody1272 {
	this := BTExportModelBody1272{}
	return &this
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetClosed() bool {
	if o == nil || o.Closed == nil {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetClosedOk() (*bool, bool) {
	if o == nil || o.Closed == nil {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasClosed() bool {
	if o != nil && o.Closed != nil {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *BTExportModelBody1272) SetClosed(v bool) {
	o.Closed = &v
}

// GetConstituentBodyIds returns the ConstituentBodyIds field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetConstituentBodyIds() []string {
	if o == nil || o.ConstituentBodyIds == nil {
		var ret []string
		return ret
	}
	return o.ConstituentBodyIds
}

// GetConstituentBodyIdsOk returns a tuple with the ConstituentBodyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetConstituentBodyIdsOk() ([]string, bool) {
	if o == nil || o.ConstituentBodyIds == nil {
		return nil, false
	}
	return o.ConstituentBodyIds, true
}

// HasConstituentBodyIds returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasConstituentBodyIds() bool {
	if o != nil && o.ConstituentBodyIds != nil {
		return true
	}

	return false
}

// SetConstituentBodyIds gets a reference to the given []string and assigns it to the ConstituentBodyIds field.
func (o *BTExportModelBody1272) SetConstituentBodyIds(v []string) {
	o.ConstituentBodyIds = v
}

// GetConsumedByComposite returns the ConsumedByComposite field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetConsumedByComposite() bool {
	if o == nil || o.ConsumedByComposite == nil {
		var ret bool
		return ret
	}
	return *o.ConsumedByComposite
}

// GetConsumedByCompositeOk returns a tuple with the ConsumedByComposite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetConsumedByCompositeOk() (*bool, bool) {
	if o == nil || o.ConsumedByComposite == nil {
		return nil, false
	}
	return o.ConsumedByComposite, true
}

// HasConsumedByComposite returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasConsumedByComposite() bool {
	if o != nil && o.ConsumedByComposite != nil {
		return true
	}

	return false
}

// SetConsumedByComposite gets a reference to the given bool and assigns it to the ConsumedByComposite field.
func (o *BTExportModelBody1272) SetConsumedByComposite(v bool) {
	o.ConsumedByComposite = &v
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetEdges() []BTExportModelEdge1782 {
	if o == nil || o.Edges == nil {
		var ret []BTExportModelEdge1782
		return ret
	}
	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetEdgesOk() ([]BTExportModelEdge1782, bool) {
	if o == nil || o.Edges == nil {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasEdges() bool {
	if o != nil && o.Edges != nil {
		return true
	}

	return false
}

// SetEdges gets a reference to the given []BTExportModelEdge1782 and assigns it to the Edges field.
func (o *BTExportModelBody1272) SetEdges(v []BTExportModelEdge1782) {
	o.Edges = v
}

// GetFaces returns the Faces field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetFaces() []BTExportModelFace1363 {
	if o == nil || o.Faces == nil {
		var ret []BTExportModelFace1363
		return ret
	}
	return o.Faces
}

// GetFacesOk returns a tuple with the Faces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetFacesOk() ([]BTExportModelFace1363, bool) {
	if o == nil || o.Faces == nil {
		return nil, false
	}
	return o.Faces, true
}

// HasFaces returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasFaces() bool {
	if o != nil && o.Faces != nil {
		return true
	}

	return false
}

// SetFaces gets a reference to the given []BTExportModelFace1363 and assigns it to the Faces field.
func (o *BTExportModelBody1272) SetFaces(v []BTExportModelFace1363) {
	o.Faces = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportModelBody1272) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetProperties() BTExportBodyProperties3559 {
	if o == nil || o.Properties == nil {
		var ret BTExportBodyProperties3559
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetPropertiesOk() (*BTExportBodyProperties3559, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given BTExportBodyProperties3559 and assigns it to the Properties field.
func (o *BTExportModelBody1272) SetProperties(v BTExportBodyProperties3559) {
	o.Properties = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTExportModelBody1272) SetType(v string) {
	o.Type = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *BTExportModelBody1272) GetVertices() []BTExportModelVertex858 {
	if o == nil || o.Vertices == nil {
		var ret []BTExportModelVertex858
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBody1272) GetVerticesOk() ([]BTExportModelVertex858, bool) {
	if o == nil || o.Vertices == nil {
		return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *BTExportModelBody1272) HasVertices() bool {
	if o != nil && o.Vertices != nil {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []BTExportModelVertex858 and assigns it to the Vertices field.
func (o *BTExportModelBody1272) SetVertices(v []BTExportModelVertex858) {
	o.Vertices = v
}

func (o BTExportModelBody1272) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Closed != nil {
		toSerialize["closed"] = o.Closed
	}
	if o.ConstituentBodyIds != nil {
		toSerialize["constituentBodyIds"] = o.ConstituentBodyIds
	}
	if o.ConsumedByComposite != nil {
		toSerialize["consumedByComposite"] = o.ConsumedByComposite
	}
	if o.Edges != nil {
		toSerialize["edges"] = o.Edges
	}
	if o.Faces != nil {
		toSerialize["faces"] = o.Faces
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Vertices != nil {
		toSerialize["vertices"] = o.Vertices
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelBody1272 struct {
	value *BTExportModelBody1272
	isSet bool
}

func (v NullableBTExportModelBody1272) Get() *BTExportModelBody1272 {
	return v.value
}

func (v *NullableBTExportModelBody1272) Set(val *BTExportModelBody1272) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelBody1272) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelBody1272) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelBody1272(val *BTExportModelBody1272) *NullableBTExportModelBody1272 {
	return &NullableBTExportModelBody1272{value: val, isSet: true}
}

func (v NullableBTExportModelBody1272) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelBody1272) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
