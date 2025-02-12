// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Name string

const (
	NameGoTrue Name = "GoTrue"
)

func (e Name) ToPointer() *Name {
	return &e
}
func (e *Name) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GoTrue":
		*e = Name(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Name: %v", v)
	}
}

type AuthHealthResponse struct {
	Name Name `json:"name"`
}

func (o *AuthHealthResponse) GetName() Name {
	if o == nil {
		return Name("")
	}
	return o.Name
}
