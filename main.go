package main

import (
	"log"
	"runtime"
	"strings"
)

func main() {

	config, err := parseFlags()
	if err != nil {
		return
	}

	switch runtime.GOOS {
	case "windows":
		if !(strings.HasSuffix(config.BCK_PATH, "\\")) {
			config.BCK_PATH = config.BCK_PATH + "\\"
		}
	default:
		if !(strings.HasSuffix(config.BCK_PATH, "/")) {
			config.BCK_PATH = config.BCK_PATH + "/"
		}
	}

	if config.SSH_KEY_FILE != "" {
		client, err := sshConnectKey(config.IP, config.LOGIN, config.SSH_KEY_FILE, config.SSH_KEY_PASSWORD, config.PORT)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case config.BCK_TYPE == "export" || config.BCK_TYPE == "all":
			startExportMikrotikConfig(client, config.BCK_PATH, config.IP)
		case config.BCK_TYPE == "backup" && config.BCK_PASSWORD != "" || config.BCK_TYPE == "all":
			createMikrotikBackupPwd(client, config.BCK_PATH, config.IP, config.BCK_PASSWORD)
		case config.BCK_TYPE == "backup" && config.BCK_PASSWORD == "" || config.BCK_TYPE == "all":
			createMikrotikBackup(client, config.BCK_PATH, config.IP)
		}
	} else {
		client, err := sshConnectPwd(config.IP, config.LOGIN, config.PASSWORD, config.PORT)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case config.BCK_TYPE == "export" || config.BCK_TYPE == "all":
			startExportMikrotikConfig(client, config.BCK_PATH, config.IP)
		case config.BCK_TYPE == "backup" && config.BCK_PASSWORD != "" || config.BCK_TYPE == "all":
			createMikrotikBackupPwd(client, config.BCK_PATH, config.IP, config.BCK_PASSWORD)
		case config.BCK_TYPE == "backup" && config.BCK_PASSWORD == "" || config.BCK_TYPE == "all":
			createMikrotikBackup(client, config.BCK_PATH, config.IP)
		}
	}
}
