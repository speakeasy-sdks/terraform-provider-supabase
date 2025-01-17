// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
)

func (r *BranchResourceModel) RefreshFromSharedBranchResetResponse(resp *shared.BranchResetResponse) {
	if resp != nil {
		r.Message = types.StringValue(resp.Message)
		r.WorkflowRunID = types.StringValue(resp.WorkflowRunID)
	}
}

func (r *BranchResourceModel) RefreshFromSharedBranchDetailResponse(resp *shared.BranchDetailResponse) {
	if resp != nil {
		r.DbHost = types.StringValue(resp.DbHost)
		r.DbPass = types.StringPointerValue(resp.DbPass)
		r.DbPort = types.Int64Value(resp.DbPort)
		r.DbUser = types.StringPointerValue(resp.DbUser)
		r.JwtSecret = types.StringPointerValue(resp.JwtSecret)
		r.PostgresEngine = types.StringValue(resp.PostgresEngine)
		r.PostgresVersion = types.StringValue(resp.PostgresVersion)
		r.Ref = types.StringValue(resp.Ref)
		r.ReleaseChannel = types.StringValue(resp.ReleaseChannel)
		r.Status = types.StringValue(string(resp.Status))
	}
}

func (r *BranchResourceModel) ToSharedUpdateBranchBody() *shared.UpdateBranchBody {
	status := new(shared.UpdateBranchBodyStatus)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.UpdateBranchBodyStatus(r.Status.ValueString())
	} else {
		status = nil
	}
	out := shared.UpdateBranchBody{
		Status: status,
	}
	return &out
}

func (r *BranchResourceModel) RefreshFromSharedBranchResponse(resp *shared.BranchResponse) {
	if resp != nil {
		r.Status = types.StringValue(string(resp.Status))
	}
}
