// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Provider struct {
	ID        string          `json:"id"`
	Saml      *SamlDescriptor `json:"saml,omitempty"`
	Domains   []Domain        `json:"domains,omitempty"`
	CreatedAt *string         `json:"created_at,omitempty"`
	UpdatedAt *string         `json:"updated_at,omitempty"`
}

func (o *Provider) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Provider) GetSaml() *SamlDescriptor {
	if o == nil {
		return nil
	}
	return o.Saml
}

func (o *Provider) GetDomains() []Domain {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *Provider) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Provider) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}