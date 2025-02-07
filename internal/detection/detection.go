package detection

import (
	"github.com/begtodfir/containment/internal/config"
	"github.com/begtodfir/containment/internal/logging"
)

func Detect(cfg *config.DetectionConfig, sysblacklist *config.BlacklistConfig) {
	if cfg.EnableSyscalls {
		logging.Get().LogMessage(logging.INFO, "Syscall monitoring enabled, starting subroutine...\n")
		go DetectSyscalls(sysblacklist)
	}

	if cfg.EnableNetwork {
		logging.Get().LogMessage(logging.INFO, "Network monitoring enabled, starting subroutine...\n")
		go DetectNetwork()
	}

	if cfg.EnableMounts {
		logging.Get().LogMessage(logging.INFO, "Mount monitoring enabled, starting subroutine...\n")
		go DetectMounts()
	}

	if cfg.EnableProc {
		logging.Get().LogMessage(logging.INFO, "Process monitoring enabled, starting subroutine...\n")
		go DetectProc()
	}
}
