# ScanDataExportExecution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the execution | [optional] [default to null]
**UserId** | **int32** | The ID if the user triggering the export job | [optional] [default to null]
**Status** | **string** | The status of the execution | [optional] [default to null]
**Trigger** | **string** | The trigger mode | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The start time | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | The end time | [optional] [default to null]
**StatusText** | **string** | The status text | [optional] [default to null]
**JobName** | **string** | The name of the job as specified by the user | [optional] [default to null]
**UserName** | **string** | The name of the user triggering the job | [optional] [default to null]
**FilePresent** | **bool** | Indicates whether the export artifact is present in registry | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


