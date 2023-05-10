// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type CreateSecretsRequest struct {
	RequestBody []shared.CreateSecretBody `request:"mediaType=application/json"`
	// Project ref
	Ref string `pathParam:"style=simple,explode=false,name=ref"`
}

type CreateSecretsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
