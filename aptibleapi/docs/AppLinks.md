# AppLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AppLinksAccount**](AppLinksAccount.md) |  | [optional] 
**CurrentConfiguration** | Pointer to [**AppLinksCurrentConfiguration**](AppLinksCurrentConfiguration.md) |  | [optional] 
**CurrentImage** | Pointer to [**AppLinksCurrentImage**](AppLinksCurrentImage.md) |  | [optional] 
**Operations** | Pointer to [**AppLinksOperations**](AppLinksOperations.md) |  | [optional] 
**Images** | Pointer to [**AppLinksImages**](AppLinksImages.md) |  | [optional] 
**Configurations** | Pointer to [**AppLinksConfigurations**](AppLinksConfigurations.md) |  | [optional] 
**Services** | Pointer to [**AppLinksServices**](AppLinksServices.md) |  | [optional] 
**Vhosts** | Pointer to [**AppLinksVhosts**](AppLinksVhosts.md) |  | [optional] 
**EphemeralSessions** | Pointer to [**AppLinksEphemeralSessions**](AppLinksEphemeralSessions.md) |  | [optional] 
**ServiceDefinitions** | Pointer to [**AppLinksServiceDefinitions**](AppLinksServiceDefinitions.md) |  | [optional] 
**PrereleaseCommands** | Pointer to [**AppLinksPrereleaseCommands**](AppLinksPrereleaseCommands.md) |  | [optional] 
**CodeScanResults** | Pointer to [**AppLinksCodeScanResults**](AppLinksCodeScanResults.md) |  | [optional] 
**LastCodeScanResult** | Pointer to [**AppLinksLastCodeScanResult**](AppLinksLastCodeScanResult.md) |  | [optional] 
**Deployments** | Pointer to [**AppLinksDeployments**](AppLinksDeployments.md) |  | [optional] 
**CurrentDeployment** | Pointer to [**AppLinksCurrentDeployment**](AppLinksCurrentDeployment.md) |  | [optional] 
**Self** | Pointer to [**AppLinksSelf**](AppLinksSelf.md) |  | [optional] 

## Methods

### NewAppLinks

`func NewAppLinks() *AppLinks`

NewAppLinks instantiates a new AppLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinksWithDefaults

`func NewAppLinksWithDefaults() *AppLinks`

NewAppLinksWithDefaults instantiates a new AppLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AppLinks) GetAccount() AppLinksAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AppLinks) GetAccountOk() (*AppLinksAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AppLinks) SetAccount(v AppLinksAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AppLinks) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCurrentConfiguration

`func (o *AppLinks) GetCurrentConfiguration() AppLinksCurrentConfiguration`

GetCurrentConfiguration returns the CurrentConfiguration field if non-nil, zero value otherwise.

### GetCurrentConfigurationOk

`func (o *AppLinks) GetCurrentConfigurationOk() (*AppLinksCurrentConfiguration, bool)`

GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfiguration

`func (o *AppLinks) SetCurrentConfiguration(v AppLinksCurrentConfiguration)`

SetCurrentConfiguration sets CurrentConfiguration field to given value.

### HasCurrentConfiguration

`func (o *AppLinks) HasCurrentConfiguration() bool`

HasCurrentConfiguration returns a boolean if a field has been set.

### GetCurrentImage

`func (o *AppLinks) GetCurrentImage() AppLinksCurrentImage`

GetCurrentImage returns the CurrentImage field if non-nil, zero value otherwise.

### GetCurrentImageOk

`func (o *AppLinks) GetCurrentImageOk() (*AppLinksCurrentImage, bool)`

GetCurrentImageOk returns a tuple with the CurrentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentImage

`func (o *AppLinks) SetCurrentImage(v AppLinksCurrentImage)`

SetCurrentImage sets CurrentImage field to given value.

### HasCurrentImage

`func (o *AppLinks) HasCurrentImage() bool`

HasCurrentImage returns a boolean if a field has been set.

### GetOperations

`func (o *AppLinks) GetOperations() AppLinksOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AppLinks) GetOperationsOk() (*AppLinksOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AppLinks) SetOperations(v AppLinksOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *AppLinks) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetImages

`func (o *AppLinks) GetImages() AppLinksImages`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AppLinks) GetImagesOk() (*AppLinksImages, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AppLinks) SetImages(v AppLinksImages)`

SetImages sets Images field to given value.

### HasImages

`func (o *AppLinks) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetConfigurations

`func (o *AppLinks) GetConfigurations() AppLinksConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *AppLinks) GetConfigurationsOk() (*AppLinksConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *AppLinks) SetConfigurations(v AppLinksConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *AppLinks) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetServices

