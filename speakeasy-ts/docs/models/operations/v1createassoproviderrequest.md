# V1CreateASsoProviderRequest

## Example Usage

```typescript
import { V1CreateASsoProviderRequest } from "supabase/models/operations";

let value: V1CreateASsoProviderRequest = {
  ref: "<value>",
  createProviderBody: {
    type: "saml",
  },
};
```

## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ref`                                                                          | *string*                                                                       | :heavy_check_mark:                                                             | Project ref                                                                    |
| `createProviderBody`                                                           | [components.CreateProviderBody](../../models/components/createproviderbody.md) | :heavy_check_mark:                                                             | N/A                                                                            |