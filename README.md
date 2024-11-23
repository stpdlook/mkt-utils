# Утилита для создания резервных копий Mikrotik

## Примеры использования

Экспорт конфигурации. Создает в папке `~/Documents` `.rsc` файл со всей конфигурацией.
```bash
./mkt-utils --ip-address 172.25.0.1 --login admin --password "Pa$w@rd" --bck-path ~/Documents --bck-type export
```

Подключение к нестандартному порту по SSH ключу.
```bash
./mkt-utils --ip-address 172.25.0.1 --port 2222 --login admin --ssh-key-file ~/.ssh/id_ed25519 --bck-path ~/Documents --bck-type export
```

Создание резервной копии.
```
./mkt-utils --ip-address 172.25.0.1 --port 2222 --login admin --ssh-key-file ~/.ssh/id_ed25519 --bck-path ~/Documents --bck-type backup
```

Создание зашифрованной резервной копии.
```
./mkt-utils --ip-address 172.25.0.1 --port 2222 --login admin --ssh-key-file ~/.ssh/id_ed25519 --bck-path ~/Documents --bck-type backup bck-password "S&perSecr@t"
```