/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6859-c8a2bd7d8338
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Subscription struct for Subscription
type Subscription struct {
	ApplicationFeePercent *float64                    `json:"applicationFeePercent,omitempty"`
	Billing               *string                     `json:"billing,omitempty"`
	CancelAtPeriodEnd     *bool                       `json:"cancelAtPeriodEnd,omitempty"`
	CanceledAt            *int64                      `json:"canceledAt,omitempty"`
	Created               *int64                      `json:"created,omitempty"`
	CurrentPeriodEnd      *int64                      `json:"currentPeriodEnd,omitempty"`
	CurrentPeriodStart    *int64                      `json:"currentPeriodStart,omitempty"`
	Customer              *string                     `json:"customer,omitempty"`
	CustomerObject        *Customer                   `json:"customerObject,omitempty"`
	DaysUntilDue          *int32                      `json:"daysUntilDue,omitempty"`
	Discount              *Discount                   `json:"discount,omitempty"`
	EndedAt               *int64                      `json:"endedAt,omitempty"`
	Id                    *string                     `json:"id,omitempty"`
	Metadata              *map[string]string          `json:"metadata,omitempty"`
	Object                *string                     `json:"object,omitempty"`
	Plan                  *Plan                       `json:"plan,omitempty"`
	Quantity              *int32                      `json:"quantity,omitempty"`
	Start                 *int64                      `json:"start,omitempty"`
	Status                *string                     `json:"status,omitempty"`
	SubscriptionItems     *SubscriptionItemCollection `json:"subscriptionItems,omitempty"`
	TaxPercent            *float64                    `json:"taxPercent,omitempty"`
	TrialEnd              *int64                      `json:"trialEnd,omitempty"`
	TrialStart            *int64                      `json:"trialStart,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription() *Subscription {
	this := Subscription{}
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetApplicationFeePercent returns the ApplicationFeePercent field value if set, zero value otherwise.
func (o *Subscription) GetApplicationFeePercent() float64 {
	if o == nil || o.ApplicationFeePercent == nil {
		var ret float64
		return ret
	}
	return *o.ApplicationFeePercent
}

// GetApplicationFeePercentOk returns a tuple with the ApplicationFeePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetApplicationFeePercentOk() (*float64, bool) {
	if o == nil || o.ApplicationFeePercent == nil {
		return nil, false
	}
	return o.ApplicationFeePercent, true
}

// HasApplicationFeePercent returns a boolean if a field has been set.
func (o *Subscription) HasApplicationFeePercent() bool {
	if o != nil && o.ApplicationFeePercent != nil {
		return true
	}

	return false
}

// SetApplicationFeePercent gets a reference to the given float64 and assigns it to the ApplicationFeePercent field.
func (o *Subscription) SetApplicationFeePercent(v float64) {
	o.ApplicationFeePercent = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *Subscription) GetBilling() string {
	if o == nil || o.Billing == nil {
		var ret string
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetBillingOk() (*string, bool) {
	if o == nil || o.Billing == nil {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *Subscription) HasBilling() bool {
	if o != nil && o.Billing != nil {
		return true
	}

	return false
}

// SetBilling gets a reference to the given string and assigns it to the Billing field.
func (o *Subscription) SetBilling(v string) {
	o.Billing = &v
}

// GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field value if set, zero value otherwise.
func (o *Subscription) GetCancelAtPeriodEnd() bool {
	if o == nil || o.CancelAtPeriodEnd == nil {
		var ret bool
		return ret
	}
	return *o.CancelAtPeriodEnd
}

// GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCancelAtPeriodEndOk() (*bool, bool) {
	if o == nil || o.CancelAtPeriodEnd == nil {
		return nil, false
	}
	return o.CancelAtPeriodEnd, true
}

// HasCancelAtPeriodEnd returns a boolean if a field has been set.
func (o *Subscription) HasCancelAtPeriodEnd() bool {
	if o != nil && o.CancelAtPeriodEnd != nil {
		return true
	}

	return false
}

// SetCancelAtPeriodEnd gets a reference to the given bool and assigns it to the CancelAtPeriodEnd field.
func (o *Subscription) SetCancelAtPeriodEnd(v bool) {
	o.CancelAtPeriodEnd = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *Subscription) GetCanceledAt() int64 {
	if o == nil || o.CanceledAt == nil {
		var ret int64
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCanceledAtOk() (*int64, bool) {
	if o == nil || o.CanceledAt == nil {
		return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *Subscription) HasCanceledAt() bool {
	if o != nil && o.CanceledAt != nil {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given int64 and assigns it to the CanceledAt field.
func (o *Subscription) SetCanceledAt(v int64) {
	o.CanceledAt = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Subscription) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Subscription) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Subscription) SetCreated(v int64) {
	o.Created = &v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value if set, zero value otherwise.
func (o *Subscription) GetCurrentPeriodEnd() int64 {
	if o == nil || o.CurrentPeriodEnd == nil {
		var ret int64
		return ret
	}
	return *o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPeriodEndOk() (*int64, bool) {
	if o == nil || o.CurrentPeriodEnd == nil {
		return nil, false
	}
	return o.CurrentPeriodEnd, true
}

// HasCurrentPeriodEnd returns a boolean if a field has been set.
func (o *Subscription) HasCurrentPeriodEnd() bool {
	if o != nil && o.CurrentPeriodEnd != nil {
		return true
	}

	return false
}

// SetCurrentPeriodEnd gets a reference to the given int64 and assigns it to the CurrentPeriodEnd field.
func (o *Subscription) SetCurrentPeriodEnd(v int64) {
	o.CurrentPeriodEnd = &v
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value if set, zero value otherwise.
func (o *Subscription) GetCurrentPeriodStart() int64 {
	if o == nil || o.CurrentPeriodStart == nil {
		var ret int64
		return ret
	}
	return *o.CurrentPeriodStart
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPeriodStartOk() (*int64, bool) {
	if o == nil || o.CurrentPeriodStart == nil {
		return nil, false
	}
	return o.CurrentPeriodStart, true
}

// HasCurrentPeriodStart returns a boolean if a field has been set.
func (o *Subscription) HasCurrentPeriodStart() bool {
	if o != nil && o.CurrentPeriodStart != nil {
		return true
	}

	return false
}

// SetCurrentPeriodStart gets a reference to the given int64 and assigns it to the CurrentPeriodStart field.
func (o *Subscription) SetCurrentPeriodStart(v int64) {
	o.CurrentPeriodStart = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Subscription) GetCustomer() string {
	if o == nil || o.Customer == nil {
		var ret string
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCustomerOk() (*string, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Subscription) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given string and assigns it to the Customer field.
func (o *Subscription) SetCustomer(v string) {
	o.Customer = &v
}

// GetCustomerObject returns the CustomerObject field value if set, zero value otherwise.
func (o *Subscription) GetCustomerObject() Customer {
	if o == nil || o.CustomerObject == nil {
		var ret Customer
		return ret
	}
	return *o.CustomerObject
}

// GetCustomerObjectOk returns a tuple with the CustomerObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCustomerObjectOk() (*Customer, bool) {
	if o == nil || o.CustomerObject == nil {
		return nil, false
	}
	return o.CustomerObject, true
}

// HasCustomerObject returns a boolean if a field has been set.
func (o *Subscription) HasCustomerObject() bool {
	if o != nil && o.CustomerObject != nil {
		return true
	}

	return false
}

// SetCustomerObject gets a reference to the given Customer and assigns it to the CustomerObject field.
func (o *Subscription) SetCustomerObject(v Customer) {
	o.CustomerObject = &v
}

// GetDaysUntilDue returns the DaysUntilDue field value if set, zero value otherwise.
func (o *Subscription) GetDaysUntilDue() int32 {
	if o == nil || o.DaysUntilDue == nil {
		var ret int32
		return ret
	}
	return *o.DaysUntilDue
}

// GetDaysUntilDueOk returns a tuple with the DaysUntilDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetDaysUntilDueOk() (*int32, bool) {
	if o == nil || o.DaysUntilDue == nil {
		return nil, false
	}
	return o.DaysUntilDue, true
}

// HasDaysUntilDue returns a boolean if a field has been set.
func (o *Subscription) HasDaysUntilDue() bool {
	if o != nil && o.DaysUntilDue != nil {
		return true
	}

	return false
}

// SetDaysUntilDue gets a reference to the given int32 and assigns it to the DaysUntilDue field.
func (o *Subscription) SetDaysUntilDue(v int32) {
	o.DaysUntilDue = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Subscription) GetDiscount() Discount {
	if o == nil || o.Discount == nil {
		var ret Discount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetDiscountOk() (*Discount, bool) {
	if o == nil || o.Discount == nil {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Subscription) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given Discount and assigns it to the Discount field.
func (o *Subscription) SetDiscount(v Discount) {
	o.Discount = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *Subscription) GetEndedAt() int64 {
	if o == nil || o.EndedAt == nil {
		var ret int64
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetEndedAtOk() (*int64, bool) {
	if o == nil || o.EndedAt == nil {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *Subscription) HasEndedAt() bool {
	if o != nil && o.EndedAt != nil {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given int64 and assigns it to the EndedAt field.
func (o *Subscription) SetEndedAt(v int64) {
	o.EndedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subscription) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Subscription) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Subscription) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Subscription) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Subscription) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Subscription) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Subscription) SetObject(v string) {
	o.Object = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *Subscription) GetPlan() Plan {
	if o == nil || o.Plan == nil {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPlanOk() (*Plan, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *Subscription) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *Subscription) SetPlan(v Plan) {
	o.Plan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Subscription) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Subscription) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Subscription) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Subscription) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Subscription) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *Subscription) SetStart(v int64) {
	o.Start = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Subscription) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Subscription) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Subscription) SetStatus(v string) {
	o.Status = &v
}

// GetSubscriptionItems returns the SubscriptionItems field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionItems() SubscriptionItemCollection {
	if o == nil || o.SubscriptionItems == nil {
		var ret SubscriptionItemCollection
		return ret
	}
	return *o.SubscriptionItems
}

// GetSubscriptionItemsOk returns a tuple with the SubscriptionItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionItemsOk() (*SubscriptionItemCollection, bool) {
	if o == nil || o.SubscriptionItems == nil {
		return nil, false
	}
	return o.SubscriptionItems, true
}

// HasSubscriptionItems returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionItems() bool {
	if o != nil && o.SubscriptionItems != nil {
		return true
	}

	return false
}

// SetSubscriptionItems gets a reference to the given SubscriptionItemCollection and assigns it to the SubscriptionItems field.
func (o *Subscription) SetSubscriptionItems(v SubscriptionItemCollection) {
	o.SubscriptionItems = &v
}

// GetTaxPercent returns the TaxPercent field value if set, zero value otherwise.
func (o *Subscription) GetTaxPercent() float64 {
	if o == nil || o.TaxPercent == nil {
		var ret float64
		return ret
	}
	return *o.TaxPercent
}

// GetTaxPercentOk returns a tuple with the TaxPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTaxPercentOk() (*float64, bool) {
	if o == nil || o.TaxPercent == nil {
		return nil, false
	}
	return o.TaxPercent, true
}

// HasTaxPercent returns a boolean if a field has been set.
func (o *Subscription) HasTaxPercent() bool {
	if o != nil && o.TaxPercent != nil {
		return true
	}

	return false
}

// SetTaxPercent gets a reference to the given float64 and assigns it to the TaxPercent field.
func (o *Subscription) SetTaxPercent(v float64) {
	o.TaxPercent = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *Subscription) GetTrialEnd() int64 {
	if o == nil || o.TrialEnd == nil {
		var ret int64
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTrialEndOk() (*int64, bool) {
	if o == nil || o.TrialEnd == nil {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *Subscription) HasTrialEnd() bool {
	if o != nil && o.TrialEnd != nil {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int64 and assigns it to the TrialEnd field.
func (o *Subscription) SetTrialEnd(v int64) {
	o.TrialEnd = &v
}

// GetTrialStart returns the TrialStart field value if set, zero value otherwise.
func (o *Subscription) GetTrialStart() int64 {
	if o == nil || o.TrialStart == nil {
		var ret int64
		return ret
	}
	return *o.TrialStart
}

// GetTrialStartOk returns a tuple with the TrialStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTrialStartOk() (*int64, bool) {
	if o == nil || o.TrialStart == nil {
		return nil, false
	}
	return o.TrialStart, true
}

// HasTrialStart returns a boolean if a field has been set.
func (o *Subscription) HasTrialStart() bool {
	if o != nil && o.TrialStart != nil {
		return true
	}

	return false
}

// SetTrialStart gets a reference to the given int64 and assigns it to the TrialStart field.
func (o *Subscription) SetTrialStart(v int64) {
	o.TrialStart = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationFeePercent != nil {
		toSerialize["applicationFeePercent"] = o.ApplicationFeePercent
	}
	if o.Billing != nil {
		toSerialize["billing"] = o.Billing
	}
	if o.CancelAtPeriodEnd != nil {
		toSerialize["cancelAtPeriodEnd"] = o.CancelAtPeriodEnd
	}
	if o.CanceledAt != nil {
		toSerialize["canceledAt"] = o.CanceledAt
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CurrentPeriodEnd != nil {
		toSerialize["currentPeriodEnd"] = o.CurrentPeriodEnd
	}
	if o.CurrentPeriodStart != nil {
		toSerialize["currentPeriodStart"] = o.CurrentPeriodStart
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.CustomerObject != nil {
		toSerialize["customerObject"] = o.CustomerObject
	}
	if o.DaysUntilDue != nil {
		toSerialize["daysUntilDue"] = o.DaysUntilDue
	}
	if o.Discount != nil {
		toSerialize["discount"] = o.Discount
	}
	if o.EndedAt != nil {
		toSerialize["endedAt"] = o.EndedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SubscriptionItems != nil {
		toSerialize["subscriptionItems"] = o.SubscriptionItems
	}
	if o.TaxPercent != nil {
		toSerialize["taxPercent"] = o.TaxPercent
	}
	if o.TrialEnd != nil {
		toSerialize["trialEnd"] = o.TrialEnd
	}
	if o.TrialStart != nil {
		toSerialize["trialStart"] = o.TrialStart
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
