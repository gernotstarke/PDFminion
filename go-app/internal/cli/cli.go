package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Options struct {
	SourceDir string
	TargetDir string
	Force     bool
}

const (
	defaultSourceDir     = "_pdfs"
	defaultTargetDir     = "_target"
	PageNrPrefix         = ""
	ChapterPrefix        = "Kap."
	ChapterPageSeparator = " - "
)

// Version information - injected at build time
var (
	buildTime    string
	hostPlatform string
	appVersion   string
)

func SetAppVersion(version string) {
	appVersion = version
}

func ParseOptions() (*Options, error) {
	opts := &Options{
		SourceDir: defaultSourceDir,
		TargetDir: defaultTargetDir,
	}

	// Define flags once
	flag.StringVar(&opts.SourceDir, "source", defaultSourceDir, "Specify the source directory")
	flag.StringVar(&opts.TargetDir, "target", defaultTargetDir, "Specify the target directory")
	flag.BoolVar(&opts.Force, "force", false, "Forces overwrite of existing files")

	version := flag.Bool("version", false, "Show version information")
	help := flag.Bool("help", false, "Show help information")

	// Add short aliases
	flag.StringVar(&opts.SourceDir, "s", defaultSourceDir, "")
	flag.StringVar(&opts.TargetDir, "t", defaultTargetDir, "")
	flag.BoolVar(help, "h", false, "")
	flag.BoolVar(version, "v", false, "")

	flag.Usage = printHelp
	flag.Parse()

	// Handle help/version first
	switch {
	case *help:
		printHelp()
		os.Exit(0)
	case *version:
		printVersion()
		os.Exit(0)
	}

	// Validate the options
	if err := opts.validate(); err != nil {
		return nil, err
	}

	return opts, nil
}

func printHelp() {
	fmt.Printf(`PDFMinion adds page numbers to existing PDF files.
It will take all PDF files from the source directory and put the numbered copies into the target directory.
Furthermore, it will ensure that every chapter (aka file) starts with an odd number
by adding a single blank page to files with an un-even page count.
When printed double-sided, every chapter will start on a right side with an odd pagenumber.

Usage:
  -s, --source string
        Specify the source directory (default "%s")
  -t, --target string
        Specify the target directory (default "%s")
  --force
        Forces overwrite of existing files
  -h, --help
        Show this help message
  -v, --version
        Show version information
`, defaultSourceDir, defaultTargetDir)
}

func printVersion() {
	fmt.Printf("PDFminion version %s\n", appVersion)
	fmt.Printf("Built on: %s\n", hostPlatform)
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

func (o *Options) validate() error {
	if err := o.validateSourceDir(); err != nil {
		return err
	}
	return o.validateTargetDir()
}
func (o *Options) validateSourceDir() error {
	if _, err := os.Stat(o.SourceDir); os.IsNotExist(err) {
		return fmt.Errorf("Source directory '%s' does not exist", o.SourceDir)
	}
	return nil
}

func (o *Options) validateTargetDir() error {
	if _, err := os.Stat(o.TargetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory '%s' does not exist. Creating it...\n", o.TargetDir)
		if err := os.MkdirAll(o.TargetDir, os.ModePerm); err != nil {
			return fmt.Errorf("Failed to create directory '%s': %v", o.TargetDir, err)
		}
		return nil
	}

	if o.Force {
		return nil
	}

	files, err := ioutil.ReadDir(o.TargetDir)
	if err != nil {
		return fmt.Errorf("Cannot read directory '%s': %v", o.TargetDir, err)
	}

	if len(files) > 0 {
		return fmt.Errorf("Target directory '%s' is not empty. Use --force to override", o.TargetDir)
	}

	return nil
}
