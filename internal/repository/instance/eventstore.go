package instance

import (
	"github.com/zitadel/zitadel/internal/eventstore"
)

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, DefaultOrgSetEventType, DefaultOrgSetMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, ProjectSetEventType, ProjectSetMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, ConsoleSetEventType, ConsoleSetMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DefaultLanguageSetEventType, DefaultLanguageSetMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SecretGeneratorAddedEventType, SecretGeneratorAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SecretGeneratorChangedEventType, SecretGeneratorChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SecretGeneratorRemovedEventType, SecretGeneratorRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigAddedEventType, eventstore.GenericEventMapper[SMTPConfigAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigChangedEventType, eventstore.GenericEventMapper[SMTPConfigChangedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigActivatedEventType, eventstore.GenericEventMapper[SMTPConfigActivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigDeactivatedEventType, eventstore.GenericEventMapper[SMTPConfigDeactivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigPasswordChangedEventType, eventstore.GenericEventMapper[SMTPConfigPasswordChangedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigHTTPAddedEventType, eventstore.GenericEventMapper[SMTPConfigHTTPAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigHTTPChangedEventType, eventstore.GenericEventMapper[SMTPConfigHTTPChangedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMTPConfigRemovedEventType, eventstore.GenericEventMapper[SMTPConfigRemovedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigTwilioAddedEventType, eventstore.GenericEventMapper[SMSConfigTwilioAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigTwilioChangedEventType, eventstore.GenericEventMapper[SMSConfigTwilioChangedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigTwilioTokenChangedEventType, eventstore.GenericEventMapper[SMSConfigTwilioTokenChangedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigHTTPAddedEventType, eventstore.GenericEventMapper[SMSConfigHTTPAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigHTTPChangedEventType, eventstore.GenericEventMapper[SMSConfigHTTPChangedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigTwilioActivatedEventType, eventstore.GenericEventMapper[SMSConfigTwilioActivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigTwilioDeactivatedEventType, eventstore.GenericEventMapper[SMSConfigTwilioDeactivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigTwilioRemovedEventType, eventstore.GenericEventMapper[SMSConfigTwilioRemovedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigActivatedEventType, eventstore.GenericEventMapper[SMSConfigActivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigDeactivatedEventType, eventstore.GenericEventMapper[SMSConfigDeactivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SMSConfigRemovedEventType, eventstore.GenericEventMapper[SMSConfigRemovedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, DebugNotificationProviderFileAddedEventType, DebugNotificationProviderFileAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DebugNotificationProviderFileChangedEventType, DebugNotificationProviderFileChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DebugNotificationProviderFileRemovedEventType, DebugNotificationProviderFileRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DebugNotificationProviderLogAddedEventType, DebugNotificationProviderLogAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DebugNotificationProviderLogChangedEventType, DebugNotificationProviderLogChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DebugNotificationProviderLogRemovedEventType, DebugNotificationProviderLogRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OIDCSettingsAddedEventType, OIDCSettingsAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OIDCSettingsChangedEventType, OIDCSettingsChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SecurityPolicySetEventType, SecurityPolicySetEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyAddedEventType, LabelPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyChangedEventType, LabelPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyActivatedEventType, LabelPolicyActivatedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyLogoAddedEventType, LabelPolicyLogoAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyLogoRemovedEventType, LabelPolicyLogoRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyIconAddedEventType, LabelPolicyIconAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyIconRemovedEventType, LabelPolicyIconRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyLogoDarkAddedEventType, LabelPolicyLogoDarkAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyLogoDarkRemovedEventType, LabelPolicyLogoDarkRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyIconDarkAddedEventType, LabelPolicyIconDarkAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyIconDarkRemovedEventType, LabelPolicyIconDarkRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyFontAddedEventType, LabelPolicyFontAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyFontRemovedEventType, LabelPolicyFontRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LabelPolicyAssetsRemovedEventType, LabelPolicyAssetsRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyAddedEventType, LoginPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyChangedEventType, LoginPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DomainPolicyAddedEventType, DomainPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, DomainPolicyChangedEventType, DomainPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PasswordAgePolicyAddedEventType, PasswordAgePolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PasswordAgePolicyChangedEventType, PasswordAgePolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PasswordComplexityPolicyAddedEventType, PasswordComplexityPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PasswordComplexityPolicyChangedEventType, PasswordComplexityPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LockoutPolicyAddedEventType, LockoutPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LockoutPolicyChangedEventType, LockoutPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PrivacyPolicyAddedEventType, PrivacyPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PrivacyPolicyChangedEventType, PrivacyPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MemberAddedEventType, MemberAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MemberChangedEventType, MemberChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MemberRemovedEventType, MemberRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MemberCascadeRemovedEventType, MemberCascadeRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPConfigAddedEventType, IDPConfigAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPConfigChangedEventType, IDPConfigChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPConfigRemovedEventType, IDPConfigRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPConfigDeactivatedEventType, IDPConfigDeactivatedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPConfigReactivatedEventType, IDPConfigReactivatedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPOIDCConfigAddedEventType, IDPOIDCConfigAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPOIDCConfigChangedEventType, IDPOIDCConfigChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPJWTConfigAddedEventType, IDPJWTConfigAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPJWTConfigChangedEventType, IDPJWTConfigChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OAuthIDPAddedEventType, OAuthIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OAuthIDPChangedEventType, OAuthIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OIDCIDPAddedEventType, OIDCIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OIDCIDPChangedEventType, OIDCIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OIDCIDPMigratedAzureADEventType, OIDCIDPMigratedAzureADEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, OIDCIDPMigratedGoogleEventType, OIDCIDPMigratedGoogleEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, JWTIDPAddedEventType, JWTIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, JWTIDPChangedEventType, JWTIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, AzureADIDPAddedEventType, AzureADIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, AzureADIDPChangedEventType, AzureADIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitHubIDPAddedEventType, GitHubIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitHubIDPChangedEventType, GitHubIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitHubEnterpriseIDPAddedEventType, GitHubEnterpriseIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitHubEnterpriseIDPChangedEventType, GitHubEnterpriseIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitLabIDPAddedEventType, GitLabIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitLabIDPChangedEventType, GitLabIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitLabSelfHostedIDPAddedEventType, GitLabSelfHostedIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GitLabSelfHostedIDPChangedEventType, GitLabSelfHostedIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GoogleIDPAddedEventType, GoogleIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, GoogleIDPChangedEventType, GoogleIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LDAPIDPAddedEventType, LDAPIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LDAPIDPChangedEventType, LDAPIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, AppleIDPAddedEventType, AppleIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, AppleIDPChangedEventType, AppleIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SAMLIDPAddedEventType, SAMLIDPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, SAMLIDPChangedEventType, SAMLIDPChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, IDPRemovedEventType, IDPRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyIDPProviderAddedEventType, IdentityProviderAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyIDPProviderRemovedEventType, IdentityProviderRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyIDPProviderCascadeRemovedEventType, IdentityProviderCascadeRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicySecondFactorAddedEventType, SecondFactorAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicySecondFactorRemovedEventType, SecondFactorRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyMultiFactorAddedEventType, MultiFactorAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, LoginPolicyMultiFactorRemovedEventType, MultiFactorRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MailTemplateAddedEventType, MailTemplateAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MailTemplateChangedEventType, MailTemplateChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MailTextAddedEventType, MailTextAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MailTextChangedEventType, MailTextChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, CustomTextSetEventType, CustomTextSetEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, CustomTextRemovedEventType, CustomTextRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, CustomTextTemplateRemovedEventType, CustomTextTemplateRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, InstanceDomainAddedEventType, DomainAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, InstanceDomainPrimarySetEventType, DomainPrimarySetEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, InstanceDomainRemovedEventType, DomainRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, InstanceAddedEventType, InstanceAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, InstanceChangedEventType, InstanceChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, InstanceRemovedEventType, InstanceRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, NotificationPolicyAddedEventType, NotificationPolicyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, NotificationPolicyChangedEventType, NotificationPolicyChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, TrustedDomainAddedEventType, eventstore.GenericEventMapper[TrustedDomainAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, TrustedDomainRemovedEventType, eventstore.GenericEventMapper[TrustedDomainRemovedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HostedLoginTranslationSet, HostedLoginTranslationSetEventMapper)
}
