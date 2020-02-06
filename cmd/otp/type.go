package main

type Globals struct {
	Debug   bool        `short:"D" help:"Enable debug mode"`
	Verbose bool        `help:"Verbose mode." short:"V"`
	Version VersionFlag `help:"Print Version." short:"v"`
}

type VersionFlag string

type CLI struct {
	Globals

	Get GetCmd `cmd:"" help:"Get otp code."`
}

type GetCmd struct {
	Secret  string `help:"Secret code from provider, doesn't work with --profile or -p.'" short:"s"`
	Profile string `help:"Profile name, this has higher priority." short:"p"` //TODO: add profile support
}
