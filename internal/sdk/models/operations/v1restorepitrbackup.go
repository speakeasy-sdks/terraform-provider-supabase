// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
	"net/http"
)

type V1RestorePitrBackupRequest struct {
	// Project ref
	Ref               string                   `pathParam:"style=simple,explode=false,name=ref"`
	V1RestorePitrBody shared.V1RestorePitrBody `request:"mediaType=application/json"`
}

func (o *V1RestorePitrBackupRequest) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1RestorePitrBackupRequest) GetV1RestorePitrBody() shared.V1RestorePitrBody {
	if o == nil {
		return shared.V1RestorePitrBody{}
	}
	return o.V1RestorePitrBody
}

type V1RestorePitrBackupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V1RestorePitrBackupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1RestorePitrBackupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1RestorePitrBackupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
