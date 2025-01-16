// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeSQL Type = "sql"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sql":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Visibility string

const (
	VisibilityUser    Visibility = "user"
	VisibilityProject Visibility = "project"
	VisibilityOrg     Visibility = "org"
	VisibilityPublic  Visibility = "public"
)

func (e Visibility) ToPointer() *Visibility {
	return &e
}
func (e *Visibility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		fallthrough
	case "project":
		fallthrough
	case "org":
		fallthrough
	case "public":
		*e = Visibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Visibility: %v", v)
	}
}

type SnippetMeta struct {
	ID          string         `json:"id"`
	InsertedAt  string         `json:"inserted_at"`
	UpdatedAt   string         `json:"updated_at"`
	Type        Type           `json:"type"`
	Visibility  Visibility     `json:"visibility"`
	Name        string         `json:"name"`
	Description *string        `json:"description,omitempty"`
	Project     SnippetProject `json:"project"`
	Owner       SnippetUser    `json:"owner"`
	UpdatedBy   SnippetUser    `json:"updated_by"`
}

func (o *SnippetMeta) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SnippetMeta) GetInsertedAt() string {
	if o == nil {
		return ""
	}
	return o.InsertedAt
}

func (o *SnippetMeta) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *SnippetMeta) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *SnippetMeta) GetVisibility() Visibility {
	if o == nil {
		return Visibility("")
	}
	return o.Visibility
}

func (o *SnippetMeta) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SnippetMeta) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SnippetMeta) GetProject() SnippetProject {
	if o == nil {
		return SnippetProject{}
	}
	return o.Project
}

func (o *SnippetMeta) GetOwner() SnippetUser {
	if o == nil {
		return SnippetUser{}
	}
	return o.Owner
}

func (o *SnippetMeta) GetUpdatedBy() SnippetUser {
	if o == nil {
		return SnippetUser{}
	}
	return o.UpdatedBy
}
