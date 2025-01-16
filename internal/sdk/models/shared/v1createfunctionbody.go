// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V1CreateFunctionBody struct {
	Slug              string   `json:"slug"`
	Name              string   `json:"name"`
	Body              string   `json:"body"`
	VerifyJwt         *bool    `json:"verify_jwt,omitempty"`
	ComputeMultiplier *float64 `json:"compute_multiplier,omitempty"`
}

func (o *V1CreateFunctionBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *V1CreateFunctionBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *V1CreateFunctionBody) GetBody() string {
	if o == nil {
		return ""
	}
	return o.Body
}

func (o *V1CreateFunctionBody) GetVerifyJwt() *bool {
	if o == nil {
		return nil
	}
	return o.VerifyJwt
}

func (o *V1CreateFunctionBody) GetComputeMultiplier() *float64 {
	if o == nil {
		return nil
	}
	return o.ComputeMultiplier
}
