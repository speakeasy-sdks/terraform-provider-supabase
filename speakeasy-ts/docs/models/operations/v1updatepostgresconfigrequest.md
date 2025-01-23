# V1UpdatePostgresConfigRequest

## Example Usage

```typescript
import { V1UpdatePostgresConfigRequest } from "supabase/models/operations";

let value: V1UpdatePostgresConfigRequest = {
  ref: "<value>",
  updatePostgresConfigBody: {},
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ref`                                                                                      | *string*                                                                                   | :heavy_check_mark:                                                                         | Project ref                                                                                |
| `updatePostgresConfigBody`                                                                 | [components.UpdatePostgresConfigBody](../../models/components/updatepostgresconfigbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |