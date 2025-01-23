# PostgrestConfigWithJWTSecretResponse

## Example Usage

```typescript
import { PostgrestConfigWithJWTSecretResponse } from "supabase/models/components";

let value: PostgrestConfigWithJWTSecretResponse = {
  maxRows: 998848,
  dbPool: 149448,
  dbSchema: "<value>",
  dbExtraSearchPath: "<value>",
};
```

## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `maxRows`                                                               | *number*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `dbPool`                                                                | *number*                                                                | :heavy_check_mark:                                                      | If `null`, the value is automatically configured based on compute size. |
| `dbSchema`                                                              | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `dbExtraSearchPath`                                                     | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `jwtSecret`                                                             | *string*                                                                | :heavy_minus_sign:                                                      | N/A                                                                     |