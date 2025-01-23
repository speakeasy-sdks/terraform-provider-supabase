# V1RestorePitrBackupRequest

## Example Usage

```typescript
import { V1RestorePitrBackupRequest } from "supabase/models/operations";

let value: V1RestorePitrBackupRequest = {
  ref: "<value>",
  v1RestorePitrBody: {
    recoveryTimeTargetUnix: 355613,
  },
};
```

## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ref`                                                                        | *string*                                                                     | :heavy_check_mark:                                                           | Project ref                                                                  |
| `v1RestorePitrBody`                                                          | [components.V1RestorePitrBody](../../models/components/v1restorepitrbody.md) | :heavy_check_mark:                                                           | N/A                                                                          |