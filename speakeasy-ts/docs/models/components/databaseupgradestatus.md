# DatabaseUpgradeStatus

## Example Usage

```typescript
import { DatabaseUpgradeStatus } from "supabase/models/components";

let value: DatabaseUpgradeStatus = {
  targetVersion: 943749,
  status: 2,
  initiatedAt: "<value>",
  latestStatusAt: "<value>",
};
```

## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `targetVersion`                                                                                                  | *number*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `status`                                                                                                         | [components.DatabaseUpgradeStatusResponseStatus](../../models/components/databaseupgradestatusresponsestatus.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `initiatedAt`                                                                                                    | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `latestStatusAt`                                                                                                 | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `error`                                                                                                          | [components.ErrorT](../../models/components/errort.md)                                                           | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `progress`                                                                                                       | [components.Progress](../../models/components/progress.md)                                                       | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |