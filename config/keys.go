// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

// #nosec G101
const (
	DataDirKey                                         = "data-dir"
	ConfigFileKey                                      = "config-file"
	ConfigContentKey                                   = "config-file-content"
	ConfigContentTypeKey                               = "config-file-content-type"
	VersionKey                                         = "version"
	GenesisConfigFileKey                               = "genesis"         // TODO: deprecated, remove
	GenesisConfigContentKey                            = "genesis-content" // TODO: deprecated, remove this
	GenesisFileKey                                     = "genesis-file"
	GenesisFileContentKey                              = "genesis-file-content"
	NetworkNameKey                                     = "network-id"
	TxFeeKey                                           = "tx-fee"
	CreateAssetTxFeeKey                                = "create-asset-tx-fee"
	CreateSubnetTxFeeKey                               = "create-subnet-tx-fee"
	TransformSubnetTxFeeKey                            = "transform-subnet-tx-fee"
	CreateBlockchainTxFeeKey                           = "create-blockchain-tx-fee"
	AddPrimaryNetworkValidatorFeeKey                   = "add-primary-network-validator-fee"
	AddPrimaryNetworkDelegatorFeeKey                   = "add-primary-network-delegator-fee"
	AddSubnetValidatorFeeKey                           = "add-subnet-validator-fee"
	AddSubnetDelegatorFeeKey                           = "add-subnet-delegator-fee"
	UptimeRequirementKey                               = "uptime-requirement"
	MinValidatorStakeKey                               = "min-validator-stake"
	MaxValidatorStakeKey                               = "max-validator-stake"
	MinDelegatorStakeKey                               = "min-delegator-stake"
	MinDelegatorFeeKey                                 = "min-delegation-fee"
	MinStakeDurationKey                                = "min-stake-duration"
	MaxStakeDurationKey                                = "max-stake-duration"
	StakeMaxConsumptionRateKey                         = "stake-max-consumption-rate"
	StakeMinConsumptionRateKey                         = "stake-min-consumption-rate"
	StakeMintingPeriodKey                              = "stake-minting-period"
	StakeSupplyCapKey                                  = "stake-supply-cap"
	DBTypeKey                                          = "db-type"
	DBPathKey                                          = "db-dir"
	DBConfigFileKey                                    = "db-config-file"
	DBConfigContentKey                                 = "db-config-file-content"
	PublicIPKey                                        = "public-ip"
	PublicIPResolutionFreqKey                          = "public-ip-resolution-frequency"
	PublicIPResolutionServiceKey                       = "public-ip-resolution-service"
	InboundConnUpgradeThrottlerCooldownKey             = "inbound-connection-throttling-cooldown"          // TODO: deprecated, remove this
	InboundThrottlerMaxConnsPerSecKey                  = "inbound-connection-throttling-max-conns-per-sec" // TODO: deprecated, remove this
	OutboundConnectionThrottlingRpsKey                 = "outbound-connection-throttling-rps"              // TODO: deprecated, remove this
	OutboundConnectionTimeoutKey                       = "outbound-connection-timeout"                     // TODO: deprecated, remove this
	HTTPHostKey                                        = "http-host"
	HTTPPortKey                                        = "http-port"
	HTTPSEnabledKey                                    = "http-tls-enabled"
	HTTPSKeyFileKey                                    = "http-tls-key-file"
	HTTPSKeyContentKey                                 = "http-tls-key-file-content"
	HTTPSCertFileKey                                   = "http-tls-cert-file"
	HTTPSCertContentKey                                = "http-tls-cert-file-content"
	HTTPAllowedOrigins                                 = "http-allowed-origins"
	HTTPShutdownTimeoutKey                             = "http-shutdown-timeout"
	HTTPShutdownWaitKey                                = "http-shutdown-wait"
	HTTPReadTimeoutKey                                 = "http-read-timeout"
	HTTPReadHeaderTimeoutKey                           = "http-read-header-timeout"
	HTTPWriteTimeoutKey                                = "http-write-timeout"
	HTTPIdleTimeoutKey                                 = "http-idle-timeout"
	APIAuthRequiredKey                                 = "api-auth-required"
	APIAuthPasswordKey                                 = "api-auth-password"
	APIAuthPasswordFileKey                             = "api-auth-password-file"
	StateSyncIPsKey                                    = "state-sync-ips"
	StateSyncIDsKey                                    = "state-sync-ids"
	BootstrapIPsKey                                    = "bootstrap-ips"
	BootstrapIDsKey                                    = "bootstrap-ids"
	StakingPortKey                                     = "staking-port"
	StakingEnabledKey                                  = "staking-enabled" // TODO: deprecated, remove this
	StakingEphemeralCertEnabledKey                     = "staking-ephemeral-cert-enabled"
	StakingTLSKeyPathKey                               = "staking-tls-key-file"
	StakingTLSKeyContentKey                            = "staking-tls-key-file-content"
	StakingCertPathKey                                 = "staking-tls-cert-file"
	StakingCertContentKey                              = "staking-tls-cert-file-content"
	StakingEphemeralSignerEnabledKey                   = "staking-ephemeral-signer-enabled"
	StakingSignerKeyPathKey                            = "staking-signer-key-file"
	StakingSignerKeyContentKey                         = "staking-signer-key-file-content"
	StakingDisabledWeightKey                           = "staking-disabled-weight" // TODO: deprecated, remove this
	SybilProtectionEnabledKey                          = "sybil-protection-enabled"
	SybilProtectionDisabledWeightKey                   = "sybil-protection-disabled-weight"
	NetworkInitialTimeoutKey                           = "network-initial-timeout"
	NetworkMinimumTimeoutKey                           = "network-minimum-timeout"
	NetworkMaximumTimeoutKey                           = "network-maximum-timeout"
	NetworkMaximumInboundTimeoutKey                    = "network-maximum-inbound-timeout"
	NetworkTimeoutHalflifeKey                          = "network-timeout-halflife"
	NetworkTimeoutCoefficientKey                       = "network-timeout-coefficient"
	NetworkHealthMinPeersKey                           = "network-health-min-conn-peers"
	NetworkHealthMaxTimeSinceMsgReceivedKey            = "network-health-max-time-since-msg-received"
	NetworkHealthMaxTimeSinceMsgSentKey                = "network-health-max-time-since-msg-sent"
	NetworkHealthMaxPortionSendQueueFillKey            = "network-health-max-portion-send-queue-full"
	NetworkHealthMaxSendFailRateKey                    = "network-health-max-send-fail-rate"
	NetworkHealthMaxOutstandingDurationKey             = "network-health-max-outstanding-request-duration"
	NetworkPeerListNumValidatorIPsKey                  = "network-peer-list-num-validator-ips"
	NetworkPeerListValidatorGossipSizeKey              = "network-peer-list-validator-gossip-size"
	NetworkPeerListNonValidatorGossipSizeKey           = "network-peer-list-non-validator-gossip-size"
	NetworkPeerListPeersGossipSizeKey                  = "network-peer-list-peers-gossip-size"
	NetworkPeerListGossipFreqKey                       = "network-peer-list-gossip-frequency"
	NetworkInitialReconnectDelayKey                    = "network-initial-reconnect-delay"
	NetworkReadHandshakeTimeoutKey                     = "network-read-handshake-timeout"
	NetworkPingTimeoutKey                              = "network-ping-timeout"
	NetworkPingFrequencyKey                            = "network-ping-frequency"
	NetworkMaxReconnectDelayKey                        = "network-max-reconnect-delay"
	NetworkCompressionEnabledKey                       = "network-compression-enabled" // TODO this is deprecated. Eventually remove it and constants.DefaultNetworkCompressionEnabled
	NetworkCompressionTypeKey                          = "network-compression-type"
	NetworkMaxClockDifferenceKey                       = "network-max-clock-difference"
	NetworkAllowPrivateIPsKey                          = "network-allow-private-ips"
	NetworkRequireValidatorToConnectKey                = "network-require-validator-to-connect"
	NetworkPeerReadBufferSizeKey                       = "network-peer-read-buffer-size"
	NetworkPeerWriteBufferSizeKey                      = "network-peer-write-buffer-size"
	NetworkTCPProxyEnabledKey                          = "network-tcp-proxy-enabled"
	NetworkTCPProxyReadTimeoutKey                      = "network-tcp-proxy-read-timeout"
	NetworkTLSKeyLogFileKey                            = "network-tls-key-log-file-unsafe"
	NetworkInboundConnUpgradeThrottlerCooldownKey      = "network-inbound-connection-throttling-cooldown"
	NetworkInboundThrottlerMaxConnsPerSecKey           = "network-inbound-connection-throttling-max-conns-per-sec"
	NetworkOutboundConnectionThrottlingRpsKey          = "network-outbound-connection-throttling-rps"
	NetworkOutboundConnectionTimeoutKey                = "network-outbound-connection-timeout"
	BenchlistFailThresholdKey                          = "benchlist-fail-threshold"
	BenchlistDurationKey                               = "benchlist-duration"
	BenchlistMinFailingDurationKey                     = "benchlist-min-failing-duration"
	LogsDirKey                                         = "log-dir"
	LogLevelKey                                        = "log-level"
	LogDisplayLevelKey                                 = "log-display-level"
	LogFormatKey                                       = "log-format"
	LogRotaterMaxSizeKey                               = "log-rotater-max-size"
	LogRotaterMaxFilesKey                              = "log-rotater-max-files"
	LogRotaterMaxAgeKey                                = "log-rotater-max-age"
	LogRotaterCompressEnabledKey                       = "log-rotater-compress-enabled"
	LogDisableDisplayPluginLogsKey                     = "log-disable-display-plugin-logs"
	SnowSampleSizeKey                                  = "snow-sample-size"
	SnowQuorumSizeKey                                  = "snow-quorum-size"
	SnowVirtuousCommitThresholdKey                     = "snow-virtuous-commit-threshold"
	SnowRogueCommitThresholdKey                        = "snow-rogue-commit-threshold"
	SnowAvalancheNumParentsKey                         = "snow-avalanche-num-parents"
	SnowAvalancheBatchSizeKey                          = "snow-avalanche-batch-size"
	SnowConcurrentRepollsKey                           = "snow-concurrent-repolls"
	SnowOptimalProcessingKey                           = "snow-optimal-processing"
	SnowMaxProcessingKey                               = "snow-max-processing"
	SnowMaxTimeProcessingKey                           = "snow-max-time-processing"
	SnowMixedQueryNumPushVdrKey                        = "snow-mixed-query-num-push-vdr"
	SnowMixedQueryNumPushNonVdrKey                     = "snow-mixed-query-num-push-non-vdr"
	TrackSubnetsKey                                    = "track-subnets"
	AdminAPIEnabledKey                                 = "api-admin-enabled"
	InfoAPIEnabledKey                                  = "api-info-enabled"
	KeystoreAPIEnabledKey                              = "api-keystore-enabled"
	MetricsAPIEnabledKey                               = "api-metrics-enabled"
	HealthAPIEnabledKey                                = "api-health-enabled"
	IpcAPIEnabledKey                                   = "api-ipcs-enabled"
	IpcsChainIDsKey                                    = "ipcs-chain-ids"
	IpcsPathKey                                        = "ipcs-path"
	MeterVMsEnabledKey                                 = "meter-vms-enabled"
	ConsensusGossipFrequencyKey                        = "consensus-gossip-frequency" // TODO: deprecated, remove this
	ConsensusAcceptedFrontierGossipFrequencyKey        = "consensus-accepted-frontier-gossip-frequency"
	ConsensusAppConcurrencyKey                         = "consensus-app-concurrency"
	ConsensusGossipAcceptedFrontierValidatorSizeKey    = "consensus-accepted-frontier-gossip-validator-size"
	ConsensusGossipAcceptedFrontierNonValidatorSizeKey = "consensus-accepted-frontier-gossip-non-validator-size"
	ConsensusGossipAcceptedFrontierPeerSizeKey         = "consensus-accepted-frontier-gossip-peer-size"
	ConsensusGossipOnAcceptValidatorSizeKey            = "consensus-on-accept-gossip-validator-size"
	ConsensusGossipOnAcceptNonValidatorSizeKey         = "consensus-on-accept-gossip-non-validator-size"
	ConsensusGossipOnAcceptPeerSizeKey                 = "consensus-on-accept-gossip-peer-size"
	AppGossipValidatorSizeKey                          = "consensus-app-gossip-validator-size"
	AppGossipNonValidatorSizeKey                       = "consensus-app-gossip-non-validator-size"
	AppGossipPeerSizeKey                               = "consensus-app-gossip-peer-size"
	ConsensusShutdownTimeoutKey                        = "consensus-shutdown-timeout"
	ProposerVMUseCurrentHeightKey                      = "proposervm-use-current-height"
	FdLimitKey                                         = "fd-limit"
	IndexEnabledKey                                    = "index-enabled"
	IndexAllowIncompleteKey                            = "index-allow-incomplete"
	RouterHealthMaxDropRateKey                         = "router-health-max-drop-rate"
	RouterHealthMaxOutstandingRequestsKey              = "router-health-max-outstanding-requests"
	HealthCheckFreqKey                                 = "health-check-frequency"
	HealthCheckAveragerHalflifeKey                     = "health-check-averager-halflife"
	RetryBootstrapKey                                  = "bootstrap-retry-enabled"
	RetryBootstrapWarnFrequencyKey                     = "bootstrap-retry-warn-frequency"
	PluginDirKey                                       = "plugin-dir"
	BootstrapBeaconConnectionTimeoutKey                = "bootstrap-beacon-connection-timeout"
	BootstrapMaxTimeGetAncestorsKey                    = "bootstrap-max-time-get-ancestors"
	BootstrapAncestorsMaxContainersSentKey             = "bootstrap-ancestors-max-containers-sent"
	BootstrapAncestorsMaxContainersReceivedKey         = "bootstrap-ancestors-max-containers-received"
	ChainDataDirKey                                    = "chain-data-dir"
	ChainConfigDirKey                                  = "chain-config-dir"
	ChainConfigContentKey                              = "chain-config-content"
	SubnetConfigDirKey                                 = "subnet-config-dir"
	SubnetConfigContentKey                             = "subnet-config-content"
	ProfileDirKey                                      = "profile-dir"
	ProfileContinuousEnabledKey                        = "profile-continuous-enabled"
	ProfileContinuousFreqKey                           = "profile-continuous-freq"
	ProfileContinuousMaxFilesKey                       = "profile-continuous-max-files"
	InboundThrottlerAtLargeAllocSizeKey                = "throttler-inbound-at-large-alloc-size"
	InboundThrottlerVdrAllocSizeKey                    = "throttler-inbound-validator-alloc-size"
	InboundThrottlerNodeMaxAtLargeBytesKey             = "throttler-inbound-node-max-at-large-bytes"
	InboundThrottlerMaxProcessingMsgsPerNodeKey        = "throttler-inbound-node-max-processing-msgs"
	InboundThrottlerBandwidthRefillRateKey             = "throttler-inbound-bandwidth-refill-rate"
	InboundThrottlerBandwidthMaxBurstSizeKey           = "throttler-inbound-bandwidth-max-burst-size"
	InboundThrottlerCPUMaxRecheckDelayKey              = "throttler-inbound-cpu-max-recheck-delay"
	InboundThrottlerDiskMaxRecheckDelayKey             = "throttler-inbound-disk-max-recheck-delay"
	CPUVdrAllocKey                                     = "throttler-inbound-cpu-validator-alloc"
	CPUMaxNonVdrUsageKey                               = "throttler-inbound-cpu-max-non-validator-usage"
	CPUMaxNonVdrNodeUsageKey                           = "throttler-inbound-cpu-max-non-validator-node-usage"
	SystemTrackerFrequencyKey                          = "system-tracker-frequency"
	SystemTrackerProcessingHalflifeKey                 = "system-tracker-processing-halflife"
	SystemTrackerCPUHalflifeKey                        = "system-tracker-cpu-halflife"
	SystemTrackerDiskHalflifeKey                       = "system-tracker-disk-halflife"
	SystemTrackerRequiredAvailableDiskSpaceKey         = "system-tracker-disk-required-available-space"
	SystemTrackerWarningThresholdAvailableDiskSpaceKey = "system-tracker-disk-warning-threshold-available-space"
	DiskVdrAllocKey                                    = "throttler-inbound-disk-validator-alloc"
	DiskMaxNonVdrUsageKey                              = "throttler-inbound-disk-max-non-validator-usage"
	DiskMaxNonVdrNodeUsageKey                          = "throttler-inbound-disk-max-non-validator-node-usage"
	OutboundThrottlerAtLargeAllocSizeKey               = "throttler-outbound-at-large-alloc-size"
	OutboundThrottlerVdrAllocSizeKey                   = "throttler-outbound-validator-alloc-size"
	OutboundThrottlerNodeMaxAtLargeBytesKey            = "throttler-outbound-node-max-at-large-bytes"
	UptimeMetricFreqKey                                = "uptime-metric-freq"
	VMAliasesFileKey                                   = "vm-aliases-file"
	VMAliasesContentKey                                = "vm-aliases-file-content"
	ChainAliasesFileKey                                = "chain-aliases-file"
	ChainAliasesContentKey                             = "chain-aliases-file-content"
	TracingEnabledKey                                  = "tracing-enabled"
	TracingEndpointKey                                 = "tracing-endpoint"
	TracingInsecureKey                                 = "tracing-insecure"
	TracingSampleRateKey                               = "tracing-sample-rate"
	TracingExporterTypeKey                             = "tracing-exporter-type"
)
