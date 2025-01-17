// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &FunctionDataSource{}
var _ datasource.DataSourceWithConfigure = &FunctionDataSource{}

func NewFunctionDataSource() datasource.DataSource {
	return &FunctionDataSource{}
}

// FunctionDataSource is the data source implementation.
type FunctionDataSource struct {
	client *sdk.Supabase
}

// FunctionDataSourceModel describes the data model.
type FunctionDataSourceModel struct {
	ComputeMultiplier types.Number `tfsdk:"compute_multiplier"`
	CreatedAt         types.Int64  `tfsdk:"created_at"`
	EntrypointPath    types.String `tfsdk:"entrypoint_path"`
	ID                types.String `tfsdk:"id"`
	ImportMap         types.Bool   `tfsdk:"import_map"`
	ImportMapPath     types.String `tfsdk:"import_map_path"`
	Name              types.String `tfsdk:"name"`
	Ref               types.String `tfsdk:"ref"`
	Slug              types.String `tfsdk:"slug"`
	Status            types.String `tfsdk:"status"`
	UpdatedAt         types.Int64  `tfsdk:"updated_at"`
	VerifyJwt         types.Bool   `tfsdk:"verify_jwt"`
	Version           types.Int64  `tfsdk:"version"`
}

// Metadata returns the data source type name.
func (r *FunctionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function"
}

// Schema defines the schema for the data source.
func (r *FunctionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Function DataSource",

		Attributes: map[string]schema.Attribute{
			"compute_multiplier": schema.NumberAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed: true,
			},
			"entrypoint_path": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"import_map": schema.BoolAttribute{
				Computed: true,
			},
			"import_map_path": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"ref": schema.StringAttribute{
				Required:    true,
				Description: `Project ref`,
			},
			"slug": schema.StringAttribute{
				Required:    true,
				Description: `Function slug`,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.Int64Attribute{
				Computed: true,
			},
			"verify_jwt": schema.BoolAttribute{
				Computed: true,
			},
			"version": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *FunctionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Supabase)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Supabase, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *FunctionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *FunctionDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var ref string
	ref = data.Ref.ValueString()

	var slug string
	slug = data.Slug.ValueString()

	request := operations.V1GetAFunctionRequest{
		Ref:  ref,
		Slug: slug,
	}
	res, err := r.client.EdgeFunctions.V1GetAFunction(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.FunctionSlugResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFunctionSlugResponse(res.FunctionSlugResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
