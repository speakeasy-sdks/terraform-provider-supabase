# CreateBranchBody

## Example Usage

```typescript
import { CreateBranchBody } from "supabase/models/components";

let value: CreateBranchBody = {
  branchName: "<value>",
};
```

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `desiredInstanceSize`                                                            | [components.DesiredInstanceSize](../../models/components/desiredinstancesize.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `releaseChannel`                                                                 | [components.ReleaseChannel](../../models/components/releasechannel.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |
| `postgresEngine`                                                                 | [components.PostgresEngine](../../models/components/postgresengine.md)           | :heavy_minus_sign:                                                               | Postgres engine version. If not provided, the latest version will be used.       |
| `branchName`                                                                     | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `gitBranch`                                                                      | *string*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `persistent`                                                                     | *boolean*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `region`                                                                         | *string*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |