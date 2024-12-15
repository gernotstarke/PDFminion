package cli

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"pdfminion/internal/domain"
	"sort"
	"strings"
	"time"
)

// Version information - injected at build time
var (
	buildTime     string
	buildPlatform string
)

func ParseOptions() (*domain.MinionConfig, error) {
	opts := &domain.MinionConfig{
		SourceDir: domain.DefaultSourceDir,
		TargetDir: domain.DefaultTargetDir,
	}

	// Register all flags
	flag.StringVar(&opts.SourceDir, "source", domain.DefaultSourceDir, domain.Flags["source"].Help)
	flag.StringVar(&opts.SourceDir, "s", domain.DefaultSourceDir, "")
	flag.StringVar(&opts.TargetDir, "target", domain.DefaultTargetDir, domain.Flags["target"].Help)
	flag.StringVar(&opts.TargetDir, "t", domain.DefaultTargetDir, "")
	flag.BoolVar(&opts.Force, "force", false, domain.Flags["force"].Help)
	flag.BoolVar(&opts.Force, "f", false, "")
	flag.BoolVar(&opts.Debug, "debug", false, domain.Flags["debug"].Help)
	flag.BoolVar(&opts.Debug, "d", false, "")
	flag.BoolVar(&opts.Help, "help", false, domain.Flags["help"].Help)
	flag.BoolVar(&opts.Help, "h", false, "")
	flag.BoolVar(&opts.Version, "version", false, domain.Flags["version"].Help)
	flag.BoolVar(&opts.Version, "v", false, "")

	flag.BoolVar(&opts.ListLanguages, "list-languages", false, domain.Flags["list-languages"].Help)
	flag.BoolVar(&opts.ListLanguages, "ll", false, "")

	flag.Usage = printHelp
	flag.Parse()

	// Handle help and version first
	if opts.Help || flag.Arg(0) == "help" || flag.Arg(0) == "?" {
		printHelp()
		os.Exit(0)
	}

	if opts.Version {
		printVersion()
		os.Exit(0)
	}

	if opts.ListLanguages {
		log.Debug().Msg("List available languages")
		domain.PrintLanguages()
		os.Exit(0)
	}

	// Only validate directories if we're actually going to process files
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return opts, nil
}

func printHelp() {
	progName := filepath.Base(os.Args[0])

	fmt.Printf(`PDFMinion adds page numbers to existing PDF files.
It will take all PDF files from the source directory and put the numbered copies 
into the target directory. Every chapter (aka file) starts with an odd number
by adding a single blank page to files with an un-even page count.
When printed double-sided, every chapter will start on a right side.

Usage: %s [options]

MinionConfig:
`, progName)

	// Create a sorted list of flags for consistent output
	var flagNames []string
	for name := range domain.Flags {
		flagNames = append(flagNames, name)
	}
	sort.Strings(flagNames)

	// Print each flag with its short form
	for _, name := range flagNames {
		f := domain.Flags[name]
		switch def := f.Default.(type) {
		case string:
			fmt.Printf("  -%s, --%-12s %s (default: %q)\n",
				f.Short, f.Long, f.Help, def)
		case bool:
			fmt.Printf("  -%s, --%-12s %s\n",
				f.Short, f.Long, f.Help)
		}
	}
}

func printVersion() {
	fmt.Printf("PDFminion version %s\n", domain.AppVersion)
	fmt.Printf("Built on: %s\n", buildPlatform)
	if buildTime != "" {
		t, err := time.Parse("2006 Jan 02 15:04", buildTime)
		if err == nil {
			formattedTime := formatBuildTime(t)
			fmt.Printf("Build time: %s\n", formattedTime)
		} else {
			fmt.Printf("Build time: %s\n", buildTime)
		}
	} else {
		fmt.Println("Build time: Not available")
	}
}

func formatBuildTime(t time.Time) string {
	formatted := t.Format("2006 Jan 02 15:04")
	parts := strings.Split(formatted, " ")
	if len(parts) == 4 {
		day := parts[2]
		parts[2] = day + getDaySuffix(day)
		parts[3] += "h"
		return strings.Join(parts, " ")
	}
	return formatted
}

func getDaySuffix(day string) string {
	switch day {
	case "01", "21", "31":
		return "st"
	case "02", "22":
		return "nd"
	case "03", "23":
		return "rd"
	default:
		return "th"
	}
}
