/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7802-6b1a0d70e58f
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

// VariablesApiService VariablesApi service
type VariablesApiService service

type ApiCreateVariableStudioRequest struct {
	ctx                  context.Context
	ApiService           *VariablesApiService
	did                  string
	wid                  string
	bTModelElementParams *BTModelElementParams
	linkDocumentId       *string
}

func (r ApiCreateVariableStudioRequest) BTModelElementParams(bTModelElementParams BTModelElementParams) ApiCreateVariableStudioRequest {
	r.bTModelElementParams = &bTModelElementParams
	return r
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiCreateVariableStudioRequest) LinkDocumentId(linkDocumentId string) ApiCreateVariableStudioRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiCreateVariableStudioRequest) Execute() (*BTDocumentElementInfo, *http.Response, error) {
	return r.ApiService.CreateVariableStudioExecute(r)
}

/*
CreateVariableStudio Create a variable studio

Create a Variable studio

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wid The id of the workspace in which to perform the operation.
 @return ApiCreateVariableStudioRequest
*/
func (a *VariablesApiService) CreateVariableStudio(ctx context.Context, did string, wid string) ApiCreateVariableStudioRequest {
	return ApiCreateVariableStudioRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
	}
}

// Execute executes the request
//  @return BTDocumentElementInfo
func (a *VariablesApiService) CreateVariableStudioExecute(r ApiCreateVariableStudioRequest) (*BTDocumentElementInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTDocumentElementInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.CreateVariableStudio")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/w/{wid}/variablestudio"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTModelElementParams == nil {
		return localVarReturnValue, nil, reportError("bTModelElementParams is required and must be specified")
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
	localVarPostBody = r.bTModelElementParams
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

type ApiGetVariableStudioReferencesRequest struct {
	ctx            context.Context
	ApiService     *VariablesApiService
	did            string
	wv             string
	wvid           string
	eid            string
	linkDocumentId *string
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiGetVariableStudioReferencesRequest) LinkDocumentId(linkDocumentId string) ApiGetVariableStudioReferencesRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetVariableStudioReferencesRequest) Execute() (*BTVariableStudioReferenceListInfo, *http.Response, error) {
	return r.ApiService.GetVariableStudioReferencesExecute(r)
}

/*
GetVariableStudioReferences Retrieve the variable studio references from an element

Get an element's variable studio references

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wv
 @param wvid
 @param eid The id of the element in which to perform the operation.
 @return ApiGetVariableStudioReferencesRequest
*/
func (a *VariablesApiService) GetVariableStudioReferences(ctx context.Context, did string, wv string, wvid string, eid string) ApiGetVariableStudioReferencesRequest {
	return ApiGetVariableStudioReferencesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTVariableStudioReferenceListInfo
func (a *VariablesApiService) GetVariableStudioReferencesExecute(r ApiGetVariableStudioReferencesRequest) (*BTVariableStudioReferenceListInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTVariableStudioReferenceListInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.GetVariableStudioReferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/{wv}/{wvid}/e/{eid}/variablestudioreferences"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
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
		var v BTVariableStudioReferenceListInfo
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

type ApiGetVariableStudioScopeRequest struct {
	ctx            context.Context
	ApiService     *VariablesApiService
	did            string
	wv             string
	wvid           string
	eid            string
	linkDocumentId *string
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiGetVariableStudioScopeRequest) LinkDocumentId(linkDocumentId string) ApiGetVariableStudioScopeRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetVariableStudioScopeRequest) Execute() (*BTVariableStudioScopeInfo, *http.Response, error) {
	return r.ApiService.GetVariableStudioScopeExecute(r)
}

/*
GetVariableStudioScope Get the scope of a variable studio

Get the scope of a variable studio

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wv
 @param wvid
 @param eid The id of the element in which to perform the operation.
 @return ApiGetVariableStudioScopeRequest
*/
func (a *VariablesApiService) GetVariableStudioScope(ctx context.Context, did string, wv string, wvid string, eid string) ApiGetVariableStudioScopeRequest {
	return ApiGetVariableStudioScopeRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTVariableStudioScopeInfo
func (a *VariablesApiService) GetVariableStudioScopeExecute(r ApiGetVariableStudioScopeRequest) (*BTVariableStudioScopeInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTVariableStudioScopeInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.GetVariableStudioScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/{wv}/{wvid}/e/{eid}/variablestudioscope"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
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
		var v BTVariableStudioScopeInfo
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

type ApiGetVariablesRequest struct {
	ctx                                 context.Context
	ApiService                          *VariablesApiService
	did                                 string
	wv                                  string
	wvid                                string
	eid                                 string
	linkDocumentId                      *string
	includeValuesAndReferencedVariables *bool
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiGetVariablesRequest) LinkDocumentId(linkDocumentId string) ApiGetVariablesRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetVariablesRequest) IncludeValuesAndReferencedVariables(includeValuesAndReferencedVariables bool) ApiGetVariablesRequest {
	r.includeValuesAndReferencedVariables = &includeValuesAndReferencedVariables
	return r
}

func (r ApiGetVariablesRequest) Execute() (*BTVariableTableInfo, *http.Response, error) {
	return r.ApiService.GetVariablesExecute(r)
}

/*
GetVariables Retrieve the variables from a variable table

Gets contents of variable tables

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wv
 @param wvid
 @param eid The id of the element in which to perform the operation.
 @return ApiGetVariablesRequest
*/
func (a *VariablesApiService) GetVariables(ctx context.Context, did string, wv string, wvid string, eid string) ApiGetVariablesRequest {
	return ApiGetVariablesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTVariableTableInfo
func (a *VariablesApiService) GetVariablesExecute(r ApiGetVariablesRequest) (*BTVariableTableInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTVariableTableInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.GetVariables")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/{wv}/{wvid}/e/{eid}/variables"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	if r.includeValuesAndReferencedVariables != nil {
		localVarQueryParams.Add("includeValuesAndReferencedVariables", parameterToString(*r.includeValuesAndReferencedVariables, ""))
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
		var v BTVariableTableInfo
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

type ApiSetVariableStudioReferencesRequest struct {
	ctx                               context.Context
	ApiService                        *VariablesApiService
	did                               string
	wid                               string
	eid                               string
	bTVariableStudioReferenceListInfo *BTVariableStudioReferenceListInfo
	linkDocumentId                    *string
}

func (r ApiSetVariableStudioReferencesRequest) BTVariableStudioReferenceListInfo(bTVariableStudioReferenceListInfo BTVariableStudioReferenceListInfo) ApiSetVariableStudioReferencesRequest {
	r.bTVariableStudioReferenceListInfo = &bTVariableStudioReferenceListInfo
	return r
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiSetVariableStudioReferencesRequest) LinkDocumentId(linkDocumentId string) ApiSetVariableStudioReferencesRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiSetVariableStudioReferencesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SetVariableStudioReferencesExecute(r)
}

/*
SetVariableStudioReferences Set the variable studio references for an element

Set an element's variable studio references

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wid The id of the workspace in which to perform the operation.
 @param eid
 @return ApiSetVariableStudioReferencesRequest
*/
func (a *VariablesApiService) SetVariableStudioReferences(ctx context.Context, did string, wid string, eid string) ApiSetVariableStudioReferencesRequest {
	return ApiSetVariableStudioReferencesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *VariablesApiService) SetVariableStudioReferencesExecute(r ApiSetVariableStudioReferencesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.SetVariableStudioReferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/w/{wid}/e/{eid}/variablestudioreferences"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTVariableStudioReferenceListInfo == nil {
		return localVarReturnValue, nil, reportError("bTVariableStudioReferenceListInfo is required and must be specified")
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
	localVarPostBody = r.bTVariableStudioReferenceListInfo
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

type ApiSetVariableStudioScopeRequest struct {
	ctx                       context.Context
	ApiService                *VariablesApiService
	did                       string
	wid                       string
	eid                       string
	bTVariableStudioScopeInfo *BTVariableStudioScopeInfo
	linkDocumentId            *string
}

func (r ApiSetVariableStudioScopeRequest) BTVariableStudioScopeInfo(bTVariableStudioScopeInfo BTVariableStudioScopeInfo) ApiSetVariableStudioScopeRequest {
	r.bTVariableStudioScopeInfo = &bTVariableStudioScopeInfo
	return r
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiSetVariableStudioScopeRequest) LinkDocumentId(linkDocumentId string) ApiSetVariableStudioScopeRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiSetVariableStudioScopeRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SetVariableStudioScopeExecute(r)
}

/*
SetVariableStudioScope Set the scope of a variable studio

Set the scope of a variable studio

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wid The id of the workspace in which to perform the operation.
 @param eid
 @return ApiSetVariableStudioScopeRequest
*/
func (a *VariablesApiService) SetVariableStudioScope(ctx context.Context, did string, wid string, eid string) ApiSetVariableStudioScopeRequest {
	return ApiSetVariableStudioScopeRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *VariablesApiService) SetVariableStudioScopeExecute(r ApiSetVariableStudioScopeRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.SetVariableStudioScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/w/{wid}/e/{eid}/variablestudioscope"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTVariableStudioScopeInfo == nil {
		return localVarReturnValue, nil, reportError("bTVariableStudioScopeInfo is required and must be specified")
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
	localVarPostBody = r.bTVariableStudioScopeInfo
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

type ApiSetVariablesRequest struct {
	ctx              context.Context
	ApiService       *VariablesApiService
	did              string
	wid              string
	eid              string
	bTVariableParams *[]BTVariableParams
	linkDocumentId   *string
}

func (r ApiSetVariablesRequest) BTVariableParams(bTVariableParams []BTVariableParams) ApiSetVariablesRequest {
	r.bTVariableParams = &bTVariableParams
	return r
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiSetVariablesRequest) LinkDocumentId(linkDocumentId string) ApiSetVariablesRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiSetVariablesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SetVariablesExecute(r)
}

/*
SetVariables Assign variables to a variable studio

Assign variables to a variable studio

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wid The id of the workspace in which to perform the operation.
 @param eid
 @return ApiSetVariablesRequest
*/
func (a *VariablesApiService) SetVariables(ctx context.Context, did string, wid string, eid string) ApiSetVariablesRequest {
	return ApiSetVariablesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *VariablesApiService) SetVariablesExecute(r ApiSetVariablesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.SetVariables")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/d/{did}/w/{wid}/e/{eid}/variables"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTVariableParams == nil {
		return localVarReturnValue, nil, reportError("bTVariableParams is required and must be specified")
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
	localVarPostBody = r.bTVariableParams
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
