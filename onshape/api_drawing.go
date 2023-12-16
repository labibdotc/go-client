/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DrawingApiService DrawingApi service
type DrawingApiService service

type ApiCreateDrawingAppElementRequest struct {
	ctx             context.Context
	ApiService      *DrawingApiService
	did             string
	wid             string
	bTDrawingParams *BTDrawingParams
}

func (r ApiCreateDrawingAppElementRequest) BTDrawingParams(bTDrawingParams BTDrawingParams) ApiCreateDrawingAppElementRequest {
	r.bTDrawingParams = &bTDrawingParams
	return r
}

func (r ApiCreateDrawingAppElementRequest) Execute() (*BTDocumentElementInfo, *http.Response, error) {
	return r.ApiService.CreateDrawingAppElementExecute(r)
}

/*
CreateDrawingAppElement Create a new drawing in a document.

This endpoint takes a JSON Schema as input. See the schema docs below for details, and see [API Guide: Drawings](https://onshape-public.github.io/docs/api-adv/drawings/) for more information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did ID of the document in which to create the drawing.
 @param wid ID of the workspace in which to create the drawing.
 @return ApiCreateDrawingAppElementRequest
*/
func (a *DrawingApiService) CreateDrawingAppElement(ctx context.Context, did string, wid string) ApiCreateDrawingAppElementRequest {
	return ApiCreateDrawingAppElementRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
	}
}

// Execute executes the request
//  @return BTDocumentElementInfo
func (a *DrawingApiService) CreateDrawingAppElementExecute(r ApiCreateDrawingAppElementRequest) (*BTDocumentElementInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTDocumentElementInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.CreateDrawingAppElement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/d/{did}/w/{wid}/create"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTDrawingParams == nil {
		return localVarReturnValue, nil, reportError("bTDrawingParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTDrawingParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTDocumentElementInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateDrawingTranslationRequest struct {
	ctx                     context.Context
	ApiService              *DrawingApiService
	did                     string
	wv                      string
	wvid                    string
	eid                     string
	bTTranslateFormatParams *BTTranslateFormatParams
}

func (r ApiCreateDrawingTranslationRequest) BTTranslateFormatParams(bTTranslateFormatParams BTTranslateFormatParams) ApiCreateDrawingTranslationRequest {
	r.bTTranslateFormatParams = &bTTranslateFormatParams
	return r
}

func (r ApiCreateDrawingTranslationRequest) Execute() (*BTTranslationRequestInfo, *http.Response, error) {
	return r.ApiService.CreateDrawingTranslationExecute(r)
}

/*
CreateDrawingTranslation Translate (export) a drawing to a different format.

Export a drawing to a different format within a document. Use `getDrawingTranslatorFormats` for a list of supported translation (i.e., import/export) formats.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wv
 @param wvid
 @param eid
 @return ApiCreateDrawingTranslationRequest
*/
func (a *DrawingApiService) CreateDrawingTranslation(ctx context.Context, did string, wv string, wvid string, eid string) ApiCreateDrawingTranslationRequest {
	return ApiCreateDrawingTranslationRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTTranslationRequestInfo
func (a *DrawingApiService) CreateDrawingTranslationExecute(r ApiCreateDrawingTranslationRequest) (*BTTranslationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTTranslationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.CreateDrawingTranslation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/d/{did}/{wv}/{wvid}/e/{eid}/translations"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTTranslateFormatParams == nil {
		return localVarReturnValue, nil, reportError("bTTranslateFormatParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTTranslateFormatParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTTranslationRequestInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDrawingTranslatorFormatsRequest struct {
	ctx        context.Context
	ApiService *DrawingApiService
	did        string
	wid        string
	eid        string
}

func (r ApiGetDrawingTranslatorFormatsRequest) Execute() ([]BTModelFormatInfo, *http.Response, error) {
	return r.ApiService.GetDrawingTranslatorFormatsExecute(r)
}

/*
GetDrawingTranslatorFormats Get a list of all valid formats the drawing can be translated (exported) to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @param eid
 @return ApiGetDrawingTranslatorFormatsRequest
*/
func (a *DrawingApiService) GetDrawingTranslatorFormats(ctx context.Context, did string, wid string, eid string) ApiGetDrawingTranslatorFormatsRequest {
	return ApiGetDrawingTranslatorFormatsRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return []BTModelFormatInfo
func (a *DrawingApiService) GetDrawingTranslatorFormatsExecute(r ApiGetDrawingTranslatorFormatsRequest) ([]BTModelFormatInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTModelFormatInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.GetDrawingTranslatorFormats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/d/{did}/w/{wid}/e/{eid}/translationformats"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v []BTModelFormatInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetModificationStatusRequest struct {
	ctx        context.Context
	ApiService *DrawingApiService
	mrid       string
}

func (r ApiGetModificationStatusRequest) Execute() (*BTAppModificationRequestInfo, *http.Response, error) {
	return r.ApiService.GetModificationStatusExecute(r)
}

/*
GetModificationStatus Get the status of a drawing modification operation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mrid
 @return ApiGetModificationStatusRequest
*/
func (a *DrawingApiService) GetModificationStatus(ctx context.Context, mrid string) ApiGetModificationStatusRequest {
	return ApiGetModificationStatusRequest{
		ApiService: a,
		ctx:        ctx,
		mrid:       mrid,
	}
}

// Execute executes the request
//  @return BTAppModificationRequestInfo
func (a *DrawingApiService) GetModificationStatusExecute(r ApiGetModificationStatusRequest) (*BTAppModificationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAppModificationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.GetModificationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/modify/status/{mrid}"
	localVarPath = strings.Replace(localVarPath, "{"+"mrid"+"}", url.PathEscape(parameterToString(r.mrid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTAppModificationRequestInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiModifyDrawingRequest struct {
	ctx                         context.Context
	ApiService                  *DrawingApiService
	did                         string
	wid                         string
	eid                         string
	bTDrawingModificationParams *BTDrawingModificationParams
	linkDocumentId              *string
}

func (r ApiModifyDrawingRequest) BTDrawingModificationParams(bTDrawingModificationParams BTDrawingModificationParams) ApiModifyDrawingRequest {
	r.bTDrawingModificationParams = &bTDrawingModificationParams
	return r
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiModifyDrawingRequest) LinkDocumentId(linkDocumentId string) ApiModifyDrawingRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiModifyDrawingRequest) Execute() (*BTAppModificationRequestInfo, *http.Response, error) {
	return r.ApiService.ModifyDrawingExecute(r)
}

/*
ModifyDrawing Modify a drawing via JSON payload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wid The id of the workspace in which to perform the operation.
 @param eid The id of the element in which to perform the operation.
 @return ApiModifyDrawingRequest
*/
func (a *DrawingApiService) ModifyDrawing(ctx context.Context, did string, wid string, eid string) ApiModifyDrawingRequest {
	return ApiModifyDrawingRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTAppModificationRequestInfo
func (a *DrawingApiService) ModifyDrawingExecute(r ApiModifyDrawingRequest) (*BTAppModificationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAppModificationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.ModifyDrawing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/d/{did}/w/{wid}/e/{eid}/modify"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTDrawingModificationParams == nil {
		return localVarReturnValue, nil, reportError("bTDrawingModificationParams is required and must be specified")
	}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTDrawingModificationParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTAppModificationRequestInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
