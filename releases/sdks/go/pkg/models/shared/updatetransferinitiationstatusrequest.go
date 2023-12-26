// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UpdateTransferInitiationStatusRequestStatus string

const (
	UpdateTransferInitiationStatusRequestStatusWaitingForValidation UpdateTransferInitiationStatusRequestStatus = "WAITING_FOR_VALIDATION"
	UpdateTransferInitiationStatusRequestStatusProcessing           UpdateTransferInitiationStatusRequestStatus = "PROCESSING"
	UpdateTransferInitiationStatusRequestStatusProcessed            UpdateTransferInitiationStatusRequestStatus = "PROCESSED"
	UpdateTransferInitiationStatusRequestStatusFailed               UpdateTransferInitiationStatusRequestStatus = "FAILED"
	UpdateTransferInitiationStatusRequestStatusRejected             UpdateTransferInitiationStatusRequestStatus = "REJECTED"
	UpdateTransferInitiationStatusRequestStatusValidated            UpdateTransferInitiationStatusRequestStatus = "VALIDATED"
)

func (e UpdateTransferInitiationStatusRequestStatus) ToPointer() *UpdateTransferInitiationStatusRequestStatus {
	return &e
}

func (e *UpdateTransferInitiationStatusRequestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WAITING_FOR_VALIDATION":
		fallthrough
	case "PROCESSING":
		fallthrough
	case "PROCESSED":
		fallthrough
	case "FAILED":
		fallthrough
	case "REJECTED":
		fallthrough
	case "VALIDATED":
		*e = UpdateTransferInitiationStatusRequestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateTransferInitiationStatusRequestStatus: %v", v)
	}
}

type UpdateTransferInitiationStatusRequest struct {
	Status UpdateTransferInitiationStatusRequestStatus `json:"status"`
}

func (o *UpdateTransferInitiationStatusRequest) GetStatus() UpdateTransferInitiationStatusRequestStatus {
	if o == nil {
		return UpdateTransferInitiationStatusRequestStatus("")
	}
	return o.Status
}