// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type GetFunctionsRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type GetFunctionsResponse struct {
	ContentType       string
	FunctionResponses []shared.FunctionResponse
	StatusCode        int
	RawResponse       *http.Response
}
