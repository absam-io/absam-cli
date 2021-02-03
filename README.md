Absam CLI
=========
`absam-cli` é uma aplicação de linha de comando (CLI) para a API da Absam.

```
absam-cli is a command-line interface for the Absam API

Usage:
  absam-cli [command]

Available Commands:
  firewall    Manage a server's firewall
  help        Help about any command
  init        Setup absam-cli
  server      Manage servers

Flags:
  -h, --help   help for absam-cli

Use "absam-cli [command] --help" for more information about a command.

```

- [Absam CLI](#absam-cli)
  - [Instalação](#instalação)
  - [Uso](#uso)
    - [init](#init)
  - [Desinstalar](#desinstalar)


## Instalação
Antes de rodar os seguintes comandos, confira se o `Go` (1.14 ou maior) está instalado na máquina.
Você pode instalar o Go através do gerenciador de pacotes do seu sistema operacional ou baixando diretamente
do [site oficial](https://golang.org/dl/).

Se o Go já estiver instalado, existem duas maneiras de instalar a CLI em sua máquina.

Utilizando `make`:

```
# make
# make install
```

Fazendo manualmente com o `Go`:
```
$ go build -o absam-cli main.go
# mv absam-cli /usr/bin
```

## Uso
Para utilizar a `absam-cli` é necessário gerar os tokens de nossa API no [painel da Absam](https://dashboard.absam.io/api).

### init
Após gerar os tokens, rode o seguinte comando em seu terminal passando cada token para suas flags respectivas.

```
$ absam-cli init --access-token=token_string --secret-access-token=secret_token_string
```

Todos os detalhes de cada comando podem ser encontrados utilizando a flag `--help` para cada comando.

## Desinstalar
Para desinstalar a `absam-cli` basta rodar o seguinte comando na pasta do repositório:

```
# make uninstall
```
