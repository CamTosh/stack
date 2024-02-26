// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"time"
)

type TriggerOccurrence struct {
	Date               time.Time              `json:"date"`
	Error              *string                `json:"error,omitempty"`
	Event              map[string]interface{} `json:"event"`
	TriggerID          string                 `json:"triggerID"`
	WorkflowInstance   *WorkflowInstance      `json:"workflowInstance,omitempty"`
	WorkflowInstanceID *string                `json:"workflowInstanceID,omitempty"`
}

func (t TriggerOccurrence) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TriggerOccurrence) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TriggerOccurrence) GetDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Date
}

func (o *TriggerOccurrence) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TriggerOccurrence) GetEvent() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Event
}

func (o *TriggerOccurrence) GetTriggerID() string {
	if o == nil {
		return ""
	}
	return o.TriggerID
}

func (o *TriggerOccurrence) GetWorkflowInstance() *WorkflowInstance {
	if o == nil {
		return nil
	}
	return o.WorkflowInstance
}

func (o *TriggerOccurrence) GetWorkflowInstanceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowInstanceID
}
