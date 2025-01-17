// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateProviderBody struct {
	MetadataXML      *string           `json:"metadata_xml,omitempty"`
	MetadataURL      *string           `json:"metadata_url,omitempty"`
	Domains          []string          `json:"domains,omitempty"`
	AttributeMapping *AttributeMapping `json:"attribute_mapping,omitempty"`
}

func (o *UpdateProviderBody) GetMetadataXML() *string {
	if o == nil {
		return nil
	}
	return o.MetadataXML
}

func (o *UpdateProviderBody) GetMetadataURL() *string {
	if o == nil {
		return nil
	}
	return o.MetadataURL
}

func (o *UpdateProviderBody) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *UpdateProviderBody) GetAttributeMapping() *AttributeMapping {
	if o == nil {
		return nil
	}
	return o.AttributeMapping
}
