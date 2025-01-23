# Oauth
(*oauth*)

## Overview

### Available Operations

* [authorizeUser](#authorizeuser) - [Beta] Authorize user through oauth
* [exchangeToken](#exchangetoken) - [Beta] Exchange auth code for user's access and refresh token

## authorizeUser

[Beta] Authorize user through oauth

### Example Usage

```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  await supabase.oauth.authorizeUser({
    clientId: "<id>",
    responseType: "id_token token",
    redirectUri: "https://well-worn-jacket.info",
  });


}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SupabaseCore } from "supabase/core.js";
import { oauthAuthorizeUser } from "supabase/funcs/oauthAuthorizeUser.js";

// Use `SupabaseCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const supabase = new SupabaseCore({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const res = await oauthAuthorizeUser(supabase, {
    clientId: "<id>",
    responseType: "id_token token",
    redirectUri: "https://well-worn-jacket.info",
  });

  if (!res.ok) {
    throw res.error;
  }

  const { value: result } = res;

  
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.V1AuthorizeUserRequest](../../models/operations/v1authorizeuserrequest.md)                                                                                         | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<void\>**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## exchangeToken

[Beta] Exchange auth code for user's access and refresh token

### Example Usage

```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await supabase.oauth.exchangeToken({
    grantType: "authorization_code",
    clientId: "<id>",
    clientSecret: "<value>",
  });

  // Handle the result
  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SupabaseCore } from "supabase/core.js";
import { oauthExchangeToken } from "supabase/funcs/oauthExchangeToken.js";

// Use `SupabaseCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const supabase = new SupabaseCore({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const res = await oauthExchangeToken(supabase, {
    grantType: "authorization_code",
    clientId: "<id>",
    clientSecret: "<value>",
  });

  if (!res.ok) {
    throw res.error;
  }

  const { value: result } = res;

  // Handle the result
  console.log(result);
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.OAuthTokenBody](../../models/components/oauthtokenbody.md)                                                                                                         | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.OAuthTokenResponse](../../models/components/oauthtokenresponse.md)\>**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |