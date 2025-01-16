// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1ListAllSsoProviderRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

func (o *V1ListAllSsoProviderRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

type V1ListAllSsoProviderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse           *http.Response
	ListProvidersResponse *shared.ListProvidersResponse
}

func (o *V1ListAllSsoProviderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1ListAllSsoProviderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1ListAllSsoProviderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1ListAllSsoProviderResponse) GetListProvidersResponse() *shared.ListProvidersResponse {
	if o == nil {
		return nil
	}
	return o.ListProvidersResponse
}
