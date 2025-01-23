# V1ProjectResponse

## Example Usage

```typescript
import { V1ProjectResponse } from "supabase/models/components";

let value: V1ProjectResponse = {
  id: "<id>",
  organizationId: "<id>",
  name: "<value>",
  region: "us-east-1",
  createdAt: "2023-03-29T16:32:59Z",
  status: "PAUSE_FAILED",
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Id of your project                                                                       |                                                                                          |
| `organizationId`                                                                         | *string*                                                                                 | :heavy_check_mark:                                                                       | Slug of your organization                                                                |                                                                                          |
| `name`                                                                                   | *string*                                                                                 | :heavy_check_mark:                                                                       | Name of your project                                                                     |                                                                                          |
| `region`                                                                                 | *string*                                                                                 | :heavy_check_mark:                                                                       | Region of your project                                                                   | us-east-1                                                                                |
| `createdAt`                                                                              | *string*                                                                                 | :heavy_check_mark:                                                                       | Creation timestamp                                                                       | 2023-03-29T16:32:59Z                                                                     |
| `status`                                                                                 | [components.V1ProjectResponseStatus](../../models/components/v1projectresponsestatus.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |