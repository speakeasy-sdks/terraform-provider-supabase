// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateProviderBodyTypeEnum - What type of provider will be created
type CreateProviderBodyTypeEnum string

const (
	CreateProviderBodyTypeEnumSaml CreateProviderBodyTypeEnum = "saml"
)

func (e CreateProviderBodyTypeEnum) ToPointer() *CreateProviderBodyTypeEnum {
	return &e
}

func (e *CreateProviderBodyTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saml":
		*e = CreateProviderBodyTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProviderBodyTypeEnum: %v", v)
	}
}

type CreateProviderBody struct {
	AttributeMapping *AttributeMapping `json:"attribute_mapping,omitempty"`
	Domains          []string          `json:"domains,omitempty"`
	MetadataURL      *string           `json:"metadata_url,omitempty"`
	MetadataXML      *string           `json:"metadata_xml,omitempty"`
	// What type of provider will be created
	Type CreateProviderBodyTypeEnum `json:"type"`
}
