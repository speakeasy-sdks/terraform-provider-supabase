# UpdateCustomHostnameResponse

## Example Usage

```typescript
import { UpdateCustomHostnameResponse } from "supabase/models/components";

let value: UpdateCustomHostnameResponse = {
  status: "4_origin_setup_completed",
  customHostname: "<value>",
  data: {
    success: false,
    errors: [
      {},
    ],
    messages: [
      {},
    ],
    result: {
      id: "<id>",
      hostname: "qualified-hose.org",
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
  },
};
```

## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `status`                                                                                                       | [components.UpdateCustomHostnameResponseStatus](../../models/components/updatecustomhostnameresponsestatus.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `customHostname`                                                                                               | *string*                                                                                                       | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `data`                                                                                                         | [components.CfResponse](../../models/components/cfresponse.md)                                                 | :heavy_check_mark:                                                                                             | N/A                                                                                                            |