package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SubscriptionAccountOwner represents the SubscriptionAccountOwner schema from the OpenAPI specification
type SubscriptionAccountOwner struct {
	Email string `json:"email,omitempty"`
	Puid string `json:"puid,omitempty"`
}

// ComputeUsageDetail represents the ComputeUsageDetail schema from the OpenAPI specification
type ComputeUsageDetail struct {
	Usage float64 `json:"usage,omitempty"`
	Sku string `json:"sku,omitempty"`
}

// CreateEnvironmentStateChangeBody represents the CreateEnvironmentStateChangeBody schema from the OpenAPI specification
type CreateEnvironmentStateChangeBody struct {
	Newvalue int `json:"newValue,omitempty"`
	Oldvalue int `json:"oldValue,omitempty"`
	Time string `json:"time,omitempty"`
}

// UpdateSystemConfigurationBody represents the UpdateSystemConfigurationBody schema from the OpenAPI specification
type UpdateSystemConfigurationBody struct {
	Comment string `json:"comment,omitempty"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// RuntimeConstraintsBody represents the RuntimeConstraintsBody schema from the OpenAPI specification
type RuntimeConstraintsBody struct {
	Allowedportprivacysettings []int `json:"allowedPortPrivacySettings,omitempty"`
	Imageallowlist []string `json:"imageAllowList,omitempty"`
}

// CreatePoolGroupBody represents the CreatePoolGroupBody schema from the OpenAPI specification
type CreatePoolGroupBody struct {
	Region int `json:"region"`
	Tags map[string]interface{} `json:"tags,omitempty"`
	Displayname string `json:"displayName"`
}

// LocationsResult represents the LocationsResult schema from the OpenAPI specification
type LocationsResult struct {
	Current int `json:"current,omitempty"`
	Hostnames map[string]interface{} `json:"hostnames,omitempty"`
	Available []int `json:"available,omitempty"`
}

// PrebuildReadinessResult represents the PrebuildReadinessResult schema from the OpenAPI specification
type PrebuildReadinessResult struct {
	Location int `json:"location,omitempty"`
	Poolskus []string `json:"poolSkus,omitempty"`
	Prebuildhash string `json:"prebuildHash,omitempty"`
	Repoid string `json:"repoId,omitempty"`
	Supportedskus []string `json:"supportedSkus,omitempty"`
	Templateskus []string `json:"templateSkus,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
}

// GitConfigOptionsBody represents the GitConfigOptionsBody schema from the OpenAPI specification
type GitConfigOptionsBody struct {
	Useremail string `json:"userEmail,omitempty"`
	Username string `json:"userName,omitempty"`
}

// TunnelPortInfoResult represents the TunnelPortInfoResult schema from the OpenAPI specification
type TunnelPortInfoResult struct {
	Portvisibility string `json:"portVisibility,omitempty"`
	Tunneltoken string `json:"tunnelToken,omitempty"`
}

// BillSummary represents the BillSummary schema from the OpenAPI specification
type BillSummary struct {
	Billgenerationtime string `json:"billGenerationTime,omitempty"`
	Location int `json:"location,omitempty"`
	Periodstart string `json:"periodStart,omitempty"`
	Environmentid string `json:"environmentId,omitempty"`
	Id string `json:"id,omitempty"`
	Plan VsoPlanInfo `json:"plan,omitempty"`
	Usage map[string]interface{} `json:"usage,omitempty"`
	Partitionkey string `json:"partitionKey,omitempty"`
	Periodend string `json:"periodEnd,omitempty"`
	Usagedetail []EnvironmentUsage `json:"usageDetail,omitempty"`
}

// VMResult represents the VMResult schema from the OpenAPI specification
type VMResult struct {
	Connection VMConnectionInfo `json:"connection,omitempty"`
	Provisioningstatus ProvisioningStatusResult `json:"provisioningStatus,omitempty"`
	Status int `json:"status,omitempty"`
}

// CreatePrebuildTemplateBody represents the CreatePrebuildTemplateBody schema from the OpenAPI specification
type CreatePrebuildTemplateBody struct {
	Githubprebuildtemplateendpoint string `json:"gitHubPrebuildTemplateEndpoint,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	Storagetype int `json:"storageType,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Friendlyname string `json:"friendlyName"`
	Planid string `json:"planId,omitempty"`
	Templateinfo PrebuildTemplateInfo `json:"templateInfo,omitempty"`
	Experimentalfeatures ExperimentalFeaturesBody `json:"experimentalFeatures,omitempty"`
	Githubprebuildinstanceendpoint string `json:"gitHubPrebuildInstanceEndpoint,omitempty"`
}

// PoolStatusResponseBody represents the PoolStatusResponseBody schema from the OpenAPI specification
type PoolStatusResponseBody struct {
	Readyunassignednotlatestversioncount int `json:"readyUnassignedNotLatestVersionCount,omitempty"`
	Readyunassignedcount int `json:"readyUnassignedCount,omitempty"`
	Sku string `json:"sku,omitempty"`
	Isenvironmentpool bool `json:"isEnvironmentPool,omitempty"`
	Resourcetype string `json:"resourceType,omitempty"`
	Location string `json:"location,omitempty"`
	Readyunassignednotlatestversionandidlecount int `json:"readyUnassignedNotLatestVersionAndIdleCount,omitempty"`
	Allwithlatestversion bool `json:"allWithLatestVersion,omitempty"`
	Poolcode string `json:"poolCode,omitempty"`
	Readyunassignedlatestversioncount int `json:"readyUnassignedLatestVersionCount,omitempty"`
}

