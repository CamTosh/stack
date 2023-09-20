// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type AddMetadataOnTransactionRequest struct {
	// Use an idempotency key
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
	// metadata
	RequestBody map[string]string `request:"mediaType=application/json"`
	// Set the dryRun mode. Dry run mode doesn't add the logs to the database or publish a message to the message broker.
	DryRun *bool `queryParam:"style=form,explode=true,name=dryRun"`
	// Transaction ID.
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *AddMetadataOnTransactionRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *AddMetadataOnTransactionRequest) GetRequestBody() map[string]string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddMetadataOnTransactionRequest) GetDryRun() *bool {
	if o == nil {
		return nil
	}
	return o.DryRun
}

func (o *AddMetadataOnTransactionRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *AddMetadataOnTransactionRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type AddMetadataOnTransactionResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}

func (o *AddMetadataOnTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddMetadataOnTransactionResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *AddMetadataOnTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddMetadataOnTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
