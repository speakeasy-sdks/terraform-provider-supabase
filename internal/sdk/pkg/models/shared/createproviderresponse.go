// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateProviderResponse struct {
	CreatedAt *string         `json:"created_at,omitempty"`
	Domains   []Domain        `json:"domains,omitempty"`
	ID        string          `json:"id"`
	Saml      *SamlDescriptor `json:"saml,omitempty"`
	UpdatedAt *string         `json:"updated_at,omitempty"`
}
