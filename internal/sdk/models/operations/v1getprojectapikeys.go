// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1GetProjectAPIKeysRequest struct {
	// Project ref
	Ref    string `pathParam:"style=simple,explode=false,name=ref"`
	Reveal bool   `queryParam:"style=form,explode=true,name=reveal"`
}

func (o *V1GetProjectAPIKeysRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1GetProjectAPIKeysRequest) GetReveal() bool {
	if o == nil {
		return false
	}
	return o.Reveal
}

type V1GetProjectAPIKeysResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse     *http.Response
	APIKeyResponses []shared.APIKeyResponse
}

func (o *V1GetProjectAPIKeysResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1GetProjectAPIKeysResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1GetProjectAPIKeysResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V1GetProjectAPIKeysResponse) GetAPIKeyResponses() []shared.APIKeyResponse {
	if o == nil {
		return nil
	}
	return o.APIKeyResponses
}