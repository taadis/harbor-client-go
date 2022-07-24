# ScanDataExportRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobName** | **string** | Name of the scan data export job | [optional] [default to null]
**Projects** | **[]int64** | A list of one or more projects for which to export the scan data, defaults to all if empty | [optional] [default to null]
**Labels** | **[]int64** | A list of one or more labels for which to export the scan data, defaults to all if empty | [optional] [default to null]
**Repositories** | **string** | A list of repositories for which to export the scan data, defaults to all if empty | [optional] [default to null]
**CveIds** | **string** | CVE-IDs for which to export data. Multiple CVE-IDs can be specified by separating using &#39;,&#39; and enclosed between &#39;{}&#39;. Defaults to all if empty | [optional] [default to null]
**Tags** | **string** | A list of tags enclosed within &#39;{}&#39;. Defaults to all if empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


