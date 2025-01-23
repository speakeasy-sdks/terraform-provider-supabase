# V1UpdatePgsodiumConfigRequest

## Example Usage

```typescript
import { V1UpdatePgsodiumConfigRequest } from "supabase/models/operations";

let value: V1UpdatePgsodiumConfigRequest = {
  ref: "<value>",
  updatePgsodiumConfigBody: {
    rootKey: "<value>",
  },
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ref`                                                                                      | *string*                                                                                   | :heavy_check_mark:                                                                         | Project ref                                                                                |
| `updatePgsodiumConfigBody`                                                                 | [components.UpdatePgsodiumConfigBody](../../models/components/updatepgsodiumconfigbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |