# CreateProviderBody

## Example Usage

```typescript
import { CreateProviderBody } from "supabase/models/components";

let value: CreateProviderBody = {
  type: "saml",
};
```

## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `type`                                                                                 | [components.CreateProviderBodyType](../../models/components/createproviderbodytype.md) | :heavy_check_mark:                                                                     | What type of provider will be created                                                  |
| `metadataXml`                                                                          | *string*                                                                               | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `metadataUrl`                                                                          | *string*                                                                               | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `domains`                                                                              | *string*[]                                                                             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `attributeMapping`                                                                     | [components.AttributeMapping](../../models/components/attributemapping.md)             | :heavy_minus_sign:                                                                     | N/A                                                                                    |