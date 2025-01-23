# CreateApiKeyRequest

## Example Usage

```typescript
import { CreateApiKeyRequest } from "supabase/models/operations";

let value: CreateApiKeyRequest = {
  ref: "<value>",
  reveal: false,
  createApiKeyBody: {
    type: "secret",
  },
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ref`                                                                      | *string*                                                                   | :heavy_check_mark:                                                         | Project ref                                                                |
| `reveal`                                                                   | *boolean*                                                                  | :heavy_check_mark:                                                         | N/A                                                                        |
| `createApiKeyBody`                                                         | [components.CreateApiKeyBody](../../models/components/createapikeybody.md) | :heavy_check_mark:                                                         | N/A                                                                        |