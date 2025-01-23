# SetUpReadReplicaBody

## Example Usage

```typescript
import { SetUpReadReplicaBody } from "supabase/models/components";

let value: SetUpReadReplicaBody = {
  readReplicaRegion: "us-east-1",
};
```

## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `readReplicaRegion`                                                          | [components.ReadReplicaRegion](../../models/components/readreplicaregion.md) | :heavy_check_mark:                                                           | Region you want your read replica to reside in                               | us-east-1                                                                    |