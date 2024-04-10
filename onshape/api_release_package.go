/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

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

// ReleasePackageApiService ReleasePackageApi service
type ReleasePackageApiService service

type ApiCreateObsoletionPackageRequest struct {
	ctx        context.Context
	ApiService *ReleasePackageApiService
	wfid       string
	revisionId *string
	debugMode  *bool
}

func (r ApiCreateObsoletionPackageRequest) RevisionId(revisionId string) ApiCreateObsoletionPackageRequest {
	r.revisionId = &revisionId
	return r
}

func (r ApiCreateObsoletionPackageRequest) DebugMode(debugMode bool) ApiCreateObsoletionPackageRequest {
	r.debugMode = &debugMode
	return r
}

func (r ApiCreateObsoletionPackageRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateObsoletionPackageExecute(r)
}

/*
CreateObsoletionPackage Create an obsoletion package to make an existing revision obsolete.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wfid
	@return ApiCreateObsoletionPackageRequest
*/
func (a *ReleasePackageApiService) CreateObsoletionPackage(ctx context.Context, wfid string) ApiCreateObsoletionPackageRequest {
	return ApiCreateObsoletionPackageRequest{
		ApiService: a,
		ctx:        ctx,
		wfid:       wfid,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ReleasePackageApiService) CreateObsoletionPackageExecute(r ApiCreateObsoletionPackageRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleasePackageApiService.CreateObsoletionPackage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/releasepackages/obsoletion/{wfid}"
	localVarPath = strings.Replace(localVarPath, "{"+"wfid"+"}", url.PathEscape(parameterToString(r.wfid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revisionId == nil {
		return localVarReturnValue, nil, reportError("revisionId is required and must be specified")
	}

	localVarQueryParams.Add("revisionId", parameterToString(*r.revisionId, ""))
	if r.debugMode != nil {
		localVarQueryParams.Add("debugMode", parameterToString(*r.debugMode, ""))
	}
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
		var v map[string]interface{}
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

type ApiCreateReleasePackageRequest struct {
	ctx                    context.Context
	ApiService             *ReleasePackageApiService
	wfid                   string
	bTReleasePackageParams *BTReleasePackageParams
	debugMode              *bool
}

func (r ApiCreateReleasePackageRequest) BTReleasePackageParams(bTReleasePackageParams BTReleasePackageParams) ApiCreateReleasePackageRequest {
	r.bTReleasePackageParams = &bTReleasePackageParams
	return r
}

func (r ApiCreateReleasePackageRequest) DebugMode(debugMode bool) ApiCreateReleasePackageRequest {
	r.debugMode = &debugMode
	return r
}

func (r ApiCreateReleasePackageRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateReleasePackageExecute(r)
}

/*
CreateReleasePackage Create a new release package for one or more items.

All revisionable items must be from the same document. Once a release package is successfully created, use `updateReleasePackage` to update all desired item/package properties, and transition it to the desired state.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wfid
	@return ApiCreateReleasePackageRequest
*/
func (a *ReleasePackageApiService) CreateReleasePackage(ctx context.Context, wfid string) ApiCreateReleasePackageRequest {
	return ApiCreateReleasePackageRequest{
		ApiService: a,
		ctx:        ctx,
		wfid:       wfid,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ReleasePackageApiService) CreateReleasePackageExecute(r ApiCreateReleasePackageRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleasePackageApiService.CreateReleasePackage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/releasepackages/release/{wfid}"
	localVarPath = strings.Replace(localVarPath, "{"+"wfid"+"}", url.PathEscape(parameterToString(r.wfid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTReleasePackageParams == nil {
		return localVarReturnValue, nil, reportError("bTReleasePackageParams is required and must be specified")
	}

	if r.debugMode != nil {
		localVarQueryParams.Add("debugMode", parameterToString(*r.debugMode, ""))
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
	localVarPostBody = r.bTReleasePackageParams
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
		var v map[string]interface{}
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

type ApiGetCompanyReleaseWorkflowRequest struct {
	ctx        context.Context
	ApiService *ReleasePackageApiService
	documentId *string
}

func (r ApiGetCompanyReleaseWorkflowRequest) DocumentId(documentId string) ApiGetCompanyReleaseWorkflowRequest {
	r.documentId = &documentId
	return r
}

func (r ApiGetCompanyReleaseWorkflowRequest) Execute() (*BTActiveWorkflowInfo, *http.Response, error) {
	return r.ApiService.GetCompanyReleaseWorkflowExecute(r)
}

/*
GetCompanyReleaseWorkflow Get information about the release/obsoletion workflow for a company-owned document.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCompanyReleaseWorkflowRequest
*/
func (a *ReleasePackageApiService) GetCompanyReleaseWorkflow(ctx context.Context) ApiGetCompanyReleaseWorkflowRequest {
	return ApiGetCompanyReleaseWorkflowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BTActiveWorkflowInfo
func (a *ReleasePackageApiService) GetCompanyReleaseWorkflowExecute(r ApiGetCompanyReleaseWorkflowRequest) (*BTActiveWorkflowInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTActiveWorkflowInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleasePackageApiService.GetCompanyReleaseWorkflow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/releasepackages/companyreleaseworkflow"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.documentId != nil {
		localVarQueryParams.Add("documentId", parameterToString(*r.documentId, ""))
	}
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
		var v BTActiveWorkflowInfo
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

type ApiGetReleasePackageRequest struct {
	ctx        context.Context
	ApiService *ReleasePackageApiService
	rpid       string
	detailed   *bool
}

func (r ApiGetReleasePackageRequest) Detailed(detailed bool) ApiGetReleasePackageRequest {
	r.detailed = &detailed
	return r
}

func (r ApiGetReleasePackageRequest) Execute() (*BTReleasePackageInfo, *http.Response, error) {
	return r.ApiService.GetReleasePackageExecute(r)
}

/*
GetReleasePackage Get details about the specified release package.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param rpid
	@return ApiGetReleasePackageRequest
*/
func (a *ReleasePackageApiService) GetReleasePackage(ctx context.Context, rpid string) ApiGetReleasePackageRequest {
	return ApiGetReleasePackageRequest{
		ApiService: a,
		ctx:        ctx,
		rpid:       rpid,
	}
}

// Execute executes the request
//
//	@return BTReleasePackageInfo
func (a *ReleasePackageApiService) GetReleasePackageExecute(r ApiGetReleasePackageRequest) (*BTReleasePackageInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTReleasePackageInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleasePackageApiService.GetReleasePackage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/releasepackages/{rpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpid"+"}", url.PathEscape(parameterToString(r.rpid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.detailed != nil {
		localVarQueryParams.Add("detailed", parameterToString(*r.detailed, ""))
	}
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
		var v BTReleasePackageInfo
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

type ApiUpdateReleasePackageRequest struct {
	ctx                          context.Context
	ApiService                   *ReleasePackageApiService
	rpid                         string
	bTUpdateReleasePackageParams *BTUpdateReleasePackageParams
	action                       *string
	wfaction                     *string
}

func (r ApiUpdateReleasePackageRequest) BTUpdateReleasePackageParams(bTUpdateReleasePackageParams BTUpdateReleasePackageParams) ApiUpdateReleasePackageRequest {
	r.bTUpdateReleasePackageParams = &bTUpdateReleasePackageParams
	return r
}

func (r ApiUpdateReleasePackageRequest) Action(action string) ApiUpdateReleasePackageRequest {
	r.action = &action
	return r
}

func (r ApiUpdateReleasePackageRequest) Wfaction(wfaction string) ApiUpdateReleasePackageRequest {
	r.wfaction = &wfaction
	return r
}

func (r ApiUpdateReleasePackageRequest) Execute() (*BTReleasePackageInfo, *http.Response, error) {
	return r.ApiService.UpdateReleasePackageExecute(r)
}

/*
UpdateReleasePackage Update the release/obsoletion package/item properties.

Use the `wfaction` query param to also perform a workflow transition.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param rpid
	@return ApiUpdateReleasePackageRequest
*/
func (a *ReleasePackageApiService) UpdateReleasePackage(ctx context.Context, rpid string) ApiUpdateReleasePackageRequest {
	return ApiUpdateReleasePackageRequest{
		ApiService: a,
		ctx:        ctx,
		rpid:       rpid,
	}
}

// Execute executes the request
//
//	@return BTReleasePackageInfo
func (a *ReleasePackageApiService) UpdateReleasePackageExecute(r ApiUpdateReleasePackageRequest) (*BTReleasePackageInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTReleasePackageInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleasePackageApiService.UpdateReleasePackage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/releasepackages/{rpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpid"+"}", url.PathEscape(parameterToString(r.rpid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTUpdateReleasePackageParams == nil {
		return localVarReturnValue, nil, reportError("bTUpdateReleasePackageParams is required and must be specified")
	}

	if r.action != nil {
		localVarQueryParams.Add("action", parameterToString(*r.action, ""))
	}
	if r.wfaction != nil {
		localVarQueryParams.Add("wfaction", parameterToString(*r.wfaction, ""))
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
	localVarPostBody = r.bTUpdateReleasePackageParams
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
		var v BTReleasePackageInfo
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
