# BranchResponse

## Example Usage

```typescript
import { BranchResponse } from "supabase/models/components";

let value: BranchResponse = {
  id: "<id>",
  name: "<value>",
  projectRef: "<value>",
  parentProjectRef: "<value>",
  isDefault: false,
  resetOnPush: false,
  persistent: false,
  status: "MIGRATIONS_PASSED",
  createdAt: "1726432078023",
  updatedAt: "1737550569016",
};
```

## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `prNumber`                                                                         | *number*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `latestCheckRunId`                                                                 | *number*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `id`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `name`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `projectRef`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `parentProjectRef`                                                                 | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `isDefault`                                                                        | *boolean*                                                                          | :heavy_check_mark:                                                                 | N/A                                                                                |
| `gitBranch`                                                                        | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `resetOnPush`                                                                      | *boolean*                                                                          | :heavy_check_mark:                                                                 | N/A                                                                                |
| `persistent`                                                                       | *boolean*                                                                          | :heavy_check_mark:                                                                 | N/A                                                                                |
| `status`                                                                           | [components.BranchResponseStatus](../../models/components/branchresponsestatus.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `createdAt`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `updatedAt`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |