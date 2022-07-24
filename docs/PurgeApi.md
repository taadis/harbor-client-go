# \PurgeApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePurgeSchedule**](PurgeApi.md#CreatePurgeSchedule) | **Post** /system/purgeaudit/schedule | Create a purge job schedule.
[**GetPurgeHistory**](PurgeApi.md#GetPurgeHistory) | **Get** /system/purgeaudit | Get purge job results.
[**GetPurgeJob**](PurgeApi.md#GetPurgeJob) | **Get** /system/purgeaudit/{purge_id} | Get purge job status.
[**GetPurgeJobLog**](PurgeApi.md#GetPurgeJobLog) | **Get** /system/purgeaudit/{purge_id}/log | Get purge job log.
[**GetPurgeSchedule**](PurgeApi.md#GetPurgeSchedule) | **Get** /system/purgeaudit/schedule | Get purge&#39;s schedule.
[**StopPurge**](PurgeApi.md#StopPurge) | **Put** /system/purgeaudit/{purge_id} | Stop the specific purge audit log execution
[**UpdatePurgeSchedule**](PurgeApi.md#UpdatePurgeSchedule) | **Put** /system/purgeaudit/schedule | Update purge job&#39;s schedule.


# **CreatePurgeSchedule**
> CreatePurgeSchedule(ctx, schedule, optional)
Create a purge job schedule.

This endpoint is for update purge job schedule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**Schedule**](Schedule.md)| The purge job&#39;s schedule, it is a json object. ｜ The sample format is ｜ {\&quot;parameters\&quot;:{\&quot;audit_retention_hour\&quot;:168,\&quot;dry_run\&quot;:true, \&quot;include_operations\&quot;:\&quot;create,delete,pull\&quot;},\&quot;schedule\&quot;:{\&quot;type\&quot;:\&quot;Hourly\&quot;,\&quot;cron\&quot;:\&quot;0 0 * * * *\&quot;}} ｜ the include_operation should be a comma separated string, e.g. create,delete,pull, if it is empty, no operation will be purged.  | 
 **optional** | ***PurgeApiCreatePurgeScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiCreatePurgeScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurgeHistory**
> []ExecHistory GetPurgeHistory(ctx, optional)
Get purge job results.

get purge job execution history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PurgeApiGetPurgeHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiGetPurgeHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]

### Return type

[**[]ExecHistory**](ExecHistory.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurgeJob**
> ExecHistory GetPurgeJob(ctx, purgeId, optional)
Get purge job status.

This endpoint let user get purge job status filtered by specific ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purgeId** | **int64**| The ID of the purge log | 
 **optional** | ***PurgeApiGetPurgeJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiGetPurgeJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ExecHistory**](ExecHistory.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurgeJobLog**
> string GetPurgeJobLog(ctx, purgeId, optional)
Get purge job log.

This endpoint let user get purge job logs filtered by specific ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purgeId** | **int64**| The ID of the purge log | 
 **optional** | ***PurgeApiGetPurgeJobLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiGetPurgeJobLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurgeSchedule**
> ExecHistory GetPurgeSchedule(ctx, optional)
Get purge's schedule.

This endpoint is for get schedule of purge job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PurgeApiGetPurgeScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiGetPurgeScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ExecHistory**](ExecHistory.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopPurge**
> StopPurge(ctx, purgeId, optional)
Stop the specific purge audit log execution

Stop the purge audit log execution specified by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purgeId** | **int64**| The ID of the purge log | 
 **optional** | ***PurgeApiStopPurgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiStopPurgeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePurgeSchedule**
> UpdatePurgeSchedule(ctx, schedule, optional)
Update purge job's schedule.

This endpoint is for update purge job schedule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**Schedule**](Schedule.md)| The purge job&#39;s schedule, it is a json object. ｜ The sample format is ｜ {\&quot;parameters\&quot;:{\&quot;audit_retention_hour\&quot;:168,\&quot;dry_run\&quot;:true, \&quot;include_operations\&quot;:\&quot;create,delete,pull\&quot;},\&quot;schedule\&quot;:{\&quot;type\&quot;:\&quot;Hourly\&quot;,\&quot;cron\&quot;:\&quot;0 0 * * * *\&quot;}} ｜ the include_operation should be a comma separated string, e.g. create,delete,pull, if it is empty, no operation will be purged.  | 
 **optional** | ***PurgeApiUpdatePurgeScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurgeApiUpdatePurgeScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

