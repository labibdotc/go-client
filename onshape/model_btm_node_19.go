/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17666-197c9d1638c5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMNode19 - struct for BTMNode19
type BTMNode19 struct {
	implBTMNode19 interface{}
}

// BTMUnitsDefault160AsBTMNode19 is a convenience function that returns BTMUnitsDefault160 wrapped in BTMNode19
func (o *BTMUnitsDefault160) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTInstanceControlNode750AsBTMNode19 is a convenience function that returns BTInstanceControlNode750 wrapped in BTMNode19
func (o *BTInstanceControlNode750) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTAssemblyPattern1974AsBTMNode19 is a convenience function that returns BTAssemblyPattern1974 wrapped in BTMNode19
func (o *BTAssemblyPattern1974) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMFeature134AsBTMNode19 is a convenience function that returns BTMFeature134 wrapped in BTMNode19
func (o *BTMFeature134) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMIndividualQueryBase139AsBTMNode19 is a convenience function that returns BTMIndividualQueryBase139 wrapped in BTMNode19
func (o *BTMIndividualQueryBase139) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMParameter1AsBTMNode19 is a convenience function that returns BTMParameter1 wrapped in BTMNode19
func (o *BTMParameter1) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMArrayParameterItem1843AsBTMNode19 is a convenience function that returns BTMArrayParameterItem1843 wrapped in BTMNode19
func (o *BTMArrayParameterItem1843) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMEnumOption592AsBTMNode19 is a convenience function that returns BTMEnumOption592 wrapped in BTMNode19
func (o *BTMEnumOption592) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMNodeInvalid1772AsBTMNode19 is a convenience function that returns BTMNodeInvalid1772 wrapped in BTMNode19
func (o *BTMNodeInvalid1772) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMConfigurationData1560AsBTMNode19 is a convenience function that returns BTMConfigurationData1560 wrapped in BTMNode19
func (o *BTMConfigurationData1560) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMRecordMetrics1169AsBTMNode19 is a convenience function that returns BTMRecordMetrics1169 wrapped in BTMNode19
func (o *BTMRecordMetrics1169) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTInstanceBase2263AsBTMNode19 is a convenience function that returns BTInstanceBase2263 wrapped in BTMNode19
func (o *BTInstanceBase2263) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMModel141AsBTMNode19 is a convenience function that returns BTMModel141 wrapped in BTMNode19
func (o *BTMModel141) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMSuppressionStateConfigured2598AsBTMNode19 is a convenience function that returns BTMSuppressionStateConfigured2598 wrapped in BTMNode19
func (o *BTMSuppressionStateConfigured2598) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMConfiguredValue1341AsBTMNode19 is a convenience function that returns BTMConfiguredValue1341 wrapped in BTMNode19
func (o *BTMConfiguredValue1341) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTParametricInstance2641AsBTMNode19 is a convenience function that returns BTParametricInstance2641 wrapped in BTMNode19
func (o *BTParametricInstance2641) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTAssemblySimulationData978AsBTMNode19 is a convenience function that returns BTAssemblySimulationData978 wrapped in BTMNode19
func (o *BTAssemblySimulationData978) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTSimulationInstance3093AsBTMNode19 is a convenience function that returns BTSimulationInstance3093 wrapped in BTMNode19
func (o *BTSimulationInstance3093) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTClonedInstance2505AsBTMNode19 is a convenience function that returns BTClonedInstance2505 wrapped in BTMNode19
func (o *BTClonedInstance2505) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMImport136AsBTMNode19 is a convenience function that returns BTMImport136 wrapped in BTMNode19
func (o *BTMImport136) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTInstanceFolder3627AsBTMNode19 is a convenience function that returns BTInstanceFolder3627 wrapped in BTMNode19
func (o *BTInstanceFolder3627) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTToleranceString3274AsBTMNode19 is a convenience function that returns BTToleranceString3274 wrapped in BTMNode19
func (o *BTToleranceString3274) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMConfigurationParameter819AsBTMNode19 is a convenience function that returns BTMConfigurationParameter819 wrapped in BTMNode19
func (o *BTMConfigurationParameter819) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMSketchInvalid1601AsBTMNode19 is a convenience function that returns BTMSketchInvalid1601 wrapped in BTMNode19
func (o *BTMSketchInvalid1601) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMSuppressionState1924AsBTMNode19 is a convenience function that returns BTMSuppressionState1924 wrapped in BTMNode19
func (o *BTMSuppressionState1924) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTAssemblyInstance947AsBTMNode19 is a convenience function that returns BTAssemblyInstance947 wrapped in BTMNode19
func (o *BTAssemblyInstance947) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMUserCode161AsBTMNode19 is a convenience function that returns BTMUserCode161 wrapped in BTMNode19
func (o *BTMUserCode161) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMFolder3208AsBTMNode19 is a convenience function that returns BTMFolder3208 wrapped in BTMNode19
func (o *BTMFolder3208) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMRollback150AsBTMNode19 is a convenience function that returns BTMRollback150 wrapped in BTMNode19
func (o *BTMRollback150) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTPartInstance81AsBTMNode19 is a convenience function that returns BTPartInstance81 wrapped in BTMNode19
func (o *BTPartInstance81) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTAssemblyReplicate2774AsBTMNode19 is a convenience function that returns BTAssemblyReplicate2774 wrapped in BTMNode19
func (o *BTAssemblyReplicate2774) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTInstance642AsBTMNode19 is a convenience function that returns BTInstance642 wrapped in BTMNode19
func (o *BTInstance642) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMSketchCompositeEntity893AsBTMNode19 is a convenience function that returns BTMSketchCompositeEntity893 wrapped in BTMNode19
func (o *BTMSketchCompositeEntity893) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// BTMSketchEntity3AsBTMNode19 is a convenience function that returns BTMSketchEntity3 wrapped in BTMNode19
func (o *BTMSketchEntity3) AsBTMNode19() *BTMNode19 {
	return &BTMNode19{o}
}

// NewBTMNode19 instantiates a new BTMNode19 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMNode19() *BTMNode19 {
	this := BTMNode19{Newbase_BTMNode19()}
	return &this
}

// NewBTMNode19WithDefaults instantiates a new BTMNode19 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMNode19WithDefaults() *BTMNode19 {
	this := BTMNode19{Newbase_BTMNode19WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMNode19) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNode19) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMNode19) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMNode19) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMNode19) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNode19) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMNode19) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMNode19) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMNode19) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNode19) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMNode19) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMNode19) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMNode19) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTAssemblyInstance-947'
	if jsonDict["btType"] == "BTAssemblyInstance-947" {
		// try to unmarshal JSON data into BTAssemblyInstance947
		var qr *BTAssemblyInstance947
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTAssemblyInstance947: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAssemblyPattern-1974'
	if jsonDict["btType"] == "BTAssemblyPattern-1974" {
		// try to unmarshal JSON data into BTAssemblyPattern1974
		var qr *BTAssemblyPattern1974
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTAssemblyPattern1974: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAssemblyReplicate-2774'
	if jsonDict["btType"] == "BTAssemblyReplicate-2774" {
		// try to unmarshal JSON data into BTAssemblyReplicate2774
		var qr *BTAssemblyReplicate2774
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTAssemblyReplicate2774: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAssemblySimulationData-978'
	if jsonDict["btType"] == "BTAssemblySimulationData-978" {
		// try to unmarshal JSON data into BTAssemblySimulationData978
		var qr *BTAssemblySimulationData978
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTAssemblySimulationData978: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTClonedInstance-2505'
	if jsonDict["btType"] == "BTClonedInstance-2505" {
		// try to unmarshal JSON data into BTClonedInstance2505
		var qr *BTClonedInstance2505
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTClonedInstance2505: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTInstance-642'
	if jsonDict["btType"] == "BTInstance-642" {
		// try to unmarshal JSON data into BTInstance642
		var qr *BTInstance642
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTInstance642: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTInstanceBase-2263'
	if jsonDict["btType"] == "BTInstanceBase-2263" {
		// try to unmarshal JSON data into BTInstanceBase2263
		var qr *BTInstanceBase2263
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTInstanceBase2263: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTInstanceControlNode-750'
	if jsonDict["btType"] == "BTInstanceControlNode-750" {
		// try to unmarshal JSON data into BTInstanceControlNode750
		var qr *BTInstanceControlNode750
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTInstanceControlNode750: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTInstanceFolder-3627'
	if jsonDict["btType"] == "BTInstanceFolder-3627" {
		// try to unmarshal JSON data into BTInstanceFolder3627
		var qr *BTInstanceFolder3627
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTInstanceFolder3627: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMArrayParameterItem-1843'
	if jsonDict["btType"] == "BTMArrayParameterItem-1843" {
		// try to unmarshal JSON data into BTMArrayParameterItem1843
		var qr *BTMArrayParameterItem1843
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMArrayParameterItem1843: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfigurationData-1560'
	if jsonDict["btType"] == "BTMConfigurationData-1560" {
		// try to unmarshal JSON data into BTMConfigurationData1560
		var qr *BTMConfigurationData1560
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMConfigurationData1560: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfigurationParameter-819'
	if jsonDict["btType"] == "BTMConfigurationParameter-819" {
		// try to unmarshal JSON data into BTMConfigurationParameter819
		var qr *BTMConfigurationParameter819
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMConfigurationParameter819: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfiguredValue-1341'
	if jsonDict["btType"] == "BTMConfiguredValue-1341" {
		// try to unmarshal JSON data into BTMConfiguredValue1341
		var qr *BTMConfiguredValue1341
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMConfiguredValue1341: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMEnumOption-592'
	if jsonDict["btType"] == "BTMEnumOption-592" {
		// try to unmarshal JSON data into BTMEnumOption592
		var qr *BTMEnumOption592
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMEnumOption592: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMFeature-134'
	if jsonDict["btType"] == "BTMFeature-134" {
		// try to unmarshal JSON data into BTMFeature134
		var qr *BTMFeature134
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMFeature134: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMFolder-3208'
	if jsonDict["btType"] == "BTMFolder-3208" {
		// try to unmarshal JSON data into BTMFolder3208
		var qr *BTMFolder3208
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMFolder3208: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMImport-136'
	if jsonDict["btType"] == "BTMImport-136" {
		// try to unmarshal JSON data into BTMImport136
		var qr *BTMImport136
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMImport136: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualQueryBase-139'
	if jsonDict["btType"] == "BTMIndividualQueryBase-139" {
		// try to unmarshal JSON data into BTMIndividualQueryBase139
		var qr *BTMIndividualQueryBase139
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMIndividualQueryBase139: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMModel-141'
	if jsonDict["btType"] == "BTMModel-141" {
		// try to unmarshal JSON data into BTMModel141
		var qr *BTMModel141
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMModel141: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMNodeInvalid-1772'
	if jsonDict["btType"] == "BTMNodeInvalid-1772" {
		// try to unmarshal JSON data into BTMNodeInvalid1772
		var qr *BTMNodeInvalid1772
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMNodeInvalid1772: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameter-1'
	if jsonDict["btType"] == "BTMParameter-1" {
		// try to unmarshal JSON data into BTMParameter1
		var qr *BTMParameter1
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMParameter1: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMRecordMetrics-1169'
	if jsonDict["btType"] == "BTMRecordMetrics-1169" {
		// try to unmarshal JSON data into BTMRecordMetrics1169
		var qr *BTMRecordMetrics1169
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMRecordMetrics1169: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMRollback-150'
	if jsonDict["btType"] == "BTMRollback-150" {
		// try to unmarshal JSON data into BTMRollback150
		var qr *BTMRollback150
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMRollback150: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSketchCompositeEntity-893'
	if jsonDict["btType"] == "BTMSketchCompositeEntity-893" {
		// try to unmarshal JSON data into BTMSketchCompositeEntity893
		var qr *BTMSketchCompositeEntity893
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMSketchCompositeEntity893: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSketchEntity-3'
	if jsonDict["btType"] == "BTMSketchEntity-3" {
		// try to unmarshal JSON data into BTMSketchEntity3
		var qr *BTMSketchEntity3
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMSketchEntity3: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSketchInvalid-1601'
	if jsonDict["btType"] == "BTMSketchInvalid-1601" {
		// try to unmarshal JSON data into BTMSketchInvalid1601
		var qr *BTMSketchInvalid1601
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMSketchInvalid1601: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSuppressionState-1924'
	if jsonDict["btType"] == "BTMSuppressionState-1924" {
		// try to unmarshal JSON data into BTMSuppressionState1924
		var qr *BTMSuppressionState1924
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMSuppressionState1924: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSuppressionStateConfigured-2598'
	if jsonDict["btType"] == "BTMSuppressionStateConfigured-2598" {
		// try to unmarshal JSON data into BTMSuppressionStateConfigured2598
		var qr *BTMSuppressionStateConfigured2598
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMSuppressionStateConfigured2598: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMUnitsDefault-160'
	if jsonDict["btType"] == "BTMUnitsDefault-160" {
		// try to unmarshal JSON data into BTMUnitsDefault160
		var qr *BTMUnitsDefault160
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMUnitsDefault160: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMUserCode-161'
	if jsonDict["btType"] == "BTMUserCode-161" {
		// try to unmarshal JSON data into BTMUserCode161
		var qr *BTMUserCode161
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTMUserCode161: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTParametricInstance-2641'
	if jsonDict["btType"] == "BTParametricInstance-2641" {
		// try to unmarshal JSON data into BTParametricInstance2641
		var qr *BTParametricInstance2641
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTParametricInstance2641: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPartInstance-81'
	if jsonDict["btType"] == "BTPartInstance-81" {
		// try to unmarshal JSON data into BTPartInstance81
		var qr *BTPartInstance81
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTPartInstance81: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSimulationInstance-3093'
	if jsonDict["btType"] == "BTSimulationInstance-3093" {
		// try to unmarshal JSON data into BTSimulationInstance3093
		var qr *BTSimulationInstance3093
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTSimulationInstance3093: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTToleranceString-3274'
	if jsonDict["btType"] == "BTToleranceString-3274" {
		// try to unmarshal JSON data into BTToleranceString3274
		var qr *BTToleranceString3274
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMNode19 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMNode19 = nil
			return fmt.Errorf("Failed to unmarshal BTMNode19 as BTToleranceString3274: %s", err.Error())
		}
	}

	var qtx *base_BTMNode19
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMNode19 = qtx
		return nil // data stored in dst.base_BTMNode19, return on the first match
	} else {
		dst.implBTMNode19 = nil
		return fmt.Errorf("Failed to unmarshal BTMNode19 as base_BTMNode19: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMNode19) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMNode19) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMNode19
}

type NullableBTMNode19 struct {
	value *BTMNode19
	isSet bool
}

func (v NullableBTMNode19) Get() *BTMNode19 {
	return v.value
}

func (v *NullableBTMNode19) Set(val *BTMNode19) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMNode19) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMNode19) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMNode19(val *BTMNode19) *NullableBTMNode19 {
	return &NullableBTMNode19{value: val, isSet: true}
}

func (v NullableBTMNode19) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMNode19) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMNode19 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
}

// Newbase_BTMNode19 instantiates a new base_BTMNode19 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMNode19() *base_BTMNode19 {
	this := base_BTMNode19{}
	return &this
}

// Newbase_BTMNode19WithDefaults instantiates a new base_BTMNode19 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMNode19WithDefaults() *base_BTMNode19 {
	this := base_BTMNode19{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMNode19) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMNode19) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMNode19) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMNode19) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMNode19) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMNode19) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMNode19) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMNode19) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMNode19) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMNode19) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMNode19) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMNode19) SetNodeId(v string) {
	o.NodeId = &v
}

func (o base_BTMNode19) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}