// VMConnectionInfo represents the VMConnectionInfo schema from the OpenAPI specification
type VMConnectionInfo struct {
	Connectiontype int `json:"connectionType,omitempty"`
	Liveshareworkspaceid string `json:"liveShareWorkspaceId,omitempty"`
}

// RefreshProfileTelemetryPropertiesResponse represents the RefreshProfileTelemetryPropertiesResponse schema from the OpenAPI specification
type RefreshProfileTelemetryPropertiesResponse struct {
	Failed []ProfileSpecifier `json:"failed,omitempty"`
	Succeeded []ProfileSpecifier `json:"succeeded,omitempty"`
}

// CreateEnvironmentPoolResourceBody represents the CreateEnvironmentPoolResourceBody schema from the OpenAPI specification
type CreateEnvironmentPoolResourceBody struct {
	Environmentoptions PrebuildEnvironmentOptions `json:"environmentOptions,omitempty"`
	Secrets []SecretDataBody `json:"secrets,omitempty"`
}

// IdentityBody represents the IdentityBody schema from the OpenAPI specification
type IdentityBody struct {
	Displayname string `json:"displayName,omitempty"`
	Id string `json:"id,omitempty"`
	Username string `json:"userName,omitempty"`
}

// ClientUsageData represents the ClientUsageData schema from the OpenAPI specification
type ClientUsageData struct {
	Activeminutes int `json:"activeMinutes,omitempty"`
	Lastactivity string `json:"lastActivity,omitempty"`
}

// ChangeResourceDeletionRequestBody represents the ChangeResourceDeletionRequestBody schema from the OpenAPI specification
type ChangeResourceDeletionRequestBody struct {
	Poolcode string `json:"poolCode,omitempty"`
	Pooltype string `json:"poolType,omitempty"`
	Region string `json:"region,omitempty"`
	Comment string `json:"comment,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// CollectedData represents the CollectedData schema from the OpenAPI specification
type CollectedData struct {
	Timestamp string `json:"timestamp,omitempty"`
	Environmentid string `json:"environmentId,omitempty"`
	Name string `json:"name,omitempty"`
	Parentactivityid string `json:"parentActivityId,omitempty"`
}

// PrebuildTemplateInfoResult represents the PrebuildTemplateInfoResult schema from the OpenAPI specification
type PrebuildTemplateInfoResult struct {
	Lastusedtime string `json:"lastUsedTime,omitempty"`
	Id string `json:"id,omitempty"`
	Logicalskus []string `json:"logicalSkus,omitempty"`
	Prebuildhash string `json:"prebuildHash,omitempty"`
	Isprebuild bool `json:"isPrebuild,omitempty"`
	Repoid int64 `json:"repoId,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Devcontainerpath string `json:"devcontainerPath,omitempty"`
	Prebuildconfigurationid string `json:"prebuildConfigurationId,omitempty"`
	Templatestatus string `json:"templateStatus,omitempty"`
	Commitid string `json:"commitId,omitempty"`
}

// SkuInfoResult represents the SkuInfoResult schema from the OpenAPI specification
type SkuInfoResult struct {
	Os string `json:"os,omitempty"`
	Availablesettings AvailableSettingsResult `json:"availableSettings,omitempty"`
	Displayname string `json:"displayName,omitempty"`
	Name string `json:"name,omitempty"`
}

// CloudEnvironmentFolderBody represents the CloudEnvironmentFolderBody schema from the OpenAPI specification
type CloudEnvironmentFolderBody struct {
	Recentfolderpaths []string `json:"recentFolderPaths,omitempty"`
}

// SubscriptionAdditionalProperties represents the SubscriptionAdditionalProperties schema from the OpenAPI specification
type SubscriptionAdditionalProperties struct {
	Billingproperties BillingProperties `json:"billingProperties,omitempty"`
	Resourceproviderproperties string `json:"resourceProviderProperties,omitempty"`
}

