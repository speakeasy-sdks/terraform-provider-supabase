# SamlDescriptor

## Example Usage

```typescript
import { SamlDescriptor } from "supabase/models/components";

let value: SamlDescriptor = {
  id: "<id>",
  entityId: "<id>",
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `entityId`                                                                 | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `metadataUrl`                                                              | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `metadataXml`                                                              | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `attributeMapping`                                                         | [components.AttributeMapping](../../models/components/attributemapping.md) | :heavy_minus_sign:                                                         | N/A                                                                        |