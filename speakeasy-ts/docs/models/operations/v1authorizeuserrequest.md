# V1AuthorizeUserRequest

## Example Usage

```typescript
import { V1AuthorizeUserRequest } from "supabase/models/operations";

let value: V1AuthorizeUserRequest = {
  clientId: "<id>",
  responseType: "token",
  redirectUri: "https://dental-chasuble.info",
};
```

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `clientId`                                                                       | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `responseType`                                                                   | [operations.ResponseType](../../models/operations/responsetype.md)               | :heavy_check_mark:                                                               | N/A                                                                              |
| `redirectUri`                                                                    | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `scope`                                                                          | *string*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `state`                                                                          | *string*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `responseMode`                                                                   | *string*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `codeChallenge`                                                                  | *string*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `codeChallengeMethod`                                                            | [operations.CodeChallengeMethod](../../models/operations/codechallengemethod.md) | :heavy_minus_sign:                                                               | N/A                                                                              |