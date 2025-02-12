// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CustomJwks struct {
}

type CreateThirdPartyAuthBody struct {
	OidcIssuerURL *string     `json:"oidc_issuer_url,omitempty"`
	JwksURL       *string     `json:"jwks_url,omitempty"`
	CustomJwks    *CustomJwks `json:"custom_jwks,omitempty"`
}

func (o *CreateThirdPartyAuthBody) GetOidcIssuerURL() *string {
	if o == nil {
		return nil
	}
	return o.OidcIssuerURL
}

func (o *CreateThirdPartyAuthBody) GetJwksURL() *string {
	if o == nil {
		return nil
	}
	return o.JwksURL
}

func (o *CreateThirdPartyAuthBody) GetCustomJwks() *CustomJwks {
	if o == nil {
		return nil
	}
	return o.CustomJwks
}
