package config

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

type ContainmentConfig struct {
	Detection *DetectionConfig
	Blacklist *BlacklistConfig
	Logging   *LoggingConfig
	Response  *ResponseConfig
	Alerting  *AlertingConfig
}

type DetectionConfig struct {
	EnableSyscalls bool
	EnableNetwork  bool
	EnableMounts   bool
	EnableProc     bool
}

type BlacklistConfig struct {
	Syscalls []string
}

type LoggingConfig struct {
	Level   string
	OutFile string
}

type ResponseConfig struct {
	Threshold    int
	KillOnEscape bool
	Quarantine   bool
	NotifyOnly   bool
}

type AlertingConfig struct {
	Threshold int
	Webhooks  []string
	Emails    []string
}

func Retrieve() *ContainmentConfig {
	detectionConfig, err := loadDetection(fmt.Sprintf("%s%s", configDir, detectionFile))
	if err != nil {
		log.Fatalf("Failed to read config file %s: %v", detectionFile, err)
	}

	blacklistConfig, err := loadBlacklist(fmt.Sprintf("%s%s", configDir, blacklistFile))
	if err != nil {
		log.Fatalf("Failed to read config file %s: %v", blacklistFile, err)
	}

	responseConfig, err := loadResponse(fmt.Sprintf("%s%s", configDir, responseFile))
	if err != nil {
		log.Fatalf("Failed to read config file: %s: %v", responseFile, err)
	}

	loggingConfig, err := loadLogging(fmt.Sprintf("%s%s", configDir, loggingFile))
	if err != nil {
		log.Fatalf("Failed to read config file: %s: %v", loggingFile, err)
	}

	alertingConfg, err := loadAlerting(fmt.Sprintf("%s%s", configDir, alertingFile))
	if err != nil {
		log.Fatalf("Failed to read config file: %s: %v", alertingFile, err)
	}

	return &ContainmentConfig{
		Detection: detectionConfig,
		Blacklist: blacklistConfig,
		Response:  responseConfig,
		Logging:   loggingConfig,
		Alerting:  alertingConfg,
	}
}

func loadDetection(filename string) (*DetectionConfig, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	syscallEnable, err := cfg.Section(DetectionSection).Key(DetectionEnableSyscalls).Bool()
	if err != nil {
		return nil, err
	}

	networkEnable, err := cfg.Section(DetectionSection).Key(DetectionEnableNetwork).Bool()
	if err != nil {
		return nil, err
	}

	mountEnable, err := cfg.Section(DetectionSection).Key(DetectionEnableMounts).Bool()
	if err != nil {
		return nil, err
	}

	procEnable, err := cfg.Section(DetectionSection).Key(DetectionEnableProc).Bool()
	if err != nil {
		return nil, err
	}

	return &DetectionConfig{
		EnableSyscalls: syscallEnable,
		EnableNetwork:  networkEnable,
		EnableMounts:   mountEnable,
		EnableProc:     procEnable,
	}, nil
}

func loadBlacklist(filename string) (*BlacklistConfig, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	syscalls := cfg.Section(BlacklistSection).Key(BlacklistSyscalls).String()
	syscallList := strings.Split(syscalls, ",")

	for i := range syscallList {
		syscallList[i] = strings.TrimSpace(syscallList[i])
	}

	return &BlacklistConfig{
		Syscalls: syscallList,
	}, nil
}

func loadResponse(filename string) (*ResponseConfig, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	threshold, err := cfg.Section(ResponseSection).Key(ResponseThreshold).Int()
	if err != nil {
		return nil, err
	}

	kill, err := cfg.Section(ResponseSection).Key(ResponseKillOnEscape).Bool()
	if err != nil {
		return nil, err
	}

	quarantine, err := cfg.Section(ResponseSection).Key(ResponseQuarantine).Bool()
	if err != nil {
		return nil, err
	}

	notifyOnly, err := cfg.Section(ResponseSection).Key(ResponseNotifyOnly).Bool()
	if err != nil {
		return nil, err
	}

	return &ResponseConfig{
		Threshold:    threshold,
		KillOnEscape: kill,
		Quarantine:   quarantine,
		NotifyOnly:   notifyOnly,
	}, nil
}

func loadLogging(filename string) (*LoggingConfig, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	level := cfg.Section(LoggingSection).Key(LoggingLevel).String()

	output := cfg.Section(LoggingSection).Key(LoggingOutput).String()

	return &LoggingConfig{
		Level:   level,
		OutFile: output,
	}, nil
}

func loadAlerting(filename string) (*AlertingConfig, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	threshold, err := cfg.Section(AlertingSection).Key(AlertingThreshold).Int()
	if err != nil {
		return nil, err
	}

	webhooks := cfg.Section(AlertingSection).Key(AlertingWebhooks).String()
	webhookList := strings.Split(webhooks, ",")

	for i := range webhookList {
		webhookList[i] = strings.TrimSpace(webhookList[i])
	}

	emails := cfg.Section(AlertingSection).Key(AlertingEmails).String()
	emailList := strings.Split(emails, ",")

	for i := range emailList {
		emailList[i] = strings.TrimSpace(emailList[i])
	}

	return &AlertingConfig{
		Threshold: threshold,
		Webhooks:  webhookList,
		Emails:    emailList,
	}, nil

}
