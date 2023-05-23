/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16301-d273853a12e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTActiveWorkflowInfo struct for BTActiveWorkflowInfo
type BTActiveWorkflowInfo struct {
	AllowReleaseItemsFromOtherDocuments *bool                     `json:"allowReleaseItemsFromOtherDocuments,omitempty"`
	CanCurrentUserCreateReleases        *bool                     `json:"canCurrentUserCreateReleases,omitempty"`
	CanCurrentUserManageWorkflows       *bool                     `json:"canCurrentUserManageWorkflows,omitempty"`
	CanCurrentUserSeeArenaItemLink      *bool                     `json:"canCurrentUserSeeArenaItemLink,omitempty"`
	CanCurrentUserSyncBomToArena        *bool                     `json:"canCurrentUserSyncBomToArena,omitempty"`
	CanCurrentUserSyncToArena           *bool                     `json:"canCurrentUserSyncToArena,omitempty"`
	CanCurrentUserSyncVersionsToArena   *bool                     `json:"canCurrentUserSyncVersionsToArena,omitempty"`
	CompanyId                           *string                   `json:"companyId,omitempty"`
	DocumentId                          *string                   `json:"documentId,omitempty"`
	DrawingCanDuplicatePartNumber       *bool                     `json:"drawingCanDuplicatePartNumber,omitempty"`
	EnabledActiveMultipleWorkflows      *bool                     `json:"enabledActiveMultipleWorkflows,omitempty"`
	ObsoletionWorkflow                  *BTPublishedWorkflowInfo  `json:"obsoletionWorkflow,omitempty"`
	ObsoletionWorkflowId                *string                   `json:"obsoletionWorkflowId,omitempty"`
	PartNumberingSchemeId               *string                   `json:"partNumberingSchemeId,omitempty"`
	PickableWorkflows                   []BTPublishedWorkflowInfo `json:"pickableWorkflows,omitempty"`
	ReleaseWorkflow                     *BTPublishedWorkflowInfo  `json:"releaseWorkflow,omitempty"`
	ReleaseWorkflowId                   *string                   `json:"releaseWorkflowId,omitempty"`
	ReleaseableApplications             []string                  `json:"releaseableApplications,omitempty"`
	UsingAutoPartNumbering              *bool                     `json:"usingAutoPartNumbering,omitempty"`
	UsingManagedWorkflow                *bool                     `json:"usingManagedWorkflow,omitempty"`
	UsingThirdPartyPartNumbering        *bool                     `json:"usingThirdPartyPartNumbering,omitempty"`
}

// NewBTActiveWorkflowInfo instantiates a new BTActiveWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTActiveWorkflowInfo() *BTActiveWorkflowInfo {
	this := BTActiveWorkflowInfo{}
	return &this
}

// NewBTActiveWorkflowInfoWithDefaults instantiates a new BTActiveWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTActiveWorkflowInfoWithDefaults() *BTActiveWorkflowInfo {
	this := BTActiveWorkflowInfo{}
	return &this
}

// GetAllowReleaseItemsFromOtherDocuments returns the AllowReleaseItemsFromOtherDocuments field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetAllowReleaseItemsFromOtherDocuments() bool {
	if o == nil || o.AllowReleaseItemsFromOtherDocuments == nil {
		var ret bool
		return ret
	}
	return *o.AllowReleaseItemsFromOtherDocuments
}

// GetAllowReleaseItemsFromOtherDocumentsOk returns a tuple with the AllowReleaseItemsFromOtherDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetAllowReleaseItemsFromOtherDocumentsOk() (*bool, bool) {
	if o == nil || o.AllowReleaseItemsFromOtherDocuments == nil {
		return nil, false
	}
	return o.AllowReleaseItemsFromOtherDocuments, true
}

// HasAllowReleaseItemsFromOtherDocuments returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasAllowReleaseItemsFromOtherDocuments() bool {
	if o != nil && o.AllowReleaseItemsFromOtherDocuments != nil {
		return true
	}

	return false
}

// SetAllowReleaseItemsFromOtherDocuments gets a reference to the given bool and assigns it to the AllowReleaseItemsFromOtherDocuments field.
func (o *BTActiveWorkflowInfo) SetAllowReleaseItemsFromOtherDocuments(v bool) {
	o.AllowReleaseItemsFromOtherDocuments = &v
}

// GetCanCurrentUserCreateReleases returns the CanCurrentUserCreateReleases field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleases() bool {
	if o == nil || o.CanCurrentUserCreateReleases == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserCreateReleases
}

// GetCanCurrentUserCreateReleasesOk returns a tuple with the CanCurrentUserCreateReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleasesOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserCreateReleases == nil {
		return nil, false
	}
	return o.CanCurrentUserCreateReleases, true
}

// HasCanCurrentUserCreateReleases returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserCreateReleases() bool {
	if o != nil && o.CanCurrentUserCreateReleases != nil {
		return true
	}

	return false
}

// SetCanCurrentUserCreateReleases gets a reference to the given bool and assigns it to the CanCurrentUserCreateReleases field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserCreateReleases(v bool) {
	o.CanCurrentUserCreateReleases = &v
}

// GetCanCurrentUserManageWorkflows returns the CanCurrentUserManageWorkflows field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserManageWorkflows() bool {
	if o == nil || o.CanCurrentUserManageWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserManageWorkflows
}

// GetCanCurrentUserManageWorkflowsOk returns a tuple with the CanCurrentUserManageWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserManageWorkflowsOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserManageWorkflows == nil {
		return nil, false
	}
	return o.CanCurrentUserManageWorkflows, true
}

// HasCanCurrentUserManageWorkflows returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserManageWorkflows() bool {
	if o != nil && o.CanCurrentUserManageWorkflows != nil {
		return true
	}

	return false
}

// SetCanCurrentUserManageWorkflows gets a reference to the given bool and assigns it to the CanCurrentUserManageWorkflows field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserManageWorkflows(v bool) {
	o.CanCurrentUserManageWorkflows = &v
}

// GetCanCurrentUserSeeArenaItemLink returns the CanCurrentUserSeeArenaItemLink field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSeeArenaItemLink() bool {
	if o == nil || o.CanCurrentUserSeeArenaItemLink == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserSeeArenaItemLink
}

// GetCanCurrentUserSeeArenaItemLinkOk returns a tuple with the CanCurrentUserSeeArenaItemLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSeeArenaItemLinkOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserSeeArenaItemLink == nil {
		return nil, false
	}
	return o.CanCurrentUserSeeArenaItemLink, true
}

// HasCanCurrentUserSeeArenaItemLink returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserSeeArenaItemLink() bool {
	if o != nil && o.CanCurrentUserSeeArenaItemLink != nil {
		return true
	}

	return false
}

// SetCanCurrentUserSeeArenaItemLink gets a reference to the given bool and assigns it to the CanCurrentUserSeeArenaItemLink field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserSeeArenaItemLink(v bool) {
	o.CanCurrentUserSeeArenaItemLink = &v
}

// GetCanCurrentUserSyncBomToArena returns the CanCurrentUserSyncBomToArena field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncBomToArena() bool {
	if o == nil || o.CanCurrentUserSyncBomToArena == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserSyncBomToArena
}

// GetCanCurrentUserSyncBomToArenaOk returns a tuple with the CanCurrentUserSyncBomToArena field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncBomToArenaOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserSyncBomToArena == nil {
		return nil, false
	}
	return o.CanCurrentUserSyncBomToArena, true
}

// HasCanCurrentUserSyncBomToArena returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncBomToArena() bool {
	if o != nil && o.CanCurrentUserSyncBomToArena != nil {
		return true
	}

	return false
}

// SetCanCurrentUserSyncBomToArena gets a reference to the given bool and assigns it to the CanCurrentUserSyncBomToArena field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncBomToArena(v bool) {
	o.CanCurrentUserSyncBomToArena = &v
}

// GetCanCurrentUserSyncToArena returns the CanCurrentUserSyncToArena field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncToArena() bool {
	if o == nil || o.CanCurrentUserSyncToArena == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserSyncToArena
}

// GetCanCurrentUserSyncToArenaOk returns a tuple with the CanCurrentUserSyncToArena field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncToArenaOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserSyncToArena == nil {
		return nil, false
	}
	return o.CanCurrentUserSyncToArena, true
}

// HasCanCurrentUserSyncToArena returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncToArena() bool {
	if o != nil && o.CanCurrentUserSyncToArena != nil {
		return true
	}

	return false
}

// SetCanCurrentUserSyncToArena gets a reference to the given bool and assigns it to the CanCurrentUserSyncToArena field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncToArena(v bool) {
	o.CanCurrentUserSyncToArena = &v
}

// GetCanCurrentUserSyncVersionsToArena returns the CanCurrentUserSyncVersionsToArena field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncVersionsToArena() bool {
	if o == nil || o.CanCurrentUserSyncVersionsToArena == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserSyncVersionsToArena
}

// GetCanCurrentUserSyncVersionsToArenaOk returns a tuple with the CanCurrentUserSyncVersionsToArena field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncVersionsToArenaOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserSyncVersionsToArena == nil {
		return nil, false
	}
	return o.CanCurrentUserSyncVersionsToArena, true
}

// HasCanCurrentUserSyncVersionsToArena returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncVersionsToArena() bool {
	if o != nil && o.CanCurrentUserSyncVersionsToArena != nil {
		return true
	}

	return false
}

// SetCanCurrentUserSyncVersionsToArena gets a reference to the given bool and assigns it to the CanCurrentUserSyncVersionsToArena field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncVersionsToArena(v bool) {
	o.CanCurrentUserSyncVersionsToArena = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTActiveWorkflowInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTActiveWorkflowInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDrawingCanDuplicatePartNumber returns the DrawingCanDuplicatePartNumber field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumber() bool {
	if o == nil || o.DrawingCanDuplicatePartNumber == nil {
		var ret bool
		return ret
	}
	return *o.DrawingCanDuplicatePartNumber
}

// GetDrawingCanDuplicatePartNumberOk returns a tuple with the DrawingCanDuplicatePartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumberOk() (*bool, bool) {
	if o == nil || o.DrawingCanDuplicatePartNumber == nil {
		return nil, false
	}
	return o.DrawingCanDuplicatePartNumber, true
}

// HasDrawingCanDuplicatePartNumber returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasDrawingCanDuplicatePartNumber() bool {
	if o != nil && o.DrawingCanDuplicatePartNumber != nil {
		return true
	}

	return false
}

// SetDrawingCanDuplicatePartNumber gets a reference to the given bool and assigns it to the DrawingCanDuplicatePartNumber field.
func (o *BTActiveWorkflowInfo) SetDrawingCanDuplicatePartNumber(v bool) {
	o.DrawingCanDuplicatePartNumber = &v
}

// GetEnabledActiveMultipleWorkflows returns the EnabledActiveMultipleWorkflows field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflows() bool {
	if o == nil || o.EnabledActiveMultipleWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.EnabledActiveMultipleWorkflows
}

// GetEnabledActiveMultipleWorkflowsOk returns a tuple with the EnabledActiveMultipleWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflowsOk() (*bool, bool) {
	if o == nil || o.EnabledActiveMultipleWorkflows == nil {
		return nil, false
	}
	return o.EnabledActiveMultipleWorkflows, true
}

// HasEnabledActiveMultipleWorkflows returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasEnabledActiveMultipleWorkflows() bool {
	if o != nil && o.EnabledActiveMultipleWorkflows != nil {
		return true
	}

	return false
}

// SetEnabledActiveMultipleWorkflows gets a reference to the given bool and assigns it to the EnabledActiveMultipleWorkflows field.
func (o *BTActiveWorkflowInfo) SetEnabledActiveMultipleWorkflows(v bool) {
	o.EnabledActiveMultipleWorkflows = &v
}

// GetObsoletionWorkflow returns the ObsoletionWorkflow field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflow() BTPublishedWorkflowInfo {
	if o == nil || o.ObsoletionWorkflow == nil {
		var ret BTPublishedWorkflowInfo
		return ret
	}
	return *o.ObsoletionWorkflow
}

// GetObsoletionWorkflowOk returns a tuple with the ObsoletionWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowOk() (*BTPublishedWorkflowInfo, bool) {
	if o == nil || o.ObsoletionWorkflow == nil {
		return nil, false
	}
	return o.ObsoletionWorkflow, true
}

// HasObsoletionWorkflow returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasObsoletionWorkflow() bool {
	if o != nil && o.ObsoletionWorkflow != nil {
		return true
	}

	return false
}

// SetObsoletionWorkflow gets a reference to the given BTPublishedWorkflowInfo and assigns it to the ObsoletionWorkflow field.
func (o *BTActiveWorkflowInfo) SetObsoletionWorkflow(v BTPublishedWorkflowInfo) {
	o.ObsoletionWorkflow = &v
}

// GetObsoletionWorkflowId returns the ObsoletionWorkflowId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowId() string {
	if o == nil || o.ObsoletionWorkflowId == nil {
		var ret string
		return ret
	}
	return *o.ObsoletionWorkflowId
}

// GetObsoletionWorkflowIdOk returns a tuple with the ObsoletionWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowIdOk() (*string, bool) {
	if o == nil || o.ObsoletionWorkflowId == nil {
		return nil, false
	}
	return o.ObsoletionWorkflowId, true
}

// HasObsoletionWorkflowId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasObsoletionWorkflowId() bool {
	if o != nil && o.ObsoletionWorkflowId != nil {
		return true
	}

	return false
}

// SetObsoletionWorkflowId gets a reference to the given string and assigns it to the ObsoletionWorkflowId field.
func (o *BTActiveWorkflowInfo) SetObsoletionWorkflowId(v string) {
	o.ObsoletionWorkflowId = &v
}

// GetPartNumberingSchemeId returns the PartNumberingSchemeId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetPartNumberingSchemeId() string {
	if o == nil || o.PartNumberingSchemeId == nil {
		var ret string
		return ret
	}
	return *o.PartNumberingSchemeId
}

// GetPartNumberingSchemeIdOk returns a tuple with the PartNumberingSchemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetPartNumberingSchemeIdOk() (*string, bool) {
	if o == nil || o.PartNumberingSchemeId == nil {
		return nil, false
	}
	return o.PartNumberingSchemeId, true
}

// HasPartNumberingSchemeId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasPartNumberingSchemeId() bool {
	if o != nil && o.PartNumberingSchemeId != nil {
		return true
	}

	return false
}

// SetPartNumberingSchemeId gets a reference to the given string and assigns it to the PartNumberingSchemeId field.
func (o *BTActiveWorkflowInfo) SetPartNumberingSchemeId(v string) {
	o.PartNumberingSchemeId = &v
}

// GetPickableWorkflows returns the PickableWorkflows field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetPickableWorkflows() []BTPublishedWorkflowInfo {
	if o == nil || o.PickableWorkflows == nil {
		var ret []BTPublishedWorkflowInfo
		return ret
	}
	return o.PickableWorkflows
}

// GetPickableWorkflowsOk returns a tuple with the PickableWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetPickableWorkflowsOk() ([]BTPublishedWorkflowInfo, bool) {
	if o == nil || o.PickableWorkflows == nil {
		return nil, false
	}
	return o.PickableWorkflows, true
}

// HasPickableWorkflows returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasPickableWorkflows() bool {
	if o != nil && o.PickableWorkflows != nil {
		return true
	}

	return false
}

// SetPickableWorkflows gets a reference to the given []BTPublishedWorkflowInfo and assigns it to the PickableWorkflows field.
func (o *BTActiveWorkflowInfo) SetPickableWorkflows(v []BTPublishedWorkflowInfo) {
	o.PickableWorkflows = v
}

// GetReleaseWorkflow returns the ReleaseWorkflow field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflow() BTPublishedWorkflowInfo {
	if o == nil || o.ReleaseWorkflow == nil {
		var ret BTPublishedWorkflowInfo
		return ret
	}
	return *o.ReleaseWorkflow
}

// GetReleaseWorkflowOk returns a tuple with the ReleaseWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflowOk() (*BTPublishedWorkflowInfo, bool) {
	if o == nil || o.ReleaseWorkflow == nil {
		return nil, false
	}
	return o.ReleaseWorkflow, true
}

// HasReleaseWorkflow returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasReleaseWorkflow() bool {
	if o != nil && o.ReleaseWorkflow != nil {
		return true
	}

	return false
}

// SetReleaseWorkflow gets a reference to the given BTPublishedWorkflowInfo and assigns it to the ReleaseWorkflow field.
func (o *BTActiveWorkflowInfo) SetReleaseWorkflow(v BTPublishedWorkflowInfo) {
	o.ReleaseWorkflow = &v
}

// GetReleaseWorkflowId returns the ReleaseWorkflowId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflowId() string {
	if o == nil || o.ReleaseWorkflowId == nil {
		var ret string
		return ret
	}
	return *o.ReleaseWorkflowId
}

// GetReleaseWorkflowIdOk returns a tuple with the ReleaseWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflowIdOk() (*string, bool) {
	if o == nil || o.ReleaseWorkflowId == nil {
		return nil, false
	}
	return o.ReleaseWorkflowId, true
}

// HasReleaseWorkflowId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasReleaseWorkflowId() bool {
	if o != nil && o.ReleaseWorkflowId != nil {
		return true
	}

	return false
}

// SetReleaseWorkflowId gets a reference to the given string and assigns it to the ReleaseWorkflowId field.
func (o *BTActiveWorkflowInfo) SetReleaseWorkflowId(v string) {
	o.ReleaseWorkflowId = &v
}

// GetReleaseableApplications returns the ReleaseableApplications field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetReleaseableApplications() []string {
	if o == nil || o.ReleaseableApplications == nil {
		var ret []string
		return ret
	}
	return o.ReleaseableApplications
}

// GetReleaseableApplicationsOk returns a tuple with the ReleaseableApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetReleaseableApplicationsOk() ([]string, bool) {
	if o == nil || o.ReleaseableApplications == nil {
		return nil, false
	}
	return o.ReleaseableApplications, true
}

// HasReleaseableApplications returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasReleaseableApplications() bool {
	if o != nil && o.ReleaseableApplications != nil {
		return true
	}

	return false
}

// SetReleaseableApplications gets a reference to the given []string and assigns it to the ReleaseableApplications field.
func (o *BTActiveWorkflowInfo) SetReleaseableApplications(v []string) {
	o.ReleaseableApplications = v
}

// GetUsingAutoPartNumbering returns the UsingAutoPartNumbering field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumbering() bool {
	if o == nil || o.UsingAutoPartNumbering == nil {
		var ret bool
		return ret
	}
	return *o.UsingAutoPartNumbering
}

// GetUsingAutoPartNumberingOk returns a tuple with the UsingAutoPartNumbering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingOk() (*bool, bool) {
	if o == nil || o.UsingAutoPartNumbering == nil {
		return nil, false
	}
	return o.UsingAutoPartNumbering, true
}

// HasUsingAutoPartNumbering returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasUsingAutoPartNumbering() bool {
	if o != nil && o.UsingAutoPartNumbering != nil {
		return true
	}

	return false
}

// SetUsingAutoPartNumbering gets a reference to the given bool and assigns it to the UsingAutoPartNumbering field.
func (o *BTActiveWorkflowInfo) SetUsingAutoPartNumbering(v bool) {
	o.UsingAutoPartNumbering = &v
}

// GetUsingManagedWorkflow returns the UsingManagedWorkflow field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflow() bool {
	if o == nil || o.UsingManagedWorkflow == nil {
		var ret bool
		return ret
	}
	return *o.UsingManagedWorkflow
}

// GetUsingManagedWorkflowOk returns a tuple with the UsingManagedWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflowOk() (*bool, bool) {
	if o == nil || o.UsingManagedWorkflow == nil {
		return nil, false
	}
	return o.UsingManagedWorkflow, true
}

// HasUsingManagedWorkflow returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasUsingManagedWorkflow() bool {
	if o != nil && o.UsingManagedWorkflow != nil {
		return true
	}

	return false
}

// SetUsingManagedWorkflow gets a reference to the given bool and assigns it to the UsingManagedWorkflow field.
func (o *BTActiveWorkflowInfo) SetUsingManagedWorkflow(v bool) {
	o.UsingManagedWorkflow = &v
}

// GetUsingThirdPartyPartNumbering returns the UsingThirdPartyPartNumbering field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetUsingThirdPartyPartNumbering() bool {
	if o == nil || o.UsingThirdPartyPartNumbering == nil {
		var ret bool
		return ret
	}
	return *o.UsingThirdPartyPartNumbering
}

// GetUsingThirdPartyPartNumberingOk returns a tuple with the UsingThirdPartyPartNumbering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetUsingThirdPartyPartNumberingOk() (*bool, bool) {
	if o == nil || o.UsingThirdPartyPartNumbering == nil {
		return nil, false
	}
	return o.UsingThirdPartyPartNumbering, true
}

// HasUsingThirdPartyPartNumbering returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasUsingThirdPartyPartNumbering() bool {
	if o != nil && o.UsingThirdPartyPartNumbering != nil {
		return true
	}

	return false
}

// SetUsingThirdPartyPartNumbering gets a reference to the given bool and assigns it to the UsingThirdPartyPartNumbering field.
func (o *BTActiveWorkflowInfo) SetUsingThirdPartyPartNumbering(v bool) {
	o.UsingThirdPartyPartNumbering = &v
}

func (o BTActiveWorkflowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowReleaseItemsFromOtherDocuments != nil {
		toSerialize["allowReleaseItemsFromOtherDocuments"] = o.AllowReleaseItemsFromOtherDocuments
	}
	if o.CanCurrentUserCreateReleases != nil {
		toSerialize["canCurrentUserCreateReleases"] = o.CanCurrentUserCreateReleases
	}
	if o.CanCurrentUserManageWorkflows != nil {
		toSerialize["canCurrentUserManageWorkflows"] = o.CanCurrentUserManageWorkflows
	}
	if o.CanCurrentUserSeeArenaItemLink != nil {
		toSerialize["canCurrentUserSeeArenaItemLink"] = o.CanCurrentUserSeeArenaItemLink
	}
	if o.CanCurrentUserSyncBomToArena != nil {
		toSerialize["canCurrentUserSyncBomToArena"] = o.CanCurrentUserSyncBomToArena
	}
	if o.CanCurrentUserSyncToArena != nil {
		toSerialize["canCurrentUserSyncToArena"] = o.CanCurrentUserSyncToArena
	}
	if o.CanCurrentUserSyncVersionsToArena != nil {
		toSerialize["canCurrentUserSyncVersionsToArena"] = o.CanCurrentUserSyncVersionsToArena
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DrawingCanDuplicatePartNumber != nil {
		toSerialize["drawingCanDuplicatePartNumber"] = o.DrawingCanDuplicatePartNumber
	}
	if o.EnabledActiveMultipleWorkflows != nil {
		toSerialize["enabledActiveMultipleWorkflows"] = o.EnabledActiveMultipleWorkflows
	}
	if o.ObsoletionWorkflow != nil {
		toSerialize["obsoletionWorkflow"] = o.ObsoletionWorkflow
	}
	if o.ObsoletionWorkflowId != nil {
		toSerialize["obsoletionWorkflowId"] = o.ObsoletionWorkflowId
	}
	if o.PartNumberingSchemeId != nil {
		toSerialize["partNumberingSchemeId"] = o.PartNumberingSchemeId
	}
	if o.PickableWorkflows != nil {
		toSerialize["pickableWorkflows"] = o.PickableWorkflows
	}
	if o.ReleaseWorkflow != nil {
		toSerialize["releaseWorkflow"] = o.ReleaseWorkflow
	}
	if o.ReleaseWorkflowId != nil {
		toSerialize["releaseWorkflowId"] = o.ReleaseWorkflowId
	}
	if o.ReleaseableApplications != nil {
		toSerialize["releaseableApplications"] = o.ReleaseableApplications
	}
	if o.UsingAutoPartNumbering != nil {
		toSerialize["usingAutoPartNumbering"] = o.UsingAutoPartNumbering
	}
	if o.UsingManagedWorkflow != nil {
		toSerialize["usingManagedWorkflow"] = o.UsingManagedWorkflow
	}
	if o.UsingThirdPartyPartNumbering != nil {
		toSerialize["usingThirdPartyPartNumbering"] = o.UsingThirdPartyPartNumbering
	}
	return json.Marshal(toSerialize)
}

type NullableBTActiveWorkflowInfo struct {
	value *BTActiveWorkflowInfo
	isSet bool
}

func (v NullableBTActiveWorkflowInfo) Get() *BTActiveWorkflowInfo {
	return v.value
}

func (v *NullableBTActiveWorkflowInfo) Set(val *BTActiveWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTActiveWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTActiveWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTActiveWorkflowInfo(val *BTActiveWorkflowInfo) *NullableBTActiveWorkflowInfo {
	return &NullableBTActiveWorkflowInfo{value: val, isSet: true}
}

func (v NullableBTActiveWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTActiveWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
