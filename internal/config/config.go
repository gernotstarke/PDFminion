package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	SourceDir   string
	TargetDir   string
	showHelp    bool
	showVersion bool
}

const (
	defaultSourceDir = "_pdfs"
	defaultTargetDir = "_target"
)

// Version information
var (
	Version   = "development"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func New() *Config {
	return &Config{
		SourceDir: defaultSourceDir,
		TargetDir: defaultTargetDir,
	}
}

func (c *Config) ParseFlags() {
	var sourceFlag, targetFlag string

	flag.StringVar(&sourceFlag, "s", "", "Specify the source directory")
	flag.StringVar(&sourceFlag, "source", "", "Specify the source directory")
	flag.StringVar(&targetFlag, "t", "", "Specify the target directory")
	flag.StringVar(&targetFlag, "target", "", "Specify the target directory")
	flag.BoolVar(&c.showHelp, "h", false, "Show help information")
	flag.BoolVar(&c.showHelp, "help", false, "Show help information")
	flag.BoolVar(&c.showVersion, "v", false, "Show version information")
	flag.BoolVar(&c.showVersion, "version", false, "Show version information")

	flag.Parse()

	// Store the parsed values
	if sourceFlag != "" {
		c.SourceDir = sourceFlag
	}
	if targetFlag != "" {
		c.TargetDir = targetFlag
	}

	// Check for "help" or "?" as the first argument
	if flag.Arg(0) == "help" || flag.Arg(0) == "?" {
		c.showHelp = true
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
	fmt.Println("Usage of PDFminion:")
	fmt.Printf("  -s, -source string\n\tSpecify the source directory (default \"%s\")\n", defaultSourceDir)
	fmt.Printf("  -t, -target string\n\tSpecify the target directory (default \"%s\")\n", defaultTargetDir)
	fmt.Println("  -h, -help\n\tShow this help message")
	fmt.Println("  -v, -version\n\tShow version information")
	fmt.Println("  help, ?\n\tShow this help message")
}

func (c *Config) printVersion() {
	fmt.Printf("PDFminion version %s\n", Version)
	fmt.Printf("Build time: %s\n", BuildTime)
	fmt.Printf("Git commit: %s\n", GitCommit)
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
	} else {
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
