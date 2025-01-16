// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type UpdateAPIKeyRequest struct {
	// Project ref
	Ref              string                  `pathParam:"style=simple,explode=false,name=ref"`
	ID               string                  `pathParam:"style=simple,explode=false,name=id"`
	Reveal           bool                    `queryParam:"style=form,explode=true,name=reveal"`
	UpdateAPIKeyBody shared.UpdateAPIKeyBody `request:"mediaType=application/json"`
}

func (o *UpdateAPIKeyRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *UpdateAPIKeyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAPIKeyRequest) GetReveal() bool {
	if o == nil {
		return false
	}
	return o.Reveal
}

func (o *UpdateAPIKeyRequest) GetUpdateAPIKeyBody() shared.UpdateAPIKeyBody {
	if o == nil {
		return shared.UpdateAPIKeyBody{}
	}
	return o.UpdateAPIKeyBody
}

type UpdateAPIKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse    *http.Response
	APIKeyResponse *shared.APIKeyResponse
}

func (o *UpdateAPIKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAPIKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAPIKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAPIKeyResponse) GetAPIKeyResponse() *shared.APIKeyResponse {
	if o == nil {
		return nil
	}
	return o.APIKeyResponse
}
