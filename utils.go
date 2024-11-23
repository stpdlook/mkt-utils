package main

import (
	"log"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

func sshConnectPwd(ip, login, password string, port uint) (*goph.Client, error) {

	// client, err := goph.NewUnknown(login, ip, goph.Password(password))

	client, err := goph.NewConn(&goph.Config{
		User:     login,
		Addr:     ip,
		Port:     port,
		Auth:     goph.Password(password),
		Timeout:  goph.DefaultTimeout,
		Callback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		return nil, err
	}

	return client, err
}

func sshConnectKey(ip, login, sshKeyPath, sshKeyPwd string, port uint) (*goph.Client, error) {
	auth, err := goph.Key(sshKeyPath, sshKeyPwd)

	if err != nil {
		return nil, err
	}

	// client, err := goph.NewUnknown(login, ip, auth)
	client, err := goph.NewConn(&goph.Config{
		User:     login,
		Addr:     ip,
		Port:     port,
		Auth:     auth,
		Timeout:  goph.DefaultTimeout,
		Callback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		return nil, err
	}

	return client, err
}

func startExportMikrotikConfig(client *goph.Client, bckPath, ip string) {
	_, err := client.Run(ExportCommand)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Download("/bck.rsc", bckPath+ip+ExportFileName)
	if err != nil {
		log.Fatal(err)
	}
}

func createMikrotikBackup(client *goph.Client, bckPath, ip string) {
	_, err := client.Run(BackupCommand)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Download("/bck.backup", bckPath+ip+BackupFileName)
	if err != nil {
		log.Fatal(err)
	}
}

func createMikrotikBackupPwd(client *goph.Client, bckPath, ip, bckPwd string) {
	_, err := client.Run(BackupCommandPwd + bckPwd)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Download("/bck.backup", bckPath+ip+BackupFileName)
	if err != nil {
		log.Fatal(err)
	}
}
