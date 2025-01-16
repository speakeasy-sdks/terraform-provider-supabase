// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseType string

const (
	ResponseTypeCode         ResponseType = "code"
	ResponseTypeToken        ResponseType = "token"
	ResponseTypeIDTokenToken ResponseType = "id_token token"
)

func (e ResponseType) ToPointer() *ResponseType {
	return &e
}
func (e *ResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "code":
		fallthrough
	case "token":
		fallthrough
	case "id_token token":
		*e = ResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseType: %v", v)
	}
}

type CodeChallengeMethod string

const (
	CodeChallengeMethodPlain  CodeChallengeMethod = "plain"
	CodeChallengeMethodSha256 CodeChallengeMethod = "sha256"
	CodeChallengeMethodS256   CodeChallengeMethod = "S256"
)

func (e CodeChallengeMethod) ToPointer() *CodeChallengeMethod {
	return &e
}
func (e *CodeChallengeMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plain":
		fallthrough
	case "sha256":
		fallthrough
	case "S256":
		*e = CodeChallengeMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CodeChallengeMethod: %v", v)
	}
}

type V1AuthorizeUserRequest struct {
	ClientID            string               `queryParam:"style=form,explode=true,name=client_id"`
	ResponseType        ResponseType         `queryParam:"style=form,explode=true,name=response_type"`
	RedirectURI         string               `queryParam:"style=form,explode=true,name=redirect_uri"`
	Scope               *string              `queryParam:"style=form,explode=true,name=scope"`
	State               *string              `queryParam:"style=form,explode=true,name=state"`
	ResponseMode        *string              `queryParam:"style=form,explode=true,name=response_mode"`
	CodeChallenge       *string              `queryParam:"style=form,explode=true,name=code_challenge"`
	CodeChallengeMethod *CodeChallengeMethod `queryParam:"style=form,explode=true,name=code_challenge_method"`
}

func (o *V1AuthorizeUserRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *V1AuthorizeUserRequest) GetResponseType() ResponseType {
	if o == nil {
		return ResponseType("")
	}
	return o.ResponseType
}

func (o *V1AuthorizeUserRequest) GetRedirectURI() string {
	if o == nil {
		return ""
	}
	return o.RedirectURI
}

func (o *V1AuthorizeUserRequest) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *V1AuthorizeUserRequest) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *V1AuthorizeUserRequest) GetResponseMode() *string {
	if o == nil {
		return nil
	}
	return o.ResponseMode
}

func (o *V1AuthorizeUserRequest) GetCodeChallenge() *string {
	if o == nil {
		return nil
	}
	return o.CodeChallenge
}

func (o *V1AuthorizeUserRequest) GetCodeChallengeMethod() *CodeChallengeMethod {
	if o == nil {
		return nil
	}
	return o.CodeChallengeMethod
}

type V1AuthorizeUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V1AuthorizeUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V1AuthorizeUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V1AuthorizeUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}