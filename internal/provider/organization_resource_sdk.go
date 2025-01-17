// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
)

func (r *OrganizationResourceModel) ToSharedCreateOrganizationV1Dto() *shared.CreateOrganizationV1Dto {
	var name string
	name = r.Name.ValueString()

	out := shared.CreateOrganizationV1Dto{
		Name: name,
	}
	return &out
}

func (r *OrganizationResourceModel) RefreshFromSharedOrganizationResponseV1(resp *shared.OrganizationResponseV1) {
	if resp != nil {
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
	}
}
