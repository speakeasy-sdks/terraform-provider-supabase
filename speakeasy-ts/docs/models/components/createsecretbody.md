# CreateSecretBody

## Example Usage

```typescript
import { CreateSecretBody } from "supabase/models/components";

let value: CreateSecretBody = {
  name: "string",
  value: "<value>",
};
```

## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | Secret name must not start with the SUPABASE_ prefix. | string                                                |
| `value`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |                                                       |