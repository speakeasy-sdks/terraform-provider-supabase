# SupavisorConfigResponse

## Example Usage

```typescript
import { SupavisorConfigResponse } from "supabase/models/components";

let value: SupavisorConfigResponse = {
  dbPort: 423855,
  defaultPoolSize: 606393,
  maxClientConn: 19193,
  identifier: "<value>",
  databaseType: "PRIMARY",
  isUsingScramAuth: false,
  dbUser: "<value>",
  dbHost: "<value>",
  dbName: "<value>",
  connectionString: "<value>",
  poolMode: "session",
};
```

## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `dbPort`                                                                                                 | *number*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `defaultPoolSize`                                                                                        | *number*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `maxClientConn`                                                                                          | *number*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `identifier`                                                                                             | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `databaseType`                                                                                           | [components.DatabaseType](../../models/components/databasetype.md)                                       | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `isUsingScramAuth`                                                                                       | *boolean*                                                                                                | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `dbUser`                                                                                                 | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `dbHost`                                                                                                 | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `dbName`                                                                                                 | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `connectionString`                                                                                       | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `poolMode`                                                                                               | [components.SupavisorConfigResponsePoolMode](../../models/components/supavisorconfigresponsepoolmode.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |