# OTP in command line

OTP is a simple cli tool that implemented in Golang to output the OTP code.

This project is using the lib from "github.com/creachadair/otp" as the HOTP (RFC 4226) and TOTP (RFC 6238) algorithms.

# Build

```bash
mkdir bin
go build -o bin/otp ./cmd/otp
```

# Usage

## OTP help
```bash
cd bin
./otp --help
Usage: otp <command>

An OTP command tool

Flags:
      --help       Show context-sensitive help.
  -D, --debug      Enable debug mode
  -V, --verbose    Verbose mode.
  -v, --version    Print Version.

Commands:
  get    Get otp code.

Run "otp <command> --help" for more information on a command.
```

## Get an otp code from secret

```bash
./otp get --secret <secret>
```

## Pipe with openconnect VPN to automatic pass the MFA
```bash
# make sure to remove "<>" with following
echo "<your password>\n`otp get -s <secret>`" | sudo openconnect -u <your username> https://<your vpn url with port> --passwd-on-stdin
```