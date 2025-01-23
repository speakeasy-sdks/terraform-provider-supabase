# V1ServiceHealthResponse

## Example Usage

```typescript
import { V1ServiceHealthResponse } from "supabase/models/components";

let value: V1ServiceHealthResponse = {
  name: "rest",
  healthy: false,
  status: "COMING_UP",
};
```

## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `info`                                                                                               | *components.Info*                                                                                    | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `name`                                                                                               | [components.V1ServiceHealthResponseName](../../models/components/v1servicehealthresponsename.md)     | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `healthy`                                                                                            | *boolean*                                                                                            | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `status`                                                                                             | [components.V1ServiceHealthResponseStatus](../../models/components/v1servicehealthresponsestatus.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `error`                                                                                              | *string*                                                                                             | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |