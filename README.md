# duckdns

```
 ____                _     ____   _   _  ____
|  _ \  _   _   ___ | | __|  _ \ | \ | |/ ___|
| | | || | | | / __|| |/ /| | | ||  \| |\___ \
| |_| || |_| || (__ |   < | |_| || |\  | ___) |
|____/  \__,_| \___||_|\_\|____/ |_| \_||____/

DuckDNS is a command line tool (cli) for the popular DuckDNS.org service. https://www.duckdns.org/

Usage:
  duckdns [command]

Available Commands:
  help        Help about any command
  update      Update duckdns domain(s). See: `duckdns update --help` for further details

Flags:
  -h, --help      help for duckdns
      --version   Displays the current version of this tool

Use "duckdns [command] --help" for more information about a command.

```

#

# Building and running duckdns #

```
go get github.com/enda-mullally/duckdns
        
cd %GOPATH%/src/github.com/enda-mullally/duckdns
    
go build
```

# Help

```
duckdns --help
```

# Usage

## Updating your DuckDNS.org domains via the update command using the flags --domains, --token ##
<br />

```
duckdns update --domains "a9275" --token "496462d1-0cf8-4292-aa00-************"
```

### If you need to use a specific ip, provide the --ip flag
<br />

```
duckdns update --domains "a9275" --token "496462d1-0cf8-4292-aa00-************" --ip 10.0.0.1
```

## Updating your DuckDNS.org domains via environment flags ##

### If you have set your env vars DUCKDNS_DOMAINS, DUCKDNS_TOKEN and optionally DUCKDNS_IP you can run duckdns in *--env mode* ###
<br />

```
duckdns update --env
```
<br />

![Screenshot](/screenshots/Scr01.png)