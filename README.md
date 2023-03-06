# keyfile
Keychain-based file encryption

## Install

```shell
CGO_ENABLED=1 go install github.com/chyroc/keyfile@latest
```

## Usage

### Encrypt File

```shell
keyfile encrypt --account `<account_name>` --file filepath > encrypt_filepath
```

### Decrypt File

```shell
keyfile decrypt --account `<account_name>` --file encrypt_filepath
```

### Decrypt File and Re-Encrypt to File

```shell
keyfile decrypt --account `<account_name>` --file encrypt_filepath --editor vim
```

### Get Keychain Secret

```shell
keyfile get-secret --account `<account_name>`
```

### Set Keychain Secret

```shell
keyfile set-secret --account `<account_name>` --secret `<secret>`
```

### Del Keychain Secret

```shell
keyfile del-secret --account `<account_name>`
```

### Help

```shell
NAME:
   keyfile - A new cli application

USAGE:
   keyfile [global options] command [command options] [arguments...]

DESCRIPTION:
   Keychain-based file encryption

COMMANDS:
   decrypt, dec  decrypt content from file
   encrypt, enc  encrypt content from file
   get-secret    get secret from keychain
   set-secret    set secret to keychain
   del-secret    del secret from keychain
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
