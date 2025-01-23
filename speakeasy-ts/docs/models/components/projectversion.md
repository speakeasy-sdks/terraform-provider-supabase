# ProjectVersion

## Example Usage

```typescript
import { ProjectVersion } from "supabase/models/components";

let value: ProjectVersion = {
  postgresVersion: "15",
  releaseChannel: "internal",
  appVersion: "<value>",
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `postgresVersion`                                                          | [components.PostgresEngine](../../models/components/postgresengine.md)     | :heavy_check_mark:                                                         | Postgres engine version. If not provided, the latest version will be used. |
| `releaseChannel`                                                           | [components.ReleaseChannel](../../models/components/releasechannel.md)     | :heavy_check_mark:                                                         | N/A                                                                        |
| `appVersion`                                                               | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |