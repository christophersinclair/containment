package config

const (
	configDir     = "/etc/containment.d/"
	detectionFile = "detection.conf"
	loggingFile   = "logging.conf"
	alertingFile  = "alerting.conf"
	responseFile  = "response.conf"
	blacklistFile = "blacklist.conf"

	DetectionSection = "detection"
	LoggingSection   = "logging"
	AlertingSection  = "alerting"
	ResponseSection  = "response"
	BlacklistSection = "blacklist"

	// Detection config keys
	DetectionEnableSyscalls = "enable_syscalls"
	DetectionEnableNetwork  = "enable_network"
	DetectionEnableMounts   = "enable_mounts"
	DetectionEnableProc     = "enable_proc_scans"

	// Blacklist config keys
	BlacklistSyscalls = "syscalls"

	// Response config keys
	ResponseThreshold    = "threshold"
	ResponseKillOnEscape = "kill_on_escape"
	ResponseQuarantine   = "quarantine"
	ResponseNotifyOnly   = "notify_only"

	// Logging config keys
	LoggingLevel  = "level"
	LoggingOutput = "output"

	// Alerting config keys
	AlertingThreshold = "threshold"
	AlertingWebhooks  = "webhook_list"
	AlertingEmails    = "email_list"
)
