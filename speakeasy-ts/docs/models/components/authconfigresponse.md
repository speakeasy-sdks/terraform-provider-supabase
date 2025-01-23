# AuthConfigResponse

## Example Usage

```typescript
import { AuthConfigResponse } from "supabase/models/components";

let value: AuthConfigResponse = {
  apiMaxRequestDuration: 453543,
  dbMaxPoolSize: 722056,
  jwtExp: 866383,
  mailerOtpExp: 975522,
  mailerOtpLength: 855804,
  mfaMaxEnrolledFactors: 11714,
  mfaPhoneOtpLength: 359978,
  mfaPhoneMaxFrequency: 729991,
  passwordMinLength: 171629,
  rateLimitAnonymousUsers: 521037,
  rateLimitEmailSent: 54338,
  rateLimitSmsSent: 199996,
  rateLimitTokenRefresh: 18521,
  rateLimitVerify: 793698,
  rateLimitOtp: 223924,
  securityRefreshTokenReuseInterval: 345352,
  sessionsInactivityTimeout: 928082,
  sessionsTimebox: 704415,
  smsMaxFrequency: 31838,
  smsOtpExp: 164694,
  smsOtpLength: 621479,
  smtpMaxFrequency: 577229,
  disableSignup: false,
  externalAnonymousUsersEnabled: false,
  externalAppleAdditionalClientIds: "<value>",
  externalAppleClientId: "<id>",
  externalAppleEnabled: false,
  externalAppleSecret: "<value>",
  externalAzureClientId: "<id>",
  externalAzureEnabled: false,
  externalAzureSecret: "<value>",
  externalAzureUrl: "https://vivid-obedience.name/",
  externalBitbucketClientId: "<id>",
  externalBitbucketEnabled: false,
  externalBitbucketSecret: "<value>",
  externalDiscordClientId: "<id>",
  externalDiscordEnabled: false,
  externalDiscordSecret: "<value>",
  externalEmailEnabled: false,
  externalFacebookClientId: "<id>",
  externalFacebookEnabled: false,
  externalFacebookSecret: "<value>",
  externalFigmaClientId: "<id>",
  externalFigmaEnabled: false,
  externalFigmaSecret: "<value>",
  externalGithubClientId: "<id>",
  externalGithubEnabled: false,
  externalGithubSecret: "<value>",
  externalGitlabClientId: "<id>",
  externalGitlabEnabled: false,
  externalGitlabSecret: "<value>",
  externalGitlabUrl: "https://shallow-exploration.info",
  externalGoogleAdditionalClientIds: "<value>",
  externalGoogleClientId: "<id>",
  externalGoogleEnabled: false,
  externalGoogleSecret: "<value>",
  externalGoogleSkipNonceCheck: false,
  externalKakaoClientId: "<id>",
  externalKakaoEnabled: false,
  externalKakaoSecret: "<value>",
  externalKeycloakClientId: "<id>",
  externalKeycloakEnabled: false,
  externalKeycloakSecret: "<value>",
  externalKeycloakUrl: "https://descriptive-unblinking.net/",
  externalLinkedinOidcClientId: "<id>",
  externalLinkedinOidcEnabled: false,
  externalLinkedinOidcSecret: "<value>",
  externalSlackOidcClientId: "<id>",
  externalSlackOidcEnabled: false,
  externalSlackOidcSecret: "<value>",
  externalNotionClientId: "<id>",
  externalNotionEnabled: false,
  externalNotionSecret: "<value>",
  externalPhoneEnabled: false,
  externalSlackClientId: "<id>",
  externalSlackEnabled: false,
  externalSlackSecret: "<value>",
  externalSpotifyClientId: "<id>",
  externalSpotifyEnabled: false,
  externalSpotifySecret: "<value>",
  externalTwitchClientId: "<id>",
  externalTwitchEnabled: false,
  externalTwitchSecret: "<value>",
  externalTwitterClientId: "<id>",
  externalTwitterEnabled: false,
  externalTwitterSecret: "<value>",
  externalWorkosClientId: "<id>",
  externalWorkosEnabled: false,
  externalWorkosSecret: "<value>",
  externalWorkosUrl: "https://enchanted-disclosure.com/",
  externalZoomClientId: "<id>",
  externalZoomEnabled: false,
  externalZoomSecret: "<value>",
  hookCustomAccessTokenEnabled: false,
  hookCustomAccessTokenUri: "https://general-possession.info/",
  hookCustomAccessTokenSecrets: "<value>",
  hookMfaVerificationAttemptEnabled: false,
  hookMfaVerificationAttemptUri: "https://alive-bathhouse.net/",
  hookMfaVerificationAttemptSecrets: "<value>",
  hookPasswordVerificationAttemptEnabled: false,
  hookPasswordVerificationAttemptUri: "https://muted-technician.org/",
  hookPasswordVerificationAttemptSecrets: "<value>",
  hookSendSmsEnabled: false,
  hookSendSmsUri: "https://radiant-dividend.com/",
  hookSendSmsSecrets: "<value>",
  hookSendEmailEnabled: false,
  hookSendEmailUri: "https://glittering-godparent.name",
  hookSendEmailSecrets: "<value>",
  mailerAllowUnverifiedEmailSignIns: false,
  mailerAutoconfirm: false,
  mailerSecureEmailChangeEnabled: false,
  mailerSubjectsConfirmation: "<value>",
  mailerSubjectsEmailChange: "<value>",
  mailerSubjectsInvite: "<value>",
  mailerSubjectsMagicLink: "<value>",
  mailerSubjectsReauthentication: "<value>",
  mailerSubjectsRecovery: "<value>",
  mailerTemplatesConfirmationContent: "<value>",
  mailerTemplatesEmailChangeContent: "<value>",
  mailerTemplatesInviteContent: "<value>",
  mailerTemplatesMagicLinkContent: "<value>",
  mailerTemplatesReauthenticationContent: "<value>",
  mailerTemplatesRecoveryContent: "<value>",
  mfaTotpEnrollEnabled: false,
  mfaTotpVerifyEnabled: false,
  mfaPhoneEnrollEnabled: false,
  mfaPhoneVerifyEnabled: false,
  mfaWebAuthnEnrollEnabled: false,
  mfaWebAuthnVerifyEnabled: false,
  mfaPhoneTemplate: "<value>",
  passwordHibpEnabled: false,
  passwordRequiredCharacters: "<value>",
  refreshTokenRotationEnabled: false,
  samlEnabled: false,
  samlExternalUrl: "https://practical-supplier.biz",
  samlAllowEncryptedAssertions: false,
  securityCaptchaEnabled: false,
  securityCaptchaProvider: "<value>",
  securityCaptchaSecret: "<value>",
  securityManualLinkingEnabled: false,
  securityUpdatePasswordRequireReauthentication: false,
  sessionsSinglePerUser: false,
  sessionsTags: "<value>",
  siteUrl: "https://dependent-valentine.net",
  smsAutoconfirm: false,
  smsMessagebirdAccessKey: "<value>",
  smsMessagebirdOriginator: "<value>",
  smsProvider: "<value>",
  smsTemplate: "<value>",
  smsTestOtp: "<value>",
  smsTestOtpValidUntil: "<value>",
  smsTextlocalApiKey: "<value>",
  smsTextlocalSender: "<value>",
  smsTwilioAccountSid: "<id>",
  smsTwilioAuthToken: "<value>",
  smsTwilioContentSid: "<id>",
  smsTwilioMessageServiceSid: "<id>",
  smsTwilioVerifyAccountSid: "<id>",
  smsTwilioVerifyAuthToken: "<value>",
  smsTwilioVerifyMessageServiceSid: "<id>",
  smsVonageApiKey: "<value>",
  smsVonageApiSecret: "<value>",
  smsVonageFrom: "<value>",
  smtpAdminEmail: "<value>",
  smtpHost: "<value>",
  smtpPass: "<value>",
  smtpPort: "<value>",
  smtpSenderName: "<value>",
  smtpUser: "<value>",
  uriAllowList: "<value>",
};
```

## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `apiMaxRequestDuration`                         | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `dbMaxPoolSize`                                 | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `jwtExp`                                        | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerOtpExp`                                  | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerOtpLength`                               | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `mfaMaxEnrolledFactors`                         | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `mfaPhoneOtpLength`                             | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `mfaPhoneMaxFrequency`                          | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `passwordMinLength`                             | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `rateLimitAnonymousUsers`                       | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `rateLimitEmailSent`                            | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `rateLimitSmsSent`                              | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `rateLimitTokenRefresh`                         | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `rateLimitVerify`                               | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `rateLimitOtp`                                  | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `securityRefreshTokenReuseInterval`             | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `sessionsInactivityTimeout`                     | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `sessionsTimebox`                               | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsMaxFrequency`                               | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsOtpExp`                                     | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsOtpLength`                                  | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpMaxFrequency`                              | *number*                                        | :heavy_check_mark:                              | N/A                                             |
| `disableSignup`                                 | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalAnonymousUsersEnabled`                 | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalAppleAdditionalClientIds`              | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalAppleClientId`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalAppleEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalAppleSecret`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalAzureClientId`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalAzureEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalAzureSecret`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalAzureUrl`                              | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalBitbucketClientId`                     | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalBitbucketEnabled`                      | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalBitbucketSecret`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalDiscordClientId`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalDiscordEnabled`                        | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalDiscordSecret`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalEmailEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalFacebookClientId`                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalFacebookEnabled`                       | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalFacebookSecret`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalFigmaClientId`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalFigmaEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalFigmaSecret`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGithubClientId`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGithubEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalGithubSecret`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGitlabClientId`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGitlabEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalGitlabSecret`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGitlabUrl`                             | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGoogleAdditionalClientIds`             | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGoogleClientId`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGoogleEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalGoogleSecret`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalGoogleSkipNonceCheck`                  | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalKakaoClientId`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalKakaoEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalKakaoSecret`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalKeycloakClientId`                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalKeycloakEnabled`                       | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalKeycloakSecret`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalKeycloakUrl`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalLinkedinOidcClientId`                  | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalLinkedinOidcEnabled`                   | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalLinkedinOidcSecret`                    | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalSlackOidcClientId`                     | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalSlackOidcEnabled`                      | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalSlackOidcSecret`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalNotionClientId`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalNotionEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalNotionSecret`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalPhoneEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalSlackClientId`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalSlackEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalSlackSecret`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalSpotifyClientId`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalSpotifyEnabled`                        | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalSpotifySecret`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalTwitchClientId`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalTwitchEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalTwitchSecret`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalTwitterClientId`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalTwitterEnabled`                        | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalTwitterSecret`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalWorkosClientId`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalWorkosEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalWorkosSecret`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalWorkosUrl`                             | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalZoomClientId`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `externalZoomEnabled`                           | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `externalZoomSecret`                            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookCustomAccessTokenEnabled`                  | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `hookCustomAccessTokenUri`                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookCustomAccessTokenSecrets`                  | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookMfaVerificationAttemptEnabled`             | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `hookMfaVerificationAttemptUri`                 | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookMfaVerificationAttemptSecrets`             | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookPasswordVerificationAttemptEnabled`        | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `hookPasswordVerificationAttemptUri`            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookPasswordVerificationAttemptSecrets`        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookSendSmsEnabled`                            | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `hookSendSmsUri`                                | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookSendSmsSecrets`                            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookSendEmailEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `hookSendEmailUri`                              | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `hookSendEmailSecrets`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerAllowUnverifiedEmailSignIns`             | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mailerAutoconfirm`                             | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mailerSecureEmailChangeEnabled`                | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mailerSubjectsConfirmation`                    | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerSubjectsEmailChange`                     | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerSubjectsInvite`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerSubjectsMagicLink`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerSubjectsReauthentication`                | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerSubjectsRecovery`                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerTemplatesConfirmationContent`            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerTemplatesEmailChangeContent`             | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerTemplatesInviteContent`                  | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerTemplatesMagicLinkContent`               | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerTemplatesReauthenticationContent`        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mailerTemplatesRecoveryContent`                | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `mfaTotpEnrollEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mfaTotpVerifyEnabled`                          | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mfaPhoneEnrollEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mfaPhoneVerifyEnabled`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mfaWebAuthnEnrollEnabled`                      | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mfaWebAuthnVerifyEnabled`                      | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `mfaPhoneTemplate`                              | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `passwordHibpEnabled`                           | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `passwordRequiredCharacters`                    | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `refreshTokenRotationEnabled`                   | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `samlEnabled`                                   | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `samlExternalUrl`                               | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `samlAllowEncryptedAssertions`                  | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `securityCaptchaEnabled`                        | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `securityCaptchaProvider`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `securityCaptchaSecret`                         | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `securityManualLinkingEnabled`                  | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `securityUpdatePasswordRequireReauthentication` | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `sessionsSinglePerUser`                         | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `sessionsTags`                                  | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `siteUrl`                                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsAutoconfirm`                                | *boolean*                                       | :heavy_check_mark:                              | N/A                                             |
| `smsMessagebirdAccessKey`                       | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsMessagebirdOriginator`                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsProvider`                                   | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTemplate`                                   | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTestOtp`                                    | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTestOtpValidUntil`                          | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTextlocalApiKey`                            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTextlocalSender`                            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioAccountSid`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioAuthToken`                            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioContentSid`                           | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioMessageServiceSid`                    | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioVerifyAccountSid`                     | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioVerifyAuthToken`                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsTwilioVerifyMessageServiceSid`              | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsVonageApiKey`                               | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsVonageApiSecret`                            | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smsVonageFrom`                                 | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpAdminEmail`                                | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpHost`                                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpPass`                                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpPort`                                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpSenderName`                                | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `smtpUser`                                      | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `uriAllowList`                                  | *string*                                        | :heavy_check_mark:                              | N/A                                             |