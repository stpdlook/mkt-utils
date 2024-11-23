package main

// var today string = time.Now().Format("2006-01-02T15:04:05")

var (
	ExportFile       string = "bck.rsc"
	BackupFile       string = "bck.backup"
	ExportCommand    string = "export file=" + ExportFile
	BackupCommand    string = "system backup save name=" + BackupFile + " dont-encrypt=yes"
	BackupCommandPwd string = "system backup save name=" + BackupFile + " password="
	ExportFileName   string = "_" + GetDate() + "_" + ExportFile
	BackupFileName   string = "_" + GetDate() + "_" + BackupFile
)
