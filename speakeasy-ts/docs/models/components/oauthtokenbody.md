# OAuthTokenBody

## Example Usage

```typescript
import { OAuthTokenBody } from "supabase/models/components";

let value: OAuthTokenBody = {
  grantType: "authorization_code",
  clientId: "<id>",
  clientSecret: "<value>",
};
```

## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `grantType`                                                  | [components.GrantType](../../models/components/granttype.md) | :heavy_check_mark:                                           | N/A                                                          |
| `clientId`                                                   | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `clientSecret`                                               | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `code`                                                       | *string*                                                     | :heavy_minus_sign:                                           | N/A                                                          |
| `codeVerifier`                                               | *string*                                                     | :heavy_minus_sign:                                           | N/A                                                          |
| `redirectUri`                                                | *string*                                                     | :heavy_minus_sign:                                           | N/A                                                          |
| `refreshToken`                                               | *string*                                                     | :heavy_minus_sign:                                           | N/A                                                          |