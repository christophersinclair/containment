# Containment - Linux Container Escape Detection Daemon

## Overview
Containment is a Linux daemon written in Go that monitors system calls and Linux namespaces to detect potential container escape attempts. It is designed to be lightweight and efficient while providing security monitoring for Kubernetes nodes and other containerized environments.

## Features
- Monitors syscalls for suspicious activity
- Tracks Linux namespaces to identify privilege escalation
- Configurable logging to file or systemd journal
- Minimal system overhead with eBPF support
- Runs as a systemd service with security hardening

## Installation

### Run the setup script
```bash
chmod +x setup.sh
./setup.sh
```
This script setups up the configurations, service, eBPF maps, etc.


## Configuration
Containment reads configurations from `/etc/containment.d/`. These configurations should be in normal INI format.

For example, the detection configuration would look like this:
```ini
[detection]
enable_syscalls = true    # Monitor suspicious syscalls (e.g., unshare, setns)
enable_network = true     # Detect unauthorized network activity
enable_mounts = true      # Monitor abnormal mounts inside containers
enable_proc_scans = true  # Scan /proc for hidden processes
```

## Security Considerations
- Runs with minimal privileges using `CapabilityBoundingSet`
- Uses `ProtectSystem=strict` to prevent unauthorized modifications
- `NoNewPrivileges=true` ensures no privilege escalation
- eBPF integration planned for efficient syscall tracing

## Future Enhancements
- More granular event filtering
- Kubernetes-aware security policies

## License
This project is licensed under the MIT License.

## Contact
For questions or support, reach out via [GitHub Issues](https://github.com/begtodfir/containment/issues).

