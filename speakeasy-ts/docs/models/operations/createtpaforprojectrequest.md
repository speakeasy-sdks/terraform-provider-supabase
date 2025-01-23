# CreateTPAForProjectRequest

## Example Usage

```typescript
import { CreateTPAForProjectRequest } from "supabase/models/operations";

let value: CreateTPAForProjectRequest = {
  ref: "<value>",
  createThirdPartyAuthBody: {},
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ref`                                                                                      | *string*                                                                                   | :heavy_check_mark:                                                                         | Project ref                                                                                |
| `createThirdPartyAuthBody`                                                                 | [components.CreateThirdPartyAuthBody](../../models/components/createthirdpartyauthbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |