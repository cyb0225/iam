/**
@author: yeebing
@date: 2022/9/25
**/

package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type App struct {
	name        string // the command server name
	basename    string // the name of binary file
	description string
	cmd         *cobra.Command // root command
	args        cobra.PositionalArgs
	noVersion   bool // if the app have version command
	noConfig    bool // if the app have a config file
	opt         CliOption
}

type Option func(app *App)

// NewApp init a command server frame.
func NewApp(name string, basename string, opts ...Option) *App {
	app := &App{
		name:     name,
		basename: basename,
	}

	for _, o := range opts {
		o(app)
	}

	buildCommand(app)

	return app
}

// Run set up app command server.
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// WithDescription set app's description.
func WithDescription(desc string) Option {
	return func(app *App) {
		app.description = desc
	}
}

// WithArgs set app's args operator func.
func WithArgs(args cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = args
	}
}

// WithOption set app's option.
func WithOption(opt CliOption) Option {
	return func(app *App) {
		app.opt = opt
	}
}

// WithNoVersion set app not to set version command.
func WithNoVersion() Option {
	return func(app *App) {
		app.noVersion = true
	}
}

// WithNoConfig set app not to set config command.
func WithNoConfig() Option {
	return func(app *App) {
		app.noConfig = true
	}
}

func buildCommand(a *App) {
	cmd := cobra.Command{
		Use:   a.basename, // the binary (exec) filename
		Short: a.name,
		Long:  a.description,
		Args:  a.args,
	}

	a.cmd = &cmd

	// set global flags

	if !a.noConfig {
	}

	if !a.noVersion {
	}
}
