// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SslEnforcementRequest struct {
	RequestedConfig SslEnforcements `json:"requestedConfig"`
}

func (o *SslEnforcementRequest) GetRequestedConfig() SslEnforcements {
	if o == nil {
		return SslEnforcements{}
	}
	return o.RequestedConfig
}
