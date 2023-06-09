// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type GetPostgresConfigRequest struct {
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type GetPostgresConfigResponse struct {
	ContentType            string
	PostgresConfigResponse *shared.PostgresConfigResponse
	StatusCode             int
	RawResponse            *http.Response
}
