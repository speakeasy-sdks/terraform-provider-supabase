// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1DeleteABranchRequest struct {
	// Branch ID
	BranchID string `pathParam:"style=simple,explode=false,name=branch_id"`
}

func (o *V1DeleteABranchRequest) GetBranchID() string {
	if o == nil {
		return ""
	}
	return o.BranchID
}

type V1DeleteABranchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse          *http.Response
	BranchDeleteResponse *shared.BranchDeleteResponse
}

func (o *V1DeleteABranchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1DeleteABranchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1DeleteABranchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1DeleteABranchResponse) GetBranchDeleteResponse() *shared.BranchDeleteResponse {
	if o == nil {
		return nil
	}
	return o.BranchDeleteResponse
}
