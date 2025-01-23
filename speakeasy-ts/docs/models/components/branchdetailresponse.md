# BranchDetailResponse

## Example Usage

```typescript
import { BranchDetailResponse } from "supabase/models/components";

let value: BranchDetailResponse = {
  dbPort: 715190,
  ref: "<value>",
  postgresVersion: "<value>",
  postgresEngine: "<value>",
  releaseChannel: "<value>",
  status: "UPGRADING",
  dbHost: "<value>",
};
```

## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `dbPort`                                               | *number*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `ref`                                                  | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `postgresVersion`                                      | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `postgresEngine`                                       | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `releaseChannel`                                       | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `status`                                               | [components.Status](../../models/components/status.md) | :heavy_check_mark:                                     | N/A                                                    |
| `dbHost`                                               | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `dbUser`                                               | *string*                                               | :heavy_minus_sign:                                     | N/A                                                    |
| `dbPass`                                               | *string*                                               | :heavy_minus_sign:                                     | N/A                                                    |
| `jwtSecret`                                            | *string*                                               | :heavy_minus_sign:                                     | N/A                                                    |