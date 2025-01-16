// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1DeleteASsoProviderRequest struct {
	// Project ref
	Ref        string `pathParam:"style=simple,explode=false,name=ref"`
	ProviderID string `pathParam:"style=simple,explode=false,name=provider_id"`
}

func (o *V1DeleteASsoProviderRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1DeleteASsoProviderRequest) GetProviderID() string {
	if o == nil {
		return ""
	}
	return o.ProviderID
}

type V1DeleteASsoProviderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse            *http.Response
	DeleteProviderResponse *shared.DeleteProviderResponse
}

func (o *V1DeleteASsoProviderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1DeleteASsoProviderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1DeleteASsoProviderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1DeleteASsoProviderResponse) GetDeleteProviderResponse() *shared.DeleteProviderResponse {
	if o == nil {
		return nil
	}
	return o.DeleteProviderResponse
}