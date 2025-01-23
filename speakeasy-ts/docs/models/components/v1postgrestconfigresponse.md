# V1PostgrestConfigResponse

## Example Usage

```typescript
import { V1PostgrestConfigResponse } from "supabase/models/components";

let value: V1PostgrestConfigResponse = {
  maxRows: 868126,
  dbPool: 162493,
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