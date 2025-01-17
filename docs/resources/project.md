---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "supabase_project Resource - terraform-provider-supabase"
subcategory: ""
description: |-
  Project Resource
---

# supabase_project (Resource)

Project Resource

## Example Usage

```terraform
resource "supabase_project" "my_project" {
  db_pass               = "...my_db_pass..."
  desired_instance_size = "2xlarge"
  kps_enabled           = false
  name                  = "...my_name..."
  organization_id       = "...my_organization_id..."
  plan                  = "free"
  postgres_engine       = "15"
  region                = "us-east-1"
  release_channel       = "withdrawn"
  template_url          = "https://github.com/supabase/supabase/tree/master/examples/slack-clone/nextjs-slack-clone"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `db_pass` (String) Database password. Requires replacement if changed.
- `name` (String) Name of your project, should not contain dots. Requires replacement if changed.
- `organization_id` (String) Slug of your organization. Requires replacement if changed.
- `region` (String) Region you want your server to reside in. must be one of ["us-east-1", "us-east-2", "us-west-1", "us-west-2", "ap-east-1", "ap-southeast-1", "ap-northeast-1", "ap-northeast-2", "ap-southeast-2", "eu-west-1", "eu-west-2", "eu-west-3", "eu-north-1", "eu-central-1", "eu-central-2", "ca-central-1", "ap-south-1", "sa-east-1"]; Requires replacement if changed.

### Optional

- `desired_instance_size` (String) must be one of ["micro", "small", "medium", "large", "xlarge", "2xlarge", "4xlarge", "8xlarge", "12xlarge", "16xlarge"]; Requires replacement if changed.
- `kps_enabled` (Boolean) This field is deprecated and is ignored in this request. Requires replacement if changed.
- `plan` (String) Subscription Plan is now set on organization level and is ignored in this request. must be one of ["free", "pro"]; Requires replacement if changed.
- `postgres_engine` (String) Postgres engine version. If not provided, the latest version will be used. must be "15"; Requires replacement if changed.
- `release_channel` (String) Release channel. If not provided, GA will be used. must be one of ["internal", "alpha", "beta", "ga", "withdrawn"]; Requires replacement if changed.
- `template_url` (String) Template URL used to create the project from the CLI. Requires replacement if changed.

### Read-Only

- `created_at` (String) Creation timestamp
- `id` (String) Id of your project
- `status` (String) must be one of ["ACTIVE_HEALTHY", "ACTIVE_UNHEALTHY", "COMING_UP", "GOING_DOWN", "INACTIVE", "INIT_FAILED", "REMOVED", "RESTARTING", "UNKNOWN", "UPGRADING", "PAUSING", "RESTORING", "RESTORE_FAILED", "PAUSE_FAILED", "RESIZING"]