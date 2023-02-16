# keyfile
Keychain-based file encryption

## Install

```shell
go install github.com/chyroc/keyfile@latest
```

## Usage

### Encrypt File

```shell
keyfile -a `<account_name>` filepath > encrypt_filepath
```

### Decrypt File

```shell
keyfile -r -a `<account_name>` encrypt_filepath
```

### Help

```shell
NAME:
   keyfile - keyfile [--read] [--account] filepath

USAGE:
   keyfile [global options] command [command options] [arguments...]

DESCRIPTION:
   Keychain-based file encryption

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --read, -r                 read a file (default: false)
   --account value, -a value  account name [$KEYFILE_ACCOUNT]
   --help, -h                 show help
```
