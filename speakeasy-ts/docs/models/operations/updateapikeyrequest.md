# UpdateApiKeyRequest

## Example Usage

```typescript
import { UpdateApiKeyRequest } from "supabase/models/operations";

let value: UpdateApiKeyRequest = {
  ref: "<value>",
  id: "<id>",
  reveal: false,
  updateApiKeyBody: {},
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ref`                                                                      | *string*                                                                   | :heavy_check_mark:                                                         | Project ref                                                                |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `reveal`                                                                   | *boolean*                                                                  | :heavy_check_mark:                                                         | N/A                                                                        |
| `updateApiKeyBody`                                                         | [components.UpdateApiKeyBody](../../models/components/updateapikeybody.md) | :heavy_check_mark:                                                         | N/A                                                                        |