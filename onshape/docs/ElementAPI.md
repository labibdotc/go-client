# \ElementAPI

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyElementFromSourceDocument**](ElementAPI.md#CopyElementFromSourceDocument) | **Post** /elements/copyelement/{did}/workspace/{wid} | Copy tab by document ID and workspace ID.
[**DecodeConfiguration**](ElementAPI.md#DecodeConfiguration) | **Get** /elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configurationencodings/{cid} | Decode configuration string by documentation ID, workspace or version or microversion ID, and tab ID.
[**DeleteElement**](ElementAPI.md#DeleteElement) | **Delete** /elements/d/{did}/w/{wid}/e/{eid} | 
[**EncodeConfigurationMap**](ElementAPI.md#EncodeConfigurationMap) | **Post** /elements/d/{did}/e/{eid}/configurationencodings | Encode configuration by documentation ID and tab ID.
[**GetConfiguration**](ElementAPI.md#GetConfiguration) | **Get** /elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Retrieve configuration by document ID, workspace or version or microversion ID, and tab ID.
[**GetElementTranslatorFormatsByVersionOrWorkspace**](ElementAPI.md#GetElementTranslatorFormatsByVersionOrWorkspace) | **Get** /elements/translatorFormats/{did}/{wv}/{wvid}/{eid} | 
[**UpdateConfiguration**](ElementAPI.md#UpdateConfiguration) | **Post** /elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Update configuration by document ID, workspace or microversion ID, and tab ID.
[**UpdateReferences**](ElementAPI.md#UpdateReferences) | **Post** /elements/d/{did}/w/{wid}/e/{eid}/updatereferences | Update or replace node references by document ID, workspace ID, and tab ID.



## CopyElementFromSourceDocument

> BTDocumentElementInfo CopyElementFromSourceDocument(ctx, did, wid).BTCopyElementParams(bTCopyElementParams).Execute()

Copy tab by document ID and workspace ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    bTCopyElementParams := *openapiclient.NewBTCopyElementParams() // BTCopyElementParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.CopyElementFromSourceDocument(context.Background(), did, wid).BTCopyElementParams(bTCopyElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.CopyElementFromSourceDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyElementFromSourceDocument`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.CopyElementFromSourceDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyElementFromSourceDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTCopyElementParams** | [**BTCopyElementParams**](BTCopyElementParams.md) |  | 

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecodeConfiguration

> BTConfigurationInfo DecodeConfiguration(ctx, did, wvm, wvmid, eid, cid).LinkDocumentId(linkDocumentId).IncludeDisplay(includeDisplay).ConfigurationIsId(configurationIsId).Execute()

Decode configuration string by documentation ID, workspace or version or microversion ID, and tab ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    cid := "cid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    includeDisplay := true // bool |  (optional) (default to false)
    configurationIsId := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.DecodeConfiguration(context.Background(), did, wvm, wvmid, eid, cid).LinkDocumentId(linkDocumentId).IncludeDisplay(includeDisplay).ConfigurationIsId(configurationIsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.DecodeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeConfiguration`: BTConfigurationInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.DecodeConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecodeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** |  | 
 **includeDisplay** | **bool** |  | [default to false]
 **configurationIsId** | **bool** |  | [default to false]

### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteElement

> map[string]interface{} DeleteElement(ctx, did, wid, eid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.DeleteElement(context.Background(), did, wid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.DeleteElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteElement`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.DeleteElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncodeConfigurationMap

> BTEncodedConfigurationInfo EncodeConfigurationMap(ctx, did, eid).BTConfigurationParams(bTConfigurationParams).VersionId(versionId).LinkDocumentId(linkDocumentId).Execute()

Encode configuration by documentation ID and tab ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    bTConfigurationParams := *openapiclient.NewBTConfigurationParams() // BTConfigurationParams | 
    versionId := "versionId_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.EncodeConfigurationMap(context.Background(), did, eid).BTConfigurationParams(bTConfigurationParams).VersionId(versionId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.EncodeConfigurationMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncodeConfigurationMap`: BTEncodedConfigurationInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.EncodeConfigurationMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncodeConfigurationMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTConfigurationParams** | [**BTConfigurationParams**](BTConfigurationParams.md) |  | 
 **versionId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTEncodedConfigurationInfo**](BTEncodedConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> BTConfigurationResponse2019 GetConfiguration(ctx, did, wvm, wvmid, eid).Execute()

Retrieve configuration by document ID, workspace or version or microversion ID, and tab ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.GetConfiguration(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: BTConfigurationResponse2019
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTConfigurationResponse2019**](BTConfigurationResponse2019.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementTranslatorFormatsByVersionOrWorkspace

> []BTModelFormatInfo GetElementTranslatorFormatsByVersionOrWorkspace(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).CheckContent(checkContent).Configuration(configuration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    checkContent := true // bool |  (optional) (default to true)
    configuration := "configuration_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.GetElementTranslatorFormatsByVersionOrWorkspace(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).CheckContent(checkContent).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.GetElementTranslatorFormatsByVersionOrWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementTranslatorFormatsByVersionOrWorkspace`: []BTModelFormatInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.GetElementTranslatorFormatsByVersionOrWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementTranslatorFormatsByVersionOrWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **checkContent** | **bool** |  | [default to true]
 **configuration** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]BTModelFormatInfo**](BTModelFormatInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> BTConfigurationResponse2019 UpdateConfiguration(ctx, did, wvm, wvmid, eid).BTConfigurationUpdateCall2933(bTConfigurationUpdateCall2933).Execute()

Update configuration by document ID, workspace or microversion ID, and tab ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    bTConfigurationUpdateCall2933 := *openapiclient.NewBTConfigurationUpdateCall2933() // BTConfigurationUpdateCall2933 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.UpdateConfiguration(context.Background(), did, wvm, wvmid, eid).BTConfigurationUpdateCall2933(bTConfigurationUpdateCall2933).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: BTConfigurationResponse2019
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTConfigurationUpdateCall2933** | [**BTConfigurationUpdateCall2933**](BTConfigurationUpdateCall2933.md) |  | 

### Return type

[**BTConfigurationResponse2019**](BTConfigurationResponse2019.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReferences

> map[string]interface{} UpdateReferences(ctx, did, wid, eid).BTUpdateReferenceParams(bTUpdateReferenceParams).Execute()

Update or replace node references by document ID, workspace ID, and tab ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    eid := "eid_example" // string | 
    bTUpdateReferenceParams := *openapiclient.NewBTUpdateReferenceParams() // BTUpdateReferenceParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementAPI.UpdateReferences(context.Background(), did, wid, eid).BTUpdateReferenceParams(bTUpdateReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementAPI.UpdateReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReferences`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ElementAPI.UpdateReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateReferenceParams** | [**BTUpdateReferenceParams**](BTUpdateReferenceParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
