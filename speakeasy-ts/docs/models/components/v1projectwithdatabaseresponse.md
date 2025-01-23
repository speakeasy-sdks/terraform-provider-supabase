# V1ProjectWithDatabaseResponse

## Example Usage

```typescript
import { V1ProjectWithDatabaseResponse } from "supabase/models/components";

let value: V1ProjectWithDatabaseResponse = {
  id: "<id>",
  organizationId: "<id>",
  name: "<value>",
  region: "us-east-1",
  createdAt: "2023-03-29T16:32:59Z",
  database: {
    host: "motionless-metabolite.org",
    version: "<value>",
    postgresEngine: "<value>",
    releaseChannel: "<value>",
  },
  status: "ACTIVE_UNHEALTHY",
};
```

## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `id`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Id of your project                                                                                               |                                                                                                                  |
| `organizationId`                                                                                                 | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Slug of your organization                                                                                        |                                                                                                                  |
| `name`                                                                                                           | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Name of your project                                                                                             |                                                                                                                  |
| `region`                                                                                                         | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Region of your project                                                                                           | us-east-1                                                                                                        |
| `createdAt`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Creation timestamp                                                                                               | 2023-03-29T16:32:59Z                                                                                             |
| `database`                                                                                                       | [components.V1DatabaseResponse](../../models/components/v1databaseresponse.md)                                   | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `status`                                                                                                         | [components.V1ProjectWithDatabaseResponseStatus](../../models/components/v1projectwithdatabaseresponsestatus.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |