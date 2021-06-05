Feature: Convenience functions for PDF files
  In order to handle PDF files
  as a developer
  I need convenience functions for various file-related operations.

  Scenario: Recognize directory without PDF files
    Given An empty directory is selected as source-dir
    When PDF files are counted
    Then the number of PDF files should be 0

  Scenario: Recognize a single PDF file in directory
    Given A directory containing a single file
    And this file has type PDF
    When PDF files are counted
    Then the number of PDF files should be 1

