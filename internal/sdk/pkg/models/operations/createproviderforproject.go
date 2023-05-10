// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type CreateProviderForProjectRequest struct {
	CreateProviderBody shared.CreateProviderBody `request:"mediaType=application/json"`
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type CreateProviderForProjectResponse struct {
	ContentType            string
	CreateProviderResponse *shared.CreateProviderResponse
	StatusCode             int
	RawResponse            *http.Response
}
