/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14179-e31166ce6183
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCommentInfo struct for BTCommentInfo
type BTCommentInfo struct {
	AssemblyFeatures   []string                 `json:"assemblyFeatures,omitempty"`
	AssignedAt         *JSONTime                `json:"assignedAt,omitempty"`
	Assignee           *BTUserSummaryInfo       `json:"assignee,omitempty"`
	Attachment         *BTCommentAttachmentInfo `json:"attachment,omitempty"`
	CanDelete          *bool                    `json:"canDelete,omitempty"`
	CanResolveOrReopen *bool                    `json:"canResolveOrReopen,omitempty"`
	CreatedAt          *JSONTime                `json:"createdAt,omitempty"`
	DocumentId         *string                  `json:"documentId,omitempty"`
	ElementFeature     *string                  `json:"elementFeature,omitempty"`
	ElementId          *string                  `json:"elementId,omitempty"`
	ElementOccurrences []string                 `json:"elementOccurrences,omitempty"`
	ElementQuery       *string                  `json:"elementQuery,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id      *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
	// Name of the resource.
	Name             *string                  `json:"name,omitempty"`
	ObjectId         *string                  `json:"objectId,omitempty"`
	ObjectType       *int32                   `json:"objectType,omitempty"`
	ParentId         *string                  `json:"parentId,omitempty"`
	ReleasePackageId *string                  `json:"releasePackageId,omitempty"`
	ReopenedAt       *JSONTime                `json:"reopenedAt,omitempty"`
	ReopenedBy       *BTUserSummaryInfo       `json:"reopenedBy,omitempty"`
	ReplyCount       *int64                   `json:"replyCount,omitempty"`
	ResolvedAt       *JSONTime                `json:"resolvedAt,omitempty"`
	ResolvedBy       *BTUserSummaryInfo       `json:"resolvedBy,omitempty"`
	State            *int32                   `json:"state,omitempty"`
	Thumbnail        *BTCommentAttachmentInfo `json:"thumbnail,omitempty"`
	TopLevel         *bool                    `json:"topLevel,omitempty"`
	User             *BTUserSummaryInfo       `json:"user,omitempty"`
	VersionId        *string                  `json:"versionId,omitempty"`
	ViewData         *BTViewDataInfo          `json:"viewData,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef     *string `json:"viewRef,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

// NewBTCommentInfo instantiates a new BTCommentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCommentInfo() *BTCommentInfo {
	this := BTCommentInfo{}
	return &this
}

// NewBTCommentInfoWithDefaults instantiates a new BTCommentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCommentInfoWithDefaults() *BTCommentInfo {
	this := BTCommentInfo{}
	return &this
}

// GetAssemblyFeatures returns the AssemblyFeatures field value if set, zero value otherwise.
func (o *BTCommentInfo) GetAssemblyFeatures() []string {
	if o == nil || o.AssemblyFeatures == nil {
		var ret []string
		return ret
	}
	return o.AssemblyFeatures
}

// GetAssemblyFeaturesOk returns a tuple with the AssemblyFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetAssemblyFeaturesOk() ([]string, bool) {
	if o == nil || o.AssemblyFeatures == nil {
		return nil, false
	}
	return o.AssemblyFeatures, true
}

// HasAssemblyFeatures returns a boolean if a field has been set.
func (o *BTCommentInfo) HasAssemblyFeatures() bool {
	if o != nil && o.AssemblyFeatures != nil {
		return true
	}

	return false
}

// SetAssemblyFeatures gets a reference to the given []string and assigns it to the AssemblyFeatures field.
func (o *BTCommentInfo) SetAssemblyFeatures(v []string) {
	o.AssemblyFeatures = v
}

// GetAssignedAt returns the AssignedAt field value if set, zero value otherwise.
func (o *BTCommentInfo) GetAssignedAt() JSONTime {
	if o == nil || o.AssignedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.AssignedAt
}

// GetAssignedAtOk returns a tuple with the AssignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetAssignedAtOk() (*JSONTime, bool) {
	if o == nil || o.AssignedAt == nil {
		return nil, false
	}
	return o.AssignedAt, true
}

// HasAssignedAt returns a boolean if a field has been set.
func (o *BTCommentInfo) HasAssignedAt() bool {
	if o != nil && o.AssignedAt != nil {
		return true
	}

	return false
}

// SetAssignedAt gets a reference to the given JSONTime and assigns it to the AssignedAt field.
func (o *BTCommentInfo) SetAssignedAt(v JSONTime) {
	o.AssignedAt = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *BTCommentInfo) GetAssignee() BTUserSummaryInfo {
	if o == nil || o.Assignee == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetAssigneeOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *BTCommentInfo) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given BTUserSummaryInfo and assigns it to the Assignee field.
func (o *BTCommentInfo) SetAssignee(v BTUserSummaryInfo) {
	o.Assignee = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *BTCommentInfo) GetAttachment() BTCommentAttachmentInfo {
	if o == nil || o.Attachment == nil {
		var ret BTCommentAttachmentInfo
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetAttachmentOk() (*BTCommentAttachmentInfo, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *BTCommentInfo) HasAttachment() bool {
	if o != nil && o.Attachment != nil {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given BTCommentAttachmentInfo and assigns it to the Attachment field.
func (o *BTCommentInfo) SetAttachment(v BTCommentAttachmentInfo) {
	o.Attachment = &v
}

// GetCanDelete returns the CanDelete field value if set, zero value otherwise.
func (o *BTCommentInfo) GetCanDelete() bool {
	if o == nil || o.CanDelete == nil {
		var ret bool
		return ret
	}
	return *o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetCanDeleteOk() (*bool, bool) {
	if o == nil || o.CanDelete == nil {
		return nil, false
	}
	return o.CanDelete, true
}

// HasCanDelete returns a boolean if a field has been set.
func (o *BTCommentInfo) HasCanDelete() bool {
	if o != nil && o.CanDelete != nil {
		return true
	}

	return false
}

// SetCanDelete gets a reference to the given bool and assigns it to the CanDelete field.
func (o *BTCommentInfo) SetCanDelete(v bool) {
	o.CanDelete = &v
}

// GetCanResolveOrReopen returns the CanResolveOrReopen field value if set, zero value otherwise.
func (o *BTCommentInfo) GetCanResolveOrReopen() bool {
	if o == nil || o.CanResolveOrReopen == nil {
		var ret bool
		return ret
	}
	return *o.CanResolveOrReopen
}

// GetCanResolveOrReopenOk returns a tuple with the CanResolveOrReopen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetCanResolveOrReopenOk() (*bool, bool) {
	if o == nil || o.CanResolveOrReopen == nil {
		return nil, false
	}
	return o.CanResolveOrReopen, true
}

// HasCanResolveOrReopen returns a boolean if a field has been set.
func (o *BTCommentInfo) HasCanResolveOrReopen() bool {
	if o != nil && o.CanResolveOrReopen != nil {
		return true
	}

	return false
}

// SetCanResolveOrReopen gets a reference to the given bool and assigns it to the CanResolveOrReopen field.
func (o *BTCommentInfo) SetCanResolveOrReopen(v bool) {
	o.CanResolveOrReopen = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTCommentInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTCommentInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTCommentInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTCommentInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementFeature returns the ElementFeature field value if set, zero value otherwise.
func (o *BTCommentInfo) GetElementFeature() string {
	if o == nil || o.ElementFeature == nil {
		var ret string
		return ret
	}
	return *o.ElementFeature
}

// GetElementFeatureOk returns a tuple with the ElementFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetElementFeatureOk() (*string, bool) {
	if o == nil || o.ElementFeature == nil {
		return nil, false
	}
	return o.ElementFeature, true
}

// HasElementFeature returns a boolean if a field has been set.
func (o *BTCommentInfo) HasElementFeature() bool {
	if o != nil && o.ElementFeature != nil {
		return true
	}

	return false
}

// SetElementFeature gets a reference to the given string and assigns it to the ElementFeature field.
func (o *BTCommentInfo) SetElementFeature(v string) {
	o.ElementFeature = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTCommentInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementOccurrences returns the ElementOccurrences field value if set, zero value otherwise.
func (o *BTCommentInfo) GetElementOccurrences() []string {
	if o == nil || o.ElementOccurrences == nil {
		var ret []string
		return ret
	}
	return o.ElementOccurrences
}

// GetElementOccurrencesOk returns a tuple with the ElementOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetElementOccurrencesOk() ([]string, bool) {
	if o == nil || o.ElementOccurrences == nil {
		return nil, false
	}
	return o.ElementOccurrences, true
}

// HasElementOccurrences returns a boolean if a field has been set.
func (o *BTCommentInfo) HasElementOccurrences() bool {
	if o != nil && o.ElementOccurrences != nil {
		return true
	}

	return false
}

// SetElementOccurrences gets a reference to the given []string and assigns it to the ElementOccurrences field.
func (o *BTCommentInfo) SetElementOccurrences(v []string) {
	o.ElementOccurrences = v
}

// GetElementQuery returns the ElementQuery field value if set, zero value otherwise.
func (o *BTCommentInfo) GetElementQuery() string {
	if o == nil || o.ElementQuery == nil {
		var ret string
		return ret
	}
	return *o.ElementQuery
}

// GetElementQueryOk returns a tuple with the ElementQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetElementQueryOk() (*string, bool) {
	if o == nil || o.ElementQuery == nil {
		return nil, false
	}
	return o.ElementQuery, true
}

// HasElementQuery returns a boolean if a field has been set.
func (o *BTCommentInfo) HasElementQuery() bool {
	if o != nil && o.ElementQuery != nil {
		return true
	}

	return false
}

// SetElementQuery gets a reference to the given string and assigns it to the ElementQuery field.
func (o *BTCommentInfo) SetElementQuery(v string) {
	o.ElementQuery = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCommentInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCommentInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCommentInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCommentInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCommentInfo) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BTCommentInfo) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BTCommentInfo) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BTCommentInfo) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCommentInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCommentInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCommentInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTCommentInfo) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTCommentInfo) GetObjectType() int32 {
	if o == nil || o.ObjectType == nil {
		var ret int32
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetObjectTypeOk() (*int32, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTCommentInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int32 and assigns it to the ObjectType field.
func (o *BTCommentInfo) SetObjectType(v int32) {
	o.ObjectType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTCommentInfo) SetParentId(v string) {
	o.ParentId = &v
}

// GetReleasePackageId returns the ReleasePackageId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetReleasePackageId() string {
	if o == nil || o.ReleasePackageId == nil {
		var ret string
		return ret
	}
	return *o.ReleasePackageId
}

// GetReleasePackageIdOk returns a tuple with the ReleasePackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetReleasePackageIdOk() (*string, bool) {
	if o == nil || o.ReleasePackageId == nil {
		return nil, false
	}
	return o.ReleasePackageId, true
}

// HasReleasePackageId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasReleasePackageId() bool {
	if o != nil && o.ReleasePackageId != nil {
		return true
	}

	return false
}

// SetReleasePackageId gets a reference to the given string and assigns it to the ReleasePackageId field.
func (o *BTCommentInfo) SetReleasePackageId(v string) {
	o.ReleasePackageId = &v
}

// GetReopenedAt returns the ReopenedAt field value if set, zero value otherwise.
func (o *BTCommentInfo) GetReopenedAt() JSONTime {
	if o == nil || o.ReopenedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ReopenedAt
}

// GetReopenedAtOk returns a tuple with the ReopenedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetReopenedAtOk() (*JSONTime, bool) {
	if o == nil || o.ReopenedAt == nil {
		return nil, false
	}
	return o.ReopenedAt, true
}

// HasReopenedAt returns a boolean if a field has been set.
func (o *BTCommentInfo) HasReopenedAt() bool {
	if o != nil && o.ReopenedAt != nil {
		return true
	}

	return false
}

// SetReopenedAt gets a reference to the given JSONTime and assigns it to the ReopenedAt field.
func (o *BTCommentInfo) SetReopenedAt(v JSONTime) {
	o.ReopenedAt = &v
}

// GetReopenedBy returns the ReopenedBy field value if set, zero value otherwise.
func (o *BTCommentInfo) GetReopenedBy() BTUserSummaryInfo {
	if o == nil || o.ReopenedBy == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.ReopenedBy
}

// GetReopenedByOk returns a tuple with the ReopenedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetReopenedByOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.ReopenedBy == nil {
		return nil, false
	}
	return o.ReopenedBy, true
}

// HasReopenedBy returns a boolean if a field has been set.
func (o *BTCommentInfo) HasReopenedBy() bool {
	if o != nil && o.ReopenedBy != nil {
		return true
	}

	return false
}

// SetReopenedBy gets a reference to the given BTUserSummaryInfo and assigns it to the ReopenedBy field.
func (o *BTCommentInfo) SetReopenedBy(v BTUserSummaryInfo) {
	o.ReopenedBy = &v
}

// GetReplyCount returns the ReplyCount field value if set, zero value otherwise.
func (o *BTCommentInfo) GetReplyCount() int64 {
	if o == nil || o.ReplyCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplyCount
}

// GetReplyCountOk returns a tuple with the ReplyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetReplyCountOk() (*int64, bool) {
	if o == nil || o.ReplyCount == nil {
		return nil, false
	}
	return o.ReplyCount, true
}

// HasReplyCount returns a boolean if a field has been set.
func (o *BTCommentInfo) HasReplyCount() bool {
	if o != nil && o.ReplyCount != nil {
		return true
	}

	return false
}

// SetReplyCount gets a reference to the given int64 and assigns it to the ReplyCount field.
func (o *BTCommentInfo) SetReplyCount(v int64) {
	o.ReplyCount = &v
}

// GetResolvedAt returns the ResolvedAt field value if set, zero value otherwise.
func (o *BTCommentInfo) GetResolvedAt() JSONTime {
	if o == nil || o.ResolvedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ResolvedAt
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetResolvedAtOk() (*JSONTime, bool) {
	if o == nil || o.ResolvedAt == nil {
		return nil, false
	}
	return o.ResolvedAt, true
}

// HasResolvedAt returns a boolean if a field has been set.
func (o *BTCommentInfo) HasResolvedAt() bool {
	if o != nil && o.ResolvedAt != nil {
		return true
	}

	return false
}

// SetResolvedAt gets a reference to the given JSONTime and assigns it to the ResolvedAt field.
func (o *BTCommentInfo) SetResolvedAt(v JSONTime) {
	o.ResolvedAt = &v
}

// GetResolvedBy returns the ResolvedBy field value if set, zero value otherwise.
func (o *BTCommentInfo) GetResolvedBy() BTUserSummaryInfo {
	if o == nil || o.ResolvedBy == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.ResolvedBy
}

// GetResolvedByOk returns a tuple with the ResolvedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetResolvedByOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.ResolvedBy == nil {
		return nil, false
	}
	return o.ResolvedBy, true
}

// HasResolvedBy returns a boolean if a field has been set.
func (o *BTCommentInfo) HasResolvedBy() bool {
	if o != nil && o.ResolvedBy != nil {
		return true
	}

	return false
}

// SetResolvedBy gets a reference to the given BTUserSummaryInfo and assigns it to the ResolvedBy field.
func (o *BTCommentInfo) SetResolvedBy(v BTUserSummaryInfo) {
	o.ResolvedBy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTCommentInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTCommentInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTCommentInfo) SetState(v int32) {
	o.State = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *BTCommentInfo) GetThumbnail() BTCommentAttachmentInfo {
	if o == nil || o.Thumbnail == nil {
		var ret BTCommentAttachmentInfo
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetThumbnailOk() (*BTCommentAttachmentInfo, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *BTCommentInfo) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given BTCommentAttachmentInfo and assigns it to the Thumbnail field.
func (o *BTCommentInfo) SetThumbnail(v BTCommentAttachmentInfo) {
	o.Thumbnail = &v
}

// GetTopLevel returns the TopLevel field value if set, zero value otherwise.
func (o *BTCommentInfo) GetTopLevel() bool {
	if o == nil || o.TopLevel == nil {
		var ret bool
		return ret
	}
	return *o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetTopLevelOk() (*bool, bool) {
	if o == nil || o.TopLevel == nil {
		return nil, false
	}
	return o.TopLevel, true
}

// HasTopLevel returns a boolean if a field has been set.
func (o *BTCommentInfo) HasTopLevel() bool {
	if o != nil && o.TopLevel != nil {
		return true
	}

	return false
}

// SetTopLevel gets a reference to the given bool and assigns it to the TopLevel field.
func (o *BTCommentInfo) SetTopLevel(v bool) {
	o.TopLevel = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *BTCommentInfo) GetUser() BTUserSummaryInfo {
	if o == nil || o.User == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetUserOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *BTCommentInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given BTUserSummaryInfo and assigns it to the User field.
func (o *BTCommentInfo) SetUser(v BTUserSummaryInfo) {
	o.User = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTCommentInfo) SetVersionId(v string) {
	o.VersionId = &v
}

// GetViewData returns the ViewData field value if set, zero value otherwise.
func (o *BTCommentInfo) GetViewData() BTViewDataInfo {
	if o == nil || o.ViewData == nil {
		var ret BTViewDataInfo
		return ret
	}
	return *o.ViewData
}

// GetViewDataOk returns a tuple with the ViewData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetViewDataOk() (*BTViewDataInfo, bool) {
	if o == nil || o.ViewData == nil {
		return nil, false
	}
	return o.ViewData, true
}

// HasViewData returns a boolean if a field has been set.
func (o *BTCommentInfo) HasViewData() bool {
	if o != nil && o.ViewData != nil {
		return true
	}

	return false
}

// SetViewData gets a reference to the given BTViewDataInfo and assigns it to the ViewData field.
func (o *BTCommentInfo) SetViewData(v BTViewDataInfo) {
	o.ViewData = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCommentInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCommentInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCommentInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTCommentInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTCommentInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTCommentInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTCommentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyFeatures != nil {
		toSerialize["assemblyFeatures"] = o.AssemblyFeatures
	}
	if o.AssignedAt != nil {
		toSerialize["assignedAt"] = o.AssignedAt
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}
	if o.CanDelete != nil {
		toSerialize["canDelete"] = o.CanDelete
	}
	if o.CanResolveOrReopen != nil {
		toSerialize["canResolveOrReopen"] = o.CanResolveOrReopen
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementFeature != nil {
		toSerialize["elementFeature"] = o.ElementFeature
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementOccurrences != nil {
		toSerialize["elementOccurrences"] = o.ElementOccurrences
	}
	if o.ElementQuery != nil {
		toSerialize["elementQuery"] = o.ElementQuery
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ReleasePackageId != nil {
		toSerialize["releasePackageId"] = o.ReleasePackageId
	}
	if o.ReopenedAt != nil {
		toSerialize["reopenedAt"] = o.ReopenedAt
	}
	if o.ReopenedBy != nil {
		toSerialize["reopenedBy"] = o.ReopenedBy
	}
	if o.ReplyCount != nil {
		toSerialize["replyCount"] = o.ReplyCount
	}
	if o.ResolvedAt != nil {
		toSerialize["resolvedAt"] = o.ResolvedAt
	}
	if o.ResolvedBy != nil {
		toSerialize["resolvedBy"] = o.ResolvedBy
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Thumbnail != nil {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if o.TopLevel != nil {
		toSerialize["topLevel"] = o.TopLevel
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.ViewData != nil {
		toSerialize["viewData"] = o.ViewData
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTCommentInfo struct {
	value *BTCommentInfo
	isSet bool
}

func (v NullableBTCommentInfo) Get() *BTCommentInfo {
	return v.value
}

func (v *NullableBTCommentInfo) Set(val *BTCommentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCommentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCommentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCommentInfo(val *BTCommentInfo) *NullableBTCommentInfo {
	return &NullableBTCommentInfo{value: val, isSet: true}
}

func (v NullableBTCommentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCommentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
