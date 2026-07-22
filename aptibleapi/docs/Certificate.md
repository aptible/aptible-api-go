# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CertificateBody** | **string** |  | 
**PrivateKey** | Pointer to **string** |  | [optional] 
**CommonName** | **string** |  | 
**NotBefore** | **string** |  | 
**NotAfter** | **string** |  | 
**IssuerCountry** | **string** |  | 
**IssuerOrganization** | **string** |  | 
**IssuerWebsite** | **NullableString** |  | 
**IssuerCommonName** | **string** |  | 
**SubjectCountry** | **NullableString** |  | 
**SubjectState** | **NullableString** |  | 
**SubjectLocale** | **NullableString** |  | 
**SubjectOrganization** | **NullableString** |  | 
**Acme** | **bool** |  | 
**LeafCertificate** | **string** |  | 
**CertificateChain** | **string** |  | 
**Sha256Fingerprint** | **string** |  | 
**Trusted** | **bool** |  | 
**SelfSigned** | **bool** |  | 
**CertificateArn** | **NullableString** |  | 
**SubjectAlternativeNames** | **[]string** |  | 
**PrivateKeyAlgorithm** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**CertificateLinks**](CertificateLinks.md) |  | [optional] 

## Methods

### NewCertificate

`func NewCertificate(id int32, metaType string, certificateBody string, commonName string, notBefore string, notAfter string, issuerCountry string, issuerOrganization string, issuerWebsite NullableString, issuerCommonName string, subjectCountry NullableString, subjectState NullableString, subjectLocale NullableString, subjectOrganization NullableString, acme bool, leafCertificate string, certificateChain string, sha256Fingerprint string, trusted bool, selfSigned bool, certificateArn NullableString, subjectAlternativeNames []string, privateKeyAlgorithm string, createdAt string, updatedAt string, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Certificate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Certificate) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Certificate) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Certificate) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCertificateBody

`func (o *Certificate) GetCertificateBody() string`

GetCertificateBody returns the CertificateBody field if non-nil, zero value otherwise.

### GetCertificateBodyOk

`func (o *Certificate) GetCertificateBodyOk() (*string, bool)`

GetCertificateBodyOk returns a tuple with the CertificateBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBody

`func (o *Certificate) SetCertificateBody(v string)`

SetCertificateBody sets CertificateBody field to given value.


### GetPrivateKey

`func (o *Certificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Certificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Certificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *Certificate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetCommonName

`func (o *Certificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *Certificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *Certificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetNotBefore

`func (o *Certificate) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *Certificate) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *Certificate) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *Certificate) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *Certificate) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *Certificate) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### GetIssuerCountry

`func (o *Certificate) GetIssuerCountry() string`

GetIssuerCountry returns the IssuerCountry field if non-nil, zero value otherwise.

### GetIssuerCountryOk

`func (o *Certificate) GetIssuerCountryOk() (*string, bool)`

GetIssuerCountryOk returns a tuple with the IssuerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCountry

`func (o *Certificate) SetIssuerCountry(v string)`

SetIssuerCountry sets IssuerCountry field to given value.


### GetIssuerOrganization

`func (o *Certificate) GetIssuerOrganization() string`

GetIssuerOrganization returns the IssuerOrganization field if non-nil, zero value otherwise.

### GetIssuerOrganizationOk

`func (o *Certificate) GetIssuerOrganizationOk() (*string, bool)`

GetIssuerOrganizationOk returns a tuple with the IssuerOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrganization

`func (o *Certificate) SetIssuerOrganization(v string)`

SetIssuerOrganization sets IssuerOrganization field to given value.


### GetIssuerWebsite

`func (o *Certificate) GetIssuerWebsite() string`

GetIssuerWebsite returns the IssuerWebsite field if non-nil, zero value otherwise.

### GetIssuerWebsiteOk

`func (o *Certificate) GetIssuerWebsiteOk() (*string, bool)`

GetIssuerWebsiteOk returns a tuple with the IssuerWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerWebsite

`func (o *Certificate) SetIssuerWebsite(v string)`

SetIssuerWebsite sets IssuerWebsite field to given value.


### SetIssuerWebsiteNil

`func (o *Certificate) SetIssuerWebsiteNil(b bool)`

 SetIssuerWebsiteNil sets the value for IssuerWebsite to be an explicit nil

### UnsetIssuerWebsite
`func (o *Certificate) UnsetIssuerWebsite()`

UnsetIssuerWebsite ensures that no value is present for IssuerWebsite, not even an explicit nil
### GetIssuerCommonName

`func (o *Certificate) GetIssuerCommonName() string`

GetIssuerCommonName returns the IssuerCommonName field if non-nil, zero value otherwise.

### GetIssuerCommonNameOk

`func (o *Certificate) GetIssuerCommonNameOk() (*string, bool)`

GetIssuerCommonNameOk returns a tuple with the IssuerCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCommonName

`func (o *Certificate) SetIssuerCommonName(v string)`

SetIssuerCommonName sets IssuerCommonName field to given value.


### GetSubjectCountry

`func (o *Certificate) GetSubjectCountry() string`

GetSubjectCountry returns the SubjectCountry field if non-nil, zero value otherwise.

### GetSubjectCountryOk

`func (o *Certificate) GetSubjectCountryOk() (*string, bool)`

GetSubjectCountryOk returns a tuple with the SubjectCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCountry

`func (o *Certificate) SetSubjectCountry(v string)`

SetSubjectCountry sets SubjectCountry field to given value.


### SetSubjectCountryNil

`func (o *Certificate) SetSubjectCountryNil(b bool)`

 SetSubjectCountryNil sets the value for SubjectCountry to be an explicit nil

### UnsetSubjectCountry
`func (o *Certificate) UnsetSubjectCountry()`

UnsetSubjectCountry ensures that no value is present for SubjectCountry, not even an explicit nil
### GetSubjectState

`func (o *Certificate) GetSubjectState() string`

GetSubjectState returns the SubjectState field if non-nil, zero value otherwise.

### GetSubjectStateOk

`func (o *Certificate) GetSubjectStateOk() (*string, bool)`

GetSubjectStateOk returns a tuple with the SubjectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectState

`func (o *Certificate) SetSubjectState(v string)`

SetSubjectState sets SubjectState field to given value.


### SetSubjectStateNil

`func (o *Certificate) SetSubjectStateNil(b bool)`

 SetSubjectStateNil sets the value for SubjectState to be an explicit nil

### UnsetSubjectState
`func (o *Certificate) UnsetSubjectState()`

UnsetSubjectState ensures that no value is present for SubjectState, not even an explicit nil
### GetSubjectLocale

`func (o *Certificate) GetSubjectLocale() string`

GetSubjectLocale returns the SubjectLocale field if non-nil, zero value otherwise.

### GetSubjectLocaleOk

`func (o *Certificate) GetSubjectLocaleOk() (*string, bool)`

GetSubjectLocaleOk returns a tuple with the SubjectLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectLocale

`func (o *Certificate) SetSubjectLocale(v string)`

SetSubjectLocale sets SubjectLocale field to given value.


### SetSubjectLocaleNil

`func (o *Certificate) SetSubjectLocaleNil(b bool)`

 SetSubjectLocaleNil sets the value for SubjectLocale to be an explicit nil

### UnsetSubjectLocale
`func (o *Certificate) UnsetSubjectLocale()`

UnsetSubjectLocale ensures that no value is present for SubjectLocale, not even an explicit nil
### GetSubjectOrganization

`func (o *Certificate) GetSubjectOrganization() string`

GetSubjectOrganization returns the SubjectOrganization field if non-nil, zero value otherwise.

### GetSubjectOrganizationOk

`func (o *Certificate) GetSubjectOrganizationOk() (*string, bool)`

GetSubjectOrganizationOk returns a tuple with the SubjectOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrganization

`func (o *Certificate) SetSubjectOrganization(v string)`

SetSubjectOrganization sets SubjectOrganization field to given value.


### SetSubjectOrganizationNil

`func (o *Certificate) SetSubjectOrganizationNil(b bool)`

 SetSubjectOrganizationNil sets the value for SubjectOrganization to be an explicit nil

### UnsetSubjectOrganization
`func (o *Certificate) UnsetSubjectOrganization()`

UnsetSubjectOrganization ensures that no value is present for SubjectOrganization, not even an explicit nil
### GetAcme

`func (o *Certificate) GetAcme() bool`

GetAcme returns the Acme field if non-nil, zero value otherwise.

### GetAcmeOk

`func (o *Certificate) GetAcmeOk() (*bool, bool)`

GetAcmeOk returns a tuple with the Acme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcme

`func (o *Certificate) SetAcme(v bool)`

SetAcme sets Acme field to given value.


### GetLeafCertificate

`func (o *Certificate) GetLeafCertificate() string`

GetLeafCertificate returns the LeafCertificate field if non-nil, zero value otherwise.

### GetLeafCertificateOk

`func (o *Certificate) GetLeafCertificateOk() (*string, bool)`

GetLeafCertificateOk returns a tuple with the LeafCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCertificate

`func (o *Certificate) SetLeafCertificate(v string)`

SetLeafCertificate sets LeafCertificate field to given value.


### GetCertificateChain

`func (o *Certificate) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *Certificate) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *Certificate) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.


### GetSha256Fingerprint

`func (o *Certificate) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *Certificate) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *Certificate) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.


### GetTrusted

`func (o *Certificate) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *Certificate) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *Certificate) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.


### GetSelfSigned

`func (o *Certificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *Certificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *Certificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.


### GetCertificateArn

`func (o *Certificate) GetCertificateArn() string`

GetCertificateArn returns the CertificateArn field if non-nil, zero value otherwise.

### GetCertificateArnOk

`func (o *Certificate) GetCertificateArnOk() (*string, bool)`

GetCertificateArnOk returns a tuple with the CertificateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateArn

`func (o *Certificate) SetCertificateArn(v string)`

SetCertificateArn sets CertificateArn field to given value.


### SetCertificateArnNil

`func (o *Certificate) SetCertificateArnNil(b bool)`

 SetCertificateArnNil sets the value for CertificateArn to be an explicit nil

### UnsetCertificateArn
`func (o *Certificate) UnsetCertificateArn()`

UnsetCertificateArn ensures that no value is present for CertificateArn, not even an explicit nil
### GetSubjectAlternativeNames

`func (o *Certificate) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *Certificate) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *Certificate) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.


### GetPrivateKeyAlgorithm

`func (o *Certificate) GetPrivateKeyAlgorithm() string`

GetPrivateKeyAlgorithm returns the PrivateKeyAlgorithm field if non-nil, zero value otherwise.

### GetPrivateKeyAlgorithmOk

`func (o *Certificate) GetPrivateKeyAlgorithmOk() (*string, bool)`

GetPrivateKeyAlgorithmOk returns a tuple with the PrivateKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAlgorithm

`func (o *Certificate) SetPrivateKeyAlgorithm(v string)`

SetPrivateKeyAlgorithm sets PrivateKeyAlgorithm field to given value.


### GetCreatedAt

`func (o *Certificate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Certificate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Certificate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Certificate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Certificate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Certificate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *Certificate) GetLinks() CertificateLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Certificate) GetLinksOk() (*CertificateLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Certificate) SetLinks(v CertificateLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Certificate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


