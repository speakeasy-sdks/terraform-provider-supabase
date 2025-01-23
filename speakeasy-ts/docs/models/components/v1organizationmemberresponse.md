# V1OrganizationMemberResponse

## Example Usage

```typescript
import { V1OrganizationMemberResponse } from "supabase/models/components";

let value: V1OrganizationMemberResponse = {
  userId: "<id>",
  userName: "Emerald.Marks98",
  roleName: "<value>",
  mfaEnabled: false,
};
```

## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `userId`           | *string*           | :heavy_check_mark: | N/A                |
| `userName`         | *string*           | :heavy_check_mark: | N/A                |
| `email`            | *string*           | :heavy_minus_sign: | N/A                |
| `roleName`         | *string*           | :heavy_check_mark: | N/A                |
| `mfaEnabled`       | *boolean*          | :heavy_check_mark: | N/A                |