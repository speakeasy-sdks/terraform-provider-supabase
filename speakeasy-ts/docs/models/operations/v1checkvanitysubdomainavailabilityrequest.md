# V1CheckVanitySubdomainAvailabilityRequest

## Example Usage

```typescript
import { V1CheckVanitySubdomainAvailabilityRequest } from "supabase/models/operations";

let value: V1CheckVanitySubdomainAvailabilityRequest = {
  ref: "<value>",
  vanitySubdomainBody: {
    vanitySubdomain: "<value>",
  },
};
```

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ref`                                                                            | *string*                                                                         | :heavy_check_mark:                                                               | Project ref                                                                      |
| `vanitySubdomainBody`                                                            | [components.VanitySubdomainBody](../../models/components/vanitysubdomainbody.md) | :heavy_check_mark:                                                               | N/A                                                                              |