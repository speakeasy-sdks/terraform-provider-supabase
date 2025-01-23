# CustomHostnameDetails

## Example Usage

```typescript
import { CustomHostnameDetails } from "supabase/models/components";

let value: CustomHostnameDetails = {
  id: "<id>",
  hostname: "strict-pressure.com",
  ssl: {
    status: "<value>",
    validationRecords: [
      {
        txtName: "<value>",
        txtValue: "<value>",
      },
    ],
  },
  ownershipVerification: {
    type: "<value>",
    name: "<value>",
    value: "<value>",
  },
  customOriginServer: "<value>",
  status: "<value>",
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `hostname`                                                                           | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `ssl`                                                                                | [components.SslValidation](../../models/components/sslvalidation.md)                 | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `ownershipVerification`                                                              | [components.OwnershipVerification](../../models/components/ownershipverification.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `customOriginServer`                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `verificationErrors`                                                                 | *string*[]                                                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `status`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |