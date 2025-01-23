# NetworkRestrictionsResponse

## Example Usage

```typescript
import { NetworkRestrictionsResponse } from "supabase/models/components";

let value: NetworkRestrictionsResponse = {
  entitlement: "allowed",
  config: {},
  status: "applied",
};
```

## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `entitlement`                                                                                                | [components.Entitlement](../../models/components/entitlement.md)                                             | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `config`                                                                                                     | [components.NetworkRestrictionsRequest](../../models/components/networkrestrictionsrequest.md)               | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `oldConfig`                                                                                                  | [components.NetworkRestrictionsRequest](../../models/components/networkrestrictionsrequest.md)               | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `status`                                                                                                     | [components.NetworkRestrictionsResponseStatus](../../models/components/networkrestrictionsresponsestatus.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |