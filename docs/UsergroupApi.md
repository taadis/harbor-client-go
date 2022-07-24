# \UsergroupApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserGroup**](UsergroupApi.md#CreateUserGroup) | **Post** /usergroups | Create user group
[**DeleteUserGroup**](UsergroupApi.md#DeleteUserGroup) | **Delete** /usergroups/{group_id} | Delete user group
[**GetUserGroup**](UsergroupApi.md#GetUserGroup) | **Get** /usergroups/{group_id} | Get user group information
[**ListUserGroups**](UsergroupApi.md#ListUserGroups) | **Get** /usergroups | Get all user groups information
[**SearchUserGroups**](UsergroupApi.md#SearchUserGroups) | **Get** /usergroups/search | Search groups by groupname
[**UpdateUserGroup**](UsergroupApi.md#UpdateUserGroup) | **Put** /usergroups/{group_id} | Update group information


# **CreateUserGroup**
> CreateUserGroup(ctx, optional)
Create user group

Create user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupApiCreateUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupApiCreateUserGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **usergroup** | [**optional.Interface of UserGroup**](UserGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserGroup**
> DeleteUserGroup(ctx, groupId, optional)
Delete user group

Delete user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int32**|  | 
 **optional** | ***UsergroupApiDeleteUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupApiDeleteUserGroupOpts struct

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

# **GetUserGroup**
> UserGroup GetUserGroup(ctx, groupId, optional)
Get user group information

Get user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***UsergroupApiGetUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupApiGetUserGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserGroups**
> []UserGroup ListUserGroups(ctx, optional)
Get all user groups information

Get all user groups information, it is open for system admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupApiListUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupApiListUserGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **ldapGroupDn** | **optional.String**| search with ldap group DN | 
 **groupName** | **optional.String**| group name need to search, fuzzy matches | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUserGroups**
> []UserGroupSearchItem SearchUserGroups(ctx, groupname, optional)
Search groups by groupname

This endpoint is to search groups by group name.  It's open for all authenticated requests. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupname** | **string**| Group name for filtering results. | 
 **optional** | ***UsergroupApiSearchUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupApiSearchUserGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]

### Return type

[**[]UserGroupSearchItem**](UserGroupSearchItem.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserGroup**
> UpdateUserGroup(ctx, groupId, optional)
Update group information

Update user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***UsergroupApiUpdateUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupApiUpdateUserGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **usergroup** | [**optional.Interface of UserGroup**](UserGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

