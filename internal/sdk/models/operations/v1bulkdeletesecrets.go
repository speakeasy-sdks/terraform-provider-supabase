// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type V1BulkDeleteSecretsRequest struct {
	// Project ref
	Ref         string   `pathParam:"style=simple,explode=false,name=ref"`
	RequestBody []string `request:"mediaType=application/json"`
}

func (o *V1BulkDeleteSecretsRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1BulkDeleteSecretsRequest) GetRequestBody() []string {
	if o == nil {
		return []string{}
	}
	return o.RequestBody
}

type V1BulkDeleteSecretsResponseBody struct {
}

type V1BulkDeleteSecretsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *V1BulkDeleteSecretsResponseBody
}

func (o *V1BulkDeleteSecretsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1BulkDeleteSecretsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1BulkDeleteSecretsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1BulkDeleteSecretsResponse) GetObject() *V1BulkDeleteSecretsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
