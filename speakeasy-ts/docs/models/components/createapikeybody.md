# CreateApiKeyBody

## Example Usage

```typescript
import { CreateApiKeyBody } from "supabase/models/components";

let value: CreateApiKeyBody = {
  type: "secret",
};
```

## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `type`                                                                                                       | [components.CreateApiKeyBodyType](../../models/components/createapikeybodytype.md)                           | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `description`                                                                                                | *string*                                                                                                     | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `secretJwtTemplate`                                                                                          | [components.CreateApiKeyBodySecretJwtTemplate](../../models/components/createapikeybodysecretjwttemplate.md) | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |