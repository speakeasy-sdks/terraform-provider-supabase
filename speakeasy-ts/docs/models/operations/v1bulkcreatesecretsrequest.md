# V1BulkCreateSecretsRequest

## Example Usage

```typescript
import { V1BulkCreateSecretsRequest } from "supabase/models/operations";

let value: V1BulkCreateSecretsRequest = {
  ref: "<value>",
  requestBody: [
    {
      name: "string",
      value: "<value>",
    },
  ],
};
```

## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ref`                                                                        | *string*                                                                     | :heavy_check_mark:                                                           | Project ref                                                                  |
| `requestBody`                                                                | [components.CreateSecretBody](../../models/components/createsecretbody.md)[] | :heavy_check_mark:                                                           | N/A                                                                          |