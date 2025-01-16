// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeRefreshToken      GrantType = "refresh_token"
)

func (e GrantType) ToPointer() *GrantType {
	return &e
}
func (e *GrantType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "authorization_code":
		fallthrough
	case "refresh_token":
		*e = GrantType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GrantType: %v", v)
	}
}

type OAuthTokenBody struct {
	GrantType    GrantType `form:"name=grant_type"`
	ClientID     string    `form:"name=client_id"`
	ClientSecret string    `form:"name=client_secret"`
	Code         *string   `form:"name=code"`
	CodeVerifier *string   `form:"name=code_verifier"`
	RedirectURI  *string   `form:"name=redirect_uri"`
	RefreshToken *string   `form:"name=refresh_token"`
}

func (o *OAuthTokenBody) GetGrantType() GrantType {
	if o == nil {
		return GrantType("")
	}
	return o.GrantType
}

func (o *OAuthTokenBody) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *OAuthTokenBody) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *OAuthTokenBody) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *OAuthTokenBody) GetCodeVerifier() *string {
	if o == nil {
		return nil
	}
	return o.CodeVerifier
}

func (o *OAuthTokenBody) GetRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.RedirectURI
}

func (o *OAuthTokenBody) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}
