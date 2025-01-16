// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1GetSupavisorConfigRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

func (o *V1GetSupavisorConfigRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

type V1GetSupavisorConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse              *http.Response
	SupavisorConfigResponses []shared.SupavisorConfigResponse
}

func (o *V1GetSupavisorConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1GetSupavisorConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1GetSupavisorConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1GetSupavisorConfigResponse) GetSupavisorConfigResponses() []shared.SupavisorConfigResponse {
	if o == nil {
		return nil
	}
	return o.SupavisorConfigResponses
}
