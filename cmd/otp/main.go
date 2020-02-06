package main

import (
	"github.com/alecthomas/kong"
)

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
