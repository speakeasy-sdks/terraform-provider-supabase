# V1UpdateSslEnforcementConfigRequest

## Example Usage

```typescript
import { V1UpdateSslEnforcementConfigRequest } from "supabase/models/operations";

let value: V1UpdateSslEnforcementConfigRequest = {
  ref: "<value>",
  sslEnforcementRequest: {
    requestedConfig: {
      database: false,
    },
  },
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ref`                                                                                | *string*                                                                             | :heavy_check_mark:                                                                   | Project ref                                                                          |
| `sslEnforcementRequest`                                                              | [components.SslEnforcementRequest](../../models/components/sslenforcementrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |