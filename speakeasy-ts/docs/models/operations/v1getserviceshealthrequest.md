# V1GetServicesHealthRequest

## Example Usage

```typescript
import { V1GetServicesHealthRequest } from "supabase/models/operations";

let value: V1GetServicesHealthRequest = {
  ref: "<value>",
  services: [
    "realtime",
  ],
};
```

## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ref`                                                        | *string*                                                     | :heavy_check_mark:                                           | Project ref                                                  |
| `timeoutMs`                                                  | *number*                                                     | :heavy_minus_sign:                                           | N/A                                                          |
| `services`                                                   | [operations.Services](../../models/operations/services.md)[] | :heavy_check_mark:                                           | N/A                                                          |