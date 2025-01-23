# V1UpdateHostnameConfigRequest

## Example Usage

```typescript
import { V1UpdateHostnameConfigRequest } from "supabase/models/operations";

let value: V1UpdateHostnameConfigRequest = {
  ref: "<value>",
  updateCustomHostnameBody: {
    customHostname: "<value>",
  },
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ref`                                                                                      | *string*                                                                                   | :heavy_check_mark:                                                                         | Project ref                                                                                |
| `updateCustomHostnameBody`                                                                 | [components.UpdateCustomHostnameBody](../../models/components/updatecustomhostnamebody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |