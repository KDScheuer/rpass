# Work in Progress
This repo is a placeholder for an idea that I have to build a password generator. This would make my life a litte easier and be a good simple project to learn Go.
Below is the desired function of the executable.

# RPass
RPass is a lightweight, fast CLI utility written in Go for generating secure random passwords on demand.

## Why This Tool Exists
As an SRE or sysadmin, you often need to generate one-off passwords—for SSL certs, SFTP users, temporary credentials, or other short-lived tasks. Opening a full password manager every time feels excessive, especially when the password isn’t worth storing.

RPass was built to solve that frustration:
- No account required
- No bloat
- Just a fast, portable binary that does one thing well: generates secure passwords from the terminal.


## Arguments

| Flag |  Alias              |  Description                                                   |
|------|---------------------|-----------------------------------------------------------------|
|  -l  |  --length           |  Length of the password (default: 24)                           |
|  -s  |  --symbols          |  Include symbols in the password (default: true)                |
|  -u  |  --upper            |  Include uppercase characters (default: true)                   |
|  -x  |  --lower            |  Include lowercase characters (default: true)                   |
|  -n  |  --number           |  Include numbers (default: true)                                |
|  -S  |  --symbols-allowed  |  Restrict to specific symbols (e.g. -S "!@#")                   |
|  -p  |  --prompt           |  Interactive mode: prompts for options                          |
|  -h  |  --help             |  Show help message                                              |


## Examples
Generate a 24 charcter password with upper, lower, numbers, and symbols (default behavivior)
```bash
rpass
```
Generate a 19-character password with upper, lower, and numbers (no symbols)
```bash
rpass -l 19 -uxn
```
Generate a 24-character password with numbers only
```bash
rpass -n
```
Generate a password with upper, lower, numbers and only the symbols !#"
```bash
rpass -S !#"
```
Generate a 10-character password with numbers and symbols !#'
```bash
rpass -S !#' -l 10 -n
```