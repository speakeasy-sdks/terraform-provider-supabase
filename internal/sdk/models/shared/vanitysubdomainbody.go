// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type VanitySubdomainBody struct {
	VanitySubdomain string `json:"vanity_subdomain"`
}

func (o *VanitySubdomainBody) GetVanitySubdomain() string {
	if o == nil {
		return ""
	}
	return o.VanitySubdomain
}
