// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1GetABranchConfigRequest struct {
	// Branch ID
	BranchID string `pathParam:"style=simple,explode=false,name=branch_id"`
}

func (o *V1GetABranchConfigRequest) GetBranchID() string {
	if o == nil {
		return ""
	}
	return o.BranchID
}

type V1GetABranchConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse          *http.Response
	BranchDetailResponse *shared.BranchDetailResponse
}

func (o *V1GetABranchConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1GetABranchConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1GetABranchConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1GetABranchConfigResponse) GetBranchDetailResponse() *shared.BranchDetailResponse {
	if o == nil {
		return nil
	}
	return o.BranchDetailResponse
}