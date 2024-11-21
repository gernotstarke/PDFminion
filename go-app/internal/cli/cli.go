package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Options struct {
	SourceDir string
	TargetDir string
	Force     bool
	Debug     bool
	Help      bool
	Version   bool
}

// FlagDef represents a flag definition with possible short and long forms
type FlagDef struct {
	Long    string
	Short   string
	Default interface{}
	Help    string
}

// Flag definitions
var flags = map[string]FlagDef{
	"source": {
		Long:    "source",
		Short:   "s",
		Default: defaultSourceDir,
		Help:    "Specify the source directory",
	},
	"target": {
		Long:    "target",
		Short:   "t",
		Default: defaultTargetDir,
		Help:    "Specify the target directory",
	},
	"force": {
		Long:    "force",
		Short:   "f",
		Default: false,
		Help:    "Forces overwrite of existing files",
	},
	"debug": {
		Long:    "debug",
		Short:   "d",
		Default: false,
		Help:    "Enable debug logging",
	},
	"help": {
		Long:    "help",
		Short:   "h",
		Default: false,
		Help:    "Show this help message",
	},
	"version": {
		Long:    "version",
		Short:   "v",
		Default: false,
		Help:    "Show version information",
	},
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

	// Register all flags
	for _, f := range flags {
		switch def := f.Default.(type) {
		case string:
			var value string
			flag.StringVar(&value, f.Long, def, f.Help)
			if f.Short != "" {
				flag.StringVar(&value, f.Short, def, f.Help)
			}
			switch f.Long {
			case "source":
				opts.SourceDir = value
			case "target":
				opts.TargetDir = value
			}
		case bool:
			var value bool
			flag.BoolVar(&value, f.Long, def, f.Help)
			if f.Short != "" {
				flag.BoolVar(&value, f.Short, def, f.Help)
			}
			switch f.Long {
			case "force":
				opts.Force = value
			case "debug":
				opts.Debug = value
			case "help":
				opts.Help = value
			case "version":
				opts.Version = value
			}
		}
	}

	flag.Usage = printHelp
	flag.Parse()

	// Handle help and version first, before any validation
	if opts.Help || (len(os.Args) > 1 && (os.Args[1] == "help" || os.Args[1] == "?")) {
		printHelp()
		os.Exit(0)
	}

	if opts.Version {
		printVersion()
		os.Exit(0)
	}

	// Only validate directories if we're actually going to process files
	if err := opts.validate(); err != nil {
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

Options:
`, progName)

	// Create a sorted list of flags for consistent output
	var flagNames []string
	for name := range flags {
		flagNames = append(flagNames, name)
	}
	sort.Strings(flagNames)

	// Print each flag with its short form
	for _, name := range flagNames {
		f := flags[name]
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
		return fmt.Errorf("source directory %q does not exist", o.SourceDir)
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
