// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V1ProjectRefResponse struct {
	ID   string `json:"id"`
	Ref  string `json:"ref"`
	Name string `json:"name"`
}

func (o *V1ProjectRefResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *V1ProjectRefResponse) GetRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}

func (o *V1ProjectRefResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
