# V1CreateABranchRequest

## Example Usage

```typescript
import { V1CreateABranchRequest } from "supabase/models/operations";

let value: V1CreateABranchRequest = {
  ref: "<value>",
  createBranchBody: {
    branchName: "<value>",
  },
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ref`                                                                      | *string*                                                                   | :heavy_check_mark:                                                         | Project ref                                                                |
| `createBranchBody`                                                         | [components.CreateBranchBody](../../models/components/createbranchbody.md) | :heavy_check_mark:                                                         | N/A                                                                        |