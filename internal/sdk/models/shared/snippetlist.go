// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SnippetList struct {
	Data []SnippetMeta `json:"data"`
}

func (o *SnippetList) GetData() []SnippetMeta {
	if o == nil {
		return []SnippetMeta{}
	}
	return o.Data
}