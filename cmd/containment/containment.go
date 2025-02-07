package main

import (
	"fmt"

	"github.com/begtodfir/containment/internal/config"
	"github.com/begtodfir/containment/internal/detection"
	"github.com/begtodfir/containment/internal/logging"
)

func main() {
	printGreeting()

	cfg := config.Retrieve()

	logging.Setup(cfg.Logging)
	detection.Detect(cfg.Detection, cfg.Blacklist)

}

func printGreeting() {
	greeting := `
	..-::::::-..
.:-::::::::::::::-:.
._:::    ::    :::_.     _____             _        _                            _   
 .:( ^   :: ^   ):.     /  __ \           | |      (_)                          | |  
 .:::   (::)   :::.     | /  \/ ___  _ __ | |_ __ _ _ _ __  _ __ ___   ___ _ __ | |_ 
 .:::::::UU:::::::.     | |    / _ \| '_ \| __/ _' | | '_ \| '_ ' _ \ / _ \ '_ \| __|
 .::::::::::::::::.     | \__/\ (_) | | | | || (_| | | | | | | | | | |  __/ | | | |_ 
 -::::::::::::::::-      \____/\___/|_| |_|\__\__,_|_|_| |_|_| |_| |_|\___|_| |_|\__|
 .::::::::::::::::.
  .::::::::::::::.
    oO:::::::Oo

Welcome to Containment! Author: @begtodfir
	`

	fmt.Println(greeting)
}
