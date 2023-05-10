// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supabase/internal/sdk/pkg/models/shared"
)

type CreateFunctionRequest struct {
	CreateFunctionBody shared.CreateFunctionBody `request:"mediaType=application/json"`
	ImportMap          *bool                     `queryParam:"style=form,explode=true,name=import_map"`
	Name               *string                   `queryParam:"style=form,explode=true,name=name"`
	// Project ref
	Ref       string  `pathParam:"style=simple,explode=false,name=ref"`
	Slug      *string `queryParam:"style=form,explode=true,name=slug"`
	VerifyJwt *bool   `queryParam:"style=form,explode=true,name=verify_jwt"`
}

type CreateFunctionResponse struct {
	ContentType      string
	FunctionResponse *shared.FunctionResponse
	StatusCode       int
	RawResponse      *http.Response
}