// CreateOrUpdatePoolBody represents the CreateOrUpdatePoolBody schema from the OpenAPI specification
type CreateOrUpdatePoolBody struct {
	Usergroupname string `json:"userGroupName,omitempty"`
	Vmspecs VMSpecs `json:"vmSpecs"`
	Domainusercredentials DomainUserCredentials `json:"domainUserCredentials,omitempty"`
	Hotpoolsettings HotPoolSettings `json:"hotPoolSettings,omitempty"`
	Poolgroupname string `json:"poolGroupName"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// EnvironmentUsage represents the EnvironmentUsage schema from the OpenAPI specification
type EnvironmentUsage struct {
	Endstate int `json:"endState,omitempty"`
	Id string `json:"id,omitempty"`
	Resourceusage ResourceUsageDetail `json:"resourceUsage,omitempty"`
	Sku Sku `json:"sku,omitempty"`
}

// RefreshProfileTelemetryPropertiesRequest represents the RefreshProfileTelemetryPropertiesRequest schema from the OpenAPI specification
type RefreshProfileTelemetryPropertiesRequest struct {
	Partner string `json:"partner,omitempty"`
	Tenantid string `json:"tenantId,omitempty"`
	Userids string `json:"userIds,omitempty"`
}

// ResourceUsageDetail represents the ResourceUsageDetail schema from the OpenAPI specification
type ResourceUsageDetail struct {
	Compute []ComputeUsageDetail `json:"compute,omitempty"`
	Storage []StorageUsageDetail `json:"storage,omitempty"`
}

// CloudEnvironmentResult represents the CloudEnvironmentResult schema from the OpenAPI specification
type CloudEnvironmentResult struct {
	TypeField string `json:"type,omitempty"`
	Accesstoken string `json:"accessToken,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Connection ConnectionInfoBody `json:"connection,omitempty"`
	Skudisplayname string `json:"skuDisplayName,omitempty"`
	Portforwardingconnection ConnectionInfoBody `json:"portForwardingConnection,omitempty"`
	Recentfolders []string `json:"recentFolders,omitempty"`
	Id string `json:"id,omitempty"`
	Subscriptiondata SubscriptionData `json:"subscriptionData,omitempty"`
	Laststateupdatereason string `json:"lastStateUpdateReason,omitempty"`
	Billableownertype int `json:"billableOwnerType,omitempty"`
	Ownerid string `json:"ownerId,omitempty"`
	Storageutilizationinkb int64 `json:"storageUtilizationInKb,omitempty"`
	Gitstatus GitStatus `json:"gitStatus,omitempty"`
	Location string `json:"location,omitempty"`
	Planid string `json:"planId,omitempty"`
	Organizationid string `json:"organizationId,omitempty"`
	Container ContainerInfoBody `json:"container,omitempty"`
	Failoverdetails FailoverDetails `json:"failoverDetails,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Exportedbloburl string `json:"exportedBlobUrl,omitempty"`
	Lastused string `json:"lastUsed,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	Clientusage ClientUsageSession `json:"clientUsage,omitempty"`
	Templatestatus string `json:"templateStatus,omitempty"`
	Active string `json:"active,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Containerimage string `json:"containerImage,omitempty"`
	Friendlyname string `json:"friendlyName,omitempty"`
	Platform string `json:"platform,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Displaystorageutilizationinkb bool `json:"displayStorageUtilizationInKb,omitempty"`
	Createfromprebuild bool `json:"createFromPrebuild,omitempty"`
	Resourcetier int `json:"resourceTier,omitempty"`
	State string `json:"state,omitempty"`
	Runtimeconstraints RuntimeConstraintsBody `json:"runtimeConstraints,omitempty"`
	Prebuildtype string `json:"prebuildType,omitempty"`
}

// UnfilteredCloudEnvironmentResult represents the UnfilteredCloudEnvironmentResult schema from the OpenAPI specification
type UnfilteredCloudEnvironmentResult struct {
	Ownerid string `json:"ownerId,omitempty"`
	Platform string `json:"platform,omitempty"`
	Prebuildtype string `json:"prebuildType,omitempty"`
	Subscriptiondata SubscriptionData `json:"subscriptionData,omitempty"`
	TypeField string `json:"type,omitempty"`
	Exportedbloburl string `json:"exportedBlobUrl,omitempty"`
	Gitstatus GitStatus `json:"gitStatus,omitempty"`
	Laststateupdatereason string `json:"lastStateUpdateReason,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Displaystorageutilizationinkb bool `json:"displayStorageUtilizationInKb,omitempty"`
	Lastused string `json:"lastUsed,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	Containerimage string `json:"containerImage,omitempty"`
	Runtimeconstraints RuntimeConstraintsBody `json:"runtimeConstraints,omitempty"`
	Container ContainerInfoBody `json:"container,omitempty"`
	Skudisplayname string `json:"skuDisplayName,omitempty"`
	Portforwardingconnection ConnectionInfoBody `json:"portForwardingConnection,omitempty"`
	Planid string `json:"planId,omitempty"`
	Templatestatus string `json:"templateStatus,omitempty"`
	Resourcetier int `json:"resourceTier,omitempty"`
	Id string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
	Recentfolders []string `json:"recentFolders,omitempty"`
	Friendlyname string `json:"friendlyName,omitempty"`
	Storageutilizationinkb int64 `json:"storageUtilizationInKb,omitempty"`
	Created string `json:"created,omitempty"`
	Accesstoken string `json:"accessToken,omitempty"`
	Createfromprebuild bool `json:"createFromPrebuild,omitempty"`
	Updated string `json:"updated,omitempty"`
	Location string `json:"location,omitempty"`
	Organizationid string `json:"organizationId,omitempty"`
	Billableownertype int `json:"billableOwnerType,omitempty"`
	Failoverdetails FailoverDetails `json:"failoverDetails,omitempty"`
	Connection ConnectionInfoBody `json:"connection,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Clientusage ClientUsageSession `json:"clientUsage,omitempty"`
	Active string `json:"active,omitempty"`
}

// PlanResourceProperties represents the PlanResourceProperties schema from the OpenAPI specification
type PlanResourceProperties struct {
	Defaultenvironmentsku string `json:"defaultEnvironmentSku,omitempty"`
	Encryption PlanResourceEncryptionProperties `json:"encryption,omitempty"`
	Userid string `json:"userId,omitempty"`
	Vnetproperties VnetProperties `json:"vnetProperties,omitempty"`
	Defaultcodespacesku string `json:"defaultCodespaceSku,omitempty"`
}

// IssueDelegatePlanAccessTokenBody represents the IssueDelegatePlanAccessTokenBody schema from the OpenAPI specification
type IssueDelegatePlanAccessTokenBody struct {
	Scope string `json:"scope,omitempty"`
	Environmentids []string `json:"environmentIds,omitempty"`
	Expiration string `json:"expiration,omitempty"`
	Identity DelegateIdentity `json:"identity,omitempty"`
	Portnumbers []int `json:"portNumbers,omitempty"`
}

// ClientUsageSession represents the ClientUsageSession schema from the OpenAPI specification
type ClientUsageSession struct {
	Sessionid string `json:"sessionId,omitempty"`
	Usagedata map[string]interface{} `json:"usageData,omitempty"`
}

// VMSpecs represents the VMSpecs schema from the OpenAPI specification
type VMSpecs struct {
	Subnetresourceid string `json:"subnetResourceId"`
	Disktype int `json:"diskType"`
	Imageresourceid string `json:"imageResourceId"`
	Size string `json:"size"`
}

// ProvisioningStatusResult represents the ProvisioningStatusResult schema from the OpenAPI specification
type ProvisioningStatusResult struct {
	Totalsteps int `json:"totalSteps,omitempty"`
	Completedsteps int `json:"completedSteps,omitempty"`
	Currentstepdescription string `json:"currentStepDescription,omitempty"`
	Isready bool `json:"isReady,omitempty"`
	Operationstartedtimeutc string `json:"operationStartedTimeUtc,omitempty"`
}

// ClaimVMBody represents the ClaimVMBody schema from the OpenAPI specification
type ClaimVMBody struct {
	User UserIdentity `json:"user"`
}

// UnderInvestigationResponseBody represents the UnderInvestigationResponseBody schema from the OpenAPI specification
type UnderInvestigationResponseBody struct {
	Underinvestigation bool `json:"underInvestigation,omitempty"`
	Updated bool `json:"updated,omitempty"`
	Investigationstarted string `json:"investigationStarted,omitempty"`
	Resourceid string `json:"resourceId,omitempty"`
}

// UserIdentity represents the UserIdentity schema from the OpenAPI specification
type UserIdentity struct {
	Userprincipalname string `json:"userPrincipalName"`
}

// Sku represents the Sku schema from the OpenAPI specification
type Sku struct {
	Name string `json:"name,omitempty"`
	Tier string `json:"tier,omitempty"`
}

// PlanResourceIdentity represents the PlanResourceIdentity schema from the OpenAPI specification
type PlanResourceIdentity struct {
	Principalid string `json:"principalId,omitempty"`
	Tenantid string `json:"tenantId,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// EnvironmentBillingInfo represents the EnvironmentBillingInfo schema from the OpenAPI specification
type EnvironmentBillingInfo struct {
	Sku Sku `json:"sku,omitempty"`
	Userid string `json:"userId,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Pool represents the Pool schema from the OpenAPI specification
type Pool struct {
	Pooltype int `json:"poolType,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Targetcount int64 `json:"targetCount,omitempty"`
}

// EnvironmentRegistrationCallbackPayloadBody represents the EnvironmentRegistrationCallbackPayloadBody schema from the OpenAPI specification
type EnvironmentRegistrationCallbackPayloadBody struct {
	Sessionid string `json:"sessionId,omitempty"`
	Sessionpath string `json:"sessionPath,omitempty"`
}

// FailoverDetails represents the FailoverDetails schema from the OpenAPI specification
type FailoverDetails struct {
	Failoverenabled bool `json:"failoverEnabled,omitempty"`
	Failoverregion int `json:"failoverRegion,omitempty"`
}

// SubscriptionData represents the SubscriptionData schema from the OpenAPI specification
type SubscriptionData struct {
	Computequota int `json:"computeQuota,omitempty"`
	Computeusage int `json:"computeUsage,omitempty"`
	Subscriptionid string `json:"subscriptionId,omitempty"`
	Subscriptionstate string `json:"subscriptionState,omitempty"`
}

// PlanResourceEncryptionProperties represents the PlanResourceEncryptionProperties schema from the OpenAPI specification
type PlanResourceEncryptionProperties struct {
	Keysource string `json:"keySource,omitempty"`
	Keyvaultproperties PlanResourceKeyVaultProperties `json:"keyVaultProperties,omitempty"`
}

// UpdateCloudEnvironmentBody represents the UpdateCloudEnvironmentBody schema from the OpenAPI specification
type UpdateCloudEnvironmentBody struct {
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Failoverdetails FailoverDetails `json:"failoverDetails,omitempty"`
	Friendlyname string `json:"friendlyName,omitempty"`
	Planaccesstoken string `json:"planAccessToken,omitempty"`
	Planid string `json:"planId,omitempty"`
	Skuname string `json:"skuName,omitempty"`
}

// EnvironmentRegistrationCallbackBody represents the EnvironmentRegistrationCallbackBody schema from the OpenAPI specification
type EnvironmentRegistrationCallbackBody struct {
	Payload EnvironmentRegistrationCallbackPayloadBody `json:"payload,omitempty"`
	TypeField string `json:"type"`
}

// StorageUsageDetail represents the StorageUsageDetail schema from the OpenAPI specification
type StorageUsageDetail struct {
	Sku string `json:"sku,omitempty"`
	Usage float64 `json:"usage,omitempty"`
	Size int `json:"size,omitempty"`
	Sizeinkb int64 `json:"sizeInKB,omitempty"`
}

// UpdatePoolGroupBody represents the UpdatePoolGroupBody schema from the OpenAPI specification
type UpdatePoolGroupBody struct {
	Displayname string `json:"displayName"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// UpdatePrebuildTemplateVersionsBody represents the UpdatePrebuildTemplateVersionsBody schema from the OpenAPI specification
type UpdatePrebuildTemplateVersionsBody struct {
	Repoid int64 `json:"repoId"`
	Branchname string `json:"branchName"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Maxprebuildtemplateversions int `json:"maxPrebuildTemplateVersions"`
}

// PlanResourceKeyVaultProperties represents the PlanResourceKeyVaultProperties schema from the OpenAPI specification
type PlanResourceKeyVaultProperties struct {
	Keyname string `json:"keyName,omitempty"`
	Keyvaulturi string `json:"keyVaultUri,omitempty"`
	Keyversion string `json:"keyVersion,omitempty"`
}

// AddForwardedPortSettings represents the AddForwardedPortSettings schema from the OpenAPI specification
type AddForwardedPortSettings struct {
	Tunneltype int `json:"tunnelType,omitempty"`
	Privacy int `json:"privacy,omitempty"`
}

// RPSubscriptionProperties represents the RPSubscriptionProperties schema from the OpenAPI specification
type RPSubscriptionProperties struct {
	Additionalproperties SubscriptionAdditionalProperties `json:"additionalProperties,omitempty"`
	Locationplacementid string `json:"locationPlacementId,omitempty"`
	Managedbytenants []StringStringKeyValuePair `json:"managedByTenants,omitempty"`
	Quotaid string `json:"quotaId,omitempty"`
	Registeredfeatures []StringStringKeyValuePair `json:"registeredFeatures,omitempty"`
	Tenantid string `json:"tenantId,omitempty"`
	Accountowner SubscriptionAccountOwner `json:"accountOwner,omitempty"`
}

// UpdatePrebuildTemplateBody represents the UpdatePrebuildTemplateBody schema from the OpenAPI specification
type UpdatePrebuildTemplateBody struct {
	Issuccess bool `json:"isSuccess"`
}

// VmLogsUploadInfo represents the VmLogsUploadInfo schema from the OpenAPI specification
type VmLogsUploadInfo struct {
	Containername string `json:"containerName,omitempty"`
	Message string `json:"message,omitempty"`
	Pathincontainer string `json:"pathInContainer,omitempty"`
	Storageuri string `json:"storageUri,omitempty"`
	Vmresourceid string `json:"vmResourceId,omitempty"`
}

// PrebuildTemplateInfo represents the PrebuildTemplateInfo schema from the OpenAPI specification
type PrebuildTemplateInfo struct {
	Prebuildconfigurationid string `json:"prebuildConfigurationId,omitempty"`
	Templatesizeingb float64 `json:"templateSizeInGB,omitempty"`
	Totaltimesavingsinseconds string `json:"totalTimeSavingsInSeconds,omitempty"`
	Workflowrunid string `json:"workFlowRunId,omitempty"`
	Container ContainerInfo `json:"container,omitempty"`
}

// BillingProperties represents the BillingProperties schema from the OpenAPI specification
type BillingProperties struct {
	Workloadtype string `json:"workloadType,omitempty"`
	Billingtype string `json:"billingType,omitempty"`
	Channeltype string `json:"channelType,omitempty"`
	Paymenttype string `json:"paymentType,omitempty"`
	Tier string `json:"tier,omitempty"`
}

// StringStringKeyValuePair represents the StringStringKeyValuePair schema from the OpenAPI specification
type StringStringKeyValuePair struct {
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// EnvironmentStateChange represents the EnvironmentStateChange schema from the OpenAPI specification
type EnvironmentStateChange struct {
	Environment EnvironmentBillingInfo `json:"environment,omitempty"`
	Id string `json:"id,omitempty"`
	Newvalue int `json:"newValue,omitempty"`
	Oldvalue int `json:"oldValue,omitempty"`
	Partitionkey string `json:"partitionKey,omitempty"`
	Time string `json:"time,omitempty"`
}

// UpdateUserSecretsRequestBody represents the UpdateUserSecretsRequestBody schema from the OpenAPI specification
type UpdateUserSecretsRequestBody struct {
	Secrets []SecretDataBody `json:"secrets,omitempty"`
}

// PlanResourceUpdateBody represents the PlanResourceUpdateBody schema from the OpenAPI specification
type PlanResourceUpdateBody struct {
	Identity PlanResourceIdentity `json:"identity,omitempty"`
	Properties PlanResourceProperties `json:"properties,omitempty"`
}

// ReplayBillRequestBody represents the ReplayBillRequestBody schema from the OpenAPI specification
type ReplayBillRequestBody struct {
	Endtime string `json:"endTime,omitempty"`
	Starttime string `json:"startTime,omitempty"`
}

// HeartBeatBody represents the HeartBeatBody schema from the OpenAPI specification
type HeartBeatBody struct {
	Agentversion string `json:"agentVersion,omitempty"`
	Collecteddatalist []CollectedData `json:"collectedDataList,omitempty"`
	Environmentid string `json:"environmentId,omitempty"`
	Resourceid string `json:"resourceId,omitempty"`
	Timestamp string `json:"timeStamp,omitempty"`
}

// ProfileSpecifier represents the ProfileSpecifier schema from the OpenAPI specification
type ProfileSpecifier struct {
	Tid string `json:"tid,omitempty"`
	Oid string `json:"oid,omitempty"`
	Provider string `json:"provider,omitempty"`
}

// DeletePrebuildTemplatesBody represents the DeletePrebuildTemplatesBody schema from the OpenAPI specification
type DeletePrebuildTemplatesBody struct {
	Branchname string `json:"branchName"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Prebuildconfigurationid int64 `json:"prebuildConfigurationId,omitempty"`
	Repoid int64 `json:"repoId"`
}

// RPSubscriptionNotification represents the RPSubscriptionNotification schema from the OpenAPI specification
type RPSubscriptionNotification struct {
	Properties RPSubscriptionProperties `json:"properties,omitempty"`
	Registrationdate string `json:"registrationDate,omitempty"`
	State string `json:"state,omitempty"`
}

// AgentResponse represents the AgentResponse schema from the OpenAPI specification
type AgentResponse struct {
	Asseturi string `json:"assetUri,omitempty"`
	Family string `json:"family,omitempty"`
	Name string `json:"name,omitempty"`
}

// ContainerInfoBody represents the ContainerInfoBody schema from the OpenAPI specification
type ContainerInfoBody struct {
	Id string `json:"id,omitempty"`
	Schemaversion string `json:"schemaVersion,omitempty"`
}

// NetworkSettingsResource represents the NetworkSettingsResource schema from the OpenAPI specification
type NetworkSettingsResource struct {
	TypeField string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	Properties NetworkSettingsResourceProperties `json:"properties,omitempty"`
	Provisioningstate string `json:"provisioningState,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// ScopedCreateSecretBody represents the ScopedCreateSecretBody schema from the OpenAPI specification
type ScopedCreateSecretBody struct {
	Notes string `json:"notes,omitempty"`
	Scope int `json:"scope,omitempty"`
	Secretname string `json:"secretName,omitempty"`
	TypeField int `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Filters []SecretFilterBody `json:"filters,omitempty"`
}

// GitStatus represents the GitStatus schema from the OpenAPI specification
type GitStatus struct {
	Hasunpushedchanges bool `json:"hasUnpushedChanges,omitempty"`
	Nogitrepo bool `json:"noGitRepo,omitempty"`
	Ahead int `json:"ahead,omitempty"`
	Behind int `json:"behind,omitempty"`
	Branch string `json:"branch,omitempty"`
	Commit string `json:"commit,omitempty"`
	Hasuncommittedchanges bool `json:"hasUncommittedChanges,omitempty"`
}

// PoolResult represents the PoolResult schema from the OpenAPI specification
type PoolResult struct {
	Poolgroupname string `json:"poolGroupName"`
	Provisioningstatus ProvisioningStatusResult `json:"provisioningStatus,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
	Usergroupname string `json:"userGroupName,omitempty"`
	Vmspecs VMSpecs `json:"vmSpecs"`
	Domainusercredentials DomainUserCredentials `json:"domainUserCredentials,omitempty"`
	Hotpoolsettings HotPoolSettings `json:"hotPoolSettings,omitempty"`
}

// SeedInfoBody represents the SeedInfoBody schema from the OpenAPI specification
type SeedInfoBody struct {
	Recurseclone bool `json:"recurseClone,omitempty"`
	Repository RepositoryInfoBody `json:"repository,omitempty"`
	Seedmoniker string `json:"seedMoniker,omitempty"`
	Seedtype string `json:"seedType,omitempty"`
	Cloneurl string `json:"cloneUrl,omitempty"`
	Gitconfig GitConfigOptionsBody `json:"gitConfig,omitempty"`
}

// ContainerInfo represents the ContainerInfo schema from the OpenAPI specification
type ContainerInfo struct {
	Id string `json:"id,omitempty"`
	Imagename string `json:"imageName,omitempty"`
	Schemaversion string `json:"schemaVersion,omitempty"`
}

// PoolSettingsBody represents the PoolSettingsBody schema from the OpenAPI specification
type PoolSettingsBody struct {
	Subscription string `json:"subscription,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Pools []Pool `json:"pools,omitempty"`
	Repoid string `json:"repoId,omitempty"`
	Storagetype int `json:"storageType,omitempty"`
}

// ScopedSecretResultBody represents the ScopedSecretResultBody schema from the OpenAPI specification
type ScopedSecretResultBody struct {
	Filters []SecretFilterBody `json:"filters,omitempty"`
	Id string `json:"id,omitempty"`
	Lastmodified string `json:"lastModified,omitempty"`
	Notes string `json:"notes,omitempty"`
	Scope int `json:"scope,omitempty"`
	Secretname string `json:"secretName,omitempty"`
	TypeField int `json:"type,omitempty"`
}

// PoolGroupResult represents the PoolGroupResult schema from the OpenAPI specification
type PoolGroupResult struct {
	Displayname string `json:"displayName,omitempty"`
	Region int `json:"region,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// PersonalizationInfoBody represents the PersonalizationInfoBody schema from the OpenAPI specification
type PersonalizationInfoBody struct {
	Dotfilesinstallcommand string `json:"dotfilesInstallCommand,omitempty"`
	Dotfilesrepository string `json:"dotfilesRepository,omitempty"`
	Dotfilestargetpath string `json:"dotfilesTargetPath,omitempty"`
}

// TelemetryData represents the TelemetryData schema from the OpenAPI specification
type TelemetryData struct {
	Level string `json:"level,omitempty"`
	Message string `json:"message,omitempty"`
	Optionalvalues map[string]interface{} `json:"optionalValues,omitempty"`
	Time string `json:"time,omitempty"`
}

// DelegateIdentity represents the DelegateIdentity schema from the OpenAPI specification
type DelegateIdentity struct {
	Displayname string `json:"displayName,omitempty"`
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// ProblemDetails represents the ProblemDetails schema from the OpenAPI specification
type ProblemDetails struct {
	Status int `json:"status,omitempty"`
	Title string `json:"title,omitempty"`
	TypeField string `json:"type,omitempty"`
	Detail string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

// VnetProperties represents the VnetProperties schema from the OpenAPI specification
type VnetProperties struct {
	Subnetid string `json:"subnetId,omitempty"`
}

// CreateCloudEnvironmentBody represents the CreateCloudEnvironmentBody schema from the OpenAPI specification
type CreateCloudEnvironmentBody struct {
	Hasdevcontainerjson bool `json:"hasDevcontainerJson,omitempty"`
	Devcontainerjson string `json:"devContainerJson,omitempty"`
	Personalization PersonalizationInfoBody `json:"personalization,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Githubappurl string `json:"gitHubAppUrl,omitempty"`
	Experimentalfeatures ExperimentalFeaturesBody `json:"experimentalFeatures,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Friendlyname string `json:"friendlyName"`
	Label string `json:"label,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Netmoncorrelationdata NetmonCorrelationDataBody `json:"netmonCorrelationData,omitempty"`
	Identity IdentityBody `json:"identity,omitempty"`
	Secrets []SecretDataBody `json:"secrets,omitempty"`
	Githubapiurl string `json:"gitHubApiUrl,omitempty"`
	TypeField string `json:"type"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Connection ConnectionInfoBody `json:"connection,omitempty"`
	Containerimage string `json:"containerImage,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	Location string `json:"location,omitempty"`
	Platform string `json:"platform,omitempty"`
	Planid string `json:"planId,omitempty"`
	Workingdirectory string `json:"workingDirectory,omitempty"`
	Analyticstrackingid string `json:"analyticsTrackingId,omitempty"`
	Createasprebuild bool `json:"createAsPrebuild,omitempty"`
	Githubenvironmentendpoint string `json:"githubEnvironmentEndpoint,omitempty"`
	Runtimeconstraints RuntimeConstraintsBody `json:"runtimeConstraints,omitempty"`
	Usertier string `json:"userTier,omitempty"`
	Billableowner BillableOwnerBody `json:"billableOwner,omitempty"`
	Testaccount bool `json:"testAccount,omitempty"`
	Githubpfsauthendpoint string `json:"gitHubPfsAuthEndpoint,omitempty"`
}

// ExperimentalFeaturesBody represents the ExperimentalFeaturesBody schema from the OpenAPI specification
type ExperimentalFeaturesBody struct {
	Enabledynamichttpsdetection bool `json:"enableDynamicHttpsDetection,omitempty"`
	Queueresourceallocation bool `json:"queueResourceAllocation,omitempty"`
	Useprebuildfastpathifavailable bool `json:"usePrebuildFastPathIfAvailable,omitempty"`
	Useprebuiltimages bool `json:"usePrebuiltImages,omitempty"`
	Usestoragev2 bool `json:"useStorageV2,omitempty"`
}

// CreateTemplateResult represents the CreateTemplateResult schema from the OpenAPI specification
type CreateTemplateResult struct {
	Templateid string `json:"templateId,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Sasurl string `json:"sasUrl,omitempty"`
}

// ConnectionInfoBody represents the ConnectionInfoBody schema from the OpenAPI specification
type ConnectionInfoBody struct {
	Hostpublickeys []string `json:"hostPublicKeys,omitempty"`
	Relayendpoint string `json:"relayEndpoint,omitempty"`
	Relaysastoken string `json:"relaySasToken,omitempty"`
	Sessiontoken string `json:"sessionToken,omitempty"`
	Tunnelproperties TunnelProperties `json:"tunnelProperties,omitempty"`
	Connectionserviceuri string `json:"connectionServiceUri,omitempty"`
	Connectionsessionid string `json:"connectionSessionId,omitempty"`
	Connectionsessionpath string `json:"connectionSessionPath,omitempty"`
}

// NetworkSettingsResourceList represents the NetworkSettingsResourceList schema from the OpenAPI specification
type NetworkSettingsResourceList struct {
	Value []NetworkSettingsResource `json:"value,omitempty"`
}

// DomainUserCredentials represents the DomainUserCredentials schema from the OpenAPI specification
type DomainUserCredentials struct {
	Domain string `json:"domain"`
	Organizationalunit string `json:"organizationalUnit,omitempty"`
	Passwordsecretidentifier string `json:"passwordSecretIdentifier"`
	Username string `json:"userName"`
}

// PoolConfigRequestBody represents the PoolConfigRequestBody schema from the OpenAPI specification
type PoolConfigRequestBody struct {
	Maxtargetcount string `json:"maxTargetCount,omitempty"`
	Mintargetcount string `json:"minTargetCount,omitempty"`
	Poolcode string `json:"poolCode,omitempty"`
	Pooltype string `json:"poolType,omitempty"`
	Region string `json:"region,omitempty"`
	Targetcount string `json:"targetCount,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// TunnelProperties represents the TunnelProperties schema from the OpenAPI specification
type TunnelProperties struct {
	Serviceuri string `json:"serviceUri,omitempty"`
	Tunnelid string `json:"tunnelId,omitempty"`
	Tunnelname string `json:"tunnelName,omitempty"`
	Clusterid string `json:"clusterId,omitempty"`
	Connectaccesstoken string `json:"connectAccessToken,omitempty"`
	Domain string `json:"domain,omitempty"`
	Manageportsaccesstoken string `json:"managePortsAccessToken,omitempty"`
}

// ScopedUpdateSecretBody represents the ScopedUpdateSecretBody schema from the OpenAPI specification
type ScopedUpdateSecretBody struct {
	Filters []SecretFilterBody `json:"filters,omitempty"`
	Notes string `json:"notes,omitempty"`
	Scope int `json:"scope,omitempty"`
	Secretname string `json:"secretName,omitempty"`
	Value string `json:"value,omitempty"`
}

// SecretFilterBody represents the SecretFilterBody schema from the OpenAPI specification
type SecretFilterBody struct {
	TypeField int `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// HotPoolSettings represents the HotPoolSettings schema from the OpenAPI specification
type HotPoolSettings struct {
	Size int `json:"size,omitempty"`
}

// PlanResource represents the PlanResource schema from the OpenAPI specification
type PlanResource struct {
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	Properties PlanResourceProperties `json:"properties,omitempty"`
	Provisioningstate string `json:"provisioningState,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
	TypeField string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Identity PlanResourceIdentity `json:"identity,omitempty"`
}

// PlanResourceHeaders represents the PlanResourceHeaders schema from the OpenAPI specification
type PlanResourceHeaders struct {
	Hometenantid string `json:"homeTenantId,omitempty"`
	Identityprincipalid string `json:"identityPrincipalId,omitempty"`
	Identityurl string `json:"identityUrl,omitempty"`
	Clienttenantid string `json:"clientTenantId,omitempty"`
}

// VsoPlanInfo represents the VsoPlanInfo schema from the OpenAPI specification
type VsoPlanInfo struct {
	Subscription string `json:"subscription,omitempty"`
	Location int `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	Providernamespace string `json:"providerNamespace,omitempty"`
	Resourcegroup string `json:"resourceGroup,omitempty"`
	Resourceid string `json:"resourceId,omitempty"`
}

// LocationInfoResult represents the LocationInfoResult schema from the OpenAPI specification
type LocationInfoResult struct {
	Skus []SkuInfoResult `json:"skus,omitempty"`
}

// BillableOwnerBody represents the BillableOwnerBody schema from the OpenAPI specification
type BillableOwnerBody struct {
	TypeField int `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
}

// TenantInfoResult represents the TenantInfoResult schema from the OpenAPI specification
type TenantInfoResult struct {
	Id string `json:"id,omitempty"`
}

// NetworkSettingsResourceProperties represents the NetworkSettingsResourceProperties schema from the OpenAPI specification
type NetworkSettingsResourceProperties struct {
	Subnetid string `json:"subnetId,omitempty"`
}

// RepositoryInfoBody represents the RepositoryInfoBody schema from the OpenAPI specification
type RepositoryInfoBody struct {
	Createtype string `json:"createType,omitempty"`
	Prebuildhash string `json:"prebuildHash,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Name string `json:"name,omitempty"`
	Repoid int64 `json:"repoId,omitempty"`
	Diskusage string `json:"diskUsage,omitempty"`
	Owner string `json:"owner,omitempty"`
	Url string `json:"url,omitempty"`
	Commitid string `json:"commitId,omitempty"`
}

// SecretDataBody represents the SecretDataBody schema from the OpenAPI specification
type SecretDataBody struct {
	Value string `json:"value,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField int `json:"type,omitempty"`
}

// AvailableSettingsResult represents the AvailableSettingsResult schema from the OpenAPI specification
type AvailableSettingsResult struct {
	Sku []string `json:"sku,omitempty"`
}

// NotificationDataBody represents the NotificationDataBody schema from the OpenAPI specification
type NotificationDataBody struct {
	Message string `json:"message,omitempty"`
	Modal bool `json:"modal,omitempty"`
	Details string `json:"details,omitempty"`
	Displaymode string `json:"displayMode,omitempty"`
}

// PlanResourceList represents the PlanResourceList schema from the OpenAPI specification
type PlanResourceList struct {
	Value []PlanResource `json:"value,omitempty"`
}

// NetmonCorrelationDataBody represents the NetmonCorrelationDataBody schema from the OpenAPI specification
type NetmonCorrelationDataBody struct {
	Repositorycreatedat string `json:"repositoryCreatedAt,omitempty"`
	Billableownerdatabaseid string `json:"billableOwnerDatabaseId,omitempty"`
	Billableownerplan string `json:"billableOwnerPlan,omitempty"`
	Repositoryprivate bool `json:"repositoryPrivate,omitempty"`
	Ownercreatedat string `json:"ownerCreatedAt,omitempty"`
	Ownerdatabaseid string `json:"ownerDatabaseId,omitempty"`
	Ownerglobalrelayid string `json:"ownerGlobalRelayId,omitempty"`
	Repositorydatabaseid string `json:"repositoryDatabaseId,omitempty"`
	Repositoryglobalrelayid string `json:"repositoryGlobalRelayId,omitempty"`
	Billableownercreatedat string `json:"billableOwnerCreatedAt,omitempty"`
	Billableownerglobalrelayid string `json:"billableOwnerGlobalRelayId,omitempty"`
	Ownerplan string `json:"ownerPlan,omitempty"`
}

// SystemConfigurationResponse represents the SystemConfigurationResponse schema from the OpenAPI specification
type SystemConfigurationResponse struct {
	Comment string `json:"comment,omitempty"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// PoolDefinitionInput represents the PoolDefinitionInput schema from the OpenAPI specification
type PoolDefinitionInput struct {
	Targetcount int `json:"targetCount"`
	TypeField int `json:"type"`
	Dimensions map[string]interface{} `json:"dimensions"`
	Isenabled bool `json:"isEnabled"`
	Location int `json:"location"`
	Logicalskus []string `json:"logicalSkus,omitempty"`
	Subtype int `json:"subtype"`
}

// PrebuildEnvironmentOptions represents the PrebuildEnvironmentOptions schema from the OpenAPI specification
type PrebuildEnvironmentOptions struct {
	Correlationid string `json:"correlationId,omitempty"`
}
