// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SslEnforcementResponse struct {
	CurrentConfig       SslEnforcements `json:"currentConfig"`
	AppliedSuccessfully bool            `json:"appliedSuccessfully"`
}

func (o *SslEnforcementResponse) GetCurrentConfig() SslEnforcements {
	if o == nil {
		return SslEnforcements{}
	}
	return o.CurrentConfig
}

func (o *SslEnforcementResponse) GetAppliedSuccessfully() bool {
	if o == nil {
		return false
	}
	return o.AppliedSuccessfully
}