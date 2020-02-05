package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/creachadair/otp"
	"time"
)

func Codes(activationCode string) (token string, _ error) {
	var conf otp.Config
	if err := conf.ParseKey(activationCode); err != nil {
		return "", err
	}
	token = conf.TOTP()
	return
}

type Globals struct {
	Debug   bool        `short:"D" help:"Enable debug mode"`
	Verbose bool        `help:"Verbose mode." short:"V"`
	Version VersionFlag `help:"Print Version." short:"v"`
}

type VersionFlag string

//noinspection ALL
type CLI struct {
	Globals

	Get GetCmd `cmd:"" help:"Get otp code."`
}

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type GetCmd struct {
	Secret  string `help:"Secret code from provider, doesn't work with --profile or -p.'" short:"s"`
	Profile string `help:"Profile name, this has higher priority." short:"p"` //TODO: add profile support
}

func (g *GetCmd) Run(ctx *Globals) error {
	secret := g.Secret
	if ctx.Debug {
		fmt.Println("Current Time is", time.Now())
	}
	token, e := Codes(secret)
	if e != nil {
		panic(e)
	}
	fmt.Println(token)
	return nil
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.0.1"),
		},
	}
	ctx := kong.Parse(&cli,
		kong.Name("otp"),
		kong.Description("An OTP command tool"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"version": string(cli.Globals.Version),
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
