package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"golang.org/x/text/language"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"pdfminion/internal/domain"
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "pdfminion",
	Short: "PDFMinion adds page numbers to PDF files with custom options.",
	Long:  "PDFMinion is a CLI tool to add page numbers to existing PDF files. It provides customizable options like chapter numbers, running header,  merging and more.",
	RunE:  runNormalProcessing, // This runs when no subcommand is given
}

var minionConfig *domain.MinionConfig // Config variable to be populated

// LoadConfig initializes the configuration
func LoadConfig() (*domain.MinionConfig, error) {
	// Define flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to the configuration file")
	rootCmd.PersistentFlags().Bool("settings", false, "Display the final configuration settings")
	rootCmd.Flags().StringP("source", "s", "", "Source directory for input files")
	rootCmd.Flags().StringP("target", "t", "", "Target directory for output files")
	rootCmd.Flags().String("page-prefix", "", "Prefix for page numbers")
	rootCmd.Flags().Bool("evenify", false, "Ensure even number of pages")
	rootCmd.Flags().Bool("debug", false, "Enable debug mode")

	// Bind flags to Viper
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("settings", rootCmd.PersistentFlags().Lookup("settings"))
	viper.BindPFlag("source", rootCmd.Flags().Lookup("source"))
	viper.BindPFlag("target", rootCmd.Flags().Lookup("target"))
	viper.BindPFlag("page-prefix", rootCmd.Flags().Lookup("page-prefix"))
	viper.BindPFlag("evenify", rootCmd.Flags().Lookup("evenify"))
	viper.BindPFlag("debug", rootCmd.Flags().Lookup("debug"))

	// Hook into PreRun for early config loading and merging
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Skip config loading for subcommands
		if cmd.CalledAs() != "" && cmd.CalledAs() != "pdfminion" {
			return nil
		}
		return loadAndMergeConfig()
	}

	// Add subcommands
	addCommands()

	// Set the main logic
	rootCmd.RunE = runNormalProcessing

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		return nil, err
	}

	return minionConfig, nil
}

func loadAndMergeConfig() error {
	// Step 1: Initialize defaults
	minionConfig = initializeDefaults()
	log.Debug().Msg("Initialized defaults.")

	// Step 2: Check for --config flag
	configFile := viper.GetString("config")
	if configFile != "" {
		log.Debug().Str("filename", configFile).Msg("Detected --config flag.")

		// Search in the home directory first
		homeDir, err := os.UserHomeDir()
		if err == nil {
			homeConfigPath := fmt.Sprintf("%s/%s", homeDir, configFile)
			log.Debug().Str("path", homeConfigPath).Msg("Searching for config file in the home directory.")
			if _, err := os.Stat(homeConfigPath); err == nil {
				viper.SetConfigFile(homeConfigPath)
				if err := viper.MergeInConfig(); err == nil {
					log.Info().Str("path", homeConfigPath).Msg("Configuration file successfully read from home directory.")
				}
			}
		}

		// Search in the current directory
		localConfigPath := configFile
		log.Debug().Str("path", localConfigPath).Msg("Searching for config file in the current directory.")
		if _, err := os.Stat(localConfigPath); err == nil {
			viper.SetConfigFile(localConfigPath)
			if err := viper.MergeInConfig(); err == nil {
				log.Info().Str("path", localConfigPath).Msg("Configuration file successfully read from current directory.")
			}
		}
	}

	// Step 3: Merge Viper into minionConfig
	if err := viper.Unmarshal(minionConfig); err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	log.Debug().Msg("Merged configuration with CLI flags and defaults.")
	return nil
}

func addCommands() {
	// Subcommand: version
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show the application version",
		Run: func(cmd *cobra.Command, args []string) {
			domain.PrintVersion()
			os.Exit(0)
		},
	})

	// Subcommand: credits
	rootCmd.AddCommand(&cobra.Command{
		Use:   "credits",
		Short: "Show credits for PDFMinion",
		Run: func(cmd *cobra.Command, args []string) {
			domain.GiveCredits()
			os.Exit(0)
		},
	})

	// Subcommand: list-languages
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list-languages",
		Short: "List supported languages",
		Run: func(cmd *cobra.Command, args []string) {
			domain.PrintLanguages()
			os.Exit(0)
		},
	})
	// and its short form ll
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ll",
		Short: "List supported languages",
		Run: func(cmd *cobra.Command, args []string) {
			domain.PrintLanguages()
			os.Exit(0)
		},
	})
}

func runNormalProcessing(cmd *cobra.Command, args []string) error {
	// Step 4: Check for --settings flag to print the final config
	if viper.GetBool("settings") {
		log.Info().Msg("Displaying final configuration settings as requested.")
		domain.PrintFinalConfiguration(*minionConfig)
		fmt.Println("Final Configuration:")
		os.Exit(0)
	}

	// Step 5: Setup logging
	setupLogging(minionConfig.Debug)
	log.Info().Msg("Configuration setup complete.")

	// Step 6: Main processing logic
	fmt.Println("Processing PDFs with the following configuration:")
	fmt.Printf("%+v\n", *minionConfig)

	return nil
}

// SetupLogging initializes logging with debug or production settings.
func setupLogging(debug bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05",
		}).Level(zerolog.DebugLevel)
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05",
		}).Level(zerolog.InfoLevel)
	}
}

// Initialize default values for the configuration
func initializeDefaults() *domain.MinionConfig {
	return &domain.MinionConfig{
		ConfigFileName:       "pdfminion.yaml",
		Language:             language.English,
		SourceDir:            "./source",
		TargetDir:            "./target",
		Evenify:              true,
		Merge:                false,
		MergeFileName:        "merged.pdf",
		RunningHeader:        "",
		ChapterPrefix:        "",
		Separator:            "",
		PagePrefix:           "",
		TotalPageCountPrefix: "",
		BlankPageText:        "",
		Settings:             false,
		Debug:                false,
	}
}
