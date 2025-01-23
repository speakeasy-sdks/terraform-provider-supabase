# SslEnforcementResponse

## Example Usage

```typescript
import { SslEnforcementResponse } from "supabase/models/components";

let value: SslEnforcementResponse = {
  currentConfig: {
    database: false,
  },
  appliedSuccessfully: false,
};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `currentConfig`                                                          | [components.SslEnforcements](../../models/components/sslenforcements.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `appliedSuccessfully`                                                    | *boolean*                                                                | :heavy_check_mark:                                                       | N/A                                                                      |