# V1RemoveAReadReplicaRequest

## Example Usage

```typescript
import { V1RemoveAReadReplicaRequest } from "supabase/models/operations";

let value: V1RemoveAReadReplicaRequest = {
  ref: "<value>",
  removeReadReplicaBody: {
    databaseIdentifier: "<value>",
  },
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ref`                                                                                | *string*                                                                             | :heavy_check_mark:                                                                   | Project ref                                                                          |
| `removeReadReplicaBody`                                                              | [components.RemoveReadReplicaBody](../../models/components/removereadreplicabody.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |