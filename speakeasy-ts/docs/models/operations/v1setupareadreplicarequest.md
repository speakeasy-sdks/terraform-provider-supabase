# V1SetupAReadReplicaRequest

## Example Usage

```typescript
import { V1SetupAReadReplicaRequest } from "supabase/models/operations";

let value: V1SetupAReadReplicaRequest = {
  ref: "<value>",
  setUpReadReplicaBody: {
    readReplicaRegion: "us-east-1",
  },
};
```

## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ref`                                                                              | *string*                                                                           | :heavy_check_mark:                                                                 | Project ref                                                                        |
| `setUpReadReplicaBody`                                                             | [components.SetUpReadReplicaBody](../../models/components/setupreadreplicabody.md) | :heavy_check_mark:                                                                 | N/A                                                                                |