# SslValidation

## Example Usage

```typescript
import { SslValidation } from "supabase/models/components";

let value: SslValidation = {
  status: "<value>",
  validationRecords: [
    {
      txtName: "<value>",
      txtValue: "<value>",
    },
  ],
};
```

## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `status`                                                                     | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `validationRecords`                                                          | [components.ValidationRecord](../../models/components/validationrecord.md)[] | :heavy_check_mark:                                                           | N/A                                                                          |
| `validationErrors`                                                           | [components.ValidationError](../../models/components/validationerror.md)[]   | :heavy_minus_sign:                                                           | N/A                                                                          |