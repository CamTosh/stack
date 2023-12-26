// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ChangeConfigSecretRequest struct {
	ConfigChangeSecret *shared.ConfigChangeSecret `request:"mediaType=application/json"`
	// Config ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ChangeConfigSecretRequest) GetConfigChangeSecret() *shared.ConfigChangeSecret {
	if o == nil {
		return nil
	}
	return o.ConfigChangeSecret
}

func (o *ChangeConfigSecretRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ChangeConfigSecretResponse struct {
	// Secret successfully changed.
	ConfigResponse *shared.ConfigResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WebhooksErrorResponse *shared.WebhooksErrorResponse
}

func (o *ChangeConfigSecretResponse) GetConfigResponse() *shared.ConfigResponse {
	if o == nil {
		return nil
	}
	return o.ConfigResponse
}

func (o *ChangeConfigSecretResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ChangeConfigSecretResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ChangeConfigSecretResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ChangeConfigSecretResponse) GetWebhooksErrorResponse() *shared.WebhooksErrorResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksErrorResponse
}