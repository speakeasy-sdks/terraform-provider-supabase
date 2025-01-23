# ThirdPartyAuth

## Example Usage

```typescript
import { ThirdPartyAuth } from "supabase/models/components";

let value: ThirdPartyAuth = {
  id: "<id>",
  type: "<value>",
  insertedAt: "<value>",
  updatedAt: "1737594614030",
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `type`                                                                                     | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `oidcIssuerUrl`                                                                            | *string*                                                                                   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `jwksUrl`                                                                                  | *string*                                                                                   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `customJwks`                                                                               | [components.ThirdPartyAuthCustomJwks](../../models/components/thirdpartyauthcustomjwks.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `resolvedJwks`                                                                             | [components.ResolvedJwks](../../models/components/resolvedjwks.md)                         | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `insertedAt`                                                                               | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `updatedAt`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `resolvedAt`                                                                               | *string*                                                                                   | :heavy_minus_sign:                                                                         | N/A                                                                                        |