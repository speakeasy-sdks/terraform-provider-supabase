# V1UpdateAFunctionRequest

## Example Usage

```typescript
import { V1UpdateAFunctionRequest } from "supabase/models/operations";

let value: V1UpdateAFunctionRequest = {
  ref: "<value>",
  slug: "<value>",
  v1UpdateFunctionBody: {},
};
```

## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ref`                                                                              | *string*                                                                           | :heavy_check_mark:                                                                 | Project ref                                                                        |
| `slug`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | Function slug                                                                      |
| `name`                                                                             | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `verifyJwt`                                                                        | *boolean*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `importMap`                                                                        | *boolean*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `entrypointPath`                                                                   | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `importMapPath`                                                                    | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `computeMultiplier`                                                                | *number*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `v1UpdateFunctionBody`                                                             | [components.V1UpdateFunctionBody](../../models/components/v1updatefunctionbody.md) | :heavy_check_mark:                                                                 | N/A                                                                                |