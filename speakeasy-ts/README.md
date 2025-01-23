# supabase

Developer-friendly & type-safe Typescript SDK specifically catered to leverage *supabase* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=supabase&utm_campaign=typescript"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/speakeasy-onboarding/onboarding). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Supabase API: Supabase API
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [supabase](#supabase)
  * [SDK Installation](#sdk-installation)
  * [Requirements](#requirements)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Standalone functions](#standalone-functions)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Debugging](#debugging)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

> [!TIP]
> To finish publishing your SDK to npm and others you must [run your first generation action](https://www.speakeasy.com/docs/github-setup#step-by-step-guide).


The SDK can be installed with either [npm](https://www.npmjs.com/), [pnpm](https://pnpm.io/), [bun](https://bun.sh/) or [yarn](https://classic.yarnpkg.com/en/) package managers.

### NPM

```bash
npm add <UNSET>
```

### PNPM

```bash
pnpm add <UNSET>
```

### Bun

```bash
bun add <UNSET>
```

### Yarn

```bash
yarn add <UNSET> zod

# Note that Yarn does not install peer dependencies automatically. You will need
# to install zod as shown above.
```
<!-- End SDK Installation [installation] -->

<!-- Start Requirements [requirements] -->
## Requirements

For supported JavaScript runtimes, please consult [RUNTIMES.md](RUNTIMES.md).
<!-- End Requirements [requirements] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await supabase.environments.getBranchConfig({
    branchId: "<id>",
  });

  // Handle the result
  console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type | Scheme      |
| -------- | ---- | ----------- |
| `bearer` | http | HTTP Bearer |

To authenticate with the API the `bearer` parameter must be set when initializing the SDK client instance. For example:
```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await supabase.environments.getBranchConfig({
    branchId: "<id>",
  });

  // Handle the result
  console.log(result);
}

run();

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [auth](docs/sdks/auth/README.md)

* [getConfig](docs/sdks/auth/README.md#getconfig) - Gets project's auth config
* [updateAuthConfig](docs/sdks/auth/README.md#updateauthconfig) - Updates a project's auth config
* [createThirdPartyAuth](docs/sdks/auth/README.md#createthirdpartyauth) - Creates a new third-party auth integration
* [listThirdPartyAuth](docs/sdks/auth/README.md#listthirdpartyauth) - [Alpha] Lists all third-party auth integrations
* [deleteThirdPartyAuth](docs/sdks/auth/README.md#deletethirdpartyauth) - [Alpha] Removes a third-party auth integration
* [getThirdPartyAuth](docs/sdks/auth/README.md#getthirdpartyauth) - [Alpha] Get a third-party integration
* [createSsoProvider](docs/sdks/auth/README.md#createssoprovider) - Creates a new SSO provider
* [listSsoProviders](docs/sdks/auth/README.md#listssoproviders) - Lists all SSO providers
* [getSsoProvider](docs/sdks/auth/README.md#getssoprovider) - Gets a SSO provider by its UUID
* [updateSsoProvider](docs/sdks/auth/README.md#updatessoprovider) - Updates a SSO provider by its UUID
* [deleteSsoProvider](docs/sdks/auth/README.md#deletessoprovider) - Removes a SSO provider by its UUID

### [database](docs/sdks/database/README.md)

* [getSnippet](docs/sdks/database/README.md#getsnippet) - Gets a specific SQL snippet
* [getSSLEnforcement](docs/sdks/database/README.md#getsslenforcement) - [Beta] Get project's SSL enforcement configuration.
* [updateSSLEnforcement](docs/sdks/database/README.md#updatesslenforcement) - [Beta] Update project's SSL enforcement configuration.
* [getReadonlyModeStatus](docs/sdks/database/README.md#getreadonlymodestatus) - Returns project's readonly mode status
* [disableReadonlyModeTemporarily](docs/sdks/database/README.md#disablereadonlymodetemporarily) - Disables project's readonly mode for the next 15 minutes
* [setupReadReplica](docs/sdks/database/README.md#setupreadreplica) - [Beta] Set up a read replica
* [removeReadReplica](docs/sdks/database/README.md#removereadreplica) - [Beta] Remove a read replica
* [getPostgresConfig](docs/sdks/database/README.md#getpostgresconfig) - Gets project's Postgres config
* [getPgbouncerConfig](docs/sdks/database/README.md#getpgbouncerconfig) - Get project's pgbouncer config
* [getSupavisorConfig](docs/sdks/database/README.md#getsupavisorconfig) - Gets project's supavisor config
* [runQuery](docs/sdks/database/README.md#runquery) - [Beta] Run sql query
* [enableWebhook](docs/sdks/database/README.md#enablewebhook) - [Beta] Enables Database Webhooks on the project
* [listBackups](docs/sdks/database/README.md#listbackups) - Lists all backups
* [restorePitrBackup](docs/sdks/database/README.md#restorepitrbackup) - Restores a PITR backup for a database

### [databases](docs/sdks/databases/README.md)

* [generateTypescriptTypes](docs/sdks/databases/README.md#generatetypescripttypes) - Generate TypeScript types
* [updatePostgresConfig](docs/sdks/databases/README.md#updatepostgresconfig) - Updates project's Postgres config
* [updateSupavisorConfig](docs/sdks/databases/README.md#updatesupavisorconfig) - Updates project's supavisor config

### [domains](docs/sdks/domains/README.md)

* [getHostnameConfig](docs/sdks/domains/README.md#gethostnameconfig) - [Beta] Gets project's custom hostname config
* [deleteHostnameConfig](docs/sdks/domains/README.md#deletehostnameconfig) - [Beta] Deletes a project's custom hostname configuration
* [initializeHostnameConfig](docs/sdks/domains/README.md#initializehostnameconfig) - [Beta] Updates project's custom hostname configuration
* [verifyDnsConfig](docs/sdks/domains/README.md#verifydnsconfig) - [Beta] Attempts to verify the DNS configuration for project's custom hostname configuration
* [activateCustomHostname](docs/sdks/domains/README.md#activatecustomhostname) - [Beta] Activates a custom hostname for a project.
* [getVanitySubdomain](docs/sdks/domains/README.md#getvanitysubdomain) - [Beta] Gets current vanity subdomain config
* [deleteVanitySubdomain](docs/sdks/domains/README.md#deletevanitysubdomain) - [Beta] Deletes a project's vanity subdomain configuration
* [checkVanitySubdomainAvailability](docs/sdks/domains/README.md#checkvanitysubdomainavailability) - [Beta] Checks vanity subdomain availability
* [activateVanitySubdomain](docs/sdks/domains/README.md#activatevanitysubdomain) - [Beta] Activates a vanity subdomain for a project.

### [edgeFunctions](docs/sdks/edgefunctions/README.md)

* [create](docs/sdks/edgefunctions/README.md#create) - Create a function
* [list](docs/sdks/edgefunctions/README.md#list) - List all functions
* [getFunction](docs/sdks/edgefunctions/README.md#getfunction) - Retrieve a function
* [updateFunction](docs/sdks/edgefunctions/README.md#updatefunction) - Update a function
* [deleteFunction](docs/sdks/edgefunctions/README.md#deletefunction) - Delete a function
* [getFunctionBody](docs/sdks/edgefunctions/README.md#getfunctionbody) - Retrieve a function body

### [environments](docs/sdks/environments/README.md)

* [getBranchConfig](docs/sdks/environments/README.md#getbranchconfig) - Get database branch config
* [updateBranchConfig](docs/sdks/environments/README.md#updatebranchconfig) - Update database branch config
* [deleteBranch](docs/sdks/environments/README.md#deletebranch) - Delete a database branch
* [resetBranch](docs/sdks/environments/README.md#resetbranch) - Resets a database branch
* [listBranches](docs/sdks/environments/README.md#listbranches) - List all database branches
* [createBranch](docs/sdks/environments/README.md#createbranch) - Create a database branch
* [disablePreviewBranching](docs/sdks/environments/README.md#disablepreviewbranching) - Disables preview branching

### [oauth](docs/sdks/oauth/README.md)

* [authorizeUser](docs/sdks/oauth/README.md#authorizeuser) - [Beta] Authorize user through oauth
* [exchangeToken](docs/sdks/oauth/README.md#exchangetoken) - [Beta] Exchange auth code for user's access and refresh token

### [organizations](docs/sdks/organizations/README.md)

* [list](docs/sdks/organizations/README.md#list) - List all organizations
* [create](docs/sdks/organizations/README.md#create) - Create an organization
* [listMembers](docs/sdks/organizations/README.md#listmembers) - List members of an organization
* [get](docs/sdks/organizations/README.md#get) - Gets information about the organization

### [projects](docs/sdks/projects/README.md)

* [list](docs/sdks/projects/README.md#list) - List all projects
* [create](docs/sdks/projects/README.md#create) - Create a project
* [listNetworkBans](docs/sdks/projects/README.md#listnetworkbans) - [Beta] Gets project's network bans
* [deleteNetworkBans](docs/sdks/projects/README.md#deletenetworkbans) - [Beta] Remove network bans.
* [getNetworkRestrictions](docs/sdks/projects/README.md#getnetworkrestrictions) - [Beta] Gets project's network restrictions
* [updateNetworkRestrictions](docs/sdks/projects/README.md#updatenetworkrestrictions) - [Beta] Updates project's network restrictions
* [get](docs/sdks/projects/README.md#get) - Gets a specific project that belongs to the authenticated user
* [delete](docs/sdks/projects/README.md#delete) - Deletes the given project
* [upgradePostgresVersion](docs/sdks/projects/README.md#upgradepostgresversion) - [Beta] Upgrades the project's Postgres version
* [getUpgradeEligibility](docs/sdks/projects/README.md#getupgradeeligibility) - [Beta] Returns the project's eligibility for upgrades
* [getUpgradeStatus](docs/sdks/projects/README.md#getupgradestatus) - [Beta] Gets the latest status of the project's upgrade
* [getHealth](docs/sdks/projects/README.md#gethealth) - Gets project's service health status

### [rest](docs/sdks/rest/README.md)

* [getConfig](docs/sdks/rest/README.md#getconfig) - Gets project's postgrest config
* [updatePostgrestConfig](docs/sdks/rest/README.md#updatepostgrestconfig) - Updates project's postgrest config

### [secrets](docs/sdks/secrets/README.md)

* [getApiKeys](docs/sdks/secrets/README.md#getapikeys) - Get project api keys
* [createApiKey](docs/sdks/secrets/README.md#createapikey) - [Alpha] Creates a new API key for the project
* [updateApiKey](docs/sdks/secrets/README.md#updateapikey) - [Alpha] Updates an API key for the project
* [getApiKey](docs/sdks/secrets/README.md#getapikey) - [Alpha] Get API key
* [deleteApiKey](docs/sdks/secrets/README.md#deleteapikey) - [Alpha] Deletes an API key for the project
* [getPGSodiumConfig](docs/sdks/secrets/README.md#getpgsodiumconfig) - [Beta] Gets project's pgsodium config
* [updatePgsodiumConfig](docs/sdks/secrets/README.md#updatepgsodiumconfig) - [Beta] Updates project's pgsodium config. Updating the root_key can cause all data encrypted with the older key to become inaccessible.
* [list](docs/sdks/secrets/README.md#list) - List all secrets
* [bulkCreate](docs/sdks/secrets/README.md#bulkcreate) - Bulk create secrets
* [bulkDelete](docs/sdks/secrets/README.md#bulkdelete) - Bulk delete secrets

### [snippets](docs/sdks/snippets/README.md)

* [listAll](docs/sdks/snippets/README.md#listall) - Lists SQL snippets for the logged in user

### [storage](docs/sdks/storage/README.md)

* [getConfig](docs/sdks/storage/README.md#getconfig) - Gets project's storage config
* [patchConfig](docs/sdks/storage/README.md#patchconfig) - Updates project's storage config
* [listBuckets](docs/sdks/storage/README.md#listbuckets) - Lists all buckets


</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Standalone functions [standalone-funcs] -->
## Standalone functions

All the methods listed above are available as standalone functions. These
functions are ideal for use in applications running in the browser, serverless
runtimes or other environments where application bundle size is a primary
concern. When using a bundler to build your application, all unused
functionality will be either excluded from the final bundle or tree-shaken away.

To read more about standalone functions, check [FUNCTIONS.md](./FUNCTIONS.md).

<details>

<summary>Available standalone functions</summary>

- [`authCreateSsoProvider`](docs/sdks/auth/README.md#createssoprovider) - Creates a new SSO provider
- [`authCreateThirdPartyAuth`](docs/sdks/auth/README.md#createthirdpartyauth) - Creates a new third-party auth integration
- [`authDeleteSsoProvider`](docs/sdks/auth/README.md#deletessoprovider) - Removes a SSO provider by its UUID
- [`authDeleteThirdPartyAuth`](docs/sdks/auth/README.md#deletethirdpartyauth) - [Alpha] Removes a third-party auth integration
- [`authGetConfig`](docs/sdks/auth/README.md#getconfig) - Gets project's auth config
- [`authGetSsoProvider`](docs/sdks/auth/README.md#getssoprovider) - Gets a SSO provider by its UUID
- [`authGetThirdPartyAuth`](docs/sdks/auth/README.md#getthirdpartyauth) - [Alpha] Get a third-party integration
- [`authListSsoProviders`](docs/sdks/auth/README.md#listssoproviders) - Lists all SSO providers
- [`authListThirdPartyAuth`](docs/sdks/auth/README.md#listthirdpartyauth) - [Alpha] Lists all third-party auth integrations
- [`authUpdateAuthConfig`](docs/sdks/auth/README.md#updateauthconfig) - Updates a project's auth config
- [`authUpdateSsoProvider`](docs/sdks/auth/README.md#updatessoprovider) - Updates a SSO provider by its UUID
- [`databaseDisableReadonlyModeTemporarily`](docs/sdks/database/README.md#disablereadonlymodetemporarily) - Disables project's readonly mode for the next 15 minutes
- [`databaseEnableWebhook`](docs/sdks/database/README.md#enablewebhook) - [Beta] Enables Database Webhooks on the project
- [`databaseGetPgbouncerConfig`](docs/sdks/database/README.md#getpgbouncerconfig) - Get project's pgbouncer config
- [`databaseGetPostgresConfig`](docs/sdks/database/README.md#getpostgresconfig) - Gets project's Postgres config
- [`databaseGetReadonlyModeStatus`](docs/sdks/database/README.md#getreadonlymodestatus) - Returns project's readonly mode status
- [`databaseGetSnippet`](docs/sdks/database/README.md#getsnippet) - Gets a specific SQL snippet
- [`databaseGetSSLEnforcement`](docs/sdks/database/README.md#getsslenforcement) - [Beta] Get project's SSL enforcement configuration.
- [`databaseGetSupavisorConfig`](docs/sdks/database/README.md#getsupavisorconfig) - Gets project's supavisor config
- [`databaseListBackups`](docs/sdks/database/README.md#listbackups) - Lists all backups
- [`databaseRemoveReadReplica`](docs/sdks/database/README.md#removereadreplica) - [Beta] Remove a read replica
- [`databaseRestorePitrBackup`](docs/sdks/database/README.md#restorepitrbackup) - Restores a PITR backup for a database
- [`databaseRunQuery`](docs/sdks/database/README.md#runquery) - [Beta] Run sql query
- [`databaseSetupReadReplica`](docs/sdks/database/README.md#setupreadreplica) - [Beta] Set up a read replica
- [`databasesGenerateTypescriptTypes`](docs/sdks/databases/README.md#generatetypescripttypes) - Generate TypeScript types
- [`databasesUpdatePostgresConfig`](docs/sdks/databases/README.md#updatepostgresconfig) - Updates project's Postgres config
- [`databasesUpdateSupavisorConfig`](docs/sdks/databases/README.md#updatesupavisorconfig) - Updates project's supavisor config
- [`databaseUpdateSSLEnforcement`](docs/sdks/database/README.md#updatesslenforcement) - [Beta] Update project's SSL enforcement configuration.
- [`domainsActivateCustomHostname`](docs/sdks/domains/README.md#activatecustomhostname) - [Beta] Activates a custom hostname for a project.
- [`domainsActivateVanitySubdomain`](docs/sdks/domains/README.md#activatevanitysubdomain) - [Beta] Activates a vanity subdomain for a project.
- [`domainsCheckVanitySubdomainAvailability`](docs/sdks/domains/README.md#checkvanitysubdomainavailability) - [Beta] Checks vanity subdomain availability
- [`domainsDeleteHostnameConfig`](docs/sdks/domains/README.md#deletehostnameconfig) - [Beta] Deletes a project's custom hostname configuration
- [`domainsDeleteVanitySubdomain`](docs/sdks/domains/README.md#deletevanitysubdomain) - [Beta] Deletes a project's vanity subdomain configuration
- [`domainsGetHostnameConfig`](docs/sdks/domains/README.md#gethostnameconfig) - [Beta] Gets project's custom hostname config
- [`domainsGetVanitySubdomain`](docs/sdks/domains/README.md#getvanitysubdomain) - [Beta] Gets current vanity subdomain config
- [`domainsInitializeHostnameConfig`](docs/sdks/domains/README.md#initializehostnameconfig) - [Beta] Updates project's custom hostname configuration
- [`domainsVerifyDnsConfig`](docs/sdks/domains/README.md#verifydnsconfig) - [Beta] Attempts to verify the DNS configuration for project's custom hostname configuration
- [`edgeFunctionsCreate`](docs/sdks/edgefunctions/README.md#create) - Create a function
- [`edgeFunctionsDeleteFunction`](docs/sdks/edgefunctions/README.md#deletefunction) - Delete a function
- [`edgeFunctionsGetFunction`](docs/sdks/edgefunctions/README.md#getfunction) - Retrieve a function
- [`edgeFunctionsGetFunctionBody`](docs/sdks/edgefunctions/README.md#getfunctionbody) - Retrieve a function body
- [`edgeFunctionsList`](docs/sdks/edgefunctions/README.md#list) - List all functions
- [`edgeFunctionsUpdateFunction`](docs/sdks/edgefunctions/README.md#updatefunction) - Update a function
- [`environmentsCreateBranch`](docs/sdks/environments/README.md#createbranch) - Create a database branch
- [`environmentsDeleteBranch`](docs/sdks/environments/README.md#deletebranch) - Delete a database branch
- [`environmentsDisablePreviewBranching`](docs/sdks/environments/README.md#disablepreviewbranching) - Disables preview branching
- [`environmentsGetBranchConfig`](docs/sdks/environments/README.md#getbranchconfig) - Get database branch config
- [`environmentsListBranches`](docs/sdks/environments/README.md#listbranches) - List all database branches
- [`environmentsResetBranch`](docs/sdks/environments/README.md#resetbranch) - Resets a database branch
- [`environmentsUpdateBranchConfig`](docs/sdks/environments/README.md#updatebranchconfig) - Update database branch config
- [`oauthAuthorizeUser`](docs/sdks/oauth/README.md#authorizeuser) - [Beta] Authorize user through oauth
- [`oauthExchangeToken`](docs/sdks/oauth/README.md#exchangetoken) - [Beta] Exchange auth code for user's access and refresh token
- [`organizationsCreate`](docs/sdks/organizations/README.md#create) - Create an organization
- [`organizationsGet`](docs/sdks/organizations/README.md#get) - Gets information about the organization
- [`organizationsList`](docs/sdks/organizations/README.md#list) - List all organizations
- [`organizationsListMembers`](docs/sdks/organizations/README.md#listmembers) - List members of an organization
- [`projectsCreate`](docs/sdks/projects/README.md#create) - Create a project
- [`projectsDelete`](docs/sdks/projects/README.md#delete) - Deletes the given project
- [`projectsDeleteNetworkBans`](docs/sdks/projects/README.md#deletenetworkbans) - [Beta] Remove network bans.
- [`projectsGet`](docs/sdks/projects/README.md#get) - Gets a specific project that belongs to the authenticated user
- [`projectsGetHealth`](docs/sdks/projects/README.md#gethealth) - Gets project's service health status
- [`projectsGetNetworkRestrictions`](docs/sdks/projects/README.md#getnetworkrestrictions) - [Beta] Gets project's network restrictions
- [`projectsGetUpgradeEligibility`](docs/sdks/projects/README.md#getupgradeeligibility) - [Beta] Returns the project's eligibility for upgrades
- [`projectsGetUpgradeStatus`](docs/sdks/projects/README.md#getupgradestatus) - [Beta] Gets the latest status of the project's upgrade
- [`projectsList`](docs/sdks/projects/README.md#list) - List all projects
- [`projectsListNetworkBans`](docs/sdks/projects/README.md#listnetworkbans) - [Beta] Gets project's network bans
- [`projectsUpdateNetworkRestrictions`](docs/sdks/projects/README.md#updatenetworkrestrictions) - [Beta] Updates project's network restrictions
- [`projectsUpgradePostgresVersion`](docs/sdks/projects/README.md#upgradepostgresversion) - [Beta] Upgrades the project's Postgres version
- [`restGetConfig`](docs/sdks/rest/README.md#getconfig) - Gets project's postgrest config
- [`restUpdatePostgrestConfig`](docs/sdks/rest/README.md#updatepostgrestconfig) - Updates project's postgrest config
- [`secretsBulkCreate`](docs/sdks/secrets/README.md#bulkcreate) - Bulk create secrets
- [`secretsBulkDelete`](docs/sdks/secrets/README.md#bulkdelete) - Bulk delete secrets
- [`secretsCreateApiKey`](docs/sdks/secrets/README.md#createapikey) - [Alpha] Creates a new API key for the project
- [`secretsDeleteApiKey`](docs/sdks/secrets/README.md#deleteapikey) - [Alpha] Deletes an API key for the project
- [`secretsGetApiKey`](docs/sdks/secrets/README.md#getapikey) - [Alpha] Get API key
- [`secretsGetApiKeys`](docs/sdks/secrets/README.md#getapikeys) - Get project api keys
- [`secretsGetPGSodiumConfig`](docs/sdks/secrets/README.md#getpgsodiumconfig) - [Beta] Gets project's pgsodium config
- [`secretsList`](docs/sdks/secrets/README.md#list) - List all secrets
- [`secretsUpdateApiKey`](docs/sdks/secrets/README.md#updateapikey) - [Alpha] Updates an API key for the project
- [`secretsUpdatePgsodiumConfig`](docs/sdks/secrets/README.md#updatepgsodiumconfig) - [Beta] Updates project's pgsodium config. Updating the root_key can cause all data encrypted with the older key to become inaccessible.
- [`snippetsListAll`](docs/sdks/snippets/README.md#listall) - Lists SQL snippets for the logged in user
- [`storageGetConfig`](docs/sdks/storage/README.md#getconfig) - Gets project's storage config
- [`storageListBuckets`](docs/sdks/storage/README.md#listbuckets) - Lists all buckets
- [`storagePatchConfig`](docs/sdks/storage/README.md#patchconfig) - Updates project's storage config

</details>
<!-- End Standalone functions [standalone-funcs] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await supabase.environments.getBranchConfig({
    branchId: "<id>",
  }, {
    retries: {
      strategy: "backoff",
      backoff: {
        initialInterval: 1,
        maxInterval: 50,
        exponent: 1.1,
        maxElapsedTime: 100,
      },
      retryConnectionErrors: false,
    },
  });

  // Handle the result
  console.log(result);
}

run();

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  retryConfig: {
    strategy: "backoff",
    backoff: {
      initialInterval: 1,
      maxInterval: 50,
      exponent: 1.1,
      maxElapsedTime: 100,
    },
    retryConnectionErrors: false,
  },
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await supabase.environments.getBranchConfig({
    branchId: "<id>",
  });

  // Handle the result
  console.log(result);
}

run();

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

If the request fails due to, for example 4XX or 5XX status codes, it will throw a `APIError`.

| Error Type      | Status Code | Content Type |
| --------------- | ----------- | ------------ |
| errors.APIError | 4XX, 5XX    | \*/\*        |

```typescript
import { Supabase } from "supabase";
import { SDKValidationError } from "supabase/models/errors";

const supabase = new Supabase({
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  let result;
  try {
    result = await supabase.environments.getBranchConfig({
      branchId: "<id>",
    });

    // Handle the result
    console.log(result);
  } catch (err) {
    switch (true) {
      // The server response does not match the expected SDK schema
      case (err instanceof SDKValidationError):
        {
          // Pretty-print will provide a human-readable multi-line error message
          console.error(err.pretty());
          // Raw value may also be inspected
          console.error(err.rawValue);
          return;
        }
        apierror.js;
      // Server returned an error status code or an unknown content type
      case (err instanceof APIError): {
        console.error(err.statusCode);
        console.error(err.rawResponse.body);
        return;
      }
      default: {
        // Other errors such as network errors, see HTTPClientErrors for more details
        throw err;
      }
    }
  }
}

run();

```

Validation errors can also occur when either method arguments or data returned from the server do not match the expected format. The `SDKValidationError` that is thrown as a result will capture the raw value that failed validation in an attribute called `rawValue`. Additionally, a `pretty()` method is available on this error that can be used to log a nicely formatted multi-line string since validation errors can list many issues and the plain error string may be difficult read when debugging.

In some rare cases, the SDK can fail to get a response from the server or even make the request due to unexpected circumstances such as network conditions. These types of errors are captured in the `models/errors/httpclienterrors.ts` module:

| HTTP Client Error                                    | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- |
| RequestAbortedError                                  | HTTP request was aborted by the client               |
| RequestTimeoutError                                  | HTTP request timed out due to an AbortSignal signal  |
| ConnectionError                                      | HTTP client was unable to make a request to a server |
| InvalidRequestError                                  | Any input used to create a request is invalid        |
| UnexpectedClientError                                | Unrecognised or unexpected error                     |
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally by passing a URL to the `serverURL: string` optional parameter when initializing the SDK client instance. For example:
```typescript
import { Supabase } from "supabase";

const supabase = new Supabase({
  serverURL: "https://api.supabase.com",
  bearer: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await supabase.environments.getBranchConfig({
    branchId: "<id>",
  });

  // Handle the result
  console.log(result);
}

run();

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The TypeScript SDK makes API calls using an `HTTPClient` that wraps the native
[Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API). This
client is a thin wrapper around `fetch` and provides the ability to attach hooks
around the request lifecycle that can be used to modify the request or handle
errors and response.

The `HTTPClient` constructor takes an optional `fetcher` argument that can be
used to integrate a third-party HTTP client or when writing tests to mock out
the HTTP client and feed in fixtures.

The following example shows how to use the `"beforeRequest"` hook to to add a
custom header and a timeout to requests and how to use the `"requestError"` hook
to log errors:

```typescript
import { Supabase } from "supabase";
import { HTTPClient } from "supabase/lib/http";

const httpClient = new HTTPClient({
  // fetcher takes a function that has the same signature as native `fetch`.
  fetcher: (request) => {
    return fetch(request);
  }
});

httpClient.addHook("beforeRequest", (request) => {
  const nextRequest = new Request(request, {
    signal: request.signal || AbortSignal.timeout(5000)
  });

  nextRequest.headers.set("x-custom-header", "custom value");

  return nextRequest;
});

httpClient.addHook("requestError", (error, request) => {
  console.group("Request Error");
  console.log("Reason:", `${error}`);
  console.log("Endpoint:", `${request.method} ${request.url}`);
  console.groupEnd();
});

const sdk = new Supabase({ httpClient });
```
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Debugging [debug] -->
## Debugging

You can setup your SDK to emit debug logs for SDK requests and responses.

You can pass a logger that matches `console`'s interface as an SDK option.

> [!WARNING]
> Beware that debug logging will reveal secrets, like API tokens in headers, in log messages printed to a console or files. It's recommended to use this feature only during local development and not in production.

```typescript
import { Supabase } from "supabase";

const sdk = new Supabase({ debugLogger: console });
```
<!-- End Debugging [debug] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=supabase&utm_campaign=typescript)
