# DatabaseUpgradeStatusResponse

## Example Usage

```typescript
import { DatabaseUpgradeStatusResponse } from "supabase/models/components";

let value: DatabaseUpgradeStatusResponse = {
  databaseUpgradeStatus: {
    targetVersion: 359508,
    status: 1,
    initiatedAt: "<value>",
    latestStatusAt: "<value>",
  },
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `databaseUpgradeStatus`                                                              | [components.DatabaseUpgradeStatus](../../models/components/databaseupgradestatus.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |