package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishingable 
type OnPremisesPublishingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateUrl()(*string)
    GetApplicationServerTimeout()(*string)
    GetApplicationType()(*string)
    GetExternalAuthenticationType()(*ExternalAuthenticationType)
    GetExternalUrl()(*string)
    GetInternalUrl()(*string)
    GetIsBackendCertificateValidationEnabled()(*bool)
    GetIsHttpOnlyCookieEnabled()(*bool)
    GetIsOnPremPublishingEnabled()(*bool)
    GetIsPersistentCookieEnabled()(*bool)
    GetIsSecureCookieEnabled()(*bool)
    GetIsStateSessionEnabled()(*bool)
    GetIsTranslateHostHeaderEnabled()(*bool)
    GetIsTranslateLinksInBodyEnabled()(*bool)
    GetOnPremisesApplicationSegments()([]OnPremisesApplicationSegmentable)
    GetSingleSignOnSettings()(OnPremisesPublishingSingleSignOnable)
    GetUseAlternateUrlForTranslationAndRedirect()(*bool)
    GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable)
    GetVerifiedCustomDomainKeyCredential()(KeyCredentialable)
    GetVerifiedCustomDomainPasswordCredential()(PasswordCredentialable)
    SetAlternateUrl(value *string)()
    SetApplicationServerTimeout(value *string)()
    SetApplicationType(value *string)()
    SetExternalAuthenticationType(value *ExternalAuthenticationType)()
    SetExternalUrl(value *string)()
    SetInternalUrl(value *string)()
    SetIsBackendCertificateValidationEnabled(value *bool)()
    SetIsHttpOnlyCookieEnabled(value *bool)()
    SetIsOnPremPublishingEnabled(value *bool)()
    SetIsPersistentCookieEnabled(value *bool)()
    SetIsSecureCookieEnabled(value *bool)()
    SetIsStateSessionEnabled(value *bool)()
    SetIsTranslateHostHeaderEnabled(value *bool)()
    SetIsTranslateLinksInBodyEnabled(value *bool)()
    SetOnPremisesApplicationSegments(value []OnPremisesApplicationSegmentable)()
    SetSingleSignOnSettings(value OnPremisesPublishingSingleSignOnable)()
    SetUseAlternateUrlForTranslationAndRedirect(value *bool)()
    SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)()
    SetVerifiedCustomDomainKeyCredential(value KeyCredentialable)()
    SetVerifiedCustomDomainPasswordCredential(value PasswordCredentialable)()
}
