# V1UpdateASsoProviderRequest

## Example Usage

```typescript
import { V1UpdateASsoProviderRequest } from "supabase/models/operations";

let value: V1UpdateASsoProviderRequest = {
  ref: "<value>",
  providerId: "<id>",
  updateProviderBody: {},
};
```

## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ref`                                                                          | *string*                                                                       | :heavy_check_mark:                                                             | Project ref                                                                    |
| `providerId`                                                                   | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `updateProviderBody`                                                           | [components.UpdateProviderBody](../../models/components/updateproviderbody.md) | :heavy_check_mark:                                                             | N/A                                                                            |