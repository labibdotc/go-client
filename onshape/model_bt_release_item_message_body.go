/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5836-ea08b349dac9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTReleaseItemMessageBody struct for BTReleaseItemMessageBody
type BTReleaseItemMessageBody struct {
	AppElementSessionId *string   `json:"appElementSessionId,omitempty"`
	CommentId           *string   `json:"commentId,omitempty"`
	Data                *string   `json:"data,omitempty"`
	DocumentId          *string   `json:"documentId,omitempty"`
	DocumentState       *string   `json:"documentState,omitempty"`
	ElementId           *string   `json:"elementId,omitempty"`
	Event               *string   `json:"event,omitempty"`
	MessageId           *string   `json:"messageId,omitempty"`
	MetadataObjectType  *string   `json:"metadataObjectType,omitempty"`
	PartId              *string   `json:"partId,omitempty"`
	PartIdentity        *string   `json:"partIdentity,omitempty"`
	PartNumber          *string   `json:"partNumber,omitempty"`
	Timestamp           *JSONTime `json:"timestamp,omitempty"`
	TranslatationId     *string   `json:"translatationId,omitempty"`
	TranslationId       *string   `json:"translationId,omitempty"`
	UserId              *string   `json:"userId,omitempty"`
	VersionId           *string   `json:"versionId,omitempty"`
	WebhookId           *string   `json:"webhookId,omitempty"`
	WorkspaceId         *string   `json:"workspaceId,omitempty"`
	ElementType         *int32    `json:"elementType,omitempty"`
	IsFromInitialState  *bool     `json:"isFromInitialState,omitempty"`
	IsToTerminalState   *bool     `json:"isToTerminalState,omitempty"`
	ReleaseId           *string   `json:"releaseId,omitempty"`
	Status              *string   `json:"status,omitempty"`
	TransitionName      *string   `json:"transitionName,omitempty"`
}

// NewBTReleaseItemMessageBody instantiates a new BTReleaseItemMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleaseItemMessageBody() *BTReleaseItemMessageBody {
	this := BTReleaseItemMessageBody{}
	return &this
}

// NewBTReleaseItemMessageBodyWithDefaults instantiates a new BTReleaseItemMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleaseItemMessageBodyWithDefaults() *BTReleaseItemMessageBody {
	this := BTReleaseItemMessageBody{}
	return &this
}

// GetAppElementSessionId returns the AppElementSessionId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetAppElementSessionId() string {
	if o == nil || o.AppElementSessionId == nil {
		var ret string
		return ret
	}
	return *o.AppElementSessionId
}

// GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetAppElementSessionIdOk() (*string, bool) {
	if o == nil || o.AppElementSessionId == nil {
		return nil, false
	}
	return o.AppElementSessionId, true
}

// HasAppElementSessionId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasAppElementSessionId() bool {
	if o != nil && o.AppElementSessionId != nil {
		return true
	}

	return false
}

// SetAppElementSessionId gets a reference to the given string and assigns it to the AppElementSessionId field.
func (o *BTReleaseItemMessageBody) SetAppElementSessionId(v string) {
	o.AppElementSessionId = &v
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetCommentId() string {
	if o == nil || o.CommentId == nil {
		var ret string
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetCommentIdOk() (*string, bool) {
	if o == nil || o.CommentId == nil {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasCommentId() bool {
	if o != nil && o.CommentId != nil {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given string and assigns it to the CommentId field.
func (o *BTReleaseItemMessageBody) SetCommentId(v string) {
	o.CommentId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTReleaseItemMessageBody) SetData(v string) {
	o.Data = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTReleaseItemMessageBody) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentState returns the DocumentState field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetDocumentState() string {
	if o == nil || o.DocumentState == nil {
		var ret string
		return ret
	}
	return *o.DocumentState
}

// GetDocumentStateOk returns a tuple with the DocumentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetDocumentStateOk() (*string, bool) {
	if o == nil || o.DocumentState == nil {
		return nil, false
	}
	return o.DocumentState, true
}

// HasDocumentState returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasDocumentState() bool {
	if o != nil && o.DocumentState != nil {
		return true
	}

	return false
}

// SetDocumentState gets a reference to the given string and assigns it to the DocumentState field.
func (o *BTReleaseItemMessageBody) SetDocumentState(v string) {
	o.DocumentState = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTReleaseItemMessageBody) SetElementId(v string) {
	o.ElementId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *BTReleaseItemMessageBody) SetEvent(v string) {
	o.Event = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *BTReleaseItemMessageBody) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMetadataObjectType returns the MetadataObjectType field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetMetadataObjectType() string {
	if o == nil || o.MetadataObjectType == nil {
		var ret string
		return ret
	}
	return *o.MetadataObjectType
}

// GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetMetadataObjectTypeOk() (*string, bool) {
	if o == nil || o.MetadataObjectType == nil {
		return nil, false
	}
	return o.MetadataObjectType, true
}

// HasMetadataObjectType returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasMetadataObjectType() bool {
	if o != nil && o.MetadataObjectType != nil {
		return true
	}

	return false
}

// SetMetadataObjectType gets a reference to the given string and assigns it to the MetadataObjectType field.
func (o *BTReleaseItemMessageBody) SetMetadataObjectType(v string) {
	o.MetadataObjectType = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTReleaseItemMessageBody) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetPartIdentity() string {
	if o == nil || o.PartIdentity == nil {
		var ret string
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetPartIdentityOk() (*string, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given string and assigns it to the PartIdentity field.
func (o *BTReleaseItemMessageBody) SetPartIdentity(v string) {
	o.PartIdentity = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTReleaseItemMessageBody) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetTimestamp() JSONTime {
	if o == nil || o.Timestamp == nil {
		var ret JSONTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetTimestampOk() (*JSONTime, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given JSONTime and assigns it to the Timestamp field.
func (o *BTReleaseItemMessageBody) SetTimestamp(v JSONTime) {
	o.Timestamp = &v
}

// GetTranslatationId returns the TranslatationId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetTranslatationId() string {
	if o == nil || o.TranslatationId == nil {
		var ret string
		return ret
	}
	return *o.TranslatationId
}

// GetTranslatationIdOk returns a tuple with the TranslatationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetTranslatationIdOk() (*string, bool) {
	if o == nil || o.TranslatationId == nil {
		return nil, false
	}
	return o.TranslatationId, true
}

// HasTranslatationId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasTranslatationId() bool {
	if o != nil && o.TranslatationId != nil {
		return true
	}

	return false
}

// SetTranslatationId gets a reference to the given string and assigns it to the TranslatationId field.
func (o *BTReleaseItemMessageBody) SetTranslatationId(v string) {
	o.TranslatationId = &v
}

// GetTranslationId returns the TranslationId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetTranslationId() string {
	if o == nil || o.TranslationId == nil {
		var ret string
		return ret
	}
	return *o.TranslationId
}

// GetTranslationIdOk returns a tuple with the TranslationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetTranslationIdOk() (*string, bool) {
	if o == nil || o.TranslationId == nil {
		return nil, false
	}
	return o.TranslationId, true
}

// HasTranslationId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasTranslationId() bool {
	if o != nil && o.TranslationId != nil {
		return true
	}

	return false
}

// SetTranslationId gets a reference to the given string and assigns it to the TranslationId field.
func (o *BTReleaseItemMessageBody) SetTranslationId(v string) {
	o.TranslationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTReleaseItemMessageBody) SetUserId(v string) {
	o.UserId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTReleaseItemMessageBody) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *BTReleaseItemMessageBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTReleaseItemMessageBody) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetElementType() int32 {
	if o == nil || o.ElementType == nil {
		var ret int32
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetElementTypeOk() (*int32, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given int32 and assigns it to the ElementType field.
func (o *BTReleaseItemMessageBody) SetElementType(v int32) {
	o.ElementType = &v
}

// GetIsFromInitialState returns the IsFromInitialState field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetIsFromInitialState() bool {
	if o == nil || o.IsFromInitialState == nil {
		var ret bool
		return ret
	}
	return *o.IsFromInitialState
}

// GetIsFromInitialStateOk returns a tuple with the IsFromInitialState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetIsFromInitialStateOk() (*bool, bool) {
	if o == nil || o.IsFromInitialState == nil {
		return nil, false
	}
	return o.IsFromInitialState, true
}

// HasIsFromInitialState returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasIsFromInitialState() bool {
	if o != nil && o.IsFromInitialState != nil {
		return true
	}

	return false
}

// SetIsFromInitialState gets a reference to the given bool and assigns it to the IsFromInitialState field.
func (o *BTReleaseItemMessageBody) SetIsFromInitialState(v bool) {
	o.IsFromInitialState = &v
}

// GetIsToTerminalState returns the IsToTerminalState field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetIsToTerminalState() bool {
	if o == nil || o.IsToTerminalState == nil {
		var ret bool
		return ret
	}
	return *o.IsToTerminalState
}

// GetIsToTerminalStateOk returns a tuple with the IsToTerminalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetIsToTerminalStateOk() (*bool, bool) {
	if o == nil || o.IsToTerminalState == nil {
		return nil, false
	}
	return o.IsToTerminalState, true
}

// HasIsToTerminalState returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasIsToTerminalState() bool {
	if o != nil && o.IsToTerminalState != nil {
		return true
	}

	return false
}

// SetIsToTerminalState gets a reference to the given bool and assigns it to the IsToTerminalState field.
func (o *BTReleaseItemMessageBody) SetIsToTerminalState(v bool) {
	o.IsToTerminalState = &v
}

// GetReleaseId returns the ReleaseId field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetReleaseId() string {
	if o == nil || o.ReleaseId == nil {
		var ret string
		return ret
	}
	return *o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetReleaseIdOk() (*string, bool) {
	if o == nil || o.ReleaseId == nil {
		return nil, false
	}
	return o.ReleaseId, true
}

// HasReleaseId returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasReleaseId() bool {
	if o != nil && o.ReleaseId != nil {
		return true
	}

	return false
}

// SetReleaseId gets a reference to the given string and assigns it to the ReleaseId field.
func (o *BTReleaseItemMessageBody) SetReleaseId(v string) {
	o.ReleaseId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BTReleaseItemMessageBody) SetStatus(v string) {
	o.Status = &v
}

// GetTransitionName returns the TransitionName field value if set, zero value otherwise.
func (o *BTReleaseItemMessageBody) GetTransitionName() string {
	if o == nil || o.TransitionName == nil {
		var ret string
		return ret
	}
	return *o.TransitionName
}

// GetTransitionNameOk returns a tuple with the TransitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemMessageBody) GetTransitionNameOk() (*string, bool) {
	if o == nil || o.TransitionName == nil {
		return nil, false
	}
	return o.TransitionName, true
}

// HasTransitionName returns a boolean if a field has been set.
func (o *BTReleaseItemMessageBody) HasTransitionName() bool {
	if o != nil && o.TransitionName != nil {
		return true
	}

	return false
}

// SetTransitionName gets a reference to the given string and assigns it to the TransitionName field.
func (o *BTReleaseItemMessageBody) SetTransitionName(v string) {
	o.TransitionName = &v
}

func (o BTReleaseItemMessageBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppElementSessionId != nil {
		toSerialize["appElementSessionId"] = o.AppElementSessionId
	}
	if o.CommentId != nil {
		toSerialize["commentId"] = o.CommentId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentState != nil {
		toSerialize["documentState"] = o.DocumentState
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.MetadataObjectType != nil {
		toSerialize["metadataObjectType"] = o.MetadataObjectType
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartIdentity != nil {
		toSerialize["partIdentity"] = o.PartIdentity
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TranslatationId != nil {
		toSerialize["translatationId"] = o.TranslatationId
	}
	if o.TranslationId != nil {
		toSerialize["translationId"] = o.TranslationId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WebhookId != nil {
		toSerialize["webhookId"] = o.WebhookId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.IsFromInitialState != nil {
		toSerialize["isFromInitialState"] = o.IsFromInitialState
	}
	if o.IsToTerminalState != nil {
		toSerialize["isToTerminalState"] = o.IsToTerminalState
	}
	if o.ReleaseId != nil {
		toSerialize["releaseId"] = o.ReleaseId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TransitionName != nil {
		toSerialize["transitionName"] = o.TransitionName
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleaseItemMessageBody struct {
	value *BTReleaseItemMessageBody
	isSet bool
}

func (v NullableBTReleaseItemMessageBody) Get() *BTReleaseItemMessageBody {
	return v.value
}

func (v *NullableBTReleaseItemMessageBody) Set(val *BTReleaseItemMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleaseItemMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleaseItemMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleaseItemMessageBody(val *BTReleaseItemMessageBody) *NullableBTReleaseItemMessageBody {
	return &NullableBTReleaseItemMessageBody{value: val, isSet: true}
}

func (v NullableBTReleaseItemMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleaseItemMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
