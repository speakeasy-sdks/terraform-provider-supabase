# OAuthTokenResponse

## Example Usage

```typescript
import { OAuthTokenResponse } from "supabase/models/components";

let value: OAuthTokenResponse = {
  expiresIn: 653108,
  tokenType: "Bearer",
  accessToken: "<value>",
  refreshToken: "<value>",
};
```

## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `expiresIn`                                                  | *number*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `tokenType`                                                  | [components.TokenType](../../models/components/tokentype.md) | :heavy_check_mark:                                           | N/A                                                          |
| `accessToken`                                                | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `refreshToken`                                               | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |