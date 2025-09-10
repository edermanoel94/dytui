package main

import (
	"dytui/internal/controller"
	"dytui/internal/gui"
	"flag"
	"fmt"
	"os"
)

const (
	version     = "0.0.1\n"
	usageString = `Usage: dytui [flags]

Flags:
	-h, --help           print help information
	-v, --version        print version
	-p, --profile        select aws profile

Examples:
  cat semicolon.csv          | c2j --delimiter ";" | jq
  cat csv_without_header.csv | c2j --no-header | jq`
)

// flags
var (
	fVersion bool
	fHelp    bool
)

func main() {
	flag.BoolVar(&fVersion, "version", false, "print version")
	flag.BoolVar(&fVersion, "v", false, "print version")
	flag.BoolVar(&fHelp, "help", false, "print help")
	flag.BoolVar(&fHelp, "h", false, "print help")

	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usageString)
		os.Exit(0)
	}
	flag.Parse()

	run()
}

func run() {
	switch {
	case fHelp:
		printUsage()
		os.Exit(0)
	case fVersion:
		printVersion()
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "flag provided but not defined %s \n", flag.Args()[0])
		printUsage()
		os.Exit(-1)
	}

	ctrl, err := controller.New()

	if err != nil {
		fmt.Fprintf(os.Stdout, "controller couldn't initialize, error: %s\n", err.Error())
		os.Exit(-1)
	}

	gui := gui.New(ctrl)

	if err := gui.Run(); err != nil {
		fmt.Fprintf(os.Stdout, "gui couldn't initialize, error: %s\n", err.Error())
		os.Exit(-1)
	}
}

func printVersion() {
	fmt.Fprintf(os.Stdout, version)
}

func printUsage() {
	fmt.Fprintf(os.Stdout, usageString)
}
