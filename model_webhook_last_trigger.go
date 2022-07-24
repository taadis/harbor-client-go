/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// The webhook policy and last trigger time group by event type.
type WebhookLastTrigger struct {
	// The webhook policy name.
	PolicyName string `json:"policy_name,omitempty"`
	// The webhook event type.
	EventType string `json:"event_type,omitempty"`
	// Whether or not the webhook policy enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The creation time of webhook policy.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The last trigger time of webhook policy.
	LastTriggerTime time.Time `json:"last_trigger_time,omitempty"`
}
