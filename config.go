package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func GetDate() string {
	switch runtime.GOOS {
	case "windows":
		dateFormat := "2006-01-02T15-04-05"
		Today := time.Now().Format(dateFormat)
		return Today
	default:
		dateFormat := "2006-01-02T15:04:05"
		Today := time.Now().Format(dateFormat)
		return Today
	}
}

type Config struct {
	IP               string
	PORT             uint
	LOGIN            string
	SSH_KEY_FILE     string
	SSH_KEY_PASSWORD string
	PASSWORD         string
	BCK_TYPE         string
	BCK_PATH         string
	BCK_PASSWORD     string
}

func appUsage() {
	fmt.Println("Утилита для создания резервных копий Mikrotik\n")
	fmt.Println("Доступные аргументы:")
	fmt.Println("\t--ip-address (обязателен)\n\t\tIP адрес устройства")
	fmt.Println("\t--port\n\t\tSSH порт, по умолчанию 22")
	fmt.Println("\t--login (обязателен)\n\t\tИмя пользователя для подключения к устройству")
	fmt.Println("\t--password (обязателен)\n\t\tПароль для подключения к устройству")
	fmt.Println("\t--ssh-key-file\n\t\tПуть до SSH ключа, используется вместо --password")
	fmt.Println("\t--bck-path (обязателен)\n\t\tПапка, куда будут сохраняться резервные копии")
	fmt.Println("\t--bck-password\n\t\tПароль для создания резервной копии (не обязателен)")
	fmt.Println("\t--bck-type (обязателен)\n\t\tВыбор формата резервной копии. Допустимые значения: export, backup, all.")
}

func parseFlags() (Config, error) {
	ip := flag.String("ip-address", "", "IP Адрес устройства")
	port := flag.Uint("port", 22, "SSH порт")
	login := flag.String("login", "", "Имя пользователя")
	sshKeyFile := flag.String("ssh-key-file", "", "Путь до ssh ключа")
	sshKeyPwd := flag.String("ssh-key-pwd", "", "Пароль от ssh ключа (при наличии)")
	password := flag.String("password", "", "Пароль от устройства")
	bckPath := flag.String("bck-path", "", "Папка, куда будут сохраняться резервные копии")
	bckPassword := flag.String("bck-password", "", "Пароль для шифрования резервной копии")
	bckType := flag.String("bck-type", "", "Выбор формата")

	flag.Usage = appUsage
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *ip == "" {
		log.Fatal("IP адрес устройства не указан (--ip-address)")
	}

	switch {
	case *bckType == "export":
		fmt.Println("Режим бэкапа: экспорт конфигурации")
	case *bckType == "backup":
		if *bckPassword != "" {
			fmt.Println("Режим бэкапа: резервное копирование (зашифрованный)")
		} else {
			fmt.Println("Режим бэкапа: резервное копирование")
		}
	case *bckType == "all":
		if *bckPassword != "" {
			fmt.Println("Режим бэкапа: комбинированный (зашифрованный)")
		} else {
			fmt.Println("Режим бэкапа: комбинированный")
		}
	default:
		log.Fatal("Недопустимый формат --bck-type. Допустимые значения: export, backup, all")
	}

	if *login == "" {
		log.Fatal("Имя пользователя не указано (--login)")
	}

	if *bckPath == "" {
		log.Fatal("Не задан путь для сохранения резервных копий (--bck-path)")
	}

	if *password == "" && *sshKeyFile == "" {
		log.Fatal("Не указан --ssh-key-file или --password")
	}

	if *password != "" && *sshKeyFile != "" {
		log.Println("Указан --ssh-key-file и --password, будет использоваться подключение по SSH ключу")
	}

	return Config{
		SSH_KEY_FILE:     *sshKeyFile,
		LOGIN:            *login,
		IP:               *ip,
		PORT:             *port,
		PASSWORD:         *password,
		BCK_PATH:         *bckPath,
		BCK_PASSWORD:     *bckPassword,
		SSH_KEY_PASSWORD: *sshKeyPwd,
		BCK_TYPE:         *bckType,
	}, nil
}
