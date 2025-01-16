// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type DeleteTPAForProjectRequest struct {
	// Project ref
	Ref   string `pathParam:"style=simple,explode=false,name=ref"`
	TpaID string `pathParam:"style=simple,explode=false,name=tpa_id"`
}

func (o *DeleteTPAForProjectRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *DeleteTPAForProjectRequest) GetTpaID() string {
	if o == nil {
		return ""
	}
	return o.TpaID
}

type DeleteTPAForProjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse    *http.Response
	ThirdPartyAuth *shared.ThirdPartyAuth
}

func (o *DeleteTPAForProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTPAForProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTPAForProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTPAForProjectResponse) GetThirdPartyAuth() *shared.ThirdPartyAuth {
	if o == nil {
		return nil
	}
	return o.ThirdPartyAuth
}