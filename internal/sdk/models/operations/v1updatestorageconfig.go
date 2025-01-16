// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1UpdateStorageConfigRequest struct {
	// Project ref
	Ref                     string                         `pathParam:"style=simple,explode=false,name=ref"`
	UpdateStorageConfigBody shared.UpdateStorageConfigBody `request:"mediaType=application/json"`
}

func (o *V1UpdateStorageConfigRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1UpdateStorageConfigRequest) GetUpdateStorageConfigBody() shared.UpdateStorageConfigBody {
	if o == nil {
		return shared.UpdateStorageConfigBody{}
	}
	return o.UpdateStorageConfigBody
}

type V1UpdateStorageConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V1UpdateStorageConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1UpdateStorageConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1UpdateStorageConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
