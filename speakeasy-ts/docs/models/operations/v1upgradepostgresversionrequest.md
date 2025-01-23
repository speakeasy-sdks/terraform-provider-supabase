# V1UpgradePostgresVersionRequest

## Example Usage

```typescript
import { V1UpgradePostgresVersionRequest } from "supabase/models/operations";

let value: V1UpgradePostgresVersionRequest = {
  ref: "<value>",
  upgradeDatabaseBody: {
    releaseChannel: "withdrawn",
    targetVersion: "<value>",
  },
};
```

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ref`                                                                            | *string*                                                                         | :heavy_check_mark:                                                               | Project ref                                                                      |
| `upgradeDatabaseBody`                                                            | [components.UpgradeDatabaseBody](../../models/components/upgradedatabasebody.md) | :heavy_check_mark:                                                               | N/A                                                                              |