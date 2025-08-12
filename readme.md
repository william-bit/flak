# Flak - The Dev Environment manager

A portable, isolated, fast & powerful development environment for web development. It is fast, lightweight and easy-to-use

Enjoy!

## Quick Add

### Runtime and service

- PHP
- NodeJS
- Java
- Python
- Golang
- Ruby
- Elixir

### Server

- Nginx
- Apache
- Tomcat

### Database

- MySQL
- MariaDB
- PostgreSQL
- MonggoDB
- PocketBase
- Memcached
- Redis

### Message broker

- RabbitMQ

### Devtool

- Git
- OpenSSL
- Telnet
- MailPit
- TortoiseGit
- FileZilla
- Pandoc

### DBMS

- heidisql
- PHPMyAdmin
- PHPRedisAdmin
- Adminer

### Editor

- VsCode
- NotePad++
- Neovim

### Dependency managers

- Composer
- PNPM
- Yarn
- Gradle
- Maven

### FTP client

- FileZilla

### Terminal enhancements

- Bat
- LazyDocker
- Delta
- LazyGit
- FZF
- Zoxide
- Starship
- FileFind(fd)
- RipGrep(rg)

## Debug

Run this command to debug the application:
```bash
dlv debug --headless --listen=:2345 --api-version=2 --log --log-output=rpc
```

Attach dap server to vscode/dap client :

Example of vscode configuration:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Attach to Headless Delve",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "${workspaceFolder}",
            "port": 2345,
            "host": "127.0.0.1",
            "showLog": true,
            "trace": "verbose"
        }
    ]
}
```

## Goal

1. Support Version switching for supported tools.
2. Interactive configuration to each tool.
3. Easy download package each tool.(If tools not in zip archive it will download .msi/.exe file and initiate installation)
4. Easily add tool to env

## Not Goal

1. Provide native GUI at the moment.
