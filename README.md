<p align="center">

![Go Version](https://img.shields.io/badge/go-1.24+-00ADD8?logo=go&logoColor=white)
![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux-informational)
![Build](https://img.shields.io/badge/build-passing-brightgreen)
[![Release](https://img.shields.io/github/v/release/KDScheuer/rpass)](https://github.com/KDScheuer/rpass/releases)

</p>

# RPass
RPass is a lightweight, fast CLI utility written in Go for generating secure random passwords on demand.

- [Overview](#why-this-tool-exists)
- [Examples](#examples)
- [Arguments](#arguments)
- [Install](#install-instructions)
- [License](#license)

## Why This Tool Exists
As an SRE or sysadmin, you often need to generate one-off passwords—for SSL certs, SFTP users, temporary credentials, or other short-lived tasks. Opening a full password manager every time feels excessive, especially when the password isn’t worth storing.

RPass was built to solve that frustration:
- Quick and Secure passwords from the terminal
- No bloat or loading times
- Just a fast, portable binary that does one thing well: generates secure passwords from the terminal.

## Examples
Default settings (24 Characters, Min 2 of each type, Upper, Lower, Numbers and Symbols)
```bash
rpass
```
Generate 16 character password without symbols
```bash
rpass -uxn -l 16
```
Generate a 32 character password with only Upper and Lower
```bash
rpass -uxl 32
```
Generate a password with only `!#` as symbols with Upper, Lower, and Numbers
```bash
rpass -s "!#"
```
Generate a 32 character password with a minimum of 6 of each type of character
```bash 
rpass -t 6 -l 32
```
Display Version
```bash
rpass --version
```

## Arguments
| Flag |  Description                                                    |
|------|-----------------------------------------------------------------|
|  -l  |  Length of the password (default: 24)                           |
|  -s  |  Include symbols in the password. Can be passed alone which will use `!@#$%^&*()` or with a string of symbols to select from i.e. `-s "!#"`                |
|  -u  |  Include uppercase characters (default: true)                   |
|  -x  |  Include lowercase characters (default: true)                   |
|  -n  |  Include numbers (default: true)                                |                 
|  -t  |  Minimum occurrences of each type selected (default: 2)         |
|  -h  |  Show help message                                              |
|  -version  |  Show version                                             |

## Install Instructions
### Option 1: Download prebuilt executable (Windows, Linux)
1. Go to Releases Page
2. Download appropriate binary
    - `rpass.exe` for windows
    - `rpass` for linux
3. (If Linux) Make executable & move to PATH
```bash
mv rpass /usr/local/bin/rpass
chmod +x /usr/local/bin/rpass
```

### Option 2: Build Yourself (Requires Go)
```bash
git clone https://github.com/KDScheuer/rpass.git
cd rpass
go build -o rpass
```
The `rpass` binart will be created in current directory and will need to be moved to PATH

### Option 3: Run without Building (Requires Go)
```bash
git clone https://github.com/KDScheuer/rpass.git
cd rpass
go run . #Any Arguments
```

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.