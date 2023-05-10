// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type ReverifyRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type ReverifyResponse struct {
	ContentType                  string
	StatusCode                   int
	RawResponse                  *http.Response
	UpdateCustomHostnameResponse *shared.UpdateCustomHostnameResponse
}
