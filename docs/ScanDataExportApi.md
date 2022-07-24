# \ScanDataExportApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadScanData**](ScanDataExportApi.md#DownloadScanData) | **Get** /export/cve/download/{execution_id} | Download the scan data export file
[**ExportScanData**](ScanDataExportApi.md#ExportScanData) | **Post** /export/cve | Export scan data for selected projects
[**GetScanDataExportExecution**](ScanDataExportApi.md#GetScanDataExportExecution) | **Get** /export/cve/execution/{execution_id} | Get the specific scan data export execution
[**GetScanDataExportExecutionList**](ScanDataExportApi.md#GetScanDataExportExecutionList) | **Get** /export/cve/executions | Get a list of specific scan data export execution jobs for a specified user


# **DownloadScanData**
> *os.File DownloadScanData(ctx, executionId, optional)
Download the scan data export file

Download the scan data report. Default format is CSV

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **int32**| Execution ID | 
 **optional** | ***ScanDataExportApiDownloadScanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanDataExportApiDownloadScanDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **format** | **optional.String**| The format of the data to be exported. e.g. CSV or PDF | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportScanData**
> ScanDataExportJob ExportScanData(ctx, xScanDataType, criteria, optional)
Export scan data for selected projects

Export scan data for selected projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xScanDataType** | **string**| The type of scan data to export | 
  **criteria** | [**ScanDataExportRequest**](ScanDataExportRequest.md)| The criteria for the export | 
 **optional** | ***ScanDataExportApiExportScanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanDataExportApiExportScanDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScanDataExportJob**](ScanDataExportJob.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScanDataExportExecution**
> ScanDataExportExecution GetScanDataExportExecution(ctx, executionId, optional)
Get the specific scan data export execution

Get the scan data export execution specified by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **int32**| Execution ID | 
 **optional** | ***ScanDataExportApiGetScanDataExportExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanDataExportApiGetScanDataExportExecutionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScanDataExportExecution**](ScanDataExportExecution.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScanDataExportExecutionList**
> ScanDataExportExecutionList GetScanDataExportExecutionList(ctx, optional)
Get a list of specific scan data export execution jobs for a specified user

Get the scan data export execution specified by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScanDataExportApiGetScanDataExportExecutionListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanDataExportApiGetScanDataExportExecutionListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScanDataExportExecutionList**](ScanDataExportExecutionList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

