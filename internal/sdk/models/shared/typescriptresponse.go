// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TypescriptResponse struct {
	Types string `json:"types"`
}

func (o *TypescriptResponse) GetTypes() string {
	if o == nil {
		return ""
	}
	return o.Types
}