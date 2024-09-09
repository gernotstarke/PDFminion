package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Config struct {
	SourceDir   string
	TargetDir   string
	showHelp    bool
	showVersion bool
	Force       bool
}

const (
	defaultSourceDir = "_pdfs"
	defaultTargetDir = "_target"
	Version          = "0.3.0"
)

const PageNrPrefix = ""
const ChapterPrefix = "Kap."
const ChapterPageSeparator = " - "

// BuildTime will be injected at build time
var BuildTime string

func New() *Config {
	return &Config{
		SourceDir: defaultSourceDir,
		TargetDir: defaultTargetDir,
		Force:     false,
	}
}

func (c *Config) ParseFlags() {
	flag.StringVar(&c.SourceDir, "s", defaultSourceDir, "Specify the source directory")
	flag.StringVar(&c.SourceDir, "source", defaultSourceDir, "Specify the source directory")
	flag.StringVar(&c.TargetDir, "t", defaultTargetDir, "Specify the target directory")
	flag.StringVar(&c.TargetDir, "target", defaultTargetDir, "Specify the target directory")
	flag.BoolVar(&c.Force, "force", false, "Skips check of empty target directory, forces overwrite of existing files")
	flag.BoolVar(&c.showHelp, "h", false, "Show help information")
	flag.BoolVar(&c.showHelp, "help", false, "Show help information")
	flag.BoolVar(&c.showVersion, "v", false, "Show version information")
	flag.BoolVar(&c.showVersion, "version", false, "Show version information")

	// Custom usage function
	flag.Usage = c.printHelp

	// Parse flags
	flag.Parse()

	// Check for unrecognized arguments or "help"
	if flag.NArg() > 0 {
		arg := flag.Arg(0)
		if arg == "help" || arg == "?" || strings.HasPrefix(arg, "-") {
			c.showHelp = true
		}
	}

	// Check for "--" prefixed long-form flags
	for i, arg := range os.Args {
		switch arg {
		case "--source":
			if i+1 < len(os.Args) {
				c.SourceDir = os.Args[i+1]
			}
		case "--target":
			if i+1 < len(os.Args) {
				c.TargetDir = os.Args[i+1]
			}
		case "--help":
			c.showHelp = true
		case "--version":
			c.showVersion = true
		}
	}
}

func (c *Config) Evaluate() error {
	if c.showHelp {
		c.printHelp()
		os.Exit(0)
	}

	if c.showVersion {
		c.printVersion()
		os.Exit(0)
	}

	return c.validate()
}

func (c *Config) printHelp() {
	fmt.Println("PDFMinion adds page numbers to existing PDF files.")
	fmt.Println("It will take all PDF files from the source directory and put the numbered copies into the target directory.")
	fmt.Println("Furthermore, it will ensure that every chapter (aka file) starts with an odd number")
	fmt.Println("by adding a single blank page to files with an un-even page count.")
	fmt.Println("When printed double-sided, every chapter will start on a right side with an odd pagenumber.")
	fmt.Println("\n\nUsage:")
	fmt.Printf("  -s, --source string\n\tSpecify the source directory (default \"%s\")\n", defaultSourceDir)
	fmt.Printf("  -t, --target string\n\tSpecify the target directory (default \"%s\")\n", defaultTargetDir)
	fmt.Printf("  --force string\n\tSkips check of empty target directory, forces overwrite of existing files (default false)\n")
	fmt.Println("  -h, --help, ?, -?, help\n\tShow this help message")
	fmt.Println("  -v, --version\n\tShow version information")
}

func (c *Config) printVersion() {
	fmt.Printf("PDFminion version %s\n", Version)
	if BuildTime != "" {
		t, err := time.Parse("2006 Jan 02 15:04", BuildTime)
		if err == nil {
			formattedBuildTime := t.Format("2006 Jan 02 15:04")
			parts := strings.Split(formattedBuildTime, " ")
			if len(parts) == 4 {
				day := parts[2]
				suffix := getSuffix(day)
				parts[2] = day + suffix
				parts[3] += "h"
				formattedBuildTime = strings.Join(parts, " ")
			}
			fmt.Printf("Build time: %s\n", formattedBuildTime)
		} else {
			fmt.Printf("Build time: %s\n", BuildTime)
		}
	} else {
		fmt.Println("Build time: Not available")
	}
}

func (c *Config) validate() error {
	if err := c.validateSourceDir(); err != nil {
		return err
	}
	return c.validateTargetDir()
}

func (c *Config) validateSourceDir() error {
	if _, err := os.Stat(c.SourceDir); os.IsNotExist(err) {
		return fmt.Errorf("source directory %s does not exist", c.SourceDir)
	}
	return nil
}

func (c *Config) validateTargetDir() error {
	if _, err := os.Stat(c.TargetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory %s does not exist. Creating it...\n", c.TargetDir)
		if err := os.MkdirAll(c.TargetDir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %s: %v", c.TargetDir, err)
		}
	} else if !c.Force {
		// Check if the directory is empty
		files, err := ioutil.ReadDir(c.TargetDir)
		if err != nil {
			return fmt.Errorf("error reading directory %s: %v", c.TargetDir, err)
		}

		if len(files) > 0 {
			return fmt.Errorf("target directory %s is not empty", c.TargetDir)
		}
	}
	return nil
}

func getSuffix(day string) string {
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
