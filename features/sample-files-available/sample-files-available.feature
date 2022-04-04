@developer
Feature: Sample files are available

  Note: The name and directory path of the sample files
  are configured in the package `sample_pdfs`

  Background:
    Given The directory "sample-files-for-testing"

  Scenario: The directory containing sample files is available
    When Existence of samplePDFDir is checked
    Then The directory containing sample files is available


  Scenario: A directory without PDF files exists
    Given Directory "EmptyDir"
    When Number of PDF files is counted
    Then 0 is returned

  Scenario: A directory with 1 PDF file exists
    Given Directory "OnePDF"
    When Number of PDF files is counted
    Then 1 is returned

  Scenario: A directory with 2 PDF files exists
    Given Directory "TwoPDFs"
    When Number of PDF files is counted
    Then 2 is returned

