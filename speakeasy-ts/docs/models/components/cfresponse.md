# CfResponse

## Example Usage

```typescript
import { CfResponse } from "supabase/models/components";

let value: CfResponse = {
  success: false,
  errors: [
    {},
  ],
  messages: [
    {},
  ],
  result: {
    id: "<id>",
    hostname: "hidden-hubris.org",
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
  },
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `success`                                                                            | *boolean*                                                                            | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `errors`                                                                             | [components.Errors](../../models/components/errors.md)[]                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `messages`                                                                           | [components.Messages](../../models/components/messages.md)[]                         | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `result`                                                                             | [components.CustomHostnameDetails](../../models/components/customhostnamedetails.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |