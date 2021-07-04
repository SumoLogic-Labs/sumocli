package cmdutils

func ValidateCapabilities(capability string) bool {
	switch capability {
	case
		"viewCollectors",
		"manageCollectors",
		"manageBudgets",
		"manageDataVolumeFeed",
		"viewFieldExtraction",
		"manageFieldExtractionRules",
		"manageS3DataForwarding",
		"manageContent",
		"dataVolumeIndex",
		"viewConnections",
		"manageConnections",
		"viewScheduledViews",
		"manageScheduledViews",
		"viewPartitions",
		"managePartitions",
		"viewFields",
		"manageFields",
		"viewAccountOverview",
		"manageTokens",
		"manageDataStreams",
		"manageEntityTypeConfig",
		"manageMonitors",
		"metricsTransformation",
		"metricsExtraction",
		"metricsRules",
		"managePasswordPolicy",
		"ipWhitelisting",
		"createAccessKeys",
		"manageAccessKeys",
		"manageSupportAccountAccess",
		"manageAuditDataFeed",
		"manageSaml",
		"shareDashboardOutsideOrg",
		"manageOrgSettings",
		"changeDataAccessLevel",
		"shareDashboardWorld",
		"shareDashboardWhitelist",
		"manageUsersAndRoles",
		"searchAuditIndex",
		"auditEventIndex":
		return true
	}
	return false
}
