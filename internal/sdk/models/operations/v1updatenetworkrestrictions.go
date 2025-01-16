// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1UpdateNetworkRestrictionsRequest struct {
	// Project ref
	Ref                        string                            `pathParam:"style=simple,explode=false,name=ref"`
	NetworkRestrictionsRequest shared.NetworkRestrictionsRequest `request:"mediaType=application/json"`
}

func (o *V1UpdateNetworkRestrictionsRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1UpdateNetworkRestrictionsRequest) GetNetworkRestrictionsRequest() shared.NetworkRestrictionsRequest {
	if o == nil {
		return shared.NetworkRestrictionsRequest{}
	}
	return o.NetworkRestrictionsRequest
}

type V1UpdateNetworkRestrictionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                 *http.Response
	NetworkRestrictionsResponse *shared.NetworkRestrictionsResponse
}

func (o *V1UpdateNetworkRestrictionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1UpdateNetworkRestrictionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1UpdateNetworkRestrictionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1UpdateNetworkRestrictionsResponse) GetNetworkRestrictionsResponse() *shared.NetworkRestrictionsResponse {
	if o == nil {
		return nil
	}
	return o.NetworkRestrictionsResponse
}
