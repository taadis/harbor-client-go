/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Metrics struct {
	// The count of task
	TaskCount int32 `json:"task_count,omitempty"`
	// The count of success task
	SuccessTaskCount int32 `json:"success_task_count,omitempty"`
	// The count of error task
	ErrorTaskCount int32 `json:"error_task_count,omitempty"`
	// The count of pending task
	PendingTaskCount int32 `json:"pending_task_count,omitempty"`
	// The count of running task
	RunningTaskCount int32 `json:"running_task_count,omitempty"`
	// The count of scheduled task
	ScheduledTaskCount int32 `json:"scheduled_task_count,omitempty"`
	// The count of stopped task
	StoppedTaskCount int32 `json:"stopped_task_count,omitempty"`
}
