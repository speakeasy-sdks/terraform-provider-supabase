// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type GetVanitySubdomainConfigRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type GetVanitySubdomainConfigResponse struct {
	ContentType                   string
	StatusCode                    int
	RawResponse                   *http.Response
	VanitySubdomainConfigResponse *shared.VanitySubdomainConfigResponse
}
