// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type V1EnableDatabaseWebhookRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

func (o *V1EnableDatabaseWebhookRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

type V1EnableDatabaseWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V1EnableDatabaseWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1EnableDatabaseWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1EnableDatabaseWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
