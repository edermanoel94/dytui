package main

import (
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
  cat comma.csv              | c2j | jq        
  cat semicolon.csv          | c2j --delimiter ";" | jq
  cat csv_without_header.csv | c2j --no-header | jq`
)

// flags
var (
	fProfile string
	fVersion bool
	fHelp    bool
)

func main() {
	flag.StringVar(&fProfile, "profile", "", "select aws profile")
	flag.StringVar(&fProfile, "p", "", "select aws profile")
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
	case fProfile != "":
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "flag provided but not defined %s \n", flag.Args()[0])
		printUsage()
		os.Exit(-1)
	}

}

func printVersion() {
	fmt.Fprintf(os.Stdout, version)
}

func printUsage() {
	fmt.Fprintf(os.Stdout, usageString)
}
