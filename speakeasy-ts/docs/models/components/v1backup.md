# V1Backup

## Example Usage

```typescript
import { V1Backup } from "supabase/models/components";

let value: V1Backup = {
  status: "PENDING",
  isPhysicalBackup: false,
  insertedAt: "<value>",
};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `status`                                                               | [components.V1BackupStatus](../../models/components/v1backupstatus.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `isPhysicalBackup`                                                     | *boolean*                                                              | :heavy_check_mark:                                                     | N/A                                                                    |
| `insertedAt`                                                           | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |