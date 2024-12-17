# 8. layered-configuration-approach

Date: 2024-12-16

## Status

Accepted

## Context

PDFminion requires a mechanism to parse and process command-line arguments for a variety of commands and options, such as `--source`, `--target`, `--language`, etc.

Users shall be able to:

* save their preferred configuration in configuration files
* overwrite their saved configs with command-line flags
* rely on default settings for minimal effort

This is called a _layered configuration approach_:

## Decision


### Define the Configuration Workflow:

* Load default values first.
* Override with values from the configuration file (if provided via --file or exists as pdfminion.cfg).
* Override further with command-line flags.
* Library Choice: Use Viper for configuration file handling and defaults.
* Use the Cobra package for command-line parsing. Cobra integrates well with Viper for CLI apps.

### Priority Order:

Command-line flags > Config file > Defaults.

### Implementation Outline:

* Parse CLI flags (use Cobra).
* Determine the config file location (from --file or pdfminion.cfg).
* Load the config file (if it exists).
* Merge all sources into a unified configuration structure.




## Consequences

We test this slightly more complex configuration approach on a branch(#15-layered-configuration)

 