// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1ListAllOrganizationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse             *http.Response
	OrganizationResponseV1s []shared.OrganizationResponseV1
}

func (o *V1ListAllOrganizationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1ListAllOrganizationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1ListAllOrganizationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1ListAllOrganizationsResponse) GetOrganizationResponseV1s() []shared.OrganizationResponseV1 {
	if o == nil {
		return nil
	}
	return o.OrganizationResponseV1s
}