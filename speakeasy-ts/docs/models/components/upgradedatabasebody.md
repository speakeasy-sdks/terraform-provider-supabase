# UpgradeDatabaseBody

## Example Usage

```typescript
import { UpgradeDatabaseBody } from "supabase/models/components";

let value: UpgradeDatabaseBody = {
  releaseChannel: "alpha",
  targetVersion: "<value>",
};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `releaseChannel`                                                       | [components.ReleaseChannel](../../models/components/releasechannel.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `targetVersion`                                                        | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |