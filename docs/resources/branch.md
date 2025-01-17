---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "supabase_branch Resource - terraform-provider-supabase"
subcategory: ""
description: |-
  Branch Resource
---

# supabase_branch (Resource)

Branch Resource

## Example Usage

```terraform
resource "supabase_branch" "my_branch" {
  branch_id = "...my_branch_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `branch_id` (String) Branch ID

### Read-Only

- `db_host` (String)
- `db_pass` (String)
- `db_port` (Number)
- `db_user` (String)
- `jwt_secret` (String)
- `message` (String)
- `postgres_engine` (String)
- `postgres_version` (String)
- `ref` (String)
- `release_channel` (String)
- `status` (String) must be one of ["ACTIVE_HEALTHY", "ACTIVE_UNHEALTHY", "COMING_UP", "GOING_DOWN", "INACTIVE", "INIT_FAILED", "REMOVED", "RESTARTING", "UNKNOWN", "UPGRADING", "PAUSING", "RESTORING", "RESTORE_FAILED", "PAUSE_FAILED", "RESIZING"]
- `workflow_run_id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import supabase_branch.my_supabase_branch ""
```
