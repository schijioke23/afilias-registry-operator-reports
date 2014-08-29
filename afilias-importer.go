//
// this is the main CLI
//
// Copyright 2014 TLD dotHIV Registry GmbH | http://tld.hiv/
//
// @author Markus Tacker <m@click4life.hiv>
//
package main

import (
	"fmt"
	"github.com/dothiv/afilias-registry-operator-reports/cli"
	"github.com/dothiv/afilias-registry-operator-reports/command/importer/registrations"
	"github.com/wsxiaoys/terminal/color"
	"os"
)

func error(msg string) {
	color.Fprintln(os.Stderr, "@{!r}ERROR @{|}"+msg)
}

func Help() {
	cli.HelpBanner("@{g}<command>@{|}")
	color.Fprintln(os.Stdout, fmt.Sprintf("  @{g}command@{|} may be         help | %s\n", registrations.NAME))
	color.Fprintln(os.Stdout, fmt.Sprintf("Use %s help <command> to get help for a command", os.Args[0]))
}

func main() {
	if len(os.Args) == 1 {
		Help()
		error(fmt.Sprintf("(%s) too few arguments", os.Args[0]))
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help":
		if len(os.Args) == 2 {
			Help()
		} else if len(os.Args) > 3 {
			Help()
			error(fmt.Sprintf("(%s) too many arguments", os.Args[0]))
			os.Exit(1)
		} else {
			switch os.Args[2] {
			case registrations.NAME:
				registrations.Help()
			}
		}

		os.Exit(0)
	case registrations.NAME:
		c, err := registrations.NewConfig()
		if err != nil {
			error(err.Error())
			registrations.Help()
			os.Exit(1)
		}
		err = registrations.Import(c)
		if err != nil {
			error(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
}