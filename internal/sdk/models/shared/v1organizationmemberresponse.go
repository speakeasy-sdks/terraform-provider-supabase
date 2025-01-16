// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V1OrganizationMemberResponse struct {
	UserID     string  `json:"user_id"`
	UserName   string  `json:"user_name"`
	Email      *string `json:"email,omitempty"`
	RoleName   string  `json:"role_name"`
	MfaEnabled bool    `json:"mfa_enabled"`
}

func (o *V1OrganizationMemberResponse) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *V1OrganizationMemberResponse) GetUserName() string {
	if o == nil {
		return ""
	}
	return o.UserName
}

func (o *V1OrganizationMemberResponse) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *V1OrganizationMemberResponse) GetRoleName() string {
	if o == nil {
		return ""
	}
	return o.RoleName
}

func (o *V1OrganizationMemberResponse) GetMfaEnabled() bool {
	if o == nil {
		return false
	}
	return o.MfaEnabled
}