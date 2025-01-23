# V1OrganizationSlugResponse

## Example Usage

```typescript
import { V1OrganizationSlugResponse } from "supabase/models/components";

let value: V1OrganizationSlugResponse = {
  optInTags: [
    "AI_SQL_GENERATOR_OPT_IN",
  ],
  allowedReleaseChannels: [
    "alpha",
  ],
  id: "<id>",
  name: "<value>",
};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `plan`                                                                   | [components.BillingPlanId](../../models/components/billingplanid.md)     | :heavy_minus_sign:                                                       | N/A                                                                      |
| `optInTags`                                                              | [components.OptInTags](../../models/components/optintags.md)[]           | :heavy_check_mark:                                                       | N/A                                                                      |
| `allowedReleaseChannels`                                                 | [components.ReleaseChannel](../../models/components/releasechannel.md)[] | :heavy_check_mark:                                                       | N/A                                                                      |
| `id`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `name`                                                                   | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |