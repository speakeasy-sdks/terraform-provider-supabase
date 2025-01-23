# V1BackupsResponse

## Example Usage

```typescript
import { V1BackupsResponse } from "supabase/models/components";

let value: V1BackupsResponse = {
  region: "<value>",
  walgEnabled: false,
  pitrEnabled: false,
  backups: [
    {
      status: "FAILED",
      isPhysicalBackup: false,
      insertedAt: "<value>",
    },
  ],
  physicalBackupData: {},
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `region`                                                                   | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `walgEnabled`                                                              | *boolean*                                                                  | :heavy_check_mark:                                                         | N/A                                                                        |
| `pitrEnabled`                                                              | *boolean*                                                                  | :heavy_check_mark:                                                         | N/A                                                                        |
| `backups`                                                                  | [components.V1Backup](../../models/components/v1backup.md)[]               | :heavy_check_mark:                                                         | N/A                                                                        |
| `physicalBackupData`                                                       | [components.V1PhysicalBackup](../../models/components/v1physicalbackup.md) | :heavy_check_mark:                                                         | N/A                                                                        |