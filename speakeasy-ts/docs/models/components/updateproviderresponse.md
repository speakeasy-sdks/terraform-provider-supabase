# UpdateProviderResponse

## Example Usage

```typescript
import { UpdateProviderResponse } from "supabase/models/components";

let value: UpdateProviderResponse = {
  id: "<id>",
};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `saml`                                                                 | [components.SamlDescriptor](../../models/components/samldescriptor.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `domains`                                                              | [components.Domain](../../models/components/domain.md)[]               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `createdAt`                                                            | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `updatedAt`                                                            | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |