package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/riveryc/otp/pkg/codes"
	"time"
)

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

func (g *GetCmd) Run(ctx *Globals) error {
	secret := g.Secret
	if ctx.Debug {
		fmt.Println("Current Time is", time.Now())
	}
	token, e := codes.Codes(secret)
	if e != nil {
		panic(e)
	}
	fmt.Println(token)
	return nil
}
