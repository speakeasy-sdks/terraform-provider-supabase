# FunctionSlugResponse

## Example Usage

```typescript
import { FunctionSlugResponse } from "supabase/models/components";

let value: FunctionSlugResponse = {
  version: 374170,
  createdAt: 463575,
  updatedAt: 277628,
  id: "<id>",
  slug: "<value>",
  name: "<value>",
  status: "REMOVED",
};
```

## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `version`                                                                                      | *number*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `createdAt`                                                                                    | *number*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updatedAt`                                                                                    | *number*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `slug`                                                                                         | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `name`                                                                                         | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `status`                                                                                       | [components.FunctionSlugResponseStatus](../../models/components/functionslugresponsestatus.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `verifyJwt`                                                                                    | *boolean*                                                                                      | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `importMap`                                                                                    | *boolean*                                                                                      | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `entrypointPath`                                                                               | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `importMapPath`                                                                                | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `computeMultiplier`                                                                            | *number*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |