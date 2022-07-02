package chainlink

import (
	"math/big"
	"net"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/libocr/commontypes"
	ocrnetworking "github.com/smartcontractkit/libocr/networking"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink/core/assets"
	coreconfig "github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/chainlink/core/store/dialects"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
)

// legacyGeneralConfig is a wrapper to adapt Config to the legacy config.GeneralConfig interface.
type legacyGeneralConfig struct {
	//TODO store original input w/o defaults too?
	c *Config
}

//TODO constructor which asserts non-nils?

func (l *legacyGeneralConfig) Validate() error {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) SetLogLevel(lvl zapcore.Level) error {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) SetLogSQL(logSQL bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) FeatureExternalInitiators() bool {
	return *l.c.JobPipeline.ExternalInitiatorsEnabled
}

func (l *legacyGeneralConfig) FeatureFeedsManager() bool {
	return *l.c.Feature.FeedsManager
}

func (l *legacyGeneralConfig) FeatureOffchainReporting() bool {
	return l.c.OCR != nil //TODO && l.c.P2P.NetworkingVersion() == ??
}

func (l *legacyGeneralConfig) FeatureOffchainReporting2() bool {
	return l.c.OCR2 != nil //TODO && l.c.P2P.NetworkingVersion() == ??
}

//TODO remove
func (l *legacyGeneralConfig) FeatureUICSAKeys() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) FeatureLogPoller() bool {
	return *l.c.Feature.LogPoller
}

func (l *legacyGeneralConfig) AutoPprofEnabled() bool {
	return *l.c.AutoPprof.Enabled
}

//TODO implied from chains? or master switch to flip them all off?
func (l *legacyGeneralConfig) EVMEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) EVMRPCEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) KeeperCheckUpkeepGasPriceFeatureEnabled() bool {
	return *l.c.Keeper.UpkeepCheckGasPriceEnabled
}

func (l *legacyGeneralConfig) P2PEnabled() bool {
	p := l.c.P2P
	return p.V1 != nil || p.V2 != nil //TODO or Disabled off switch?
}

func (l *legacyGeneralConfig) SolanaEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) TerraEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) StarkNetEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) AdvisoryLockCheckInterval() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) AdvisoryLockID() int64 {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) AllowOrigins() string {
	return *l.c.WebServer.AllowOrigins
}

func (l *legacyGeneralConfig) AppID() uuid.UUID {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) AuthenticatedRateLimit() int64 {
	return *l.c.WebServer.RateLimit.Authenticated
}

func (l *legacyGeneralConfig) AuthenticatedRateLimitPeriod() models.Duration {
	return *l.c.WebServer.RateLimit.AuthenticatedPeriod
}

func (l *legacyGeneralConfig) AutoPprofBlockProfileRate() int {
	return int(*l.c.AutoPprof.BlockProfileRate)
}

func (l *legacyGeneralConfig) AutoPprofCPUProfileRate() int {
	return int(*l.c.AutoPprof.CPUProfileRate)
}

func (l *legacyGeneralConfig) AutoPprofGatherDuration() models.Duration {
	return models.MustMakeDuration(l.c.AutoPprof.GatherDuration.Duration())
}

func (l *legacyGeneralConfig) AutoPprofGatherTraceDuration() models.Duration {
	return models.MustMakeDuration(l.c.AutoPprof.GatherTraceDuration.Duration())
}

func (l *legacyGeneralConfig) AutoPprofGoroutineThreshold() int {
	return int(*l.c.AutoPprof.GoroutineThreshold)
}

func (l *legacyGeneralConfig) AutoPprofMaxProfileSize() utils.FileSize {
	return *l.c.AutoPprof.MaxProfileSize
}

func (l *legacyGeneralConfig) AutoPprofMemProfileRate() int {
	return int(*l.c.AutoPprof.MemProfileRate)
}

func (l *legacyGeneralConfig) AutoPprofMemThreshold() utils.FileSize {
	return *l.c.AutoPprof.MemThreshold
}

func (l *legacyGeneralConfig) AutoPprofMutexProfileFraction() int {
	return int(*l.c.AutoPprof.MutexProfileFraction)
}

func (l *legacyGeneralConfig) AutoPprofPollInterval() models.Duration {
	return *l.c.AutoPprof.PollInterval
}

func (l *legacyGeneralConfig) AutoPprofProfileRoot() string {
	return *l.c.AutoPprof.ProfileRoot
}

func (l *legacyGeneralConfig) BlockBackfillDepth() uint64 {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) BlockBackfillSkip() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) BridgeResponseURL() *url.URL {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) CertFile() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) DatabaseBackupDir() string {
	return *l.c.Database.Backup.Dir
}

func (l *legacyGeneralConfig) DatabaseBackupFrequency() time.Duration {
	return l.c.Database.Backup.Frequency.Duration()
}

func (l *legacyGeneralConfig) DatabaseBackupMode() coreconfig.DatabaseBackupMode {
	return *l.c.Database.Backup.Mode
}

func (l *legacyGeneralConfig) DatabaseBackupOnVersionUpgrade() bool {
	return *l.c.Database.Backup.OnVersionUpgrade
}

func (l *legacyGeneralConfig) DatabaseBackupURL() *url.URL {
	return (*url.URL)(l.c.Database.Backup.URL)
}

func (l *legacyGeneralConfig) DatabaseListenerMaxReconnectDuration() time.Duration {
	return l.c.Database.Listener.MaxReconnectDuration.Duration()
}

func (l *legacyGeneralConfig) DatabaseListenerMinReconnectInterval() time.Duration {
	return l.c.Database.Listener.MinReconnectInterval.Duration()
}

func (l *legacyGeneralConfig) DatabaseLockingMode() string {
	return *l.c.Database.Lock.Mode
}

func (l *legacyGeneralConfig) DatabaseURL() url.URL {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) DefaultChainID() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) DefaultHTTPLimit() int64 {
	return int64(*l.c.JobPipeline.HTTPRequestMaxSize)
}

func (l *legacyGeneralConfig) DefaultHTTPTimeout() models.Duration {
	return *l.c.JobPipeline.DefaultHTTPRequestTimeout
}

func (l *legacyGeneralConfig) DefaultLogLevel() zapcore.Level {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) Dev() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) ShutdownGracePeriod() time.Duration {
	return l.c.ShutdownGracePeriod.Duration()
}

func (l *legacyGeneralConfig) EthereumHTTPURL() *url.URL {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) EthereumNodes() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) EthereumSecondaryURLs() []url.URL {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) EthereumURL() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) ExplorerAccessKey() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) ExplorerSecret() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) ExplorerURL() *url.URL {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) FMDefaultTransactionQueueDepth() uint32 {
	return *l.c.FluxMonitor.DefaultTransactionQueueDepth
}

func (l *legacyGeneralConfig) FMSimulateTransactions() bool {
	return *l.c.FluxMonitor.SimulateTransactions
}

func (l *legacyGeneralConfig) GetAdvisoryLockIDConfiguredOrDefault() int64 {
	return *l.c.Database.Lock.AdvisoryID
}

func (l *legacyGeneralConfig) GetDatabaseDialectConfiguredOrDefault() dialects.DialectName {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) HTTPServerWriteTimeout() time.Duration {
	return l.c.WebServer.HTTPWriteTimeout.Duration()
}

func (l *legacyGeneralConfig) InsecureFastScrypt() bool {
	return *l.c.InsecureFastScrypt
}

func (l *legacyGeneralConfig) JSONConsole() bool {
	return *l.c.Log.JSONConsole
}

func (l *legacyGeneralConfig) JobPipelineMaxRunDuration() time.Duration {
	return l.c.JobPipeline.MaxRunDuration.Duration()
}

func (l *legacyGeneralConfig) JobPipelineReaperInterval() time.Duration {
	return l.c.JobPipeline.ReaperInterval.Duration()
}

func (l *legacyGeneralConfig) JobPipelineReaperThreshold() time.Duration {
	return l.c.JobPipeline.ReaperThreshold.Duration()
}

func (l *legacyGeneralConfig) JobPipelineResultWriteQueueDepth() uint64 {
	return uint64(*l.c.JobPipeline.ResultWriteQueueDepth)
}

func (l *legacyGeneralConfig) KeeperDefaultTransactionQueueDepth() uint32 {
	return *l.c.Keeper.DefaultTransactionQueueDepth
}

func (l *legacyGeneralConfig) KeeperGasPriceBufferPercent() uint32 {
	return *l.c.Keeper.GasPriceBufferPercent
}

func (l *legacyGeneralConfig) KeeperGasTipCapBufferPercent() uint32 {
	return *l.c.Keeper.GasTipCapBufferPercent
}

func (l *legacyGeneralConfig) KeeperBaseFeeBufferPercent() uint32 {
	return *l.c.Keeper.BaseFeeBufferPercent
}

func (l *legacyGeneralConfig) KeeperMaximumGracePeriod() int64 {
	return *l.c.Keeper.MaximumGracePeriod
}

func (l *legacyGeneralConfig) KeeperRegistryCheckGasOverhead() uint64 {
	return l.c.Keeper.RegistryCheckGasOverhead.ToInt().Uint64()
}

func (l *legacyGeneralConfig) KeeperRegistryPerformGasOverhead() uint64 {
	return l.c.Keeper.RegistryPerformGasOverhead.ToInt().Uint64()
}

func (l *legacyGeneralConfig) KeeperRegistrySyncInterval() time.Duration {
	return l.c.Keeper.RegistrySyncInterval.Duration()
}

func (l *legacyGeneralConfig) KeeperRegistrySyncUpkeepQueueSize() uint32 {
	return *l.c.Keeper.RegistrySyncUpkeepQueueSize
}

func (l *legacyGeneralConfig) KeeperTurnLookBack() int64 {
	return *l.c.Keeper.TurnLookBack
}

func (l *legacyGeneralConfig) KeeperTurnFlagEnabled() bool {
	return *l.c.Keeper.TurnFlagEnabled
}

func (l *legacyGeneralConfig) KeyFile() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) LeaseLockDuration() time.Duration {
	return l.c.Database.Lock.LeaseDuration.Duration()
}

func (l *legacyGeneralConfig) LeaseLockRefreshInterval() time.Duration {
	return l.c.Database.Lock.LeaseRefreshInterval.Duration()
}

func (l *legacyGeneralConfig) LogFileDir() string {
	return *l.c.Log.FileDir
}

func (l *legacyGeneralConfig) LogLevel() zapcore.Level {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) LogSQL() bool {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) LogFileMaxSize() utils.FileSize {
	return *l.c.Log.FileMaxSize
}

func (l *legacyGeneralConfig) LogFileMaxAge() int64 {
	return *l.c.Log.FileMaxAgeDays
}

func (l *legacyGeneralConfig) LogFileMaxBackups() int64 {
	return *l.c.Log.FileMaxBackups
}

func (l *legacyGeneralConfig) LogUnixTimestamps() bool {
	return *l.c.Log.UnixTS
}

func (l *legacyGeneralConfig) MigrateDatabase() bool {
	return *l.c.Database.MigrateOnStartup
}

func (l *legacyGeneralConfig) ORMMaxIdleConns() int {
	return int(*l.c.Database.ORMMaxIdleConns)
}

func (l *legacyGeneralConfig) ORMMaxOpenConns() int {
	return int(*l.c.Database.ORMMaxOpenConns)
}

func (l *legacyGeneralConfig) Port() uint16 {
	return *l.c.WebServer.HTTPPort
}

func (l *legacyGeneralConfig) RPID() string {
	return *l.c.WebServer.MFA.RPID
}

func (l *legacyGeneralConfig) RPOrigin() string {
	return *l.c.WebServer.MFA.RPOrigin
}

func (l *legacyGeneralConfig) ReaperExpiration() models.Duration {
	return *l.c.ReaperExpiration
}

func (l *legacyGeneralConfig) RootDir() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) SecureCookies() bool {
	return *l.c.WebServer.SecureCookies
}

func (l *legacyGeneralConfig) SessionOptions() sessions.Options {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) SessionSecret() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) SessionTimeout() models.Duration {
	return models.MustMakeDuration(l.c.WebServer.SessionTimeout.Duration())
}

func (l *legacyGeneralConfig) SolanaNodes() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) TerraNodes() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) TLSCertPath() string {
	return *l.c.WebServer.TLS.CertPath
}

func (l *legacyGeneralConfig) TLSDir() string {
	return filepath.Join(*l.c.RootDir, "tls")
}

func (l *legacyGeneralConfig) TLSHost() string {
	return *l.c.WebServer.TLS.Host
}

func (l *legacyGeneralConfig) TLSKeyPath() string {
	return *l.c.WebServer.TLS.KeyPath
}

func (l *legacyGeneralConfig) TLSPort() uint16 {
	return *l.c.WebServer.TLS.HTTPSPort
}

func (l *legacyGeneralConfig) TLSRedirect() bool {
	return *l.c.WebServer.TLS.ForceRedirect
}

func (l *legacyGeneralConfig) TelemetryIngressLogging() bool {
	return *l.c.TelemetryIngress.Logging
}

func (l *legacyGeneralConfig) TelemetryIngressUniConn() bool {
	return *l.c.TelemetryIngress.UniConn
}

func (l *legacyGeneralConfig) TelemetryIngressServerPubKey() string {
	return *l.c.TelemetryIngress.ServerPubKey
}

func (l *legacyGeneralConfig) TelemetryIngressURL() *url.URL {
	return (*url.URL)(l.c.TelemetryIngress.URL)
}

func (l *legacyGeneralConfig) TelemetryIngressBufferSize() uint {
	return uint(*l.c.TelemetryIngress.BufferSize)
}

func (l *legacyGeneralConfig) TelemetryIngressMaxBatchSize() uint {
	return uint(*l.c.TelemetryIngress.MaxBatchSize)
}

func (l *legacyGeneralConfig) TelemetryIngressSendInterval() time.Duration {
	return l.c.TelemetryIngress.SendInterval.Duration()
}

func (l *legacyGeneralConfig) TelemetryIngressSendTimeout() time.Duration {
	return l.c.TelemetryIngress.SendTimeout.Duration()
}

func (l *legacyGeneralConfig) TelemetryIngressUseBatchSend() bool {
	return *l.c.TelemetryIngress.UseBatchSend
}

func (l *legacyGeneralConfig) TriggerFallbackDBPollInterval() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) UnAuthenticatedRateLimit() int64 {
	return *l.c.WebServer.RateLimit.Unauthenticated
}

func (l *legacyGeneralConfig) UnAuthenticatedRateLimitPeriod() models.Duration {
	return *l.c.WebServer.RateLimit.UnauthenticatedPeriod
}

func (l *legacyGeneralConfig) GlobalBalanceMonitorEnabled() (bool, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalBlockEmissionIdleWarningThreshold() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalBlockHistoryEstimatorBatchSize() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalBlockHistoryEstimatorBlockDelay() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalBlockHistoryEstimatorBlockHistorySize() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalBlockHistoryEstimatorEIP1559FeeCapBufferBlocks() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalBlockHistoryEstimatorTransactionPercentile() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalChainType() (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEthTxReaperInterval() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEthTxReaperThreshold() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEthTxResendAfterThreshold() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmEIP1559DynamicFees() (bool, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmFinalityDepth() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasBumpPercent() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasBumpThreshold() (uint64, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasBumpTxDepth() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasBumpWei() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasFeeCapDefault() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasLimitDefault() (uint64, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasLimitMultiplier() (float32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasLimitTransfer() (uint64, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasPriceDefault() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasTipCapDefault() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmGasTipCapMinimum() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmHeadTrackerHistoryDepth() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmHeadTrackerMaxBufferSize() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmHeadTrackerSamplingInterval() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmLogBackfillBatchSize() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmLogPollInterval() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmMaxGasPriceWei() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmMaxInFlightTransactions() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmMaxQueuedTransactions() (uint64, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmMinGasPriceWei() (*big.Int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmNonceAutoSync() (bool, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmUseForwarders() (bool, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalEvmRPCDefaultBatchSize() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalFlagsContractAddress() (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalGasEstimatorMode() (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalLinkContractAddress() (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalOperatorFactoryAddress() (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalMinIncomingConfirmations() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalMinimumContractPayment() (*assets.Link, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalNodeNoNewHeadsThreshold() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalNodePollFailureThreshold() (uint32, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalNodePollInterval() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalOCRContractConfirmations() (uint16, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalOCRContractTransmitterTransmitTimeout() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalOCRDatabaseTimeout() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) GlobalOCRObservationGracePeriod() (time.Duration, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) OCRBlockchainTimeout() time.Duration {
	return l.c.OCR.BlockchainTimeout.Duration()
}

func (l *legacyGeneralConfig) OCRContractPollInterval() time.Duration {
	return l.c.OCR.ContractPollInterval.Duration()
}

func (l *legacyGeneralConfig) OCRContractSubscribeInterval() time.Duration {
	return l.c.OCR.ContractSubscribeInterval.Duration()
}

func (l *legacyGeneralConfig) OCRMonitoringEndpoint() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) OCRKeyBundleID() (string, error) {
	return l.c.OCR.KeyBundleID.String(), nil
}

func (l *legacyGeneralConfig) OCRObservationTimeout() time.Duration {
	return l.c.OCR.ObservationTimeout.Duration()
}

func (l *legacyGeneralConfig) OCRSimulateTransactions() bool {
	return *l.c.OCR.SimulateTransactions
}

func (l *legacyGeneralConfig) OCRTransmitterAddress() (ethkey.EIP55Address, error) {
	return *l.c.OCR.TransmitterAddress, nil
}

func (l *legacyGeneralConfig) OCRTraceLogging() bool {
	return *l.c.P2P.TraceLogging
}

func (l *legacyGeneralConfig) OCRDefaultTransactionQueueDepth() uint32 {
	return *l.c.OCR.DefaultTransactionQueueDepth
}

func (l *legacyGeneralConfig) OCR2ContractConfirmations() uint16 {
	return uint16(*l.c.OCR2.ContractConfirmations)
}

func (l *legacyGeneralConfig) OCR2ContractTransmitterTransmitTimeout() time.Duration {
	return l.c.OCR2.ContractTransmitterTransmitTimeout.Duration()
}

func (l *legacyGeneralConfig) OCR2BlockchainTimeout() time.Duration {
	return l.c.OCR2.BlockchainTimeout.Duration()
}

func (l *legacyGeneralConfig) OCR2DatabaseTimeout() time.Duration {
	return l.c.OCR2.DatabaseTimeout.Duration()
}

func (l *legacyGeneralConfig) OCR2ContractPollInterval() time.Duration {
	return l.c.OCR2.ContractPollInterval.Duration()
}

func (l *legacyGeneralConfig) OCR2ContractSubscribeInterval() time.Duration {
	return l.c.OCR2.ContractSubscribeInterval.Duration()
}

func (l *legacyGeneralConfig) OCR2MonitoringEndpoint() string {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) OCR2KeyBundleID() (string, error) {
	return l.c.OCR2.KeyBundleID.String(), nil
}

func (l *legacyGeneralConfig) OCR2TraceLogging() bool {
	return *l.c.P2P.TraceLogging
}

func (l *legacyGeneralConfig) P2PNetworkingStack() (n ocrnetworking.NetworkingStack) {
	return l.c.P2P.NetworkStack()
}

func (l *legacyGeneralConfig) P2PNetworkingStackRaw() string {
	return l.c.P2P.NetworkStack().String()
}

func (l *legacyGeneralConfig) P2PPeerID() p2pkey.PeerID {
	return *l.c.P2P.V1.PeerID
}

func (l *legacyGeneralConfig) P2PPeerIDRaw() string {
	return l.c.P2P.V1.PeerID.String()
}

func (l *legacyGeneralConfig) P2PIncomingMessageBufferSize() int {
	return int(*l.c.P2P.IncomingMessageBufferSize)
}

func (l *legacyGeneralConfig) P2POutgoingMessageBufferSize() int {
	return int(*l.c.P2P.OutgoingMessageBufferSize)
}

func (l *legacyGeneralConfig) OCRNewStreamTimeout() time.Duration {
	return l.c.P2P.V1.NewStreamTimeout.Duration()
}

func (l *legacyGeneralConfig) OCRBootstrapCheckInterval() time.Duration {
	return l.c.P2P.V1.BootstrapCheckInterval.Duration()
}

func (l *legacyGeneralConfig) OCRDHTLookupInterval() int {
	return int(*l.c.P2P.V1.DHTLookupInterval)
}

func (l *legacyGeneralConfig) OCRIncomingMessageBufferSize() int {
	return int(*l.c.P2P.IncomingMessageBufferSize)
}

func (l *legacyGeneralConfig) OCROutgoingMessageBufferSize() int {
	return int(*l.c.P2P.OutgoingMessageBufferSize)
}

func (l *legacyGeneralConfig) P2PAnnounceIP() net.IP {
	return *l.c.P2P.V1.AnnounceIP
}

func (l *legacyGeneralConfig) P2PAnnouncePort() uint16 {
	return *l.c.P2P.V1.AnnouncePort
}

func (l *legacyGeneralConfig) P2PBootstrapPeers() ([]string, error) {
	return *l.c.P2P.V1.DefaultBootstrapPeers, nil
}

func (l *legacyGeneralConfig) P2PDHTAnnouncementCounterUserPrefix() uint32 {
	return *l.c.P2P.V1.DHTAnnouncementCounterUserPrefix
}

func (l *legacyGeneralConfig) P2PListenIP() net.IP {
	return *l.c.P2P.V1.ListenIP
}

func (l *legacyGeneralConfig) P2PListenPort() uint16 {
	return *l.c.P2P.V1.ListenPort
}

func (l *legacyGeneralConfig) P2PListenPortRaw() string {
	return strconv.Itoa(int(*l.c.P2P.V1.ListenPort))
}

func (l *legacyGeneralConfig) P2PNewStreamTimeout() time.Duration {
	return l.c.P2P.V1.NewStreamTimeout.Duration()
}

func (l *legacyGeneralConfig) P2PBootstrapCheckInterval() time.Duration {
	return l.c.P2P.V1.BootstrapCheckInterval.Duration()
}

func (l *legacyGeneralConfig) P2PDHTLookupInterval() int {
	return int(*l.c.P2P.V1.DHTLookupInterval)
}

func (l *legacyGeneralConfig) P2PPeerstoreWriteInterval() time.Duration {
	return l.c.P2P.V1.PeerstoreWriteInterval.Duration()
}

func (l *legacyGeneralConfig) P2PV2AnnounceAddresses() []string {
	if p := l.c.P2P; p != nil {
		if v2 := p.V2; v2 != nil {
			if v := v2.AnnounceAddresses; v != nil {
				return *v
			}
		}
	}
	return nil
}

func (l *legacyGeneralConfig) P2PV2Bootstrappers() (locators []commontypes.BootstrapperLocator) {
	if p := l.c.P2P; p != nil {
		if v2 := p.V2; v2 != nil {
			if v := v2.DefaultBootstrappers; v != nil {
				return *v
			}
		}
	}
	return nil
}

func (l *legacyGeneralConfig) P2PV2BootstrappersRaw() []string {
	if p := l.c.P2P; p != nil {
		if v2 := p.V2; v2 != nil {
			if v := v2.DefaultBootstrappers; v != nil {
				var s []string
				for _, b := range *v {
					t, err := b.MarshalText()
					if err != nil {
						//TODO log panic matches old behavior.... only called for UI presentation anyways....
					}
					s = append(s, string(t))
				}
			}
		}
	}
	return nil
}

func (l *legacyGeneralConfig) P2PV2DeltaDial() models.Duration {
	if p := l.c.P2P; p != nil {
		if v2 := p.V2; v2 != nil {
			if v := v2.DeltaDial; v != nil {
				return *v
			}
		}
	}
	return models.Duration{}
}

func (l *legacyGeneralConfig) P2PV2DeltaReconcile() models.Duration {
	if p := l.c.P2P; p != nil {
		if v2 := p.V2; v2 != nil {
			if v := v2.DeltaReconcile; v != nil {
				return *v
			}
		}
	}
	return models.Duration{}
}

func (l *legacyGeneralConfig) P2PV2ListenAddresses() []string {
	if p := l.c.P2P; p != nil {
		if v2 := p.V2; v2 != nil {
			if v := v2.ListenAddresses; v != nil {
				return *v
			}
		}
	}
	return nil
}