`func (o *AppLinks) GetServices() AppLinksServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AppLinks) GetServicesOk() (*AppLinksServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AppLinks) SetServices(v AppLinksServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *AppLinks) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetVhosts

`func (o *AppLinks) GetVhosts() AppLinksVhosts`

GetVhosts returns the Vhosts field if non-nil, zero value otherwise.

### GetVhostsOk

`func (o *AppLinks) GetVhostsOk() (*AppLinksVhosts, bool)`

GetVhostsOk returns a tuple with the Vhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhosts

`func (o *AppLinks) SetVhosts(v AppLinksVhosts)`

SetVhosts sets Vhosts field to given value.

### HasVhosts

`func (o *AppLinks) HasVhosts() bool`

HasVhosts returns a boolean if a field has been set.

### GetEphemeralSessions

`func (o *AppLinks) GetEphemeralSessions() AppLinksEphemeralSessions`

GetEphemeralSessions returns the EphemeralSessions field if non-nil, zero value otherwise.

### GetEphemeralSessionsOk

`func (o *AppLinks) GetEphemeralSessionsOk() (*AppLinksEphemeralSessions, bool)`

GetEphemeralSessionsOk returns a tuple with the EphemeralSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralSessions

`func (o *AppLinks) SetEphemeralSessions(v AppLinksEphemeralSessions)`

SetEphemeralSessions sets EphemeralSessions field to given value.

### HasEphemeralSessions

`func (o *AppLinks) HasEphemeralSessions() bool`

HasEphemeralSessions returns a boolean if a field has been set.

### GetServiceDefinitions

`func (o *AppLinks) GetServiceDefinitions() AppLinksServiceDefinitions`

GetServiceDefinitions returns the ServiceDefinitions field if non-nil, zero value otherwise.

### GetServiceDefinitionsOk

`func (o *AppLinks) GetServiceDefinitionsOk() (*AppLinksServiceDefinitions, bool)`

GetServiceDefinitionsOk returns a tuple with the ServiceDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefinitions

`func (o *AppLinks) SetServiceDefinitions(v AppLinksServiceDefinitions)`

SetServiceDefinitions sets ServiceDefinitions field to given value.

### HasServiceDefinitions

`func (o *AppLinks) HasServiceDefinitions() bool`

HasServiceDefinitions returns a boolean if a field has been set.

### GetPrereleaseCommands

`func (o *AppLinks) GetPrereleaseCommands() AppLinksPrereleaseCommands`

GetPrereleaseCommands returns the PrereleaseCommands field if non-nil, zero value otherwise.

### GetPrereleaseCommandsOk

`func (o *AppLinks) GetPrereleaseCommandsOk() (*AppLinksPrereleaseCommands, bool)`

GetPrereleaseCommandsOk returns a tuple with the PrereleaseCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrereleaseCommands

`func (o *AppLinks) SetPrereleaseCommands(v AppLinksPrereleaseCommands)`

SetPrereleaseCommands sets PrereleaseCommands field to given value.

### HasPrereleaseCommands

`func (o *AppLinks) HasPrereleaseCommands() bool`

HasPrereleaseCommands returns a boolean if a field has been set.

### GetCodeScanResults

`func (o *AppLinks) GetCodeScanResults() AppLinksCodeScanResults`

GetCodeScanResults returns the CodeScanResults field if non-nil, zero value otherwise.

### GetCodeScanResultsOk

`func (o *AppLinks) GetCodeScanResultsOk() (*AppLinksCodeScanResults, bool)`

GetCodeScanResultsOk returns a tuple with the CodeScanResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeScanResults

`func (o *AppLinks) SetCodeScanResults(v AppLinksCodeScanResults)`

SetCodeScanResults sets CodeScanResults field to given value.

### HasCodeScanResults

`func (o *AppLinks) HasCodeScanResults() bool`

HasCodeScanResults returns a boolean if a field has been set.

### GetLastCodeScanResult

`func (o *AppLinks) GetLastCodeScanResult() AppLinksLastCodeScanResult`

GetLastCodeScanResult returns the LastCodeScanResult field if non-nil, zero value otherwise.

### GetLastCodeScanResultOk

`func (o *AppLinks) GetLastCodeScanResultOk() (*AppLinksLastCodeScanResult, bool)`

GetLastCodeScanResultOk returns a tuple with the LastCodeScanResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCodeScanResult

`func (o *AppLinks) SetLastCodeScanResult(v AppLinksLastCodeScanResult)`

SetLastCodeScanResult sets LastCodeScanResult field to given value.

### HasLastCodeScanResult

`func (o *AppLinks) HasLastCodeScanResult() bool`

HasLastCodeScanResult returns a boolean if a field has been set.

### GetDeployments

`func (o *AppLinks) GetDeployments() AppLinksDeployments`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *AppLinks) GetDeploymentsOk() (*AppLinksDeployments, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *AppLinks) SetDeployments(v AppLinksDeployments)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *AppLinks) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetCurrentDeployment

`func (o *AppLinks) GetCurrentDeployment() AppLinksCurrentDeployment`

GetCurrentDeployment returns the CurrentDeployment field if non-nil, zero value otherwise.

### GetCurrentDeploymentOk

`func (o *AppLinks) GetCurrentDeploymentOk() (*AppLinksCurrentDeployment, bool)`

GetCurrentDeploymentOk returns a tuple with the CurrentDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeployment

`func (o *AppLinks) SetCurrentDeployment(v AppLinksCurrentDeployment)`

SetCurrentDeployment sets CurrentDeployment field to given value.

### HasCurrentDeployment

`func (o *AppLinks) HasCurrentDeployment() bool`

HasCurrentDeployment returns a boolean if a field has been set.

### GetSelf

`func (o *AppLinks) GetSelf() AppLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AppLinks) GetSelfOk() (*AppLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AppLinks) SetSelf(v AppLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AppLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


