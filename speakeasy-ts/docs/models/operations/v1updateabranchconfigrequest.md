# V1UpdateABranchConfigRequest

## Example Usage

```typescript
import { V1UpdateABranchConfigRequest } from "supabase/models/operations";

let value: V1UpdateABranchConfigRequest = {
  branchId: "<id>",
  updateBranchBody: {},
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `branchId`                                                                 | *string*                                                                   | :heavy_check_mark:                                                         | Branch ID                                                                  |
| `updateBranchBody`                                                         | [components.UpdateBranchBody](../../models/components/updatebranchbody.md) | :heavy_check_mark:                                                         | N/A                                                                        |