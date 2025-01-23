# FunctionResponse

## Example Usage

```typescript
import { FunctionResponse } from "supabase/models/components";

let value: FunctionResponse = {
  version: 213312,
  createdAt: 518201,
  updatedAt: 25662,
  id: "<id>",
  slug: "<value>",
  name: "<value>",
  status: "ACTIVE",
};
```

## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `version`                                                                              | *number*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `createdAt`                                                                            | *number*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `updatedAt`                                                                            | *number*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `slug`                                                                                 | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `name`                                                                                 | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `status`                                                                               | [components.FunctionResponseStatus](../../models/components/functionresponsestatus.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `verifyJwt`                                                                            | *boolean*                                                                              | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `importMap`                                                                            | *boolean*                                                                              | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `entrypointPath`                                                                       | *string*                                                                               | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `importMapPath`                                                                        | *string*                                                                               | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `computeMultiplier`                                                                    | *number*                                                                               | :heavy_minus_sign:                                                                     | N/A                                                                                    |