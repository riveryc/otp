# OTP in command line

OTP is a simple cli tool that implemented in Golang to output the OTP code.

This project is using the lib from "github.com/creachadair/otp" as the HOTP (RFC 4226) and TOTP (RFC 6238) algorithms.

# Build

```bash
go build .
```

# Usage

## Direct output current OTP
```bash
otp <Activation Code>
```

## Pipe with openconnect VPN to automatic pass the MFA
```bash
# make sure to remove "<>" with following
echo "<your password>\n`otp <your activation code>`" | sudo openconnect -u <your username> https://<your vpn url with port> --passwd-on-stdin
```