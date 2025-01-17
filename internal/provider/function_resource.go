// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/operations"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FunctionResource{}
var _ resource.ResourceWithImportState = &FunctionResource{}

func NewFunctionResource() resource.Resource {
	return &FunctionResource{}
}

// FunctionResource defines the resource implementation.
type FunctionResource struct {
	client *sdk.Supabase
}

// FunctionResourceModel describes the resource data model.
type FunctionResourceModel struct {
	Body              types.String `tfsdk:"body"`
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

func (r *FunctionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function"
}

func (r *FunctionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Function Resource",
		Attributes: map[string]schema.Attribute{
			"body": schema.StringAttribute{
				Required: true,
			},
			"compute_multiplier": schema.NumberAttribute{
				Computed: true,
				Optional: true,
			},
			"created_at": schema.Int64Attribute{
				Computed: true,
			},
			"entrypoint_path": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"import_map": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"import_map_path": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ref": schema.StringAttribute{
				Required:    true,
				Description: `Project ref`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(20, 20),
				},
			},
			"slug": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`/^[A-Za-z0-9_-]+$/`), "must match pattern "+regexp.MustCompile(`/^[A-Za-z0-9_-]+$/`).String()),
				},
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["ACTIVE", "REMOVED", "THROTTLED"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ACTIVE",
						"REMOVED",
						"THROTTLED",
					),
				},
			},
			"updated_at": schema.Int64Attribute{
				Computed: true,
			},
			"verify_jwt": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"version": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *FunctionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Supabase)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Supabase, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *FunctionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *FunctionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var ref string
	ref = data.Ref.ValueString()

	slug := new(string)
	if !data.Slug.IsUnknown() && !data.Slug.IsNull() {
		*slug = data.Slug.ValueString()
	} else {
		slug = nil
	}
	name := new(string)
	if !data.Name.IsUnknown() && !data.Name.IsNull() {
		*name = data.Name.ValueString()
	} else {
		name = nil
	}
	verifyJwt := new(bool)
	if !data.VerifyJwt.IsUnknown() && !data.VerifyJwt.IsNull() {
		*verifyJwt = data.VerifyJwt.ValueBool()
	} else {
		verifyJwt = nil
	}
	importMap := new(bool)
	if !data.ImportMap.IsUnknown() && !data.ImportMap.IsNull() {
		*importMap = data.ImportMap.ValueBool()
	} else {
		importMap = nil
	}
	entrypointPath := new(string)
	if !data.EntrypointPath.IsUnknown() && !data.EntrypointPath.IsNull() {
		*entrypointPath = data.EntrypointPath.ValueString()
	} else {
		entrypointPath = nil
	}
	importMapPath := new(string)
	if !data.ImportMapPath.IsUnknown() && !data.ImportMapPath.IsNull() {
		*importMapPath = data.ImportMapPath.ValueString()
	} else {
		importMapPath = nil
	}
	computeMultiplier := new(float64)
	if !data.ComputeMultiplier.IsUnknown() && !data.ComputeMultiplier.IsNull() {
		*computeMultiplier, _ = data.ComputeMultiplier.ValueBigFloat().Float64()
	} else {
		computeMultiplier = nil
	}
	v1CreateFunctionBody := *data.ToSharedV1CreateFunctionBody()
	request := operations.V1CreateAFunctionRequest{
		Ref:                  ref,
		Slug:                 slug,
		Name:                 name,
		VerifyJwt:            verifyJwt,
		ImportMap:            importMap,
		EntrypointPath:       entrypointPath,
		ImportMapPath:        importMapPath,
		ComputeMultiplier:    computeMultiplier,
		V1CreateFunctionBody: v1CreateFunctionBody,
	}
	res, err := r.client.EdgeFunctions.V1CreateAFunction(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.FunctionResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFunctionResponse(res.FunctionResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var ref1 string
	ref1 = data.Ref.ValueString()

	var slug1 string
	slug1 = data.Slug.ValueString()

	request1 := operations.V1GetAFunctionRequest{
		Ref:  ref1,
		Slug: slug1,
	}
	res1, err := r.client.EdgeFunctions.V1GetAFunction(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.FunctionSlugResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedFunctionSlugResponse(res1.FunctionSlugResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FunctionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *FunctionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

func (r *FunctionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *FunctionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var ref string
	ref = data.Ref.ValueString()

	var slug string
	slug = data.Slug.ValueString()

	name := new(string)
	if !data.Name.IsUnknown() && !data.Name.IsNull() {
		*name = data.Name.ValueString()
	} else {
		name = nil
	}
	verifyJwt := new(bool)
	if !data.VerifyJwt.IsUnknown() && !data.VerifyJwt.IsNull() {
		*verifyJwt = data.VerifyJwt.ValueBool()
	} else {
		verifyJwt = nil
	}
	importMap := new(bool)
	if !data.ImportMap.IsUnknown() && !data.ImportMap.IsNull() {
		*importMap = data.ImportMap.ValueBool()
	} else {
		importMap = nil
	}
	entrypointPath := new(string)
	if !data.EntrypointPath.IsUnknown() && !data.EntrypointPath.IsNull() {
		*entrypointPath = data.EntrypointPath.ValueString()
	} else {
		entrypointPath = nil
	}
	importMapPath := new(string)
	if !data.ImportMapPath.IsUnknown() && !data.ImportMapPath.IsNull() {
		*importMapPath = data.ImportMapPath.ValueString()
	} else {
		importMapPath = nil
	}
	computeMultiplier := new(float64)
	if !data.ComputeMultiplier.IsUnknown() && !data.ComputeMultiplier.IsNull() {
		*computeMultiplier, _ = data.ComputeMultiplier.ValueBigFloat().Float64()
	} else {
		computeMultiplier = nil
	}
	v1UpdateFunctionBody := *data.ToSharedV1UpdateFunctionBody()
	request := operations.V1UpdateAFunctionRequest{
		Ref:                  ref,
		Slug:                 slug,
		Name:                 name,
		VerifyJwt:            verifyJwt,
		ImportMap:            importMap,
		EntrypointPath:       entrypointPath,
		ImportMapPath:        importMapPath,
		ComputeMultiplier:    computeMultiplier,
		V1UpdateFunctionBody: v1UpdateFunctionBody,
	}
	res, err := r.client.EdgeFunctions.V1UpdateAFunction(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.FunctionResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFunctionResponse(res.FunctionResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var ref1 string
	ref1 = data.Ref.ValueString()

	var slug1 string
	slug1 = data.Slug.ValueString()

	request1 := operations.V1GetAFunctionRequest{
		Ref:  ref1,
		Slug: slug1,
	}
	res1, err := r.client.EdgeFunctions.V1GetAFunction(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.FunctionSlugResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedFunctionSlugResponse(res1.FunctionSlugResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FunctionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *FunctionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request := operations.V1DeleteAFunctionRequest{
		Ref:  ref,
		Slug: slug,
	}
	res, err := r.client.EdgeFunctions.V1DeleteAFunction(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *FunctionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		Ref  string `json:"ref"`
		Slug string `json:"slug"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "ref": "",  "slug": ""}': `+err.Error())
		return
	}

	if len(data.Ref) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field ref is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), data.Ref)...)
	if len(data.Slug) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field slug is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("slug"), data.Slug)...)

}