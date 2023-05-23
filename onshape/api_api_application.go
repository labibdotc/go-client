/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16301-d273853a12e7
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
	"reflect"
	"strings"
)

// APIApplicationApiService APIApplicationApi service
type APIApplicationApiService service

type ApiDeleteAppSettingsRequest struct {
	ctx        context.Context
	ApiService *APIApplicationApiService
	uid        string
	cid        string
	key        *[]string
}

func (r ApiDeleteAppSettingsRequest) Key(key []string) ApiDeleteAppSettingsRequest {
	r.key = &key
	return r
}

func (r ApiDeleteAppSettingsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAppSettingsExecute(r)
}

/*
DeleteAppSettings Delete application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @param cid
 @return ApiDeleteAppSettingsRequest
*/
func (a *APIApplicationApiService) DeleteAppSettings(ctx context.Context, uid string, cid string) ApiDeleteAppSettingsRequest {
	return ApiDeleteAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
		cid:        cid,
	}
}

// Execute executes the request
func (a *APIApplicationApiService) DeleteAppSettingsExecute(r ApiDeleteAppSettingsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.DeleteAppSettings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/clients/{cid}/settings/users/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		t := *r.key
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("key", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("key", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteCompanyAppSettingsRequest struct {
	ctx        context.Context
	ApiService *APIApplicationApiService
	cpid       string
	cid        string
	key        *[]string
}

func (r ApiDeleteCompanyAppSettingsRequest) Key(key []string) ApiDeleteCompanyAppSettingsRequest {
	r.key = &key
	return r
}

func (r ApiDeleteCompanyAppSettingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteCompanyAppSettingsExecute(r)
}

/*
DeleteCompanyAppSettings Delete company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cpid
 @param cid
 @return ApiDeleteCompanyAppSettingsRequest
*/
func (a *APIApplicationApiService) DeleteCompanyAppSettings(ctx context.Context, cpid string, cid string) ApiDeleteCompanyAppSettingsRequest {
	return ApiDeleteCompanyAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		cpid:       cpid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *APIApplicationApiService) DeleteCompanyAppSettingsExecute(r ApiDeleteCompanyAppSettingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.DeleteCompanyAppSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/clients/{cid}/settings/companies/{cpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cpid"+"}", url.PathEscape(parameterToString(r.cpid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		t := *r.key
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("key", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("key", parameterToString(t, "multi"))
		}
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

type ApiGetApplicableExtensionsForClientRequest struct {
	ctx            context.Context
	ApiService     *APIApplicationApiService
	uid            string
	cid            string
	validPurchases *bool
}

func (r ApiGetApplicableExtensionsForClientRequest) ValidPurchases(validPurchases bool) ApiGetApplicableExtensionsForClientRequest {
	r.validPurchases = &validPurchases
	return r
}

func (r ApiGetApplicableExtensionsForClientRequest) Execute() ([]BTAPIApplicationExtensionInfo, *http.Response, error) {
	return r.ApiService.GetApplicableExtensionsForClientExecute(r)
}

/*
GetApplicableExtensionsForClient Method for GetApplicableExtensionsForClient

Get list of client extensions the user has granted/accepted terms for

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @param cid
 @return ApiGetApplicableExtensionsForClientRequest
*/
func (a *APIApplicationApiService) GetApplicableExtensionsForClient(ctx context.Context, uid string, cid string) ApiGetApplicableExtensionsForClientRequest {
	return ApiGetApplicableExtensionsForClientRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return []BTAPIApplicationExtensionInfo
func (a *APIApplicationApiService) GetApplicableExtensionsForClientExecute(r ApiGetApplicableExtensionsForClientRequest) ([]BTAPIApplicationExtensionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTAPIApplicationExtensionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.GetApplicableExtensionsForClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/extensions/user/{uid}/client/{cid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.validPurchases != nil {
		localVarQueryParams.Add("validPurchases", parameterToString(*r.validPurchases, ""))
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
		var v []BTAPIApplicationExtensionInfo
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

type ApiGetCompanyAppSettingsRequest struct {
	ctx        context.Context
	ApiService *APIApplicationApiService
	cpid       string
	cid        string
	documentId *string
	key        *[]string
}

// A document owned by the company. Read access to this document allows read access to its owning company&#39;s settings.
func (r ApiGetCompanyAppSettingsRequest) DocumentId(documentId string) ApiGetCompanyAppSettingsRequest {
	r.documentId = &documentId
	return r
}

func (r ApiGetCompanyAppSettingsRequest) Key(key []string) ApiGetCompanyAppSettingsRequest {
	r.key = &key
	return r
}

func (r ApiGetCompanyAppSettingsRequest) Execute() (*BTUserAppSettingsInfo, *http.Response, error) {
	return r.ApiService.GetCompanyAppSettingsExecute(r)
}

/*
GetCompanyAppSettings Retrieve company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cpid
 @param cid
 @return ApiGetCompanyAppSettingsRequest
*/
func (a *APIApplicationApiService) GetCompanyAppSettings(ctx context.Context, cpid string, cid string) ApiGetCompanyAppSettingsRequest {
	return ApiGetCompanyAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		cpid:       cpid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return BTUserAppSettingsInfo
func (a *APIApplicationApiService) GetCompanyAppSettingsExecute(r ApiGetCompanyAppSettingsRequest) (*BTUserAppSettingsInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTUserAppSettingsInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.GetCompanyAppSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/clients/{cid}/settings/companies/{cpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cpid"+"}", url.PathEscape(parameterToString(r.cpid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.documentId != nil {
		localVarQueryParams.Add("documentId", parameterToString(*r.documentId, ""))
	}
	if r.key != nil {
		t := *r.key
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("key", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("key", parameterToString(t, "multi"))
		}
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
		var v BTUserAppSettingsInfo
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

type ApiGetUserAppSettingsRequest struct {
	ctx        context.Context
	ApiService *APIApplicationApiService
	uid        string
	cid        string
	key        *[]string
}

func (r ApiGetUserAppSettingsRequest) Key(key []string) ApiGetUserAppSettingsRequest {
	r.key = &key
	return r
}

func (r ApiGetUserAppSettingsRequest) Execute() (*BTUserAppSettingsInfo, *http.Response, error) {
	return r.ApiService.GetUserAppSettingsExecute(r)
}

/*
GetUserAppSettings Retrieve application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @param cid
 @return ApiGetUserAppSettingsRequest
*/
func (a *APIApplicationApiService) GetUserAppSettings(ctx context.Context, uid string, cid string) ApiGetUserAppSettingsRequest {
	return ApiGetUserAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return BTUserAppSettingsInfo
func (a *APIApplicationApiService) GetUserAppSettingsExecute(r ApiGetUserAppSettingsRequest) (*BTUserAppSettingsInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTUserAppSettingsInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.GetUserAppSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/clients/{cid}/settings/users/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		t := *r.key
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("key", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("key", parameterToString(t, "multi"))
		}
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
		var v BTUserAppSettingsInfo
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

type ApiUpdateAppCompanySettingsRequest struct {
	ctx                     context.Context
	ApiService              *APIApplicationApiService
	cpid                    string
	cid                     string
	bTUserAppSettingsParams *BTUserAppSettingsParams
}

func (r ApiUpdateAppCompanySettingsRequest) BTUserAppSettingsParams(bTUserAppSettingsParams BTUserAppSettingsParams) ApiUpdateAppCompanySettingsRequest {
	r.bTUserAppSettingsParams = &bTUserAppSettingsParams
	return r
}

func (r ApiUpdateAppCompanySettingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateAppCompanySettingsExecute(r)
}

/*
UpdateAppCompanySettings Update or create company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cpid
 @param cid
 @return ApiUpdateAppCompanySettingsRequest
*/
func (a *APIApplicationApiService) UpdateAppCompanySettings(ctx context.Context, cpid string, cid string) ApiUpdateAppCompanySettingsRequest {
	return ApiUpdateAppCompanySettingsRequest{
		ApiService: a,
		ctx:        ctx,
		cpid:       cpid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *APIApplicationApiService) UpdateAppCompanySettingsExecute(r ApiUpdateAppCompanySettingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.UpdateAppCompanySettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/clients/{cid}/settings/companies/{cpid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cpid"+"}", url.PathEscape(parameterToString(r.cpid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTUserAppSettingsParams == nil {
		return localVarReturnValue, nil, reportError("bTUserAppSettingsParams is required and must be specified")
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
	localVarPostBody = r.bTUserAppSettingsParams
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

type ApiUpdateAppSettingsRequest struct {
	ctx                     context.Context
	ApiService              *APIApplicationApiService
	uid                     string
	cid                     string
	bTUserAppSettingsParams *BTUserAppSettingsParams
}

func (r ApiUpdateAppSettingsRequest) BTUserAppSettingsParams(bTUserAppSettingsParams BTUserAppSettingsParams) ApiUpdateAppSettingsRequest {
	r.bTUserAppSettingsParams = &bTUserAppSettingsParams
	return r
}

func (r ApiUpdateAppSettingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateAppSettingsExecute(r)
}

/*
UpdateAppSettings Update or create application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @param cid
 @return ApiUpdateAppSettingsRequest
*/
func (a *APIApplicationApiService) UpdateAppSettings(ctx context.Context, uid string, cid string) ApiUpdateAppSettingsRequest {
	return ApiUpdateAppSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *APIApplicationApiService) UpdateAppSettingsExecute(r ApiUpdateAppSettingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIApplicationApiService.UpdateAppSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/clients/{cid}/settings/users/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTUserAppSettingsParams == nil {
		return localVarReturnValue, nil, reportError("bTUserAppSettingsParams is required and must be specified")
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
	localVarPostBody = r.bTUserAppSettingsParams
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
