// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type GetTypescriptTypesRequest struct {
	IncludedSchemas *string `queryParam:"style=form,explode=true,name=included_schemas"`
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type GetTypescriptTypesResponse struct {
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
	TypescriptResponse *shared.TypescriptResponse
}
