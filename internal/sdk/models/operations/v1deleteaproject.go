// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1DeleteAProjectRequest struct {
	// Project ref
	ID string `pathParam:"style=simple,explode=false,name=ref"`
}

func (o *V1DeleteAProjectRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type V1DeleteAProjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse          *http.Response
	V1ProjectRefResponse *shared.V1ProjectRefResponse
}

func (o *V1DeleteAProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1DeleteAProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1DeleteAProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1DeleteAProjectResponse) GetV1ProjectRefResponse() *shared.V1ProjectRefResponse {
	if o == nil {
		return nil
	}
	return o.V1ProjectRefResponse
}
