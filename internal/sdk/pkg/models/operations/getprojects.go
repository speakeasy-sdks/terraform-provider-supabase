// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type GetProjectsResponse struct {
	ContentType      string
	ProjectResponses []shared.ProjectResponse
	StatusCode       int
	RawResponse      *http.Response
}