Feature: Command-line usage

  Scenario: A user wants to numberize PDFs via the command line
  Given A source directory
    And a different target directory
   And the source directory contains one PDF named "cli.pdf"
  When the user calls the application without further parameters
  Then The target directory will contain one PDF named "cli.pdf"