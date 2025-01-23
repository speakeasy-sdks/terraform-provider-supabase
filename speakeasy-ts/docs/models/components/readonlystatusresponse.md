# ReadOnlyStatusResponse

## Example Usage

```typescript
import { ReadOnlyStatusResponse } from "supabase/models/components";

let value: ReadOnlyStatusResponse = {
  enabled: false,
  overrideEnabled: false,
  overrideActiveUntil: "<value>",
};
```

## Fields

| Field                 | Type                  | Required              | Description           |
| --------------------- | --------------------- | --------------------- | --------------------- |
| `enabled`             | *boolean*             | :heavy_check_mark:    | N/A                   |
| `overrideEnabled`     | *boolean*             | :heavy_check_mark:    | N/A                   |
| `overrideActiveUntil` | *string*              | :heavy_check_mark:    | N/A                   |