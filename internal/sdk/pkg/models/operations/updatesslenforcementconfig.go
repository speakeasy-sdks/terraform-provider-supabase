// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type UpdateSslEnforcementConfigRequest struct {
	SslEnforcementRequest shared.SslEnforcementRequest `request:"mediaType=application/json"`
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type UpdateSslEnforcementConfigResponse struct {
	ContentType            string
	SslEnforcementResponse *shared.SslEnforcementResponse
	StatusCode             int
	RawResponse            *http.Response
}
