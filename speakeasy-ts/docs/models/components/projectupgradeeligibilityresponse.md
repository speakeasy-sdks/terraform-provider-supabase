# ProjectUpgradeEligibilityResponse

## Example Usage

```typescript
import { ProjectUpgradeEligibilityResponse } from "supabase/models/components";

let value: ProjectUpgradeEligibilityResponse = {
  currentAppVersionReleaseChannel: "ga",
  durationEstimateHours: 456150,
  eligible: false,
  currentAppVersion: "<value>",
  latestAppVersion: "<value>",
  targetUpgradeVersions: [
    {
      postgresVersion: "15",
      releaseChannel: "beta",
      appVersion: "<value>",
    },
  ],
  potentialBreakingChanges: [
    "<value>",
  ],
  legacyAuthCustomRoles: [
    "<value>",
  ],
  extensionDependentObjects: [
    "<value>",
  ],
};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `currentAppVersionReleaseChannel`                                        | [components.ReleaseChannel](../../models/components/releasechannel.md)   | :heavy_check_mark:                                                       | N/A                                                                      |
| `durationEstimateHours`                                                  | *number*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `eligible`                                                               | *boolean*                                                                | :heavy_check_mark:                                                       | N/A                                                                      |
| `currentAppVersion`                                                      | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `latestAppVersion`                                                       | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `targetUpgradeVersions`                                                  | [components.ProjectVersion](../../models/components/projectversion.md)[] | :heavy_check_mark:                                                       | N/A                                                                      |
| `potentialBreakingChanges`                                               | *string*[]                                                               | :heavy_check_mark:                                                       | N/A                                                                      |
| `legacyAuthCustomRoles`                                                  | *string*[]                                                               | :heavy_check_mark:                                                       | N/A                                                                      |
| `extensionDependentObjects`                                              | *string*[]                                                               | :heavy_check_mark:                                                       | N/A                                                                      |