# V1RunAQueryRequest

## Example Usage

```typescript
import { V1RunAQueryRequest } from "supabase/models/operations";

let value: V1RunAQueryRequest = {
  ref: "<value>",
  v1RunQueryBody: {
    query: "<value>",
  },
};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ref`                                                                  | *string*                                                               | :heavy_check_mark:                                                     | Project ref                                                            |
| `v1RunQueryBody`                                                       | [components.V1RunQueryBody](../../models/components/v1runquerybody.md) | :heavy_check_mark:                                                     | N/A                                                                    |