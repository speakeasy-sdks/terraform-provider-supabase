# V1CreateAFunctionRequest

## Example Usage

```typescript
import { V1CreateAFunctionRequest } from "supabase/models/operations";

let value: V1CreateAFunctionRequest = {
  ref: "<value>",
  v1CreateFunctionBody: {
    slug: "<value>",
    name: "<value>",
    body: "<value>",
  },
};
```

## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ref`                                                                              | *string*                                                                           | :heavy_check_mark:                                                                 | Project ref                                                                        |
| `slug`                                                                             | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `name`                                                                             | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `verifyJwt`                                                                        | *boolean*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `importMap`                                                                        | *boolean*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `entrypointPath`                                                                   | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `importMapPath`                                                                    | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `computeMultiplier`                                                                | *number*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `v1CreateFunctionBody`                                                             | [components.V1CreateFunctionBody](../../models/components/v1createfunctionbody.md) | :heavy_check_mark:                                                                 | N/A                                                                                |